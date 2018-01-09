package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"io"
)

func main() {
	fmt.Print("Hello")
}

type TreeNode struct {
	Hash  [32]byte
	Left  *TreeNode
	Right *TreeNode
}

func CreateTree(l []*TreeNode, start int, end int) *TreeNode {
	fmt.Println(start, end)
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

func BuildTree(r *bytes.Reader) *TreeNode {
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

func RootHash(r *bytes.Reader) [32]byte {
	return BuildTree(r).Hash
}
