package utils

import (
	"fmt"
)

type UserAgentBuilderParams struct {
	ProviderVersion  *string
	PluginVersion    *string
	PluginType       *string
	TerraformVersion *string
	OperatorSuffix   *string
}

func (p *UserAgentBuilderParams) String() string {
	var ua string
	if p.ProviderVersion != nil {
		ua += fmt.Sprintf("terraform-provider-cloudflare/%s", *p.ProviderVersion)
	}

	if p.PluginType != nil {
		ua += fmt.Sprintf(" %s", *p.PluginType)
	}

	if p.PluginVersion != nil {
		ua += fmt.Sprintf("/%s", *p.PluginVersion)
	}

	// Operator suffix and Terraform version are mutually exclusive and we should
	// only ever see one of them.
	if p.OperatorSuffix != nil {
		ua += fmt.Sprintf(" %s", *p.OperatorSuffix)
	} else if p.TerraformVersion != nil {
		ua += fmt.Sprintf(" terraform/%s", *p.TerraformVersion)
	}

	return ua
}

func BuildUserAgent(params UserAgentBuilderParams) string {
	return params.String()
}
