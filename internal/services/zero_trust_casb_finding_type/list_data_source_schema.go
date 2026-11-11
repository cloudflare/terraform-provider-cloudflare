// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_casb_finding_type

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/schemata"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ datasource.DataSourceWithConfigValidators = (*ZeroTrustCasbFindingTypesDataSource)(nil)

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
			"direction": schema.StringAttribute{
				Description: "Sort direction for the primary field. Defaults to 'desc'.\nTies are resolved by case-insensitive name, then ID, both ascending regardless of direction.\nAvailable values: \"asc\", \"desc\".",
				Optional:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("asc", "desc"),
				},
			},
			"max_updated_at": schema.StringAttribute{
				Description: "Return only finding types updated at or before this RFC 3339 timestamp (inclusive).",
				Optional:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"min_updated_at": schema.StringAttribute{
				Description: "Return only finding types updated at or after this RFC 3339 timestamp (inclusive).",
				Optional:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"order": schema.StringAttribute{
				Description: "Field to sort by (case-insensitive). Defaults to 'created'.\nName and vendor sorting is case-insensitive; vendor uses friendly provider names.\nAscending order: type Content then Posture; origin custom then system;\nseverity Low, Medium, High, Critical.\nAvailable values: \"created\", \"updated\", \"match_count\", \"name\", \"vendor\", \"type\", \"origin\", \"severity\".",
				Optional:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive(
						"created",
						"updated",
						"match_count",
						"name",
						"vendor",
						"type",
						"origin",
						"severity",
					),
				},
			},
			"quarantined": schema.BoolAttribute{
				Description: "When true, return only quarantined finding types; when false, return only non-quarantined finding types. Omit to return both.",
				Optional:    true,
			},
			"search": schema.StringAttribute{
				Description: "Filter finding types by name or ID (case-insensitive substring match).",
				Optional:    true,
			},
			"origins": schema.ListAttribute{
				Description: "Filter finding types by origin. Supports multiple comma-separated, case-insensitive values.",
				Optional:    true,
				Validators: []validator.List{
					listvalidator.ValueStringsAre(
						stringvalidator.OneOfCaseInsensitive("custom", "system"),
					),
				},
				ElementType: types.StringType,
			},
			"severities": schema.ListAttribute{
				Description: "Filter finding types by severity. Supports multiple comma-separated, case-insensitive values.",
				Optional:    true,
				Validators: []validator.List{
					listvalidator.ValueStringsAre(
						stringvalidator.OneOfCaseInsensitive(
							"critical",
							"high",
							"medium",
							"low",
						),
					),
				},
				ElementType: types.StringType,
			},
			"types": schema.ListAttribute{
				Description: "Filter finding types by category type. Supports multiple comma-separated, case-insensitive values.",
				Optional:    true,
				Validators: []validator.List{
					listvalidator.ValueStringsAre(
						stringvalidator.OneOfCaseInsensitive("Content", "Posture"),
					),
				},
				ElementType: types.StringType,
			},
			"vendors": schema.ListAttribute{
				Description: "Filter finding types by vendor. Supports multiple comma-separated values. Each value may be either the vendor enum (GOOGLE_WORKSPACE) or the friendly name returned by the read routes (Google Workspace).",
				Optional:    true,
				Validators: []validator.List{
					listvalidator.ValueStringsAre(
						stringvalidator.OneOfCaseInsensitive(
							"ANTHROPIC",
							"AWS",
							"BITBUCKET",
							"BOX",
							"CONFLUENCE",
							"DROPBOX",
							"GITHUB",
							"GOOGLE_CLOUD_PLATFORM",
							"GOOGLE_WORKSPACE",
							"JIRA",
							"MICROSOFT",
							"OPENAI",
							"SALESFORCE",
							"SERVICENOW",
							"SLACK",
							"ZOOM",
						),
					),
				},
				ElementType: types.StringType,
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
				CustomType:  customfield.NewNestedObjectListType[ZeroTrustCasbFindingTypesResultDataSourceModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Description: "The unique identifier of the finding.",
							Computed:    true,
						},
						"category": schema.SingleNestedAttribute{
							Description: "Category information for a finding.",
							Computed:    true,
							CustomType:  customfield.NewNestedObjectType[ZeroTrustCasbFindingTypesCategoryDataSourceModel](ctx),
							Attributes: map[string]schema.Attribute{
								"observation": schema.StringAttribute{
									Description: "The type of the observation.\nAvailable values: \"Issue\", \"Insight\", \"Activity\".",
									Computed:    true,
									Validators: []validator.String{
										stringvalidator.OneOfCaseInsensitive(
											"Issue",
											"Insight",
											"Activity",
										),
									},
								},
								"product": schema.StringAttribute{
									Description: "The product category.\nAvailable values: \"SaaS\", \"Cloud\".",
									Computed:    true,
									Validators: []validator.String{
										stringvalidator.OneOfCaseInsensitive("SaaS", "Cloud"),
									},
								},
								"type": schema.StringAttribute{
									Description: "The type of the finding category.\nAvailable values: \"Content\", \"Posture\".",
									Computed:    true,
									Validators: []validator.String{
										stringvalidator.OneOfCaseInsensitive("Content", "Posture"),
									},
								},
							},
						},
						"name": schema.StringAttribute{
							Description: "The name of the finding.",
							Computed:    true,
						},
						"severity": schema.StringAttribute{
							Description: "Default severity used when no integration-specific severity override exists.\nAvailable values: \"Critical\", \"High\", \"Medium\", \"Low\".",
							Computed:    true,
							Validators: []validator.String{
								stringvalidator.OneOfCaseInsensitive(
									"Critical",
									"High",
									"Medium",
									"Low",
								),
							},
						},
						"vendor": schema.StringAttribute{
							Description: "The SaaS/Cloud vendor of the platform with which the finding is associated.",
							Computed:    true,
						},
						"applies_to_all_integrations": schema.BoolAttribute{
							Description: "Whether this finding type applies to every current and future integration of the vendor.",
							Computed:    true,
						},
						"asset_class": schema.StringAttribute{
							Description: "Asset class the finding type's policy evaluates.",
							Computed:    true,
						},
						"asset_vendor": schema.StringAttribute{
							Description: "Canonical vendor enum value.\nAvailable values: \"ANTHROPIC\", \"AWS\", \"BITBUCKET\", \"BOX\", \"CONFLUENCE\", \"DROPBOX\", \"GITHUB\", \"GOOGLE_CLOUD_PLATFORM\", \"GOOGLE_WORKSPACE\", \"JIRA\", \"MICROSOFT\", \"OPENAI\", \"SALESFORCE\", \"SERVICENOW\", \"SLACK\", \"ZOOM\".",
							Computed:    true,
							Validators: []validator.String{
								stringvalidator.OneOfCaseInsensitive(
									"ANTHROPIC",
									"AWS",
									"BITBUCKET",
									"BOX",
									"CONFLUENCE",
									"DROPBOX",
									"GITHUB",
									"GOOGLE_CLOUD_PLATFORM",
									"GOOGLE_WORKSPACE",
									"JIRA",
									"MICROSOFT",
									"OPENAI",
									"SALESFORCE",
									"SERVICENOW",
									"SLACK",
									"ZOOM",
								),
							},
						},
						"created_at": schema.StringAttribute{
							Description: "Timestamp when the finding type was created.",
							Computed:    true,
							CustomType:  timetypes.RFC3339Type{},
						},
						"description": schema.StringAttribute{
							Description: "Detailed description of the finding.",
							Computed:    true,
						},
						"expression": schema.StringAttribute{
							Description: "Rego expression backing this finding type. For custom finding types, generated package and import declarations are omitted. System expressions are returned unchanged.",
							Computed:    true,
						},
						"integration_ids": schema.ListAttribute{
							Description: "Integrations this finding type applies to when applies_to_all_integrations is false.",
							Computed:    true,
							CustomType:  customfield.NewListType[types.String](ctx),
							ElementType: types.StringType,
						},
						"match_count": schema.Int64Attribute{
							Description: "Number of currently active finding instances across the account's integrations. Zero when none have been detected.",
							Computed:    true,
						},
						"origin": schema.StringAttribute{
							Description: "Whether the finding type is system-managed or customer-created.\nAvailable values: \"system\", \"custom\".",
							Computed:    true,
							Validators: []validator.String{
								stringvalidator.OneOfCaseInsensitive("system", "custom"),
							},
						},
						"quarantined": schema.BoolAttribute{
							Description: "Whether evaluation of this finding type is suspended by an operator. Present only when true.",
							Computed:    true,
						},
						"updated_at": schema.StringAttribute{
							Description: "Timestamp when the finding type was last updated.",
							Computed:    true,
							CustomType:  timetypes.RFC3339Type{},
						},
						"version": schema.Int64Attribute{
							Description: "Current version of the finding type. Incremented on each update.",
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func (d *ZeroTrustCasbFindingTypesDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = ListDataSourceSchema(ctx)
}

func (d *ZeroTrustCasbFindingTypesDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
