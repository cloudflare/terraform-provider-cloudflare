package connectivity_directory_service

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (d *ConnectivityDirectoryServiceDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Use this data source to look up a single Connectivity Directory Service by service ID.",
		Attributes: map[string]schema.Attribute{
			consts.IDSchemaKey: schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "The identifier of the service. Same as service_id.",
			},
			consts.AccountIDSchemaKey: schema.StringAttribute{
				MarkdownDescription: consts.AccountIDSchemaDescription,
				Required:            true,
			},
			"service_id": schema.StringAttribute{
				MarkdownDescription: "The service ID to look up.",
				Required:            true,
			},
			"name": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "The name of the directory service.",
			},
			"type": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "The type of directory service.",
			},
			"host": schema.SingleNestedAttribute{
				Computed:            true,
				MarkdownDescription: "The host configuration for the directory service.",
				Attributes: map[string]schema.Attribute{
					"ipv4": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: "The IPv4 address of the service.",
					},
					"ipv6": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: "The IPv6 address of the service.",
					},
					"hostname": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: "The hostname of the service.",
					},
					"network": schema.SingleNestedAttribute{
						Computed:            true,
						MarkdownDescription: "The tunnel network configuration.",
						Attributes: map[string]schema.Attribute{
							"tunnel_id": schema.StringAttribute{
								Computed:            true,
								MarkdownDescription: "The ID of the tunnel to route through.",
							},
						},
					},
					"resolver_network": schema.SingleNestedAttribute{
						Computed:            true,
						MarkdownDescription: "The resolver network configuration for hostname-based services.",
						Attributes: map[string]schema.Attribute{
							"tunnel_id": schema.StringAttribute{
								Computed:            true,
								MarkdownDescription: "The ID of the tunnel for resolver traffic.",
							},
							"resolver_ips": schema.ListAttribute{
								Computed:            true,
								MarkdownDescription: "List of resolver IP addresses.",
								ElementType:         types.StringType,
							},
						},
					},
				},
			},
			"http_port": schema.Int64Attribute{
				Computed:            true,
				MarkdownDescription: "The HTTP port number for the service.",
			},
			"https_port": schema.Int64Attribute{
				Computed:            true,
				MarkdownDescription: "The HTTPS port number for the service.",
			},
			"created_at": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Timestamp when the service was created.",
			},
			"updated_at": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Timestamp when the service was last updated.",
			},
		},
	}
}
