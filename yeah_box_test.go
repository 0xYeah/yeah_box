package yeahBox

import (
	"fmt"
	"testing"
)

func TestMethods(t *testing.T) {
	vInfo := GetVersion()
	fmt.Printf("Version: [%s]\n", vInfo)
}
