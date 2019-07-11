package gawrapper

import (
	"fmt"
	"time"
)

func ComputeTOTP(secret string) int {
	t := int64(time.Now().Unix() / 30)
	code, err := ComputeCode(secret, t)
	if err != nil {
		fmt.Println(err)
	}
	return code
}
