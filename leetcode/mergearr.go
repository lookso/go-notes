/*
@Time : 2019-08-16 07:17 
@Author : Tenlu
@File : mergearr
@Software: GoLand
*/
package main

import "fmt"

//给定两个有序整数数组 nums1 和 nums2，将 nums2 合并到 nums1 中，使得 num1 成为一个有序数组。
//
//说明:
//
//初始化 nums1 和 nums2 的元素数量分别为 m 和 n。
//你可以假设 nums1 有足够的空间（空间大小大于或等于 m + n）来保存 nums2 中的元素。
//示例:
//
//输入:
//nums1 = [1,2,3,0,0,0], m = 3
//nums2 = [2,5,6],       n = 3
//
//输出: [1,2,2,3,5,6]

func main() {
	var nums1 = []int{1, 2, 3, 0, 0, 0}
	var m = 3
	var nums2 = []int{2, 5, 6, 0, 0, 0, 0, 0}
	var n = 3
	merge(nums1, m, nums2, n)
}

func merge(nums1 []int, m int, nums2 []int, n int) {
	temp := make([]int, m)
	copy(temp, nums1)
	t, j := 0, 0 //t为temp的索引，j为nums2的索引
	for i := 0; i < len(nums1); i++ {
		//当t大于temp的长度，那就是说temp全部放进去了nums1中，那剩下的就是放nums2剩余的值了
		if t >= len(temp) {
			nums1[i] = nums2[j]
			j++
			continue
		}
		//当j大于nums2的长度的时候，那就是说明nums2全部都放进去了nums1中，那剩下的就是放temp剩余的值了
		if j >= n {
			nums1[i] = temp[t]
			t++
			continue
		}
		//比较nums2与temp对应值的大小，小的那个就放进nums1中
		if nums2[j] <= temp[t] {
			nums1[i] = nums2[j]
			j++
		} else {
			nums1[i] = temp[t]
			t++
		}
	}
	fmt.Println(nums1)

}
