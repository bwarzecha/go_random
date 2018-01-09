package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: merkle file_path")
		return
	}
	path := os.Args[1]
	f, err := os.Open(path)
	if err != nil {
		fmt.Errorf("Unable to open file %s. Message:  %v", path, err)
	}
	tree := BuildTree(f)
	PrintTree(tree, 0)
}

type TreeNode struct {
	Hash  [32]byte
	Left  *TreeNode
	Right *TreeNode
}

func CreateTree(l []*TreeNode, start int, end int) *TreeNode {
	if start == end {
		return l[start]
	}
	var left *TreeNode
	var right *TreeNode
	if end-start == 1 {
		left = l[start]
		right = l[end]
	} else {
		lStart := start
		lEnd := end - ((end - start) / 2)
		if ((lEnd-lStart)+1)%2 != 0 {
			lEnd = lEnd - 1
		}
		rStart := lEnd + 1
		rEnd := end
		left = CreateTree(l, lStart, lEnd)
		right = CreateTree(l, rStart, rEnd)
	}
	concat := append(left.Hash[:], right.Hash[:]...)
	var hash = sha256.Sum256(concat)
	return &TreeNode{hash, left, right}
}

func PrintTree(n *TreeNode, level int) {
	fmt.Printf("%s%v \n", strings.Repeat("---- ", level), hex.EncodeToString(n.Hash[:])[0:5])
	if n.Left != nil {
		PrintTree(n.Left, level+1)
	}
	if n.Right != nil {
		PrintTree(n.Right, level+1)
	}
}

func BuildTree(r io.Reader) *TreeNode {
	var nodes []*TreeNode
	chunk := make([]byte, 16)

	var hash [32]byte
	for {
		n, err := r.Read(chunk)
		if n == 16 {
			hash = sha256.Sum256(chunk)
		} else {
			hash = sha256.Sum256(chunk[0:n])
		}
		node := &TreeNode{hash, nil, nil}
		nodes = append(nodes, node)
		if err == io.EOF {
			break
		}
	}
	return CreateTree(nodes, 0, len(nodes)-1)
}

func RootHash(r io.Reader) [32]byte {
	return BuildTree(r).Hash
}
