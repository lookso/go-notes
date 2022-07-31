package main

import "fmt"

func main() {

	fmt.Println([]byte("1"))
	a := make([][]byte, 0)
	a = append(a, []byte{48, 48, 48, 48, 48})
	a = append(a, []byte{48, 49, 49, 49, 49})
	a = append(a, []byte{48, 49, 49, 49, 49})
	a = append(a, []byte{49, 49, 49, 48, 48})
	fmt.Println(maximalRectangle(a))
}
func maximalRectangle(matrix [][]byte) (ans int) {
	if len(matrix) == 0 {
		return
	}
	m, n := len(matrix), len(matrix[0])
	left := make([][]int, m)
	for i, row := range matrix {
		left[i] = make([]int, n)
		for j, v := range row {
			if v == '0' {
				continue
			}
			if j == 0 {
				left[i][j] = 1
			} else {
				left[i][j] = left[i][j-1] + 1
			}
		}
	}
	fmt.Println(left)
	for i, row := range matrix {
		for j, v := range row {
			if v == '0' {
				continue
			}
			width := left[i][j]
			area := width
			for k := i - 1; k >= 0; k-- {
				width = min(width, left[k][j])
				area = max(area, (i-k+1)*width)
			}
			ans = max(ans, area)
		}
	}
	return
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
