package pages_project_test

import (
	"context"
	"fmt"
	"os"
	"regexp"
	"testing"

	"github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/cloudflare-go/v6/pages"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
)

func TestMain(m *testing.M) {
	resource.TestMain(m)
}

func init() {
	resource.AddTestSweepers("cloudflare_pages_project", &resource.Sweeper{
		Name: "cloudflare_pages_project",
		F:    testSweepCloudflarePagesProjects,
	})
}

func testSweepCloudflarePagesProjects(r string) error {
	ctx := context.Background()
	client := acctest.SharedClient()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	if accountID == "" {
		tflog.Info(ctx, "Skipping Pages projects sweep: CLOUDFLARE_ACCOUNT_ID not set")
		return nil
	}

	// List all pages projects in the account (this actually returns deployments)
	list, err := client.Pages.Projects.List(ctx, pages.ProjectListParams{
		AccountID: cloudflare.F(accountID),
	})
	if err != nil {
		tflog.Error(ctx, fmt.Sprintf("Failed to list pages projects: %s", err))
		return fmt.Errorf("failed to list pages projects: %w", err)
	}

	if len(list.Result) == 0 {
		tflog.Info(ctx, "No Pages projects to sweep")
		return nil
	}

	// Track unique project names to avoid duplicate deletions
	projectNames := make(map[string]bool)

	// Delete all pages projects with test prefix
	for _, project := range list.Result {
		// Use standard filtering helper
		if !utils.ShouldSweepResource(project.Name) {
			continue
		}

		// Only delete each project once (deployments can have multiple entries per project)
		if !projectNames[project.Name] {
			projectNames[project.Name] = true
			tflog.Info(ctx, fmt.Sprintf("Deleting Pages project: %s (account: %s)", project.Name, accountID))
			_, err := client.Pages.Projects.Delete(ctx, project.Name, pages.ProjectDeleteParams{
				AccountID: cloudflare.F(accountID),
			})
			if err != nil {
				tflog.Error(ctx, fmt.Sprintf("Failed to delete Pages project %s: %s", project.Name, err))
				continue
			}
			tflog.Info(ctx, fmt.Sprintf("Deleted Pages project: %s", project.Name))
		}
	}

	return nil
}

func testPagesProjectSource(resourceID, accountID, projectName, repoOwner, repoName string) string {
	return acctest.LoadTestCase("pagesprojectsource.tf", resourceID, accountID, projectName, repoOwner, repoName)
}

func testPagesProjectBuildConfig(resourceID, accountID, projectName string) string {
	return acctest.LoadTestCase("pagesprojectbuildconfig.tf", resourceID, accountID, projectName)
}

func testPagesProjectDeploymentConfig(resourceID, accountID, projectName string) string {
	return acctest.LoadTestCase("pagesprojectdeploymentconfig.tf", resourceID, accountID, projectName)
}

func testPagesProjectDirectUpload(resourceID, accountID, projectName string) string {
	return acctest.LoadTestCase("pagesprojectdirectupload.tf", resourceID, accountID, projectName)
}

func testPagesProjectMinimal(resourceID, accountID, projectName string) string {
	return acctest.LoadTestCase("pagesprojectminimal.tf", resourceID, accountID, projectName)
}

func testPagesProjectMinimalWithDeploymentConfigs(resourceID, accountID, projectName string) string {
	return acctest.LoadTestCase("pagesprojectminimalwithdeploymentconfigs.tf", resourceID, accountID, projectName)
}

func testPagesProjectFullConfig(resourceID, accountID, projectName, owner, repo string) string {
	return acctest.LoadTestCase("pagesprojectfullconfig.tf", resourceID, accountID, projectName, owner, repo)
}

func testPagesProjectEnvVars(resourceID, accountID, projectName string) string {
	return acctest.LoadTestCase("pagesprojectenvvars.tf", resourceID, accountID, projectName)
}

func testPagesProjectWithEnvVars(resourceID, accountID, projectName string) string {
	return acctest.LoadTestCase("pagesprojectwithenvvars.tf", resourceID, accountID, projectName)
}

func testPagesProjectWithoutEnvVars(resourceID, accountID, projectName string) string {
	return acctest.LoadTestCase("pagesprojectwithoutenvvars.tf", resourceID, accountID, projectName)
}

func testPagesProjectWithOnlyOneEnvVar(resourceID, accountID, projectName string) string {
	return acctest.LoadTestCase("pagesprojectwithonlyoneenvvar.tf", resourceID, accountID, projectName)
}

func testPagesProjectPreviewSettings(resourceID, accountID, projectName, owner, repo, setting, extraConfig string) string {
	return acctest.LoadTestCase("pagesprojectpreviewsettings.tf", resourceID, accountID, projectName, owner, repo, setting, extraConfig)
}

func testPagesProjectUpdated(resourceID, accountID, projectName string) string {
	return acctest.LoadTestCase("pagesprojectupdated.tf", resourceID, accountID, projectName)
}

func testPagesProjectDeploymentConfigsNoBuildConfig(resourceID, accountID, projectName string) string {
	return acctest.LoadTestCase("pagesprojectdeploymentconfigsnobuildconfig.tf", resourceID, accountID, projectName)
}

