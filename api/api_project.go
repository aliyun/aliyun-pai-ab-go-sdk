package api

import (
	"strconv"

	"github.com/alibabacloud-go/paiabtest-20240119/client"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/aliyun/aliyun-pai-ab-go-sdk/swagger"
)

type ProjectApiService service

func (a *ProjectApiService) ListProjects() ([]*swagger.Project, error) {
	request := client.ListProjectsRequest{}
	request.SetAll(true)
	runtime := &util.RuntimeOptions{
		IgnoreSSL: tea.Bool(true),
	}
	headers := make(map[string]*string)
	response, err := a.client.ListProjectsWithOptions(&request, headers, runtime)
	var ret []*swagger.Project
	if err != nil {
		return ret, err
	}

	for _, item := range response.Body.Projects {
		if id, err := strconv.Atoi(*item.ProjectId); err == nil {
			project := &swagger.Project{
				ProjectName: *item.Name,
				ProjectInfo: *item.Description,
			}
			project.ExpProjectId = int32(id)
			project.WorkspaceId = *item.WorkspaceId
			ret = append(ret, project)
		}
	}
	return ret, nil
}
