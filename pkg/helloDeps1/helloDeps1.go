package helloDeps1

import "fmt"

func PrintPhrase(phrase string) string {
	fmt.Println(phrase, "called helloDeps1.PrintPhrase")
	return phrase
}
