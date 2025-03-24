package utils

import (
	"fmt"
)

type UserAgentBuilderParams struct {
	// Version of `terraform-provider-cloudflare`.
	ProviderVersion *string

	// Version of `terraform-plugin-*` libraries that we rely on for the internal
	// operations.
	PluginVersion *string

	// Which plugin is in use. Currently only available options are
	// `terraform-plugin-sdk` and `terraform-plugin-framework`.
	PluginType *string

	// Version of Terraform that is initiating the operation. Mutually exclusive
	// with `OperatorSuffix`.
	TerraformVersion *string

	// Customised operation suffix to append to the user agent for identifying
	// traffic. Mutually exclusive with `TerraformVersion`.
	OperatorSuffix *string
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

// BuildUserAgent takes the `UserAgentBuilderParams` and contextually builds
// a HTTP user agent for making API calls.
func BuildUserAgent(params UserAgentBuilderParams) string {
	return params.String()
}
