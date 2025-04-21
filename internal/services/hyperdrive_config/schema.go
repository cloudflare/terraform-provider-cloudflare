// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package hyperdrive_config

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
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
				Description:   "Identifier",
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"account_id": schema.StringAttribute{
				Description:   "Identifier",
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
						Description: "The name of your origin database.",
						Required:    true,
					},
					"host": schema.StringAttribute{
						Description: "The host (hostname or IP) of your origin database.",
						Required:    true,
					},
					"password": schema.StringAttribute{
						Description: "The password required to access your origin database. This value is write-only and never returned by the API.",
						Required:    true,
						Sensitive:   true,
					},
					"port": schema.Int64Attribute{
						Description: "The port (default: 5432 for Postgres) of your origin database.",
						Optional:    true,
					},
					"scheme": schema.StringAttribute{
						Description: "Specifies the URL scheme used to connect to your origin database.\nAvailable values: \"postgres\", \"postgresql\", \"mysql\".",
						Required:    true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive("postgres", "postgresql", "mysql"),
						},
					},
					"user": schema.StringAttribute{
						Description: "The user of your origin database.",
						Required:    true,
					},
					"access_client_id": schema.StringAttribute{
						Description: "The Client ID of the Access token to use when connecting to the origin database.",
						Optional:    true,
					},
					"access_client_secret": schema.StringAttribute{
						Description: "The Client Secret of the Access token to use when connecting to the origin database. This value is write-only and never returned by the API.",
						Optional:    true,
						Sensitive:   true,
					},
				},
			},
			"caching": schema.SingleNestedAttribute{
				Computed:   true,
				Optional:   true,
				CustomType: customfield.NewNestedObjectType[HyperdriveConfigCachingModel](ctx),
				Attributes: map[string]schema.Attribute{
					"disabled": schema.BoolAttribute{
						Description: "When set to true, disables the caching of SQL responses. (Default: false)",
						Optional:    true,
						Computed:    true,
						Default:     booldefault.StaticBool(false),
					},
					"max_age": schema.Int64Attribute{
						Description: "When present, specifies max duration for which items should persist in the cache. Not returned if set to default. (Default: 60)",
						Optional:    true,
					},
					"stale_while_revalidate": schema.Int64Attribute{
						Description: "When present, indicates the number of seconds cache may serve the response after it becomes stale. Not returned if set to default. (Default: 15)",
						Optional:    true,
					},
				},
			},
			"created_on": schema.StringAttribute{
				Description: "When the Hyperdrive configuration was created.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"modified_on": schema.StringAttribute{
				Description: "When the Hyperdrive configuration was last modified.",
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
