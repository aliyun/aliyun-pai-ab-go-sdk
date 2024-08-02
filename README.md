# aliyun-pai-ab-go-sdk

Go sdk for PAI-ABTest server. Aliyun product docoment [link](https://help.aliyun.com/zh/pai/user-guide/a-b-experiment/?spm=a2c4g.11174283.0.0.3416527fVzoan7).

## Installation

```
go get github.com/aliyun/aliyun-pai-ab-go-sdk
```

## Usage

```go
   package main

import (
	"fmt"
	"log"
	"os"

	"github.com/aliyun/aliyun-pai-ab-go-sdk/api"
	"github.com/aliyun/aliyun-pai-ab-go-sdk/experiments"
	"github.com/aliyun/aliyun-pai-ab-go-sdk/model"
)

func main() {

	// init config
	region := "cn-beijing"
	accessId := os.Getenv("ALIBABA_CLOUD_ACCESS_KEY_ID")
	accessKey := os.Getenv("ALIBABA_CLOUD_ACCESS_KEY_SECRET")
	config := api.NewConfiguration(region, accessId, accessKey)

	// init client
	client, err := experiments.NewExperimentClient(config, experiments.WithLogger(experiments.LoggerFunc(log.Printf)))
	if err != nil {
		log.Fatal(err)
	}

	// set up experiment context
	experimentContext := model.ExperimentContext{
		RequestId: "pvid",
		Uid:       "157",
		FilterParams: map[string]interface{}{
			"sex": "male",
			"age": 35,
		},
	}

	// match experiment
	// DefaultProject is project name
	experimentResult := client.MatchExperiment("DefaultProject", &experimentContext)

	// print experiment info
	fmt.Println(experimentResult.Info())
	// print exp id
	fmt.Println(experimentResult.GetExpId())

	// get experiment param value
	param := experimentResult.GetExperimentParams().GetString("ab_param_name", "not_exist")
	if param != "not_exist" {
		// experiment logic

	} else {
		// default logic

	}
}

```

If you call the ABTest service in Alibaba Cloud VPC, you can specify using the VPC to connect to the service.
```go
	// init config
	region := "cn-beijing"
	accessId := os.Getenv("ALIBABA_CLOUD_ACCESS_KEY_ID")
	accessKey := os.Getenv("ALIBABA_CLOUD_ACCESS_KEY_SECRET")
	config := api.NewConfiguration(region, accessId, accessKey)
	config.UseVpc(true)
```

## Version Release Notes 
1.0.0 (2024-08-01) 
* Initial release