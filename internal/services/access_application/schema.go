// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package access_application

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:   "UUID.",
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"account_id": schema.StringAttribute{
				Description:   "The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.",
				Optional:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"zone_id": schema.StringAttribute{
				Description:   "The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.",
				Optional:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"name": schema.StringAttribute{
				Description: "Friendly name of the Access Application.",
				Required:    true,
			},
			"domain": schema.StringAttribute{
				Description: "The primary hostname and path secured by Access. This domain will be displayed if the app is visible in the App Launcher.",
				Optional:    true,
				Computed:    true,
			},
			"type": schema.StringAttribute{
				Description: "The application type. Available values: `app_launcher`, `bookmark`, `biso`, `dash_sso`, `saas`, `self_hosted`, `ssh`, `vnc`, `warp`.",
				Optional:    true,
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive(
						"app_launcher",
						"bookmark",
						"biso",
						"dash_sso",
						"saas",
						"self_hosted",
						"ssh",
						"vnc",
						"warp",
					),
				},
				Default: stringdefault.StaticString("self_hosted"),
			},
			"session_duration": schema.StringAttribute{
				Description: "How often a user will be forced to re-authorise. Must be in the format `48h` or `2h45m`. Valid units are `ns`, `us` (or `µs`), `ms`, `s`, `m`, `h`.",
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString("24h"),
			},
			"auto_redirect_to_identity": schema.BoolAttribute{
				Description: "Option to skip identity provider banner during SSO when this application is set as the default one for the account.",
				Optional:    true,
				Computed:    true,
			},
			"enable_binding_cookie": schema.BoolAttribute{
				Description: "Option to provide increased security against compromised authorization tokens and CSRF attacks by requiring an additional \"binding\" cookie when accessing the application.",
				Optional:    true,
				Computed:    true,
			},
			"allowed_idps": schema.ListAttribute{
				ElementType: types.StringType,
				Description: "The identity providers selected for the application.",
				Optional:    true,
			},
			"custom_deny_message": schema.StringAttribute{
				Description: "Option that returns a custom error message when a user is denied access to the application.",
				Optional:    true,
			},
			"custom_deny_url": schema.StringAttribute{
				Description: "Option that redirects to a custom URL when a user is denied access to the application via identity based rules.",
				Optional:    true,
			},
			"custom_non_identity_deny_url": schema.StringAttribute{
				Description: "Option that redirects to a custom URL when a user is denied access to the application via non identity rules.",
				Optional:    true,
			},
			"logo_url": schema.StringAttribute{
				Description: "Image URL for the logo shown in the app launcher dashboard.",
				Optional:    true,
			},
			"created_at": schema.StringAttribute{
				Description: "Timestamp of when the resource was created.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"updated_at": schema.StringAttribute{
				Description: "Timestamp of when the resource was last updated.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
		},
	}
}
