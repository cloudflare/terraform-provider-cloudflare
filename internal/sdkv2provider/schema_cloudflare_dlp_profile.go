package sdkv2provider

import (
	"fmt"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

const (
	DLPProfileTypeCustom     = "custom"
	DLPProfileTypePredefined = "predefined"
)

func resourceCloudflareDLPPatternSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"regex": {
			Description: "The regex that defines the pattern.",
			Type:        schema.TypeString,
			Required:    true,
		},
		"validation": {
			Description: "The validation algorithm to apply with this pattern.",
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
}

func resourceCloudflareDLPEntrySchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"id": {
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "Unique entry identifier.",
		},
		"name": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "Name of the entry to deploy.",
		},
		"enabled": {
			Type:        schema.TypeBool,
			Optional:    true,
			Default:     false,
			Description: "Whether the entry is active.",
		},
		"pattern": {
			Type:     schema.TypeList,
			MaxItems: 1,
			Optional: true,
			Elem: &schema.Resource{
				Schema: resourceCloudflareDLPPatternSchema(),
			},
		},
	}
}

func resourceCloudflareDLPProfileSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		consts.AccountIDSchemaKey: {
			Description: consts.AccountIDSchemaDescription,
			Type:        schema.TypeString,
			Required:    true,
			ForceNew:    true,
		},
		"name": {
			Description: "Name of the profile.",
			Type:        schema.TypeString,
			Required:    true,
			ForceNew:    true,
		},
		"description": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Brief summary of the profile and its intended use.",
		},
		"type": {
			Type:         schema.TypeString,
			Required:     true,
			ForceNew:     true,
			ValidateFunc: validation.StringInSlice([]string{DLPProfileTypeCustom, DLPProfileTypePredefined}, false),
			Description:  fmt.Sprintf("The type of the profile. %s", renderAvailableDocumentationValuesStringSlice([]string{DLPProfileTypeCustom, DLPProfileTypePredefined})),
		},
		"entry": {
			Type:        schema.TypeSet,
			Description: "List of entries to apply to the profile.",
			Required:    true,
			Elem: &schema.Resource{
				Schema: resourceCloudflareDLPEntrySchema(),
			},
		},
		"allowed_match_count": {
			Type:         schema.TypeInt,
			Description:  "Related DLP policies will trigger when the match count exceeds the number set.",
			Required:     true,
			ValidateFunc: validation.IntBetween(0, 1000),
		},
	}
}
