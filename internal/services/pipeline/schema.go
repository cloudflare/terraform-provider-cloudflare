// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package pipeline

import (
  "context"

  "github.com/hashicorp/terraform-plugin-framework-validators/float64validator"
  "github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
  "github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
  "github.com/hashicorp/terraform-plugin-framework/resource"
  "github.com/hashicorp/terraform-plugin-framework/resource/schema"
  "github.com/hashicorp/terraform-plugin-framework/resource/schema/float64default"
  "github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
  "github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
  "github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
  "github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
  "github.com/hashicorp/terraform-plugin-framework/schema/validator"
  "github.com/hashicorp/terraform-plugin-framework/types"
)

var _ resource.ResourceWithConfigValidators = (*PipelineResource)(nil)

func ResourceSchema(ctx context.Context) (schema.Schema) {
  return schema.Schema{
    Attributes: map[string]schema.Attribute{
      "id": schema.StringAttribute{
        Description: "Specifies the Pipeline identifier.",
        Computed: true,
        PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
      },
      "account_id": schema.StringAttribute{
        Description: "Specifies the public ID of the account.",
        Required: true,
        PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
      },
      "name": schema.StringAttribute{
        Description: "Defines the name of Pipeline.",
        Required: true,
      },
      "destination": schema.SingleNestedAttribute{
        Required: true,
        Attributes: map[string]schema.Attribute{
          "batch": schema.SingleNestedAttribute{
            Required: true,
            Attributes: map[string]schema.Attribute{
              "max_bytes": schema.Int64Attribute{
                Description: "Specifies rough maximum size of files.",
                Computed: true,
                Optional: true,
                Validators: []validator.Int64{
                int64validator.Between(1000, 100000000),
                },
                Default: int64default.  StaticInt64(100000000),
              },
              "max_duration_s": schema.Float64Attribute{
                Description: "Specifies duration to wait to aggregate batches files.",
                Computed: true,
                Optional: true,
                Validators: []validator.Float64{
                float64validator.Between(0.25, 300),
                },
                Default: float64default.  StaticFloat64(300),
              },
              "max_rows": schema.Int64Attribute{
                Description: "Specifies rough maximum number of rows per file.",
                Computed: true,
                Optional: true,
                Validators: []validator.Int64{
                int64validator.Between(100, 10000000),
                },
                Default: int64default.  StaticInt64(10000000),
              },
            },
          },
          "compression": schema.SingleNestedAttribute{
            Required: true,
            Attributes: map[string]schema.Attribute{
              "type": schema.StringAttribute{
                Description: "Specifies the desired compression algorithm and format.\nAvailable values: \"none\", \"gzip\", \"deflate\".",
                Computed: true,
                Optional: true,
                Validators: []validator.String{
                stringvalidator.OneOfCaseInsensitive(
                  "none",
                  "gzip",
                  "deflate",
                ),
                },
                Default: stringdefault.  StaticString("gzip"),
              },
            },
          },
          "credentials": schema.SingleNestedAttribute{
            Required: true,
            Attributes: map[string]schema.Attribute{
              "access_key_id": schema.StringAttribute{
                Description: "Specifies the R2 Bucket Access Key Id.",
                Required: true,
              },
              "endpoint": schema.StringAttribute{
                Description: "Specifies the R2 Endpoint.",
                Required: true,
              },
              "secret_access_key": schema.StringAttribute{
                Description: "Specifies the R2 Bucket Secret Access Key.",
                Required: true,
              },
            },
          },
          "format": schema.StringAttribute{
            Description: "Specifies the format of data to deliver.\nAvailable values: \"json\".",
            Required: true,
            Validators: []validator.String{
            stringvalidator.OneOfCaseInsensitive("json"),
            },
          },
          "path": schema.SingleNestedAttribute{
            Required: true,
            Attributes: map[string]schema.Attribute{
              "bucket": schema.StringAttribute{
                Description: "Specifies the R2 Bucket to store files.",
                Required: true,
              },
              "filename": schema.StringAttribute{
                Description: "Specifies the name pattern to for individual data files.",
                Optional: true,
              },
              "filepath": schema.StringAttribute{
                Description: "Specifies the name pattern for directory.",
                Optional: true,
              },
              "prefix": schema.StringAttribute{
                Description: "Specifies the base directory within the bucket.",
                Optional: true,
              },
            },
          },
          "type": schema.StringAttribute{
            Description: "Specifies the type of destination.\nAvailable values: \"r2\".",
            Required: true,
            Validators: []validator.String{
            stringvalidator.OneOfCaseInsensitive("r2"),
            },
          },
        },
      },
      "source": schema.ListNestedAttribute{
        Required: true,
        NestedObject: schema.NestedAttributeObject{
          Attributes: map[string]schema.Attribute{
            "format": schema.StringAttribute{
              Description: "Specifies the format of source data.\nAvailable values: \"json\".",
              Required: true,
              Validators: []validator.String{
              stringvalidator.OneOfCaseInsensitive("json"),
              },
            },
            "type": schema.StringAttribute{
              Required: true,
            },
            "authentication": schema.BoolAttribute{
              Description: "Specifies authentication is required to send to this Pipeline.",
              Optional: true,
            },
            "cors": schema.SingleNestedAttribute{
              Optional: true,
              Attributes: map[string]schema.Attribute{
                "origins": schema.ListAttribute{
                  Description: "Specifies allowed origins to allow Cross Origin HTTP Requests.",
                  Optional: true,
                  ElementType: types.StringType,
                },
              },
            },
          },
        },
      },
      "endpoint": schema.StringAttribute{
        Description: "Indicates the endpoint URL to send traffic.",
        Computed: true,
      },
      "version": schema.Float64Attribute{
        Description: "Indicates the version number of last saved configuration.",
        Computed: true,
      },
    },
  }
}

func (r *PipelineResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
  resp.Schema = ResourceSchema(ctx)
}

func (r *PipelineResource) ConfigValidators(_ context.Context) ([]resource.ConfigValidator) {
  return []resource.ConfigValidator{
  }
}
