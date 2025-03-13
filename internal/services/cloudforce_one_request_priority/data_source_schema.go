// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloudforce_one_request_priority

import (
  "context"

  "github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
  "github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
  "github.com/hashicorp/terraform-plugin-framework/datasource"
  "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
  "github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ datasource.DataSourceWithConfigValidators = (*CloudforceOneRequestPriorityDataSource)(nil)

func DataSourceSchema(ctx context.Context) (schema.Schema) {
  return schema.Schema{
    Attributes: map[string]schema.Attribute{
      "account_identifier": schema.StringAttribute{
        Description: "Identifier",
        Required: true,
      },
      "priority_identifer": schema.StringAttribute{
        Description: "UUID",
        Required: true,
      },
      "completed": schema.StringAttribute{
        Computed: true,
        CustomType: timetypes.RFC3339Type{

        },
      },
      "content": schema.StringAttribute{
        Description: "Request content",
        Computed: true,
      },
      "created": schema.StringAttribute{
        Computed: true,
        CustomType: timetypes.RFC3339Type{

        },
      },
      "id": schema.StringAttribute{
        Description: "UUID",
        Computed: true,
      },
      "message_tokens": schema.Int64Attribute{
        Description: "Tokens for the request messages",
        Computed: true,
      },
      "priority": schema.StringAttribute{
        Computed: true,
        CustomType: timetypes.RFC3339Type{

        },
      },
      "readable_id": schema.StringAttribute{
        Description: "Readable Request ID",
        Computed: true,
      },
      "request": schema.StringAttribute{
        Description: "Requested information from request",
        Computed: true,
      },
      "status": schema.StringAttribute{
        Description: "Request Status\nAvailable values: \"open\", \"accepted\", \"reported\", \"approved\", \"completed\", \"declined\".",
        Computed: true,
        Validators: []validator.String{
        stringvalidator.OneOfCaseInsensitive(
          "open",
          "accepted",
          "reported",
          "approved",
          "completed",
          "declined",
        ),
        },
      },
      "summary": schema.StringAttribute{
        Description: "Brief description of the request",
        Computed: true,
      },
      "tlp": schema.StringAttribute{
        Description: "The CISA defined Traffic Light Protocol (TLP)\nAvailable values: \"clear\", \"amber\", \"amber-strict\", \"green\", \"red\".",
        Computed: true,
        Validators: []validator.String{
        stringvalidator.OneOfCaseInsensitive(
          "clear",
          "amber",
          "amber-strict",
          "green",
          "red",
        ),
        },
      },
      "tokens": schema.Int64Attribute{
        Description: "Tokens for the request",
        Computed: true,
      },
      "updated": schema.StringAttribute{
        Computed: true,
        CustomType: timetypes.RFC3339Type{

        },
      },
    },
  }
}

func (d *CloudforceOneRequestPriorityDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
  resp.Schema = DataSourceSchema(ctx)
}

func (d *CloudforceOneRequestPriorityDataSource) ConfigValidators(_ context.Context) ([]datasource.ConfigValidator) {
  return []datasource.ConfigValidator{
  }
}
