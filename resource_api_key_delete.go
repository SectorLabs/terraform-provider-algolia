package main

import (
	"github.com/algolia/algoliasearch-client-go/algolia/search"
	"github.com/hashicorp/terraform/helper/schema"
)

// Deletes an existing API key
func resourceAPIKeyDelete(d *schema.ResourceData, m interface{}) error {
	client := *m.(*search.Client)

	result, err := client.DeleteAPIKey(d.Id())
	if err != nil {
		return err
	}

	result.Wait()

	d.SetId("")
	d.Set("key", "")
	d.Set("indexes", "")
	d.Set("acl", "")
	d.Set("description", "")
	d.Set("validity", "")
	d.Set("max_queries_per_ip_per_hour", "")
	d.Set("max_hits_per_query", "")
	return nil
}
