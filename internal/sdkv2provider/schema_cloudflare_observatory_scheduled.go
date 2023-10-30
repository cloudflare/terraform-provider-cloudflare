package sdkv2provider

import (
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
)

var observatoryTestScheduleFrequency = []string{
	"DAILY",
	"WEEKLY",
}

var observatoryTestRegions = []string{
	"us-central1",
	"us-east1",
	"us-east4",
	"us-south1",
	"us-west1",
	"southamerica-east1",
	"europe-north1",
	"europe-southwest1",
	"europe-west1",
	"europe-west2",
	"europe-west3",
	"europe-west4",
	"europe-west8",
	"europe-west9",
	"asia-east1",
	"asia-south1",
	"asia-southeast1",
	"me-west1",
	"australia-southeast1",
}

func resourceCloudflareObservatoryScheduledTestSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		consts.ZoneIDSchemaKey: {
			Description: consts.ZoneIDSchemaDescription,
			Type:        schema.TypeString,
			Required:    true,
			ForceNew:    true,
		},
		"url": {
			Description: "The page to run the test on.",
			Type:        schema.TypeString,
			Required:    true,
			ForceNew:    true,
			// ignore trailing "/" as observatory automatically adds it to tests
			DiffSuppressFunc: func(k, oldValue, newValue string, d *schema.ResourceData) bool {
				return strings.TrimSuffix(oldValue, "/") == newValue
			},
			DiffSuppressOnRefresh: true,
		},
		"region": {
			Description:  fmt.Sprintf("The region to run the test in. %s", renderAvailableDocumentationValuesStringSlice(observatoryTestRegions)),
			Type:         schema.TypeString,
			Required:     true,
			ValidateFunc: validation.StringInSlice(observatoryTestRegions, false),
			ForceNew:     true,
		},
		"frequency": {
			Description:  fmt.Sprintf("The frequency to run the test. %s", renderAvailableDocumentationValuesStringSlice(observatoryTestScheduleFrequency)),
			Type:         schema.TypeString,
			Required:     true,
			ValidateFunc: validation.StringInSlice(observatoryTestScheduleFrequency, false),
			ForceNew:     true,
		},
	}
}
