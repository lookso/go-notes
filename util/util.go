/*
@Time : 2020-01-01 19:29 
@Author : Tenlu
@File : util
@Software: GoLand
*/
package util

import "errors"

func Sum(a, b int) (c int) {
	return a + b
}

func Division(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("除数不能为0")
	}
	return a / b, nil
}
