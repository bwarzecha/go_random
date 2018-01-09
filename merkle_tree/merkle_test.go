package main

import (
	"bytes"
	"testing"
)

func TestCreateRootHash(t *testing.T) {
	data := []byte("Gdzies jest lecz nie wiadomo gdzie, malenka pszczolka co sie zwie, Malutka Maja mieszka tu")
	r := bytes.NewReader(data)
	tree := BuildTree(r)
	PrintTree(tree, 0)
}
