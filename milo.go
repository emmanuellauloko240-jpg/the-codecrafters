package main

import (
	"fmt"
	"strings"
)

func upper(me string) string {
	return strings.ToUpper(me)
}

func cap(is string) string {
	return strings.ToUpper(is[1:]) + strings.ToLower(is[:1])
}

func lower(If string) string {
	return strings.ToLower(If)
}

// func reverse(two string) string {
// 	r := []rune(two)
// 	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
// 		r[i], r[j] = r[j], r[i]
// 	}
// 	return strings.Join(r, " ")

// }

func title(one string) string {
	the := map[string]bool{
		"a":   true,
		"an":  true,
		"the": true,
		"and": true,
		"but": true,
		"or":  true,
		"for": true,
		"nor": true,
		"on":  true,
		"at":  true,
		"to":  true,
		"by":  true,
		"in":  true,
		"of":  true,
		"up":  true,
		"as":  true,
		"is":  true,
		"it":  true,
	}
	text := strings.Fields(one)
	for i, w := range text {
		if i == 0 || !the[w] {
			text[i] = strings.ToUpper(string(w[0])) + w[1:]
		}

	}
	return strings.Join(text, " ")

}
func snake(lolipop string) string {
	w := strings.ToLower(lolipop)
	return strings.ReplaceAll(w, " ", "_")
}

func main() {
	var me string
	fmt.Println("Enter word:")
	fmt.Scanln(&me)

	var operation string
	fmt.Println("upper", "lower", "title", "reverse", "Quit")
	fmt.Scanln(&operation)
	switch operation {
	case "upper":
		fmt.Println(upper(me))
	case "lower":
		fmt.Println(lower(me))
	case "title":
		fmt.Println(title(me))
	// case "reverse":
	// 	fmt.Println(reverse(me))
	case "quit":
		fmt.Println("Goodbye")
	}

}
