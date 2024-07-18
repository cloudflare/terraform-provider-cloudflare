// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package custom_hostname

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/float64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ datasource.DataSourceWithConfigValidators = &CustomHostnameDataSource{}
var _ datasource.DataSourceWithValidateConfig = &CustomHostnameDataSource{}

func (r CustomHostnameDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"zone_id": schema.StringAttribute{
				Description: "Identifier",
				Optional:    true,
			},
			"custom_hostname_id": schema.StringAttribute{
				Description: "Identifier",
				Optional:    true,
			},
			"id": schema.StringAttribute{
				Description: "Identifier",
				Computed:    true,
			},
			"hostname": schema.StringAttribute{
				Description: "The custom hostname that will point to your hostname via CNAME.",
				Computed:    true,
			},
			"created_at": schema.StringAttribute{
				Description: "This is the time the hostname was created.",
				Computed:    true,
				Optional:    true,
			},
			"custom_metadata": schema.SingleNestedAttribute{
				Description: "These are per-hostname (customer) settings.",
				Computed:    true,
				Optional:    true,
				Attributes: map[string]schema.Attribute{
					"key": schema.StringAttribute{
						Description: "Unique metadata for this hostname.",
						Computed:    true,
						Optional:    true,
					},
				},
			},
			"custom_origin_server": schema.StringAttribute{
				Description: "a valid hostname that’s been added to your DNS zone as an A, AAAA, or CNAME record.",
				Computed:    true,
				Optional:    true,
			},
			"custom_origin_sni": schema.StringAttribute{
				Description: "A hostname that will be sent to your custom origin server as SNI for TLS handshake. This can be a valid subdomain of the zone or custom origin server name or the string ':request_host_header:' which will cause the host header in the request to be used as SNI. Not configurable with default/fallback origin server.",
				Computed:    true,
				Optional:    true,
			},
			"ownership_verification": schema.SingleNestedAttribute{
				Description: "This is a record which can be placed to activate a hostname.",
				Computed:    true,
				Optional:    true,
				Attributes: map[string]schema.Attribute{
					"name": schema.StringAttribute{
						Description: "DNS Name for record.",
						Computed:    true,
						Optional:    true,
					},
					"type": schema.StringAttribute{
						Description: "DNS Record type.",
						Computed:    true,
						Optional:    true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive("txt"),
						},
					},
					"value": schema.StringAttribute{
						Description: "Content for the record.",
						Computed:    true,
						Optional:    true,
					},
				},
			},
			"ownership_verification_http": schema.SingleNestedAttribute{
				Description: "This presents the token to be served by the given http url to activate a hostname.",
				Computed:    true,
				Optional:    true,
				Attributes: map[string]schema.Attribute{
					"http_body": schema.StringAttribute{
						Description: "Token to be served.",
						Computed:    true,
						Optional:    true,
					},
					"http_url": schema.StringAttribute{
						Description: "The HTTP URL that will be checked during custom hostname verification and where the customer should host the token.",
						Computed:    true,
						Optional:    true,
					},
				},
			},
			"status": schema.StringAttribute{
				Description: "Status of the hostname's activation.",
				Computed:    true,
				Optional:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("active", "pending", "active_redeploying", "moved", "pending_deletion", "deleted", "pending_blocked", "pending_migration", "pending_provisioned", "test_pending", "test_active", "test_active_apex", "test_blocked", "test_failed", "provisioned", "blocked"),
				},
			},
			"verification_errors": schema.ListAttribute{
				Description: "These are errors that were encountered while trying to activate a hostname.",
				Computed:    true,
				Optional:    true,
				ElementType: types.StringType,
			},
			"find_one_by": schema.SingleNestedAttribute{
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"zone_id": schema.StringAttribute{
						Description: "Identifier",
						Required:    true,
					},
					"id": schema.StringAttribute{
						Description: "Hostname ID to match against. This ID was generated and returned during the initial custom_hostname creation. This parameter cannot be used with the 'hostname' parameter.",
						Optional:    true,
					},
					"direction": schema.StringAttribute{
						Description: "Direction to order hostnames.",
						Optional:    true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive("asc", "desc"),
						},
					},
					"hostname": schema.StringAttribute{
						Description: "Fully qualified domain name to match against. This parameter cannot be used with the 'id' parameter.",
						Optional:    true,
					},
					"order": schema.StringAttribute{
						Description: "Field to order hostnames by.",
						Computed:    true,
						Optional:    true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive("ssl", "ssl_status"),
						},
					},
					"page": schema.Float64Attribute{
						Description: "Page number of paginated results.",
						Computed:    true,
						Optional:    true,
					},
					"per_page": schema.Float64Attribute{
						Description: "Number of hostnames per page.",
						Computed:    true,
						Optional:    true,
					},
					"ssl": schema.Float64Attribute{
						Description: "Whether to filter hostnames based on if they have SSL enabled.",
						Optional:    true,
						Validators: []validator.Float64{
							float64validator.OneOf(0, 1),
						},
					},
				},
			},
		},
	}
}

func (r *CustomHostnameDataSource) ConfigValidators(ctx context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}

func (r *CustomHostnameDataSource) ValidateConfig(ctx context.Context, req datasource.ValidateConfigRequest, resp *datasource.ValidateConfigResponse) {
}
