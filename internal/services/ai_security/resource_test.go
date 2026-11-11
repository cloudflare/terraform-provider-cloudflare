package ai_security_test

import (
	"context"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/cloudflare/cloudflare-go/v7"
	"github.com/cloudflare/cloudflare-go/v7/ai_security"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func testAISecurityClient(t *testing.T) (*cloudflare.Client, string) {
	t.Helper()
	if os.Getenv("TF_ACC") == "" {
		t.Skip("set TF_ACC=1 to run acceptance tests")
	}
	acctest.TestAccPreCheck(t)
	acctest.TestAccPreCheck_ZoneID(t)
	return cloudflare.NewClient(), os.Getenv("CLOUDFLARE_ZONE_ID")
}

func TestAccCloudflareAISecurity_Basic(t *testing.T) {
	client, zoneID := testAISecurityClient(t)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	initial, err := client.AISecurity.Get(ctx, ai_security.AISecurityGetParams{ZoneID: cloudflare.F(zoneID)})
	cancel()
	if err != nil {
		t.Fatalf("read original AI Security setting: %v", err)
	}

	// Terraform destroy only removes state, so restore the zone even if a test step fails.
	t.Cleanup(func() {
		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()
		_, err := client.AISecurity.Update(ctx, ai_security.AISecurityUpdateParams{
			ZoneID:  cloudflare.F(zoneID),
			Enabled: cloudflare.F(initial.Enabled),
		})
		if err != nil {
			t.Errorf("restore AI Security setting on test zone: %v", err)
			return
		}
		current, err := client.AISecurity.Get(ctx, ai_security.AISecurityGetParams{ZoneID: cloudflare.F(zoneID)})
		if err != nil || current.Enabled != initial.Enabled {
			t.Errorf("verify restored AI Security setting: got %v, error %v, want enabled=%t", current, err, initial.Enabled)
		}
	})

	config := func(enabled bool) string {
		return fmt.Sprintf(`
resource "cloudflare_ai_security" "test" {
  zone_id = %q
  enabled = %t
}

data "cloudflare_ai_security" "current" {
  zone_id = cloudflare_ai_security.test.zone_id
  depends_on = [cloudflare_ai_security.test]
}
`, zoneID, enabled)
	}
	check := func(enabled bool) resource.TestCheckFunc {
		return resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr("cloudflare_ai_security.test", "id", zoneID),
			resource.TestCheckResourceAttr("cloudflare_ai_security.test", "zone_id", zoneID),
			resource.TestCheckResourceAttr("cloudflare_ai_security.test", "enabled", fmt.Sprint(enabled)),
			resource.TestCheckResourceAttr("data.cloudflare_ai_security.current", "enabled", fmt.Sprint(enabled)),
		)
	}

	// Leave the API in the changed state until destroy, so CheckDestroy can
	// distinguish state removal from an unintended change to the zone setting.
	wantAfterDestroy := !initial.Enabled
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy: func(_ *terraform.State) error {
			ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
			defer cancel()
			current, err := client.AISecurity.Get(ctx, ai_security.AISecurityGetParams{ZoneID: cloudflare.F(zoneID)})
			if err != nil {
				return fmt.Errorf("read AI Security setting after Terraform destroy: %w", err)
			}
			if current.Enabled != wantAfterDestroy {
				return fmt.Errorf("AI Security setting changed after Terraform destroy: got enabled=%t, want %t", current.Enabled, wantAfterDestroy)
			}
			return nil
		},
		Steps: []resource.TestStep{
			{Config: config(initial.Enabled), Check: check(initial.Enabled)},
			{
				Config: config(wantAfterDestroy),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction("cloudflare_ai_security.test", plancheck.ResourceActionUpdate),
					},
				},
				Check: check(wantAfterDestroy),
			},
			{RefreshState: true, Check: check(wantAfterDestroy)},
			{
				ResourceName:      "cloudflare_ai_security.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}
