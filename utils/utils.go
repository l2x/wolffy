package utils

import "strings"

func DelEmptySlice(arr []string) []string {
	var res []string
	for _, v := range arr {
		if strings.TrimSpace(v) != "" {
			res = append(res, v)
		}
	}
	return res
}

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
