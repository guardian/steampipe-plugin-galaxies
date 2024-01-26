package galaxies

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/guardian/steampipe-plugin-galaxies/store"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
	"os"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

type Person struct {
	Name    string   `json:"name"`
	EmailId string   `json:"emailId"`
	Role    string   `json:"role"`
	Teams   []string `json:"teams"`
	Streams []string `json:"streams"`
}

func tablePeople() *plugin.Table {
	return &plugin.Table{
		Name:        "galaxies_people",
		Description: "People in the Guardian P&E department",
		List: &plugin.ListConfig{
			Hydrate: getPeople,
		},
		Columns: []*plugin.Column{
			{Name: "name", Type: proto.ColumnType_STRING, Description: "TODO"},
			{Name: "email_id", Type: proto.ColumnType_STRING, Description: "TODO", Transform: transform.FromCamel()},
			{Name: "role", Type: proto.ColumnType_STRING, Description: "TODO"},
			{Name: "teams", Type: proto.ColumnType_JSON, Description: "TODO"},
			{Name: "streams", Type: proto.ColumnType_JSON, Description: "TODO"},
		},
	}
}

func getPeople(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	cfg, err := config.LoadDefaultConfig(
		ctx,
		config.WithRegion("eu-west-1"),
		config.WithSharedConfigProfile("deployTools"),
	)

	if err != nil {
		return nil, fmt.Errorf("unable to load AWS config, %w", err)
	}

	client := s3.NewFromConfig(cfg)
	bucket := os.Getenv("GALAXIES_BUCKET")
	s := store.New(client, bucket, "galaxies.gutools.co.uk/data")

	data, err := s.Get("people.json")
	if err != nil {
		return err, fmt.Errorf("failed to get file from S3, %w", err)
	}

	var records []Person
	err = json.Unmarshal(data, &records)
	plugin.Logger(ctx).Info("records", records)
	if err != nil {
		return err, fmt.Errorf("failed to parse file, %w", err)
	}

	for _, record := range records {
		d.StreamListItem(ctx, record)
	}

	return nil, nil

}
