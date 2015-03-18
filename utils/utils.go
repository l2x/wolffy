package utils

import (
	"crypto/md5"
	"io"
	"math/rand"
	"strings"
	"time"
)

type StringReverse []string

func (s StringReverse) Len() int {
	return len(s)
}

func (s StringReverse) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s StringReverse) Less(i, j int) bool {
	return s[i] > s[j]
}

func Compress(path string, spath string) error {

	return nil
}

func DelEmptySlice(arr []string) []string {
	var res []string
	for _, v := range arr {
		if strings.TrimSpace(v) != "" {
			res = append(res, v)
		}
	}
	return res
}

func Md5(s string) string {
	h := md5.New()
	io.WriteString(h, s)
	return string(h.Sum(nil))
}

func RandInt(min int, max int) int {
	rand.Seed(time.Now().UnixNano())
	return min + rand.Intn(max-min)
}
