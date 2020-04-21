export PROVIDER_PLUGIN="terraform-provider-panos_v1.6.2_x4"

dep:
	go mod download

run:
	go build -o build/provider && ./build/provider
