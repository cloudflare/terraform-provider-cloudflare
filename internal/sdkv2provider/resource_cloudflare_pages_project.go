package sdkv2provider

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/MakeNowJust/heredoc/v2"
	"github.com/cloudflare/cloudflare-go"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
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
	deploymentVariables := cloudflare.EnvironmentVariableMap{}
	for key, value := range parsed {
		switch key {
		case "environment_variables":
			variables := value.(map[string]interface{})
			for i, variable := range variables {
				envVar := cloudflare.EnvironmentVariable{
					Value: variable.(string),
					Type:  cloudflare.PlainText,
				}
				deploymentVariables[i] = &envVar
			}

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
		case "fail_open":
			config.FailOpen = value.(bool)
			break
		case "always_use_latest_compatibility_date":
			config.AlwaysUseLatestCompatibilityDate = value.(bool)
			break
		case "usage_model":
			config.UsageModel = cloudflare.UsageModel(value.(string))
			break
		case "service_binding":
			serviceMap := cloudflare.ServiceBindingMap{}
			for _, item := range value.(*schema.Set).List() {
				data := item.(map[string]interface{})
				serviceMap[data["name"].(string)] = &cloudflare.ServiceBinding{
					Service:     data["service"].(string),
					Environment: data["environment"].(string),
				}
			}
			config.ServiceBindings = serviceMap
			break
		}
	}
	config.EnvVars = deploymentVariables
	return config
}