func testPagesProjectPartialBuildConfig(resourceID, accountID, projectName string) string {
	return acctest.LoadTestCase("pagesprojectpartialbuildconfig.tf", resourceID, accountID, projectName)
}

func testAccCheckCloudflarePageProjectDestroy(s *terraform.State) error {
	client := acctest.SharedClient()

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "cloudflare_pages_project" {
			continue
		}

		accountID := rs.Primary.Attributes[consts.AccountIDSchemaKey]
		_, err := client.Pages.Projects.Get(context.Background(), rs.Primary.ID, pages.ProjectGetParams{
			AccountID: cloudflare.F(accountID),
		})
		if err == nil {
			return fmt.Errorf("pages project still exists")
		}
	}

	return nil
}

func TestAccCloudflarePagesProject_Basic(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	name := "cloudflare_pages_project." + rnd
	projectName := rnd
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	pagesOwner := os.Getenv("CLOUDFLARE_PAGES_OWNER")
	pagesRepo := os.Getenv("CLOUDFLARE_PAGES_REPO")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_Pages(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflarePageProjectDestroy,
		Steps: []resource.TestStep{
			{
				Config: testPagesProjectSource(rnd, accountID, projectName, pagesOwner, pagesRepo),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(name, tfjsonpath.New("name"), knownvalue.StringExact(projectName)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("source").AtMapKey("type"), knownvalue.StringExact("github")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("source").AtMapKey("config").AtMapKey("owner"), knownvalue.StringExact(pagesOwner)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("source").AtMapKey("config").AtMapKey("repo_name"), knownvalue.StringExact(pagesRepo)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("source").AtMapKey("config").AtMapKey("production_branch"), knownvalue.StringExact("main")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("source").AtMapKey("config").AtMapKey("pr_comments_enabled"), knownvalue.Bool(true)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("source").AtMapKey("config").AtMapKey("deployments_enabled"), knownvalue.Bool(true)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("source").AtMapKey("config").AtMapKey("production_deployments_enabled"), knownvalue.Bool(true)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("source").AtMapKey("config").AtMapKey("preview_deployment_setting"), knownvalue.StringExact("custom")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("source").AtMapKey("config").AtMapKey("preview_branch_includes"), knownvalue.ListSizeExact(2)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("source").AtMapKey("config").AtMapKey("preview_branch_includes").AtSliceIndex(0), knownvalue.StringExact("dev")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("source").AtMapKey("config").AtMapKey("preview_branch_includes").AtSliceIndex(1), knownvalue.StringExact("preview")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("source").AtMapKey("config").AtMapKey("preview_branch_excludes"), knownvalue.ListSizeExact(2)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("source").AtMapKey("config").AtMapKey("preview_branch_excludes").AtSliceIndex(0), knownvalue.StringExact("main")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("source").AtMapKey("config").AtMapKey("preview_branch_excludes").AtSliceIndex(1), knownvalue.StringExact("prod")),
				},
				//ExpectNonEmptyPlan: true, // Computed fields like canonical_deployment, latest_deployment can change
			},
			{
				ResourceName:        name,
				ImportState:         true,
				ImportStateVerify:   true,
				ImportStateIdPrefix: fmt.Sprintf("%s/", accountID),
			},
		},
	})
}

func TestAccCloudflarePagesProject_BuildConfig(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	name := "cloudflare_pages_project." + rnd
	projectName := rnd
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_Pages(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflarePageProjectDestroy,
		Steps: []resource.TestStep{
			{
				Config: testPagesProjectBuildConfig(rnd, accountID, projectName),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(name, tfjsonpath.New("name"), knownvalue.StringExact(projectName)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("build_config").AtMapKey("build_caching"), knownvalue.Bool(true)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("build_config").AtMapKey("build_command"), knownvalue.StringExact("npm run build")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("build_config").AtMapKey("destination_dir"), knownvalue.StringExact("build")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("build_config").AtMapKey("root_dir"), knownvalue.StringExact("/")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("build_config").AtMapKey("web_analytics_tag"), knownvalue.StringExact("cee1c73f6e4743d0b5e6bb1a0bcaabcc")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("build_config").AtMapKey("web_analytics_token"), knownvalue.StringExact("021e1057c18547eca7b79f2516f06o7x")),
				},
			},
			{
				ResourceName:        name,
				ImportState:         true,
				ImportStateVerify:   true,
				ImportStateIdPrefix: fmt.Sprintf("%s/", accountID),
			},
		},
	})
}

