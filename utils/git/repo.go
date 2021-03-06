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
	repo.Path = fmt.Sprintf("%s/%s", strings.TrimRight(repo.BasePath, "/"), repo.RemotePath)

	return repo
}

// get repository name
func (repo *Repository) GetName() string {
	parts := strings.Split(repo.RemotePath, "/")
	return strings.TrimRight(parts[len(parts)-1], ".git")
}

// clone repository
func (repo *Repository) Clone() (string, error) {
	stdout, stderr, err := com.ExecCmdDir(repo.BasePath, "git", "clone", repo.RemotePath, repo.RemotePath)
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
	tags, _ := repo.GetTags()
	for _, v := range tags {
		repo.DelTags(v)
	}

	_, stderr, err := com.ExecCmdDir(repo.Path, "git", "fetch", "--tags")
	if err != nil {
		return errors.New(stderr)
	}
	return nil
}

//delete local tags
func (repo *Repository) DelTags(commit string) error {
	_, stderr, err := com.ExecCmdDir(repo.Path, "git", "tag", "-d", commit)
	if err != nil {
		return errors.New(err.Error() + "\n" + stderr)
	}
	return nil
}

//archive
//git archive --format=tar.gz --prefix=wolffy-0.1.3/ 0.1.3 --output /tmp/wolffy-0.1.3.tag.gz
func (repo *Repository) Archive(commit, spath string) (string, error) {
	prefix := fmt.Sprintf("%s-%s/", repo.Name, commit)
	spath = fmt.Sprintf("%s-%s%s", spath, commit, ".tar.gz")

	_, stderr, err := com.ExecCmdDir(repo.Path, "git", "archive", "--format", "tar.gz", "--prefix", prefix, commit, "--output", spath)

	if err != nil {
		return "", errors.New(err.Error() + "\n" + stderr)
	}
	return spath, nil
}

// diff
func (repo *Repository) Diff(commita, commitb string) (string, error) {
	stdout, stderr, err := com.ExecCmdDir(repo.Path, "git", "diff", commita, commitb)
	if err != nil {
		return "", errors.New(stderr)
	}

	return stdout, nil
}
