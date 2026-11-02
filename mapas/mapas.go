package mapas

import "fmt"

func CalcMillion() {
	N := 1_000_000
	nums := make([]int, N)
	for i := range N {
		nums[i] = i + 1
	}
	sum := 0
	for _, item := range nums {
		sum += item
	}
	fmt.Print(sum)
}