func TestAccCloudflarePagesProject_DeploymentConfig(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	name := "cloudflare_pages_project." + rnd
	projectName := rnd
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_Pages(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflarePageProjectDestroy,
		Steps: []resource.TestStep{
			{
				Config: testPagesProjectDeploymentConfig(rnd, accountID, projectName),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(name, tfjsonpath.New("name"), knownvalue.StringExact(projectName)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),

					// Preview
					statecheck.ExpectKnownValue(name, tfjsonpath.New("deployment_configs").AtMapKey("preview").AtMapKey("env_vars"), knownvalue.MapSizeExact(2)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("deployment_configs").AtMapKey("preview").AtMapKey("env_vars").AtMapKey("ENVIRONMENT").AtMapKey("type"), knownvalue.StringExact("plain_text")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("deployment_configs").AtMapKey("preview").AtMapKey("env_vars").AtMapKey("ENVIRONMENT").AtMapKey("value"), knownvalue.StringExact("preview")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("deployment_configs").AtMapKey("preview").AtMapKey("env_vars").AtMapKey("TURNSTILE_SECRET").AtMapKey("type"), knownvalue.StringExact("secret_text")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("deployment_configs").AtMapKey("preview").AtMapKey("env_vars").AtMapKey("TURNSTILE_SECRET").AtMapKey("value"), knownvalue.StringExact("1x0000000000000000000000000000000AA")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("deployment_configs").AtMapKey("preview").AtMapKey("kv_namespaces"), knownvalue.MapSizeExact(1)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("deployment_configs").AtMapKey("preview").AtMapKey("kv_namespaces").AtMapKey("KV_BINDING").AtMapKey("namespace_id"), knownvalue.StringRegexp(regexp.MustCompile("^[0-9a-f]{32}$"))),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("deployment_configs").AtMapKey("preview").AtMapKey("durable_object_namespaces"), knownvalue.MapSizeExact(1)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("deployment_configs").AtMapKey("preview").AtMapKey("durable_object_namespaces").AtMapKey("DO_BINDING").AtMapKey("namespace_id"), knownvalue.StringExact("5eb63bbbe01eeed093cb22bb8f5acdc3")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("deployment_configs").AtMapKey("preview").AtMapKey("d1_databases"), knownvalue.MapSizeExact(1)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("deployment_configs").AtMapKey("preview").AtMapKey("d1_databases").AtMapKey("D1_BINDING").AtMapKey("id"), knownvalue.StringExact("445e2955-951a-4358-a35b-a4d0c813f63")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("deployment_configs").AtMapKey("preview").AtMapKey("r2_buckets"), knownvalue.MapSizeExact(1)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("deployment_configs").AtMapKey("preview").AtMapKey("r2_buckets").AtMapKey("R2_BINDING").AtMapKey("name"), knownvalue.StringExact("some-bucket")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("deployment_configs").AtMapKey("preview").AtMapKey("compatibility_date"), knownvalue.StringExact("2022-08-15")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("deployment_configs").AtMapKey("preview").AtMapKey("compatibility_flags"), knownvalue.ListSizeExact(1)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("deployment_configs").AtMapKey("preview").AtMapKey("compatibility_flags").AtSliceIndex(0), knownvalue.StringExact("preview_flag")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("deployment_configs").AtMapKey("preview").AtMapKey("fail_open"), knownvalue.Bool(true)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("deployment_configs").AtMapKey("preview").AtMapKey("always_use_latest_compatibility_date"), knownvalue.Bool(true)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("deployment_configs").AtMapKey("preview").AtMapKey("usage_model"), knownvalue.StringExact("standard")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("deployment_configs").AtMapKey("preview").AtMapKey("placement"), knownvalue.Null()),

					// Production
					statecheck.ExpectKnownValue(name, tfjsonpath.New("deployment_configs").AtMapKey("production").AtMapKey("env_vars"), knownvalue.MapSizeExact(4)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("deployment_configs").AtMapKey("production").AtMapKey("env_vars").AtMapKey("ENVIRONMENT").AtMapKey("type"), knownvalue.StringExact("plain_text")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("deployment_configs").AtMapKey("production").AtMapKey("env_vars").AtMapKey("ENVIRONMENT").AtMapKey("value"), knownvalue.StringExact("production")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("deployment_configs").AtMapKey("production").AtMapKey("env_vars").AtMapKey("OTHER_VALUE").AtMapKey("type"), knownvalue.StringExact("plain_text")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("deployment_configs").AtMapKey("production").AtMapKey("env_vars").AtMapKey("OTHER_VALUE").AtMapKey("value"), knownvalue.StringExact("other value")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("deployment_configs").AtMapKey("production").AtMapKey("env_vars").AtMapKey("TURNSTILE_SECRET").AtMapKey("type"), knownvalue.StringExact("secret_text")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("deployment_configs").AtMapKey("production").AtMapKey("env_vars").AtMapKey("TURNSTILE_SECRET").AtMapKey("value"), knownvalue.StringExact("1x0000000000000000000000000000000AA")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("deployment_configs").AtMapKey("production").AtMapKey("env_vars").AtMapKey("TURNSTILE_INVIS_SECRET").AtMapKey("type"), knownvalue.StringExact("secret_text")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("deployment_configs").AtMapKey("production").AtMapKey("env_vars").AtMapKey("TURNSTILE_INVIS_SECRET").AtMapKey("value"), knownvalue.StringExact("2x0000000000000000000000000000000AA")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("deployment_configs").AtMapKey("production").AtMapKey("kv_namespaces"), knownvalue.MapSizeExact(2)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("deployment_configs").AtMapKey("production").AtMapKey("kv_namespaces").AtMapKey("KV_BINDING_1").AtMapKey("namespace_id"), knownvalue.StringRegexp(regexp.MustCompile("^[0-9a-f]{32}$"))),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("deployment_configs").AtMapKey("production").AtMapKey("kv_namespaces").AtMapKey("KV_BINDING_2").AtMapKey("namespace_id"), knownvalue.StringRegexp(regexp.MustCompile("^[0-9a-f]{32}$"))),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("deployment_configs").AtMapKey("production").AtMapKey("durable_object_namespaces"), knownvalue.MapSizeExact(2)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("deployment_configs").AtMapKey("production").AtMapKey("durable_object_namespaces").AtMapKey("DO_BINDING_1").AtMapKey("namespace_id"), knownvalue.StringExact("5eb63bbbe01eeed093cb22bb8f5acdc3")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("deployment_configs").AtMapKey("production").AtMapKey("durable_object_namespaces").AtMapKey("DO_BINDING_2").AtMapKey("namespace_id"), knownvalue.StringExact("3cdca5f8bb22bc390deee10ebbb36be5")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("deployment_configs").AtMapKey("production").AtMapKey("d1_databases"), knownvalue.MapSizeExact(2)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("deployment_configs").AtMapKey("production").AtMapKey("d1_databases").AtMapKey("D1_BINDING_1").AtMapKey("id"), knownvalue.StringExact("445e2955-951a-4358-a35b-a4d0c813f63")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("deployment_configs").AtMapKey("production").AtMapKey("d1_databases").AtMapKey("D1_BINDING_2").AtMapKey("id"), knownvalue.StringExact("a399414b-c697-409a-a688-377db6433cd9")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("deployment_configs").AtMapKey("production").AtMapKey("r2_buckets"), knownvalue.MapSizeExact(2)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("deployment_configs").AtMapKey("production").AtMapKey("r2_buckets").AtMapKey("R2_BINDING_1").AtMapKey("name"), knownvalue.StringExact("some-bucket")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("deployment_configs").AtMapKey("production").AtMapKey("r2_buckets").AtMapKey("R2_BINDING_2").AtMapKey("name"), knownvalue.StringExact("other-bucket")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("deployment_configs").AtMapKey("production").AtMapKey("compatibility_date"), knownvalue.StringExact("2022-08-16")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("deployment_configs").AtMapKey("production").AtMapKey("compatibility_flags"), knownvalue.ListSizeExact(2)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("deployment_configs").AtMapKey("production").AtMapKey("compatibility_flags").AtSliceIndex(0), knownvalue.StringExact("production_flag")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("deployment_configs").AtMapKey("production").AtMapKey("compatibility_flags").AtSliceIndex(1), knownvalue.StringExact("second flag")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("deployment_configs").AtMapKey("production").AtMapKey("fail_open"), knownvalue.Bool(true)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("deployment_configs").AtMapKey("production").AtMapKey("always_use_latest_compatibility_date"), knownvalue.Bool(false)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("deployment_configs").AtMapKey("production").AtMapKey("usage_model"), knownvalue.StringExact("standard")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("deployment_configs").AtMapKey("production").AtMapKey("placement").AtMapKey("mode"), knownvalue.StringExact("smart")),
				},
			},
			{
				ResourceName:        name,
				ImportState:         true,
				ImportStateVerify:   true,
				ImportStateIdPrefix: fmt.Sprintf("%s/", accountID),
				ImportStateVerifyIgnore: []string{
					"deployment_configs.preview.env_vars.TURNSTILE_SECRET.value",
					"deployment_configs.production.env_vars.TURNSTILE_SECRET.value",
					"deployment_configs.production.env_vars.TURNSTILE_INVIS_SECRET.value",
				},
			},
		},
	})
}

