package model

import (
	"encoding/json"
	"strings"

	"github.com/aliyun/aliyun-pai-ab-go-sdk/swagger"
)

type ExpParameter struct {
	Key   string `json:"key"`
	Type  string `json:"type"`
	Value any    `json:"value"`
}
type ExperimentVersion struct {
	*swagger.ExperimentVersion
	debugUserMap map[string]bool `json:"-"`

	diversionBucket DiversionBucket

	params map[string]any

	experiment *Experiment
}

func NewExperimentVersion(experimentVersion *swagger.ExperimentVersion) *ExperimentVersion {
	return &ExperimentVersion{
		ExperimentVersion: experimentVersion,
		params:            make(map[string]any),
		debugUserMap:      make(map[string]bool),
	}
}

// Init
func (e *ExperimentVersion) Init() error {
	if e.DebugUsers != "" {
		uids := strings.Split(e.DebugUsers, ",")
		for _, uid := range uids {
			e.debugUserMap[uid] = true
		}
	}
	e.diversionBucket = NewUidDiversionBucket(100, e.Buckets)

	expParams := make([]ExpParameter, 0)

	if err := json.Unmarshal([]byte(e.ExpVersionConfig), &expParams); err != nil {
		return err
	}

	for _, p := range expParams {
		e.params[p.Key] = p.Value
	}
	return nil
}

// SetExperiment
func (e *ExperimentVersion) SetExperiment(experiment *Experiment) {
	e.experiment = experiment
}

// Experiment
func (e *ExperimentVersion) Experiment() *Experiment {
	return e.experiment
}

// Params
func (e *ExperimentVersion) Params() map[string]any {
	return e.params
}

// AddDebugUsers
func (e *ExperimentVersion) AddDebugUsers(users []string) {
	for _, uid := range users {
		e.debugUserMap[uid] = true
	}
}

// MatchDebugUsers return true if debug_users is set and debug_users contain of uid
func (e *ExperimentVersion) MatchDebugUsers(experimentContext *ExperimentContext) bool {
	if _, found := e.debugUserMap[experimentContext.Uid]; found {
		return true
	}

	return false
}

func (e *ExperimentVersion) Match(experimentContext *ExperimentContext) bool {

	if e.ExperimentFlow == 0 {
		return false
	}
	if e.ExperimentFlow == 100 {
		return true
	}

	if e.diversionBucket != nil {
		return e.diversionBucket.Match(experimentContext)
	}

	return false
}
