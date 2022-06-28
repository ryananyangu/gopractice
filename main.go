package main

import (
	"fmt"
	"strings"
)

//
func twoSum(nums []int, target int) []int {
	for index, value := range nums {
		for index2, value2 := range nums {
			if index2 == index {
				continue
			}
			if (value + value2) == target {
				return []int{index, index2}
			}

		}
	}
	return []int{0, 0}

}

func lengthOfLongestSubstring(s string) int {
	current_long := ""

	for index, char := range s {
		now := fmt.Sprintf("%c", char)
		if index == 0 {
			current_long += now
			continue
		}

		// strings.Contains()
	}

	return len(current_long)

}

func isPalindrome(x int) bool {
	tostr := fmt.Sprintf("%d", x)
	intsize := len(tostr) - 1
	halved := intsize / 2
	for i := 0; i <= halved; i++ {
		if tostr[i] == tostr[intsize-i] {

			continue
		}
		return false
	}
	return true

}

func longestCommonPrefix(strs []string) string {
	location := 0
	prefix := ""
	run := true
	size := len(strs)
	count := 0
	current_letter := ""

	if size == 1 {
		return strs[0]
	}
	if size == 0 {
		return ""
	}
	splt := strings.Split(strs[0], "")
	previous_letter := ""
	if len(splt) != 0 {
		previous_letter = splt[location]
	}

	for run {

		for _, word := range strs {
			splt = strings.Split(word, "")
			if len(splt) != 0 {
				if location >= len(splt) {
					return prefix
				}
				current_letter = splt[location]
			} else {
				current_letter = ""
			}

			if current_letter == previous_letter {
				count += 1
			} else {
				break
			}
		}

		if count == size {
			prefix += current_letter
			location += 1
			// previous_letter = strings.Split(strs[0], "")[location]

			splt := strings.Split(strs[0], "")
			previous_letter = ""
			if len(splt) != 0 {
				if location >= len(splt) {
					return prefix
				}
				previous_letter = splt[location]
			} else {
				return prefix
			}

			count = 0
			continue

		} else {

			run = false
		}

	}
	return prefix

}

func longestCommonPrefix2(strs []string) string {
	obj := map[string]string{}
	result := ""
	for r_index, word := range strs {
		for c_index, letter := range strings.Split(word, "") {
			key := fmt.Sprintf("%d%d", r_index, c_index)
			obj[key] = letter
		}
	}
	for 
	fmt.Println(obj)
	return result
}
func main() {
	// fmt.Println(lengthOfLongestSubstring("pwwkew"))
	// fmt.Println(twoSum([]int{3, 3}, 6))
	// fmt.Println(isPalindrome(101001))
	fmt.Println(longestCommonPrefix2([]string{"flower", "flower", "flower", "flower"}))

}
