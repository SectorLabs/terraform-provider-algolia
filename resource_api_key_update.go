package main

import (
	"github.com/algolia/algoliasearch-client-go/v3/algolia/search"
	"github.com/hashicorp/terraform/helper/schema"
	"time"
)

// Updates an existing API key
func resourceAPIKeyUpdate(d *schema.ResourceData, m interface{}) error {
	key := search.Key{
		Value:                  d.Id(),
		ACL:                    castStringList(d.Get("acl").([]interface{})),
		Validity:               time.Duration(d.Get("validity").(int)) * time.Second,
		MaxQueriesPerIPPerHour: d.Get("max_queries_per_ip_per_hour").(int),
		MaxHitsPerQuery:        d.Get("max_hits_per_query").(int),
		Indexes:                castStringList(d.Get("indexes").([]interface{})),
		Description:            d.Get("description").(string),
	}

	client := *m.(*search.Client)
	result, err := client.UpdateAPIKey(key)
	if err != nil {
		return err
	}

	result.Wait()

	d.SetId(result.Key)
	return resourceAPIKeyRead(d, m)
}
