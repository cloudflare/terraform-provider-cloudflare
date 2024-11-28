package leaked_credential_check_rules_test

import (
	"context"
	"errors"
	"fmt"
	"os"
	"regexp"
	"testing"

	cfv1 "github.com/cloudflare/cloudflare-go"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"

	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func init() {
	resource.AddTestSweepers("cloudflare_leaked_credential_check_rules", &resource.Sweeper{
		Name: "cloudflare_leaked_credential_check_rules",
		F:    testSweepCloudflareLCCRules,
	})
}

func testSweepCloudflareLCCRules(r string) error {
	ctx := context.Background()
	client, clientErr := acctest.SharedV1Client()
	if clientErr != nil {
		tflog.Error(ctx, fmt.Sprintf("Failed to create Cloudflare client: %s", clientErr))
	}

	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	if zoneID == "" {
		return errors.New("CLOUDFLARE_ZONE_ID must be set")
	}
	// fetch existing rules from API
	rules, err := client.LeakedCredentialCheckListDetections(ctx, cfv1.ZoneIdentifier(zoneID), cfv1.LeakedCredentialCheckListDetectionsParams{})
	if err != nil {
		tflog.Error(ctx, fmt.Sprintf("Error fetching Leaked Credential Check user-defined detection patterns: %s", err))
		return err
	}
	for _, rule := range rules {
		deleteParam := cfv1.LeakedCredentialCheckDeleteDetectionParams{DetectionID: rule.ID}
		_, err := client.LeakedCredentialCheckDeleteDetection(ctx, cfv1.ZoneIdentifier(zoneID), deleteParam)
		if err != nil {
			tflog.Error(ctx, fmt.Sprintf("Error deleting a user-defined detection patter for Leaked Credential Check: %s", err))
		}

	}
	return nil
}

func TestAccCloudflareLeakedCredentialCheckRules_CRUD(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("cloudflare_leaked_credential_check_rules.%s", rnd)
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")

	var detectionRules []cfv1.LeakedCredentialCheckDetectionEntry

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccConfigAddHeader(rnd, zoneID, testAccLCCTwoSimpleRules(rnd)),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "zone_id", zoneID),
					resource.TestCheckResourceAttr(name, "rule.#", "2"),
					resource.TestCheckResourceAttr(name, "rule.0.%", "3"),
					resource.TestCheckResourceAttr(name, "rule.1.%", "3"),
					resource.TestMatchTypeSetElemNestedAttrs(
						name,
						"rule.*",
						map[string]*regexp.Regexp{
							"id":       regexp.MustCompile(`^[a-z0-9]+$`),
							"username": regexp.MustCompile(`"id"`),
							"password": regexp.MustCompile(`"secret"`),
						},
					),
					resource.TestMatchTypeSetElemNestedAttrs(
						name,
						"rule.*",
						map[string]*regexp.Regexp{
							"id":       regexp.MustCompile(`^[a-z0-9]+$`),
							"username": regexp.MustCompile(`"user"`),
							"password": regexp.MustCompile(`"pass"`),
						},
					),
					testAccCheckLCCNumRules(name, 2, &detectionRules),
				),
			},
			{ // remove one rule, keep one, and add a new one
				Config: testAccConfigAddHeader(rnd, zoneID, testAccLCCSimpleChangeOneRule(rnd)),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "zone_id", zoneID),
					resource.TestCheckResourceAttr(name, "rule.#", "2"),
					resource.TestCheckResourceAttr(name, "rule.0.%", "3"),
					resource.TestCheckResourceAttr(name, "rule.1.%", "3"),
					resource.TestMatchTypeSetElemNestedAttrs(
						name,
						"rule.*",
						map[string]*regexp.Regexp{
							"id":       regexp.MustCompile(`^[a-z0-9]+$`),
							"username": regexp.MustCompile(`"name"`),
							"password": regexp.MustCompile(`"key"`),
						},
					),
					testAccCheckLCCOneRuleChange(name, 2, &detectionRules),
				),
			},
			{ // clear rules
				Config: testAccConfigAddHeader(rnd, zoneID, testAccLCCNoRules(rnd)),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "zone_id", zoneID),
					resource.TestCheckResourceAttr(name, "rule.#", "0"),
					testAccCheckLCCNumRules(name, 0, &detectionRules),
				),
			},
			{ // add a single rule
				Config: testAccConfigAddHeader(rnd, zoneID, testAccLCCOneSimpleRule(rnd)),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "zone_id", zoneID),
					resource.TestCheckResourceAttr(name, "rule.#", "1"),
					resource.TestCheckResourceAttr(name, "rule.0.%", "3"),
					resource.TestCheckResourceAttr(name, "rule.0.username", "lookup_json_string(http.request.body.raw, \"username\")"),
					resource.TestCheckResourceAttr(name, "rule.0.password", "lookup_json_string(http.request.body.raw, \"password\")"),
					testAccCheckLCCNumRules(name, 1, &detectionRules),
				),
			},
		},
	})
}

