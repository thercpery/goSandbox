package main

type EnglishBot struct{}
type SpanishBot struct{}

func main() {
	eb := EnglishBot{}
	sb := SpanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

// This will cause an error because the function name "printGreeting" is redeclared

// func printGreeting(eb EnglishBot) {
// 	fmt.Println(eb.getGreeting())
// }

// func printGreeting(sb SpanishBot) {
// 	fmt.Println(sb.getGreeting())
// }

func (EnglishBot) getGreeting() string {
	// VERY custom logic for generating an English greeting.
	return "Hi There!"
}

func (SpanishBot) getGreeting() string {
	return "Hola!"
}
