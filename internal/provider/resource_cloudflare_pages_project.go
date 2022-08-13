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
		buildConfig = cloudflare.PagesProjectBuildConfig{
			BuildCommand:      d.Get("build_config.0.build_command").(string),
			DestinationDir:    d.Get("build_config.0.destination_dir").(string),
			RootDir:           d.Get("build_config.0.root_dir").(string),
			WebAnalyticsTag:   d.Get("build_config.0.web_analytics_tag").(string),
			WebAnalyticsToken: d.Get("build_config.0.web_analytics_tag").(string),
		}
	}

	return cloudflare.PagesProject{
		Name:        name,
		BuildConfig: buildConfig,
	}
}

func resourceCloudflarePagesProjectRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	accountID := d.Get("account_id").(string)
	projectName := d.Get("name").(string)

	pageProject := buildPagesProject(d)

	res, err := client.PagesProject(ctx, accountID, projectName)
	if err != nil {
		return diag.FromErr(fmt.Errorf("error reading cloudflare pages project %q: %w", pageProject.Name, err))
	}

	d.SetId(res.ID)
	d.Set("subdomain", res.SubDomain)
	d.Set("created_on", res.CreatedOn)
	d.Set("domains", res.Domains)
	d.Set("source.0.type", res.Source.Type)
	d.Set("source.0.config.owner", res.Source.Config.Owner)
	d.Set("source.0.config.repo_name", res.Source.Config.RepoName)
	d.Set("source.0.config.production_branch", res.Source.Config.ProductionBranch)
	d.Set("source.0.config.pr_comments_enabled", res.Source.Config.PRCommentsEnabled)
	d.Set("source.0.config.deployments_enabled", res.Source.Config.DeploymentsEnabled)

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

	pageProject := buildPagesProject(d)

	_, err := client.CreatePagesProject(ctx, accountID, pageProject)
	if err != nil {
		return diag.FromErr(fmt.Errorf("error updating cloudflare pages project %q: %w", pageProject.Name, err))
	}

	return resourceCloudflarePagesProjectRead(ctx, d, meta)
}

func resourceCloudflarePagesProjectDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	accountID := d.Get("account_id").(string)
	projectName := d.Get("project_name").(string)

	err := client.DeletePagesProject(ctx, accountID, projectName)
	if err != nil {
		return diag.FromErr(fmt.Errorf("error deleting cloudflare pages project %q: %w", projectName, err))
	}
	return nil
}
