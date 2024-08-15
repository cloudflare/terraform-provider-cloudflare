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

var _ datasource.DataSourceWithConfigValidators = &Web3HostnameDataSource{}

func (d *Web3HostnameDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"identifier": schema.StringAttribute{
				Description: "Identifier",
				Optional:    true,
			},
			"zone_identifier": schema.StringAttribute{
				Description: "Identifier",
				Optional:    true,
			},
			"created_on": schema.StringAttribute{
				Computed:   true,
				CustomType: timetypes.RFC3339Type{},
			},
			"id": schema.StringAttribute{
				Description: "Identifier",
				Computed:    true,
			},
			"modified_on": schema.StringAttribute{
				Computed:   true,
				CustomType: timetypes.RFC3339Type{},
			},
			"name": schema.StringAttribute{
				Description: "The hostname that will point to the target gateway via CNAME.",
				Computed:    true,
			},
			"status": schema.StringAttribute{
				Description: "Status of the hostname's activation.",
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
			"description": schema.StringAttribute{
				Description: "An optional description of the hostname.",
				Computed:    true,
				Optional:    true,
			},
			"dnslink": schema.StringAttribute{
				Description: "DNSLink value used if the target is ipfs.",
				Computed:    true,
				Optional:    true,
			},
			"target": schema.StringAttribute{
				Description: "Target gateway of the hostname.",
				Computed:    true,
				Optional:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive(
						"ethereum",
						"ipfs",
						"ipfs_universal_path",
					),
				},
			},
			"filter": schema.SingleNestedAttribute{
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"zone_identifier": schema.StringAttribute{
						Description: "Identifier",
						Required:    true,
					},
				},
			},
		},
	}
}

func (d *Web3HostnameDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
