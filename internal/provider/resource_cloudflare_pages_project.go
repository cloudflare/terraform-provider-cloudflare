package provider

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"strings"
	"time"

	"github.com/MakeNowJust/heredoc/v2"
	"github.com/cloudflare/cloudflare-go"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCloudflarePagesProject() *schema.Resource {
	return &schema.Resource{
		Schema:        resourceCloudflarePagesProjectSchema(),
		CreateContext: resourceCloudflarePagesProjectCreate,
		ReadContext:   resourceCloudflarePagesProjectRead,
		UpdateContext: resourceCloudflarePagesProjectUpdate,
		DeleteContext: resourceCloudflarePagesProjectDelete,
		Importer: &schema.ResourceImporter{
			StateContext: resourceCloudflarePagesProjectImport,
		},
		Description: heredoc.Doc(`
			Provides a resource which manages Cloudflare Pages projects.
		`),
	}
}

func buildDeploymentConfig(environment interface{}) cloudflare.PagesProjectDeploymentConfigEnvironment {
	config := cloudflare.PagesProjectDeploymentConfigEnvironment{}
	parsed := environment.(map[string]interface{})
	for key, value := range parsed {
		switch key {
		case "environment_variables":
			deploymentVariables := make(map[string]cloudflare.PagesProjectDeploymentVar)
			variables := value.(map[string]interface{})
			for i, variable := range variables {
				deploymentVariables[i] = cloudflare.PagesProjectDeploymentVar{Value: variable.(string)}
			}
			config.EnvVars = deploymentVariables
			break
		case "kv_namespaces":
			namespace := cloudflare.NamespaceBindingMap{}
			variables := value.(map[string]interface{})
			for i, variable := range variables {
				namespace[i] = &cloudflare.NamespaceBindingValue{Value: variable.(string)}
			}
			config.KvNamespaces = namespace
			break
		case "durable_object_namespaces":
			namespace := cloudflare.NamespaceBindingMap{}
			variables := value.(map[string]interface{})
			for i, variable := range variables {
				namespace[i] = &cloudflare.NamespaceBindingValue{Value: variable.(string)}
			}
			config.DoNamespaces = namespace
			break
		case "d1_databases":
			bindingMap := cloudflare.D1BindingMap{}
			variables := value.(map[string]interface{})
			for i, variable := range variables {
				bindingMap[i] = &cloudflare.D1Binding{ID: variable.(string)}
			}
			config.D1Databases = bindingMap
			break
		case "r2_buckets":
			bindingMap := cloudflare.R2BindingMap{}
			variables := value.(map[string]interface{})
			for i, variable := range variables {
				bindingMap[i] = &cloudflare.R2BindingValue{Name: variable.(string)}
			}
			config.R2Bindings = bindingMap
			break
		case "compatibility_date":
			config.CompatibilityDate = value.(string)
			break
		case "compatibility_flags":
			for _, item := range value.([]interface{}) {
				config.CompatibilityFlags = append(config.CompatibilityFlags, item.(string))
			}
			break
		}
	}

	return config
}

