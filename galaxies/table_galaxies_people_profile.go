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

type PersonProfileInfo struct {
	ID            string `json:"id"`
	Email         string `json:"email"`
	Pronouns      string `json:"pronouns"`
	GithubHandle  string `json:"gitHubHandle"`
	TwitterHandle string `json:"twitterHandle"`
	Website       any    `json:"website"`
	PictureUrl    string `json:"pictureUrl"`
}

func tablePeopleProfile() *plugin.Table {
	return &plugin.Table{
		Name:        "galaxies_people_profile",
		Description: "Profiles for people in the Guardian P&E department",
		List: &plugin.ListConfig{
			Hydrate: getPeopleProfile,
		},
		Columns: []*plugin.Column{
			{Name: "id", Type: proto.ColumnType_STRING, Description: "TODO"},
			{Name: "email", Type: proto.ColumnType_STRING, Description: "TODO"},
			{Name: "pronouns", Type: proto.ColumnType_STRING, Description: "TODO"},
			{Name: "github_handle", Type: proto.ColumnType_STRING, Description: "TODO", Transform: transform.FromCamel()},
			{Name: "twitter_handle", Type: proto.ColumnType_STRING, Description: "TODO", Transform: transform.FromCamel()},
			{Name: "website", Type: proto.ColumnType_STRING, Description: "TODO"},
			{Name: "picture_url", Type: proto.ColumnType_STRING, Description: "TODO", Transform: transform.FromCamel()},
		},
	}
}

func getPeopleProfile(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	s, err := store.GalaxiesS3(ctx)
	if err != nil {
		return nil, fmt.Errorf("unable to load AWS config, %w", err)
	}

	data, err := s.Get("peopleProfileInfo.json")
	if err != nil {
		return err, fmt.Errorf("failed to get file from S3, %w", err)
	}

	var records map[string]PersonProfileInfo
	err = json.Unmarshal(data, &records)
	plugin.Logger(ctx).Info("records", records)
	if err != nil {
		return err, fmt.Errorf("failed to parse file, %w", err)
	}

	for ID, record := range records {
		record.ID = ID // Galaxies profile info uses email ID as the map keys, so pull this out.
		d.StreamListItem(ctx, record)
	}

	return nil, nil

}
