package connectivity_directory_service

import (
	"context"

	"github.com/MakeNowJust/heredoc/v2"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r *ConnectivityDirectoryServiceResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: heredoc.Doc(`
			Manages a Cloudflare Connectivity Directory Service.

			Directory services represent applications accessible through Cloudflare tunnels,
			acting as a registry for internal services that need to be discoverable and
			accessible via Cloudflare's connectivity layer.
		`),
		Attributes: map[string]schema.Attribute{
			consts.IDSchemaKey: schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "The identifier of this resource. Same as service_id.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			consts.AccountIDSchemaKey: schema.StringAttribute{
				MarkdownDescription: consts.AccountIDSchemaDescription,
				Required:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"service_id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "The service ID assigned by Cloudflare.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"name": schema.StringAttribute{
				MarkdownDescription: "The name of the directory service.",
				Required:            true,
			},
			"type": schema.StringAttribute{
				MarkdownDescription: "The type of directory service. Currently only `http` is supported.",
				Required:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("http"),
				},
			},
			"host": schema.SingleNestedAttribute{
				MarkdownDescription: "The host configuration for the directory service.",
				Required:            true,
				Attributes: map[string]schema.Attribute{
					"ipv4": schema.StringAttribute{
						MarkdownDescription: "The IPv4 address of the service.",
						Optional:            true,
					},
					"ipv6": schema.StringAttribute{
						MarkdownDescription: "The IPv6 address of the service.",
						Optional:            true,
					},
					"hostname": schema.StringAttribute{
						MarkdownDescription: "The hostname of the service.",
						Optional:            true,
					},
					"network": schema.SingleNestedAttribute{
						MarkdownDescription: "The tunnel network configuration.",
						Optional:            true,
						Attributes: map[string]schema.Attribute{
							"tunnel_id": schema.StringAttribute{
								MarkdownDescription: "The ID of the tunnel to route through.",
								Required:            true,
							},
						},
					},
					"resolver_network": schema.SingleNestedAttribute{
						MarkdownDescription: "The resolver network configuration for hostname-based services.",
						Optional:            true,
						Attributes: map[string]schema.Attribute{
							"tunnel_id": schema.StringAttribute{
								MarkdownDescription: "The ID of the tunnel for resolver traffic.",
								Required:            true,
							},
							"resolver_ips": schema.ListAttribute{
								MarkdownDescription: "List of resolver IP addresses.",
								Optional:            true,
								ElementType:         types.StringType,
							},
						},
					},
				},
			},
			"http_port": schema.Int64Attribute{
				MarkdownDescription: "The HTTP port number for the service.",
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{
					int64validator.AtLeast(1),
					int64validator.AtMost(65535),
				},
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
			},
			"https_port": schema.Int64Attribute{
				MarkdownDescription: "The HTTPS port number for the service.",
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{
					int64validator.AtLeast(1),
					int64validator.AtMost(65535),
				},
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
			},
			"created_at": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Timestamp when the service was created.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"updated_at": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Timestamp when the service was last updated.",
			},
		},
	}
}
