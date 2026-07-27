// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package moq_relay

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/datasourcevalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ datasource.DataSourceWithConfigValidators = (*MoQRelayDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"relay_id": schema.StringAttribute{
				Optional: true,
			},
			"account_id": schema.StringAttribute{
				Description: "Cloudflare account identifier.",
				Required:    true,
			},
			"created": schema.StringAttribute{
				Computed:   true,
				CustomType: timetypes.RFC3339Type{},
			},
			"modified": schema.StringAttribute{
				Computed:   true,
				CustomType: timetypes.RFC3339Type{},
			},
			"name": schema.StringAttribute{
				Computed: true,
			},
			"status": schema.StringAttribute{
				Description: "\"connected\" when active, omitted otherwise.\nAvailable values: \"connected\".",
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("connected"),
				},
			},
			"uid": schema.StringAttribute{
				Computed: true,
			},
			"config": schema.SingleNestedAttribute{
				Description: "upstreams and lingering_subscribe are mutually exclusive.",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectType[MoQRelayConfigDataSourceModel](ctx),
				Attributes: map[string]schema.Attribute{
					"lingering_subscribe": schema.SingleNestedAttribute{
						Computed:   true,
						CustomType: customfield.NewNestedObjectType[MoQRelayConfigLingeringSubscribeDataSourceModel](ctx),
						Attributes: map[string]schema.Attribute{
							"enabled": schema.BoolAttribute{
								Computed: true,
							},
							"max_timeout_ms": schema.Int64Attribute{
								Description: "Relay-level ceiling on lingering subscribe timeout (ms). Default 30000.",
								Computed:    true,
								Validators: []validator.Int64{
									int64validator.Between(0, 300000),
								},
							},
						},
					},
					"upstreams": schema.SingleNestedAttribute{
						Description: "Upstreams are external MOQT server publishers that a relay falls back\nto when it has no local publisher for a requested namespace/track.",
						Computed:    true,
						CustomType:  customfield.NewNestedObjectType[MoQRelayConfigUpstreamsDataSourceModel](ctx),
						Attributes: map[string]schema.Attribute{
							"enabled": schema.BoolAttribute{
								Computed: true,
							},
							"upstreams": schema.ListNestedAttribute{
								Description: "Ordered list of upstream MOQT server publishers. Each entry is an\nobject (not a bare string) so per-upstream configuration can be\nadded in the future without another breaking change.",
								Computed:    true,
								CustomType:  customfield.NewNestedObjectListType[MoQRelayConfigUpstreamsUpstreamsDataSourceModel](ctx),
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"url": schema.StringAttribute{
											Description: "Upstream MOQT server publisher URL.",
											Computed:    true,
										},
									},
								},
							},
						},
					},
				},
			},
			"filter": schema.SingleNestedAttribute{
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"asc": schema.BoolAttribute{
						Description: "Sort order by `created`. When true, results are returned oldest-first\n(ascending); otherwise newest-first (descending, the default).",
						Computed:    true,
						Optional:    true,
					},
					"created_after": schema.StringAttribute{
						Description: "Cursor for pagination. Returns relays created strictly after this\nRFC 3339 timestamp (typically the `created` value of the last item\non the current page, to fetch the next page).",
						Optional:    true,
						CustomType:  timetypes.RFC3339Type{},
					},
					"created_before": schema.StringAttribute{
						Description: "Cursor for pagination. Returns relays created strictly before this\nRFC 3339 timestamp (typically the `created` value of the first item\non the current page, to fetch the previous page).",
						Optional:    true,
						CustomType:  timetypes.RFC3339Type{},
					},
					"per_page": schema.Int64Attribute{
						Description: "Maximum number of relays to return per page.",
						Optional:    true,
						Validators: []validator.Int64{
							int64validator.AtLeast(1),
						},
					},
				},
			},
		},
	}
}

func (d *MoQRelayDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *MoQRelayDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{
		datasourcevalidator.ExactlyOneOf(path.MatchRoot("relay_id"), path.MatchRoot("filter")),
	}
}
