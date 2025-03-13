// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package api_shield_operation

import (
  "context"

  "github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
  "github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
  "github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
  "github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
  "github.com/hashicorp/terraform-plugin-framework/resource"
  "github.com/hashicorp/terraform-plugin-framework/resource/schema"
  "github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
  "github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
  "github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ resource.ResourceWithConfigValidators = (*APIShieldOperationResource)(nil)

func ResourceSchema(ctx context.Context) (schema.Schema) {
  return schema.Schema{
    Attributes: map[string]schema.Attribute{
      "id": schema.StringAttribute{
        Description: "UUID",
        Computed: true,
        PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
      },
      "operation_id": schema.StringAttribute{
        Description: "UUID",
        Computed: true,
        PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
      },
      "zone_id": schema.StringAttribute{
        Description: "Identifier",
        Required: true,
        PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
      },
      "endpoint": schema.StringAttribute{
        Description: "The endpoint which can contain path parameter templates in curly braces, each will be replaced from left to right with {varN}, starting with {var1}, during insertion. This will further be Cloudflare-normalized upon insertion. See: https://developers.cloudflare.com/rules/normalization/how-it-works/.",
        Required: true,
        PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
      },
      "host": schema.StringAttribute{
        Description: "RFC3986-compliant host.",
        Required: true,
        PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
      },
      "method": schema.StringAttribute{
        Description: "The HTTP method used to access the endpoint.\nAvailable values: \"GET\", \"POST\", \"HEAD\", \"OPTIONS\", \"PUT\", \"DELETE\", \"CONNECT\", \"PATCH\", \"TRACE\".",
        Required: true,
        Validators: []validator.String{
        stringvalidator.OneOfCaseInsensitive(
          "GET",
          "POST",
          "HEAD",
          "OPTIONS",
          "PUT",
          "DELETE",
          "CONNECT",
          "PATCH",
          "TRACE",
        ),
        },
        PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
      },
      "last_updated": schema.StringAttribute{
        Computed: true,
        CustomType: timetypes.RFC3339Type{

        },
      },
      "features": schema.SingleNestedAttribute{
        Computed: true,
        CustomType: customfield.NewNestedObjectType[APIShieldOperationFeaturesModel](ctx),
        Attributes: map[string]schema.Attribute{
          "thresholds": schema.SingleNestedAttribute{
            Computed: true,
            CustomType: customfield.NewNestedObjectType[APIShieldOperationFeaturesThresholdsModel](ctx),
            Attributes: map[string]schema.Attribute{
              "auth_id_tokens": schema.Int64Attribute{
                Description: "The total number of auth-ids seen across this calculation.",
                Computed: true,
              },
              "data_points": schema.Int64Attribute{
                Description: "The number of data points used for the threshold suggestion calculation.",
                Computed: true,
              },
              "last_updated": schema.StringAttribute{
                Computed: true,
                CustomType: timetypes.RFC3339Type{

                },
              },
              "p50": schema.Int64Attribute{
                Description: "The p50 quantile of requests (in period_seconds).",
                Computed: true,
              },
              "p90": schema.Int64Attribute{
                Description: "The p90 quantile of requests (in period_seconds).",
                Computed: true,
              },
              "p99": schema.Int64Attribute{
                Description: "The p99 quantile of requests (in period_seconds).",
                Computed: true,
              },
              "period_seconds": schema.Int64Attribute{
                Description: "The period over which this threshold is suggested.",
                Computed: true,
              },
              "requests": schema.Int64Attribute{
                Description: "The estimated number of requests covered by these calculations.",
                Computed: true,
              },
              "suggested_threshold": schema.Int64Attribute{
                Description: "The suggested threshold in requests done by the same auth_id or period_seconds.",
                Computed: true,
              },
            },
          },
          "parameter_schemas": schema.SingleNestedAttribute{
            Computed: true,
            CustomType: customfield.NewNestedObjectType[APIShieldOperationFeaturesParameterSchemasModel](ctx),
            Attributes: map[string]schema.Attribute{
              "last_updated": schema.StringAttribute{
                Computed: true,
                CustomType: timetypes.RFC3339Type{

                },
              },
              "parameter_schemas": schema.SingleNestedAttribute{
                Description: "An operation schema object containing a response.",
                Computed: true,
                CustomType: customfield.NewNestedObjectType[APIShieldOperationFeaturesParameterSchemasParameterSchemasModel](ctx),
                Attributes: map[string]schema.Attribute{
                  "parameters": schema.ListAttribute{
                    Description: "An array containing the learned parameter schemas.",
                    Computed: true,
                    CustomType: customfield.NewListType[jsontypes.Normalized](ctx),
                    ElementType: jsontypes.NormalizedType{

                    },
                  },
                  "responses": schema.StringAttribute{
                    Description: "An empty response object. This field is required to yield a valid operation schema.",
                    Computed: true,
                    CustomType: jsontypes.NormalizedType{

                    },
                  },
                },
              },
            },
          },
          "api_routing": schema.SingleNestedAttribute{
            Description: "API Routing settings on endpoint.",
            Computed: true,
            CustomType: customfield.NewNestedObjectType[APIShieldOperationFeaturesAPIRoutingModel](ctx),
            Attributes: map[string]schema.Attribute{
              "last_updated": schema.StringAttribute{
                Computed: true,
                CustomType: timetypes.RFC3339Type{

                },
              },
              "route": schema.StringAttribute{
                Description: "Target route.",
                Computed: true,
              },
            },
          },
          "confidence_intervals": schema.SingleNestedAttribute{
            Computed: true,
            CustomType: customfield.NewNestedObjectType[APIShieldOperationFeaturesConfidenceIntervalsModel](ctx),
            Attributes: map[string]schema.Attribute{
              "last_updated": schema.StringAttribute{
                Computed: true,
                CustomType: timetypes.RFC3339Type{

                },
              },
              "suggested_threshold": schema.SingleNestedAttribute{
                Computed: true,
                CustomType: customfield.NewNestedObjectType[APIShieldOperationFeaturesConfidenceIntervalsSuggestedThresholdModel](ctx),
                Attributes: map[string]schema.Attribute{
                  "confidence_intervals": schema.SingleNestedAttribute{
                    Computed: true,
                    CustomType: customfield.NewNestedObjectType[APIShieldOperationFeaturesConfidenceIntervalsSuggestedThresholdConfidenceIntervalsModel](ctx),
                    Attributes: map[string]schema.Attribute{
                      "p90": schema.SingleNestedAttribute{
                        Description: "Upper and lower bound for percentile estimate",
                        Computed: true,
                        CustomType: customfield.NewNestedObjectType[APIShieldOperationFeaturesConfidenceIntervalsSuggestedThresholdConfidenceIntervalsP90Model](ctx),
                        Attributes: map[string]schema.Attribute{
                          "lower": schema.Float64Attribute{
                            Description: "Lower bound for percentile estimate",
                            Computed: true,
                          },
                          "upper": schema.Float64Attribute{
                            Description: "Upper bound for percentile estimate",
                            Computed: true,
                          },
                        },
                      },
                      "p95": schema.SingleNestedAttribute{
                        Description: "Upper and lower bound for percentile estimate",
                        Computed: true,
                        CustomType: customfield.NewNestedObjectType[APIShieldOperationFeaturesConfidenceIntervalsSuggestedThresholdConfidenceIntervalsP95Model](ctx),
                        Attributes: map[string]schema.Attribute{
                          "lower": schema.Float64Attribute{
                            Description: "Lower bound for percentile estimate",
                            Computed: true,
                          },
                          "upper": schema.Float64Attribute{
                            Description: "Upper bound for percentile estimate",
                            Computed: true,
                          },
                        },
                      },
                      "p99": schema.SingleNestedAttribute{
                        Description: "Upper and lower bound for percentile estimate",
                        Computed: true,
                        CustomType: customfield.NewNestedObjectType[APIShieldOperationFeaturesConfidenceIntervalsSuggestedThresholdConfidenceIntervalsP99Model](ctx),
                        Attributes: map[string]schema.Attribute{
                          "lower": schema.Float64Attribute{
                            Description: "Lower bound for percentile estimate",
                            Computed: true,
                          },
                          "upper": schema.Float64Attribute{
                            Description: "Upper bound for percentile estimate",
                            Computed: true,
                          },
                        },
                      },
                    },
                  },
                  "mean": schema.Float64Attribute{
                    Description: "Suggested threshold.",
                    Computed: true,
                  },
                },
              },
            },
          },
          "schema_info": schema.SingleNestedAttribute{
            Computed: true,
            CustomType: customfield.NewNestedObjectType[APIShieldOperationFeaturesSchemaInfoModel](ctx),
            Attributes: map[string]schema.Attribute{
              "active_schema": schema.SingleNestedAttribute{
                Description: "Schema active on endpoint.",
                Computed: true,
                CustomType: customfield.NewNestedObjectType[APIShieldOperationFeaturesSchemaInfoActiveSchemaModel](ctx),
                Attributes: map[string]schema.Attribute{
                  "id": schema.StringAttribute{
                    Description: "UUID",
                    Computed: true,
                  },
                  "created_at": schema.StringAttribute{
                    Computed: true,
                    CustomType: timetypes.RFC3339Type{

                    },
                  },
                  "is_learned": schema.BoolAttribute{
                    Description: "True if schema is Cloudflare-provided.",
                    Computed: true,
                  },
                  "name": schema.StringAttribute{
                    Description: "Schema file name.",
                    Computed: true,
                  },
                },
              },
              "learned_available": schema.BoolAttribute{
                Description: "True if a Cloudflare-provided learned schema is available for this endpoint.",
                Computed: true,
              },
              "mitigation_action": schema.StringAttribute{
                Description: "Action taken on requests failing validation.\nAvailable values: \"none\", \"log\", \"block\".",
                Computed: true,
                Validators: []validator.String{
                stringvalidator.OneOfCaseInsensitive(
                  "none",
                  "log",
                  "block",
                ),
                },
              },
            },
          },
        },
      },
    },
  }
}

func (r *APIShieldOperationResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
  resp.Schema = ResourceSchema(ctx)
}

func (r *APIShieldOperationResource) ConfigValidators(_ context.Context) ([]resource.ConfigValidator) {
  return []resource.ConfigValidator{
  }
}
