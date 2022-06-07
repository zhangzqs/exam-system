package repository

import (
	"fmt"
	"testing"
)

func TestUserValid(t *testing.T) {
	uid, err := UserValid("zzq", "pwd")
	fmt.Println(uid, err)
}
