package mobileauth

import (
	"fmt"
	"testing"
)

func TestAuth(t *testing.T) {
	user, err := Auth("anil", "password")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(user)
}
