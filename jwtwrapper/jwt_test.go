package jwtwrapper

import (
	"fmt"
	"testing"
	"time"
)

func TestIssueAndGet(t *testing.T) {
	JWTSecret := "xihaijie"
	c := map[string]interface{}{
		"foo": "bar",
	}

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := IssueTokenStr(c, JWTSecret)

	fmt.Println(tokenString, err)

	cGot, err := GetMapClaims(tokenString, JWTSecret)
	fmt.Printf("%v,%v\n", cGot, err)

	if v := cGot["foo"]; v != "bar" {
		t.Error("Expected bar, got", v)
	}
}

func TestIssueWithExp(t *testing.T) {
	JWTSecret := "xihaijie"
	c := map[string]interface{}{
		"foo": "bar",
	}

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := IssueTokenStrWithExp(c, JWTSecret, 1)

	fmt.Println(tokenString, err)

	cGot, err := GetMapClaims(tokenString, JWTSecret)
	fmt.Printf("%v,%v\n", cGot, err)

	if v := cGot["foo"]; v != "bar" {
		t.Error("Expected bar, got", v)
	}

	// has to be exp more than one sec
	time.Sleep(2 * time.Second)
	cGot, err = GetMapClaims(tokenString, JWTSecret)

	fmt.Printf("%v,%v\n", cGot, err)
	if err == nil {
		t.Error("Expected err, got nil")
	}

}
