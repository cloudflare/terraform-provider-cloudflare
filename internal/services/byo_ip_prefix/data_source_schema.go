// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package byo_ip_prefix

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/datasourcevalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/path"
)

var _ datasource.DataSourceWithConfigValidators = (*ByoIPPrefixDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"prefix_id": schema.StringAttribute{
				Description: "Identifier",
				Optional:    true,
			},
			"account_id": schema.StringAttribute{
				Description: "Identifier",
				Computed:    true,
				Optional:    true,
			},
			"advertised_modified_at": schema.StringAttribute{
				Description: "Last time the advertisement status was changed. This field is only not 'null' if on demand is enabled.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"created_at": schema.StringAttribute{
				Computed:   true,
				CustomType: timetypes.RFC3339Type{},
			},
			"modified_at": schema.StringAttribute{
				Computed:   true,
				CustomType: timetypes.RFC3339Type{},
			},
			"advertised": schema.BoolAttribute{
				Description: "Prefix advertisement status to the Internet. This field is only not 'null' if on demand is enabled.",
				Computed:    true,
				Optional:    true,
			},
			"approved": schema.StringAttribute{
				Description: "Approval state of the prefix (P = pending, V = active).",
				Computed:    true,
				Optional:    true,
			},
			"asn": schema.Int64Attribute{
				Description: "Autonomous System Number (ASN) the prefix will be advertised under.",
				Computed:    true,
				Optional:    true,
			},
			"cidr": schema.StringAttribute{
				Description: "IP Prefix in Classless Inter-Domain Routing format.",
				Computed:    true,
				Optional:    true,
			},
			"description": schema.StringAttribute{
				Description: "Description of the prefix.",
				Computed:    true,
				Optional:    true,
			},
			"id": schema.StringAttribute{
				Description: "Identifier",
				Computed:    true,
				Optional:    true,
			},
			"loa_document_id": schema.StringAttribute{
				Description: "Identifier for the uploaded LOA document.",
				Computed:    true,
				Optional:    true,
			},
			"on_demand_enabled": schema.BoolAttribute{
				Description: "Whether advertisement of the prefix to the Internet may be dynamically enabled or disabled.",
				Computed:    true,
				Optional:    true,
			},
			"on_demand_locked": schema.BoolAttribute{
				Description: "Whether advertisement status of the prefix is locked, meaning it cannot be changed.",
				Computed:    true,
				Optional:    true,
			},
			"filter": schema.SingleNestedAttribute{
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"account_id": schema.StringAttribute{
						Description: "Identifier",
						Required:    true,
					},
				},
			},
		},
	}
}

func (d *ByoIPPrefixDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *ByoIPPrefixDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{
		datasourcevalidator.RequiredTogether(path.MatchRoot("account_id"), path.MatchRoot("prefix_id")),
		datasourcevalidator.ExactlyOneOf(path.MatchRoot("filter"), path.MatchRoot("account_id")),
		datasourcevalidator.ExactlyOneOf(path.MatchRoot("filter"), path.MatchRoot("prefix_id")),
	}
}
