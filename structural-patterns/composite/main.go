package main

import "design-patterns/structural-patterns/composite/framework"

func main() {
	file1 := &framework.File{Name: "File1"}
	file2 := &framework.File{Name: "File2"}
	file3 := &framework.File{Name: "File3"}

	folder1 := &framework.Folder{
		Name: "Folder1",
	}

	folder1.Add(file1)

	folder2 := &framework.Folder{
		Name: "Folder2",
	}
	folder2.Add(file2)
	folder2.Add(file3)
	folder2.Add(folder1)

	folder2.Search("rose")
}