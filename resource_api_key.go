package main

import (
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
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"indexes": &schema.Schema{
				Type: schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Required: true,
			},
			"acl": &schema.Schema{
				Type: schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Required: true,
			},

			// Optional
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Required: false,
				Optional: true,
				Default:  "",
			},
			"validity": &schema.Schema{
				Type:     schema.TypeInt,
				Required: false,
				Optional: true,
				Default:  0,
			},
			"max_queries_per_ip_per_hour": &schema.Schema{
				Type:     schema.TypeInt,
				Required: false,
				Optional: true,
				Default:  0,
			},
			"max_hits_per_query": &schema.Schema{
				Type:     schema.TypeInt,
				Required: false,
				Optional: true,
				Default:  0,
			},
		},
	}
}

func resourceAPIKeyCreate(d *schema.ResourceData, m interface{}) error {
	return resourceAPIKeyRead(d, m)
}

func resourceAPIKeyRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceAPIKeyUpdate(d *schema.ResourceData, m interface{}) error {
	return resourceAPIKeyRead(d, m)
}

func resourceAPIKeyDelete(d *schema.ResourceData, m interface{}) error {
	return nil
}
