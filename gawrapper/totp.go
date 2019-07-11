package gawrapper

import (
	"fmt"
	"strconv"
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

func VerifyTOTP(secret string, codeStr string) (bool, error) {
	t := int64(time.Now().Unix() / 30)

	yourCode, err := strconv.Atoi(codeStr)
	if err != nil {
		return false, err
	}

	myCode, err := ComputeCode(secret, t)
	if err != nil {
		return false, err
	}

	if yourCode == myCode {
		return true, nil
	}

	myCodeOld, err := ComputeCode(secret, t-1)
	if err != nil {
		return false, err
	}

	if yourCode == myCodeOld {
		return true, nil
	}

	return false, nil
}
