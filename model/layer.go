package model

import "github.com/aliyun/aliyun-pai-ab-go-sdk/swagger"

type Layer struct {
	*swagger.Layer
	domains     []*Domain
	experiments []*Experiment
}

func NewLayer(layer *swagger.Layer) *Layer {
	return &Layer{
		Layer: layer,
	}
}

// AddDomain assign a domain to layer
func (l *Layer) AddDomain(domain *Domain) {
	l.domains = append(l.domains, domain)
}

// GetDomain get domain by id
// If not found, return nil
/**
func (l *Layer) GetDomain(id int) *Domain {
	return l.domainMap[id]
}
**/

// Domains return all domains
func (l *Layer) Domains() []*Domain {
	return l.domains
}

// AddExperiment add experiment to layer
func (l *Layer) AddExperiment(experiment *Experiment) {
	l.experiments = append(l.experiments, experiment)
}

// Experiments return all experiments
func (l *Layer) Experiments() []*Experiment {
	return l.experiments
}
