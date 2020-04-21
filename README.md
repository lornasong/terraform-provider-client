# terraform-provider-client
Client to consume terraform providers using terraform-plugin-sdk

## Get Started
To download dependencies and run with the Palo Alto plugin
```
make deps
make run
```

## Switch Plugins
Download appropriate terraform provider from: https://releases.hashicorp.com
Unzip file and move binary into base of repository
Update the `Makefile` `PROVIDER_PLUGIN` value with binary name
