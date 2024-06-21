// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package device_posture_integration

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

func (r DevicePostureIntegrationResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:   "API UUID.",
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"account_id": schema.StringAttribute{
				Required: true,
			},
			"config": schema.SingleNestedAttribute{
				Description: "The configuration object containing third-party integration information.",
				Required:    true,
				Attributes: map[string]schema.Attribute{
					"api_url": schema.StringAttribute{
						Description: "The Workspace One API URL provided in the Workspace One Admin Dashboard.",
						Optional:    true,
					},
					"auth_url": schema.StringAttribute{
						Description: "The Workspace One Authorization URL depending on your region.",
						Optional:    true,
					},
					"client_id": schema.StringAttribute{
						Description: "The Workspace One client ID provided in the Workspace One Admin Dashboard.",
						Optional:    true,
					},
					"client_secret": schema.StringAttribute{
						Description: "The Workspace One client secret provided in the Workspace One Admin Dashboard.",
						Required:    true,
					},
					"customer_id": schema.StringAttribute{
						Description: "The Crowdstrike customer ID.",
						Optional:    true,
					},
					"client_key": schema.StringAttribute{
						Description: "The Uptycs client secret.",
						Optional:    true,
					},
					"access_client_id": schema.StringAttribute{
						Description: "If present, this id will be passed in the `CF-Access-Client-ID` header when hitting the `api_url`",
						Optional:    true,
					},
					"access_client_secret": schema.StringAttribute{
						Description: "If present, this secret will be passed in the `CF-Access-Client-Secret` header when hitting the `api_url`",
						Optional:    true,
					},
				},
			},
			"interval": schema.StringAttribute{
				Description: "The interval between each posture check with the third-party API. Use `m` for minutes (e.g. `5m`) and `h` for hours (e.g. `12h`).",
				Required:    true,
			},
			"name": schema.StringAttribute{
				Description: "The name of the device posture integration.",
				Required:    true,
			},
			"type": schema.StringAttribute{
				Description: "The type of device posture integration.",
				Required:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("workspace_one", "crowdstrike_s2s", "uptycs", "intune", "kolide", "tanium", "sentinelone_s2s"),
				},
			},
		},
	}
}
