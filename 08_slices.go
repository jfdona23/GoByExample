package main

import "fmt"

func main() {

	s := make([]string, 4)
	fmt.Println("emp:", s)

	s[0] = "a"
	s[1] = "b"
	s[3] = "d"

	fmt.Println("set:", s)
	fmt.Println("get:", s[3])

	fmt.Println("len:", len(s))

	s[2] = "c"
	s = append(s, "e")
	s = append(s, "f")
	fmt.Println("apd:", s)

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy:", c)

	l := s[2:5]
	fmt.Println("sl1:", l)

	l = s[:5]
	fmt.Println("sl2:", l)

	l = s[2:]
	fmt.Println("sl3:", l)

	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)

	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}