func parseDeploymentConfig(deployment cloudflare.PagesProjectDeploymentConfigEnvironment) (returnValue []map[string]interface{}) {
	config := make(map[string]interface{})

	config["compatibility_date"] = deployment.CompatibilityDate
	config["compatibility_flags"] = deployment.CompatibilityFlags

	config["fail_open"] = deployment.FailOpen
	config["always_use_latest_compatibility_date"] = deployment.AlwaysUseLatestCompatibilityDate
	config["usage_model"] = deployment.UsageModel

	deploymentVars := map[string]string{}
	for key, value := range deployment.EnvVars {
		if value.Type == cloudflare.PlainText {
			deploymentVars[key] = value.Value
		}
	}
	config["environment_variables"] = deploymentVars

	deploymentVars = map[string]string{}
	for key, value := range deployment.KvNamespaces {
		deploymentVars[key] = value.Value
	}
	config["kv_namespaces"] = deploymentVars

	deploymentVars = map[string]string{}
	for key, value := range deployment.DoNamespaces {
		deploymentVars[key] = value.Value
	}
	config["durable_object_namespaces"] = deploymentVars

	deploymentVars = map[string]string{}
	for key, value := range deployment.R2Bindings {
		deploymentVars[key] = value.Name
	}
	config["r2_buckets"] = deploymentVars

	deploymentVars = map[string]string{}
	for key, value := range deployment.D1Databases {
		deploymentVars[key] = value.ID
	}
	config["d1_databases"] = deploymentVars

	serviceBindings := &schema.Set{F: schema.HashResource(serviceBindingResource)}
	for key, value := range deployment.ServiceBindings {
		serviceBindings.Add(map[string]interface{}{
			"name":        key,
			"service":     value.Service,
			"environment": value.Environment,
		})
	}
	config["service_binding"] = serviceBindings

	returnValue = append(returnValue, config)
	return
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
		if webAnalyticsToken, ok := d.GetOk("build_config.0.web_analytics_token"); ok {
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
			if productionBranchEnable, ok := d.GetOk("source.0.config.0.production_deployment_enabled"); ok {
				sourceConfig.ProductionDeploymentsEnabled = productionBranchEnable.(bool)
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
		project.Source = &source
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
	accountID := d.Get(consts.AccountIDSchemaKey).(string)

	project, err := client.PagesProject(ctx, accountID, d.Id())
	if err != nil {
		return diag.FromErr(fmt.Errorf("error reading cloudflare pages project %q: %w", d.Id(), err))
	}

	tflog.Debug(ctx, fmt.Sprintf("Cloudflare Pages Project Response: %#v", project))
	d.Set("subdomain", project.SubDomain)
	d.Set("production_branch", project.ProductionBranch)
	d.Set(consts.AccountIDSchemaKey, accountID)
	d.Set("domains", project.Domains)
	d.Set("created_on", project.CreatedOn.Format(time.RFC3339))

	if project.Source != nil {
		var source []map[string]interface{}
		source = append(source, map[string]interface{}{
			"type": project.Source.Type,
			"config": []map[string]interface{}{
				{
					"owner":                         project.Source.Config.Owner,
					"repo_name":                     project.Source.Config.RepoName,
					"production_branch":             project.Source.Config.ProductionBranch,
					"pr_comments_enabled":           project.Source.Config.PRCommentsEnabled,
					"deployments_enabled":           project.Source.Config.DeploymentsEnabled,
					"production_deployment_enabled": project.Source.Config.ProductionDeploymentsEnabled,
					"preview_branch_includes":       project.Source.Config.PreviewBranchIncludes,
					"preview_branch_excludes":       project.Source.Config.PreviewBranchExcludes,
					"preview_deployment_setting":    project.Source.Config.PreviewDeploymentSetting,
				},
			},
		},
		)
		d.Set("source", source)
	}
	emptyProjectBuildConfig := cloudflare.PagesProjectBuildConfig{}
	if project.BuildConfig != emptyProjectBuildConfig {
		var buildConfig []map[string]interface{}
		buildConfig = append(buildConfig, map[string]interface{}{
			"build_command":       project.BuildConfig.BuildCommand,
			"destination_dir":     project.BuildConfig.DestinationDir,
			"root_dir":            project.BuildConfig.RootDir,
			"web_analytics_tag":   project.BuildConfig.WebAnalyticsTag,
			"web_analytics_token": project.BuildConfig.WebAnalyticsToken,
		},
		)
		d.Set("build_config", buildConfig)
	}

	var deploymentConfigs []map[string]interface{}
	deploymentConfig := make(map[string]interface{})
	deploymentConfig["preview"] = parseDeploymentConfig(project.DeploymentConfigs.Preview)
	deploymentConfig["production"] = parseDeploymentConfig(project.DeploymentConfigs.Production)
	deploymentConfigs = append(deploymentConfigs, deploymentConfig)
	d.Set("deployment_configs", deploymentConfigs)

	return nil
}

func resourceCloudflarePagesProjectCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	accountID := d.Get(consts.AccountIDSchemaKey).(string)
	pageProject := buildPagesProject(d)

	project, err := client.CreatePagesProject(ctx, accountID, pageProject)
	if err != nil {
		return diag.FromErr(fmt.Errorf("error creating cloudflare pages project %q: %w", pageProject.Name, err))
	}

	d.SetId(project.Name)
	return resourceCloudflarePagesProjectRead(ctx, d, meta)
}

func resourceCloudflarePagesProjectUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	accountID := d.Get(consts.AccountIDSchemaKey).(string)

	pageProject := buildPagesProject(d)

	_, err := client.UpdatePagesProject(ctx, accountID, d.Id(), pageProject)
	if err != nil {
		return diag.FromErr(fmt.Errorf("error updating cloudflare pages project %q: %w", d.Id(), err))
	}

	return resourceCloudflarePagesProjectRead(ctx, d, meta)
}

func resourceCloudflarePagesProjectDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	accountID := d.Get(consts.AccountIDSchemaKey).(string)

	err := client.DeletePagesProject(ctx, accountID, d.Id())
	if err != nil {
		return diag.FromErr(fmt.Errorf("error deleting cloudflare pages project %q: %w", d.Id(), err))
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
	d.Set(consts.AccountIDSchemaKey, accountID)

	resourceCloudflarePagesProjectRead(ctx, d, meta)

	return []*schema.ResourceData{d}, nil
}
