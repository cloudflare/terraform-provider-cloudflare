// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package queue_consumer

import (
  "context"

  "github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
  "github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
  "github.com/hashicorp/terraform-plugin-framework/datasource"
  "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
  "github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ datasource.DataSourceWithConfigValidators = (*QueueConsumerDataSource)(nil)

func DataSourceSchema(ctx context.Context) (schema.Schema) {
  return schema.Schema{
    Attributes: map[string]schema.Attribute{
      "account_id": schema.StringAttribute{
        Description: "A Resource identifier.",
        Required: true,
      },
      "queue_id": schema.StringAttribute{
        Description: "A Resource identifier.",
        Required: true,
      },
      "consumer_id": schema.StringAttribute{
        Description: "A Resource identifier.",
        Computed: true,
      },
      "created_on": schema.StringAttribute{
        Computed: true,
      },
      "script": schema.StringAttribute{
        Description: "Name of a Worker",
        Computed: true,
      },
      "script_name": schema.StringAttribute{
        Description: "Name of a Worker",
        Computed: true,
      },
      "type": schema.StringAttribute{
        Description: `Available values: "worker".`,
        Computed: true,
        Validators: []validator.String{
        stringvalidator.OneOfCaseInsensitive("worker", "http_pull"),
        },
      },
      "settings": schema.SingleNestedAttribute{
        Computed: true,
        CustomType: customfield.NewNestedObjectType[QueueConsumerSettingsDataSourceModel](ctx),
        Attributes: map[string]schema.Attribute{
          "batch_size": schema.Float64Attribute{
            Description: "The maximum number of messages to include in a batch.",
            Computed: true,
          },
          "max_concurrency": schema.Float64Attribute{
            Description: "Maximum number of concurrent consumers that may consume from this Queue. Set to `null` to automatically opt in to the platform's maximum (recommended).",
            Computed: true,
          },
          "max_retries": schema.Float64Attribute{
            Description: "The maximum number of retries",
            Computed: true,
          },
          "max_wait_time_ms": schema.Float64Attribute{
            Description: "The number of milliseconds to wait for a batch to fill up before attempting to deliver it",
            Computed: true,
          },
          "retry_delay": schema.Float64Attribute{
            Description: "The number of seconds to delay before making the message available for another attempt.",
            Computed: true,
          },
          "visibility_timeout_ms": schema.Float64Attribute{
            Description: "The number of milliseconds that a message is exclusively leased. After the timeout, the message becomes available for another attempt.",
            Computed: true,
          },
        },
      },
    },
  }
}

func (d *QueueConsumerDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
  resp.Schema = DataSourceSchema(ctx)
}

func (d *QueueConsumerDataSource) ConfigValidators(_ context.Context) ([]datasource.ConfigValidator) {
  return []datasource.ConfigValidator{
  }
}
