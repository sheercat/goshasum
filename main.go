package main

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

var doHelp = flag.Bool("h", false, "help")
var algorithm = flag.String("a", "1", "algorithm md5, 1, 224, 256, 384, 512, 512224, 512256 (default 1).")

func help() {
	flag.Usage()
}
func main() {
	flag.Parse()
	if *doHelp {
		help()
		return
	}
	file := os.Args[len(os.Args)-1]
	data, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println(err)
		return
	}

	switch strings.ToLower(*algorithm) {
	case "md5":
		fmt.Printf("%x\n", md5.Sum(data))
	case "1":
		fmt.Printf("%x\n", sha1.Sum(data))
	case "224":
		fmt.Printf("%x\n", sha256.Sum224(data))
	case "256":
		fmt.Printf("%x\n", sha256.Sum256(data))
	case "384":
		fmt.Printf("%x\n", sha512.Sum384(data))
	case "512":
		fmt.Printf("%x\n", sha512.Sum512(data))
	case "512224":
		fmt.Printf("%x\n", sha512.Sum512_224(data))
	case "512256":
		fmt.Printf("%x\n", sha512.Sum512_256(data))
	}
}
