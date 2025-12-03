package pages_project_test

import (
	"context"
	"fmt"
	"os"
	"regexp"
	"strings"
	"testing"

	"github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/cloudflare-go/v6/pages"
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

const resourcePrefix = "tfacctest-"

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
		// Skip sweeping if no account ID is set
		return nil
	}

	// List all pages projects in the account (this actually returns deployments)
	list, err := client.Pages.Projects.List(ctx, pages.ProjectListParams{
		AccountID: cloudflare.F(accountID),
	})
	if err != nil {
		return fmt.Errorf("failed to list pages projects: %w", err)
	}

	// Track unique project names to avoid duplicate deletions
	projectNames := make(map[string]bool)

	// Delete all pages projects with test prefix
	for _, deployment := range list.Result {
		if !strings.HasPrefix(deployment.Name, resourcePrefix) {
			continue
		}

		// Only delete each project once (deployments can have multiple entries per project)
		if !projectNames[deployment.Name] {
			projectNames[deployment.Name] = true
			_, err := client.Pages.Projects.Delete(ctx, deployment.Name, pages.ProjectDeleteParams{
				AccountID: cloudflare.F(accountID),
			})
			if err != nil {
				// Log but continue sweeping other projects
				continue
			}
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

func testPagesProjectFullConfig(resourceID, accountID, projectName, owner, repo string) string {
	return acctest.LoadTestCase("pagesprojectfullconfig.tf", resourceID, accountID, projectName, owner, repo)
}

func testPagesProjectEnvVars(resourceID, accountID, projectName string) string {
	return acctest.LoadTestCase("pagesprojectenvvars.tf", resourceID, accountID, projectName)
}

func testPagesProjectPreviewSettings(resourceID, accountID, projectName, owner, repo, setting, extraConfig string) string {
	return acctest.LoadTestCase("pagesprojectpreviewsettings.tf", resourceID, accountID, projectName, owner, repo, setting, extraConfig)
}

func testPagesProjectUpdated(resourceID, accountID, projectName string) string {
	return acctest.LoadTestCase("pagesprojectupdated.tf", resourceID, accountID, projectName)
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
	projectName := resourcePrefix + rnd
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
				Config:             testPagesProjectSource(rnd, accountID, projectName, pagesOwner, pagesRepo),
				ExpectNonEmptyPlan: true, // Drift from deployment_configs bindings and path_includes/path_excludes
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
			},
			{
				ResourceName:            name,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateIdPrefix:     fmt.Sprintf("%s/", accountID),
				ImportStateVerifyIgnore: []string{"build_config", "source.config.path_includes"},
			},
		},
	})
}

func TestAccCloudflarePagesProject_BuildConfig(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	name := "cloudflare_pages_project." + rnd
	projectName := resourcePrefix + rnd
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
				Config:             testPagesProjectBuildConfig(rnd, accountID, projectName),
				ExpectNonEmptyPlan: true, // Drift from deployment_configs bindings
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
	projectName := resourcePrefix + rnd
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
					"build_config",
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
	projectName := resourcePrefix + rnd
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflarePageProjectDestroy,
		Steps: []resource.TestStep{
			{
				Config:             testPagesProjectDirectUpload(rnd, accountID, projectName),
				ExpectNonEmptyPlan: true, // Drift from deployment_configs bindings
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
	projectName := resourcePrefix + rnd
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflarePageProjectDestroy,
		Steps: []resource.TestStep{
			{
				Config:             testPagesProjectMinimal(rnd, accountID, projectName),
				ExpectNonEmptyPlan: true, // Drift from deployment_configs bindings
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(name, tfjsonpath.New("name"), knownvalue.StringExact(projectName)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("production_branch"), knownvalue.StringExact("main")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("build_config"), knownvalue.Null()),
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
	projectName := resourcePrefix + rnd
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
				Config:             testPagesProjectMinimal(rnd, accountID, projectName),
				ExpectNonEmptyPlan: true, // Drift from deployment_configs bindings
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(name, plancheck.ResourceActionUpdate),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(name, tfjsonpath.New("name"), knownvalue.StringExact(projectName)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("production_branch"), knownvalue.StringExact("main")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("deployment_configs"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("deployment_configs").AtMapKey("preview").AtMapKey("env_vars"), knownvalue.MapSizeExact(0)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("deployment_configs").AtMapKey("production").AtMapKey("env_vars"), knownvalue.MapSizeExact(0)),
				},
			},
			{
				ResourceName:            name,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateIdPrefix:     fmt.Sprintf("%s/", accountID),
				ImportStateVerifyIgnore: []string{"build_config"},
			},
		},
	})
}
func TestAccCloudflarePagesProject_FullConfiguration(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	name := "cloudflare_pages_project." + rnd
	projectName := resourcePrefix + rnd
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
					"deployment_configs.preview.ai_bindings",
					"deployment_configs.preview.analytics_engine_datasets",
					"deployment_configs.preview.browsers",
					"deployment_configs.preview.hyperdrive_bindings",
					"deployment_configs.preview.mtls_certificates",
					"deployment_configs.preview.queue_producers",
					"deployment_configs.preview.services",
					"deployment_configs.preview.vectorize_bindings",
				},
			},
		},
	})
}

