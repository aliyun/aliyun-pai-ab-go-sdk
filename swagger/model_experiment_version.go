package swagger

import (
	"time"
)

type ExperimentVersion struct {
	// 实验版本id
	ExpVersionId int32 `json:"exp_version_id,omitempty"`
	// 实验项目id
	ExpProjectId int32 `json:"exp_project_id,omitempty"`
	// 所属实验id
	ExpId int32 `json:"exp_id,omitempty"`
	// 实验版本名称
	ExpVersionName string `json:"exp_version_name,omitempty"`
	// 项目名称
	ProjectName string `json:"project_name,omitempty"`
	// 实验版本介绍
	ExpVersionInfo string `json:"exp_version_info,omitempty"`
	// 实验类型：1 对照组 2 实验组
	ExpVersionType int32 `json:"exp_version_type,omitempty"`
	// 流量占比，取值范围为0~100，单位%
	ExperimentFlow int32 `json:"experiment_flow,omitempty"`
	// 桶号,从实验分配而来
	Buckets string `json:"buckets,omitempty"`
	// 灰度测试用户列表
	DebugUsers string `json:"debug_users,omitempty"`
	// 灰度测试用户人群列表
	DebugCrowdIds string `json:"debug_crowd_ids,omitempty"`
	// 实验配置参数 json array 格式
	ExpVersionConfig string `json:"exp_version_config,omitempty"`
	// 创建时间
	CreateTime time.Time `json:"create_time,omitempty"`
	// 修改时间
	UpdateTime time.Time `json:"update_time,omitempty"`
}
