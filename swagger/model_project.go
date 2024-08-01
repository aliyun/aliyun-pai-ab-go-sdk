package swagger

import (
	"time"
)

type Project struct {
	// 工作空间 ID
	WorkspaceId string `json:"workspace_id,omitempty"`
	// 实验项目id
	ExpProjectId int32 `json:"exp_project_id,omitempty"`
	// 项目名称
	ProjectName string `json:"project_name,omitempty"`
	// 项目介绍
	ProjectInfo string `json:"project_info,omitempty"`
	// 创建时间
	CreateTime time.Time `json:"create_time,omitempty"`
	// 修改时间
	UpdateTime time.Time `json:"update_time,omitempty"`
}
