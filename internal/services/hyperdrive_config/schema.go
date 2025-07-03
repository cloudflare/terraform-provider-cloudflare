// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package hyperdrive_config

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ resource.ResourceWithConfigValidators = (*HyperdriveConfigResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:   "Define configurations using a unique string identifier.",
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"account_id": schema.StringAttribute{
				Description:   "Define configurations using a unique string identifier.",
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"name": schema.StringAttribute{
				Required: true,
			},
			"origin": schema.SingleNestedAttribute{
				Required: true,
				Attributes: map[string]schema.Attribute{
					"database": schema.StringAttribute{
						Description: "Set the name of your origin database.",
						Required:    true,
					},
					"host": schema.StringAttribute{
						Description: "Defines the host (hostname or IP) of your origin database.",
						Required:    true,
					},
					"password": schema.StringAttribute{
						Description: "Set the password needed to access your origin database. The API never returns this write-only value.",
						Required:    true,
						Sensitive:   true,
					},
					"port": schema.Int64Attribute{
						Description: "Defines the port (default: 5432 for Postgres) of your origin database.",
						Optional:    true,
					},
					"scheme": schema.StringAttribute{
						Description: "Specifies the URL scheme used to connect to your origin database.\nAvailable values: \"postgres\", \"postgresql\", \"mysql\".",
						Required:    true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive(
								"postgres",
								"postgresql",
								"mysql",
							),
						},
					},
					"user": schema.StringAttribute{
						Description: "Set the user of your origin database.",
						Required:    true,
					},
					"access_client_id": schema.StringAttribute{
						Description: "Defines the Client ID of the Access token to use when connecting to the origin database.",
						Optional:    true,
					},
					"access_client_secret": schema.StringAttribute{
						Description: "Defines the Client Secret of the Access Token to use when connecting to the origin database. The API never returns this write-only value.",
						Optional:    true,
						Sensitive:   true,
					},
				},
			},
			"origin_connection_limit": schema.Int64Attribute{
				Description: "The (soft) maximum number of connections the Hyperdrive is allowed to make to the origin database.",
				Optional:    true,
				Validators: []validator.Int64{
					int64validator.Between(5, 100),
				},
			},
			"caching": schema.SingleNestedAttribute{
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"disabled": schema.BoolAttribute{
						Description: "Set to true to disable caching of SQL responses. Default is false.",
						Optional:    true,
						Computed:    true,
						Default:     booldefault.StaticBool(false),
					},
					"max_age": schema.Int64Attribute{
						Description: "Specify the maximum duration items should persist in the cache. Not returned if set to the default (60).",
						Optional:    true,
					},
					"stale_while_revalidate": schema.Int64Attribute{
						Description: "Specify the number of seconds the cache may serve a stale response. Omitted if set to the default (15).",
						Optional:    true,
					},
				},
			},
			"mtls": schema.SingleNestedAttribute{
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"ca_certificate_id": schema.StringAttribute{
						Description: "Define CA certificate ID obtained after uploading CA cert.",
						Optional:    true,
					},
					"mtls_certificate_id": schema.StringAttribute{
						Description: "Define mTLS certificate ID obtained after uploading client cert.",
						Optional:    true,
					},
					"sslmode": schema.StringAttribute{
						Description: "Set SSL mode to 'require', 'verify-ca', or 'verify-full' to verify the CA.",
						Optional:    true,
					},
				},
			},
			"created_on": schema.StringAttribute{
				Description: "Defines the creation time of the Hyperdrive configuration.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"modified_on": schema.StringAttribute{
				Description: "Defines the last modified time of the Hyperdrive configuration.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
		},
	}
}

func (r *HyperdriveConfigResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *HyperdriveConfigResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}
