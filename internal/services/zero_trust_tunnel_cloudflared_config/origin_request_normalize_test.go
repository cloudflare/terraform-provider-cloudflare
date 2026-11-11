package zero_trust_tunnel_cloudflared_config

import (
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Response from BUGS-2130 (GET .../cfd_tunnel/<id>/configurations).
const emptyOriginRequestResponse = `{"result":{"tunnel_id":"b6896a3f-8ddf-45bb-9c7a-6408492744d0","version":0,"config":{"ingress":[{"service":"http://opal-secondary-opal-web.opal:80","hostname":"staging.iga.sso.cloudflare.dev","originRequest":{}},{"service":"http_status:404"}],"warp-routing":{},"originRequest":{}},"source":"cloudflare","created_at":"2025-01-09T22:48:43.152743Z"}}`

func decodeWithPrior(t *testing.T, prior *ZeroTrustTunnelCloudflaredConfigModel) *ZeroTrustTunnelCloudflaredConfigModel {
	t.Helper()
	presence := captureOriginRequestPresence(prior)
	env := ZeroTrustTunnelCloudflaredConfigResultEnvelope{*prior}
	if err := apijson.Unmarshal([]byte(emptyOriginRequestResponse), &env); err != nil {
		t.Fatal(err)
	}
	data := &env.Result
	presence.normalize(data)
	return data
}

func TestEmptyOriginRequestDroppedWhenPriorHadNone(t *testing.T) {
	t.Parallel()
	prior := &ZeroTrustTunnelCloudflaredConfigModel{
		TunnelID:  types.StringValue("b6896a3f-8ddf-45bb-9c7a-6408492744d0"),
		AccountID: types.StringValue("acc"),
		Config: &ZeroTrustTunnelCloudflaredConfigConfigModel{
			Ingress: &[]*ZeroTrustTunnelCloudflaredConfigConfigIngressModel{
				{Service: types.StringValue("http://opal-secondary-opal-web.opal:80"), Hostname: types.StringValue("staging.iga.sso.cloudflare.dev")},
				{Service: types.StringValue("http_status:404")},
			},
		},
	}
	data := decodeWithPrior(t, prior)
	if data.Config.OriginRequest != nil {
		t.Errorf("config.origin_request = %+v, want nil", data.Config.OriginRequest)
	}
	if ing := (*data.Config.Ingress)[0]; ing.OriginRequest != nil {
		t.Errorf("ingress[0].origin_request = %+v, want nil", ing.OriginRequest)
	}
}

func TestExplicitEmptyOriginRequestKept(t *testing.T) {
	t.Parallel()
	prior := &ZeroTrustTunnelCloudflaredConfigModel{
		TunnelID:  types.StringValue("b6896a3f-8ddf-45bb-9c7a-6408492744d0"),
		AccountID: types.StringValue("acc"),
		Config: &ZeroTrustTunnelCloudflaredConfigConfigModel{
			OriginRequest: &ZeroTrustTunnelCloudflaredConfigConfigOriginRequestModel{},
			Ingress: &[]*ZeroTrustTunnelCloudflaredConfigConfigIngressModel{
				{Service: types.StringValue("http://opal-secondary-opal-web.opal:80"), OriginRequest: &ZeroTrustTunnelCloudflaredConfigConfigIngressOriginRequestModel{}},
				{Service: types.StringValue("http_status:404")},
			},
		},
	}
	data := decodeWithPrior(t, prior)
	if data.Config.OriginRequest == nil {
		t.Error("config.origin_request was dropped, want the explicit {} kept")
	}
	if (*data.Config.Ingress)[0].OriginRequest == nil {
		t.Error("ingress[0].origin_request was dropped, want the explicit {} kept")
	}
}

func TestNonEmptyOriginRequestKept(t *testing.T) {
	t.Parallel()
	o := &ZeroTrustTunnelCloudflaredConfigModel{Config: &ZeroTrustTunnelCloudflaredConfigConfigModel{
		OriginRequest: &ZeroTrustTunnelCloudflaredConfigConfigOriginRequestModel{NoTLSVerify: types.BoolValue(true)},
	}}
	originRequestPresence{topWasNil: true}.normalize(o)
	if o.Config.OriginRequest == nil {
		t.Error("non-empty origin_request was dropped")
	}
}
