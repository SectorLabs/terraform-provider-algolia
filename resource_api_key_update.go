package main

// Updates an existing API key
func resourceAPIKeyUpdate(d *schema.ResourceData, m interface{}) error {
	return resourceAPIKeyRead(d, m)
}
