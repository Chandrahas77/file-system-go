package main

type FileSystemNode struct {
	isFile   bool 
	content  string
	children map[string]*FileSystemNode
}

func newDirectoryNode() *FileSystemNode {
	return &FileSystemNode{
		isFile:   false,
		children: make(map[string]*FileSystemNode),
	}
}

func newFileNode() *FileSystemNode {
	return &FileSystemNode{
		isFile:   true,
		children: make(map[string]*FileSystemNode),
	}
}