func TestAccCloudflarePagesProject_EnvVarTypes(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	name := "cloudflare_pages_project." + rnd
	projectName := resourcePrefix + rnd
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
				ImportStateVerifyIgnore: []string{"build_config", "deployment_configs.preview.env_vars", "deployment_configs.production.env_vars"},
			},
		},
	})
}

func TestAccCloudflarePagesProject_PreviewDeploymentSettings(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	name := "cloudflare_pages_project." + rnd
	projectName := resourcePrefix + rnd
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
				Config:             testPagesProjectPreviewSettings(rnd, accountID, projectName, pagesOwner, pagesRepo, "all", ""),
				ExpectNonEmptyPlan: true, // Drift from deployment_configs bindings and path_includes/path_excludes
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(name, tfjsonpath.New("name"), knownvalue.StringExact(projectName)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("source").AtMapKey("config").AtMapKey("preview_deployment_setting"), knownvalue.StringExact("all")),
				},
			},
			{
				Config:             testPagesProjectPreviewSettings(rnd, accountID, projectName, pagesOwner, pagesRepo, "none", ""),
				ExpectNonEmptyPlan: true, // Drift from deployment_configs bindings
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
				ExpectNonEmptyPlan: true, // Drift from deployment_configs bindings
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
				ResourceName:            name,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateIdPrefix:     fmt.Sprintf("%s/", accountID),
				ImportStateVerifyIgnore: []string{"build_config", "source.config.path_includes"},
			},
		},
	})
}

