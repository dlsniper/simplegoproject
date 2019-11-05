package hello

import "fmt"

//Hi says hello!
func Hi(message string) string {
	return fmt.Sprintf("Hello World!\n%s\n", message)
}
