package main

import (
	"github.com/hashicorp/terraform/helper/schema"
)

// Data source for getting data for an existing API key
func dataSourceAPIKey() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceAPIKeyRead,
		Schema: map[string]*schema.Schema{
			// INPUT REQUIRED
			"id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The key to retrieve data for.",
			},

			// OUTPUT
			"key": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The actual API key.",
			},
			"indexes": &schema.Schema{
				Type: schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Computed:    true,
				Description: "List of target indexes. Supports the wildcard syntax.",
			},
			"acl": &schema.Schema{
				Type: schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Computed:    true,
				Description: "List of permissions associated with this key.",
			},
			"description": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Description of this key. Informative only.",
			},
			"validity": &schema.Schema{
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "A Unix timestamp used to define the expiration date of this key.",
			},
			"max_queries_per_ip_per_hour": &schema.Schema{
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Specify the maximum number of API calls allowed from an IP address per hour. 0 = unlimited.",
			},
			"max_hits_per_query": &schema.Schema{
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Specify the maximum number of hits this key can retrieve in one call. 0 = unlimited.",
			},
		},
	}
}
