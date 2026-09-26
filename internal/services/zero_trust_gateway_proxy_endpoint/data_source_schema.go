// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_gateway_proxy_endpoint

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/datasourcevalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ datasource.DataSourceWithConfigValidators = (*ZeroTrustGatewayProxyEndpointDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"proxy_endpoint_id": schema.StringAttribute{
				Optional: true,
			},
			"account_id": schema.StringAttribute{
				Optional:    true,
			},
			"created_at": schema.StringAttribute{
				Computed:   true,
				CustomType: timetypes.RFC3339Type{},
			},
			"kind": schema.StringAttribute{
				Description: "The proxy endpoint kind\nAvailable values: \"ip\", \"identity\".",
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("ip", "identity"),
				},
			},
			"name": schema.StringAttribute{
				Description: "Specify the name of the proxy endpoint.",
				Computed:    true,
			},
			"subdomain": schema.StringAttribute{
				Description: "Specify the subdomain to use as the destination in the proxy client.",
				Computed:    true,
			},
			"updated_at": schema.StringAttribute{
				Computed:   true,
				CustomType: timetypes.RFC3339Type{},
			},
			"ips": schema.ListAttribute{
				Description: "Specify the list of CIDRs to restrict ingress connections.",
				Computed:    true,
				CustomType:  customfield.NewListType[types.String](ctx),
				ElementType: types.StringType,
			},
			"filter": schema.SingleNestedAttribute{
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"direction": schema.StringAttribute{
						Description: "Sort direction. Only takes effect when `order_by` is also provided; it\nis ignored otherwise. When `direction` is omitted the effective\ndirection is field-specific: `created_at` and `updated_at` default to\ndescending (newest first); `name` defaults to ascending.\n  * `asc` — ascending.\n  * `desc` — descending.\nAvailable values: \"asc\", \"desc\".",
						Optional:    true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive("asc", "desc"),
						},
					},
					"filter": schema.ListAttribute{
						Description: "Filter the returned proxy endpoints by one or more `field:value` pairs.\nRepeat the parameter to apply multiple filters; they are combined with\nlogical AND (an endpoint must satisfy every filter to be returned).\n\nSupported fields and their matching behaviour:\n  * `name` — case-insensitive substring match on the endpoint name.\n  * `id` — substring match on the endpoint ID (UUID), with or without dashes.\n  * `kind` — exact match on the endpoint kind. The value must be `ip` or `identity`; any other value returns `400`.\n\nEach entry must match one of the per-field patterns below: the field\nmust be one of `name`, `id`, or `kind`; `name`/`id` accept any value,\nwhile `kind` only accepts `ip` or `identity`.",
						Optional:    true,
						ElementType: types.StringType,
					},
					"order_by": schema.StringAttribute{
						Description: "Field to sort the returned endpoints by. When omitted, the order of\nresults is unspecified. Supported values:\n  * `name` — sort alphabetically by endpoint name.\n  * `created_at` — sort by creation time; defaults to descending unless `direction` is set.\n  * `updated_at` — sort by last-modified time; defaults to descending unless `direction` is set.\nAvailable values: \"name\", \"created_at\", \"updated_at\".",
						Optional:    true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive(
								"name",
								"created_at",
								"updated_at",
							),
						},
					},
					"search": schema.StringAttribute{
						Description: "Case-insensitive substring match on the endpoint name. When combined\nwith `filter`, both must match (logical AND).",
						Optional:    true,
					},
				},
			},
		},
	}
}

func (d *ZeroTrustGatewayProxyEndpointDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *ZeroTrustGatewayProxyEndpointDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{
		datasourcevalidator.ExactlyOneOf(path.MatchRoot("proxy_endpoint_id"), path.MatchRoot("filter")),
	}
}
