package main

import (
	"github.com/hashicorp/terraform/helper/customdiff"
	"github.com/hashicorp/terraform/helper/schema"
)

// Defines the "algolia.api_key" resource
func resourceAPIKey() *schema.Resource {
	return &schema.Resource{
		Create: resourceAPIKeyCreate,
		Read:   resourceAPIKeyRead,
		Update: resourceAPIKeyUpdate,
		Delete: resourceAPIKeyDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		CustomizeDiff: customdiff.All(),

		// Naming is according to the names used in the offical Golang
		// client for Algolia
		// See: https://www.algolia.com/doc/api-reference/api-methods/add-api-key/
		Schema: map[string]*schema.Schema{
			// Required Input
			"indexes": &schema.Schema{
				Type: schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Required:    true,
				Description: "List of target indexes. Supports the wildcard syntax.",
			},
			"acl": &schema.Schema{
				Type: schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Required:    true,
				Description: "List of permissions associated with this key.",
			},

			// Optional Input
			"description": &schema.Schema{
				Type:        schema.TypeString,
				Required:    false,
				Optional:    true,
				Default:     "",
				Description: "Description of this key. Informative only.",
			},
			"validity": &schema.Schema{
				Type:        schema.TypeInt,
				Required:    false,
				Optional:    true,
				Default:     0,
				Description: "A Unix timestamp used to define the expiration date of this key.",
			},
			"max_queries_per_ip_per_hour": &schema.Schema{
				Type:        schema.TypeInt,
				Required:    false,
				Optional:    true,
				Default:     0,
				Description: "Specify the maximum number of API calls allowed from an IP address per hour. 0 = unlimited.",
			},
			"max_hits_per_query": &schema.Schema{
				Type:        schema.TypeInt,
				Required:    false,
				Optional:    true,
				Default:     0,
				Description: "Specify the maximum number of hits this key can retrieve in one call. 0 = unlimited.",
			},

			// Output
			"key": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "API key that was created.",
			},
		},
	}
}
