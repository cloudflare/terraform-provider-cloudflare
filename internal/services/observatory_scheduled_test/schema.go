// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package observatory_scheduled_test

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

var _ resource.ResourceWithConfigValidators = (*ObservatoryScheduledTestResource)(nil)

func ResourceSchema(ctx context.Context) (schema.Schema) {
  return schema.Schema{
    Attributes: map[string]schema.Attribute{
      "id": schema.StringAttribute{
        Description: "A URL.",
        Computed: true,
        PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
      },
      "url": schema.StringAttribute{
        Description: "A URL.",
        Required: true,
        PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
      },
      "zone_id": schema.StringAttribute{
        Description: "Identifier",
        Required: true,
        PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
      },
      "frequency": schema.StringAttribute{
        Description: "The frequency of the test.\nAvailable values: \"DAILY\", \"WEEKLY\".",
        Computed: true,
        Validators: []validator.String{
        stringvalidator.OneOfCaseInsensitive("DAILY", "WEEKLY"),
        },
      },
      "region": schema.StringAttribute{
        Description: "A test region.\nAvailable values: \"asia-east1\", \"asia-northeast1\", \"asia-northeast2\", \"asia-south1\", \"asia-southeast1\", \"australia-southeast1\", \"europe-north1\", \"europe-southwest1\", \"europe-west1\", \"europe-west2\", \"europe-west3\", \"europe-west4\", \"europe-west8\", \"europe-west9\", \"me-west1\", \"southamerica-east1\", \"us-central1\", \"us-east1\", \"us-east4\", \"us-south1\", \"us-west1\".",
        Computed: true,
        Validators: []validator.String{
        stringvalidator.OneOfCaseInsensitive(
          "asia-east1",
          "asia-northeast1",
          "asia-northeast2",
          "asia-south1",
          "asia-southeast1",
          "australia-southeast1",
          "europe-north1",
          "europe-southwest1",
          "europe-west1",
          "europe-west2",
          "europe-west3",
          "europe-west4",
          "europe-west8",
          "europe-west9",
          "me-west1",
          "southamerica-east1",
          "us-central1",
          "us-east1",
          "us-east4",
          "us-south1",
          "us-west1",
        ),
        },
      },
      "schedule": schema.SingleNestedAttribute{
        Description: "The test schedule.",
        Computed: true,
        CustomType: customfield.NewNestedObjectType[ObservatoryScheduledTestScheduleModel](ctx),
        Attributes: map[string]schema.Attribute{
          "frequency": schema.StringAttribute{
            Description: "The frequency of the test.\nAvailable values: \"DAILY\", \"WEEKLY\".",
            Computed: true,
            Validators: []validator.String{
            stringvalidator.OneOfCaseInsensitive("DAILY", "WEEKLY"),
            },
          },
          "region": schema.StringAttribute{
            Description: "A test region.\nAvailable values: \"asia-east1\", \"asia-northeast1\", \"asia-northeast2\", \"asia-south1\", \"asia-southeast1\", \"australia-southeast1\", \"europe-north1\", \"europe-southwest1\", \"europe-west1\", \"europe-west2\", \"europe-west3\", \"europe-west4\", \"europe-west8\", \"europe-west9\", \"me-west1\", \"southamerica-east1\", \"us-central1\", \"us-east1\", \"us-east4\", \"us-south1\", \"us-west1\".",
            Computed: true,
            Validators: []validator.String{
            stringvalidator.OneOfCaseInsensitive(
              "asia-east1",
              "asia-northeast1",
              "asia-northeast2",
              "asia-south1",
              "asia-southeast1",
              "australia-southeast1",
              "europe-north1",
              "europe-southwest1",
              "europe-west1",
              "europe-west2",
              "europe-west3",
              "europe-west4",
              "europe-west8",
              "europe-west9",
              "me-west1",
              "southamerica-east1",
              "us-central1",
              "us-east1",
              "us-east4",
              "us-south1",
              "us-west1",
            ),
            },
          },
          "url": schema.StringAttribute{
            Description: "A URL.",
            Computed: true,
          },
        },
      },
      "test": schema.SingleNestedAttribute{
        Computed: true,
        CustomType: customfield.NewNestedObjectType[ObservatoryScheduledTestTestModel](ctx),
        Attributes: map[string]schema.Attribute{
          "id": schema.StringAttribute{
            Description: "UUID",
            Computed: true,
          },
          "date": schema.StringAttribute{
            Computed: true,
            CustomType: timetypes.RFC3339Type{

            },
          },
          "desktop_report": schema.SingleNestedAttribute{
            Description: "The Lighthouse report.",
            Computed: true,
            CustomType: customfield.NewNestedObjectType[ObservatoryScheduledTestTestDesktopReportModel](ctx),
            Attributes: map[string]schema.Attribute{
              "cls": schema.Float64Attribute{
                Description: "Cumulative Layout Shift.",
                Computed: true,
              },
              "device_type": schema.StringAttribute{
                Description: "The type of device.\nAvailable values: \"DESKTOP\", \"MOBILE\".",
                Computed: true,
                Validators: []validator.String{
                stringvalidator.OneOfCaseInsensitive("DESKTOP", "MOBILE"),
                },
              },
              "error": schema.SingleNestedAttribute{
                Computed: true,
                CustomType: customfield.NewNestedObjectType[ObservatoryScheduledTestTestDesktopReportErrorModel](ctx),
                Attributes: map[string]schema.Attribute{
                  "code": schema.StringAttribute{
                    Description: "The error code of the Lighthouse result.\nAvailable values: \"NOT_REACHABLE\", \"DNS_FAILURE\", \"NOT_HTML\", \"LIGHTHOUSE_TIMEOUT\", \"UNKNOWN\".",
                    Computed: true,
                    Validators: []validator.String{
                    stringvalidator.OneOfCaseInsensitive(
                      "NOT_REACHABLE",
                      "DNS_FAILURE",
                      "NOT_HTML",
                      "LIGHTHOUSE_TIMEOUT",
                      "UNKNOWN",
                    ),
                    },
                  },
                  "detail": schema.StringAttribute{
                    Description: "Detailed error message.",
                    Computed: true,
                  },
                  "final_displayed_url": schema.StringAttribute{
                    Description: "The final URL displayed to the user.",
                    Computed: true,
                  },
                },
              },
              "fcp": schema.Float64Attribute{
                Description: "First Contentful Paint.",
                Computed: true,
              },
              "json_report_url": schema.StringAttribute{
                Description: "The URL to the full Lighthouse JSON report.",
                Computed: true,
              },
              "lcp": schema.Float64Attribute{
                Description: "Largest Contentful Paint.",
                Computed: true,
              },
              "performance_score": schema.Float64Attribute{
                Description: "The Lighthouse performance score.",
                Computed: true,
              },
              "si": schema.Float64Attribute{
                Description: "Speed Index.",
                Computed: true,
              },
              "state": schema.StringAttribute{
                Description: "The state of the Lighthouse report.\nAvailable values: \"RUNNING\", \"COMPLETE\", \"FAILED\".",
                Computed: true,
                Validators: []validator.String{
                stringvalidator.OneOfCaseInsensitive(
                  "RUNNING",
                  "COMPLETE",
                  "FAILED",
                ),
                },
              },
              "tbt": schema.Float64Attribute{
                Description: "Total Blocking Time.",
                Computed: true,
              },
              "ttfb": schema.Float64Attribute{
                Description: "Time To First Byte.",
                Computed: true,
              },
              "tti": schema.Float64Attribute{
                Description: "Time To Interactive.",
                Computed: true,
              },
            },
          },
          "mobile_report": schema.SingleNestedAttribute{
            Description: "The Lighthouse report.",
            Computed: true,
            CustomType: customfield.NewNestedObjectType[ObservatoryScheduledTestTestMobileReportModel](ctx),
            Attributes: map[string]schema.Attribute{
              "cls": schema.Float64Attribute{
                Description: "Cumulative Layout Shift.",
                Computed: true,
              },
              "device_type": schema.StringAttribute{
                Description: "The type of device.\nAvailable values: \"DESKTOP\", \"MOBILE\".",
                Computed: true,
                Validators: []validator.String{
                stringvalidator.OneOfCaseInsensitive("DESKTOP", "MOBILE"),
                },
              },
              "error": schema.SingleNestedAttribute{
                Computed: true,
                CustomType: customfield.NewNestedObjectType[ObservatoryScheduledTestTestMobileReportErrorModel](ctx),
                Attributes: map[string]schema.Attribute{
                  "code": schema.StringAttribute{
                    Description: "The error code of the Lighthouse result.\nAvailable values: \"NOT_REACHABLE\", \"DNS_FAILURE\", \"NOT_HTML\", \"LIGHTHOUSE_TIMEOUT\", \"UNKNOWN\".",
                    Computed: true,
                    Validators: []validator.String{
                    stringvalidator.OneOfCaseInsensitive(
                      "NOT_REACHABLE",
                      "DNS_FAILURE",
                      "NOT_HTML",
                      "LIGHTHOUSE_TIMEOUT",
                      "UNKNOWN",
                    ),
                    },
                  },
                  "detail": schema.StringAttribute{
                    Description: "Detailed error message.",
                    Computed: true,
                  },
                  "final_displayed_url": schema.StringAttribute{
                    Description: "The final URL displayed to the user.",
                    Computed: true,
                  },
                },
              },
              "fcp": schema.Float64Attribute{
                Description: "First Contentful Paint.",
                Computed: true,
              },
              "json_report_url": schema.StringAttribute{
                Description: "The URL to the full Lighthouse JSON report.",
                Computed: true,
              },
              "lcp": schema.Float64Attribute{
                Description: "Largest Contentful Paint.",
                Computed: true,
              },
              "performance_score": schema.Float64Attribute{
                Description: "The Lighthouse performance score.",
                Computed: true,
              },
              "si": schema.Float64Attribute{
                Description: "Speed Index.",
                Computed: true,
              },
              "state": schema.StringAttribute{
                Description: "The state of the Lighthouse report.\nAvailable values: \"RUNNING\", \"COMPLETE\", \"FAILED\".",
                Computed: true,
                Validators: []validator.String{
                stringvalidator.OneOfCaseInsensitive(
                  "RUNNING",
                  "COMPLETE",
                  "FAILED",
                ),
                },
              },
              "tbt": schema.Float64Attribute{
                Description: "Total Blocking Time.",
                Computed: true,
              },
              "ttfb": schema.Float64Attribute{
                Description: "Time To First Byte.",
                Computed: true,
              },
              "tti": schema.Float64Attribute{
                Description: "Time To Interactive.",
                Computed: true,
              },
            },
          },
          "region": schema.SingleNestedAttribute{
            Description: "A test region with a label.",
            Computed: true,
            CustomType: customfield.NewNestedObjectType[ObservatoryScheduledTestTestRegionModel](ctx),
            Attributes: map[string]schema.Attribute{
              "label": schema.StringAttribute{
                Computed: true,
              },
              "value": schema.StringAttribute{
                Description: "A test region.\nAvailable values: \"asia-east1\", \"asia-northeast1\", \"asia-northeast2\", \"asia-south1\", \"asia-southeast1\", \"australia-southeast1\", \"europe-north1\", \"europe-southwest1\", \"europe-west1\", \"europe-west2\", \"europe-west3\", \"europe-west4\", \"europe-west8\", \"europe-west9\", \"me-west1\", \"southamerica-east1\", \"us-central1\", \"us-east1\", \"us-east4\", \"us-south1\", \"us-west1\".",
                Computed: true,
                Validators: []validator.String{
                stringvalidator.OneOfCaseInsensitive(
                  "asia-east1",
                  "asia-northeast1",
                  "asia-northeast2",
                  "asia-south1",
                  "asia-southeast1",
                  "australia-southeast1",
                  "europe-north1",
                  "europe-southwest1",
                  "europe-west1",
                  "europe-west2",
                  "europe-west3",
                  "europe-west4",
                  "europe-west8",
                  "europe-west9",
                  "me-west1",
                  "southamerica-east1",
                  "us-central1",
                  "us-east1",
                  "us-east4",
                  "us-south1",
                  "us-west1",
                ),
                },
              },
            },
          },
          "schedule_frequency": schema.StringAttribute{
            Description: "The frequency of the test.\nAvailable values: \"DAILY\", \"WEEKLY\".",
            Computed: true,
            Validators: []validator.String{
            stringvalidator.OneOfCaseInsensitive("DAILY", "WEEKLY"),
            },
          },
          "url": schema.StringAttribute{
            Description: "A URL.",
            Computed: true,
          },
        },
      },
    },
  }
}

func (r *ObservatoryScheduledTestResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
  resp.Schema = ResourceSchema(ctx)
}

func (r *ObservatoryScheduledTestResource) ConfigValidators(_ context.Context) ([]resource.ConfigValidator) {
  return []resource.ConfigValidator{
  }
}
