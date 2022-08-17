package provider

import (
	"context"
	"fmt"

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
		Description: heredoc.Doc(`
			Provides a resource which manages Cloudflare Pages projects.
		`),
	}
}

func buildPagesProject(d *schema.ResourceData) cloudflare.PagesProject {
	name := d.Get("name").(string)

	buildConfig := cloudflare.PagesProjectBuildConfig{}
	if _, ok := d.GetOk("build_config"); ok {
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
			if sourceProducationBranch, ok := d.GetOk("source.0.config.0.production_branch"); ok {
				sourceConfig.ProductionBranch = sourceProducationBranch.(string)
			}
			if sourcePRComments, ok := d.GetOk("source.0.config.0.pr_comments_enabled"); ok {
				sourceConfig.PRCommentsEnabled = sourcePRComments.(bool)
			}
			if sourceDeploymentsEnabled, ok := d.GetOk("source.0.config.0.deployments_enabled"); ok {
				sourceConfig.DeploymentsEnabled = sourceDeploymentsEnabled.(bool)
			}
			source.Config = &sourceConfig
		}
	}

	// I hate everything
	deploymentVariables := cloudflare.PagesProjectDeploymentConfigs{}
	if previewVars, ok := d.GetOk("deployment_configs.0.preview.0.environment_variables"); ok {
		preview := cloudflare.PagesProjectDeploymentConfigEnvironment{}
		variables := previewVars.(map[string]interface{})
		for i, variable := range variables {
			value := variable.(map[string]interface{})["value"]
			deploymentVar := cloudflare.PagesProjectDeploymentVar{Value: value.(string)}
			preview.EnvVars[i] = deploymentVar
		}
		deploymentVariables.Preview = preview
	}
	if productionVars, ok := d.GetOk("deployment_configs.0.production.0.environment_variables"); ok {
		production := cloudflare.PagesProjectDeploymentConfigEnvironment{}
		variables := productionVars.(map[string]interface{})
		for i, variable := range variables {
			value := variable.(map[string]interface{})["value"]
			deploymentVar := cloudflare.PagesProjectDeploymentVar{Value: value.(string)}
			production.EnvVars[i] = deploymentVar
		}
		deploymentVariables.Production = production
	}

	return cloudflare.PagesProject{
		Name:              name,
		BuildConfig:       buildConfig,
		DeploymentConfigs: deploymentVariables,
		Source: source,
	}
}

func resourceCloudflarePagesProjectRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	accountID := d.Get("account_id").(string)
	projectName := d.Get("name").(string)

	res, err := client.PagesProject(ctx, accountID, projectName)
	if err != nil {
		return diag.FromErr(fmt.Errorf("error reading cloudflare pages project %q: %w", projectName, err))
	}

	d.SetId(res.ID)
	d.Set("subdomain", res.SubDomain)
	d.Set("created_on", res.CreatedOn.String())
	d.Set("domains", res.Domains)
	d.Set("source", res.Source)

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
