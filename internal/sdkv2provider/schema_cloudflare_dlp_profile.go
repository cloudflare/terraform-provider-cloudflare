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

// Custom hash function used on DLP entries. Extracts the "name" property
// to provide a stable hash for profile entries and prevent spurious differences
// between the state/infra.
func hashResourceCloudflareDLPEntry(i interface{}) int {
	v := i.(map[string]interface{})
	return schema.HashString(v["name"])
}

func resourceCloudflareDLPContextAwarenessSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"enabled": {
			Type:        schema.TypeBool,
			Required:    true,
			Description: "Scan the context of predefined entries to only return matches surrounded by keywords.",
		},
		"skip": {
			Type:        schema.TypeList,
			Description: "Content types to exclude from context analysis and return all matches.",
			Required:    true,
			MaxItems:    1,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"files": {
						Type:        schema.TypeBool,
						Required:    true,
						Description: "Return all matches, regardless of context analysis result, if the data is a file.",
					},
				},
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
			Set: hashResourceCloudflareDLPEntry,
		},
		"allowed_match_count": {
			Type:         schema.TypeInt,
			Description:  "Related DLP policies will trigger when the match count exceeds the number set.",
			Required:     true,
			ValidateFunc: validation.IntBetween(0, 1000),
		},
		"context_awareness": {
			Type:        schema.TypeList,
			Description: "Scan the context of predefined entries to only return matches surrounded by keywords.",
			Computed:    true,
			Optional:    true,
			MaxItems:    1,
			Elem: &schema.Resource{
				Schema: resourceCloudflareDLPContextAwarenessSchema(),
			},
		},
	}
}
