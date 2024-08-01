package swagger

import (
	"time"
)

type Experiment struct {
	// 工作空间 ID
	WorkspaceId string `json:"workspace_id,omitempty"`
	// 实验id
	ExpId int32 `json:"exp_id,omitempty"`
	// 实验项目id
	ExpProjectId int32 `json:"exp_project_id,omitempty"`
	// 实验域id
	ExpDomainId int32 `json:"exp_domain_id,omitempty"`
	// 所在实验层id
	ExpLayerId int32 `json:"exp_layer_id,omitempty"`
	// 实验名称
	ExpName string `json:"exp_name,omitempty"`
	// 项目名称
	ProjectName string `json:"project_name,omitempty"`
	// 实验域名称
	ExpDomainName string `json:"exp_domain_name,omitempty"`
	// 层名称，校验规则：[a-zA-Z][a-zA-Z1-9-]+
	LayerName string `json:"layer_name,omitempty"`
	// 实验介绍
	ExpInfo string `json:"exp_info,omitempty"`
	// 实验负责人
	Owner string `json:"owner,omitempty"`
	// 灰度测试用户列表
	DebugUsers string `json:"debug_users,omitempty"`
	// 灰度测试用户人群列表
	DebugCrowdIds string `json:"debug_crowd_ids,omitempty"`
	// 分流类型，1：随机流量，2：根据条件过滤，默认1
	BucketType int32 `json:"bucket_type,omitempty"`
	// 流量占比，取值范围为0~100，单位%
	ExperimentFlow int32 `json:"experiment_flow,omitempty"`
	// 桶号,从实验层分配而来
	Buckets string `json:"buckets,omitempty"`
	// 过滤条件
	Filter string `json:"filter,omitempty"`
	// 启动时间
	StartTime string `json:"start_time,omitempty"`
	// 结束时间
	EndTime string `json:"end_time,omitempty"`
	// 状态 1: 停止，2 ：运行中， 3： 已推全
	Status int32 `json:"status,omitempty"`
	// 创建时间
	CreateTime time.Time `json:"create_time,omitempty"`
	// 修改时间
	UpdateTime time.Time `json:"update_time,omitempty"`
}
