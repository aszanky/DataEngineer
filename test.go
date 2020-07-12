// Please write a code that handles this use case :

// 	Find first N prime number, and print the result
// 	Input: 4
// 	Output : 2, 3, 5, 7

// 	Sort and combine these 2 lists and print the result
// 	Input 1: 4,3,6,5,1,2
// 	Input 2: F,C,D,B,A
// 	Output : [1:A],[2:B],[3:C],[4:D],[5:F],[6:NULL]

// 	The rules for question no. 4:
// 	Code should be written in your favorite language
// 	Code has to be hosted in Github or Gitlab
// 	Be creative :)
// 	Bonus Points :
// 	You have the Unit Test
// 	You implement the Command Pattern

package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	o := primeNumber(10)
	fmt.Println(o)

	a := []int{4, 3, 6, 5, 1, 2}
	b := []string{"F", "C", "D", "B", "A"}

	combine := combineList(a, b)
	fmt.Println(combine)
}

// question 4 part 1
func primeNumber(input int) (output []int) {
	i := 2
	for {
		if len(output) == input {
			break
		}

		if i == 2 {
			output = append(output, i)
		}

		if i%2 != 0 {
			output = append(output, i)
		}

		i++
	}

	return
}

//question 4 part 2
func combineList(list1 []int, list2 []string) (combined string) {
	sort.Ints(list1)
	sort.Strings(list2)

	var saveHere []string
	for idx, val := range list1 {
		var pair string
		if idx >= len(list2) {
			pair = fmt.Sprintf("[%d:NULL]", val)
			saveHere = append(saveHere, pair)
			continue
		}
		pair = fmt.Sprintf("[%d:%s]", val, list2[idx])
		saveHere = append(saveHere, pair)
	}
	combined = strings.Join(saveHere, ",")
	return
}
