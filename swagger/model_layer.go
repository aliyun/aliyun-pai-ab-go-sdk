package swagger

import (
	"time"
)

type Layer struct {
	// 工作空间 ID
	WorkspaceId string `json:"workspace_id,omitempty"`
	// 实验层id
	ExpLayerId int32 `json:"exp_layer_id,omitempty"`
	// 所属实验域id
	ExpDomainId int32 `json:"exp_domain_id,omitempty"`
	// 实验项目id
	ExpProjectId int32 `json:"exp_project_id,omitempty"`
	// 是否是默认层
	IsDefaultLayer bool `json:"is_default_layer,omitempty"`
	// 层名称，校验规则：[a-zA-Z][a-zA-Z1-9-]+
	LayerName string `json:"layer_name,omitempty"`
	// 项目名称
	ProjectName string `json:"project_name,omitempty"`
	// 实验域名称
	ExpDomainName string `json:"exp_domain_name,omitempty"`
	// 层说明
	LayerInfo string `json:"layer_info,omitempty"`
	// 创建时间
	CreateTime time.Time `json:"create_time,omitempty"`
	// 修改时间
	UpdateTime time.Time `json:"update_time,omitempty"`
}
