package v500

import (
	"context"
	"fmt"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Transform converts source (legacy v4) state to target (current v5) state.
func Transform(ctx context.Context, source SourceSpectrumApplicationModel) (*TargetSpectrumApplicationModel, diag.Diagnostics) {
	var diags diag.Diagnostics

	target := &TargetSpectrumApplicationModel{
		// Direct copies (same type, same name)
		ID:               source.ID,
		ZoneID:           source.ZoneID,
		Protocol:         source.Protocol,
		ArgoSmartRouting: source.ArgoSmartRouting,
		IPFirewall:       source.IPFirewall,
		ProxyProtocol:    source.ProxyProtocol,
		TLS:              source.TLS,
		TrafficType:      source.TrafficType,
	}

	// Convert dns: array[0] → pointer to object
	if len(source.DNS) > 0 {
		target.DNS = &TargetDNSModel{
			Name: source.DNS[0].Name,
			Type: source.DNS[0].Type,
		}
	}

	// Convert origin_dns: array[0] → pointer to object
	if len(source.OriginDNS) > 0 {
		target.OriginDNS = &TargetOriginDNSModel{
			Name: source.OriginDNS[0].Name,
			TTL:  source.OriginDNS[0].TTL,
			Type: source.OriginDNS[0].Type,
		}
	}

	// Convert origin_direct: types.List → *[]types.String
	if !source.OriginDirect.IsNull() && !source.OriginDirect.IsUnknown() {
		var elements []types.String
		diags.Append(source.OriginDirect.ElementsAs(ctx, &elements, false)...)
		if diags.HasError() {
			return nil, diags
		}
		target.OriginDirect = &elements
	}

	// Convert origin_port_range → origin_port DynamicAttribute (string range)
	// OR convert origin_port integer → DynamicAttribute (number)
	if len(source.OriginPortRange) > 0 {
		start := source.OriginPortRange[0].Start.ValueInt64()
		end := source.OriginPortRange[0].End.ValueInt64()
		portRange := fmt.Sprintf("%d-%d", start, end)
		target.OriginPort = customfield.RawNormalizedDynamicValueFrom(types.StringValue(portRange))
	} else if !source.OriginPort.IsNull() && !source.OriginPort.IsUnknown() {
		target.OriginPort = customfield.RawNormalizedDynamicValueFrom(source.OriginPort)
	} else {
		target.OriginPort = customfield.RawNormalizedDynamicValue(types.DynamicNull())
	}

	// Convert edge_ips: array[0] → customfield.NestedObject
	if len(source.EdgeIPs) > 0 {
		var ips *[]types.String
		if !source.EdgeIPs[0].IPs.IsNull() && !source.EdgeIPs[0].IPs.IsUnknown() {
			var ipElements []types.String
			diags.Append(source.EdgeIPs[0].IPs.ElementsAs(ctx, &ipElements, false)...)
			if diags.HasError() {
				return nil, diags
			}
			ips = &ipElements
		}
		edgeIPsModel := &TargetEdgeIPsModel{
			Connectivity: source.EdgeIPs[0].Connectivity,
			Type:         source.EdgeIPs[0].Type,
			IPs:          ips,
		}
		target.EdgeIPs = customfield.NewObjectMust[TargetEdgeIPsModel](ctx, edgeIPsModel)
	} else {
		target.EdgeIPs = customfield.NullObject[TargetEdgeIPsModel](ctx)
	}

	// Convert timestamps: string → timetypes.RFC3339
	if !source.CreatedOn.IsNull() && !source.CreatedOn.IsUnknown() {
		createdOn, d := timetypes.NewRFC3339Value(source.CreatedOn.ValueString())
		diags.Append(d...)
		target.CreatedOn = createdOn
	} else {
		target.CreatedOn = timetypes.NewRFC3339Null()
	}
	if !source.ModifiedOn.IsNull() && !source.ModifiedOn.IsUnknown() {
		modifiedOn, d := timetypes.NewRFC3339Value(source.ModifiedOn.ValueString())
		diags.Append(d...)
		target.ModifiedOn = modifiedOn
	} else {
		target.ModifiedOn = timetypes.NewRFC3339Null()
	}

	return target, diags
}
