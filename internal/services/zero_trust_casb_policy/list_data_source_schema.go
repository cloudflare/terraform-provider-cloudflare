// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_casb_policy

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/schemata"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ datasource.DataSourceWithConfigValidators = (*ZeroTrustCasbPoliciesDataSource)(nil)

func ListDataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		MarkdownDescription: schemata.Description{
			Scopes: []string{
				"Zero Trust Read",
				"Zero Trust Write",
			},
		}.String(),
		Attributes: map[string]schema.Attribute{
			"account_id": schema.StringAttribute{
				Required: true,
			},
			"max_items": schema.Int64Attribute{
				Description: "Max items to fetch, default: 1000",
				Optional:    true,
				Validators: []validator.Int64{
					int64validator.AtLeast(0),
				},
			},
			"result": schema.ListNestedAttribute{
				Description: "The items returned by the data source",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectListType[ZeroTrustCasbPoliciesResultDataSourceModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Description: "Unique identifier for the policy configuration.",
							Computed:    true,
						},
						"actions": schema.SingleNestedAttribute{
							Description: "The actions configured for this policy.",
							Computed:    true,
							CustomType:  customfield.NewNestedObjectType[ZeroTrustCasbPoliciesActionsDataSourceModel](ctx),
							Attributes: map[string]schema.Attribute{
								"remediation_types": schema.ListNestedAttribute{
									Description: "List of remediation types that will be executed.",
									Computed:    true,
									CustomType:  customfield.NewNestedObjectListType[ZeroTrustCasbPoliciesActionsRemediationTypesDataSourceModel](ctx),
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"display_name": schema.StringAttribute{
												Description: "Display name/label of the remediation type.",
												Computed:    true,
											},
											"remediation_type": schema.StringAttribute{
												Description: "The system name of the remediation type.",
												Computed:    true,
											},
											"remediation_type_id": schema.StringAttribute{
												Description: "Unique identifier for the remediation type.",
												Computed:    true,
											},
										},
									},
								},
								"webhook_configs": schema.ListNestedAttribute{
									Description: "List of webhook configurations that will be triggered.",
									Computed:    true,
									CustomType:  customfield.NewNestedObjectListType[ZeroTrustCasbPoliciesActionsWebhookConfigsDataSourceModel](ctx),
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"display_name": schema.StringAttribute{
												Description: "Display name/label of the webhook configuration.",
												Computed:    true,
											},
											"webhook_config_id": schema.StringAttribute{
												Description: "Unique identifier for the webhook configuration.",
												Computed:    true,
											},
										},
									},
								},
							},
						},
						"applies_to_all_integrations": schema.BoolAttribute{
							Description: "When true, the policy applies to all integrations for the account. When false, it applies only to the specified integration_ids.",
							Computed:    true,
						},
						"created_at": schema.StringAttribute{
							Description: "Timestamp when the policy was created.",
							Computed:    true,
							CustomType:  timetypes.RFC3339Type{},
						},
						"description": schema.StringAttribute{
							Description: "User-set description of what this policy does. Limited to 1000 characters.",
							Computed:    true,
						},
						"display_name": schema.StringAttribute{
							Description: "Display name for the policy configuration. Limited to 255 characters.",
							Computed:    true,
						},
						"enabled": schema.BoolAttribute{
							Description: "Whether the policy is enabled. Derived from disabled_at (enabled when disabled_at is unset).",
							Computed:    true,
						},
						"finding_type_id": schema.StringAttribute{
							Description: "The finding type this policy is associated with. Immutable after creation; changing it replaces the policy.",
							Computed:    true,
						},
						"integration_ids": schema.ListAttribute{
							Description: "The integrations this policy applies to.",
							Computed:    true,
							CustomType:  customfield.NewListType[types.String](ctx),
							ElementType: types.StringType,
						},
						"updated_at": schema.StringAttribute{
							Description: "Timestamp when the policy was last updated.",
							Computed:    true,
							CustomType:  timetypes.RFC3339Type{},
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
					},
				},
			},
		},
	}
}

func (d *ZeroTrustCasbPoliciesDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = ListDataSourceSchema(ctx)
}

func (d *ZeroTrustCasbPoliciesDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
