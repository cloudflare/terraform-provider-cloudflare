// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_cws_tenant_structure_node

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ datasource.DataSourceWithConfigValidators = (*ZeroTrustCWSTenantStructureNodesDataSource)(nil)

func ListDataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"account_id": schema.StringAttribute{
				Description: "Cloudflare account ID (32-character hexadecimal string)",
				Required:    true,
			},
			"structure_id": schema.StringAttribute{
				Description: "ID of the parent tenant structure",
				Required:    true,
			},
			"name": schema.StringAttribute{
				Description: "Filter by node name (case-insensitive substring match)",
				Optional:    true,
			},
			"native_id": schema.StringAttribute{
				Description: "Filter by exact cloud-native identifier",
				Optional:    true,
			},
			"node_kind": schema.StringAttribute{
				Description: "Filter by exact node kind",
				Optional:    true,
			},
			"parent": schema.StringAttribute{
				Description: "Filter by exact parent native ID",
				Optional:    true,
			},
			"direction": schema.StringAttribute{
				Description: "Sort direction (default: asc)\nAvailable values: \"asc\", \"desc\".",
				Computed:    true,
				Optional:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("asc", "desc"),
				},
			},
			"order": schema.StringAttribute{
				Description: "Field to order results by (default: created_at)\nAvailable values: \"name\", \"created_at\", \"updated_at\".",
				Computed:    true,
				Optional:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive(
						"name",
						"created_at",
						"updated_at",
					),
				},
			},
			"max_items": schema.Int64Attribute{
				Description: "Max items to fetch, default: 1000",
				Optional:    true,
				Validators: []validator.Int64{
					int64validator.AtLeast(0),
				},
			},
			"result": schema.ListNestedAttribute{
				Description: "The items returned by the data source",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectListType[ZeroTrustCWSTenantStructureNodesResultDataSourceModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Description: "Server-assigned UUID",
							Computed:    true,
						},
						"created_at": schema.StringAttribute{
							Computed:   true,
							CustomType: timetypes.RFC3339Type{},
						},
						"native_id": schema.StringAttribute{
							Description: "Cloud-native identifier of this node, unique within the structure; the parent-reference target. Shape depends on the structure's application. Immutable after create.",
							Computed:    true,
						},
						"node_kind": schema.StringAttribute{
							Description: "Position in the tree. Legal values depend on the structure's application (aws: root/ou/account; cloudflare: org/account/zone). At most one root per structure. Immutable after create.",
							Computed:    true,
						},
						"updated_at": schema.StringAttribute{
							Computed:   true,
							CustomType: timetypes.RFC3339Type{},
						},
						"name": schema.StringAttribute{
							Description: "Display name (e.g. an OU or account name), mutable. Optional; empty by default.",
							Computed:    true,
						},
						"parent": schema.StringAttribute{
							Description: "Native ID of this node's parent. Required for non-root nodes and omitted for the root. The parent must be in the same structure and able to contain this node kind.",
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func (d *ZeroTrustCWSTenantStructureNodesDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = ListDataSourceSchema(ctx)
}

func (d *ZeroTrustCWSTenantStructureNodesDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
