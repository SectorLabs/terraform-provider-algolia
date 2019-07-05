package main

import (
	"github.com/algolia/algoliasearch-client-go/algoliasearch"
	"github.com/hashicorp/terraform/helper/schema"
)

// Reads an exiting API key
func resourceAPIKeyRead(d *schema.ResourceData, m interface{}) error {
	client := *m.(*algoliasearch.Client)

	result, err := client.GetAPIKey(d.Id())
	if err != nil {
		d.SetId("")
		return err
	}

	d.Set("key", d.Id())
	d.Set("indexes", result.Indexes)
	d.Set("acl", result.ACL)
	d.Set("description", result.Description)
	d.Set("validity", result.Validity)
	d.Set("max_queries_per_ip_per_hour", result.MaxQueriesPerIPPerHour)
	d.Set("max_hits_per_query", result.MaxHitsPerQuery)

	return nil
}
