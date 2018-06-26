package main

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"strings"
)

var (
	doHelp    = flag.Bool("h", false, "help")
	isIgnore  = flag.Bool("i", false, "ignore print filename")
	algorithm = flag.String("a", "1", "algorithm md5, 1, 224, 256, 384, 512, 512224, 512256")
)

func main() {
	flag.Parse()
	if *doHelp {
		flag.Usage()
		return
	}

	for _, file := range flag.Args() {
		data, err := ioutil.ReadFile(file)
		if err != nil {
			fmt.Println(err)
			continue
		}
		sum, err := shasum(&data, *algorithm)
		if err != nil {
			fmt.Println(err)
			continue
		}
		if *isIgnore {
			fmt.Printf("%x", sum)
		} else {
			fmt.Printf("%x\t%s\n", sum, file)
		}
	}
}

func shasum(data *[]byte, algorithm string) (interface{}, error) {
	var sum interface{}
	switch strings.ToLower(algorithm) {
	case "md5":
		sum = md5.Sum(*data)
	case "1":
		sum = sha1.Sum(*data)
	case "224":
		sum = sha256.Sum224(*data)
	case "256":
		sum = sha256.Sum256(*data)
	case "384":
		sum = sha512.Sum384(*data)
	case "512":
		sum = sha512.Sum512(*data)
	case "512224":
		sum = sha512.Sum512_224(*data)
	case "512256":
		sum = sha512.Sum512_256(*data)
	default:
		return nil, errors.New("unsupported algorithm")
	}
	return sum, nil
}
