// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package byo_ip_prefix

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

var _ datasource.DataSourceWithConfigValidators = &ByoIPPrefixesDataSource{}
var _ datasource.DataSourceWithValidateConfig = &ByoIPPrefixesDataSource{}

func (r ByoIPPrefixesDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"account_id": schema.StringAttribute{
				Description: "Identifier",
				Required:    true,
			},
			"max_items": schema.Int64Attribute{
				Description: "Max items to fetch, default: 1000",
				Optional:    true,
			},
			"result": schema.ListNestedAttribute{
				Description: "The items returned by the data source",
				Computed:    true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Description: "Identifier",
							Computed:    true,
							Optional:    true,
						},
						"account_id": schema.StringAttribute{
							Description: "Identifier",
							Computed:    true,
							Optional:    true,
						},
						"advertised": schema.BoolAttribute{
							Description: "Prefix advertisement status to the Internet. This field is only not 'null' if on demand is enabled.",
							Computed:    true,
							Optional:    true,
						},
						"advertised_modified_at": schema.StringAttribute{
							Description: "Last time the advertisement status was changed. This field is only not 'null' if on demand is enabled.",
							Computed:    true,
							CustomType:  timetypes.RFC3339Type{},
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
						"created_at": schema.StringAttribute{
							Computed:   true,
							CustomType: timetypes.RFC3339Type{},
						},
						"description": schema.StringAttribute{
							Description: "Description of the prefix.",
							Computed:    true,
							Optional:    true,
						},
						"loa_document_id": schema.StringAttribute{
							Description: "Identifier for the uploaded LOA document.",
							Computed:    true,
							Optional:    true,
						},
						"modified_at": schema.StringAttribute{
							Computed:   true,
							CustomType: timetypes.RFC3339Type{},
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
					},
				},
			},
		},
	}
}

func (r *ByoIPPrefixesDataSource) ConfigValidators(ctx context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}

func (r *ByoIPPrefixesDataSource) ValidateConfig(ctx context.Context, req datasource.ValidateConfigRequest, resp *datasource.ValidateConfigResponse) {
}
