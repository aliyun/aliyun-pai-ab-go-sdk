package model

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
)

type ExperimentResult struct {
	ExperimentContext *ExperimentContext

	projectName string

	project *Project

	expId string

	experimentVersions []*ExperimentVersion

	experimentParams ExperimentParams
}

func NewExperimentResult(projectName string, experimentContext *ExperimentContext, project *Project) *ExperimentResult {
	result := ExperimentResult{
		projectName:        projectName,
		project:            project,
		ExperimentContext:  experimentContext,
		experimentVersions: make([]*ExperimentVersion, 0),
		experimentParams:   NewExperimentParams(),
	}

	return &result
}

// AddExperimentVersion
func (r *ExperimentResult) AddExperimentVersion(experimentVersion *ExperimentVersion) {
	r.experimentVersions = append(r.experimentVersions, experimentVersion)
}

func (r *ExperimentResult) GetExpId() string {
	return r.expId
}

func (r *ExperimentResult) Init() {
	buf := bytes.NewBuffer(nil)
	buf.WriteString("ER")
	buf.WriteString(strconv.Itoa(int(r.project.ExpProjectId)))
	if len(r.experimentVersions) > 0 {
		for _, experimentVersion := range r.experimentVersions {
			buf.WriteString("_E")
			buf.WriteString(strconv.Itoa(int(experimentVersion.Experiment().ExpId)))
			buf.WriteString("#EV")
			buf.WriteString(strconv.Itoa(int(experimentVersion.ExpVersionId)))
		}

		r.expId = buf.String()

		for _, experimentVersion := range r.experimentVersions {
			r.experimentParams.AddParams(experimentVersion.Params())
		}

	}
}

func (r *ExperimentResult) GetExperimentParams() ExperimentParams {
	return r.experimentParams
}

func (r *ExperimentResult) Info() string {
	var info []string

	if r.ExperimentContext != nil {
		info = append(info, fmt.Sprintf("requestId=%s", r.ExperimentContext.RequestId))
		info = append(info, fmt.Sprintf("uid=%s", r.ExperimentContext.Uid))
	}
	info = append(info, fmt.Sprintf("project_name=%s", r.projectName))
	info = append(info, fmt.Sprintf("exp_id=%s", r.expId))

	return strings.Join(info, "\t")
}
