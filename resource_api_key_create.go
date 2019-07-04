package main

import (
	"github.com/algolia/algoliasearch-client-go/algoliasearch"
	"github.com/hashicorp/terraform/helper/schema"
)

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
