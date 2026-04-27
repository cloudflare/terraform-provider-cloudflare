// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package email_security_impersonation_registry

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

var _ datasource.DataSourceWithConfigValidators = (*EmailSecurityImpersonationRegistriesDataSource)(nil)

func ListDataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		MarkdownDescription: schemata.Description{
			Scopes: []string{
				"Cloud Email Security: Read",
				"Cloud Email Security: Write",
			},
		}.String(),
		Attributes: map[string]schema.Attribute{
			"account_id": schema.StringAttribute{
				Description: "Identifier.",
				Required:    true,
			},
			"direction": schema.StringAttribute{
				Description: "The sorting direction.\nAvailable values: \"asc\", \"desc\".",
				Optional:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("asc", "desc"),
				},
			},
			"order": schema.StringAttribute{
				Description: "Field to sort by.\nAvailable values: \"name\", \"email\", \"created_at\".",
				Optional:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive(
						"name",
						"email",
						"created_at",
					),
				},
			},
			"provenance": schema.StringAttribute{
				Description: `Available values: "A1S_INTERNAL", "SNOOPY-CASB_OFFICE_365", "SNOOPY-OFFICE_365", "SNOOPY-GOOGLE_DIRECTORY".`,
				Optional:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive(
						"A1S_INTERNAL",
						"SNOOPY-CASB_OFFICE_365",
						"SNOOPY-OFFICE_365",
						"SNOOPY-GOOGLE_DIRECTORY",
					),
				},
			},
			"search": schema.StringAttribute{
				Description: "Search term for filtering records. Behavior may change.",
				Optional:    true,
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
				CustomType:  customfield.NewNestedObjectListType[EmailSecurityImpersonationRegistriesResultDataSourceModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Description: "Impersonation registry entry identifier",
							Computed:    true,
						},
						"comments": schema.StringAttribute{
							Computed: true,
						},
						"created_at": schema.StringAttribute{
							Computed:   true,
							CustomType: timetypes.RFC3339Type{},
						},
						"directory_id": schema.Int64Attribute{
							Computed: true,
						},
						"directory_node_id": schema.Int64Attribute{
							Computed: true,
						},
						"email": schema.StringAttribute{
							Computed: true,
						},
						"external_directory_node_id": schema.StringAttribute{
							Computed:           true,
							DeprecationMessage: "This attribute is deprecated.",
						},
						"is_email_regex": schema.BoolAttribute{
							Computed: true,
						},
						"last_modified": schema.StringAttribute{
							Description:        "Deprecated, use `modified_at` instead. End of life: November 1, 2026.",
							Computed:           true,
							DeprecationMessage: "This attribute is deprecated.",
							CustomType:         timetypes.RFC3339Type{},
						},
						"modified_at": schema.StringAttribute{
							Computed:   true,
							CustomType: timetypes.RFC3339Type{},
						},
						"name": schema.StringAttribute{
							Computed: true,
						},
						"provenance": schema.StringAttribute{
							Description: `Available values: "A1S_INTERNAL", "SNOOPY-CASB_OFFICE_365", "SNOOPY-OFFICE_365", "SNOOPY-GOOGLE_DIRECTORY".`,
							Computed:    true,
							Validators: []validator.String{
								stringvalidator.OneOfCaseInsensitive(
									"A1S_INTERNAL",
									"SNOOPY-CASB_OFFICE_365",
									"SNOOPY-OFFICE_365",
									"SNOOPY-GOOGLE_DIRECTORY",
								),
							},
						},
					},
				},
			},
		},
	}
}

func (d *EmailSecurityImpersonationRegistriesDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = ListDataSourceSchema(ctx)
}

func (d *EmailSecurityImpersonationRegistriesDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