func TestAccCloudflarePagesProject_DirectUpload(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	name := "cloudflare_pages_project." + rnd
	projectName := rnd
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflarePageProjectDestroy,
		Steps: []resource.TestStep{
			{
				Config: testPagesProjectDirectUpload(rnd, accountID, projectName),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(name, tfjsonpath.New("name"), knownvalue.StringExact(projectName)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("production_branch"), knownvalue.StringExact("main")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("created_on"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("subdomain"), knownvalue.NotNull()),
				},
			},
			{
				ResourceName:            name,
				ImportStateIdPrefix:     fmt.Sprintf("%s/", accountID),
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"build_config", "deployment_configs", "canonical_deployment", "latest_deployment", "created_on", "subdomain", "domains"},
			},
		},
	})
}

func TestAccCloudflarePagesProject_Update_AddOptionalAttributes(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	name := "cloudflare_pages_project." + rnd
	projectName := rnd
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflarePageProjectDestroy,
		Steps: []resource.TestStep{
			{
				Config: testPagesProjectMinimal(rnd, accountID, projectName),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(name, tfjsonpath.New("name"), knownvalue.StringExact(projectName)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("production_branch"), knownvalue.StringExact("main")),
					// build_config is Optional and will be null when not specified
					statecheck.ExpectKnownValue(name, tfjsonpath.New("deployment_configs"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("source"), knownvalue.Null()),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("created_on"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("subdomain"), knownvalue.NotNull()),
				},
			},
			{
				Config: testPagesProjectUpdated(rnd, accountID, projectName),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(name, plancheck.ResourceActionUpdate),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(name, tfjsonpath.New("name"), knownvalue.StringExact(projectName)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("production_branch"), knownvalue.StringExact("develop")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("build_config").AtMapKey("build_caching"), knownvalue.Bool(true)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("build_config").AtMapKey("build_command"), knownvalue.StringExact("yarn build")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("deployment_configs").AtMapKey("preview").AtMapKey("compatibility_date"), knownvalue.StringExact("2023-06-01")),
				},
			},
			{
				ResourceName:            name,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateIdPrefix:     fmt.Sprintf("%s/", accountID),
				ImportStateVerifyIgnore: []string{"deployment_configs.production.env_vars.PROD_UPDATED.value"},
			},
		},
	})
}

