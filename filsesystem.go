package main

import (
	"strings"
	"sync"
)

type FileSystem struct {
	mu   sync.RWMutex
	root *FileSystemNode
}

func NewFileSystem() *FileSystem {
	return &FileSystem{
		//root is always a directory node, so we initialize it as such ("/").
		root: newDirectoryNode(),
	}
}

func parsePath(path string) []string {
	parts := strings.Split(path, "/")
	var tokens []string
	for _, part := range parts {
		if part != "" {
			tokens = append(tokens, part)
		}
	}
	return tokens
}

func (fs *FileSystem) Mkdir(path string) {
	tokens := parsePath(path)
	curr := fs.root
	for _, token := range tokens {
		if _, exists := curr.children[token]; !exists {
			curr.children[token] = newDirectoryNode()
		}
		curr = curr.children[token]
	}
}

func (fs *FileSystem) Ls(path string) []string {
	tokens := parsePath(path)
	curr := fs.root
	for _, token := range tokens {
		curr = curr.children[token]
	}
	if curr.isFile {
		return []string{tokens[len(tokens)-1]}
		//If it's a directory, we return the sorted list of its children.
	}

	sortedChildren := make([]string, 0, len(curr.children))
	for child := range curr.children {
		sortedChildren = append(sortedChildren, child)
	}
	return sortedChildren
}

func (fs *FileSystem) AddContentToFile(path string, content string) {
	tokens := parsePath(path)
	curr := fs.root
	for i := 0; i < len(tokens)-1; i++ {
		token := tokens[i]
		if _, exists := curr.children[token]; !exists {
			curr.children[token] = newDirectoryNode()
		}
		curr = curr.children[token]
	}
	fileName := tokens[len(tokens)-1]
	if _, exists := curr.children[fileName]; !exists {
		curr.children[fileName] = newFileNode()
	}

	curr.children[fileName].content += content
}

func (fs *FileSystem) ReadContentFromFile(path string) string {
	tokens := parsePath(path)
	curr := fs.root
	for _, token := range tokens {
		curr = curr.children[token]
	}
	return curr.content
}
