package framework

type Inode interface {
	Print(string)
	Clone() Inode
}