package main

import (
	"github.com/hashicorp/terraform/plugin"
	"github.com/terraform-providers/terraform-provider-ucloud/ucloud"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: ucloud.Provider,
	})
}
