package main

import (
	"fmt"

	"design-patterns/creational-patterns/prototype/framework"
)

func main() {
	file1 := &framework.File{Name: "File1"}
	file2 := &framework.File{Name: "File2"}
	file3 := &framework.File{Name: "File3"}

	folder1 := &framework.Folder{
		Children: []framework.Inode{file1},
		Name:      "Folder1",
	}

	folder2 := &framework.Folder{
		Children: []framework.Inode{folder1, file2, file3},
		Name:      "Folder2",
	}
	fmt.Println("\nPrinting hierarchy for Folder2")
	folder2.Print("  ")

	cloneFolder := folder2.Clone()
	fmt.Println("\nPrinting hierarchy for clone Folder")
	cloneFolder.Print("  ")
}