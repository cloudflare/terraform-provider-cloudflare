// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package api_shield_schema

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
  "github.com/hashicorp/terraform-plugin-framework/types"
)

var _ resource.ResourceWithConfigValidators = (*APIShieldSchemaResource)(nil)

func ResourceSchema(ctx context.Context) (schema.Schema) {
  return schema.Schema{
    Attributes: map[string]schema.Attribute{
      "zone_id": schema.StringAttribute{
        Description: "Identifier.",
        Required: true,
        PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
      },
      "schema_id": schema.StringAttribute{
        Optional: true,
        PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
      },
      "file": schema.StringAttribute{
        Description: "Schema file bytes",
        Required: true,
        PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
      },
      "kind": schema.StringAttribute{
        Description: "Kind of schema\nAvailable values: \"openapi_v3\".",
        Required: true,
        Validators: []validator.String{
        stringvalidator.OneOfCaseInsensitive("openapi_v3"),
        },
        PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
      },
      "name": schema.StringAttribute{
        Description: "Name of the schema",
        Optional: true,
        PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
      },
      "validation_enabled": schema.StringAttribute{
        Description: "Flag whether schema is enabled for validation.\nAvailable values: \"true\", \"false\".",
        Optional: true,
        Validators: []validator.String{
        stringvalidator.OneOfCaseInsensitive("true", "false"),
        },
      },
      "created_at": schema.StringAttribute{
        Computed: true,
        CustomType: timetypes.RFC3339Type{

        },
      },
      "source": schema.StringAttribute{
        Description: "Source of the schema",
        Computed: true,
      },
      "schema": schema.SingleNestedAttribute{
        Computed: true,
        CustomType: customfield.NewNestedObjectType[APIShieldSchemaSchemaModel](ctx),
        Attributes: map[string]schema.Attribute{
          "created_at": schema.StringAttribute{
            Computed: true,
            CustomType: timetypes.RFC3339Type{

            },
          },
          "kind": schema.StringAttribute{
            Description: "Kind of schema\nAvailable values: \"openapi_v3\".",
            Computed: true,
            Validators: []validator.String{
            stringvalidator.OneOfCaseInsensitive("openapi_v3"),
            },
          },
          "name": schema.StringAttribute{
            Description: "Name of the schema",
            Computed: true,
          },
          "schema_id": schema.StringAttribute{
            Description: "UUID.",
            Computed: true,
          },
          "source": schema.StringAttribute{
            Description: "Source of the schema",
            Computed: true,
          },
          "validation_enabled": schema.BoolAttribute{
            Description: "Flag whether schema is enabled for validation.",
            Computed: true,
          },
        },
      },
      "upload_details": schema.SingleNestedAttribute{
        Computed: true,
        CustomType: customfield.NewNestedObjectType[APIShieldSchemaUploadDetailsModel](ctx),
        Attributes: map[string]schema.Attribute{
          "warnings": schema.ListNestedAttribute{
            Description: "Diagnostic warning events that occurred during processing. These events are non-critical errors found within the schema.",
            Computed: true,
            CustomType: customfield.NewNestedObjectListType[APIShieldSchemaUploadDetailsWarningsModel](ctx),
            NestedObject: schema.NestedAttributeObject{
              Attributes: map[string]schema.Attribute{
                "code": schema.Int64Attribute{
                  Description: "Code that identifies the event that occurred.",
                  Computed: true,
                },
                "locations": schema.ListAttribute{
                  Description: "JSONPath location(s) in the schema where these events were encountered.  See [https://goessner.net/articles/JsonPath/](https://goessner.net/articles/JsonPath/) for JSONPath specification.",
                  Computed: true,
                  CustomType: customfield.NewListType[types.String](ctx),
                  ElementType: types.StringType,
                },
                "message": schema.StringAttribute{
                  Description: "Diagnostic message that describes the event.",
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

func (r *APIShieldSchemaResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
  resp.Schema = ResourceSchema(ctx)
}

func (r *APIShieldSchemaResource) ConfigValidators(_ context.Context) ([]resource.ConfigValidator) {
  return []resource.ConfigValidator{
  }
}
