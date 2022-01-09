package main

import "design-patterns/behavioral-patterns/strategy/framework"

func main() {
	lfu := &framework.Lfu{}
	cache := framework.InitCache(lfu)

	cache.Add("a", "1")
	cache.Add("b", "2")

	cache.Add("c", "3")

	lru := &framework.Lru{}
	cache.SetEvictionAlgo(lru)

	cache.Add("d", "4")

	fifo := &framework.Fifo{}
	cache.SetEvictionAlgo(fifo)

	cache.Add("e", "5")
}