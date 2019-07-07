# terraform-provider-algolia

[![MIT licensed](https://img.shields.io/badge/license-MIT-blue.svg)](https://raw.githubusercontent.com/hyperium/hyper/master/LICENSE)
[![CircleCI](https://circleci.com/gh/SectorLabs/terraform-provider-algolia/tree/master.svg?style=svg&circle-token=b7df36916b22d9a96de05df22ee0ec83f2c102fc)](https://circleci.com/gh/SectorLabs/terraform-provider-algolia/tree/master)

A Terraform provider for Algolia.

Only supports creating, updating and deleting API keys for now. It might be extended with support for other resources in the future.

Pre-built binaries for macOS and Linux can downloaded from the [releases page](https://github.com/SectorLabs/terraform-provider-algolia/releases).

Requires at least Go 1.12

## Installation 

- clone the repository:

```
$ git clone https://github.com/SectorLabs/terraform-provider-algolia.git
```

- build the provider:

```
$ make
```

- move the binary to the Terraform plugins directory:

    Or, use one of the other supported methods of using custom plugins. See:
    [https://www.terraform.io/docs/extend/how-terraform-works.html#discovery](https://www.terraform.io/docs/extend/how-terraform-works.html#discovery)

```
$ mv ./terraform-provider-algolia_v0.2 ~/.terraform.d/plugins
```

## Usage

### Provider configuration

#### `algolia`

```
    provider "algolia" {
        version = "~> 0.2"

        application_id = "algolia app id"
        api_key = "algolia admin api key"
    }
```

##### Argument Reference

The following arguments are supported:

* `application_id` - (required) The algolia app id to connect to
* `api_key` - (required) The Admin API key with permissions to manage keys in the app


### Data Source Configuration


#### `algolia_api_key`

```
    data "algolia_api_key" "test" {
        id = "the_key_value"
    } 
```

##### Argument Reference

The following arguments are supported:

* `id` - (required) The key value itself 
    
### Resources Configuration

#### `algolia_api_key` 

##### Example usage

```
    resource "algolia_api_key" "test" {
        indexes = ["my-index-*"]

        acl = ["search"]

        description = "my api key name"

        validity = 0

        max_queries_per_ip_per_hour = 0

        max_hits_per_query = 0
    }

    output "my_new_api_key" {
        value = algolia_api_key.test.key
    }
```

##### Argument reference

The following arguments are supported:

* `indexes` - (required) List of target indexes. Supports the wildcard syntax.
* `acl` - (required) List of permissions for this key
* `description` - (optional) Description of this key. Informative only.
* `validity`- (optional) A unix timestamp used to define the expiration date of the key. Default: 0 (never expire).
* `max_queries_per_ip_per_hour` - (optional) Maximum amount of API calls allowed from an IP per hour. Default: 0 (no limit)
* `max_hits_per_query` - (optional) Maximum amount of hits that can be retrieved with one query. Default: 0 (no limit)

##### Import 

The keys can be imported using the key value itself:

```
    terraform import algolia_key.test "the_key_value"
```