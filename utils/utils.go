package utils

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/Unknwon/com"
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
	hasher := md5.New()
	hasher.Write([]byte(s))
	return hex.EncodeToString(hasher.Sum(nil))
}

func RandInt(min int, max int) int {
	rand.Seed(time.Now().UnixNano())
	return min + rand.Intn(max-min)
}

func ClientIp(r *http.Request) string {
	ip := r.Header.Get("X-Real-Ip")
	if ip == "" {
		s := strings.Split(r.RemoteAddr, ":")
		ip = s[0]
	}
	return ip
}

func Mkdir(args ...string) error {
	for _, v := range args {
		err := os.MkdirAll(v, 0755)
		if err != nil {
			return err
		}
	}
	return nil
}

func Unzip(path, file string) error {
	_, stderr, err := com.ExecCmdDir(path, "tar", "xvf", file)
	if err != nil {
		return errors.New(err.Error() + "\n" + stderr)
	}
	return nil
}

func UnzipToFolder(path, file, folder string) error {
	_, stderr, err := com.ExecCmdDir(path, "tar", "xvf", file, "-C", folder, "--strip-components=1")
	if err != nil {
		return errors.New(err.Error() + "\n" + stderr)
	}
	return nil
}

func RunCmd(path, cmd string) error {
	_, stderr, err := com.ExecCmdDir(path, "bash", "-c", cmd)
	if err != nil {
		return errors.New(err.Error() + "\n" + stderr)
	}
	return nil
}
