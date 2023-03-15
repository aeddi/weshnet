package main

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
)

type Tcase struct {
	PeersAmount *int
	TimeoutMS   int
	Steps       *[]Step

	json []byte
	hash string
}

type Step struct {
	MsgAmount  *int
	TimeoutMS  int
	OutOfOrder bool
	Peers      *[]int
}

func newCase(path string) Tcase {
	c := Tcase{}
	testFile, err := os.Open(path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: can't open test file at path: %s\n", path)
		os.Exit(2)
	}
	defer testFile.Close()

	jsonBytes, _ := ioutil.ReadAll(testFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: can't read test file at path: %s\n", path)
		os.Exit(2)
	}

	err = json.Unmarshal(jsonBytes, &c)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: can't unmarshal json of test file at path: %s\n", path)
		fmt.Fprintf(os.Stderr, "%s\n", err.Error())
		os.Exit(2)
	}

	if c.PeersAmount == nil {
		fmt.Fprintf(os.Stderr, "Error: missing required field 'PeersAmount' in test file at path: %s\n", path)
		os.Exit(2)
	}
	if c.TimeoutMS < 0 {
		fmt.Fprintf(os.Stderr, "Error: wrong value 'TimeoutMS < 0' in test file at path: %s\n", path)
		os.Exit(2)
	}

	for i, step := range *c.Steps {
		if step.TimeoutMS < 0 {
			fmt.Fprintf(os.Stderr, "Error: wrong value 'TimeoutMS < 0' in step %d of test file at path: %s\n", i, path)
		} else if step.MsgAmount == nil {
			fmt.Fprintf(os.Stderr, "Error: missing required field 'MsgAmount' in step %d of test file at path: %s\n", i, path)
		} else if *step.MsgAmount < 0 {
			fmt.Fprintf(os.Stderr, "Error: wrong value 'MsgAmount < 0' in step %d of test file at path: %s\n", i, path)
		} else if step.Peers == nil {
			fmt.Fprintf(os.Stderr, "Error: missing required field 'Peers' in step %d of test file at path: %s\n", i, path)
		} else if len(*step.Peers) == 0 {
			fmt.Fprintf(os.Stderr, "Error: wrong value 'Peers == []' in step %d of test file at path: %s\n", i, path)
		} else if len(*step.Peers) > 0 {
			prev := -1
			sorted := make([]int, len(*step.Peers))
			copy(sorted, *step.Peers)
			sort.Ints(sorted)
			for _, peer := range sorted {
				if prev != -1 && peer == prev {
					fmt.Fprintf(os.Stderr, "Error: duplicated peer index '%d' in 'Peers' field in step %d of test file at path: %s\n", peer, i, path)
				} else if peer < 0 || peer >= *c.PeersAmount {
					fmt.Fprintf(os.Stderr, "Error: peer index '%d' is out of bound (0 to %d) in 'Peers' field in step %d of test file at path: %s\n", peer, *c.PeersAmount-1, i, path)
				} else {
					continue
				}
				os.Exit(2)
			}
			continue
		} else {
			continue
		}
		os.Exit(2)
	}

	c.json = jsonBytes
	c.hash = fmt.Sprintf("%x", md5.Sum(jsonBytes))

	return c
}
