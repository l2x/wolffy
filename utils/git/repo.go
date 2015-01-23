// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package git

import (
	"errors"
	"strings"

	"github.com/Unknwon/com"
)

// Repository represents a Git repository.
type Repository struct {
	Path       string
	RemotePath string
	Name       string
}

func NewRepository(path, remotePath string) *Repository {
	repo := &Repository{
		Path:       path,
		RemotePath: remotePath,
	}
	repo.Name = repo.GetName()

	return repo
}

func (repo *Repository) Clone() (string, error) {
	stdout, stderr, err := com.ExecCmdDir(repo.Path, "git", "clone", repo.RemotePath, repo.Name)
	if err != nil {
		return "", errors.New(stderr)
	}
	return strings.Split(stdout, " ")[0], nil
}

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

func (repo *Repository) GetName() string {
	parts := strings.Split(repo.RemotePath, "/")
	return strings.TrimRight(parts[len(parts)-1], ".git")
}

func (repo *Repository) GetTags() []string {
}