func buildPagesProject(d *schema.ResourceData) cloudflare.PagesProject {
	project := cloudflare.PagesProject{}
	project.Name = d.Get("name").(string)
	project.ProductionBranch = d.Get("production_branch").(string)

	if _, ok := d.GetOk("build_config"); ok {
		buildConfig := cloudflare.PagesProjectBuildConfig{}
		if buildCommand, ok := d.GetOk("build_config.0.build_command"); ok {
			buildConfig.BuildCommand = buildCommand.(string)
		}
		if destinationDir, ok := d.GetOk("build_config.0.destination_dir"); ok {
			buildConfig.DestinationDir = destinationDir.(string)
		}
		if rootDir, ok := d.GetOk("build_config.0.root_dir"); ok {
			buildConfig.RootDir = rootDir.(string)
		}
		if webAnalyticsTag, ok := d.GetOk("build_config.0.web_analytics_tag"); ok {
			buildConfig.WebAnalyticsTag = webAnalyticsTag.(string)
		}
		if webAnalyticsToken, ok := d.GetOk("build_config.0.web_analytics_tag"); ok {
			buildConfig.WebAnalyticsToken = webAnalyticsToken.(string)
		}
		project.BuildConfig = buildConfig
	}

	source := cloudflare.PagesProjectSource{}
	if _, ok := d.GetOk("source"); ok {
		if sourceType, ok := d.GetOk("source.0.type"); ok {
			source.Type = sourceType.(string)
		}
		if _, ok := d.GetOk("source.0.config"); ok {
			sourceConfig := cloudflare.PagesProjectSourceConfig{}
			if sourceOwner, ok := d.GetOk("source.0.config.0.owner"); ok {
				sourceConfig.Owner = sourceOwner.(string)
			}
			if sourceRepoName, ok := d.GetOk("source.0.config.0.repo_name"); ok {
				sourceConfig.RepoName = sourceRepoName.(string)
			}
			if sourceProductionBranch, ok := d.GetOk("source.0.config.0.production_branch"); ok {
				sourceConfig.ProductionBranch = sourceProductionBranch.(string)
			}
			if sourcePRComments, ok := d.GetOk("source.0.config.0.pr_comments_enabled"); ok {
				sourceConfig.PRCommentsEnabled = sourcePRComments.(bool)
			}
			if sourceDeploymentsEnabled, ok := d.GetOk("source.0.config.0.deployments_enabled"); ok {
				sourceConfig.DeploymentsEnabled = sourceDeploymentsEnabled.(bool)
			}
			if productionBranch, ok := d.GetOk("source.0.config.0.production_branch"); ok {
				sourceConfig.ProductionBranch = productionBranch.(string)
			}
			if previewBranchIncludes, ok := d.GetOk("source.0.config.0.preview_branch_includes"); ok {
				for _, item := range previewBranchIncludes.([]interface{}) {
					sourceConfig.PreviewBranchIncludes = append(sourceConfig.PreviewBranchIncludes, item.(string))
				}
			}
			if previewBranchExcludes, ok := d.GetOk("source.0.config.0.preview_branch_excludes"); ok {
				for _, item := range previewBranchExcludes.([]interface{}) {
					sourceConfig.PreviewBranchExcludes = append(sourceConfig.PreviewBranchExcludes, item.(string))
				}
			}
			if previewDeploymentSetting, ok := d.GetOk("source.0.config.0.preview_deployment_setting"); ok {
				sourceConfig.PreviewDeploymentSetting = cloudflare.PagesPreviewDeploymentSetting(previewDeploymentSetting.(string))
			}
			source.Config = &sourceConfig
		}
	}

	if previewConfig, ok := d.GetOk("deployment_configs.0.preview.0"); ok {
		project.DeploymentConfigs.Preview = buildDeploymentConfig(previewConfig)
	}
	if productionConfig, ok := d.GetOk("deployment_configs.0.production.0"); ok {
		project.DeploymentConfigs.Production = buildDeploymentConfig(productionConfig)
	}

	return project
}

func resourceCloudflarePagesProjectRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	accountID := d.Get("account_id").(string)
	projectName := d.Get("name").(string)

	res, err := client.PagesProject(ctx, accountID, projectName)
	if err != nil {
		return diag.FromErr(fmt.Errorf("error reading cloudflare pages project %q: %w", projectName, err))
	}

	d.SetId(res.Name)
	d.Set("subdomain", res.SubDomain)
	d.Set("created_on", res.CreatedOn.Format(time.RFC3339))
	d.Set("domains", res.Domains)

	return nil
}

func resourceCloudflarePagesProjectCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	accountID := d.Get("account_id").(string)
	pageProject := buildPagesProject(d)

	_, err := client.CreatePagesProject(ctx, accountID, pageProject)
	if err != nil {
		return diag.FromErr(fmt.Errorf("error creating cloudflare pages project %q: %w", pageProject.Name, err))
	}

	return resourceCloudflarePagesProjectRead(ctx, d, meta)
}

func resourceCloudflarePagesProjectUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	accountID := d.Get("account_id").(string)
	projectName := d.Get("name").(string)

	pageProject := buildPagesProject(d)

	_, err := client.UpdatePagesProject(ctx, accountID, projectName, pageProject)
	if err != nil {
		return diag.FromErr(fmt.Errorf("error updating cloudflare pages project %q: %w", pageProject.Name, err))
	}

	return resourceCloudflarePagesProjectRead(ctx, d, meta)
}

func resourceCloudflarePagesProjectDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	accountID := d.Get("account_id").(string)
	projectName := d.Get("name").(string)

	err := client.DeletePagesProject(ctx, accountID, projectName)
	if err != nil {
		return diag.FromErr(fmt.Errorf("error deleting cloudflare pages project %q: %w", projectName, err))
	}
	return nil
}

func resourceCloudflarePagesProjectImport(ctx context.Context, d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	client := meta.(*cloudflare.API)

	// split the id so we can look up
	idAttr := strings.SplitN(d.Id(), "/", 2)
	var accountID string
	var projectName string
	if len(idAttr) == 2 {
		accountID = idAttr[0]
		projectName = idAttr[1]
	} else {
		return nil, fmt.Errorf("invalid id %q specified, should be in format \"accountID/project-name\" for import", d.Id())
	}

	project, err := client.PagesProject(ctx, accountID, projectName)
	if err != nil {
		return nil, fmt.Errorf("Unable to find record with ID %q: %w", d.Id(), err)
	}

	tflog.Info(ctx, fmt.Sprintf("Found project: %s", project.Name))

	d.SetId(project.Name)
	d.Set("name", project.Name)
	d.Set("account_id", accountID)
	d.Set("production_branch", project.ProductionBranch)

	resourceCloudflarePagesProjectRead(ctx, d, meta)

	return []*schema.ResourceData{d}, nil
}
