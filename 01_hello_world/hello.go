package hello

import "fmt"

const (
	spanish               = "Spanish"
	french                = "French"
	englishTextWithPrefix = "Hello, "
	spanishTextWithPrefix = "Hola, "
	frenchTextWithPrefix  = "Bonjour, "
)

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}

	prefix := greetingPrefix(language)

	return prefix + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case french:
		prefix = frenchTextWithPrefix
	case spanish:
		prefix = spanishTextWithPrefix
	default:
		prefix = englishTextWithPrefix
	}
	return prefix
}

func main() {
	fmt.Println(Hello("", ""))
}
