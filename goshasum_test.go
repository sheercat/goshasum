package main

import (
	"fmt"
	"testing"
)

func TestSumMD5(t *testing.T) {
	data := []byte("Hoge")
	sum := shasum(&data, "md5")
	sumstr := fmt.Sprintf("%x", sum)
	if sumstr != "a4b1b42ffe743732a437b6c1e2a16106" {
		t.Error("not equal", sumstr)
	}
}

func TestSumSha1(t *testing.T) {
	data := []byte("Hoge")
	sum := shasum(&data, "1")
	sumstr := fmt.Sprintf("%x", sum)
	if sumstr != "f73073ea771cea9c30a15407509b39d175052a4f" {
		t.Error("not equal", sumstr)
	}
}

func TestSumSha224(t *testing.T) {
	data := []byte("Hoge")
	sum := shasum(&data, "224")
	sumstr := fmt.Sprintf("%x", sum)
	if sumstr != "45ebed0d78503f661a4d52b8e4a89d6d7335c3c13a5a691f8ba8073c" {
		t.Error("not equal", sumstr)
	}
}

func TestSumSha256(t *testing.T) {
	data := []byte("Hoge")
	sum := shasum(&data, "256")
	sumstr := fmt.Sprintf("%x", sum)
	if sumstr != "45aac780f1de28911f9b1816e279cc62a0919c83193b8fcb4e0a6ac04f8bb444" {
		t.Error("not equal", sumstr)
	}
}

func TestSumSha384(t *testing.T) {
	data := []byte("Hoge")
	sum := shasum(&data, "384")
	sumstr := fmt.Sprintf("%x", sum)
	if sumstr != "03ddc807cefa471ed02f8482f1e95f2838314752f5a0a35cec082739a0e1315946d7f8025f84c2797bec3654d1734e5e" {
		t.Error("not equal", sumstr)
	}
}

func TestSumSha512(t *testing.T) {
	data := []byte("Hoge")
	sum := shasum(&data, "512")
	sumstr := fmt.Sprintf("%x", sum)
	if sumstr != "d6711ef5bf66c9612216a23fed7843b5b8801acd554e060cfdc52bcf13458e8e504663c9cb6a1dcf0c7d90ddfaaf5600d9efafd9e63fa79402c8a4dc7b2a75ab" {
		t.Error("not equal", sumstr)
	}
}

func TestSumSha512224(t *testing.T) {
	data := []byte("Hoge")
	sum := shasum(&data, "512224")
	sumstr := fmt.Sprintf("%x", sum)
	if sumstr != "15779a7d6249d73cf1d4c0ba93245add6763dc8409edf5d9b34a0812" {
		t.Error("not equal", sumstr)
	}
}

func TestSumSha512256(t *testing.T) {
	data := []byte("Hoge")
	sum := shasum(&data, "512256")
	sumstr := fmt.Sprintf("%x", sum)
	if sumstr != "22f81dac337eb00a685332b343e45f89fe71249e21157a3ba1acc21e3140c2ef" {
		t.Error("not equal", sumstr)
	}
}
