// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package connectivity_directory_service

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ resource.ResourceWithConfigValidators = (*ConnectivityDirectoryServiceResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Version: 500,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"service_id": schema.StringAttribute{
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"account_id": schema.StringAttribute{
				Description:   "Account identifier",
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"name": schema.StringAttribute{
				Required: true,
			},
			"type": schema.StringAttribute{
				Description: `Available values: "tcp", "http".`,
				Required:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("tcp", "http"),
				},
			},
			"host": schema.SingleNestedAttribute{
				Required: true,
				Attributes: map[string]schema.Attribute{
					"ipv4": schema.StringAttribute{
						Optional: true,
					},
					"network": schema.SingleNestedAttribute{
						Optional: true,
						Attributes: map[string]schema.Attribute{
							"tunnel_id": schema.StringAttribute{
								Required: true,
							},
						},
					},
					"ipv6": schema.StringAttribute{
						Optional: true,
					},
					"hostname": schema.StringAttribute{
						Optional: true,
					},
					"resolver_network": schema.SingleNestedAttribute{
						Optional: true,
						Attributes: map[string]schema.Attribute{
							"tunnel_id": schema.StringAttribute{
								Required: true,
							},
							"resolver_ips": schema.ListAttribute{
								Optional:    true,
								ElementType: types.StringType,
							},
						},
					},
				},
			},
			"app_protocol": schema.StringAttribute{
				Description: `Available values: "postgresql", "mysql".`,
				Optional:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("postgresql", "mysql"),
				},
			},
			"http_port": schema.Int64Attribute{
				Optional: true,
				Validators: []validator.Int64{
					int64validator.AtLeast(1),
				},
			},
			"https_port": schema.Int64Attribute{
				Optional: true,
				Validators: []validator.Int64{
					int64validator.AtLeast(1),
				},
			},
			"tcp_port": schema.Int64Attribute{
				Optional: true,
				Validators: []validator.Int64{
					int64validator.AtLeast(1),
				},
			},
			"tls_settings": schema.SingleNestedAttribute{
				Description: "TLS settings for a connectivity service.\n\nIf omitted, the default mode (`verify_full`) is used.",
				Optional:    true,
				Attributes: map[string]schema.Attribute{
					"cert_verification_mode": schema.StringAttribute{
						Description: "TLS certificate verification mode for the connection to the origin.\n\n- `\"verify_full\"` — verify certificate chain and hostname (default)\n- `\"verify_ca\"` — verify certificate chain only, skip hostname check\n- `\"disabled\"` — do not verify the server certificate at all",
						Required:    true,
					},
				},
			},
			"created_at": schema.StringAttribute{
				Computed:   true,
				CustomType: timetypes.RFC3339Type{},
			},
			"updated_at": schema.StringAttribute{
				Computed:   true,
				CustomType: timetypes.RFC3339Type{},
			},
		},
	}
}

func (r *ConnectivityDirectoryServiceResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *ConnectivityDirectoryServiceResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}
