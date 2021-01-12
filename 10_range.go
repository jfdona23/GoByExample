package main

import "fmt"

func main() {

	nums := []int{1, 2, 3}
	sum := 0
	for index, num := range nums {
		sum += num
		fmt.Println("idx:", index)
	}
	fmt.Println("sum:", sum)

	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	for k := range kvs {
		fmt.Println("key:", k)
	}

	for index, char := range "juan" {
		fmt.Println(index, char)
	}
}
