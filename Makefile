provider_name="terraform-provider-algolia"
provider_version="0.4"

provider_file_name="$(provider_name)_v$(provider_version)"

build:
	CGO_ENABLED=0 go build -o $(provider_file_name)

format:
	go fmt

run: build
	./$(provider_file_name)

debug_plan: build
	TF_LOG=DEBUG terraform init && terraform plan

debug_apply: build
	TF_LOG=DEBUG terraform init && terraform apply

all: build
