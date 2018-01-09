package main

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"strings"
	"testing"
)

func printTree(n *TreeNode, level int) {
	fmt.Printf("%s %v \n", strings.Repeat("--- ", level), hex.EncodeToString(n.Hash[:]))
	if n.Left != nil {
		printTree(n.Left, level+1)
	}
	if n.Right != nil {
		printTree(n.Right, level+1)
	}
}

func TestCreateRootHash(t *testing.T) {
	data := []byte("Gdzies jest lecz nie wiadomo gdzie, malenka pszczolka co sie zwie, Malutka Maja mieszka tu")
	r := bytes.NewReader(data)
	tree := BuildTree(r)
	t.Errorf("Hash: %v", tree)
	printTree(tree, 0)

}
