package core


import "fmt"

func Greet(name string) string {
	return fmt.Sprintf("%s, 你好！", name)
}