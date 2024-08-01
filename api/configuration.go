package api

import (
	"fmt"

	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	"github.com/alibabacloud-go/tea/tea"
)

type Configuration struct {
	useVpc bool
	domain string
	region string
	config *openapi.Config
}

func NewConfiguration(region, accessId, accessKey string) *Configuration {

	config := openapi.Config{
		AccessKeyId:     tea.String(accessId),
		AccessKeySecret: tea.String(accessKey),
		RegionId:        tea.String(region),
		MaxIdleConns:    tea.Int(1000),
	}
	c := &Configuration{
		useVpc: false,
		config: &config,
		region: region,
	}

	return c
}
func (c *Configuration) UseVpc(b bool) {
	c.useVpc = b
}

func (c *Configuration) GetDomain() string {
	if c.domain == "" {
		if c.useVpc {
			c.domain = fmt.Sprintf("paiabtest-vpc.%s.aliyuncs.com", c.region)
		} else {
			c.domain = fmt.Sprintf("paiabtest.%s.aliyuncs.com", c.region)
		}
	}

	return c.domain
}

func (c *Configuration) GetConfig() *openapi.Config {
	c.config.SetEndpoint(c.GetDomain())
	return c.config
}
