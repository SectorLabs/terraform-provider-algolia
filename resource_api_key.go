package main

import (
	"github.com/algolia/algoliasearch-client-go/algoliasearch"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceAPIKey() *schema.Resource {
	return &schema.Resource{
		Create: resourceAPIKeyCreate,
		Read:   resourceAPIKeyRead,
		Update: resourceAPIKeyUpdate,
		Delete: resourceAPIKeyDelete,

		// Naming is according to the names used in the offical Golang
		// client for Algolia
		// See: https://www.algolia.com/doc/api-reference/api-methods/add-api-key/
		Schema: map[string]*schema.Schema{
			// Required
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

			// Optional
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
		},
	}
}

// Creates a new API key
func resourceAPIKeyCreate(d *schema.ResourceData, m interface{}) error {
	acl := castStringList(d.Get("acl").([]interface{}))

	params := algoliasearch.Map{
		"validity":               d.Get("validity").(int),
		"maxQueriesPerIPPerHour": d.Get("max_queries_per_ip_per_hour").(int),
		"maxHitsPerQuery":        d.Get("max_hits_per_query").(int),
		"indexes":                castStringList(d.Get("indexes").([]interface{})),
		"description":            d.Get("description").(string),
	}

	client := *m.(*algoliasearch.Client)
	result, err := client.AddAPIKey(acl, params)
	if err != nil {
		return err
	}

	d.SetId(result.Key)
	return nil
}

// Reads an exiting API key
func resourceAPIKeyRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

// Updates an existing API key
func resourceAPIKeyUpdate(d *schema.ResourceData, m interface{}) error {
	return resourceAPIKeyRead(d, m)
}

// Deletes an existing API key
func resourceAPIKeyDelete(d *schema.ResourceData, m interface{}) error {
	return nil
}

func castStringList(configured []interface{}) []string {
	vs := make([]string, 0, len(configured))
	for _, v := range configured {
		val, ok := v.(string)
		if ok && val != "" {
			vs = append(vs, v.(string))
		}
	}
	return vs
}
