package galaxies

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func Plugin(ctx context.Context) *plugin.Plugin {
	p := &plugin.Plugin{
		Name:             "steampipe-plugin-galaxies",
		DefaultTransform: transform.FromGo().NullIfZero(),
		TableMap: map[string]*plugin.Table{
			"galaxies_people":         tablePeople(),
			"galaxies_people_profile": tablePeopleProfile(),
			"galaxies_teams":          tableTeams(),
			"galaxies_streams":        tableStreams(),
		},
	}
	return p
}
