package main

import (
	ch "consistentHashing/hash"
	"fmt"
)

const Nodes = 3

var CatalogNodeMap map[int]uint64

func main() {
	//CatalogNodeMap[590492522508486] = 0
	//CatalogNodeMap[428444501303217] = 1
	//CatalogNodeMap[338054854556765] = 2
	CatalogNodeMap = make(map[int]uint64)
	var (
		shop1 = 19999934
		shop2 = 183839494
		shop3 = 1727348494
		shop4 = 27246682
		shop5 = 892742939
		shop6 = 282000948
	)
	CatalogNodeMap[0] = 590492522508486
	CatalogNodeMap[1] = 428444501303217
	CatalogNodeMap[2] = 338054854556765
	h := ch.NewConsistentHash()
	//adding shops
	h.Add(0)
	h.Add(1)
	h.Add(2)
	val, _ := h.Get(shop1)
	fmt.Printf("bucket for shop %d is %d\n", shop1, val)
	val, _ = h.Get(shop2)
	fmt.Printf("bucket for shop %d is %d\n", shop2, val)
	val, _ = h.Get(shop3)
	fmt.Printf("bucket for shop %d is %d\n", shop3, val)
	val, _ = h.Get(shop4)
	fmt.Printf("bucket for shop %d is %d\n", shop4, val)
	val, _ = h.Get(shop5)
	fmt.Printf("bucket for shop %d is %d\n", shop5, val)
	val, _ = h.Get(shop6)
	fmt.Printf("bucket for shop %d is %d\n", shop6, val)
}
