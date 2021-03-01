package main

func main(){
	twoSum()
}

func twoSum(nums []int, target int) []int {
	m := make(map[int]int, len(nums))
	existsKeys := make([]int, 0)

	for k, v := range nums {
		if _, ok := m[target-v]; ok {
			existsKeys = append(existsKeys, k,m[target-v])
		}
		m[v] = k
	}
	return existsKeys
}
