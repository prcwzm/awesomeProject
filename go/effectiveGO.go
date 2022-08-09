package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

type ByteSize float64
type ByteSlice []byte

const (
	_           = iota
	KB ByteSize = 1 << (10 * iota)
	MB
	GB
	TB
	PB
	EB
	ZB
	YB
)

var (
	home   = os.Getenv("HOME")
	user   = os.Getenv("USER")
	gopath = os.Getenv("GOPATH")
)

func init() {
	if user == "" {
		log.Fatal("$USER not set")
	}
	if home == "" {
		home = "/home/" + user
	}
	if gopath == "" {
		gopath = home + "/go"
	}
	fmt.Printf("%s, %s, %s \n", user, home, gopath)
	flag.StringVar(&gopath, "gopath", gopath, "override default GOPATH")
}

func main() {
	var b ByteSlice = make([]byte, 2)
	fmt.Fprintf(&b, "This hour has %d days\n", 7)
	fmt.Println(string(b))
}

func (p *ByteSlice) Append(data []byte) {
	slice := *p
	l := len(slice)
	if l+len(data) > cap(slice) { // reallocate
		// Allocate double what's needed, for future growth.
		newSlice := make([]byte, (l+len(data))*2)
		// The copy function is predeclared and works for any slice type.
		copy(newSlice, slice)
		slice = newSlice
	}
	slice = slice[0 : l+len(data)]
	copy(slice[l:], data)
	*p = slice
}

func (p *ByteSlice) Write(data []byte) (n int, err error) {
	slice := *p
	slice.Append(data)
	*p = slice
	return len(data), nil
}

func (b ByteSize) String() string {
	switch {
	case b >= YB:
		return fmt.Sprintf("%.2fYB", b/YB)
	case b >= ZB:
		return fmt.Sprintf("%.2fZB", b/ZB)
	case b >= EB:
		return fmt.Sprintf("%.2fEB", b/EB)
	case b >= PB:
		return fmt.Sprintf("%.2fPB", b/PB)
	case b >= TB:
		return fmt.Sprintf("%.2fTB", b/TB)
	case b >= GB:
		return fmt.Sprintf("%.2fGB", b/GB)
	case b >= MB:
		return fmt.Sprintf("%.2fMB", b/MB)
	case b >= KB:
		return fmt.Sprintf("%.2fKB", b/KB)
	}
	return fmt.Sprintf("%.2fB", b)
}
