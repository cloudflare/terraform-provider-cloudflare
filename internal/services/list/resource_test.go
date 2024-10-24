package list_test

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func init() {
	resource.AddTestSweepers("cloudflare_list", &resource.Sweeper{
		Name: "cloudflare_list",
		F:    testSweepCloudflareList,
	})
}

func testSweepCloudflareList(r string) error {
	ctx := context.Background()
	client, clientErr := acctest.SharedV1Client() // TODO(terraform): replace with SharedV2Clent
	if clientErr != nil {
		tflog.Error(ctx, fmt.Sprintf("Failed to create Cloudflare client: %s", clientErr))
	}

	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	if accountID == "" {
		return errors.New("CLOUDFLARE_ACCOUNT_ID must be set")
	}

	lists, err := client.ListLists(ctx, cloudflare.AccountIdentifier(accountID), cloudflare.ListListsParams{})
	if err != nil {
		tflog.Error(ctx, fmt.Sprintf("Failed to fetch Cloudflare Lists: %s", err))
	}

	if len(lists) == 0 {
		log.Print("[DEBUG] No Cloudflare Lists to sweep")
		return nil
	}

	for _, list := range lists {
		tflog.Info(ctx, fmt.Sprintf("Deleting Cloudflare List ID: %s", list.ID))
		//nolint:errcheck
		client.DeleteList(ctx, cloudflare.AccountIdentifier(accountID), list.ID)
	}

	return nil
}

func TestAccCloudflareList_Basic(t *testing.T) {
	// Temporarily unset CLOUDFLARE_API_TOKEN if it is set as the IP List
	// endpoint does not yet support the API tokens.
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	rndIP := utils.GenerateRandomResourceName()
	rndRedirect := utils.GenerateRandomResourceName()
	rndASN := utils.GenerateRandomResourceName()
	rndHostname := utils.GenerateRandomResourceName()

	nameIP := fmt.Sprintf("cloudflare_list.%s", rndIP)
	nameRedirect := fmt.Sprintf("cloudflare_list.%s", rndRedirect)
	nameASN := fmt.Sprintf("cloudflare_list.%s", rndASN)
	nameHostname := fmt.Sprintf("cloudflare_list.%s", rndHostname)

	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	var list cloudflare.List
	var initialID string

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareList(rndIP, rndIP, rndIP, accountID, "ip"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(nameIP, "name", rndIP),
				),
			},
			{
				PreConfig: func() {
					initialID = list.ID
				},
				Config: testAccCheckCloudflareListIPUpdate(rndIP, rndIP, rndIP, accountID),
				Check: resource.ComposeTestCheckFunc(
					func(state *terraform.State) error {
						if initialID != list.ID {
							return fmt.Errorf("wanted update but List got recreated (id changed %q -> %q)", initialID, list.ID)
						}
						return nil
					},
				),
			},
			{
				Config: testAccCheckCloudflareList(rndRedirect, rndRedirect, rndRedirect, accountID, "redirect"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(nameRedirect, "name", rndRedirect)),
			},
			{
				PreConfig: func() {
					initialID = list.ID
				},
				Config: testAccCheckCloudflareListRedirectUpdate(rndRedirect, rndRedirect, rndRedirect, accountID),
				Check: resource.ComposeTestCheckFunc(
					func(state *terraform.State) error {
						if initialID != list.ID {
							return fmt.Errorf("wanted update but List got recreated (id changed %q -> %q)", initialID, list.ID)
						}
						return nil
					},
				),
			},
			{
				PreConfig: func() {
					initialID = list.ID
				},
				Config: testAccCheckCloudflareListRedirectUpdateTargetUrl(rndRedirect, rndRedirect, rndRedirect, accountID),
				Check: resource.ComposeTestCheckFunc(
					func(state *terraform.State) error {
						if initialID != list.ID {
							return fmt.Errorf("wanted update but List got recreated (id changed %q -> %q)", initialID, list.ID)
						}
						return nil
					},
				),
			},
			{
				Config: testAccCheckCloudflareList(rndASN, rndASN, rndASN, accountID, "asn"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						nameASN, "name", rndASN),
				),
			},
			{
				PreConfig: func() {
					initialID = list.ID
				},
				Config: testAccCheckCloudflareListASNUpdate(rndASN, rndASN, rndASN, accountID),
				Check: resource.ComposeTestCheckFunc(
					func(state *terraform.State) error {
						if initialID != list.ID {
							return fmt.Errorf("wanted update but List got recreated (id changed %q -> %q)", initialID, list.ID)
						}
						return nil
					},
				),
			},
			{
				Config: testAccCheckCloudflareList(rndHostname, rndHostname, rndHostname, accountID, "hostname"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						nameHostname, "name", rndHostname),
				),
			},
			{
				PreConfig: func() {
					initialID = list.ID
				},
				Config: testAccCheckCloudflareListHostnameUpdate(rndHostname, rndHostname, rndHostname, accountID),
				Check: resource.ComposeTestCheckFunc(
					func(state *terraform.State) error {
						if initialID != list.ID {
							return fmt.Errorf("wanted update but List got recreated (id changed %q -> %q)", initialID, list.ID)
						}
						return nil
					},
				),
			},
		},
	})
}

func testAccCheckCloudflareList(ID, name, description, accountID, kind string) string {
	return acctest.LoadTestCase("list.tf", ID, name, description, accountID, kind)
}

func testAccCheckCloudflareListIPUpdate(ID, name, description, accountID string) string {
	return acctest.LoadTestCase("listipupdate.tf", ID, name, description, accountID)
}

func testAccCheckCloudflareListRedirectUpdate(ID, name, description, accountID string) string {
	return acctest.LoadTestCase("listredirectupdate.tf", ID, name, description, accountID)
}

func testAccCheckCloudflareListRedirectUpdateTargetUrl(ID, name, description, accountID string) string {
	return acctest.LoadTestCase("listredirectupdatetargeturl.tf", ID, name, description, accountID)
}

func testAccCheckCloudflareListBasicIP(ID, name, description, accountID string) string {
	return acctest.LoadTestCase("listbasicip.tf", ID, name, description, accountID)
}

func testAccCheckCloudflareListASNUpdate(ID, name, description, accountID string) string {
	return acctest.LoadTestCase("listasnupdate.tf", ID, name, description, accountID)
}

func testAccCheckCloudflareListHostnameUpdate(ID, name, description, accountID string) string {
	return acctest.LoadTestCase("listhostnameupdate.tf", ID, name, description, accountID)
}
