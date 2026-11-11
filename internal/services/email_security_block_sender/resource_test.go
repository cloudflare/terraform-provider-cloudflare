package email_security_block_sender_test

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go/v7"
	"github.com/cloudflare/cloudflare-go/v7/email_security"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

func TestMain(m *testing.M) {
	resource.TestMain(m)
}

func init() {
	resource.AddTestSweepers("cloudflare_email_security_block_sender", &resource.Sweeper{
		Name: "cloudflare_email_security_block_sender",
		F:    testSweepCloudflareEmailSecurityBlockSender,
	})
}

func testSweepCloudflareEmailSecurityBlockSender(r string) error {
	ctx := context.Background()
	client := acctest.SharedClient()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	iter := client.EmailSecurity.Settings.BlockSenders.ListAutoPaging(ctx, email_security.SettingBlockSenderListParams{
		AccountID: cloudflare.F(accountID),
	})

	for iter.Next() {
		sender := iter.Current()

		// IP patterns cannot carry the test prefix, so fall back to the comment,
		// which every fixture in this package sets to the random resource name.
		if !utils.ShouldSweepResource(sender.Pattern) && !utils.ShouldSweepResource(sender.Comments) {
			continue
		}

		tflog.Info(ctx, fmt.Sprintf("Deleting Email Security blocked sender: %s (%s)", sender.ID, sender.Pattern))
		_, err := client.EmailSecurity.Settings.BlockSenders.Delete(ctx, sender.ID, email_security.SettingBlockSenderDeleteParams{
			AccountID: cloudflare.F(accountID),
		})
		if err != nil {
			tflog.Error(ctx, fmt.Sprintf("Failed to delete Email Security blocked sender %s: %s", sender.ID, err))
		}
	}

	if err := iter.Err(); err != nil {
		tflog.Error(ctx, fmt.Sprintf("Failed to list Email Security blocked senders: %s", err))
		return err
	}

	return nil
}

// TestAccCloudflareEmailSecurityBlockSender_Basic covers create, read back and
// import of an EMAIL pattern.
func TestAccCloudflareEmailSecurityBlockSender_Basic(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	resourceName := fmt.Sprintf("cloudflare_email_security_block_sender.%s", rnd)
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareEmailSecurityBlockSenderDestroy,
		Steps: []resource.TestStep{
			{
				Config: acctest.LoadTestCase("blocksender_email.tf", rnd, accountID),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("pattern"), knownvalue.StringExact(fmt.Sprintf("%s@example.com", rnd))),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("pattern_type"), knownvalue.StringExact("EMAIL")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("is_regex"), knownvalue.Bool(false)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("comments"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("id"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("created_at"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("modified_at"), knownvalue.NotNull()),
				},
			},
			{
				ResourceName:        resourceName,
				ImportState:         true,
				ImportStateVerify:   true,
				ImportStateIdPrefix: fmt.Sprintf("%s/", accountID),
			},
		},
	})
}

// TestAccCloudflareEmailSecurityBlockSender_Update changes pattern and comments
// and asserts the update happens in place: only account_id forces replacement,
// so the identifier must survive.
func TestAccCloudflareEmailSecurityBlockSender_Update(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	resourceName := fmt.Sprintf("cloudflare_email_security_block_sender.%s", rnd)
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareEmailSecurityBlockSenderDestroy,
		Steps: []resource.TestStep{
			{
				Config: acctest.LoadTestCase("blocksender_email.tf", rnd, accountID),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("pattern"), knownvalue.StringExact(fmt.Sprintf("%s@example.com", rnd))),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("comments"), knownvalue.StringExact(rnd)),
				},
			},
			{
				Config: acctest.LoadTestCase("blocksender_email_updated.tf", rnd, accountID),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionUpdate),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("pattern"), knownvalue.StringExact(fmt.Sprintf("%s_updated@example.com", rnd))),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("comments"), knownvalue.StringExact(fmt.Sprintf("%s updated", rnd))),
				},
			},
		},
	})
}

// TestAccCloudflareEmailSecurityBlockSender_Domain covers the DOMAIN pattern
// type and omits the optional comments attribute.
func TestAccCloudflareEmailSecurityBlockSender_Domain(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	resourceName := fmt.Sprintf("cloudflare_email_security_block_sender.%s", rnd)
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareEmailSecurityBlockSenderDestroy,
		Steps: []resource.TestStep{
			{
				Config: acctest.LoadTestCase("blocksender_domain.tf", rnd, accountID),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("pattern"), knownvalue.StringExact(fmt.Sprintf("%s.example.com", rnd))),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("pattern_type"), knownvalue.StringExact("DOMAIN")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("comments"), knownvalue.Null()),
				},
			},
			{
				ResourceName:        resourceName,
				ImportState:         true,
				ImportStateVerify:   true,
				ImportStateIdPrefix: fmt.Sprintf("%s/", accountID),
			},
		},
	})
}

// TestAccCloudflareEmailSecurityBlockSender_IP covers the IP pattern type with a
// CIDR block. The API rejects loopback, broadcast, link-local, unspecified and
// private addresses, so the fixture uses the RFC 5737 documentation range.
func TestAccCloudflareEmailSecurityBlockSender_IP(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	resourceName := fmt.Sprintf("cloudflare_email_security_block_sender.%s", rnd)
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareEmailSecurityBlockSenderDestroy,
		Steps: []resource.TestStep{
			{
				Config: acctest.LoadTestCase("blocksender_ip.tf", rnd, accountID),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("pattern"), knownvalue.StringExact("192.0.2.0/24")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("pattern_type"), knownvalue.StringExact("IP")),
				},
			},
		},
	})
}

// TestAccCloudflareEmailSecurityBlockSender_Regex covers is_regex = true, where
// the pattern is validated as a Java-compatible regex rather than as a literal
// address.
func TestAccCloudflareEmailSecurityBlockSender_Regex(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	resourceName := fmt.Sprintf("cloudflare_email_security_block_sender.%s", rnd)
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareEmailSecurityBlockSenderDestroy,
		Steps: []resource.TestStep{
			{
				Config: acctest.LoadTestCase("blocksender_regex.tf", rnd, accountID),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("pattern"), knownvalue.StringExact(fmt.Sprintf(`.*@%s\.example\.com`, rnd))),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("is_regex"), knownvalue.Bool(true)),
				},
			},
			{
				ResourceName:        resourceName,
				ImportState:         true,
				ImportStateVerify:   true,
				ImportStateIdPrefix: fmt.Sprintf("%s/", accountID),
			},
		},
	})
}

func testAccCheckCloudflareEmailSecurityBlockSenderDestroy(s *terraform.State) error {
	client := acctest.SharedClient()

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "cloudflare_email_security_block_sender" {
			continue
		}

		accountID := rs.Primary.Attributes[consts.AccountIDSchemaKey]
		_, err := client.EmailSecurity.Settings.BlockSenders.Get(
			context.Background(),
			rs.Primary.ID,
			email_security.SettingBlockSenderGetParams{
				AccountID: cloudflare.F(accountID),
			},
		)
		if err == nil {
			return fmt.Errorf("email security blocked sender still exists: %s", rs.Primary.ID)
		}
	}

	return nil
}
