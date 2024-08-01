package swagger

import (
	"time"
)

type Datasource struct {
	// 数据源id
	DatasourceId int32 `json:"datasource_id,omitempty"`
	// 阿里云ID
	AliyunId string `json:"aliyun_id,omitempty"`
	// 数据源类型,maxcompute/hologres
	Type_ string `json:"type,omitempty"`
	// 数据源名称
	Name string `json:"name,omitempty"`
	// 地域
	Region string `json:"region,omitempty"`
	// vpc 地址
	VpcAddress string `json:"vpc_address,omitempty"`
	// 项目名称
	Project string `json:"project,omitempty"`
	// 数据库名称
	Database string `json:"database,omitempty"`
	// 创建时间
	CreateTime time.Time `json:"create_time,omitempty"`
	// 修改时间
	UpdateTime time.Time `json:"update_time,omitempty"`
}
