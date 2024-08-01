package api

import (
	"strconv"

	"github.com/alibabacloud-go/paiabtest-20240119/client"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/aliyun/aliyun-pai-ab-go-sdk/common"
	"github.com/aliyun/aliyun-pai-ab-go-sdk/swagger"
)

type DomainApiService service

func (a *DomainApiService) ListDomains(projectId int) ([]*swagger.Domain, error) {
	request := client.ListDomainsRequest{
		ProjectId: tea.String(strconv.Itoa(projectId)),
		All:       tea.Bool(true),
	}
	var ret []*swagger.Domain

	response, err := a.client.ListDomains(&request)
	if err != nil {
		return nil, err
	}

	for _, item := range response.Body.Domains {
		if id, err := strconv.Atoi(*item.DomainId); err == nil {
			domain := swagger.Domain{
				ExpDomainId:     int32(id),
				ExpProjectId:    int32(projectId),
				ExpDomainName:   *item.Name,
				IsDefaultDomain: *item.IsDefaultDomain,
				DebugUsers:      *item.DebugUsers,
				DebugCrowdIds:   *item.CrowdIds,
				LayerName:       *item.LayerName,
			}

			if item.BucketType != nil {
				if *item.BucketType == "Random" {
					domain.BucketType = common.DomainBucketTypeRand
					domain.ExperimentFlow = int32(*item.Flow)
					domain.Buckets = *item.Buckets
				} else if *item.BucketType == "Condition" {
					domain.BucketType = common.DomainBucketTypeCond
					domain.Filter = *item.Condition
				}
			}

			if id, err := strconv.Atoi(*item.LayerId); err == nil {
				domain.ExpLayerId = int32(id)
			}

			ret = append(ret, &domain)

		}
	}

	return ret, nil
}
