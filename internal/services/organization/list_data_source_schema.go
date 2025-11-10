// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package organization

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ datasource.DataSourceWithConfigValidators = (*OrganizationsDataSource)(nil)

func ListDataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"page_size": schema.Int64Attribute{
				Description: "The amount of items to return. Defaults to 10.",
				Optional:    true,
				Validators: []validator.Int64{
					int64validator.Between(0, 1000),
				},
			},
			"page_token": schema.StringAttribute{
				Description: "An opaque token returned from the last list response that when\nprovided will retrieve the next page.\n\nParameters used to filter the retrieved list must remain in subsequent\nrequests with a page token.",
				Optional:    true,
			},
			"id": schema.ListAttribute{
				Description: "Only return organizations with the specified IDs (ex. id=foo&id=bar). Send multiple elements\nby repeating the query value.",
				Optional:    true,
				ElementType: types.StringType,
			},
			"containing": schema.SingleNestedAttribute{
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"account": schema.StringAttribute{
						Description: "Filter the list of organizations to the ones that contain this particular\naccount.",
						Optional:    true,
					},
					"organization": schema.StringAttribute{
						Description: "Filter the list of organizations to the ones that contain this particular\norganization.",
						Optional:    true,
					},
					"user": schema.StringAttribute{
						Description: "Filter the list of organizations to the ones that contain this particular\nuser.\n\nIMPORTANT: Just because an organization \"contains\" a user is not a\nrepresentation of any authorization or privilege to manage any resources\ntherein. An organization \"containing\" a user simply means the user is managed by\nthat organization.",
						Optional:    true,
					},
				},
			},
			"name": schema.SingleNestedAttribute{
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"contains": schema.StringAttribute{
						Description: "(case-insensitive) Filter the list of organizations to where the name contains a particular\nstring.",
						Optional:    true,
					},
					"ends_with": schema.StringAttribute{
						Description: "(case-insensitive) Filter the list of organizations to where the name ends with a particular\nstring.",
						Optional:    true,
					},
					"starts_with": schema.StringAttribute{
						Description: "(case-insensitive) Filter the list of organizations to where the name starts with a\nparticular string.",
						Optional:    true,
					},
				},
			},
			"parent": schema.SingleNestedAttribute{
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"id": schema.StringAttribute{
						Description: "Filter the list of organizations to the ones that are a sub-organization\nof the specified organization.\n\n\"null\" is a valid value to provide for this parameter. It means \"where\nan organization has no parent (i.e. it is a 'root' organization).\"",
						Optional:    true,
					},
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
				CustomType:  customfield.NewNestedObjectListType[OrganizationsResultDataSourceModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Computed: true,
						},
						"create_time": schema.StringAttribute{
							Computed:   true,
							CustomType: timetypes.RFC3339Type{},
						},
						"meta": schema.SingleNestedAttribute{
							Computed:   true,
							CustomType: customfield.NewNestedObjectType[OrganizationsMetaDataSourceModel](ctx),
							Attributes: map[string]schema.Attribute{
								"flags": schema.SingleNestedAttribute{
									Description: "Enable features for Organizations.",
									Computed:    true,
									CustomType:  customfield.NewNestedObjectType[OrganizationsMetaFlagsDataSourceModel](ctx),
									Attributes: map[string]schema.Attribute{
										"account_creation": schema.StringAttribute{
											Computed: true,
										},
										"account_deletion": schema.StringAttribute{
											Computed: true,
										},
										"account_migration": schema.StringAttribute{
											Computed: true,
										},
										"account_mobility": schema.StringAttribute{
											Computed: true,
										},
										"sub_org_creation": schema.StringAttribute{
											Computed: true,
										},
									},
								},
								"managed_by": schema.StringAttribute{
									Computed: true,
								},
							},
						},
						"name": schema.StringAttribute{
							Computed: true,
						},
						"parent": schema.SingleNestedAttribute{
							Computed:   true,
							CustomType: customfield.NewNestedObjectType[OrganizationsParentDataSourceModel](ctx),
							Attributes: map[string]schema.Attribute{
								"id": schema.StringAttribute{
									Computed: true,
								},
								"name": schema.StringAttribute{
									Computed: true,
								},
							},
						},
						"profile": schema.SingleNestedAttribute{
							Computed:   true,
							CustomType: customfield.NewNestedObjectType[OrganizationsProfileDataSourceModel](ctx),
							Attributes: map[string]schema.Attribute{
								"business_address": schema.StringAttribute{
									Computed: true,
								},
								"business_email": schema.StringAttribute{
									Computed: true,
								},
								"business_name": schema.StringAttribute{
									Computed: true,
								},
								"business_phone": schema.StringAttribute{
									Computed: true,
								},
								"external_metadata": schema.StringAttribute{
									Computed: true,
								},
							},
						},
					},
				},
			},
		},
	}
}

func (d *OrganizationsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = ListDataSourceSchema(ctx)
}

func (d *OrganizationsDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
