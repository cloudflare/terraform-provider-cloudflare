package provider

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCloudflarePagesProjectSchema() map[string]*schema.Schema {
	buildConfig := schema.Resource{
		Schema: map[string]*schema.Schema{
			"build_command": {
				Type:        schema.TypeString,
				Description: "Command used to build project.",
				Optional:    true,
			},
			"destination_dir": {
				Type:        schema.TypeString,
				Description: "Output directory of the build.",
				Optional:    true,
			},
			"root_dir": {
				Type:        schema.TypeString,
				Description: "Directory to run the command.",
				Optional:    true,
			},
			"web_analytics_tag": {
				Type:        schema.TypeString,
				Description: "The classifying tag for analytics.",
				Optional:    true,
			},
			"web_analytics_token": {
				Type:        schema.TypeString,
				Description: "The auth token for analytics.",
				Optional:    true,
			},
		},
	}

	source := schema.Resource{
		Schema: map[string]*schema.Schema{
			"type": {
				Type:        schema.TypeString,
				Description: "Project host type.",
				Optional:    true,
			},
			"config": {
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"owner": {
							Description: "Project owner username.",
							Type:        schema.TypeString,
							Required:    true,
						},
						"repo_name": {
							Description: "Project repository name.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"production_branch": {
							Description: "Project production branch name.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"pr_comments_enabled": {
							Description: "Enable Pages to comment on Pull Requests.",
							Type:        schema.TypeBool,
							Optional:    true,
							Computed:    true,
						},
						"deployments_enabled": {
							Description: "Toggle deployments on this repo.",
							Type:        schema.TypeBool,
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
		},
	}

	deploymentConfig := schema.Resource{
		Schema: map[string]*schema.Schema{
			"environment_variables": {
				Type:        schema.TypeMap,
				Description: "Environment variables for build configs.",
				Optional:    true,
			},
			"compatibility_date": {
				Type:        schema.TypeString,
				Description: "Compatibility date used for Pages Functions.",
				Optional:    true,
				Computed:    true,
			},
			"compatibility_flags": {
				Type:        schema.TypeList,
				Description: "Compatibility flags used for Pages Functions.",
				Optional:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Computed: true,
			},
		},
	}

	return map[string]*schema.Schema{
		"account_id": {
			Description: "The account identifier to target for the resource.",
			Type:        schema.TypeString,
			Required:    true,
		},
		"name": {
			Description: "Name of the project.",
			Type:        schema.TypeString,
			Required:    true,
		},
		"id": {
			Description: "ID of the project.",
			Type:        schema.TypeString,
			Computed:    true,
		},
		"subdomain": {
			Description: "The Cloudflare subdomain associated with the project.",
			Type:        schema.TypeString,
			Computed:    true,
		},
		"domains": {
			Description: "A list of associated custom domains for the project.",
			Type:        schema.TypeList,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			Computed: true,
		},
		"created_on": {
			Type:        schema.TypeString,
			Description: "When the project was created.",
			Computed:    true,
		},
		"build_config": {
			Description: "Configs for the project build process.",
			Type:        schema.TypeList,
			Elem:        &buildConfig,
			MaxItems:    1,
			Optional:    true,
		},
		"source": {
			Description: "Configs for the project source.",
			Optional:    true,
			Type:        schema.TypeList,
			Elem:        &source,
			MaxItems:    1,
		},
		"deployment_configs": {
			Description: "Configs for deployments in a project.",
			Type:        schema.TypeList,
			MaxItems:    1,
			Optional:    true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"preview": {
						Description: "Configs for preview deploys.",
						Type:        schema.TypeList,
						Optional:    true,
						Elem:        &deploymentConfig,
						MaxItems:    1,
					},
					"production": {
						Description: "Configs for production deploys.",
						Type:        schema.TypeList,
						Optional:    true,
						Elem:        &deploymentConfig,
						MaxItems:    1,
					},
				},
			},
		},
	}
}
