// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package web3_hostname

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ datasource.DataSourceWithConfigValidators = (*Web3HostnameDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "Specify the identifier of the hostname.",
				Computed:    true,
			},
			"identifier": schema.StringAttribute{
				Description: "Specify the identifier of the hostname.",
				Optional:    true,
			},
			"zone_id": schema.StringAttribute{
				Description: "Specify the identifier of the hostname.",
				Required:    true,
			},
			"created_on": schema.StringAttribute{
				Computed:   true,
				CustomType: timetypes.RFC3339Type{},
			},
			"description": schema.StringAttribute{
				Description: "Specify an optional description of the hostname.",
				Computed:    true,
			},
			"dnslink": schema.StringAttribute{
				Description: "Specify the DNSLink value used if the target is ipfs.",
				Computed:    true,
			},
			"modified_on": schema.StringAttribute{
				Computed:   true,
				CustomType: timetypes.RFC3339Type{},
			},
			"name": schema.StringAttribute{
				Description: "Specify the hostname that points to the target gateway via CNAME.",
				Computed:    true,
			},
			"status": schema.StringAttribute{
				Description: "Specifies the status of the hostname's activation.\nAvailable values: \"active\", \"pending\", \"deleting\", \"error\".",
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive(
						"active",
						"pending",
						"deleting",
						"error",
					),
				},
			},
			"target": schema.StringAttribute{
				Description: "Specify the target gateway of the hostname.\nAvailable values: \"ethereum\", \"ipfs\", \"ipfs_universal_path\".",
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive(
						"ethereum",
						"ipfs",
						"ipfs_universal_path",
					),
				},
			},
		},
	}
}

func (d *Web3HostnameDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *Web3HostnameDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
