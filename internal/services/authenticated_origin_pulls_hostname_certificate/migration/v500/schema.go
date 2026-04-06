package v500

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

// SourceCloudflareAuthenticatedOriginPullsCertificateSchema returns the v4 schema for
// cloudflare_authenticated_origin_pulls_certificate resource (the unified certificate resource)
// This represents the source schema for MoveState operations from v4 per-hostname certificates
func SourceCloudflareAuthenticatedOriginPullsCertificateSchema() schema.Schema {
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
				Computed:   true,
				CustomType: timetypes.RFC3339Type{},
			},
			"status": schema.StringAttribute{
				Computed: true,
			},
			"uploaded_on": schema.StringAttribute{
				Computed:   true,
				CustomType: timetypes.RFC3339Type{},
			},
		},
	}
}

// ResourceSchema returns the v5 schema for authenticated_origin_pulls_hostname_certificate
// at schema version 1 (early v5 format before bump to 500)
func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Version: 1, // v5 initial schema version (before 500)
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
			"expires_on": schema.StringAttribute{
				Computed:   true,
				CustomType: timetypes.RFC3339Type{},
			},
			"issuer": schema.StringAttribute{
				Computed: true,
			},
			"serial_number": schema.StringAttribute{
				Computed: true,
			},
			"signature": schema.StringAttribute{
				Computed: true,
			},
			"status": schema.StringAttribute{
				Computed: true,
			},
			"uploaded_on": schema.StringAttribute{
				Computed:   true,
				CustomType: timetypes.RFC3339Type{},
			},
		},
	}
}
