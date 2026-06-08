// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package share_recipient

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ datasource.DataSourceWithConfigValidators = (*ShareRecipientDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "Share Recipient identifier tag.",
				Computed:    true,
			},
			"recipient_id": schema.StringAttribute{
				Description: "Share Recipient identifier tag.",
				Required:    true,
			},
			"account_id": schema.StringAttribute{
				Description: "Account identifier.",
				Required:    true,
			},
			"share_id": schema.StringAttribute{
				Description: "Share identifier tag.",
				Required:    true,
			},
			"include_resources": schema.BoolAttribute{
				Description: "Include resources in the response.",
				Optional:    true,
			},
			"association_status": schema.StringAttribute{
				Description: "Share Recipient association status.\nAvailable values: \"associating\", \"associated\", \"disassociating\", \"disassociated\".",
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive(
						"associating",
						"associated",
						"disassociating",
						"disassociated",
					),
				},
			},
			"created": schema.StringAttribute{
				Description: "When the share was created.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"modified": schema.StringAttribute{
				Description: "When the share was modified.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"resources": schema.ListNestedAttribute{
				Computed:   true,
				CustomType: customfield.NewNestedObjectListType[ShareRecipientResourcesDataSourceModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"error": schema.StringAttribute{
							Description: "Share Recipient error message.",
							Computed:    true,
						},
						"resource_id": schema.StringAttribute{
							Description: "Share Resource identifier.",
							Computed:    true,
						},
						"resource_version": schema.Int64Attribute{
							Description: "Resource Version.",
							Computed:    true,
						},
						"terminal": schema.BoolAttribute{
							Description: "Whether the error is terminal or will be continually retried.",
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func (d *ShareRecipientDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *ShareRecipientDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
