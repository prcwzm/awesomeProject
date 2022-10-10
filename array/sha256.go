package array

import (
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
)

type shaType int

const (
	SHA256 = iota
	SHA384
	SHA512
)

func HashSHA(s []byte, conf shaType) {
	if conf == SHA512 {
		fmt.Printf("%x\n", sha512.Sum512(s))
	} else if conf == SHA384 {
		fmt.Printf("%x\n", sha512.Sum384(s))
	} else if conf == SHA256 {
		fmt.Printf("%x\n", sha256.Sum256(s))
	}
}
