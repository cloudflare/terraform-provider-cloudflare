// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_list

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ datasource.DataSourceWithConfigValidators = (*ZeroTrustListsDataSource)(nil)

func ListDataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"account_id": schema.StringAttribute{
				Optional:    true,
			},
			"direction": schema.StringAttribute{
				Description: "Sort direction. Applies to the field named in `order_by`; when `order_by`\nis omitted it applies to the default `created_at` ordering. When\n`direction` is omitted the default is field-specific: explicitly choosing\n`created_at` or `updated_at` defaults to descending (newest first); `name`\nand `item_count` default to ascending; and the default `created_at`\nordering used when `order_by` is omitted is ascending (for backwards\ncompatibility).\n  * `asc` — ascending.\n  * `desc` — descending.\nAvailable values: \"asc\", \"desc\".",
				Optional:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("asc", "desc"),
				},
			},
			"order_by": schema.StringAttribute{
				Description: "Field to sort the returned lists by. When omitted, results are ordered by\n`created_at` in ascending order (i.e. creation order) for backwards\ncompatibility. Supported values:\n  * `name` — sort alphabetically by list name.\n  * `created_at` — sort by creation time; defaults to descending unless `direction` is set.\n  * `updated_at` — sort by last-modified time; defaults to descending unless `direction` is set.\n  * `item_count` — sort by number of items in the list.\nAvailable values: \"name\", \"created_at\", \"updated_at\", \"item_count\".",
				Optional:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive(
						"name",
						"created_at",
						"updated_at",
						"item_count",
					),
				},
			},
		"search": schema.StringAttribute{
			Description: "Case-insensitive substring match on the list name or description. When\ncombined with `filter`, both must match (logical AND).",
			Optional:    true,
		},
		"type": schema.StringAttribute{
				Description: "Specify the list type.\nAvailable values: \"SERIAL\", \"URL\", \"DOMAIN\", \"EMAIL\", \"IP\", \"CATEGORY\", \"LOCATION\", \"DEVICE\", \"AAGUID\".",
				Optional:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive(
						"SERIAL",
						"URL",
						"DOMAIN",
						"EMAIL",
						"IP",
						"CATEGORY",
						"LOCATION",
						"DEVICE",
						"AAGUID",
					),
				},
			},
			"filter": schema.ListAttribute{
				Description: "Filter the returned lists by one or more `field:value` pairs.\nRepeat the parameter to apply multiple filters; they are combined with\nlogical AND (a list must satisfy every filter to be returned).\n\nSupported fields and their matching behaviour:\n  * `name` — case-insensitive substring match on the list name.\n  * `id` — substring match on the list ID (UUID), with or without dashes.\n  * `type` — exact match on the list type. Supersedes the legacy `type` query\n    parameter when both are supplied. Must be one of the valid type values.\n  * `item_count` — exact integer match on the number of items in the list.\n\nEach entry must match one of the per-field patterns below: the field must be\none of `name`, `id`, `type`, or `item_count`; `name`/`id` accept any value,\n`type` is restricted to the valid list type values, and `item_count` must be\na non-negative integer.",
				Optional:    true,
				ElementType: jsontypes.NormalizedType{},
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
				CustomType:  customfield.NewNestedObjectListType[ZeroTrustListsResultDataSourceModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Description: "Identify the API resource with a UUID.",
							Computed:    true,
						},
						"list_count": schema.Float64Attribute{
							Description: "Indicate the number of items in the list.",
							Computed:    true,
						},
						"created_at": schema.StringAttribute{
							Computed:   true,
							CustomType: timetypes.RFC3339Type{},
						},
						"description": schema.StringAttribute{
							Description: "Provide the list description.",
							Computed:    true,
						},
						"items": schema.SetNestedAttribute{
							Description: "Provide the list items.",
							Computed:    true,
							CustomType:  customfield.NewNestedObjectSetType[ZeroTrustListsItemsDataSourceModel](ctx),
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"created_at": schema.StringAttribute{
										Computed:   true,
										CustomType: timetypes.RFC3339Type{},
									},
									"description": schema.StringAttribute{
										Description: "Provide the list item description (optional).",
										Computed:    true,
									},
									"value": schema.StringAttribute{
										Description: "Specify the item value.",
										Computed:    true,
									},
								},
							},
						},
						"name": schema.StringAttribute{
							Description: "Specify the list name.",
							Computed:    true,
						},
						"type": schema.StringAttribute{
							Description: "Specify the list type.\nAvailable values: \"SERIAL\", \"URL\", \"DOMAIN\", \"EMAIL\", \"IP\", \"CATEGORY\", \"LOCATION\", \"DEVICE\", \"AAGUID\".",
							Computed:    true,
							Validators: []validator.String{
								stringvalidator.OneOfCaseInsensitive(
									"SERIAL",
									"URL",
									"DOMAIN",
									"EMAIL",
									"IP",
									"CATEGORY",
									"LOCATION",
									"DEVICE",
									"AAGUID",
								),
							},
						},
						"updated_at": schema.StringAttribute{
							Computed:   true,
							CustomType: timetypes.RFC3339Type{},
						},
					},
				},
			},
		},
	}
}

func (d *ZeroTrustListsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = ListDataSourceSchema(ctx)
}

func (d *ZeroTrustListsDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
