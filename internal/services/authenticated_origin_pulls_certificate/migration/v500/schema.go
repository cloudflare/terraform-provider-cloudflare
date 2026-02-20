package v500

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

// ResourceSchema returns the v4 schema for authenticated_origin_pulls_certificate
// This schema version is 0 (from v4 provider)
func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Version: 0, // v4 schema version
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"zone_id": schema.StringAttribute{
				Required: true,
			},
			"certificate": schema.StringAttribute{
				Required: true,
			},
			"private_key": schema.StringAttribute{
				Required: true,
			},
			"type": schema.StringAttribute{
				Required: true,
			},
			"issuer": schema.StringAttribute{
				Computed: true,
			},
			"signature": schema.StringAttribute{
				Computed: true,
			},
			"serial_number": schema.StringAttribute{
				Computed: true,
			},
			"expires_on": schema.StringAttribute{
				Computed: true,
			},
			"status": schema.StringAttribute{
				Computed: true,
			},
			"uploaded_on": schema.StringAttribute{
				Computed: true,
			},
		},
	}
}
