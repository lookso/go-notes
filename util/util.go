/*
@Time : 2020-01-01 19:29
@Author : Tenlu
@File : util
@Software: GoLand
*/
package util

import (
	"errors"
	"fmt"
)

func Sum(a, b int) (c int) {
	return a + b
}

func Division(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("除数不能为0")
	}
	return a / b, nil
}

func main() {
	var dbs = []string{"35", "36", "38", "40", "41"}
	var form = []string{"34", "35", "36", "37"}

	fmt.Println(intersect(dbs, form))
	fmt.Println(union(dbs, form))
	fmt.Println(difference(dbs, form))
	fmt.Println(difference(form, dbs))
}

func intersect(slice1, slice2 []string) []string {
	m := make(map[string]int)
	nn := make([]string, 0)
	for _, v := range slice1 {
		m[v]++
	}

	for _, v := range slice2 {
		times, _ := m[v]
		if times == 1 {
			nn = append(nn, v)
		}
	}
	return nn
}

func union(slice1, slice2 []string) []string {
	m := make(map[string]int)
	for _, v := range slice1 {
		m[v]++
	}

	for _, v := range slice2 {
		times, _ := m[v]
		if times == 0 {
			slice1 = append(slice1, v)
		}
	}
	return slice1
}

func difference(slice1, slice2 []string) []string {
	m := make(map[string]int)
	nn := make([]string, 0)
	inter := intersect(slice1, slice2)
	for _, v := range inter {
		m[v]++
	}
	for _, value := range slice1 {
		times, _ := m[value]
		if times == 0 {
			nn = append(nn, value)
		}
	}
	return nn
}

//[35 36]
//[35 36 38 40 41 34 37]
//[38 40 41]
//[34 37]