func TestAccCloudflarePagesProject_Update_RemoveOptionalAttributes(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	name := "cloudflare_pages_project." + rnd
	projectName := rnd
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflarePageProjectDestroy,
		Steps: []resource.TestStep{
			{
				Config: testPagesProjectUpdated(rnd, accountID, projectName),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(name, tfjsonpath.New("production_branch"), knownvalue.StringExact("develop")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("build_config").AtMapKey("build_caching"), knownvalue.Bool(true)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("deployment_configs").AtMapKey("preview").AtMapKey("compatibility_date"), knownvalue.StringExact("2023-06-01")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("deployment_configs").AtMapKey("preview").AtMapKey("env_vars"), knownvalue.MapSizeExact(1)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("deployment_configs").AtMapKey("production").AtMapKey("env_vars"), knownvalue.MapSizeExact(1)),
				},
			},
			{
				// Update to config with deployment_configs but no env_vars - should remove env_vars.
				Config: testPagesProjectMinimalWithDeploymentConfigs(rnd, accountID, projectName),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(name, plancheck.ResourceActionUpdate),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(name, tfjsonpath.New("name"), knownvalue.StringExact(projectName)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("production_branch"), knownvalue.StringExact("main")),
					// env_vars should be removed (null) after update
					statecheck.ExpectKnownValue(name, tfjsonpath.New("deployment_configs").AtMapKey("preview").AtMapKey("env_vars"), knownvalue.Null()),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("deployment_configs").AtMapKey("production").AtMapKey("env_vars"), knownvalue.Null()),
				},
			},
			{
				ResourceName:        name,
				ImportState:         true,
				ImportStateVerify:   true,
				ImportStateIdPrefix: fmt.Sprintf("%s/", accountID),
			},
		},
	})
}
func TestAccCloudflarePagesProject_FullConfiguration(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	name := "cloudflare_pages_project." + rnd
	projectName := rnd
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	pagesOwner := os.Getenv("CLOUDFLARE_PAGES_OWNER")
	pagesRepo := os.Getenv("CLOUDFLARE_PAGES_REPO")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_Pages(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflarePageProjectDestroy,
		Steps: []resource.TestStep{
			{
				Config: testPagesProjectFullConfig(rnd, accountID, projectName, pagesOwner, pagesRepo),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(name, tfjsonpath.New("name"), knownvalue.StringExact(projectName)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("production_branch"), knownvalue.StringExact("main")),

					// Build config
					statecheck.ExpectKnownValue(name, tfjsonpath.New("build_config").AtMapKey("build_caching"), knownvalue.Bool(true)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("build_config").AtMapKey("build_command"), knownvalue.StringExact("npm run build")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("build_config").AtMapKey("destination_dir"), knownvalue.StringExact("dist")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("build_config").AtMapKey("root_dir"), knownvalue.StringExact("/app")),

					// Source config
					statecheck.ExpectKnownValue(name, tfjsonpath.New("source").AtMapKey("type"), knownvalue.StringExact("github")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("source").AtMapKey("config").AtMapKey("preview_deployment_setting"), knownvalue.StringExact("all")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("source").AtMapKey("config").AtMapKey("path_includes"), knownvalue.ListSizeExact(2)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("source").AtMapKey("config").AtMapKey("path_excludes"), knownvalue.ListSizeExact(2)),

					// Preview deployment configs
					statecheck.ExpectKnownValue(name, tfjsonpath.New("deployment_configs").AtMapKey("preview").AtMapKey("compatibility_date"), knownvalue.StringExact("2023-01-15")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("deployment_configs").AtMapKey("preview").AtMapKey("compatibility_flags"), knownvalue.ListSizeExact(2)),

					// Environment variables
					statecheck.ExpectKnownValue(name, tfjsonpath.New("deployment_configs").AtMapKey("preview").AtMapKey("env_vars"), knownvalue.MapSizeExact(2)),

					// Bindings - test all the new binding types
					statecheck.ExpectKnownValue(name, tfjsonpath.New("deployment_configs").AtMapKey("preview").AtMapKey("kv_namespaces"), knownvalue.MapSizeExact(1)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("deployment_configs").AtMapKey("preview").AtMapKey("d1_databases"), knownvalue.MapSizeExact(1)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("deployment_configs").AtMapKey("preview").AtMapKey("r2_buckets"), knownvalue.MapSizeExact(1)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("deployment_configs").AtMapKey("preview").AtMapKey("ai_bindings"), knownvalue.MapSizeExact(1)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("deployment_configs").AtMapKey("preview").AtMapKey("analytics_engine_datasets"), knownvalue.MapSizeExact(1)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("deployment_configs").AtMapKey("preview").AtMapKey("browsers"), knownvalue.MapSizeExact(1)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("deployment_configs").AtMapKey("preview").AtMapKey("hyperdrive_bindings"), knownvalue.MapSizeExact(1)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("deployment_configs").AtMapKey("preview").AtMapKey("mtls_certificates"), knownvalue.MapSizeExact(1)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("deployment_configs").AtMapKey("preview").AtMapKey("queue_producers"), knownvalue.MapSizeExact(1)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("deployment_configs").AtMapKey("preview").AtMapKey("services"), knownvalue.MapSizeExact(1)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("deployment_configs").AtMapKey("preview").AtMapKey("vectorize_bindings"), knownvalue.MapSizeExact(1)),

					// Placement
					statecheck.ExpectKnownValue(name, tfjsonpath.New("deployment_configs").AtMapKey("preview").AtMapKey("placement").AtMapKey("mode"), knownvalue.StringExact("smart")),

					// Production deployment configs
					statecheck.ExpectKnownValue(name, tfjsonpath.New("deployment_configs").AtMapKey("production").AtMapKey("compatibility_date"), knownvalue.StringExact("2023-01-16")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("deployment_configs").AtMapKey("production").AtMapKey("env_vars"), knownvalue.MapSizeExact(1)),

					// Computed attributes
					statecheck.ExpectKnownValue(name, tfjsonpath.New("created_on"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("subdomain"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("domains"), knownvalue.NotNull()),
				},
			},
			{
				ResourceName:        name,
				ImportState:         true,
				ImportStateVerify:   true,
				ImportStateIdPrefix: fmt.Sprintf("%s/", accountID),
				ImportStateVerifyIgnore: []string{
					"deployment_configs.preview.env_vars.SECRET_KEY.value",
				},
			},
		},
	})
}

