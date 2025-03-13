// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package account_api_token_permission_groups

import (
  "context"

  "github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
  "github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
  "github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
  "github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
  "github.com/hashicorp/terraform-plugin-framework/datasource"
  "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
  "github.com/hashicorp/terraform-plugin-framework/schema/validator"
  "github.com/hashicorp/terraform-plugin-framework/types"
)

var _ datasource.DataSourceWithConfigValidators = (*AccountAPITokenPermissionGroupsListDataSource)(nil)

func ListDataSourceSchema(ctx context.Context) (schema.Schema) {
  return schema.Schema{
    Attributes: map[string]schema.Attribute{
      "account_id": schema.StringAttribute{
        Description: "Account identifier tag.",
        Required: true,
      },
      "max_items": schema.Int64Attribute{
        Description: "Max items to fetch, default: 1000",
        Optional: true,
        Validators: []validator.Int64{
        int64validator.AtLeast(0),
        },
      },
      "result": schema.ListNestedAttribute{
        Description: "The items returned by the data source",
        Computed: true,
        CustomType: customfield.NewNestedObjectListType[AccountAPITokenPermissionGroupsListResultDataSourceModel](ctx),
        NestedObject: schema.NestedAttributeObject{
          Attributes: map[string]schema.Attribute{
            "id": schema.StringAttribute{
              Description: "Public ID.",
              Computed: true,
            },
            "name": schema.StringAttribute{
              Description: "Permission Group Name",
              Computed: true,
            },
            "scopes": schema.ListAttribute{
              Description: "Resources to which the Permission Group is scoped",
              Computed: true,
              Validators: []validator.List{
              listvalidator.ValueStringsAre(
                stringvalidator.OneOfCaseInsensitive(
                  "com.cloudflare.api.account",
                  "com.cloudflare.api.account.zone",
                  "com.cloudflare.api.user",
                  "com.cloudflare.edge.r2.bucket",
                ),
              ),
              },
              CustomType: customfield.NewListType[types.String](ctx),
              ElementType: types.StringType,
            },
          },
        },
      },
    },
  }
}

func (d *AccountAPITokenPermissionGroupsListDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
  resp.Schema = ListDataSourceSchema(ctx)
}

func (d *AccountAPITokenPermissionGroupsListDataSource) ConfigValidators(_ context.Context) ([]datasource.ConfigValidator) {
  return []datasource.ConfigValidator{
  }
}
