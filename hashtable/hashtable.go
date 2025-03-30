package hashtable

import (
	"fmt"
	"hash/fnv"
)

type entry struct {
	key   string
	value string
}

type HashTable struct {
	size    int
	buckets [][]entry
}

func NewHashTable(size int) *HashTable {
	return &HashTable{
		size:    size,
		buckets: make([][]entry, size),
	}
}

func (ht *HashTable) hash(key string) int {
	h := fnv.New32a()
	h.Write([]byte(key))
	return int(h.Sum32()) % ht.size
}

func (ht *HashTable) Set(key, value string) {
	index := ht.hash(key)
	for i, e := range ht.buckets[index] {
		if e.key == key {
			ht.buckets[index][i].value = value
			return
		}
	}
	ht.buckets[index] = append(ht.buckets[index], entry{key, value})
}

func (ht *HashTable) Get(key string) (string, error) {
	index := ht.hash(key)
	for _, e := range ht.buckets[index] {
		if e.key == key {
			return e.value, nil
		}
	}
	return "", fmt.Errorf("No se encuentra el valor")
}

func (ht *HashTable) Delete(key string) error {
	index := ht.hash(key)
	bucket := ht.buckets[index]
	for i, e := range bucket {
		if e.key == key {
			ht.buckets[index] = append(bucket[:i], bucket[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("No se encuentra el valor")
}
func (ht *HashTable) Print() {
	fmt.Println("Hashtable actual:")
	for i, bucket := range ht.buckets {
		fmt.Printf("Bucket %d: ", i)
		if len(bucket) == 0 {
			fmt.Println("(empty)")
		} else {
			for _, e := range bucket {
				fmt.Printf("[%s: %s] ", e.key, e.value)
			}
			fmt.Println()
		}
	}
}
