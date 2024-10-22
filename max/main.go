package main

func Max(a []int) int {
	if len(a) == 0 {
		return 0
	}
	max := a[0]
	for _, i := range a {
		if i > max {
			max = i
		}
	}
	return max
}
