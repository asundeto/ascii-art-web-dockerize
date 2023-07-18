package ascart

import (
	"crypto/md5"
	"fmt"
)

func HashMD5(content string) string {
	h := md5.Sum([]byte(content))
	return fmt.Sprintf("%x", h)
}

func CheckHash(hash string) bool {
	hashStandard := "ac85e83127e49ec42487f272d9b9db8b"
	hashShadow := "a49d5fcb0d5c59b2e77674aa3ab8bbb1"
	hashThinkertoy := "cd9ba1cc7a1b626147cf6729ad0c6857"
	if hash == hashStandard {
		return true
	}
	if hash == hashShadow {
		return true
	}
	if hash == hashThinkertoy {
		return true
	}
	return false
}