func TestAccCloudflarePagesProject_EnvVarTypes(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	name := "cloudflare_pages_project." + rnd
	projectName := rnd
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflarePageProjectDestroy,
		Steps: []resource.TestStep{
			{
				Config: testPagesProjectEnvVars(rnd, accountID, projectName),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(name, tfjsonpath.New("name"), knownvalue.StringExact(projectName)),

					// Preview env vars - test both types
					statecheck.ExpectKnownValue(name, tfjsonpath.New("deployment_configs").AtMapKey("preview").AtMapKey("env_vars").AtMapKey("PLAIN_TEXT_VAR").AtMapKey("type"), knownvalue.StringExact("plain_text")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("deployment_configs").AtMapKey("preview").AtMapKey("env_vars").AtMapKey("SECRET_VAR").AtMapKey("type"), knownvalue.StringExact("secret_text")),

					// Production env vars - test both types
					statecheck.ExpectKnownValue(name, tfjsonpath.New("deployment_configs").AtMapKey("production").AtMapKey("env_vars").AtMapKey("PROD_PLAIN").AtMapKey("type"), knownvalue.StringExact("plain_text")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("deployment_configs").AtMapKey("production").AtMapKey("env_vars").AtMapKey("PROD_SECRET").AtMapKey("type"), knownvalue.StringExact("secret_text")),
				},
			},
			{
				ResourceName:            name,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateIdPrefix:     fmt.Sprintf("%s/", accountID),
				ImportStateVerifyIgnore: []string{"build_config", "deployment_configs.preview.compatibility_date", "deployment_configs.production.compatibility_date", "deployment_configs.preview.env_vars", "deployment_configs.production.env_vars"},
			},
		},
	})
}

func TestAccCloudflarePagesProject_PreviewDeploymentSettings(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	name := "cloudflare_pages_project." + rnd
	projectName := rnd
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	pagesOwner := os.Getenv("CLOUDFLARE_PAGES_OWNER")
	pagesRepo := os.Getenv("CLOUDFLARE_PAGES_REPO")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_Pages(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflarePageProjectDestroy,
		Steps: []resource.TestStep{
			{
				Config: testPagesProjectPreviewSettings(rnd, accountID, projectName, pagesOwner, pagesRepo, "all", ""),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(name, tfjsonpath.New("name"), knownvalue.StringExact(projectName)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("source").AtMapKey("config").AtMapKey("preview_deployment_setting"), knownvalue.StringExact("all")),
				},
			},
			{
				Config: testPagesProjectPreviewSettings(rnd, accountID, projectName, pagesOwner, pagesRepo, "none", ""),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(name, plancheck.ResourceActionUpdate),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(name, tfjsonpath.New("source").AtMapKey("config").AtMapKey("preview_deployment_setting"), knownvalue.StringExact("none")),
				},
			},
			{
				Config: testPagesProjectPreviewSettings(rnd, accountID, projectName, pagesOwner, pagesRepo, "custom", `
				preview_branch_includes = ["dev", "staging"]
				preview_branch_excludes = ["main", "prod"]
				`),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(name, plancheck.ResourceActionUpdate),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(name, tfjsonpath.New("source").AtMapKey("config").AtMapKey("preview_deployment_setting"), knownvalue.StringExact("custom")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("source").AtMapKey("config").AtMapKey("preview_branch_includes"), knownvalue.ListSizeExact(2)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("source").AtMapKey("config").AtMapKey("preview_branch_excludes"), knownvalue.ListSizeExact(2)),
				},
			},
			{
				ResourceName:        name,
				ImportState:         true,
				ImportStateVerify:   true,
				ImportStateIdPrefix: fmt.Sprintf("%s/", accountID),
			},
		},
	})
}

