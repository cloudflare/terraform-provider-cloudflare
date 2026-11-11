package ai_security_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/cloudflare/cloudflare-go/v7"
	"github.com/cloudflare/cloudflare-go/v7/ai_security"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccCloudflareAISecurityDataSource_Basic(t *testing.T) {
	client, zoneID := testAISecurityClient(t)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	current, err := client.AISecurity.Get(ctx, ai_security.AISecurityGetParams{ZoneID: cloudflare.F(zoneID)})
	if err != nil {
		t.Fatalf("read AI Security setting on test zone: %v", err)
	}

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{{
			Config: fmt.Sprintf(`data "cloudflare_ai_security" "test" {
  zone_id = %q
}`, zoneID),
			Check: resource.ComposeTestCheckFunc(
				resource.TestCheckResourceAttr("data.cloudflare_ai_security.test", "id", zoneID),
				resource.TestCheckResourceAttr("data.cloudflare_ai_security.test", "enabled", fmt.Sprint(current.Enabled)),
			),
		}},
	})
}
