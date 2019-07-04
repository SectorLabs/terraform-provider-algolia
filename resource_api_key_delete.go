package main

import (
	"github.com/algolia/algoliasearch-client-go/algoliasearch"
	"github.com/hashicorp/terraform/helper/schema"
)

// Deletes an existing API key
func resourceAPIKeyDelete(d *schema.ResourceData, m interface{}) error {
	client := *m.(*algoliasearch.Client)

	_, err := client.DeleteAPIKey(d.Id())
	if err != nil {
		return err
	}

	d.SetId("")
	return nil
}
