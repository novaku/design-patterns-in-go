package framework

type EvictionAlgo interface {
	Evict(c *Cache)
}