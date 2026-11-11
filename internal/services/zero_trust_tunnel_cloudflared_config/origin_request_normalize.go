package zero_trust_tunnel_cloudflared_config

// The API returns "originRequest": {} for configs that never set it (common for
// tunnels first configured by the v4 provider or the dashboard). Decoded as-is,
// that becomes a non-nil origin_request with every field null, so state holds {}
// while the config has nothing, and every plan shows `origin_request = {} -> null`
// (BUGS-2130). The v500 state upgrader already normalizes this on migration; this
// does the same after each API response, but only where the previous value
// (plan or state) had no origin_request, so an explicit `origin_request = {}`
// in config is kept.

type originRequestPresence struct {
	topWasNil     bool
	ingressWasNil []bool
}

func captureOriginRequestPresence(m *ZeroTrustTunnelCloudflaredConfigModel) originRequestPresence {
	p := originRequestPresence{topWasNil: true}
	if m == nil || m.Config == nil {
		return p
	}
	p.topWasNil = m.Config.OriginRequest == nil
	if m.Config.Ingress != nil {
		for _, ing := range *m.Config.Ingress {
			p.ingressWasNil = append(p.ingressWasNil, ing == nil || ing.OriginRequest == nil)
		}
	}
	return p
}

func (p originRequestPresence) normalize(m *ZeroTrustTunnelCloudflaredConfigModel) {
	if m == nil || m.Config == nil {
		return
	}
	if p.topWasNil && m.Config.OriginRequest != nil && configOriginRequestAllNull(m.Config.OriginRequest) {
		m.Config.OriginRequest = nil
	}
	if m.Config.Ingress == nil {
		return
	}
	for i, ing := range *m.Config.Ingress {
		if ing == nil || ing.OriginRequest == nil || i >= len(p.ingressWasNil) || !p.ingressWasNil[i] {
			continue
		}
		if ingressOriginRequestAllNull(ing.OriginRequest) {
			ing.OriginRequest = nil
		}
	}
}

func configOriginRequestAllNull(o *ZeroTrustTunnelCloudflaredConfigConfigOriginRequestModel) bool {
	return o.Access == nil &&
		o.CAPool.IsNull() &&
		o.ConnectTimeout.IsNull() &&
		o.DisableChunkedEncoding.IsNull() &&
		o.HTTP2Origin.IsNull() &&
		o.HTTPHostHeader.IsNull() &&
		o.KeepAliveConnections.IsNull() &&
		o.KeepAliveTimeout.IsNull() &&
		o.MatchSnItoHost.IsNull() &&
		o.NoHappyEyeballs.IsNull() &&
		o.NoTLSVerify.IsNull() &&
		o.OriginServerName.IsNull() &&
		o.ProxyType.IsNull() &&
		o.TCPKeepAlive.IsNull() &&
		o.TLSTimeout.IsNull()
}

func ingressOriginRequestAllNull(o *ZeroTrustTunnelCloudflaredConfigConfigIngressOriginRequestModel) bool {
	return o.Access == nil &&
		o.CAPool.IsNull() &&
		o.ConnectTimeout.IsNull() &&
		o.DisableChunkedEncoding.IsNull() &&
		o.HTTP2Origin.IsNull() &&
		o.HTTPHostHeader.IsNull() &&
		o.KeepAliveConnections.IsNull() &&
		o.KeepAliveTimeout.IsNull() &&
		o.MatchSnItoHost.IsNull() &&
		o.NoHappyEyeballs.IsNull() &&
		o.NoTLSVerify.IsNull() &&
		o.OriginServerName.IsNull() &&
		o.ProxyType.IsNull() &&
		o.TCPKeepAlive.IsNull() &&
		o.TLSTimeout.IsNull()
}
