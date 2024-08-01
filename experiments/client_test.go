package experiments

import (
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/aliyun/aliyun-pai-ab-go-sdk/api"
	"github.com/aliyun/aliyun-pai-ab-go-sdk/model"
)

func createExperimentClient() (*ExperimentClient, error) {
	region := "cn-beijing"
	accessId := os.Getenv("ALIBABA_CLOUD_ACCESS_KEY_ID")
	accessKey := os.Getenv("ALIBABA_CLOUD_ACCESS_KEY_SECRET")
	config := api.NewConfiguration(region, accessId, accessKey)
	return NewExperimentClient(config, WithLogger(LoggerFunc(log.Printf)))
}

func TestMatchExperiment(t *testing.T) {
	client, err := createExperimentClient()
	if err != nil {
		t.Fatal(err)
	}

	experimentContext := model.ExperimentContext{
		RequestId: "pvid",
		Uid:       "156",
		FilterParams: map[string]interface{}{
			"sex": "male",
			"age": 35,
		},
	}

	experimentResult := client.MatchExperiment("DefaultProject", &experimentContext)

	fmt.Println(experimentResult.Info())
	fmt.Println(experimentResult.GetExpId())

	fmt.Println(experimentResult.GetExperimentParams().GetString("recall_v", "not_exist"))
	fmt.Println(experimentResult.GetExperimentParams().GetString("rank_v", "not_exist"))
	fmt.Println(experimentResult.GetExperimentParams().GetString("male_v", "not_exist"))

}

func TestNotMatchExperiment(t *testing.T) {

	client, err := createExperimentClient()
	if err != nil {
		t.Fatal(err)
	}
	experimentContext := model.ExperimentContext{
		RequestId: "pvid",
		Uid:       "102441809",
		FilterParams: map[string]interface{}{
			"sex": "male",
			"age": 35,
		},
	}

	experimentResult := client.MatchExperiment("test_none", &experimentContext)

	fmt.Println(experimentResult.Info())
	fmt.Println(experimentResult.GetExpId())

	t.Log(experimentResult.GetExperimentParams())
	param := experimentResult.GetExperimentParams().GetString("recall", "none")
	if param != "none" {
		t.Fatal("not match")
	}

}
