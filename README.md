# aliyun-pai-ab-go-sdk

Go sdk for PAI-ABTest server. Aliyun product docoment [link](https://help.aliyun.com/zh/pai/user-guide/a-b-experiment/?spm=a2c4g.11174283.0.0.3416527fVzoan7).

## Installation

```
go get github.com/aliyun/aliyun-pai-ab-go-sdk
```

## Usage

```go
    // init config
	region := "cn-beijing"
	accessId := os.Getenv("ALIBABA_CLOUD_ACCESS_KEY_ID")
	accessKey := os.Getenv("ALIBABA_CLOUD_ACCESS_KEY_SECRET")
	config := api.NewConfiguration(region, accessId, accessKey)

    // init client
	client := NewExperimentClient(config, WithLogger(LoggerFunc(log.Printf)))

    // set up experiment context
	experimentContext := model.ExperimentContext{
		RequestId: "pvid",
		Uid:       "156",
		FilterParams: map[string]interface{}{
			"sex": "male",
			"age": 35,
		},
	}

    // match experiment
	experimentResult := client.MatchExperiment("DefaultProject", &experimentContext)

    // print experiment info
	fmt.Println(experimentResult.Info())
    // print exp id
	fmt.Println(experimentResult.GetExpId())

    // get experiment param value
	param := experimentResult.GetExperimentParams().GetString("ab_param_name", "not_exist"))
    if param != "not_exist" {
    // experiment logic 
        

    } else {
    // default logic

    }

```

## Version Release Notes 
1.0.0 (2024-08-01) 
* Initial release