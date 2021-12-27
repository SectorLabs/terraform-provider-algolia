package main

import (
	"github.com/algolia/algoliasearch-client-go/v3/algolia/search"
	"github.com/hashicorp/terraform/helper/schema"
	"time"
)

// Creates a new API key
func resourceAPIKeyCreate(d *schema.ResourceData, m interface{}) error {
	key := search.Key{
		ACL:                    castStringList(d.Get("acl").([]interface{})),
		Validity:               time.Duration(d.Get("validity").(int)) * time.Second,
		MaxQueriesPerIPPerHour: d.Get("max_queries_per_ip_per_hour").(int),
		MaxHitsPerQuery:        d.Get("max_hits_per_query").(int),
		Indexes:                castStringList(d.Get("indexes").([]interface{})),
		Description:            d.Get("description").(string),
	}

	client := *m.(*search.Client)
	result, err := client.AddAPIKey(key)
	if err != nil {
		return err
	}

	result.Wait()

	d.SetId(result.Key)
	return resourceAPIKeyRead(d, m)
}
