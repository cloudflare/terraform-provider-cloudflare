// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package d1_database

import (
  "context"

  "github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
  "github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
  "github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
  "github.com/hashicorp/terraform-plugin-framework/resource"
  "github.com/hashicorp/terraform-plugin-framework/resource/schema"
  "github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
  "github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
  "github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ resource.ResourceWithConfigValidators = (*D1DatabaseResource)(nil)

func ResourceSchema(ctx context.Context) (schema.Schema) {
  return schema.Schema{
    Attributes: map[string]schema.Attribute{
      "id": schema.StringAttribute{
        Description: "D1 database identifier (UUID).",
        Computed: true,
        PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
      },
      "uuid": schema.StringAttribute{
        Description: "D1 database identifier (UUID).",
        Computed: true,
        PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
      },
      "account_id": schema.StringAttribute{
        Description: "Account identifier tag.",
        Required: true,
        PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
      },
      "name": schema.StringAttribute{
        Description: "D1 database name.",
        Required: true,
        PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
      },
      "primary_location_hint": schema.StringAttribute{
        Description: "Specify the region to create the D1 primary, if available. If this option is omitted, the D1 will be created as close as possible to the current user.\nAvailable values: \"wnam\", \"enam\", \"weur\", \"eeur\", \"apac\", \"oc\".",
        Optional: true,
        Validators: []validator.String{
        stringvalidator.OneOfCaseInsensitive(
          "wnam",
          "enam",
          "weur",
          "eeur",
          "apac",
          "oc",
        ),
        },
        PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
      },
      "read_replication": schema.SingleNestedAttribute{
        Description: "Configuration for D1 read replication.",
        Computed: true,
        Optional: true,
        CustomType: customfield.NewNestedObjectType[D1DatabaseReadReplicationModel](ctx),
        Attributes: map[string]schema.Attribute{
          "mode": schema.StringAttribute{
            Description: "The read replication mode for the database. Use 'auto' to create replicas and allow D1 automatically place them around the world, or 'disabled' to not use any database replicas (it can take a few hours for all replicas to be deleted).\nAvailable values: \"auto\", \"disabled\".",
            Required: true,
            Validators: []validator.String{
            stringvalidator.OneOfCaseInsensitive("auto", "disabled"),
            },
          },
        },
      },
      "created_at": schema.StringAttribute{
        Description: "Specifies the timestamp the resource was created as an ISO8601 string.",
        Computed: true,
        CustomType: timetypes.RFC3339Type{

        },
      },
      "file_size": schema.Float64Attribute{
        Description: "The D1 database's size, in bytes.",
        Computed: true,
      },
      "num_tables": schema.Float64Attribute{
        Computed: true,
      },
      "version": schema.StringAttribute{
        Computed: true,
      },
    },
  }
}

func (r *D1DatabaseResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
  resp.Schema = ResourceSchema(ctx)
}

func (r *D1DatabaseResource) ConfigValidators(_ context.Context) ([]resource.ConfigValidator) {
  return []resource.ConfigValidator{
  }
}
