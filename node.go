package main

type AccessType int

const (
	Deny AccessType = iota
	Allow
)

type FileSystemNode struct {
	isFile      bool
	content     string
	children    map[string]*FileSystemNode
	permissions map[string]AccessType // Map of user to access type
}

func newDirectoryNode() *FileSystemNode {
	return &FileSystemNode{
		isFile:      false,
		children:    make(map[string]*FileSystemNode),
		permissions: make(map[string]AccessType),
	}
}

func newFileNode() *FileSystemNode {
	return &FileSystemNode{
		isFile:      true,
		children:    make(map[string]*FileSystemNode),
		permissions: make(map[string]AccessType),
	}
}
