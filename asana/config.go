package asana

import (
	"fmt"
	"github.com/davidji99/terraform-provider-asana/api"
	"github.com/davidji99/terraform-provider-asana/version"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
)

var (
	UserAgent = fmt.Sprintf("terraform-provider-asana/v%s", version.ProviderVersion)
)

type Config struct {
	API        *api.Client
	Headers    map[string]string
	acessToken string
	baseURL    string
}

func NewConfig() *Config {
	return &Config{}
}

func (c *Config) initializeAPI() error {
	// Initialize the custom API client for non Heroku Platform APIs
	api, clientInitErr := api.New(
		api.AccessToken(c.acessToken), api.CustomHTTPHeaders(c.Headers),
		api.UserAgent(UserAgent), api.BaseURL(c.baseURL))
	if clientInitErr != nil {
		return clientInitErr
	}
	c.API = api

	log.Printf("[INFO] Asana Client configured")

	return nil
}

func (c *Config) applySchema(d *schema.ResourceData) (err error) {
	if v, ok := d.GetOk("headers"); ok {
		headersRaw := v.(map[string]interface{})
		h := make(map[string]string)

		for k, v := range headersRaw {
			h[k] = fmt.Sprintf("%v", v)
		}

		c.Headers = h
	}

	if v, ok := d.GetOk("url"); ok {
		c.baseURL = v.(string)
	}

	return nil
}
