package swagger

import (
	"time"
)

type Feature struct {
	// 功能 id
	FeatureId int32 `json:"feature_id,omitempty"`
	// 实验项目id
	ExpProjectId int32 `json:"exp_project_id,omitempty"`
	// 所在的生效域 ID
	ExpDomainId int32 `json:"exp_domain_id,omitempty"`
	// 来源实验 id
	ExpId int32 `json:"exp_id,omitempty"`
	// 来源版本 id
	ExpVersionId int32 `json:"exp_version_id,omitempty"`
	// 功能名称
	FeatureName string `json:"feature_name,omitempty"`
	// 项目名称
	ProjectName string `json:"project_name,omitempty"`
	// 实验域名称
	ExpDomainName string `json:"exp_domain_name,omitempty"`
	// 相关实验负责人
	ExpOwner string `json:"exp_owner,omitempty"`
	// 过滤条件
	Filter string `json:"filter,omitempty"`
	// 实验配置参数 json array 格式
	Config string `json:"config,omitempty"`
	// 状态 1: 未发布，2：已发布
	Status int32 `json:"status,omitempty"`
	// 发布时间
	ReleaseTime time.Time `json:"release_time,omitempty"`
	// 创建时间
	CreateTime time.Time `json:"create_time,omitempty"`
	// 修改时间
	UpdateTime time.Time `json:"update_time,omitempty"`
}
