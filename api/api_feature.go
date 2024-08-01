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

type FeatureApiService service

func (a *FeatureApiService) ListFeatures(expDomainId int) ([]*swagger.Feature, error) {
	var ret []*swagger.Feature

	request := client.ListFeaturesRequest{
		All:      tea.Bool(true),
		DomainId: tea.String(strconv.Itoa(expDomainId)),
	}
	resonse, err := a.client.ListFeatures(&request)
	if err != nil {
		return nil, err
	}
	for _, item := range resonse.Body.Features {
		if id, err := strconv.Atoi(*item.FeatureId); err == nil {
			feature := swagger.Feature{
				FeatureId:     int32(id),
				ExpDomainId:   int32(expDomainId),
				FeatureName:   *item.Name,
				Config:        *item.Config,
				ExpDomainName: *item.DomainName,
			}

			if item.Filter != nil {
				feature.Filter = *item.Filter
			}

			if id, err := strconv.Atoi(*item.ProjectId); err == nil {
				feature.ExpProjectId = int32(id)
			}
			if id, err := strconv.Atoi(*item.ExperimentId); err == nil {
				feature.ExpId = int32(id)
			}
			if id, err := strconv.Atoi(*item.ExperimentVersionId); err == nil {
				feature.ExpVersionId = int32(id)
			}

			if item.Status != nil {
				if *item.Status == "Published" {
					feature.Status = common.FeatureStatusReleased
				} else if *item.Status == "UnPublished" {
					feature.Status = common.FeatureStatusUnreleased
				}

			}

			ret = append(ret, &feature)
		}
	}
	return ret, nil
}
