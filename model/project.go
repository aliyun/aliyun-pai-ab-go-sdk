package model

import "github.com/aliyun/aliyun-pai-ab-go-sdk/swagger"

type Project struct {
	*swagger.Project

	domainMap     map[int]*Domain
	layerMap      map[int]*Layer
	defaultDomain *Domain
}

func NewProject(project *swagger.Project) *Project {
	return &Project{
		Project:   project,
		domainMap: make(map[int]*Domain),
		layerMap:  make(map[int]*Layer),
	}
}

func (p *Project) AddDomain(domain *Domain) {
	p.domainMap[int(domain.ExpDomainId)] = domain
}

func (p *Project) AddLayer(layer *Layer) {
	p.layerMap[int(layer.ExpLayerId)] = layer
}

func (p *Project) GetDomains() map[int]*Domain {
	return p.domainMap
}

func (p *Project) GetLayer(layerId int) *Layer {
	if layer, ok := p.layerMap[layerId]; ok {
		return layer
	}
	return nil
}

func (p *Project) SetDefaultDomain(domain *Domain) {
	p.defaultDomain = domain
}

func (p *Project) DefaultDomain() *Domain {
	return p.defaultDomain
}
