// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package email_security_allow_pattern

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ resource.ResourceWithConfigValidators = (*EmailSecurityAllowPatternResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.Int64Attribute{
				Computed:      true,
				PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
			},
			"account_id": schema.StringAttribute{
				Description:   "Account Identifier",
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"body": schema.ListNestedAttribute{
				Computed:   true,
				Optional:   true,
				CustomType: customfield.NewNestedObjectListType[EmailSecurityAllowPatternBodyModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"is_recipient": schema.BoolAttribute{
							Required: true,
						},
						"is_regex": schema.BoolAttribute{
							Required: true,
						},
						"is_sender": schema.BoolAttribute{
							Required: true,
						},
						"is_spoof": schema.BoolAttribute{
							Required: true,
						},
						"pattern": schema.StringAttribute{
							Required: true,
						},
						"pattern_type": schema.StringAttribute{
							Required: true,
							Validators: []validator.String{
								stringvalidator.OneOfCaseInsensitive(
									"EMAIL",
									"DOMAIN",
									"IP",
									"UNKNOWN",
								),
							},
						},
						"verify_sender": schema.BoolAttribute{
							Required: true,
						},
						"comments": schema.StringAttribute{
							Optional: true,
						},
					},
				},
				PlanModifiers: []planmodifier.List{listplanmodifier.RequiresReplace()},
			},
			"comments": schema.StringAttribute{
				Optional: true,
			},
			"is_recipient": schema.BoolAttribute{
				Optional: true,
			},
			"is_regex": schema.BoolAttribute{
				Optional: true,
			},
			"is_sender": schema.BoolAttribute{
				Optional: true,
			},
			"is_spoof": schema.BoolAttribute{
				Optional: true,
			},
			"pattern": schema.StringAttribute{
				Optional: true,
			},
			"pattern_type": schema.StringAttribute{
				Optional: true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive(
						"EMAIL",
						"DOMAIN",
						"IP",
						"UNKNOWN",
					),
				},
			},
			"verify_sender": schema.BoolAttribute{
				Optional: true,
			},
			"created_at": schema.StringAttribute{
				Computed:   true,
				CustomType: timetypes.RFC3339Type{},
			},
			"last_modified": schema.StringAttribute{
				Computed:   true,
				CustomType: timetypes.RFC3339Type{},
			},
		},
	}
}

func (r *EmailSecurityAllowPatternResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *EmailSecurityAllowPatternResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}
