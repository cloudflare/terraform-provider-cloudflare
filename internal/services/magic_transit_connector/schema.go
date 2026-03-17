// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package magic_transit_connector

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-validators/float64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ resource.ResourceWithConfigValidators = (*MagicTransitConnectorResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Version: 500,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"account_id": schema.StringAttribute{
				Description:   "Account identifier",
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"device": schema.SingleNestedAttribute{
				Description: "Exactly one of id, serial_number, or provision_license must be provided.",
				Required:    true,
				Attributes: map[string]schema.Attribute{
					"id": schema.StringAttribute{
						Computed: true,
						Optional: true,
					},
					"provision_license": schema.BoolAttribute{
						Description: "When true, create and provision a new licence key for the connector.",
						Optional:    true,
					},
					"serial_number": schema.StringAttribute{
						Computed: true,
						Optional: true,
					},
				},
				PlanModifiers: []planmodifier.Object{objectplanmodifier.RequiresReplace()},
			},
			"provision_license": schema.BoolAttribute{
				Description: "When true, regenerate license key for the connector.",
				Optional:    true,
			},
			"activated": schema.BoolAttribute{
				Computed: true,
				Optional: true,
			},
			"interrupt_window_duration_hours": schema.Float64Attribute{
				Computed: true,
				Optional: true,
				Validators: []validator.Float64{
					float64validator.Between(1, 24),
				},
			},
			"interrupt_window_hour_of_day": schema.Float64Attribute{
				Computed: true,
				Optional: true,
			},
			"notes": schema.StringAttribute{
				Computed: true,
				Optional: true,
			},
			"timezone": schema.StringAttribute{
				Computed: true,
				Optional: true,
			},
			"interrupt_window_days_of_week": schema.ListAttribute{
				Description: "Allowed days of the week for upgrades. Default is all days.",
				Computed:    true,
				Optional:    true,
				Validators: []validator.List{
					listvalidator.ValueStringsAre(
						stringvalidator.OneOfCaseInsensitive(
							"Sunday",
							"Monday",
							"Tuesday",
							"Wednesday",
							"Thursday",
							"Friday",
							"Saturday",
						),
					),
				},
				CustomType:  customfield.NewListType[types.String](ctx),
				ElementType: types.StringType,
			},
			"interrupt_window_embargo_dates": schema.ListAttribute{
				Description: "List of dates (YYYY-MM-DD) when upgrades are blocked.",
				Computed:    true,
				Optional:    true,
				CustomType:  customfield.NewListType[types.String](ctx),
				ElementType: types.StringType,
			},
			"last_heartbeat": schema.StringAttribute{
				Computed: true,
			},
			"last_seen_version": schema.StringAttribute{
				Computed: true,
			},
			"last_updated": schema.StringAttribute{
				Computed: true,
			},
			"license_key": schema.StringAttribute{
				Computed: true,
			},
		},
	}
}

func (r *MagicTransitConnectorResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = CustomResourceSchema(ctx)
}

func (r *MagicTransitConnectorResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}
