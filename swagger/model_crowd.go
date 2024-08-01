package swagger

import (
	"time"
)

type Crowd struct {
	// 工作空间 ID
	WorkspaceId string `json:"workspace_id,omitempty"`
	// 人群id
	CrowdId int32 `json:"crowd_id,omitempty"`
	// 人群名称
	CrowdName string `json:"crowd_name,omitempty"`
	// 人群描述
	CrowdInfo string `json:"crowd_info,omitempty"`
	// 人员数量
	Quantity int32 `json:"quantity,omitempty"`
	// 标签
	Label string `json:"label,omitempty"`
	// 用户id列表
	Users string `json:"users,omitempty"`
	// 创建时间
	CreateTime time.Time `json:"create_time,omitempty"`
	// 修改时间
	UpdateTime time.Time `json:"update_time,omitempty"`
}
