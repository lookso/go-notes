package main
import "fmt"
type StuReq struct {
	ID int `json:"id"`
}
func main() {
	var newStu = new(StuReq)
	for i := 0; i < 10; i++ {
		newStu.ID = i
		PrintFunc(newStu)
	}
}
func PrintFunc(req *StuReq) {
	fmt.Println(req.ID)
}

/**
性能差:
package main
import "fmt"
type StuReq struct {
	ID int `json:"id"`
}
func main() {
	for i := 0; i < 10; i++ {
		newStu:=StuReq{ID: i}   // 需要声明拷贝n份
		PrintFunc(newStu)
	}
}
func PrintFunc(req StuReq) {
	fmt.Println(req.ID)
}
*/
