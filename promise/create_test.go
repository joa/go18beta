package promise

import (
	"fmt"
	"testing"
)

func TestCreate(t *testing.T) {
	fmt.Println(Create[string]())
}
