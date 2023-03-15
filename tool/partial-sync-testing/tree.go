package main

import (
	"fmt"
	"strings"

	"berty.tech/go-ipfs-log/iface"
	"berty.tech/go-orbit-db/stores/operation"

	"github.com/xlab/treeprint"
)

func getEntryValue(entry iface.IPFSLogEntry) string {
	op, err := operation.ParseOperation(entry)
	if err != nil {
		panic(err)
	}

	return string(op.GetValue())
}

func genSlice(p *peer) string {
	slice := "[\n"

	for _, entry := range p.store.OpLog().GetEntries().Slice() {
		msg := getEntryValue(entry)
		slice = fmt.Sprintf("%s  %s,\n", slice, msg)
	}

	return slice + "]\n"
}

func exploreBranch(seen map[string]bool, entry iface.IPFSLogEntry, entries iface.IPFSLogOrderedEntries, current string, branch treeprint.Tree) {
	first := true

	for {
		if _, ok := seen[entry.GetHash().String()]; ok {
			return
		}

		msg := getEntryValue(entry)
		println("ZOB")
		if first {
			first = false
			next := strings.Split(msg, "_")[0]
			println("NEXT", next, "CURRENT", current, branch.String())
			if next != current {
				branch = branch.AddBranch(next)
				println("AFTER", branch.String())
				current = next
			} else {
				branch = branch.FindLastNode().Branch()
			}
			continue
		}
		println("ZIB")

		branch.AddNode(msg)
		seen[entry.GetHash().String()] = true

		parents := entry.GetNext()

		println("CURRENT ", getEntryValue(entry), entry.GetHash().String())
		if len(parents) == 1 {
			entry = entries.UnsafeGet(parents[0].String())
			println("PARENT ", getEntryValue(entry), parents[0].String())
		} else {
			if len(parents) > 1 {
				for _, parent := range parents {
					entry := entries.UnsafeGet(parent.String())
					println("PARENTS ", getEntryValue(entry), parent.String())
				}
				println()
				for _, parent := range parents {
					entry := entries.UnsafeGet(parent.String())
					exploreBranch(seen, entry, entries, current, branch)
				}
			}
			return
		}
		println()
	}
}

func genDAG(p *peer) string {
	heads := p.store.OpLog().Heads().Slice()
	root := "Head"
	if len(heads) > 1 {
		root = "Heads"
	} else if len(heads) == 0 {
		root = "Empty"
	}

	dag := treeprint.NewWithRoot(root)
	seen := make(map[string]bool)

	for _, head := range heads {
		println("HEAD ", getEntryValue(head), head.GetHash().String())
	}
	for _, head := range heads {
		exploreBranch(seen, head, p.store.OpLog().GetEntries(), root, dag)
	}

	return dag.String()
}
