package ai_security_custom_topics_test

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

func TestAccCloudflareAISecurityCustomTopicsDataSource_Basic(t *testing.T) {
	client, zoneID := testCustomTopicsClient(t)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	current, err := client.AISecurity.CustomTopics.Get(ctx, ai_security.CustomTopicGetParams{ZoneID: cloudflare.F(zoneID)})
	if err != nil {
		t.Fatalf("read AI Security custom topics on test zone: %v", err)
	}

	checks := []resource.TestCheckFunc{
		resource.TestCheckResourceAttr("data.cloudflare_ai_security_custom_topics.test", "id", zoneID),
		resource.TestCheckResourceAttr("data.cloudflare_ai_security_custom_topics.test", "topics.#", fmt.Sprint(len(current.Topics))),
	}
	for i, topic := range current.Topics {
		checks = append(checks,
			resource.TestCheckResourceAttr("data.cloudflare_ai_security_custom_topics.test", fmt.Sprintf("topics.%d.label", i), topic.Label),
			resource.TestCheckResourceAttr("data.cloudflare_ai_security_custom_topics.test", fmt.Sprintf("topics.%d.topic", i), topic.Topic),
		)
	}

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{{
			Config: fmt.Sprintf(`data "cloudflare_ai_security_custom_topics" "test" {
  zone_id = %q
}`, zoneID),
			Check: resource.ComposeTestCheckFunc(checks...),
		}},
	})
}
