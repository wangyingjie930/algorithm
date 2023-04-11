// Package main
// @Description: copy from chatgpt
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type InvertedIndex map[string][]int

func BuildIndex(docs []string) InvertedIndex {
	index := make(InvertedIndex)
	for i, doc := range docs {
		words := strings.Split(doc, " ")
		for _, word := range words {
			index[word] = append(index[word], i)
		}
	}
	fmt.Println(index)
	return index
}

func main() {
	docs := []string{
		"this is a test",
		"hello world",
		"test test",
		"go programming language",
	}
	index := BuildIndex(docs)
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Enter query: ")
		if !scanner.Scan() {
			break
		}
		query := scanner.Text()
		for _, doc := range index[query] {
			fmt.Println(docs[doc])
		}
	}
}
