package v500

import "github.com/hashicorp/terraform-plugin-framework/resource/schema"

// SourceSchemaV0 returns the v4 cloudflare_worker_domain schema.
func SourceSchemaV0() schema.Schema {
	return schema.Schema{
		Version: 0,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"account_id": schema.StringAttribute{
				Required: true,
			},
			"hostname": schema.StringAttribute{
				Required: true,
			},
			"service": schema.StringAttribute{
				Required: true,
			},
			"zone_id": schema.StringAttribute{
				Required: true,
			},
			"environment": schema.StringAttribute{
				Optional: true,
			},
			"zone_name": schema.StringAttribute{
				Computed: true,
			},
			"cert_id": schema.StringAttribute{
				Computed: true,
				Optional: true,
			},
		},
	}
}
