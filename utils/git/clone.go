package git

import (
	"errors"
	"strings"

	"github.com/Unknwon/com"
)

func (repo *Repository) Clone() (string, error) {
	stdout, stderr, err := com.ExecCmdDir(repo.Path, "git", "clone", repo.RemotePath)
	if err != nil {
		return "", errors.New(stderr)
	}
	return strings.Split(stdout, " ")[0], nil
}
