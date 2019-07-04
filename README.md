# terraform-provider-algolia

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
    ```
