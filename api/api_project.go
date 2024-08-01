package api

import (
	"strconv"

	"github.com/alibabacloud-go/paiabtest-20240119/client"
	"github.com/aliyun/aliyun-pai-ab-go-sdk/swagger"
)

type ProjectApiService service

func (a *ProjectApiService) ListProjects() ([]*swagger.Project, error) {
	request := client.ListProjectsRequest{}
	request.SetAll(true)
	response, err := a.client.ListProjects(&request)
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