func TestAccCloudflarePagesProject_RemoveEnvVarsAndBindings(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	name := "cloudflare_pages_project." + rnd
	projectName := rnd
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflarePageProjectDestroy,
		Steps: []resource.TestStep{
			{
				// Step 1: Create with env_vars and bindings
				Config: testPagesProjectWithEnvVars(rnd, accountID, projectName),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(name, tfjsonpath.New("name"), knownvalue.StringExact(projectName)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("deployment_configs").AtMapKey("preview").AtMapKey("env_vars"), knownvalue.MapSizeExact(2)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("deployment_configs").AtMapKey("preview").AtMapKey("kv_namespaces"), knownvalue.MapSizeExact(1)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("deployment_configs").AtMapKey("preview").AtMapKey("d1_databases"), knownvalue.MapSizeExact(1)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("deployment_configs").AtMapKey("production").AtMapKey("env_vars"), knownvalue.MapSizeExact(2)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("deployment_configs").AtMapKey("production").AtMapKey("kv_namespaces"), knownvalue.MapSizeExact(1)),
				},
			},
			{
				// Step 2: Remove all env_vars and bindings
				Config: testPagesProjectWithoutEnvVars(rnd, accountID, projectName),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(name, plancheck.ResourceActionUpdate),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(name, tfjsonpath.New("name"), knownvalue.StringExact(projectName)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("deployment_configs").AtMapKey("preview").AtMapKey("env_vars"), knownvalue.Null()),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("deployment_configs").AtMapKey("preview").AtMapKey("kv_namespaces"), knownvalue.Null()),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("deployment_configs").AtMapKey("preview").AtMapKey("d1_databases"), knownvalue.Null()),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("deployment_configs").AtMapKey("production").AtMapKey("env_vars"), knownvalue.Null()),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("deployment_configs").AtMapKey("production").AtMapKey("kv_namespaces"), knownvalue.Null()),
				},
			},
			{
				ResourceName:        name,
				ImportState:         true,
				ImportStateVerify:   true,
				ImportStateIdPrefix: fmt.Sprintf("%s/", accountID),
			},
		},
	})
}

func TestAccCloudflarePagesProject_RemoveSpecificEnvVar(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	name := "cloudflare_pages_project." + rnd
	projectName := rnd
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflarePageProjectDestroy,
		Steps: []resource.TestStep{
			{
				// Step 1: Create with multiple env_vars
				Config: testPagesProjectWithEnvVars(rnd, accountID, projectName),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(name, tfjsonpath.New("name"), knownvalue.StringExact(projectName)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("deployment_configs").AtMapKey("preview").AtMapKey("env_vars"), knownvalue.MapSizeExact(2)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("deployment_configs").AtMapKey("production").AtMapKey("env_vars"), knownvalue.MapSizeExact(2)),
				},
			},
			{
				// Step 2: Remove one env_var (SECRET_VAR) from each environment
				Config: testPagesProjectWithOnlyOneEnvVar(rnd, accountID, projectName),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(name, plancheck.ResourceActionUpdate),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(name, tfjsonpath.New("name"), knownvalue.StringExact(projectName)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("deployment_configs").AtMapKey("preview").AtMapKey("env_vars"), knownvalue.MapSizeExact(1)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("deployment_configs").AtMapKey("preview").AtMapKey("env_vars").AtMapKey("ENVIRONMENT").AtMapKey("type"), knownvalue.StringExact("plain_text")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("deployment_configs").AtMapKey("production").AtMapKey("env_vars"), knownvalue.MapSizeExact(1)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("deployment_configs").AtMapKey("production").AtMapKey("env_vars").AtMapKey("ENVIRONMENT").AtMapKey("type"), knownvalue.StringExact("plain_text")),
				},
			},
			{
				ResourceName:        name,
				ImportState:         true,
				ImportStateVerify:   true,
				ImportStateIdPrefix: fmt.Sprintf("%s/", accountID),
			},
		},
	})
}

// TestAccCloudflarePagesProject_NoDriftWithSecretEnvVars tests the fix for issue #6526.
// It verifies that applying the same config with secret env vars does not cause
// "inconsistent result after apply" errors.
func TestAccCloudflarePagesProject_NoDriftWithSecretEnvVars(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	name := "cloudflare_pages_project." + rnd
	projectName := rnd
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflarePageProjectDestroy,
		Steps: []resource.TestStep{
			{
				// Initial apply with secret env vars
				Config: testPagesProjectEnvVars(rnd, accountID, projectName),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(name, tfjsonpath.New("deployment_configs").AtMapKey("preview").AtMapKey("env_vars").AtMapKey("SECRET_VAR").AtMapKey("type"), knownvalue.StringExact("secret_text")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("deployment_configs").AtMapKey("production").AtMapKey("env_vars").AtMapKey("PROD_SECRET").AtMapKey("type"), knownvalue.StringExact("secret_text")),
				},
			},
			{
				// Re-apply the SAME config - should show no changes (fixes #6526)
				Config: testPagesProjectEnvVars(rnd, accountID, projectName),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectEmptyPlan(),
					},
				},
			},
		},
	})
}

