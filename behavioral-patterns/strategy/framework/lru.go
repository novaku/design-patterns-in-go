package framework

import "fmt"

type Lru struct {
}

func (l *Lru) Evict(c *Cache) {
	fmt.Println("Evicting by lru strtegy")
}