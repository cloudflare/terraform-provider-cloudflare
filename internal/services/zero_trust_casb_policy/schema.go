// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_casb_policy

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/schemata"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ resource.ResourceWithConfigValidators = (*ZeroTrustCasbPolicyResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		MarkdownDescription: schemata.Description{
			Scopes: []string{
				"Zero Trust Read",
				"Zero Trust Write",
			},
		}.String(),
		Version: 500,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:   "Unique identifier for the policy configuration.",
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseNonNullStateForUnknown()},
			},
			"account_id": schema.StringAttribute{
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"finding_type_id": schema.StringAttribute{
				Description:   "The finding type this policy is associated with. All remediation actions must match this finding type.",
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"applies_to_all_integrations": schema.BoolAttribute{
				Description: "When true, the policy applies to all integrations for the account. When false, integration_ids must be provided.",
				Required:    true,
			},
			"display_name": schema.StringAttribute{
				Description: "Display name for the policy configuration.",
				Required:    true,
			},
			"enabled": schema.BoolAttribute{
				Description: "Boolean specifying if the policy is enabled or disabled.",
				Required:    true,
			},
			"actions": schema.SingleNestedAttribute{
				Description: "Actions to execute when this policy is triggered, grouped by action type.\nA policy must contain at least one action across all groups and may include\nat most one remediation.",
				Required:    true,
				Attributes: map[string]schema.Attribute{
					"remediation_types": schema.ListNestedAttribute{
						Description: "Remediation actions to execute (at most one).",
						Optional:    true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"remediation_type_id": schema.StringAttribute{
									Description: "The ID of the remediation type to execute.",
									Required:    true,
								},
							},
						},
					},
					"webhook_configs": schema.ListNestedAttribute{
						Description: "Webhook actions to execute.",
						Optional:    true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"webhook_config_id": schema.StringAttribute{
									Description: "The ID of the webhook configuration to use.",
									Required:    true,
								},
							},
						},
					},
				},
			},
			"description": schema.StringAttribute{
				Description:   "Optional description of what this policy does.",
				Computed:      true,
				Optional:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseNonNullStateForUnknown()},
			},
			"integration_ids": schema.ListAttribute{
				Description:   "The integrations this policy applies to. Required when applies_to_all_integrations is false.",
				Computed:      true,
				Optional:      true,
				CustomType:    customfield.NewListType[types.String](ctx),
				ElementType:   types.StringType,
				PlanModifiers: []planmodifier.List{listplanmodifier.UseNonNullStateForUnknown()},
			},
			"created_at": schema.StringAttribute{
				Description:   "Timestamp when the policy was created.",
				Computed:      true,
				CustomType:    timetypes.RFC3339Type{},
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseNonNullStateForUnknown()},
			},
			"disabled_at": schema.StringAttribute{
				Description: "Timestamp when the policy was disabled. Omitted from the response when the policy\nis enabled.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"last_triggered_at": schema.StringAttribute{
				Description: "Timestamp of the most recent successful policy invocation. Omitted\nfrom the response when the policy has never been successfully\ntriggered. Only populated on GET responses; absent on responses from\ncreate/update endpoints.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"updated_at": schema.StringAttribute{
				Description: "Timestamp when the policy was last updated.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
		},
	}
}

func (r *ZeroTrustCasbPolicyResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *ZeroTrustCasbPolicyResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}