// TestAccCloudflarePagesProject_CompleteBuildConfig tests that specifying build_config
// with all required fields (including build_caching) works correctly after import.
// This addresses the issue where users needed to add build_caching to their config
// to avoid "Provider produced invalid plan" errors.
func TestAccCloudflarePagesProject_CompleteBuildConfig(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	name := "cloudflare_pages_project." + rnd
	projectName := rnd
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflarePageProjectDestroy,
		Steps: []resource.TestStep{
			{
				// Create with complete build_config (including build_caching)
				Config: testPagesProjectPartialBuildConfig(rnd, accountID, projectName),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(name, tfjsonpath.New("name"), knownvalue.StringExact(projectName)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("build_config").AtMapKey("build_command"), knownvalue.StringExact("yarn build")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("build_config").AtMapKey("destination_dir"), knownvalue.StringExact("dist")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("build_config").AtMapKey("build_caching"), knownvalue.Bool(false)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("build_config").AtMapKey("root_dir"), knownvalue.StringExact("/")),
				},
			},
			{
				// Re-apply the SAME config - should show no changes
				Config: testPagesProjectPartialBuildConfig(rnd, accountID, projectName),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectEmptyPlan(),
					},
				},
			},
			{
				// Import should work without issues
				ResourceName:        name,
				ImportState:         true,
				ImportStateVerify:   true,
				ImportStateIdPrefix: fmt.Sprintf("%s/", accountID),
			},
			{
				// After import, re-apply with complete build_config should work
				Config: testPagesProjectPartialBuildConfig(rnd, accountID, projectName),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectEmptyPlan(),
					},
				},
			},
		},
	})
}

// TestAccCloudflarePagesProject_DeploymentConfigsWithoutBuildConfig tests the fix for
// the issue where specifying deployment_configs without build_config would cause
// "Provider produced invalid plan" errors because the API returns empty strings for
// build_config fields, which weren't being properly normalized to null.
func TestAccCloudflarePagesProject_DeploymentConfigsWithoutBuildConfig(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	name := "cloudflare_pages_project." + rnd
	projectName := rnd
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflarePageProjectDestroy,
		Steps: []resource.TestStep{
			{
				// Initial create with deployment_configs but NO build_config
				// This should not cause "planned value for a non-computed attribute" error
				Config: testPagesProjectDeploymentConfigsNoBuildConfig(rnd, accountID, projectName),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(name, tfjsonpath.New("name"), knownvalue.StringExact(projectName)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("production_branch"), knownvalue.StringExact("main")),
					// deployment_configs should be present
					statecheck.ExpectKnownValue(name, tfjsonpath.New("deployment_configs").AtMapKey("preview").AtMapKey("compatibility_date"), knownvalue.StringExact("2023-08-25")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("deployment_configs").AtMapKey("production").AtMapKey("compatibility_date"), knownvalue.StringExact("2023-08-25")),
					// build_config should be null since not specified
					statecheck.ExpectKnownValue(name, tfjsonpath.New("build_config"), knownvalue.Null()),
				},
			},
			{
				// Re-apply the SAME config - should show no changes
				Config: testPagesProjectDeploymentConfigsNoBuildConfig(rnd, accountID, projectName),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectEmptyPlan(),
					},
				},
			},
			{
				// Import should also work without issues
				ResourceName:        name,
				ImportState:         true,
				ImportStateVerify:   true,
				ImportStateIdPrefix: fmt.Sprintf("%s/", accountID),
			},
			{
				// After import, re-apply should still show no changes
				Config: testPagesProjectDeploymentConfigsNoBuildConfig(rnd, accountID, projectName),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectEmptyPlan(),
					},
				},
			},
		},
	})
}

// TestAccCloudflarePagesProject_NoDriftWithMinimalConfig tests the fix for issue #5928.
// It verifies that a minimal config (without build_config or deployment_configs) does not
// show drift on subsequent plans.
func TestAccCloudflarePagesProject_NoDriftWithMinimalConfig(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	name := "cloudflare_pages_project." + rnd
	projectName := rnd
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflarePageProjectDestroy,
		Steps: []resource.TestStep{
			{
				// Initial apply with minimal config
				Config: testPagesProjectMinimal(rnd, accountID, projectName),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(name, tfjsonpath.New("name"), knownvalue.StringExact(projectName)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("production_branch"), knownvalue.StringExact("main")),
					// deployment_configs will be computed by the API
					statecheck.ExpectKnownValue(name, tfjsonpath.New("deployment_configs"), knownvalue.NotNull()),
				},
			},
			{
				// Re-apply the SAME config - should show no changes (fixes #5928)
				Config: testPagesProjectMinimal(rnd, accountID, projectName),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectEmptyPlan(),
					},
				},
			},
		},
	})
}
