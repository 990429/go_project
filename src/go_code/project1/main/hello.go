package main

import "fmt"

func test(a []int) {
	a[0], a[1] = a[1], a[0]
}
func quickSort(node []int) {
	if len(node) == 0 {
		return
	}
	base := node[0]
	i, j := 0, len(node)-1
	for i < j {
		for node[j] >= base && i < j {
			j--
		}
		if i >= j {
			break
		}
		node[i], node[j] = node[j], node[i]
		for node[i] <= base && i < j {
			i++
		}
		if i >= j {
			break
		}
		node[i], node[j] = node[j], node[i]
	}
	quickSort(node[:i])
	quickSort(node[i+1:])
}
type bmap struct{
	tophash [8]uint8
	data [16]byte
	overflow *bmap
}
func main() {
	a := []int{5, 4, 3, 2, 1}
	quickSort(a)
	fmt.Println(a)
}
