package main

import (
	"github.com/algolia/algoliasearch-client-go/algolia/search"
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
		DataSourcesMap: map[string]*schema.Resource{
			"algolia_api_key": dataSourceAPIKey(),
		},
		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(data *schema.ResourceData) (interface{}, error) {
	applicationID := data.Get("application_id").(string)
	apiKey := data.Get("api_key").(string)

	client := search.NewClient(applicationID, apiKey)
	return client, nil
}
