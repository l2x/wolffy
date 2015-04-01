package controllers

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"sync"

	"github.com/l2x/wolffy/server/config"
	"github.com/l2x/wolffy/server/models"
	"github.com/l2x/wolffy/utils/git"
	"github.com/martini-contrib/render"
)

type Deploy struct{}

func (c Deploy) Push(r render.Render, req *http.Request) {
	res := NewRes()

	id := req.URL.Query().Get("id")
	pid := req.URL.Query().Get("pid")
	commit := req.URL.Query().Get("commit")
	idint, err := strconv.Atoi(id)
	if err = RenderError(r, res, err); err != nil {
		return
	}
	pidint, err := strconv.Atoi(pid)
	if err = RenderError(r, res, err); err != nil {
		return
	}

	deploy, err := models.DeployModel.GetOne(idint)
	if err = RenderError(r, res, err); err != nil {
		return
	}

	if deploy.Status == 1 {
		RenderError(r, res, errors.New(config.ERR[config.ERR_PROJECT_DEPLOYING]))
	}

	project, err := models.ProjectModel.GetOne(pidint)
	if err = RenderError(r, res, err); err != nil {
		return
	}

	repo := git.NewRepository(config.RepoPath, project.Path)
	archiveFile, err := repo.Archive(commit, repo.Path)
	if err = RenderError(r, res, err); err != nil {
		return
	}

	go c.pushCluster(project, deploy.Id, archiveFile)

	err = models.DeployModel.UpdateStatus(deploy.Id, 1)
	if err = RenderError(r, res, err); err != nil {
		return
	}

	RenderRes(r, res, deploy)
}

func (c Deploy) pushCluster(project *models.Project, did int, archiveFile string) error {
	projectClusters, err := models.ProjectClusterModel.GetAll(project.Id)
	if err != nil {
		return errors.New(config.ERR[config.ERR_PROJECT_CLUSTER_EMPTY])
	}

	var wg sync.WaitGroup
	for _, v1 := range projectClusters {
		for _, v2 := range v1.Cluster.Machines {
			deployHistory, err := models.DeployHistoryModel.Add(did, v2.Ip)
			if err != nil {
				continue
			}

			//ip := fmt.Sprintf("http://%s:%s/pull/", v2.Ip, v2.Port)
			ip := fmt.Sprintf("http://%s:%s/pull/", v2.Ip, "8001")
			go func(id int, ip, archiveFile, pushPath, bshell, eshell string) {
				wg.Add(1)
				defer wg.Done()

				status := 2
				note := ""
				err := c.pushFile(ip, archiveFile, pushPath, bshell, eshell)
				if err != nil {
					status = 3
					note = err.Error()
				}
				models.DeployHistoryModel.Update(id, status, note)
			}(deployHistory.Id, ip, archiveFile, project.PushPath, v1.Bshell, v1.Eshell)
		}
	}
	wg.Wait()

	//os.Remove(archiveFile)
	err = models.DeployModel.UpdateStatus(did, 2)
	if err != nil {
		return err
	}

	return nil
}

func (c Deploy) pushFile(ip, archiveFile, pushPath, bshell, eshell string) error {
	fmt.Println("=====================>", ip, archiveFile, pushPath, bshell, eshell)
	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)

	//关键的一步操作
	fileWriter, err := bodyWriter.CreateFormFile("file", archiveFile)
	if err != nil {
		fmt.Println("error writing to buffer")
		return err
	}

	//打开文件句柄操作
	fh, err := os.Open(archiveFile)
	if err != nil {
		fmt.Println("error opening file")
		return err
	}
	defer fh.Close()

	//iocopy
	_, err = io.Copy(fileWriter, fh)
	if err != nil {
		return err
	}

	contentType := bodyWriter.FormDataContentType()
	bodyWriter.Close()

	u, err := url.Parse(ip)
	if err != nil {
		return err
	}
	q := u.Query()
	q.Set("bshell", bshell)
	q.Set("eshell", eshell)
	q.Set("path", pushPath)
	u.RawQuery = q.Encode()

	fmt.Println("url ===============>", u.String())
	resp, err := http.Post(u.String(), contentType, bodyBuf)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	resp_body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	fmt.Println(string(resp_body))
	return nil
}

func (c Deploy) Get(r render.Render, req *http.Request) {
	res := NewRes()

	id := req.URL.Query().Get("id")
	idint, _ := strconv.Atoi(id)

	deploy, err := models.DeployModel.GetOne(idint)
	if err = RenderError(r, res, err); err != nil {
		return
	}

	RenderRes(r, res, deploy)
}

func (c Deploy) AddTag(r render.Render, req *http.Request) {
	res := NewRes()
	tag := req.URL.Query().Get("tag")
	btag := req.URL.Query().Get("btag")
	id := req.URL.Query().Get("id")
	idint, err := strconv.Atoi(id)
	if err = RenderError(r, res, err); err != nil {
		return
	}

	project, err := models.ProjectModel.GetOne(idint)
	if err = RenderError(r, res, err); err != nil {
		return
	}

	var diff string
	if btag != "" {
		repo := git.NewRepository(config.RepoPath, project.Path)
		diff, err = repo.Diff(tag, btag)
		if err = RenderError(r, res, err); err != nil {
			return
		}
	}

	deploy, err := models.DeployModel.Add(idint, tag, diff)
	if err = RenderError(r, res, err); err != nil {
		return
	}

	RenderRes(r, res, deploy)
}

func (c Deploy) History(r render.Render, req *http.Request) {
	res := NewRes()

	pid := req.URL.Query().Get("id")
	pidint, err := strconv.Atoi(pid)
	if err = RenderError(r, res, err); err != nil {
		return
	}

	deploys, err := models.DeployModel.GetAll(pidint, 15)
	if err = RenderError(r, res, err); err != nil {
		return
	}

	// get deploy history
	for k, _ := range deploys {
		deploys[k].Diff = ""
	}

	RenderRes(r, res, deploys)
}

func (c Deploy) GetDiff(r render.Render, req *http.Request) {
	res := NewRes()
	id := req.URL.Query().Get("id")
	idint, err := strconv.Atoi(id)
	if err = RenderError(r, res, err); err != nil {
		return
	}

	deploy, err := models.DeployModel.GetOne(idint)
	if err = RenderError(r, res, err); err != nil {
		return
	}

	RenderRes(r, res, deploy)
}

func (c Deploy) HistoryDetail(r render.Render, req *http.Request) {
	res := NewRes()
	id := req.URL.Query().Get("id")
	idint, err := strconv.Atoi(id)
	if err = RenderError(r, res, err); err != nil {
		return
	}

	detail, err := models.DeployHistoryModel.GetAll(idint)
	if err = RenderError(r, res, err); err != nil {
		return
	}

	RenderRes(r, res, detail)
}
