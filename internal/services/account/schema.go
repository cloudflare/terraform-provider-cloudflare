// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package account

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ resource.ResourceWithConfigValidators = (*AccountResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:   "Identifier",
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"unit": schema.SingleNestedAttribute{
				Description: "information related to the tenant unit, and optionally, an id of the unit to create the account on. see https://developers.cloudflare.com/tenant/how-to/manage-accounts/",
				Optional:    true,
				Computed:    true,
				CustomType:  customfield.NewNestedObjectType[AccountUnitModel](ctx),
				PlanModifiers: []planmodifier.Object{
					objectplanmodifier.UseStateForUnknown(),
				},
				Attributes: map[string]schema.Attribute{
					"id": schema.StringAttribute{
						Description: "Tenant unit ID",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
				},
			},
			"name": schema.StringAttribute{
				Description: "Account name",
				Required:    true,
			},
			"type": schema.StringAttribute{
				Description:        `Available values: "standard", "enterprise".`,
				Optional:           true,
				Computed:           true,
				DeprecationMessage: "The 'type' field should no longer be set through the API.",
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("standard", "enterprise"),
				},
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"managed_by": schema.SingleNestedAttribute{
				Description: "Parent container details",
				Computed:    true,
				Optional:    true,
				CustomType:  customfield.NewNestedObjectType[AccountManagedByModel](ctx),
				Attributes: map[string]schema.Attribute{
					"parent_org_id": schema.StringAttribute{
						Description: "ID of the parent Organization, if one exists",
						Computed:    true,
					},
					"parent_org_name": schema.StringAttribute{
						Description: "Name of the parent Organization, if one exists",
						Computed:    true,
					},
				},
			},
			"settings": schema.SingleNestedAttribute{
				Description: "Account settings",
				Computed:    true,
				Optional:    true,
				CustomType:  customfield.NewNestedObjectType[AccountSettingsModel](ctx),
				Attributes: map[string]schema.Attribute{
					"abuse_contact_email": schema.StringAttribute{
						Description: "Sets an abuse contact email to notify for abuse reports.",
						Optional:    true,
					},
					"enforce_twofactor": schema.BoolAttribute{
						Description: "Indicates whether membership in this account requires that\nTwo-Factor Authentication is enabled",
						Computed:    true,
						Optional:    true,
						Default:     booldefault.StaticBool(false),
					},
				},
			},
			"created_on": schema.StringAttribute{
				Description: "Timestamp for the creation of the account",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
		},
	}
}

func (r *AccountResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *AccountResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}
