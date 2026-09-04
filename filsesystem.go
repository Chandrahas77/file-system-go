package main

import (
	"sort"
	"strings"
	"sync"
)

type FileSystem struct {
	mu   sync.RWMutex
	root *FileSystemNode
}

func NewFileSystem() *FileSystem {
	return &FileSystem{
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
	fs.mu.Lock()
	defer fs.mu.Unlock()
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
	fs.mu.RLock()
	defer fs.mu.RUnlock()
	tokens := parsePath(path)
	curr := fs.root
	for _, token := range tokens {
		next, exists := curr.children[token]
		if !exists {
			return []string{}
		}
		curr = next
	}
	if curr.isFile {
		return []string{tokens[len(tokens)-1]}
	}
	sortedChildren := make([]string, 0, len(curr.children))

	for child := range curr.children {
		sortedChildren = append(sortedChildren, child)
	}
	sort.Strings(sortedChildren)
	return sortedChildren
}

func (fs *FileSystem) AddContentToFile(path string, content string) {
	fs.mu.Lock()
	defer fs.mu.Unlock()
	tokens := parsePath(path)
	if len(tokens) == 0 {
		return
	}

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
	fs.mu.RLock()
	defer fs.mu.RUnlock()
	tokens := parsePath(path)
	curr := fs.root
	for _, token := range tokens {
		curr = curr.children[token]
	}
	return curr.content
}

func (fs *FileSystem) SetPermissions(path string, user string, access AccessType) {
	fs.mu.Lock()
	defer fs.mu.Unlock()

	tokens := parsePath(path)
	curr := fs.root

	for _, token := range tokens {
		if _, exists := curr.children[token]; !exists {
			curr.children[token] = newDirectoryNode()
		}
		curr = curr.children[token]
	}
	curr.permissions[user] = access
}

func (fs *FileSystem) CanAccess(path string, user string) bool {
	fs.mu.RLock()
	defer fs.mu.RUnlock()

	curr := fs.root
	currentAccess := Deny
	tokens := parsePath(path)

	if rule, exists := curr.permissions[user]; exists {
		currentAccess = rule
	}
	for _, token := range tokens {
		next, exists := curr.children[token]
		if !exists {
			return false //path doesnt exist in furthter
		}
		curr = next
		if rule, exists := curr.permissions[user]; exists {
			currentAccess = rule
		}
	}
	return currentAccess == Allow
}
