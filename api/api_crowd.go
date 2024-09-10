package api

import (
	"context"
	"strconv"
	"strings"

	"github.com/alibabacloud-go/paiabtest-20240119/client"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
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

	runtime := &util.RuntimeOptions{
		IgnoreSSL: tea.Bool(true),
	}
	headers := make(map[string]*string)
	response, err := a.client.ListCrowdsWithOptions(&request, headers, runtime)
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
