# It Will shift each char in a inputString by shift char.
package main

import (
	"fmt"
)

// ciser-chipeur
func main() {
	var inputCipher string
	var shift rune
	fmt.Println("input cipher string and shift")
	fmt.Scanf("%s %d", &inputCipher, &shift)
	fmt.Println(inputCipher, shift)
	newString := chipheurShift(inputCipher, shift)
	fmt.Println(string(newString))
}

func chipheurShift(orignalString string, shift rune) []rune {
	var newRune []rune
	for _, v := range orignalString {
		if v >= 'A' && v <= 'Z' {
			updatedRune := shiftChipheur(v, 'A', shift)
			fmt.Println(updatedRune)
			newRune = append(newRune, updatedRune)
		} else if v >= 'a' && v <= 'z' {
			updatedRune := shiftChipheur(v, 'a', shift)
			fmt.Println(updatedRune)
			newRune = append(newRune, updatedRune)
		} else {
			newRune = append(newRune, v)
		}

	}
	return newRune
}

func shiftChipheur(index rune, base rune, shift rune) rune {
	tmp := index - base
	shift = (tmp + shift) % 26
	fmt.Println(rune(shift))
	return base + rune(shift)

}
