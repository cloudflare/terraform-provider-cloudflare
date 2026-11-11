// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_dlp_proximity_validator

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/schemata"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ datasource.DataSourceWithConfigValidators = (*ZeroTrustDLPProximityValidatorsDataSource)(nil)

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
				CustomType:  customfield.NewNestedObjectListType[ZeroTrustDLPProximityValidatorsResultDataSourceModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Computed: true,
						},
						"created_at": schema.StringAttribute{
							Computed:   true,
							CustomType: timetypes.RFC3339Type{},
						},
						"name": schema.StringAttribute{
							Description: "A name for this validator, at most 50 characters.",
							Computed:    true,
						},
						"referenced_by": schema.SingleNestedAttribute{
							Description: "What currently uses a validator. Deleting a validator that is still in use returns `409`, so\ncheck this first.\n\nRead-only. A validator is attached or detached by updating the detection entry that uses it, not\nby writing this back.",
							Computed:    true,
							CustomType:  customfield.NewNestedObjectType[ZeroTrustDLPProximityValidatorsReferencedByDataSourceModel](ctx),
							Attributes: map[string]schema.Attribute{
								"entries": schema.ListNestedAttribute{
									Description: "The detection entries using this validator, ordered by name.",
									Computed:    true,
									CustomType:  customfield.NewNestedObjectListType[ZeroTrustDLPProximityValidatorsReferencedByEntriesDataSourceModel](ctx),
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"id": schema.StringAttribute{
												Computed: true,
											},
											"name": schema.StringAttribute{
												Computed: true,
											},
											"type": schema.StringAttribute{
												Description: "The type of a detection entry.\nAvailable values: \"custom\", \"custom_prompt_topic\", \"predefined\", \"integration\", \"exact_data\", \"document_fingerprint\", \"word_list\".",
												Computed:    true,
												Validators: []validator.String{
													stringvalidator.OneOfCaseInsensitive(
														"custom",
														"custom_prompt_topic",
														"predefined",
														"integration",
														"exact_data",
														"document_fingerprint",
														"word_list",
													),
												},
											},
										},
									},
								},
								"entry_count": schema.Int64Attribute{
									Description: "Number of detection entries this validator is attached to.",
									Computed:    true,
									Validators: []validator.Int64{
										int64validator.AtLeast(0),
									},
								},
							},
						},
						"updated_at": schema.StringAttribute{
							Computed:   true,
							CustomType: timetypes.RFC3339Type{},
						},
						"words": schema.SetNestedAttribute{
							Description: "The words that raise confidence when found near a match.",
							Computed:    true,
							CustomType:  customfield.NewNestedObjectSetType[ZeroTrustDLPProximityValidatorsWordsDataSourceModel](ctx),
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"weight": schema.StringAttribute{
										Description: "How much a word contributes to a match's confidence when it is found nearby.\n\nSeparate from a profile's high/medium/low confidence threshold: this describes how telling the\nword is, not how confident the resulting match is.\nAvailable values: \"strong\", \"moderate\", \"weak\".",
										Computed:    true,
										Validators: []validator.String{
											stringvalidator.OneOfCaseInsensitive(
												"strong",
												"moderate",
												"weak",
											),
										},
									},
									"word": schema.StringAttribute{
										Description: "The word or phrase to look for. Must contain 1–128 characters and is matched\ncase-insensitively.",
										Computed:    true,
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func (d *ZeroTrustDLPProximityValidatorsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = ListDataSourceSchema(ctx)
}

func (d *ZeroTrustDLPProximityValidatorsDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
