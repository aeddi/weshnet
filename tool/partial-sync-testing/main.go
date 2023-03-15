package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"go.uber.org/zap"
)

type flags struct {
	logLevel  *string
	storeDir  *string
	outputDir *string
	testFiles []string
}

func parseFlags() *flags {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [options] test1 [test2 test3 ... testN]\n\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "Options:\n")
		flag.PrintDefaults()
	}

	cwd, err := os.Getwd()
	if err != nil {
		panic(fmt.Sprintf("can't get current working directory: %s", err.Error()))
	}

	f := flags{}
	f.logLevel = flag.String("l", "info", "set logging level among : debug, info, warn or error")
	f.storeDir = flag.String("s", filepath.Join(cwd, "store"), "set store directory path")
	f.outputDir = flag.String("d", filepath.Join(cwd, "output"), "set output directory path")
	flag.Parse()
	f.testFiles = flag.Args()

	flagError := func(err string) {
		fmt.Printf("Flag error: %s\n\n", err)
		flag.Usage()
		os.Exit(1)
	}

	if *f.logLevel != "debug" && *f.logLevel != "info" && *f.logLevel != "warn" && *f.logLevel != "error" {
		flagError(fmt.Sprintf("wrong log level: %s", *f.logLevel))
	} else if len(f.testFiles) < 1 {
		flagError("at least one test file is required")
	}

	return &f
}

func main() {
	f := parseFlags()

	logger := newLogger(*f.logLevel)
	defer logger.Sync()

	// First check if one of the test files is invalid
	testCases := make([]Tcase, len(f.testFiles))
	for i, testFile := range f.testFiles {
		testCases[i] = newCase(testFile)
	}

	// Create output and store dirs
	runDir := fmt.Sprintf("Run_%d", time.Now().Unix())
	fullOutputPath := filepath.Join(*f.outputDir, runDir)
	if err := os.MkdirAll(fullOutputPath, 0777); err != nil {
		panic(fmt.Sprintf("Error: can't create run output directory: %s", fullOutputPath))
	}
	fullStorePath := filepath.Join(*f.storeDir, runDir)
	if err := os.MkdirAll(fullStorePath, 0777); err != nil {
		panic(fmt.Sprintf("Error: can't create run output directory: %s", fullStorePath))
	}

	// Run the actual tests
	for i, testCase := range testCases {
		ctx := context.Background()
		if testCase.TimeoutMS > 0 {
			ctx, _ = context.WithTimeout(context.Background(), time.Duration(testCase.TimeoutMS)*time.Millisecond)
		}

		caseRef := fmt.Sprintf("TestCase%d_%s", i, testCase.hash)
		caseStorePath := filepath.Join(fullStorePath, caseRef)
		if err := os.MkdirAll(caseStorePath, 0777); err != nil {
			panic(fmt.Sprintf("Error: can't create case output directory: %s", caseStorePath))
		}
		g := createNewGroup(ctx, *testCase.PeersAmount, caseRef, caseStorePath, logger)

		caseOutputPath := filepath.Join(fullOutputPath, caseRef)
		if err := os.MkdirAll(caseOutputPath, 0777); err != nil {
			panic(fmt.Sprintf("Error: can't create case output directory: %s", caseOutputPath))
		}
		testCaseFile := filepath.Join(caseOutputPath, "TestCase.json")
		if err := os.WriteFile(testCaseFile, testCase.json, 0666); err != nil {
			panic(fmt.Sprintf("Error: can't create test file copy (%s): %s", caseOutputPath, err.Error()))
		}

		for j, step := range *testCase.Steps {
			stepRef := fmt.Sprintf("Step%d", j)
			stepOutputPath := filepath.Join(caseOutputPath, stepRef)
			if err := os.MkdirAll(stepOutputPath, 0777); err != nil {
				panic(fmt.Sprintf("Error: can't create case output directory: %s", stepOutputPath))
			}

			if err := g.createGroupBranch(ctx, step.TimeoutMS, stepRef, *step.MsgAmount, step.OutOfOrder, *step.Peers); err != nil {
				logger.Error(
					"Step failed, aborting test case",
					zap.String("case", caseRef),
					zap.Int("step", j),
					zap.Error(err),
				)
				break
			}

			for _, peer := range g.peers {
				peerSliceFile := filepath.Join(stepOutputPath, fmt.Sprintf("%s.slice", peer.name))
				if err := os.WriteFile(peerSliceFile, []byte(genSlice(peer)), 0666); err != nil {
					panic(fmt.Sprintf("Error: can't create peer slice file (%s): %s", peerSliceFile, err.Error()))
				}
				peerTreeFile := filepath.Join(stepOutputPath, fmt.Sprintf("%s.tree", peer.name))
				if err := os.WriteFile(peerTreeFile, []byte(genDAG(peer)), 0666); err != nil {
					panic(fmt.Sprintf("Error: can't create peer tree file (%s): %s", peerTreeFile, err.Error()))
				}
			}
		}

		g.close()
	}

	fmt.Println("You can check the results in ", fullOutputPath)
}
