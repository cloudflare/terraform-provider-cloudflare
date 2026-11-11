package ai_security_custom_topics_test

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/cloudflare/cloudflare-go/v7"
	"github.com/cloudflare/cloudflare-go/v7/ai_security"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func testCustomTopicsClient(t *testing.T) (*cloudflare.Client, string) {
	t.Helper()
	if os.Getenv("TF_ACC") == "" {
		t.Skip("set TF_ACC=1 to run acceptance tests")
	}
	acctest.TestAccPreCheck(t)
	acctest.TestAccPreCheck_ZoneID(t)
	return cloudflare.NewClient(), os.Getenv("CLOUDFLARE_ZONE_ID")
}

func TestAccCloudflareAISecurityCustomTopics_Basic(t *testing.T) {
	client, zoneID := testCustomTopicsClient(t)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	initial, err := client.AISecurity.CustomTopics.Get(ctx, ai_security.CustomTopicGetParams{ZoneID: cloudflare.F(zoneID)})
	cancel()
	if err != nil {
		t.Fatalf("read original custom topics: %v", err)
	}
	// PUT replaces every topic on the zone. Never overwrite another user's list.
	if len(initial.Topics) != 0 {
		t.Fatal("AI Security custom topics acceptance test requires an empty custom topic list on the test zone")
	}

	// Terraform destroy only removes state, so clear the test topics even if an
	// intermediate apply or assertion fails.
	t.Cleanup(func() {
		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()
		_, err := client.AISecurity.CustomTopics.Update(ctx, ai_security.CustomTopicUpdateParams{
			ZoneID: cloudflare.F(zoneID),
			Topics: cloudflare.F([]ai_security.CustomTopicUpdateParamsTopic{}),
		})
		if err != nil {
			t.Errorf("clear AI Security custom topics on test zone: %v", err)
			return
		}
		current, err := client.AISecurity.CustomTopics.Get(ctx, ai_security.CustomTopicGetParams{ZoneID: cloudflare.F(zoneID)})
		if err != nil || len(current.Topics) != 0 {
			t.Errorf("verify custom topics cleared: got %v, error %v", current, err)
		}
	})

	suffix := strconv.FormatInt(time.Now().UnixNano(), 36)
	firstLabel, secondLabel := "one-"+suffix, "two-"+suffix
	first := fmt.Sprintf(`{ label = %q, topic = "first test topic" }`, firstLabel)
	updated := fmt.Sprintf(`{ label = %q, topic = "updated test topic" }`, firstLabel)
	second := fmt.Sprintf(`{ label = %q, topic = "second test topic" }`, secondLabel)
	config := func(topics string) string {
		return fmt.Sprintf(`
resource "cloudflare_ai_security_custom_topics" "test" {
  zone_id = %q
  topics  = [%s]
}

data "cloudflare_ai_security_custom_topics" "current" {
  zone_id = cloudflare_ai_security_custom_topics.test.zone_id
  depends_on = [cloudflare_ai_security_custom_topics.test]
}
`, zoneID, topics)
	}
	check := func(labels, descriptions []string) resource.TestCheckFunc {
		checks := []resource.TestCheckFunc{
			resource.TestCheckResourceAttr("cloudflare_ai_security_custom_topics.test", "id", zoneID),
			resource.TestCheckResourceAttr("cloudflare_ai_security_custom_topics.test", "zone_id", zoneID),
		}
		for _, name := range []string{"cloudflare_ai_security_custom_topics.test", "data.cloudflare_ai_security_custom_topics.current"} {
			checks = append(checks, resource.TestCheckResourceAttr(name, "topics.#", strconv.Itoa(len(labels))))
			for i := range labels {
				checks = append(checks,
					resource.TestCheckResourceAttr(name, fmt.Sprintf("topics.%d.label", i), labels[i]),
					resource.TestCheckResourceAttr(name, fmt.Sprintf("topics.%d.topic", i), descriptions[i]),
				)
			}
		}
		return resource.ComposeTestCheckFunc(checks...)
	}

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy: func(_ *terraform.State) error {
			ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
			defer cancel()
			current, err := client.AISecurity.CustomTopics.Get(ctx, ai_security.CustomTopicGetParams{ZoneID: cloudflare.F(zoneID)})
			if err != nil {
				return fmt.Errorf("read custom topics after Terraform destroy: %w", err)
			}
			if len(current.Topics) != 1 || current.Topics[0].Label != firstLabel || current.Topics[0].Topic != "updated test topic" {
				return fmt.Errorf("Terraform destroy changed zone custom topics: got %+v", current.Topics)
			}
			return nil
		},
		Steps: []resource.TestStep{
			{Config: config(first), Check: check([]string{firstLabel}, []string{"first test topic"})},
			{
				Config: config(first + ", " + second),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction("cloudflare_ai_security_custom_topics.test", plancheck.ResourceActionUpdate),
					},
				},
				Check: check([]string{firstLabel, secondLabel}, []string{"first test topic", "second test topic"}),
			},
			{
				Config: config(updated),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction("cloudflare_ai_security_custom_topics.test", plancheck.ResourceActionUpdate),
					},
				},
				Check: check([]string{firstLabel}, []string{"updated test topic"}),
			},
			{
				Config: config(""),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction("cloudflare_ai_security_custom_topics.test", plancheck.ResourceActionUpdate),
					},
				},
				Check: check(nil, nil),
			},
			{
				Config: config(""),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectEmptyPlan(),
					},
				},
			},
			{
				Config: config(updated),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction("cloudflare_ai_security_custom_topics.test", plancheck.ResourceActionUpdate),
					},
				},
				Check: check([]string{firstLabel}, []string{"updated test topic"}),
			},
			{RefreshState: true, Check: check([]string{firstLabel}, []string{"updated test topic"})},
			{
				ResourceName:      "cloudflare_ai_security_custom_topics.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}
