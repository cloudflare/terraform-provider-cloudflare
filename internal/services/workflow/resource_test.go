package workflow_test

import (
	"context"
	"fmt"
	"os"
	"path"
	"regexp"
	"testing"

	"github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/cloudflare-go/v6/workflows"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

func TestAccCloudflareWorkflow(t *testing.T) {
	t.Parallel()

	rnd := utils.GenerateRandomResourceName()
	name := "cloudflare_workflow." + rnd
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	workflowName := fmt.Sprintf("test-workflow-%s", rnd)
	scriptName := fmt.Sprintf("test-script-%s", rnd)
	scriptNameUpdated := fmt.Sprintf("test-script-updated-%s", rnd)
	className := fmt.Sprintf("TestClass%s", rnd)
	classNameUpdated := fmt.Sprintf("TestClassUpdated%s", rnd)

	tmpDir := t.TempDir()
	contentFile := path.Join(tmpDir, "worker.js")

	err := os.WriteFile(contentFile, []byte("export default {fetch() {return new Response()}}"), 0644)
	if err != nil {
		t.Fatalf("Error creating temp file at path %s: %s", contentFile, err.Error())
	}

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareWorkflowDestroy,
		Steps: []resource.TestStep{
			{
				Config: acctest.LoadTestCase("workflow.tf", rnd, accountID, workflowName, scriptName, className, contentFile),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(name, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("workflow_name"), knownvalue.StringExact(workflowName)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("script_name"), knownvalue.StringExact(scriptName)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("class_name"), knownvalue.StringExact(className)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("name"), knownvalue.StringExact(workflowName)),
				},
				Check: testAccCheckCloudflareWorkflowExists(name, accountID, workflowName),
			},
			{
				Config: acctest.LoadTestCase("workflow.tf", rnd, accountID, workflowName, scriptNameUpdated, classNameUpdated, contentFile),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(name, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("workflow_name"), knownvalue.StringExact(workflowName)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("script_name"), knownvalue.StringExact(scriptNameUpdated)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("class_name"), knownvalue.StringExact(classNameUpdated)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("name"), knownvalue.StringExact(workflowName)),
				},
				Check: testAccCheckCloudflareWorkflowExists(name, accountID, workflowName),
			},
			{
				ResourceName:        name,
				ImportState:         true,
				ImportStateIdPrefix: fmt.Sprintf("%s/", accountID),
				ImportStateVerify:   true,
				ImportStateVerifyIgnore: []string{
                    "modified_on",
                    "version_id",
					"instances",
                },
			},
		},
	})
}

func TestAccCloudflareWorkflow_RequiredFields(t *testing.T) {
	t.Parallel()

	rnd := utils.GenerateRandomResourceName()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      acctest.LoadTestCase("workflow_missing_required.tf", rnd, accountID),
				ExpectError: regexp.MustCompile(`Missing required argument`),
			},
		},
	})
}

func testAccCheckCloudflareWorkflowDestroy(s *terraform.State) error {
	client := acctest.SharedClient()

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "cloudflare_workflow" {
			continue
		}

		accountID := rs.Primary.Attributes["account_id"]
		workflowName := rs.Primary.Attributes["workflow_name"]

		_, err := client.Workflows.Get(
			context.Background(),
			workflowName,
			workflows.WorkflowGetParams{
				AccountID: cloudflare.F(accountID),
			},
		)

		if err == nil {
			return fmt.Errorf("workflow %s still exists", workflowName)
		}
	}

	return nil
}

func testAccCheckCloudflareWorkflowExists(resourceName, accountID, workflowName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("no workflow ID is set")
		}

		client := acctest.SharedClient()
		_, err := client.Workflows.Get(
			context.Background(),
			workflowName,
			workflows.WorkflowGetParams{
				AccountID: cloudflare.F(accountID),
			},
		)

		if err != nil {
			return fmt.Errorf("workflow not found: %s", err)
		}

		return nil
	}
}
