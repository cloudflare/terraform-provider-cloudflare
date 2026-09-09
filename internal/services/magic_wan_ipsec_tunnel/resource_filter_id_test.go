package magic_wan_ipsec_tunnel_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/magic_wan_ipsec_tunnel"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

func TestMagicWANIPSECTunnelBGPFilterIDs(t *testing.T) {
	ctx := context.Background()
	filterID := "019fb3bf3b6a7d9f8bbf3d27d87cdf44"

	var r magic_wan_ipsec_tunnel.MagicWANIPSECTunnelResource
	resp := resource.SchemaResponse{}
	r.Schema(ctx, resource.SchemaRequest{}, &resp)
	realSchema := resp.Schema

	bgpSchema := schema.Schema{
		Attributes: map[string]schema.Attribute{
			"bgp": realSchema.Attributes["bgp"],
		},
	}

	rootObjectType := bgpSchema.Type().TerraformType(ctx).(tftypes.Object)
	bgpObjectType := rootObjectType.AttributeTypes["bgp"].(tftypes.Object)

	t.Run("sets filter IDs from config", func(t *testing.T) {
		cfg := tfsdk.Config{
			Raw: tftypes.NewValue(rootObjectType, map[string]tftypes.Value{
				"bgp": tftypes.NewValue(bgpObjectType, map[string]tftypes.Value{
					"customer_asn":     tftypes.NewValue(tftypes.Number, 65001),
					"extra_prefixes":   tftypes.NewValue(tftypes.List{ElementType: tftypes.String}, nil),
					"md5_key":          tftypes.NewValue(tftypes.String, nil),
					"import_filter_id": tftypes.NewValue(tftypes.String, filterID),
					"export_filter_id": tftypes.NewValue(tftypes.String, filterID),
				}),
			}),
			Schema: bgpSchema,
		}

		var got struct {
			BGP *magic_wan_ipsec_tunnel.MagicWANIPSECTunnelBGPModel `tfsdk:"bgp"`
		}

		diags := cfg.Get(ctx, &got)
		if diags.HasError() {
			t.Fatalf("config.Get failed: %v", diags)
		}

		if got.BGP == nil {
			t.Fatal("expected BGP block to be set")
		}
		if got.BGP.ImportFilterID.ValueString() != filterID {
			t.Errorf("import_filter_id = %q, want %q", got.BGP.ImportFilterID.ValueString(), filterID)
		}
		if got.BGP.ExportFilterID.ValueString() != filterID {
			t.Errorf("export_filter_id = %q, want %q", got.BGP.ExportFilterID.ValueString(), filterID)
		}
	})

	t.Run("nulls filter IDs when removed from config", func(t *testing.T) {
		cfg := tfsdk.Config{
			Raw: tftypes.NewValue(rootObjectType, map[string]tftypes.Value{
				"bgp": tftypes.NewValue(bgpObjectType, map[string]tftypes.Value{
					"customer_asn":     tftypes.NewValue(tftypes.Number, 65001),
					"extra_prefixes":   tftypes.NewValue(tftypes.List{ElementType: tftypes.String}, nil),
					"md5_key":          tftypes.NewValue(tftypes.String, nil),
					"import_filter_id": tftypes.NewValue(tftypes.String, nil),
					"export_filter_id": tftypes.NewValue(tftypes.String, nil),
				}),
			}),
			Schema: bgpSchema,
		}

		var got struct {
			BGP *magic_wan_ipsec_tunnel.MagicWANIPSECTunnelBGPModel `tfsdk:"bgp"`
		}

		diags := cfg.Get(ctx, &got)
		if diags.HasError() {
			t.Fatalf("config.Get failed: %v", diags)
		}

		if got.BGP == nil {
			t.Fatal("expected BGP block to be set")
		}
		if !got.BGP.ImportFilterID.IsNull() {
			t.Errorf("import_filter_id = %q, want null", got.BGP.ImportFilterID.ValueString())
		}
		if !got.BGP.ExportFilterID.IsNull() {
			t.Errorf("export_filter_id = %q, want null", got.BGP.ExportFilterID.ValueString())
		}
	})

	t.Run("omits BGP block entirely when not in config", func(t *testing.T) {
		cfg := tfsdk.Config{
			Raw: tftypes.NewValue(rootObjectType, map[string]tftypes.Value{
				"bgp": tftypes.NewValue(bgpObjectType, nil),
			}),
			Schema: bgpSchema,
		}

		var got struct {
			BGP *magic_wan_ipsec_tunnel.MagicWANIPSECTunnelBGPModel `tfsdk:"bgp"`
		}

		diags := cfg.Get(ctx, &got)
		if diags.HasError() {
			t.Fatalf("config.Get failed: %v", diags)
		}

		if got.BGP != nil {
			t.Errorf("expected BGP block to be nil, got %+v", got.BGP)
		}
	})
}
