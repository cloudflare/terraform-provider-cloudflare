package byo_ip_prefix_test

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	cfv3 "github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/cloudflare-go/v6/addressing"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

func TestMain(m *testing.M) {
	resource.TestMain(m)
}

func init() {
	resource.AddTestSweepers("cloudflare_byo_ip_prefix", &resource.Sweeper{
		Name: "cloudflare_byo_ip_prefix",
		F:    testSweepCloudflareBYOIPPrefixes,
	})
}

func testSweepCloudflareBYOIPPrefixes(r string) error {
	ctx := context.Background()
	client := acctest.SharedClient()

	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	if accountID == "" {
		tflog.Info(ctx, "Skipping BYO IP prefixes sweep: CLOUDFLARE_ACCOUNT_ID not set")
		return nil
	}

	prefixes, err := client.Addressing.Prefixes.List(ctx, addressing.PrefixListParams{
		AccountID: cfv3.F(accountID),
	})
	if err != nil {
		tflog.Error(ctx, fmt.Sprintf("Failed to fetch BYO IP prefixes: %s", err))
		return fmt.Errorf("failed to fetch BYO IP prefixes: %w", err)
	}

	if len(prefixes.Result) == 0 {
		tflog.Info(ctx, "No BYO IP prefixes to sweep")
		return nil
	}

	for _, prefix := range prefixes.Result {
		// Use standard filtering helper on the description field
		if !utils.ShouldSweepResource(prefix.Description) {
			continue
		}

		tflog.Info(ctx, fmt.Sprintf("Deleting BYO IP prefix: %s (account: %s)", prefix.ID, accountID))
		_, err := client.Addressing.Prefixes.Delete(ctx, prefix.ID, addressing.PrefixDeleteParams{
			AccountID: cfv3.F(accountID),
		})
		if err != nil {
			tflog.Error(ctx, fmt.Sprintf("Failed to delete BYO IP prefix %s: %s", prefix.ID, err))
			continue
		}
		tflog.Info(ctx, fmt.Sprintf("Deleted BYO IP prefix: %s", prefix.ID))
	}

	return nil
}

func TestAccCloudflareBYOIPPrefix(t *testing.T) {
	t.Parallel()

	var (
		rnd  = utils.GenerateRandomResourceName()
		name = fmt.Sprintf("cloudflare_byo_ip_prefix.%s", rnd)

		accountID     = os.Getenv("CLOUDFLARE_ACCOUNT_ID")
		loaDocumentID = os.Getenv("CLOUDFLARE_BYO_IP_LOA_DOCUMENT_ID")
		cidr          = os.Getenv("CLOUDFLARE_BYO_IP_CIDR")
		asn, _        = strconv.ParseInt(os.Getenv("CLOUDFLARE_BYO_IP_ASN"), 10, 64)
	)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
			acctest.TestAccPreCheck_BYOIPPrefix(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareBYOIPPrefixConfig(accountID, asn, cidr, loaDocumentID, "This is my BYOIP prefix old", rnd),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(
						name,
						tfjsonpath.New("description"),
						knownvalue.StringExact("This is my BYOIP prefix old"),
					),
					statecheck.ExpectKnownValue(
						name,
						tfjsonpath.New("asn"),
						knownvalue.Int64Exact(asn),
					),
					statecheck.ExpectKnownValue(
						name,
						tfjsonpath.New("loa_document_id"),
						knownvalue.StringExact(loaDocumentID),
					),
					statecheck.ExpectKnownValue(
						name,
						tfjsonpath.New("cidr"),
						knownvalue.StringExact(cidr),
					),
				},
				PreventDiskCleanup: true,
			},
			{
				Config: testAccCheckCloudflareBYOIPPrefixConfig(accountID, asn, cidr, loaDocumentID, "This is my BYOIP prefix new", rnd),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(
						name,
						tfjsonpath.New("description"),
						knownvalue.StringExact("This is my BYOIP prefix new"),
					),
					statecheck.ExpectKnownValue(
						name,
						tfjsonpath.New("asn"),
						knownvalue.Int64Exact(asn),
					),
					statecheck.ExpectKnownValue(
						name,
						tfjsonpath.New("loa_document_id"),
						knownvalue.StringExact(loaDocumentID),
					),
					statecheck.ExpectKnownValue(
						name,
						tfjsonpath.New("cidr"),
						knownvalue.StringExact(cidr),
					),
				},
				PreventDiskCleanup: true,
			},
			{
				ResourceName:            name,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateIdPrefix:     fmt.Sprintf("%s/", accountID),
				ImportStateVerifyIgnore: []string{"irr_validation_state", "rpki_validation_state", "ownership_validation_state"},
			},
		},
	})
}

func testAccCheckCloudflareBYOIPPrefixConfig(accountID string, asn int64, cidr string, loaDocumentID string, description, name string) string {
	return acctest.LoadTestCase("byoipprefixconfig.tf", accountID, asn, cidr, loaDocumentID, description, name)
}
