package experiments

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/aliyun/aliyun-pai-ab-go-sdk/common"
	"github.com/aliyun/aliyun-pai-ab-go-sdk/model"
)

func (e *ExperimentClient) logError(err error) {
	if e.ErrorLogger != nil {
		e.ErrorLogger.Printf(err.Error())
		return
	}

	if e.Logger != nil {
		e.Logger.Printf(err.Error())
	}
}

// LoadExperimentData specifies a function to load data from A/B Test Server
func (e *ExperimentClient) LoadExperimentData() {
	projectData := make(map[string]*model.Project, 0)

	projects, err := e.APIClient.ProjectApi.ListProjects()
	if err != nil {
		e.logError(fmt.Errorf("list projects error, err=%v", err))
		return
	}

	for _, p := range projects {
		project := model.NewProject(p)
		domains, err := e.APIClient.DomainApi.ListDomains(int(project.ExpProjectId))

		if err != nil {
			e.logError(fmt.Errorf("list domains error, err=%v", err))
			continue
		}
		for _, d := range domains {
			domain := model.NewDomain(d)
			if domain.DebugCrowdIds != "" {
				crowdIds := strings.Split(domain.DebugCrowdIds, ",")
				for _, crowdId := range crowdIds {
					if id, err := strconv.Atoi(crowdId); err == nil {
						users, err := e.APIClient.CrowdApi.ListCrowdUsers(id)
						if err != nil {
							e.logError(fmt.Errorf("get crowd users error, err=%v", err))
							continue
						}
						domain.AddDebugUsers(users)
					}
				}
			}
			// domain init
			if err := domain.Init(); err != nil {
				e.logError(fmt.Errorf("domain init error, err=%v", err))
				continue
			}

			if domain.IsDefaultDomain {
				project.SetDefaultDomain(domain)
			}

			project.AddDomain(domain)

			features, err := e.APIClient.FeatureApi.ListFeatures(int(domain.ExpDomainId))
			if err != nil {
				e.logError(fmt.Errorf("list features error, err=%v", err))
				continue
			}

			for _, f := range features {
				feature := model.NewFeature(f)
				if err := feature.Init(); err != nil {
					e.logError(fmt.Errorf("feature init error, err=%v", err))
					continue
				}
				domain.AddFeature(feature)
			}

			layers, err := e.APIClient.LayerApi.ListLayers(int(project.ExpProjectId))
			if err != nil {
				e.logError(fmt.Errorf("list layers error, err=%v", err))
				continue
			}
			for _, l := range layers {
				layer := model.NewLayer(l)
				domain.AddLayer(layer)
				project.AddLayer(layer)

				experiments, err := e.APIClient.ExperimentApi.ListExperiments(int(layer.ExpLayerId), common.ExpStatusRunning)
				if err != nil {
					e.logError(fmt.Errorf("list experiment  error, err=%v", err))
					continue
				}
				for _, exp := range experiments {
					experiment := model.NewExperiment(exp)
					if experiment.DebugCrowdIds != "" {
						crowdIds := strings.Split(experiment.DebugCrowdIds, ",")
						for _, crowdId := range crowdIds {

							if id, err := strconv.Atoi(crowdId); err == nil {
								users, err := e.APIClient.CrowdApi.ListCrowdUsers(id)
								if err != nil {
									e.logError(fmt.Errorf("get crowd users error, err=%v", err))
									continue
								}
								experiment.AddDebugUsers(users)
							}
						}

					}
					if err := experiment.Init(); err != nil {
						e.logError(fmt.Errorf("experiment group init error, err=%v", err))
						continue
					}
					layer.AddExperiment(experiment)
					// add experiment version
					experimentVersions, err := e.APIClient.ExperimentVersionApi.ListExperimentVersions(int(experiment.ExpId))
					if err != nil {
						e.logError(fmt.Errorf("list experiment version  error, err=%v", err))
						continue
					}

					for _, version := range experimentVersions {
						experimentVersion := model.NewExperimentVersion(version)
						if experimentVersion.DebugCrowdIds != "" {
							crowdIds := strings.Split(experimentVersion.DebugCrowdIds, ",")
							for _, crowdId := range crowdIds {

								if id, err := strconv.Atoi(crowdId); err == nil {
									users, err := e.APIClient.CrowdApi.ListCrowdUsers(id)
									if err != nil {
										e.logError(fmt.Errorf("get crowd users error, err=%v", err))
										continue
									}
									experimentVersion.AddDebugUsers(users)

								}
							}
						}
						if err := experimentVersion.Init(); err != nil {
							e.logError(fmt.Errorf("experiment version init error, err=%v", err))
							continue
						}
						experimentVersion.SetExperiment(experiment)
						experiment.AddExperimentVersion(experimentVersion)
					}

				}
			}
		}

		// assigin domain to  layer
		for _, domain := range project.GetDomains() {
			if !domain.IsDefaultDomain {
				layer := project.GetLayer(int(domain.ExpLayerId))
				if layer != nil {
					layer.AddDomain(domain)
				}
			}
		}
		projectData[project.ProjectName] = project
	}
	if len(projectData) > 0 {
		e.projectMap = projectData
	}
}

// loopLoadExperimentData async loop invoke LoadExperimentData function
func (e *ExperimentClient) loopLoadExperimentData() {

	for {
		time.Sleep(time.Minute)
		e.LoadExperimentData()
	}
}
