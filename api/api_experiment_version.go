package api

import (
	"context"
	"strconv"

	"github.com/alibabacloud-go/paiabtest-20240119/client"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/aliyun/aliyun-pai-ab-go-sdk/common"
	"github.com/aliyun/aliyun-pai-ab-go-sdk/swagger"
)

// Linger please
var (
	_ context.Context
)

type ExperimentVersionApiService service

func (a *ExperimentVersionApiService) ListExperimentVersions(expId int) ([]*swagger.ExperimentVersion, error) {
	var ret []*swagger.ExperimentVersion

	request := client.ListExperimentVersionsRequest{
		All:          tea.Bool(true),
		ExperimentId: tea.String(strconv.Itoa(expId)),
	}
	response, err := a.client.ListExperimentVersions(&request)
	if err != nil {
		return nil, err
	}

	for _, item := range response.Body.ExperimentVersions {
		if id, err := strconv.Atoi(*item.ExperimentVersionId); err == nil {
			experimentVersion := swagger.ExperimentVersion{
				ExpVersionId:     int32(id),
				ExpId:            int32(expId),
				ExpVersionName:   *item.Name,
				ExperimentFlow:   *item.Flow,
				Buckets:          *item.Buckets,
				ExpVersionConfig: *item.Config,
				DebugUsers:       *item.DebugUsers,
				DebugCrowdIds:    *item.CrowdIds,
			}

			if item.Type != nil {
				if *item.Type == "Baseline" {
					experimentVersion.ExpVersionType = common.ExpVersionTypeBase
				} else if *item.Type == "Normal" {
					experimentVersion.ExpVersionType = common.ExpVersionTypeNormal
				}

			}

			ret = append(ret, &experimentVersion)
		}
	}
	return ret, nil
}
