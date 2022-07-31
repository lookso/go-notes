package main

import (
	"fmt"
	"github.com/dlclark/regexp2"
)

func main() {
	//re1 := regexp2.MustCompile(`Windows(?=95|98|NT|2000)`, regexp2.RE2)
	//fmt.Println(re1.MatchString("Windows95"))
	//if m, _ := re1.FindStringMatch("Windows95"); m != nil {
	//	fmt.Println(m.String())
	//}

	re := regexp2.MustCompile("([一]+[加]+[一]?)(?=.)", 0)
	if isMatch, _ := re.MatchString("一加一啊"); isMatch {

		if m, _ := re.FindStringMatch("一加一啊"); m != nil {
			// the whole match is always group 0
			fmt.Printf("Group 0: %v\n", m.String())

			// you can get all the groups too
			gps := m.Groups()
			//
			//// a group can be captured multiple times, so each cap is separately addressable
			fmt.Printf("Group 1, first capture", gps[0].Captures[0].String())
			fmt.Printf("Group 1, second capture", gps[0].Captures[0].String())
		}
	}
}
