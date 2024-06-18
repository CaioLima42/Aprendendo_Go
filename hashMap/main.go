package main

import (
	"fmt"
)

type HashMap struct{
	items [100]string
}
func NewHashMap() HashMap{
	return HashMap{items: [100]string{}}
}
func (h *HashMap) GetHashKey(key string) int{
	totalSum := 0
	for _, v := range key{
		totalSum+= int(v)
	}
	hashKey := totalSum %100
	return hashKey
}

func (h *HashMap) Set(key, val string){
	hashKey := h.GetHashKey(key)
	h.items[hashKey] = val
}

func (h HashMap) Get(key string) string{
	hashKey := h.GetHashKey(key)
	return h.items[hashKey]
}

func main(){
	aa := NewHashMap()
	aa.Set("a", "sample value")
	aa.Set("ABB", "unexpected sample value")
	fmt.Println(aa.Get("a"))
}