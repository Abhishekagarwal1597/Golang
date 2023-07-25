// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"strings"
)

func main() {
	orignalString := "myself abhi"
	singleStringInstance := strings.Split(orignalString, " ")
	var reverseString []string
	for i := 0; i < len(singleStringInstance); i++ {
		fmt.Println(singleStringInstance[i])
		newInstance := []rune(singleStringInstance[i])

		for j := 0; j < len(newInstance)/2; j++ {
			newInstance[j], newInstance[len(newInstance)-1] = newInstance[len(newInstance)-1], newInstance[j]
		}
		fmt.Println("newInstance:", string(newInstance))
		reverseString = append(reverseString, string(newInstance))
	}
	fmt.Println(reverseString)
}
