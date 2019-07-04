# terraform-provider-algolia

[![MIT licensed](https://img.shields.io/badge/license-MIT-blue.svg)](https://raw.githubusercontent.com/hyperium/hyper/master/LICENSE)

A Terraform provider for Algolia.

## Usage instructions
1. Clone the repository:

    ```
    $ git clone https://github.com/SectorLabs/terraform-provider-algolia.git
    ```

2. Build the provider:

    ```
    $ make
    ```

3. Move the binary to the Terraform plugins directory:

    Or, use one of the other supported methods of using custom plugins. See:
    [https://www.terraform.io/docs/extend/how-terraform-works.html#discovery](https://www.terraform.io/docs/extend/how-terraform-works.html#discovery)

    ```
    $ mv ./terraform-provider-algolia_v0.1 ~/.terraform.d/plugins
    ```

4. Use the provider in Terraform:

    ```
    provider "algolia" {
        version = "~> 0.1"

        application_id = "algolia app id"
        api_key = "algolia admin api key"
    }

    resource "algolia_api_key" "test" {
        // Required
        // List of target indexes. Supports the wildcard syntax.
        indexes = ["my-index-*"]

        // Required
        // List of permissions for this key
        acl = ["search"]

        // Optional
        // Description of this key. Informative only.
        description = "my api key name"

        // Optional
        // A unix timestamp used to define the expiration date
        // of the key. Default: 0 (never expire).
        validity = 0

        // Optional
        // Maximum amount of API calls allowed from an IP per hour.
        // Default: 0 (no limit)
        max_queries_per_ip_per_hour = 0

        // Optional
        // Maximum amount of hits that can be retrieved with one query.
        // Default: 0 (no limit)
        max_hits_per_query = 0
    }

    output "my_new_api_key" {
        value = algolia_api_key.test.key
    }
    ```
