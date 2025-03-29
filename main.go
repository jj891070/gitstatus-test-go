package greet

import "fmt"

// Hello 回傳問候語句
func Hello(name string) string {
	return fmt.Sprintf("Hello, %s!", name)
}
