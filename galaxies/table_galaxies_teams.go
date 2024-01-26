package galaxies

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/guardian/steampipe-plugin-galaxies/store"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

type Team struct {
	ID                 string `json:"teamId"`
	Name               string `json:"teamName"`
	Description        string `json:"teamDescription"`
	ContactEmail       string `json:"teamContactEmail"`
	GoogleChatSpaceKey string `json:"teamGoogleChatSpaceKey"`
	PrimaryGithubTeam  string `json:"teamPrimaryGithubTeam"`
}

func tableTeams() *plugin.Table {
	return &plugin.Table{
		Name:        "galaxies_teams",
		Description: "Teams in the Guardian P&E department",
		List: &plugin.ListConfig{
			Hydrate: getTeams,
		},
		Columns: []*plugin.Column{
			{Name: "id", Type: proto.ColumnType_STRING, Description: "TODO"},
			{Name: "name", Type: proto.ColumnType_STRING, Description: "TODO"},
			{Name: "description", Type: proto.ColumnType_STRING, Description: "TODO"},
			{Name: "contact_email", Type: proto.ColumnType_STRING, Description: "TODO", Transform: transform.FromCamel()},
			{Name: "google_chat_space_key", Type: proto.ColumnType_STRING, Description: "TODO", Transform: transform.FromCamel()},
			{Name: "primary_github_team", Type: proto.ColumnType_STRING, Description: "TODO", Transform: transform.FromCamel()},
		},
	}
}

func getTeams(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	s, err := store.GalaxiesS3(ctx)
	if err != nil {
		return nil, fmt.Errorf("unable to load AWS config, %w", err)
	}

	data, err := s.Get("teams.json")
	if err != nil {
		return err, fmt.Errorf("failed to get file from S3, %w", err)
	}

	var records map[string]Team
	err = json.Unmarshal(data, &records)
	plugin.Logger(ctx).Info("records", records)
	if err != nil {
		return err, fmt.Errorf("failed to parse file, %w", err)
	}

	for ID, record := range records {
		record.ID = ID // Galaxies team data uses the map keys as the ID, so pull this out.
		d.StreamListItem(ctx, record)
	}

	return nil, nil

}
