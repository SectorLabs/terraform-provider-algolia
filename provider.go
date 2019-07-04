package main

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"application_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Algolia Application ID",
			},
			"api_key": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Algolia Admin API Key",
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"algolia_api_key": resourceAPIKey(),
		},
	}
}
