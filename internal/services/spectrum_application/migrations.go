// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package spectrum_application

import (
	"context"
	"fmt"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ resource.ResourceWithUpgradeState = (*SpectrumApplicationResource)(nil)

func (r *SpectrumApplicationResource) UpgradeState(ctx context.Context) map[int64]resource.StateUpgrader {
	return map[int64]resource.StateUpgrader{
		0: {
			// PriorSchema must work with BOTH v4 and v5 states (before migration)
			// Include ALL fields that exist in either version
			PriorSchema: &schema.Schema{
				Attributes: map[string]schema.Attribute{
					"id": schema.StringAttribute{
						Optional: true, // v4 allowed optional ID, v5 makes it computed-only
						Computed: true,
					},
					"zone_id": schema.StringAttribute{
						Required: true,
					},
					"protocol": schema.StringAttribute{
						Required: true,
					},
					"origin_direct": schema.ListAttribute{
						Optional:    true,
						ElementType: types.StringType,
					},
					"tls": schema.StringAttribute{
						Optional: true,
						Computed: true,
					},
					"ip_firewall": schema.BoolAttribute{
						Optional: true,
						Computed: true,
					},
					"proxy_protocol": schema.StringAttribute{
						Optional: true,
						Computed: true,
					},
					"argo_smart_routing": schema.BoolAttribute{
						Optional: true,
						Computed: true,
					},
					"traffic_type": schema.StringAttribute{
						Optional: true,
						Computed: true,
					},
					// For block-to-attribute conversions, use ListNestedAttribute
					// to handle the v4 block structure
					"dns": schema.ListNestedAttribute{
						Optional: true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"name": schema.StringAttribute{
									Optional: true,
								},
								"type": schema.StringAttribute{
									Optional: true,
								},
							},
						},
					},
					"edge_ips": schema.ListNestedAttribute{
						Optional: true,
						Computed: true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"connectivity": schema.StringAttribute{
									Optional: true,
								},
								"type": schema.StringAttribute{
									Optional: true,
								},
								"ips": schema.SetAttribute{
									Optional:    true,
									ElementType: types.StringType,
								},
							},
						},
					},
					"origin_dns": schema.ListNestedAttribute{
						Optional: true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"name": schema.StringAttribute{
									Optional: true,
								},
								"ttl": schema.Int64Attribute{
									Optional: true,
								},
								"type": schema.StringAttribute{
									Optional: true,
								},
							},
						},
					},
					// Handle both origin_port (v4 int) and origin_port_range (v4 block)
					"origin_port": schema.Int64Attribute{
						Optional: true,
					},
					"origin_port_range": schema.ListNestedAttribute{
						Optional: true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"start": schema.Int64Attribute{
									Optional: true,
								},
								"end": schema.Int64Attribute{
									Optional: true,
								},
							},
						},
					},
				},
			},

			StateUpgrader: func(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
				var priorStateData struct {
					ID               types.String `tfsdk:"id"`
					ZoneID           types.String `tfsdk:"zone_id"`
					Protocol         types.String `tfsdk:"protocol"`
					OriginDirect     []types.String `tfsdk:"origin_direct"`
					TLS              types.String `tfsdk:"tls"`
					IPFirewall       types.Bool   `tfsdk:"ip_firewall"`
					ProxyProtocol    types.String `tfsdk:"proxy_protocol"`
					ArgoSmartRouting types.Bool   `tfsdk:"argo_smart_routing"`
					TrafficType      types.String `tfsdk:"traffic_type"`
					OriginPort       types.Int64  `tfsdk:"origin_port"`

					// Block structures from v4
					DNS []struct {
						Name types.String `tfsdk:"name"`
						Type types.String `tfsdk:"type"`
					} `tfsdk:"dns"`

					EdgeIPs []struct {
						Connectivity types.String   `tfsdk:"connectivity"`
						Type         types.String   `tfsdk:"type"`
						IPs          []types.String `tfsdk:"ips"`
					} `tfsdk:"edge_ips"`

					OriginDNS []struct {
						Name types.String `tfsdk:"name"`
						TTL  types.Int64  `tfsdk:"ttl"`
						Type types.String `tfsdk:"type"`
					} `tfsdk:"origin_dns"`

					OriginPortRange []struct {
						Start types.Int64 `tfsdk:"start"`
						End   types.Int64 `tfsdk:"end"`
					} `tfsdk:"origin_port_range"`
				}

				resp.Diagnostics.Append(req.State.Get(ctx, &priorStateData)...)
				if resp.Diagnostics.HasError() {
					return
				}

				// Initialize new state with the target model
				newState := SpectrumApplicationModel{
					ID:               priorStateData.ID,
					ZoneID:           priorStateData.ZoneID,
					Protocol:         priorStateData.Protocol,
					TLS:              priorStateData.TLS,
					IPFirewall:       priorStateData.IPFirewall,
					ProxyProtocol:    priorStateData.ProxyProtocol,
					ArgoSmartRouting: priorStateData.ArgoSmartRouting,
					TrafficType:      priorStateData.TrafficType,
				}

				// Handle origin_direct list
				if len(priorStateData.OriginDirect) > 0 {
					newState.OriginDirect = &priorStateData.OriginDirect
				}

				// Convert dns block (MaxItems:1) to object
				if len(priorStateData.DNS) > 0 {
					newState.DNS = &SpectrumApplicationDNSModel{
						Name: priorStateData.DNS[0].Name,
						Type: priorStateData.DNS[0].Type,
					}
				}

				// Convert edge_ips block (MaxItems:1) to object
				if len(priorStateData.EdgeIPs) > 0 {
					edgeIPsModel := &SpectrumApplicationEdgeIPsModel{
						Connectivity: priorStateData.EdgeIPs[0].Connectivity,
						Type:         priorStateData.EdgeIPs[0].Type,
					}
					
					// Convert ips from set to list format
					if len(priorStateData.EdgeIPs[0].IPs) > 0 {
						edgeIPsModel.IPs = &priorStateData.EdgeIPs[0].IPs
					}
					
					edgeIPsValue, diags := customfield.NewObject(ctx, edgeIPsModel)
					resp.Diagnostics.Append(diags...)
					if diags.HasError() {
						return
					}
					newState.EdgeIPs = customfield.NestedObject[SpectrumApplicationEdgeIPsModel](edgeIPsValue)
				}

				// Convert origin_dns block (MaxItems:1) to object
				if len(priorStateData.OriginDNS) > 0 {
					newState.OriginDNS = &SpectrumApplicationOriginDNSModel{
						Name: priorStateData.OriginDNS[0].Name,
						TTL:  priorStateData.OriginDNS[0].TTL,
						Type: priorStateData.OriginDNS[0].Type,
					}
				}

				// Handle origin_port and origin_port_range consolidation
				if len(priorStateData.OriginPortRange) > 0 {
					// Convert origin_port_range block to origin_port string format
					start := priorStateData.OriginPortRange[0].Start
					end := priorStateData.OriginPortRange[0].End
					
					if !start.IsNull() && !end.IsNull() {
						rangeStr := fmt.Sprintf("%d-%d", start.ValueInt64(), end.ValueInt64())
						newState.OriginPort = customfield.RawNormalizedDynamicValueFrom(types.StringValue(rangeStr))
					}
				} else if !priorStateData.OriginPort.IsNull() {
					// Keep existing origin_port as integer
					newState.OriginPort = customfield.RawNormalizedDynamicValueFrom(priorStateData.OriginPort)
				}

				// Set the upgraded state
				resp.Diagnostics.Append(resp.State.Set(ctx, newState)...)
			},
		},
	}
}