func TestAccCloudflarePagesProject_RemoveEnvVars(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	name := "cloudflare_pages_project." + rnd
	projectName := resourcePrefix + rnd
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflarePageProjectDestroy,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
resource "cloudflare_pages_project" "%s" {
	account_id = "%s"
	name = "%s"
	production_branch = "main"
	deployment_configs = {
		preview = {
			compatibility_date = "2023-06-01"
			compatibility_flags = []
			env_vars = {
				VAR_ONE = { type = "plain_text", value = "value1" }
				VAR_TWO = { type = "secret_text", value = "secret1" }
			}
		}
		production = {
			compatibility_date = "2023-06-01"
			compatibility_flags = []
			env_vars = {
				PROD_VAR = { type = "plain_text", value = "prodvalue" }
			}
		}
	}
}`, rnd, accountID, projectName),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(name, tfjsonpath.New("deployment_configs").AtMapKey("preview").AtMapKey("env_vars"), knownvalue.MapSizeExact(2)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("deployment_configs").AtMapKey("production").AtMapKey("env_vars"), knownvalue.MapSizeExact(1)),
				},
			},
			{
				Config: fmt.Sprintf(`
resource "cloudflare_pages_project" "%s" {
	account_id = "%s"
	name = "%s"
	production_branch = "main"
	deployment_configs = {
		preview = {
			compatibility_date = "2023-06-01"
			compatibility_flags = []
			env_vars = {
				VAR_ONE = { type = "plain_text", value = "value1" }
			}
		}
		production = {
			compatibility_date = "2023-06-01"
			compatibility_flags = []
			env_vars = {}
		}
	}
}`, rnd, accountID, projectName),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(name, tfjsonpath.New("deployment_configs").AtMapKey("preview").AtMapKey("env_vars"), knownvalue.MapSizeExact(1)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("deployment_configs").AtMapKey("production").AtMapKey("env_vars"), knownvalue.MapSizeExact(0)),
				},
			},
			{
				Config: fmt.Sprintf(`
resource "cloudflare_pages_project" "%s" {
	account_id = "%s"
	name = "%s"
	production_branch = "main"
	deployment_configs = {
		preview = {
			compatibility_date = "2023-06-01"
			compatibility_flags = []
			env_vars = {}
		}
		production = {
			compatibility_date = "2023-06-01"
			compatibility_flags = []
			env_vars = {}
		}
	}
}`, rnd, accountID, projectName),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(name, tfjsonpath.New("deployment_configs").AtMapKey("preview").AtMapKey("env_vars"), knownvalue.MapSizeExact(0)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("deployment_configs").AtMapKey("production").AtMapKey("env_vars"), knownvalue.MapSizeExact(0)),
				},
			},
		},
	})
}

func TestAccCloudflarePagesProject_RemoveBindings(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	name := "cloudflare_pages_project." + rnd
	projectName := resourcePrefix + rnd
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflarePageProjectDestroy,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
resource "cloudflare_workers_kv_namespace" "%[1]s_kv" {
	account_id = "%[2]s"
	title = "tfacctest-pages-bindings-kv"
}

resource "cloudflare_pages_project" "%[1]s" {
	account_id = "%[2]s"
	name = "%[3]s"
	production_branch = "main"
	deployment_configs = {
		preview = {
			compatibility_date = "2023-06-01"
			compatibility_flags = []
			kv_namespaces = {
				KV_BINDING_1 = { namespace_id = cloudflare_workers_kv_namespace.%[1]s_kv.id }
				KV_BINDING_2 = { namespace_id = cloudflare_workers_kv_namespace.%[1]s_kv.id }
			}
			r2_buckets = {
				R2_BINDING = { name = "test-bucket" }
			}
			d1_databases = {
				D1_BINDING = { id = "445e2955-951a-4358-a35b-a4d0c813f63" }
			}
		}
		production = {
			compatibility_date = "2023-06-01"
			compatibility_flags = []
			kv_namespaces = {
				KV_BINDING = { namespace_id = cloudflare_workers_kv_namespace.%[1]s_kv.id }
			}
			r2_buckets = {
				R2_BINDING_1 = { name = "bucket-one" }
				R2_BINDING_2 = { name = "bucket-two" }
			}
		}
	}
}`, rnd, accountID, projectName),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(name, tfjsonpath.New("deployment_configs").AtMapKey("preview").AtMapKey("kv_namespaces"), knownvalue.MapSizeExact(2)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("deployment_configs").AtMapKey("preview").AtMapKey("r2_buckets"), knownvalue.MapSizeExact(1)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("deployment_configs").AtMapKey("preview").AtMapKey("d1_databases"), knownvalue.MapSizeExact(1)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("deployment_configs").AtMapKey("production").AtMapKey("kv_namespaces"), knownvalue.MapSizeExact(1)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("deployment_configs").AtMapKey("production").AtMapKey("r2_buckets"), knownvalue.MapSizeExact(2)),
				},
			},
			{
				Config: fmt.Sprintf(`
resource "cloudflare_workers_kv_namespace" "%[1]s_kv" {
	account_id = "%[2]s"
	title = "tfacctest-pages-bindings-kv"
}

resource "cloudflare_pages_project" "%[1]s" {
	account_id = "%[2]s"
	name = "%[3]s"
	production_branch = "main"
	deployment_configs = {
		preview = {
			compatibility_date = "2023-06-01"
			compatibility_flags = []
			kv_namespaces = {
				KV_BINDING_1 = { namespace_id = cloudflare_workers_kv_namespace.%[1]s_kv.id }
			}
			r2_buckets = {}
			d1_databases = {}
		}
		production = {
			compatibility_date = "2023-06-01"
			compatibility_flags = []
			kv_namespaces = {}
			r2_buckets = {
				R2_BINDING_1 = { name = "bucket-one" }
			}
		}
	}
}`, rnd, accountID, projectName),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(name, tfjsonpath.New("deployment_configs").AtMapKey("preview").AtMapKey("kv_namespaces"), knownvalue.MapSizeExact(1)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("deployment_configs").AtMapKey("preview").AtMapKey("r2_buckets"), knownvalue.MapSizeExact(0)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("deployment_configs").AtMapKey("preview").AtMapKey("d1_databases"), knownvalue.MapSizeExact(0)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("deployment_configs").AtMapKey("production").AtMapKey("kv_namespaces"), knownvalue.MapSizeExact(0)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("deployment_configs").AtMapKey("production").AtMapKey("r2_buckets"), knownvalue.MapSizeExact(1)),
				},
			},
			{
				Config: fmt.Sprintf(`
resource "cloudflare_workers_kv_namespace" "%[1]s_kv" {
	account_id = "%[2]s"
	title = "tfacctest-pages-bindings-kv"
}

resource "cloudflare_pages_project" "%[1]s" {
	account_id = "%[2]s"
	name = "%[3]s"
	production_branch = "main"
	deployment_configs = {
		preview = {
			compatibility_date = "2023-06-01"
			compatibility_flags = []
			kv_namespaces = {}
			r2_buckets = {}
			d1_databases = {}
		}
		production = {
			compatibility_date = "2023-06-01"
			compatibility_flags = []
			kv_namespaces = {}
			r2_buckets = {}
		}
	}
}`, rnd, accountID, projectName),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(name, tfjsonpath.New("deployment_configs").AtMapKey("preview").AtMapKey("kv_namespaces"), knownvalue.MapSizeExact(0)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("deployment_configs").AtMapKey("preview").AtMapKey("r2_buckets"), knownvalue.MapSizeExact(0)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("deployment_configs").AtMapKey("preview").AtMapKey("d1_databases"), knownvalue.MapSizeExact(0)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("deployment_configs").AtMapKey("production").AtMapKey("kv_namespaces"), knownvalue.MapSizeExact(0)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("deployment_configs").AtMapKey("production").AtMapKey("r2_buckets"), knownvalue.MapSizeExact(0)),
				},
			},
		},
	})
}
