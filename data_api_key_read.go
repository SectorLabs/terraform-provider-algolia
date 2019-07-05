package main

import (
	"github.com/hashicorp/terraform/helper/schema"
)

// Reads all details for an API key
func dataSourceAPIKeyRead(d *schema.ResourceData, m interface{}) error {
	id := d.Get("id").(string)
	d.SetId(id)

	return resourceAPIKeyRead(d, m)
}
