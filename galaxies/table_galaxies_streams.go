package galaxies

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/guardian/steampipe-plugin-galaxies/store"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

type Stream struct {
	ID          string   `json:"streamId"`
	Name        string   `json:"streamName"`
	Description string   `json:"streamDescription"`
	Members     []string `json:"streamMembers"`
}

func tableStreams() *plugin.Table {
	return &plugin.Table{
		Name:        "galaxies_streams",
		Description: "Streams in the Guardian P&E department",
		List: &plugin.ListConfig{
			Hydrate: getStreams,
		},
		Columns: []*plugin.Column{
			{Name: "id", Type: proto.ColumnType_STRING, Description: "TODO"},
			{Name: "name", Type: proto.ColumnType_STRING, Description: "TODO"},
			{Name: "description", Type: proto.ColumnType_STRING, Description: "TODO"},
			{Name: "members", Type: proto.ColumnType_JSON, Description: "TODO"},
		},
	}
}

func getStreams(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	s, err := store.GalaxiesS3(ctx)
	if err != nil {
		return nil, fmt.Errorf("unable to load AWS config, %w", err)
	}

	data, err := s.Get("streams.json")
	if err != nil {
		return err, fmt.Errorf("failed to get file from S3, %w", err)
	}

	var records map[string]Stream
	err = json.Unmarshal(data, &records)
	plugin.Logger(ctx).Info("records", records)
	if err != nil {
		return err, fmt.Errorf("failed to parse file, %w", err)
	}

	for ID, record := range records {
		record.ID = ID // Galaxies streams data uses the map keys as the ID, so pull this out.
		d.StreamListItem(ctx, record)
	}

	return nil, nil

}