func testAccConfigAddHeader(name, zoneID, config string) string {
	header := fmt.Sprintf(`
	resource "cloudflare_leaked_credential_check" "%[1]s" {
		zone_id = "%[2]s"
		enabled = true
	}`, name, zoneID)
	return header + "\n" + config
}

func testAccLCCTwoSimpleRules(name string) string {
	return fmt.Sprintf(`
  resource "cloudflare_leaked_credential_check_rules" "%[1]s" {
    zone_id = cloudflare_leaked_credential_check.%[1]s.zone_id
	rule {
		username = "lookup_json_string(http.request.body.raw, \"user\")"
		password = "lookup_json_string(http.request.body.raw, \"pass\")"
	}
	rule {
		username = "lookup_json_string(http.request.body.raw, \"id\")"
		password = "lookup_json_string(http.request.body.raw, \"secret\")"
	}
  }`, name)
}

func testAccLCCSimpleChangeOneRule(name string) string {
	return fmt.Sprintf(`
  resource "cloudflare_leaked_credential_check_rules" "%[1]s" {
    zone_id = cloudflare_leaked_credential_check.%[1]s.zone_id
	rule {
		username = "lookup_json_string(http.request.body.raw, \"name\")"
		password = "lookup_json_string(http.request.body.raw, \"key\")"
	}
	rule {
		username = "lookup_json_string(http.request.body.raw, \"id\")"
		password = "lookup_json_string(http.request.body.raw, \"secret\")"
	}
  }`, name)
}

func testAccLCCNoRules(name string) string {
	return fmt.Sprintf(`
  resource "cloudflare_leaked_credential_check_rules" "%[1]s" {
    zone_id = cloudflare_leaked_credential_check.%[1]s.zone_id
  }`, name)
}

func testAccLCCOneSimpleRule(name string) string {
	return fmt.Sprintf(`
  resource "cloudflare_leaked_credential_check_rules" "%[1]s" {
    zone_id = cloudflare_leaked_credential_check.%[1]s.zone_id
	rule {
		username = "lookup_json_string(http.request.body.raw, \"username\")"
		password = "lookup_json_string(http.request.body.raw, \"password\")"
	}
  }`, name)
}

func testAccCheckLCCNumRules(name string, length int, rules *[]cfv1.LeakedCredentialCheckDetectionEntry) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
		_, ok := s.RootModule().Resources[name]
		if !ok {
			return fmt.Errorf("not found: %s", name)
		}
		client, _ := acctest.SharedV1Client()
		detections, err := client.LeakedCredentialCheckListDetections(
			context.Background(),
			cfv1.ZoneIdentifier(zoneID),
			cfv1.LeakedCredentialCheckListDetectionsParams{},
		)
		if err != nil {
			return err
		}

		if length != len(detections) {
			return fmt.Errorf("Expected num of rules (%d) does not match the actual num (%d)", length, len(detections))
		}
		*rules = detections

		return nil
	}
}

func testAccCheckLCCOneRuleChange(name string, length int, rules *[]cfv1.LeakedCredentialCheckDetectionEntry) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
		_, ok := s.RootModule().Resources[name]
		if !ok {
			return fmt.Errorf("not found: %s", name)
		}
		client, _ := acctest.SharedV1Client()
		detections, err := client.LeakedCredentialCheckListDetections(
			context.Background(),
			cfv1.ZoneIdentifier(zoneID),
			cfv1.LeakedCredentialCheckListDetectionsParams{},
		)
		if err != nil {
			return err
		}

		if length != len(detections) {
			return fmt.Errorf("Expected num of rules (%d) does not match the actual num (%d)", length, len(detections))
		}
		previousRulesMap := make(map[string]string)
		for _, prul := range *rules {
			mapkey := prul.Password + "|" + prul.Username
			previousRulesMap[mapkey] = prul.ID
		}

		found := false
		for _, det := range detections {
			mapkey := det.Password + "|" + det.Username
			if _, exists := previousRulesMap[mapkey]; exists {
				if previousRulesMap[mapkey] != det.ID {
					return errors.New("Found the unchanged rule but the ID is different")
				}
				found = true
			}
		}
		if !found {
			return fmt.Errorf("Could not find the rule that was not changed!")
		}
		return nil
	}
}
