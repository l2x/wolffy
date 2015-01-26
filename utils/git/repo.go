// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package git

import (
	"errors"
	"fmt"
	"sort"
	"strings"

	"github.com/Unknwon/com"
	"github.com/l2x/wolffy/utils"
)

// Repository represents a Git repository.
type Repository struct {
	BasePath   string
	Path       string
	RemotePath string
	Name       string
}

func NewRepository(basePath, remotePath string) *Repository {
	repo := &Repository{
		BasePath:   basePath,
		RemotePath: remotePath,
	}
	repo.Name = repo.GetName()
	repo.Path = fmt.Sprintf("%s/%s", strings.TrimRight(repo.BasePath, "/"), repo.Name)
	fmt.Println(repo)

	return repo
}

// get repository name
func (repo *Repository) GetName() string {
	parts := strings.Split(repo.RemotePath, "/")
	return strings.TrimRight(parts[len(parts)-1], ".git")
}

// clone repository
func (repo *Repository) Clone() (string, error) {
	stdout, stderr, err := com.ExecCmdDir(repo.BasePath, "git", "clone", repo.RemotePath, repo.Name)
	if err != nil {
		return "", errors.New(stderr)
	}
	return strings.Split(stdout, " ")[0], nil
}

// list all branchs
func (repo *Repository) GetBranches() ([]string, error) {
	stdout, stderr, err := com.ExecCmdDir(repo.Path, "git", "ls-remote", "--heads", "origin")

	if err != nil {
		return nil, errors.New(stderr)
	}

	var branchs []string
	infos := strings.Split(stdout, "\n")
	for _, v := range infos {
		v = strings.Replace(v, "\t", " ", -1)
		parts := strings.Split(v, " ")
		if len(parts) != 2 {
			continue
		}
		branchs = append(branchs, strings.TrimPrefix(parts[1], "refs/heads/"))
	}
	return branchs, nil
}

// list all tags
func (repo *Repository) GetTags() ([]string, error) {
	stdout, stderr, err := com.ExecCmdDir(repo.Path, "git", "tag", "-l")
	if err != nil {
		return nil, errors.New(stderr)
	}

	tags := strings.Split(stdout, "\n")
	tags = utils.DelEmptySlice(tags)
	sort.Sort(utils.StringReverse(tags))

	return tags, nil
}

// pull tags
func (repo *Repository) PullTags() error {
	_, stderr, err := com.ExecCmdDir(repo.Path, "git", "pull", "--tags")
	if err != nil {
		return errors.New(stderr)
	}
	return nil
}

//archive
func (repo *Repository) Archive(commit, spath string) error {
	_, stderr, err := com.ExecCmdDir(repo.Path, "git", "archive", "--format", "zip", "--output", spath, commit, "-9")
	if err != nil {
		return errors.New(stderr)
	}
	return nil
}

// diff
func (repo *Repository) Diff(commit1, commit2 string) (string, error) {
	fmt.Println(commit1, commit2)

	stdout, stderr, err := com.ExecCmdDir(repo.Path, "git", "diff", commit1, commit2)
	if err != nil {
		return "", errors.New(stderr)
	}

	return stdout, nil
}
