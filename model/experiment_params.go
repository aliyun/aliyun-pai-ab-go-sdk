package model

import "strconv"

// ExperimentParams offers Get* function to get value by the key
// If not found the key, defaultValue will return
type ExperimentParams interface {
	AddParam(key string, value interface{})

	AddParams(params map[string]interface{})

	Get(key string, defaultValue interface{}) interface{}

	GetString(key, defaultValue string) string

	GetInt(key string, defaultValue int) int

	GetFloat(key string, defaultValue float64) float64

	GetInt64(key string, defaultValue int64) int64
}

type experimentParams struct {
	Parameters map[string]interface{}
}

func NewExperimentParams() *experimentParams {
	return &experimentParams{
		Parameters: make(map[string]interface{}, 0),
	}
}

func (r *experimentParams) AddParam(key string, value interface{}) {
	r.Parameters[key] = value
}

func (r *experimentParams) AddParams(params map[string]interface{}) {
	for k, v := range params {
		r.Parameters[k] = v
	}
}

func (r *experimentParams) Get(key string, defaultValue interface{}) interface{} {
	if val, ok := r.Parameters[key]; ok {
		return val
	}
	return defaultValue
}

func (r *experimentParams) GetString(key, defaultValue string) string {
	val, ok := r.Parameters[key]
	if !ok {
		return defaultValue
	}

	switch value := val.(type) {
	case string:
		return value
	case int:
		return strconv.Itoa(value)
	case float64:
		return strconv.Itoa(int(value))
	case int32:
		return strconv.Itoa(int(value))
	case int64:
		return strconv.Itoa(int(value))
	}
	return defaultValue
}

func (r *experimentParams) GetInt(key string, defaultValue int) int {
	val, ok := r.Parameters[key]
	if !ok {
		return defaultValue
	}
	switch value := val.(type) {
	case int:
		return value
	case float64:
		return int(value)
	case uint:
		return int(value)
	case int32:
		return int(value)
	case int64:
		return int(value)
	case string:
		if val, err := strconv.Atoi(value); err == nil {
			return val
		} else {
			return defaultValue
		}
	default:
		return defaultValue
	}
}
func (r *experimentParams) GetFloat(key string, defaultValue float64) float64 {
	val, ok := r.Parameters[key]
	if !ok {
		return defaultValue
	}

	switch value := val.(type) {
	case float64:
		return value
	case int:
		return float64(value)
	case string:
		if f, err := strconv.ParseFloat(value, 64); err == nil {
			return f
		} else {
			return defaultValue
		}
	default:
		return defaultValue
	}
}
func (r *experimentParams) GetInt64(key string, defaultValue int64) int64 {
	val, ok := r.Parameters[key]
	if !ok {
		return defaultValue
	}

	switch value := val.(type) {
	case int:
		return int64(value)
	case float64:
		return int64(value)
	case uint:
		return int64(value)
	case int32:
		return int64(value)
	case int64:
		return value
	case string:
		if val, err := strconv.ParseInt(value, 10, 64); err == nil {
			return val
		} else {
			return defaultValue
		}
	default:
		return defaultValue
	}
}

func MergeExperimentParams(layersParamsMap map[string]ExperimentParams) ExperimentParams {
	mergedParams := NewExperimentParams()
	for _, unmergedParams := range layersParamsMap {
		switch v := unmergedParams.(type) {
		case *experimentParams:
			for k, p := range v.Parameters {
				mergedParams.Parameters[k] = p
			}
		}
	}
	return ExperimentParams(mergedParams)
}
