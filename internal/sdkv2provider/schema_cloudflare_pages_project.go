package sdkv2provider

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
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
				Description: "Your project's root directory, where Cloudflare runs the build command. If your site is not in a subdirectory, leave this path value empty.",
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
				Description: "Configuration for the source of the Cloudflare Pages project.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"owner": {
							Description: "Project owner username.",
							Type:        schema.TypeString,
							Optional:    true,
							ForceNew:    true,
						},
						"repo_name": {
							Description: "Project repository name.",
							Type:        schema.TypeString,
							Optional:    true,
							ForceNew:    true,
						},
						"production_branch": {
							Description: "Project production branch name.",
							Type:        schema.TypeString,
							Required:    true,
						},
						"pr_comments_enabled": {
							Description: "Enable Pages to comment on Pull Requests.",
							Type:        schema.TypeBool,
							Optional:    true,
							Default:     true,
						},
						"deployments_enabled": {
							Description: "Toggle deployments on this repo.",
							Type:        schema.TypeBool,
							Optional:    true,
							Default:     true,
						},
						"production_deployment_enabled": {
							Description: "Enable production deployments.",
							Type:        schema.TypeBool,
							Optional:    true,
							Default:     true,
						},
						"preview_branch_includes": {
							Description: "Branches will be included for automatic deployment.",
							Type:        schema.TypeList,
							Optional:    true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"preview_branch_excludes": {
							Description: "Branches will be excluded from automatic deployment.",
							Type:        schema.TypeList,
							Optional:    true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"preview_deployment_setting": {
							Description:  "Preview Deployment Setting.",
							Type:         schema.TypeString,
							Optional:     true,
							ValidateFunc: validation.StringInSlice([]string{"custom", "all", "none"}, false),
							Default:      "all",
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
				Description: "Environment variables for Pages Functions.",
				Optional:    true,
				Default:     map[string]interface{}{},
			},
			"secrets": {
				Type:        schema.TypeMap,
				Description: "Encrypted environment variables for Pages Functions.",
				Optional:    true,
				Sensitive:   true,
				Default:     map[string]interface{}{},
			},
			"kv_namespaces": {
				Type:        schema.TypeMap,
				Description: "KV namespaces used for Pages Functions.",
				Optional:    true,
				Default:     map[string]interface{}{},
			},
			"durable_object_namespaces": {
				Type:        schema.TypeMap,
				Description: "Durable Object namespaces used for Pages Functions.",
				Optional:    true,
				Default:     map[string]interface{}{},
			},
			"d1_databases": {
				Type:        schema.TypeMap,
				Description: "D1 Databases used for Pages Functions.",
				Optional:    true,
				Default:     map[string]interface{}{},
			},
			"r2_buckets": {
				Type:        schema.TypeMap,
				Description: "R2 Buckets used for Pages Functions.",
				Optional:    true,
				Default:     map[string]interface{}{},
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
			"service_binding": {
				Type:        schema.TypeSet,
				Description: "Services used for Pages Functions.",
				Optional:    true,
				Elem:        serviceBindingResource,
			},
			"fail_open": {
				Type:        schema.TypeBool,
				Description: "Fail open used for Pages Functions.",
				Optional:    true,
				Default:     false,
			},
			"always_use_latest_compatibility_date": {
				Type:        schema.TypeBool,
				Description: "Use latest compatibility date for Pages Functions.",
				Optional:    true,
				Default:     false,
			},
			"usage_model": {
				Type:         schema.TypeString,
				Description:  "Usage model used for Pages Functions.",
				Optional:     true,
				ValidateFunc: validation.StringInSlice([]string{"unbound", "bundled"}, false),
				Default:      "bundled",
			},
			"placement": {
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Description: "Configuration for placement in the Cloudflare Pages project.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"mode": {
							Description: "Placement Mode for the Pages Function.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
		},
	}

	return map[string]*schema.Schema{
		consts.AccountIDSchemaKey: {
			Description: consts.AccountIDSchemaDescription,
			Type:        schema.TypeString,
			Required:    true,
		},
		"name": {
			Description: "Name of the project.",
			Type:        schema.TypeString,
			Required:    true,
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
		"production_branch": {
			Type:        schema.TypeString,
			Description: "The name of the branch that is used for the production environment.",
			Required:    true,
		},
		"build_config": {
			Description: "Configuration for the project build process. Read more about the build configuration in the [developer documentation](https://developers.cloudflare.com/pages/platform/build-configuration)",
			Type:        schema.TypeList,
			Elem:        &buildConfig,
			MaxItems:    1,
			Optional:    true,
		},
		"source": {
			Description: "Configuration for the project source. Read more about the source configuration in the [developer documentation](https://developers.cloudflare.com/pages/platform/branch-build-controls/)",
			Optional:    true,
			Type:        schema.TypeList,
			Elem:        &source,
			MaxItems:    1,
		},
		"deployment_configs": {
			Description: "Configuration for deployments in a project.",
			Type:        schema.TypeList,
			MaxItems:    1,
			Computed:    true,
			Optional:    true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"preview": {
						Description: "Configuration for preview deploys.",
						Type:        schema.TypeList,
						Computed:    true,
						Optional:    true,
						Elem:        &deploymentConfig,
						MaxItems:    1,
					},
					"production": {
						Description: "Configuration for production deploys.",
						Type:        schema.TypeList,
						Computed:    true,
						Optional:    true,
						Elem:        &deploymentConfig,
						MaxItems:    1,
					},
				},
			},
		},
	}
}
