// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package pipeline

import (
  "context"

  "github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
  "github.com/hashicorp/terraform-plugin-framework-validators/float64validator"
  "github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
  "github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
  "github.com/hashicorp/terraform-plugin-framework/datasource"
  "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
  "github.com/hashicorp/terraform-plugin-framework/schema/validator"
  "github.com/hashicorp/terraform-plugin-framework/types"
)

var _ datasource.DataSourceWithConfigValidators = (*PipelineDataSource)(nil)

func DataSourceSchema(ctx context.Context) (schema.Schema) {
  return schema.Schema{
    Attributes: map[string]schema.Attribute{
      "account_id": schema.StringAttribute{
        Description: "Specifies the public ID of the account.",
        Required: true,
      },
      "pipeline_name": schema.StringAttribute{
        Description: "Defines the name of Pipeline.",
        Required: true,
      },
      "endpoint": schema.StringAttribute{
        Description: "Indicates the endpoint URL to send traffic.",
        Computed: true,
      },
      "id": schema.StringAttribute{
        Description: "Specifies the Pipeline identifier.",
        Computed: true,
      },
      "name": schema.StringAttribute{
        Description: "Defines the name of Pipeline.",
        Computed: true,
      },
      "version": schema.Float64Attribute{
        Description: "Indicates the version number of last saved configuration.",
        Computed: true,
      },
      "destination": schema.SingleNestedAttribute{
        Computed: true,
        CustomType: customfield.NewNestedObjectType[PipelineDestinationDataSourceModel](ctx),
        Attributes: map[string]schema.Attribute{
          "batch": schema.SingleNestedAttribute{
            Computed: true,
            CustomType: customfield.NewNestedObjectType[PipelineDestinationBatchDataSourceModel](ctx),
            Attributes: map[string]schema.Attribute{
              "max_bytes": schema.Int64Attribute{
                Description: "Specifies rough maximum size of files.",
                Computed: true,
                Validators: []validator.Int64{
                int64validator.Between(1000, 100000000),
                },
              },
              "max_duration_s": schema.Float64Attribute{
                Description: "Specifies duration to wait to aggregate batches files.",
                Computed: true,
                Validators: []validator.Float64{
                float64validator.Between(0.25, 300),
                },
              },
              "max_rows": schema.Int64Attribute{
                Description: "Specifies rough maximum number of rows per file.",
                Computed: true,
                Validators: []validator.Int64{
                int64validator.Between(100, 10000000),
                },
              },
            },
          },
          "compression": schema.SingleNestedAttribute{
            Computed: true,
            CustomType: customfield.NewNestedObjectType[PipelineDestinationCompressionDataSourceModel](ctx),
            Attributes: map[string]schema.Attribute{
              "type": schema.StringAttribute{
                Description: "Specifies the desired compression algorithm and format.\nAvailable values: \"none\", \"gzip\", \"deflate\".",
                Computed: true,
                Validators: []validator.String{
                stringvalidator.OneOfCaseInsensitive(
                  "none",
                  "gzip",
                  "deflate",
                ),
                },
              },
            },
          },
          "format": schema.StringAttribute{
            Description: "Specifies the format of data to deliver.\nAvailable values: \"json\".",
            Computed: true,
            Validators: []validator.String{
            stringvalidator.OneOfCaseInsensitive("json"),
            },
          },
          "path": schema.SingleNestedAttribute{
            Computed: true,
            CustomType: customfield.NewNestedObjectType[PipelineDestinationPathDataSourceModel](ctx),
            Attributes: map[string]schema.Attribute{
              "bucket": schema.StringAttribute{
                Description: "Specifies the R2 Bucket to store files.",
                Computed: true,
              },
              "filename": schema.StringAttribute{
                Description: "Specifies the name pattern to for individual data files.",
                Computed: true,
              },
              "filepath": schema.StringAttribute{
                Description: "Specifies the name pattern for directory.",
                Computed: true,
              },
              "prefix": schema.StringAttribute{
                Description: "Specifies the base directory within the bucket.",
                Computed: true,
              },
            },
          },
          "type": schema.StringAttribute{
            Description: "Specifies the type of destination.\nAvailable values: \"r2\".",
            Computed: true,
            Validators: []validator.String{
            stringvalidator.OneOfCaseInsensitive("r2"),
            },
          },
        },
      },
      "source": schema.ListNestedAttribute{
        Computed: true,
        CustomType: customfield.NewNestedObjectListType[PipelineSourceDataSourceModel](ctx),
        NestedObject: schema.NestedAttributeObject{
          Attributes: map[string]schema.Attribute{
            "format": schema.StringAttribute{
              Description: "Specifies the format of source data.\nAvailable values: \"json\".",
              Computed: true,
              Validators: []validator.String{
              stringvalidator.OneOfCaseInsensitive("json"),
              },
            },
            "type": schema.StringAttribute{
              Computed: true,
            },
            "authentication": schema.BoolAttribute{
              Description: "Specifies authentication is required to send to this Pipeline.",
              Computed: true,
            },
            "cors": schema.SingleNestedAttribute{
              Computed: true,
              CustomType: customfield.NewNestedObjectType[PipelineSourceCORSDataSourceModel](ctx),
              Attributes: map[string]schema.Attribute{
                "origins": schema.ListAttribute{
                  Description: "Specifies allowed origins to allow Cross Origin HTTP Requests.",
                  Computed: true,
                  CustomType: customfield.NewListType[types.String](ctx),
                  ElementType: types.StringType,
                },
              },
            },
          },
        },
      },
    },
  }
}

func (d *PipelineDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
  resp.Schema = DataSourceSchema(ctx)
}

func (d *PipelineDataSource) ConfigValidators(_ context.Context) ([]datasource.ConfigValidator) {
  return []datasource.ConfigValidator{
  }
}
