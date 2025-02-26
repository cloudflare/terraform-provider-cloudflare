// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package dns_settings_internal_view

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ datasource.DataSourceWithConfigValidators = (*DNSSettingsInternalViewsDataSource)(nil)

func ListDataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"account_id": schema.StringAttribute{
				Description: "Identifier",
				Required:    true,
			},
			"order": schema.StringAttribute{
				Description: "Field to order DNS views by.\navailable values: \"name\", \"created_on\", \"modified_on\"",
				Optional:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive(
						"name",
						"created_on",
						"modified_on",
					),
				},
			},
			"zone_id": schema.StringAttribute{
				Description: "A zone ID that exists in the zones list for the view.",
				Optional:    true,
			},
			"zone_name": schema.StringAttribute{
				Description: "A zone name that exists in the zones list for the view.",
				Optional:    true,
			},
			"name": schema.SingleNestedAttribute{
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"contains": schema.StringAttribute{
						Description: "Substring of the DNS view name.",
						Optional:    true,
					},
					"endswith": schema.StringAttribute{
						Description: "Suffix of the DNS view name.",
						Optional:    true,
					},
					"exact": schema.StringAttribute{
						Description: "Exact value of the DNS view name.",
						Optional:    true,
					},
					"startswith": schema.StringAttribute{
						Description: "Prefix of the DNS view name.",
						Optional:    true,
					},
				},
			},
			"direction": schema.StringAttribute{
				Description: "Direction to order DNS views in.\navailable values: \"asc\", \"desc\"",
				Computed:    true,
				Optional:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("asc", "desc"),
				},
			},
			"match": schema.StringAttribute{
				Description: "Whether to match all search requirements or at least one (any). If set to `all`, acts like a logical AND between filters. If set to `any`, acts like a logical OR instead.\navailable values: \"any\", \"all\"",
				Computed:    true,
				Optional:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("any", "all"),
				},
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
				CustomType:  customfield.NewNestedObjectListType[DNSSettingsInternalViewsResultDataSourceModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Description: "Identifier",
							Computed:    true,
						},
						"created_time": schema.StringAttribute{
							Description: "When the view was created.",
							Computed:    true,
							CustomType:  timetypes.RFC3339Type{},
						},
						"modified_time": schema.StringAttribute{
							Description: "When the view was last modified.",
							Computed:    true,
							CustomType:  timetypes.RFC3339Type{},
						},
						"name": schema.StringAttribute{
							Description: "The name of the view.",
							Computed:    true,
						},
						"zones": schema.ListAttribute{
							Description: "The list of zones linked to this view.",
							Computed:    true,
							CustomType:  customfield.NewListType[types.String](ctx),
							ElementType: types.StringType,
						},
					},
				},
			},
		},
	}
}

func (d *DNSSettingsInternalViewsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = ListDataSourceSchema(ctx)
}

func (d *DNSSettingsInternalViewsDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
