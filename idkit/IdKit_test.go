package idkit

import (
	"testing"
	"fmt"
)

func TestCreateUniqueId(t *testing.T) {
	fmt.Printf("Unique Id %s\n", CreateUniqueId())
}

func TestCreateMd5(t *testing.T) {
	fmt.Printf("hello go md5 %s\n", CreateMd5("hello go"))
}
