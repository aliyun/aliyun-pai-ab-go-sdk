package model

import (
	"strings"

	"github.com/aliyun/aliyun-pai-ab-go-sdk/common"
	"github.com/aliyun/aliyun-pai-ab-go-sdk/swagger"
)

type Domain struct {
	*swagger.Domain

	debugUserMap map[string]bool `json:"-"`

	diversionBucket DiversionBucket

	layers []*Layer

	features []*Feature
}

func NewDomain(domain *swagger.Domain) *Domain {
	return &Domain{
		Domain:       domain,
		debugUserMap: make(map[string]bool),
	}
}

func (d *Domain) Init() error {

	if d.diversionBucket == nil {
		if d.BucketType == common.DomainBucketTypeCond {
			diversionBucket, err := NewFilterDiversionBucket(d.Filter)
			if err != nil {
				return err
			}
			d.diversionBucket = diversionBucket
		} else if d.BucketType == common.DomainBucketTypeRand {
			diversionBucket := NewUidDiversionBucket(100, d.Buckets)
			d.diversionBucket = diversionBucket
		}
	}

	if d.DebugUsers != "" {
		uids := strings.Split(d.DebugUsers, ",")
		for _, uid := range uids {
			d.debugUserMap[uid] = true
		}
	}

	return nil
}

func (d *Domain) AddLayer(layer *Layer) {
	d.layers = append(d.layers, layer)
}

// Layers
func (d *Domain) Layers() []*Layer {
	return d.layers
}

// AddFeature
func (d *Domain) AddFeature(feature *Feature) {
	d.features = append(d.features, feature)
}

func (d *Domain) Features() []*Feature {
	return d.features
}

// AddDebugUsers
func (d *Domain) AddDebugUsers(users []string) {
	for _, uid := range users {
		d.debugUserMap[uid] = true
	}
}

// MatchDebugUsers return true if debug_users is set and debug_users contain of uid
func (d *Domain) MatchDebugUsers(experimentContext *ExperimentContext) bool {
	if _, found := d.debugUserMap[experimentContext.Uid]; found {
		return true
	}

	return false
}
func (d *Domain) Match(experimentContext *ExperimentContext) bool {

	if d.ExperimentFlow == 0 {
		return false
	}
	if d.ExperimentFlow == 100 {
		return true
	}

	if d.diversionBucket != nil {
		return d.diversionBucket.Match(experimentContext)
	}

	return false
}
