package main

import (
	"github.com/guardian/steampipe-plugin-galaxies/galaxies"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{PluginFunc: galaxies.Plugin})
}
