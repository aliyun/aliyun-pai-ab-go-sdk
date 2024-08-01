package model

import (
	"strings"

	"github.com/aliyun/aliyun-pai-ab-go-sdk/swagger"
)

type Experiment struct {
	*swagger.Experiment

	debugUserMap map[string]bool `json:"-"`

	diversionBucket DiversionBucket

	experimentVersions []*ExperimentVersion
}

func NewExperiment(experiment *swagger.Experiment) *Experiment {
	return &Experiment{
		Experiment:   experiment,
		debugUserMap: make(map[string]bool),
	}
}

func (e *Experiment) Init() error {
	if e.DebugUsers != "" {
		uids := strings.Split(e.DebugUsers, ",")
		for _, uid := range uids {
			e.debugUserMap[uid] = true
		}
	}

	if e.Filter != "" {
		diversionBucket, err := NewFilterDiversionBucket(e.Filter)
		if err != nil {
			return err
		}

		e.diversionBucket = diversionBucket
	} else {
		e.diversionBucket = NewUidDiversionBucket(100, e.Buckets)
	}
	return nil
}

// AddExperimentVersion
func (e *Experiment) AddExperimentVersion(experimentVersion *ExperimentVersion) {
	e.experimentVersions = append(e.experimentVersions, experimentVersion)
}

// ExperimentVersions return all experiment versions
func (e *Experiment) ExperimentVersions() []*ExperimentVersion {
	return e.experimentVersions
}

// AddDebugUsers
func (e *Experiment) AddDebugUsers(users []string) {
	for _, uid := range users {
		e.debugUserMap[uid] = true
	}
}

// MatchDebugUsers return true if debug_users is set and debug_users contain of uid
func (e *Experiment) MatchDebugUsers(experimentContext *ExperimentContext) bool {
	if _, found := e.debugUserMap[experimentContext.Uid]; found {
		return true
	}

	return false
}

func (e *Experiment) Match(experimentContext *ExperimentContext) bool {

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
