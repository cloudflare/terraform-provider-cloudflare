package sdkv2provider

import (
	"fmt"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceCloudflareTeamsCertificateSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		consts.AccountIDSchemaKey: {
			Description: consts.AccountIDSchemaDescription,
			Type:        schema.TypeString,
			Required:    true,
		},
		"custom": {
			Type:         schema.TypeBool,
			Optional:     true,
			Description:  "The type of certificate (custom or Gateway-managed)",
			ExactlyOneOf: []string{"custom", "gateway_managed"},
		},
		"gateway_managed": {
			Type:         schema.TypeBool,
			Optional:     true,
			Description:  "The type of certificate (custom or Gateway-managed)",
			ExactlyOneOf: []string{"custom", "gateway_managed"},
		},
		"id": {
			Type:          schema.TypeString,
			Optional:      true,
			Computed:      true,
			Description:   "Certificate UUID. Computed for Gateway-managed certificates.",
			RequiredWith:  []string{"custom"},
			ConflictsWith: []string{"gateway_managed"},
		},
		"validity_period_days": {
			Type:          schema.TypeInt,
			Optional:      true,
			Description:   "Number of days the generated certificate will be valid, minimum 1 day and maximum 30 years. Defaults to 5 years.",
			ValidateFunc:  validation.IntBetween(1, 10950),
			Default:       1826,
			RequiredWith:  []string{"gateway_managed"},
			ConflictsWith: []string{"custom"},
			ForceNew:      true,
		},
		"activate": {
			Type:        schema.TypeBool,
			Optional:    true,
			Description: "Whether or not to activate a certificate. A certificate must be activated to use in Gateway certificate settings",
			Default:     false,
		},
		"in_use": {
			Type:        schema.TypeBool,
			Computed:    true,
			Description: "Whether the certificate is in use by Gateway for TLS interception and the block page",
		},
		"binding_status": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: fmt.Sprintf("The deployment status of the certificate on the edge %s", renderAvailableDocumentationValuesStringSlice([]string{"IP", "SERIAL", "URL", "DOMAIN", "EMAIL"})),
		},
		"qs_pack_id": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"uploaded_on": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"created_at": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"expires_on": {
			Type:     schema.TypeString,
			Computed: true,
		},
	}
}
