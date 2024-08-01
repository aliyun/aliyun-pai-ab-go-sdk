package model

import (
	"encoding/json"

	"github.com/aliyun/aliyun-pai-ab-go-sdk/swagger"
)

type Feature struct {
	*swagger.Feature

	params map[string]any

	diversionBucket DiversionBucket
}

func NewFeature(feature *swagger.Feature) *Feature {
	return &Feature{
		Feature: feature,
		params:  make(map[string]any),
	}
}

// Init
func (f *Feature) Init() error {
	expParams := make([]ExpParameter, 0)

	if err := json.Unmarshal([]byte(f.Config), &expParams); err != nil {
		return err
	}

	for _, p := range expParams {
		f.params[p.Key] = p.Value
	}

	if f.Filter != "" {
		diversionBucket, err := NewFilterDiversionBucket(f.Filter)
		if err != nil {
			return err
		}
		f.diversionBucket = diversionBucket
	}
	return nil
}

// Params
func (f *Feature) Params() map[string]any {
	return f.params
}

// Match
func (f *Feature) Match(experimentContext *ExperimentContext) bool {
	if f.diversionBucket == nil {
		return true
	}
	return f.diversionBucket.Match(experimentContext)
}
