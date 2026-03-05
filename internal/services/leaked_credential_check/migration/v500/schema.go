package v500

import "github.com/hashicorp/terraform-plugin-framework/resource/schema"

func SourceLeakedCredentialCheckSchema() schema.Schema {
	return schema.Schema{
		Version: 0,
		Attributes: map[string]schema.Attribute{
			"zone_id": schema.StringAttribute{
				Required: true,
			},
			"enabled": schema.BoolAttribute{
				Required: true,
			},
		},
	}
}
