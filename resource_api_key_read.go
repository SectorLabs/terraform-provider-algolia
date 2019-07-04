package main

import (
	"github.com/algolia/algoliasearch-client-go/algoliasearch"
	"github.com/hashicorp/terraform/helper/schema"
)

// Reads an exiting API key
func resourceAPIKeyRead(d *schema.ResourceData, m interface{}) error {
	client := *m.(*algoliasearch.Client)

	_, err := client.GetAPIKey(d.Id())
	if err != nil {
		d.SetId("")
		return err
	}

	return nil
}
