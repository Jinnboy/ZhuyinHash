// github.com/Jinnboy/ZhuyinHash

package zhuyinhash

import (
	"fmt"
	"io/ioutil"
)

//CJK unicode的注音編碼
var _zy0 [0x312F - 0x3100 + 1]int16
var _zy1 [0x9FFF - 0x3400 + 1]int16
var _zy2 [0x2B81F - 0x20000 + 1]int16
var _zy3 [0x2FA1F - 0x2F800 + 1]int16

func init() {
	LoadZhuyin()
}

func LoadZhuyin() {
	b, err := ioutil.ReadFile("txt/zhuyin.bin")
	if err != nil {
		fmt.Println(err)
		return
	}
	i := 0
	for j := 0; j < len(_zy0); j++ {
		n := uint16(b[i]) | uint16(b[i+1])<<8
		i += 2
		_zy0[j] = int16(n)
	}
	for j := 0; j < len(_zy1); j++ {
		n := uint16(b[i]) | uint16(b[i+1])<<8
		i += 2
		_zy1[j] = int16(n)
	}
	for j := 0; j < len(_zy2); j++ {
		n := uint16(b[i]) | uint16(b[i+1])<<8
		i += 2
		_zy2[j] = int16(n)
	}
	for j := 0; j < len(_zy3); j++ {
		n := uint16(b[i]) | uint16(b[i+1])<<8
		i += 2
		_zy3[j] = int16(n)
	}
}

func HashRune(r rune) int16 {
	if r <= 0x19 {
		return 0
	} else if r >= 0x20 && r <= 0x60 { //ASCII 空白(0x20) 、(0x60)
		return int16(r - 0x1F)
	} else if r >= 0x61 && r <= 0x7A { //a~z對應到A~Z
		return int16(r - 0x3F)
	} else if r >= 0x7B && r <= 0x7E { //ASCII {(0x7B) ~(0x7E)
		return int16(r - 0x39)
	} else if r >= 0x3100 && r <= 0x312F { //注音Bopomofo
		return _zy0[r-0x3100]
	} else if r >= 0x3400 && r <= 0x9FFF { //CJK unicode
		return _zy1[r-0x3400]
	} else if r >= 0x20000 && r <= 0x2B81F { //CJK unicode
		return _zy2[r-0x20000]
	} else if r >= 0x2f800 && r <= 0x2FA1F { //CJK unicode
		return _zy3[r-0x2f800]
	} else if r >= 0x3040 && r <= 0x30F9 { //日文字母 平假名和片假名
		return 498
	} else {
		return 499
	}
}

func Hash(s string) int64 {
	rs := []rune(s)
	length := len(rs)
	if length > 7 {
		length = 7
	}
	var n uint64 = 0
	for i := 0; i < length; i++ {
		r := rs[i]
		m := uint64(HashRune(r))
		n |= m << uint64((6-i)*9)
	}
	return int64(n)
}
