package main

import "fmt"

func main() {
	fs := NewFileSystem()
	fs.Mkdir("/a/b/c")
	fs.Mkdir("/a")
	fmt.Println("New file initialized and directory a/b/c created")
	fmt.Println(fs.Ls("/"))
	fmt.Println(fs.Ls("/a"))
	fmt.Println(fs.Ls("/a/b"))

	fs.AddContentToFile("/a/b/c/notes.txt", "hello ")
	fs.AddContentToFile("/a/b/c/notes.txt", "snowflake!")

	fmt.Println("File content:", fs.ReadContentFromFile("/a/b/c/notes.txt"))
	fmt.Println("Ls on file:", fs.Ls("/a/b/c/notes.txt"))

	fmt.Printf("Root children: %v\n", fs.root.children)
}
