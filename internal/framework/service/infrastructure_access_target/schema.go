package infrastructure_access_target

import (
	"context"

	"github.com/MakeNowJust/heredoc/v2"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	dschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
)

// Resource Schema
func (r *InfrastructureAccessTargetResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: heredoc.Doc(`
			The [Infrastructure Access Target](https://developers.cloudflare.com/cloudflare-one/insights/risk-score/) resource allows you to configure Cloudflare Risk Behaviors for an account.
		`),
		Attributes: map[string]schema.Attribute{
			consts.AccountIDSchemaKey: schema.StringAttribute{
				MarkdownDescription: consts.AccountIDSchemaDescription,
				Required:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			consts.IDSchemaKey: schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: consts.IDSchemaDescription + " This is target's unique identifier.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"hostname": schema.StringAttribute{
				MarkdownDescription: "A non-unique field that refers to a target.",
				Required:            true,
			},
			"ip": schema.SingleNestedAttribute{
				MarkdownDescription: "The IPv4/IPv6 address that identifies where to reach a target.",
				Required:            true,
				Attributes: map[string]schema.Attribute{
					"ipv4": schema.SingleNestedAttribute{
						MarkdownDescription: "The target's IPv4 address.",
						Optional:            true,
						Attributes: map[string]schema.Attribute{
							"ip_addr": schema.StringAttribute{
								MarkdownDescription: "The IP address of the target.",
								Required:            true,
							},
							"virtual_network_id": schema.StringAttribute{
								MarkdownDescription: "The private virtual network identifier for the target.",
								Required:            true,
							},
						},
					},
					"ipv6": schema.SingleNestedAttribute{
						MarkdownDescription: "The target's IPv6 address.",
						Optional:            true,
						Attributes: map[string]schema.Attribute{
							"ip_addr": schema.StringAttribute{
								MarkdownDescription: "The IP address of the target.",
								Required:            true,
							},
							"virtual_network_id": schema.StringAttribute{
								MarkdownDescription: "The private virtual network identifier for the target.",
								Required:            true,
							},
						},
					},
				},
			},
			"created_at": schema.StringAttribute{
				MarkdownDescription: "The date and time at which the target was created.",
				// Set value to read-only
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"modified_at": schema.StringAttribute{
				MarkdownDescription: "The date and time at which the target was last modified.",
				// Set value to read-only
				Computed: true,
			},
		},
	}
}

// Data Source schema
// This should be the model/request
func (d *InfrastructureAccessTargetDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = dschema.Schema{
		MarkdownDescription: "Use this data source to retrieve all Infrastructure Access Targets.",
		Attributes: map[string]dschema.Attribute{
			consts.AccountIDSchemaKey: dschema.StringAttribute{
				MarkdownDescription: consts.AccountIDSchemaDescription,
				Required:            true,
			},
			"hostname": dschema.StringAttribute{
				Optional:    true,
				Description: "The name of the app type.",
			},
			"hostname_contains": dschema.StringAttribute{
				Optional:    true,
				Description: "The name of the app type.",
			},
			"ip_v4": dschema.StringAttribute{
				Optional:    true,
				Description: "The name of the app type.",
			},
			"ip_v6": dschema.StringAttribute{
				Optional:    true,
				Description: "The name of the app type.",
			},
			"virtual_network_id": dschema.StringAttribute{
				Optional:    true,
				Description: "The name of the app type.",
			},
			"created_after": dschema.StringAttribute{
				Optional:    true,
				Description: "A date and time after a target was created to filter on.",
			},
			"modified_after": dschema.StringAttribute{
				Optional:    true,
				Description: "A date and time after a target was modified to filter on.",
			},
			// Schema for data source is separate from resource so attributes
			// are re written here but modified to be computer aka read-only
			"targets": dschema.ListNestedAttribute{
				Computed: true,
				NestedObject: dschema.NestedAttributeObject{
					Attributes: map[string]dschema.Attribute{
						consts.AccountIDSchemaKey: dschema.StringAttribute{
							MarkdownDescription: consts.AccountIDSchemaDescription,
							Computed:            true,
						},
						consts.IDSchemaKey: schema.StringAttribute{
							MarkdownDescription: consts.IDSchemaDescription + " This is target's unique identifier.",
							Computed:            true,
						},
						"hostname": dschema.StringAttribute{
							MarkdownDescription: "A non-unique field that refers to a target.",
							Computed:            true,
						},
						"ip": schema.SingleNestedAttribute{
							MarkdownDescription: "The IPv4/IPv6 address that identifies where to reach a target.",
							Required:            true,
							Attributes: map[string]schema.Attribute{
								"ipv4": schema.SingleNestedAttribute{
									MarkdownDescription: "The target's IPv4 address.",
									Optional:            true,
									Attributes: map[string]schema.Attribute{
										"ip_addr": schema.StringAttribute{
											MarkdownDescription: "The IP address of the target.",
											Required:            true,
										},
										"virtual_network_id": schema.StringAttribute{
											MarkdownDescription: "The private virtual network identifier for the target.",
											Required:            true,
										},
									},
								},
								"ipv6": schema.SingleNestedAttribute{
									MarkdownDescription: "The target's IPv6 address.",
									Optional:            true,
									Attributes: map[string]schema.Attribute{
										"ip_addr": schema.StringAttribute{
											MarkdownDescription: "The IP address of the target.",
											Required:            true,
										},
										"virtual_network_id": schema.StringAttribute{
											MarkdownDescription: "The private virtual network identifier for the target.",
											Required:            true,
										},
									},
								},
							},
						},
						"created_at": dschema.StringAttribute{
							MarkdownDescription: "The date and time at which the target was created.",
							Computed:            true,
						},
						"modified_at": dschema.StringAttribute{
							MarkdownDescription: "The date and time at which the target was last modified.",
							Computed:            true,
						},
					},
				},
			},
		},
	}
}
