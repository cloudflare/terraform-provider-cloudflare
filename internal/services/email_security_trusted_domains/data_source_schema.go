// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package email_security_trusted_domains

import (
  "context"

  "github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
  "github.com/hashicorp/terraform-plugin-framework-validators/datasourcevalidator"
  "github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
  "github.com/hashicorp/terraform-plugin-framework/datasource"
  "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
  "github.com/hashicorp/terraform-plugin-framework/path"
  "github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ datasource.DataSourceWithConfigValidators = (*EmailSecurityTrustedDomainsDataSource)(nil)

func DataSourceSchema(ctx context.Context) (schema.Schema) {
  return schema.Schema{
    Attributes: map[string]schema.Attribute{
      "id": schema.Int64Attribute{
        Description: "The unique identifier for the trusted domain.",
        Computed: true,
      },
      "trusted_domain_id": schema.Int64Attribute{
        Description: "The unique identifier for the trusted domain.",
        Optional: true,
      },
      "account_id": schema.StringAttribute{
        Description: "Account Identifier",
        Required: true,
      },
      "comments": schema.StringAttribute{
        Computed: true,
      },
      "created_at": schema.StringAttribute{
        Computed: true,
        CustomType: timetypes.RFC3339Type{

        },
      },
      "is_recent": schema.BoolAttribute{
        Description: "Select to prevent recently registered domains from triggering a\nSuspicious or Malicious disposition.",
        Computed: true,
      },
      "is_regex": schema.BoolAttribute{
        Computed: true,
      },
      "is_similarity": schema.BoolAttribute{
        Description: "Select for partner or other approved domains that have similar\nspelling to your connected domains. Prevents listed domains from\ntriggering a Spoof disposition.",
        Computed: true,
      },
      "last_modified": schema.StringAttribute{
        Computed: true,
        CustomType: timetypes.RFC3339Type{

        },
      },
      "pattern": schema.StringAttribute{
        Computed: true,
      },
      "filter": schema.SingleNestedAttribute{
        Optional: true,
        Attributes: map[string]schema.Attribute{
          "direction": schema.StringAttribute{
            Description: "The sorting direction.\nAvailable values: \"asc\", \"desc\".",
            Optional: true,
            Validators: []validator.String{
            stringvalidator.OneOfCaseInsensitive("asc", "desc"),
            },
          },
          "is_recent": schema.BoolAttribute{
            Optional: true,
          },
          "is_similarity": schema.BoolAttribute{
            Optional: true,
          },
          "order": schema.StringAttribute{
            Description: "The field to sort by.\nAvailable values: \"pattern\", \"created_at\".",
            Optional: true,
            Validators: []validator.String{
            stringvalidator.OneOfCaseInsensitive("pattern", "created_at"),
            },
          },
          "search": schema.StringAttribute{
            Description: "Allows searching in multiple properties of a record simultaneously.\nThis parameter is intended for human users, not automation. Its exact\nbehavior is intentionally left unspecified and is subject to change\nin the future.",
            Optional: true,
          },
        },
      },
    },
  }
}

func (d *EmailSecurityTrustedDomainsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
  resp.Schema = DataSourceSchema(ctx)
}

func (d *EmailSecurityTrustedDomainsDataSource) ConfigValidators(_ context.Context) ([]datasource.ConfigValidator) {
  return []datasource.ConfigValidator{
  datasourcevalidator.ExactlyOneOf(path.MatchRoot("trusted_domain_id"), path.MatchRoot("filter")),
  }
}
