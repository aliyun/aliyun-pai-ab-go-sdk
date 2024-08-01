package experiments

import (
	"crypto/md5"
	"fmt"
	"hash/fnv"
	"strconv"

	"github.com/aliyun/aliyun-pai-ab-go-sdk/api"
	"github.com/aliyun/aliyun-pai-ab-go-sdk/common"
	"github.com/aliyun/aliyun-pai-ab-go-sdk/model"
)

type ClientOption func(c *ExperimentClient)

func WithLogger(l Logger) ClientOption {
	return func(e *ExperimentClient) {
		e.Logger = l
	}
}

func WithErrorLogger(l Logger) ClientOption {
	return func(e *ExperimentClient) {
		e.ErrorLogger = l
	}
}

type ExperimentClient struct {
	// Host A/B Test server host
	Host string

	// APIClient invoke api to connect to the A/B Test Server
	APIClient *api.APIClient

	projectMap map[string]*model.Project

	// Logger specifies a logger used to report internal changes within the writer
	Logger Logger

	// ErrorLogger is the logger to report errors
	ErrorLogger Logger
}

func NewExperimentClient(config *api.Configuration, opts ...ClientOption) (*ExperimentClient, error) {
	client := ExperimentClient{
		projectMap: make(map[string]*model.Project, 0),
	}

	for _, opt := range opts {
		opt(&client)
	}

	if err := client.Validate(); err != nil {
		return nil, err
	}

	apiClient, err := api.NewAPIClient(config)
	if err != nil {
		return nil, err
	}

	client.APIClient = apiClient

	client.LoadExperimentData()

	go client.loopLoadExperimentData()

	return &client, nil
}

// Validate check the  ExperimentClient value
func (e *ExperimentClient) Validate() error {

	return nil
}

// MatchExperiment specifies to find match experiment by the ExperimentContext
// If not find the scene return error or return ExperimentResult
func (e *ExperimentClient) MatchExperiment(projectName string, experimentContext *model.ExperimentContext) *model.ExperimentResult {
	projectData := e.projectMap
	project, exist := projectData[projectName]
	if !exist {
		e.logError(fmt.Errorf("project:%s, not found the project", projectName))
		return model.NewExperimentResult(projectName, experimentContext, nil)
	}

	experimentResult := model.NewExperimentResult(projectName, experimentContext, project)

	e.matchDomain(project.DefaultDomain(), experimentResult)
	experimentResult.Init()
	return experimentResult
}

// matchDomain
func (e *ExperimentClient) matchDomain(domain *model.Domain, experimentResult *model.ExperimentResult) {
	if domain == nil {
		return
	}

	for _, feature := range domain.Features() {
		if feature.Match(experimentResult.ExperimentContext) {
			experimentResult.GetExperimentParams().AddParams(feature.Params())
		}
	}

	for _, layer := range domain.Layers() {
		e.matchLayer(layer, experimentResult)

	}
}

func (e *ExperimentClient) matchLayer(layer *model.Layer, experimentResult *model.ExperimentResult) {
	if layer == nil {
		return
	}

	// first find the debug user match
	for _, experiment := range layer.Experiments() {
		if experiment.MatchDebugUsers(experimentResult.ExperimentContext) {
			e.logInfo(fmt.Sprintf("match debug user for experiment:%s", experiment.ExpName))
			e.matchExperiment(experiment, experimentResult)
			return
		}
	}

	for _, domain := range layer.Domains() {
		if domain.MatchDebugUsers(experimentResult.ExperimentContext) {
			e.logInfo(fmt.Sprintf("match debug user for domain:%s", domain.ExpDomainName))
			e.matchDomain(domain, experimentResult)
			return
		}
	}

	hashKey := fmt.Sprintf("%s_LAYER%d", experimentResult.ExperimentContext.Uid, layer.ExpLayerId)
	hashValue := e.hashValue(hashKey)
	hashValueStr := strconv.FormatUint(hashValue, 10)
	context := model.ExperimentContext{
		Uid:          hashValueStr,
		FilterParams: experimentResult.ExperimentContext.FilterParams,
	}

	var matchExperiments []*model.Experiment
	for _, experiment := range layer.Experiments() {
		if experiment.Match(&context) {
			matchExperiments = append(matchExperiments, experiment)
		}
	}
	if len(matchExperiments) > 0 {
		if len(matchExperiments) == 1 {
			e.matchExperiment(matchExperiments[0], experimentResult)
			return
		} else {
			for _, experiment := range matchExperiments {
				if experiment.BucketType == common.ExpBucketTypeCond {
					e.matchExperiment(experiment, experimentResult)
					return
				}
			}
			// if not find the cond bucket, so here have one more rand experiment match
			// should not happen
			e.matchExperiment(matchExperiments[0], experimentResult)
			return

		}

	}

	for _, domain := range layer.Domains() {
		if domain.Match(&context) {
			e.matchDomain(domain, experimentResult)
			return
		}
	}
}

func (e *ExperimentClient) matchExperiment(experiment *model.Experiment, experimentResult *model.ExperimentResult) {
	if experiment == nil {
		return
	}

	// first find the debug user match
	for _, experimentVersion := range experiment.ExperimentVersions() {
		if experimentVersion.MatchDebugUsers(experimentResult.ExperimentContext) {
			e.logInfo(fmt.Sprintf("match debug user for experimentVersion:%s", experimentVersion.ExpVersionName))
			experimentResult.AddExperimentVersion(experimentVersion)
			return
		}
	}
	hashKey := fmt.Sprintf("%s_EXPERIMENT%d", experimentResult.ExperimentContext.Uid, experiment.ExpId)
	hashValue := e.hashValue(hashKey)
	hashValueStr := strconv.FormatUint(hashValue, 10)
	context := model.ExperimentContext{
		Uid:          hashValueStr,
		FilterParams: experimentResult.ExperimentContext.FilterParams,
	}
	for _, experimentVersion := range experiment.ExperimentVersions() {
		if experimentVersion.Match(&context) {
			experimentResult.AddExperimentVersion(experimentVersion)
			return
		}
	}
}

func (e *ExperimentClient) hashValue(hashKey string) uint64 {
	md5 := md5.Sum([]byte(hashKey))
	hash := fnv.New64()
	hash.Write(md5[:])

	return hash.Sum64()
}

func (e *ExperimentClient) logInfo(msg string, args ...interface{}) {
	if e.Logger != nil {
		e.Logger.Printf(msg, args...)
	}
}
