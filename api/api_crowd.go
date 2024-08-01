package api

import (
	"context"
	"strconv"
	"strings"

	"github.com/alibabacloud-go/paiabtest-20240119/client"
	"github.com/alibabacloud-go/tea/tea"
)

// Linger please
var (
	_ context.Context
)

type CrowdApiService service

func (a *CrowdApiService) ListCrowdUsers(crowdId int) (users []string, err error) {
	request := client.ListCrowdsRequest{
		All:     tea.Bool(true),
		CrowdId: tea.String(strconv.Itoa(crowdId)),
	}

	response, err := a.client.ListCrowds(&request)
	if err != nil {
		return nil, err
	}

	for _, crowd := range response.Body.Crowds {
		if crowd.Users != nil && *crowd.Users != "" {
			list := strings.Split(*crowd.Users, ",")
			for _, user := range list {
				if user != "" {
					users = append(users, user)
				}
			}
		}
	}

	return
}
