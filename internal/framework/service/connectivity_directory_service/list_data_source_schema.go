package connectivity_directory_service

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (d *ConnectivityDirectoryServicesDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Use this data source to list Connectivity Directory Services for an account.",
		Attributes: map[string]schema.Attribute{
			consts.IDSchemaKey: schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "The identifier for this data source. Same as account_id.",
			},
			consts.AccountIDSchemaKey: schema.StringAttribute{
				MarkdownDescription: consts.AccountIDSchemaDescription,
				Required:            true,
			},
			"type": schema.StringAttribute{
				MarkdownDescription: "Filter services by type. Currently only `http` is supported.",
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("http"),
				},
			},
			"services": schema.ListNestedAttribute{
				Computed:            true,
				MarkdownDescription: "The list of directory services.",
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: "The identifier of the service. Same as service_id.",
						},
						"service_id": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: "The service ID assigned by Cloudflare.",
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
				},
			},
		},
	}
}
