package main

import (
	"bufio"
	"fmt"
	"os"
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

func reverse(two string) string {
	r := []rune(two)
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)

}

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
	Reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("upper", "lower", "title", "reverse","snake","cap", "Quit")
		fmt.Println("Enter word:")
		me, _ := Reader.ReadString('\n')
		me = strings.TrimSpace(me)

		text := strings.Fields(me)
		operation := text[0]
		word := text[1:]
		words := strings.Join(word, " ")
		if len(words) <= 1 {
			fmt.Println("No text Provided")
			continue
		}
		switch operation {
		case "upper":
			fmt.Println(upper(words))
		case "lower":
			fmt.Println(lower(words))
		case "snake":
		    fmt.Println(snake(words))
		case "cap":
			fmt.Println(cap(words))
		case "title":
			fmt.Println(title(words))
		case "reverse":
			fmt.Println(reverse(words))
		case "quit":
			fmt.Println("Goodbye")
		}
	}
}
