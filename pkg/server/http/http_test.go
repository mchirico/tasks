package http

import (
	"fmt"
	"os"
	"testing"
)

func TestPrep(t *testing.T) {
	fmt.Println("nothing yet..")
	r := os.Getenv("TASK_KEYa")
	if r == "" {
		fmt.Println("empty")
	}
	fmt.Println(r)
}
