package main

import (
	"fmt"
)

func main() {
	fs := NewFileSystem()
	// fs.Mkdir("/a/b/c")
	// fs.Mkdir("/a")
	// fmt.Println("New file initialized and directory a/b/c created")
	// fmt.Println(fs.Ls("/"))
	// fmt.Println(fs.Ls("/a"))
	// fmt.Println(fs.Ls("/a/b"))

	// fs.AddContentToFile("/a/b/c/notes.txt", "hello ")
	// fs.AddContentToFile("/a/b/c/notes.txt", "snowflake!")

	// fmt.Println("File content:", fs.ReadContentFromFile("/a/b/c/notes.txt"))
	// fmt.Println("Ls on file:", fs.Ls("/a/b/c/notes.txt"))

	// fmt.Printf("Root children: %v\n", fs.root.children)
	// var wg sync.WaitGroup

	// // Spawn 50 writer goroutines creating dirs and writing logs
	// // numWorkers := 50
	// numWorkers := 4
	// for i := 0; i < numWorkers; i++ {
	// 	wg.Add(1)
	// 	go func(id int) {
	// 		defer wg.Done()

	// 		dirPath := fmt.Sprintf("/logs/worker_%d", id)
	// 		filePath := dirPath + "/out.log"
	// 		content := fmt.Sprintf("payload from %d", id)

	// 		fmt.Printf("[WRITER %d] Waiting for write lock to Mkdir: %s\n", id, dirPath)
	// 		fs.Mkdir(dirPath)

	// 		fmt.Printf("[WRITER %d] Waiting for write lock to write file: %s\n", id, filePath)
	// 		fs.AddContentToFile(filePath, content)

	// 		fmt.Printf("[WRITER %d] Done writing!\n", id)
	// 	}(i)
	// }

	// // Spawn 50 reader goroutines reading concurrently
	// for i := 0; i < numWorkers; i++ {
	// 	wg.Add(1)
	// 	go func(id int) {
	// 		defer wg.Done()

	// 		// Small jitter so readers attempt reads while writes are occurring
	// 		time.Sleep(1 * time.Millisecond)

	// 		fmt.Printf("  [READER %d] Requesting read lock on /logs...\n", id)
	// 		entries := fs.Ls("/logs")
	// 		fmt.Printf("  [READER %d] Read success! Saw %d directories: %v\n", id, len(entries), entries)
	// 	}(i)
	// }
	// wg.Wait()
	// fmt.Println("\n=== Final Verification ===")
	// finalDirs := fs.Ls("/logs")
	// fmt.Printf("Total directories created: %d\n", len(finalDirs))
	// fmt.Printf("All directories: %v\n", finalDirs)

	// 1. Build directory tree and create the file
	fs.Mkdir("/finance/reports")
	fs.AddContentToFile("/finance/reports/secret.txt", "confidential budget figures")
	fs.AddContentToFile("/finance/reports/summary.txt", "public summary")

	// 2. Set base permission: Alice has access to /finance
	fs.SetPermissions("/finance", "alice", Allow)

	// 3. Set local override: Alice is explicitly blocked from secret.txt
	fs.SetPermissions("/finance/reports/secret.txt", "alice", Deny)

	// --- Verifications ---

	// Direct permission
	fmt.Printf("1. /finance -> %t (expected: true)\n",
		fs.CanAccess("/finance", "alice"))

	// Inherited permission (reports has no explicit rule, inherits from /finance)
	fmt.Printf("2. /finance/reports -> %t (expected: true)\n",
		fs.CanAccess("/finance/reports", "alice"))

	// Inherited permission on child file
	fmt.Printf("3. /finance/reports/summary.txt -> %t (expected: true)\n",
		fs.CanAccess("/finance/reports/summary.txt", "alice"))

	// Local Deny override
	fmt.Printf("4. /finance/reports/secret.txt -> %t (expected: false)\n",
		fs.CanAccess("/finance/reports/secret.txt", "alice"))

	// Bob has no permissions assigned anywhere
	fmt.Printf("5. Bob on /finance -> %t (expected: false)\n",
		fs.CanAccess("/finance", "bob"))
}
