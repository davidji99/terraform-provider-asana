package asana

import (
	"context"
	"github.com/davidji99/terraform-provider-asana/api"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"access_token": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("ASANA_ACCESS_TOKEN", nil),
			},

			"url": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("ASANA_URL", api.DefaultAPIBaseURL),
			},

			"headers": {
				Type:     schema.TypeMap,
				Elem:     schema.TypeString,
				Optional: true,
			},
		},

		DataSourcesMap: map[string]*schema.Resource{},

		ResourcesMap: map[string]*schema.Resource{
			"asana_project": resourceAsanaProject(),
		},

		ConfigureContextFunc: providerConfigure,
	}
}

func providerConfigure(_ context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	log.Println("[INFO] Initializing Asana Provider")

	var diags diag.Diagnostics

	config := NewConfig()

	if applySchemaErr := config.applySchema(d); applySchemaErr != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unable to retrieve and set provider attributes",
			Detail:   applySchemaErr.Error(),
		})

		return nil, diags
	}

	if token, ok := d.GetOk("access_token"); ok {
		config.acessToken = token.(string)
	}

	if err := config.initializeAPI(); err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unable to initialize API client",
			Detail:   err.Error(),
		})

		return nil, diags
	}

	log.Printf("[DEBUG] Asana Provider initialized")

	return config, diags
}
