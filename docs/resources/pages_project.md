---
page_title: "cloudflare_pages_project Resource - Cloudflare"
subcategory: ""
description: |-
  
---

# cloudflare_pages_project (Resource)



-> If you are using a `source` block configuration, you must first have a
   connected GitHub or GitLab account connected to Cloudflare. See the
   [Getting Started with Pages](https://developers.cloudflare.com/pages/get-started/git-integration/)
   documentation on how to link your accounts.

## Example Usage

```terraform
resource "cloudflare_pages_project" "example_pages_project" {
  account_id = "023e105f4ecef8ad9ca31a8372d0c353"
  name = "my-pages-app"
  production_branch = "main"
  build_config = {
    build_caching = true
    build_command = "npm run build"
    destination_dir = "build"
    root_dir = "/"
    web_analytics_tag = "cee1c73f6e4743d0b5e6bb1a0bcaabcc"
    web_analytics_token = "021e1057c18547eca7b79f2516f06o7x"
  }
  deployment_configs = {
    preview = {
      ai_bindings = {
        AI_BINDING = {
          project_id = "some-project-id"
        }
      }
      always_use_latest_compatibility_date = false
      analytics_engine_datasets = {
        ANALYTICS_ENGINE_BINDING = {
          dataset = "api_analytics"
        }
      }
      browsers = {
        BROWSER = {

        }
      }
      build_image_major_version = 3
      compatibility_date = "2025-01-01"
      compatibility_flags = ["url_standard"]
      d1_databases = {
        D1_BINDING = {
          id = "445e2955-951a-43f8-a35b-a4d0c8138f63"
        }
      }
      durable_object_namespaces = {
        DO_BINDING = {
          namespace_id = "5eb63bbbe01eeed093cb22bb8f5acdc3"
        }
      }
      env_vars = {
        foo = {
          type = "plain_text"
          value = "hello world"
        }
      }
      fail_open = true
      hyperdrive_bindings = {
        HYPERDRIVE = {
          id = "a76a99bc342644deb02c38d66082262a"
        }
      }
      kv_namespaces = {
        KV_BINDING = {
          namespace_id = "5eb63bbbe01eeed093cb22bb8f5acdc3"
        }
      }
      limits = {
        cpu_ms = 100
      }
      mtls_certificates = {
        MTLS = {
          certificate_id = "d7cdd17c-916f-4cb7-aabe-585eb382ec4e"
        }
      }
      placement = {
        mode = "smart"
      }
      queue_producers = {
        QUEUE_PRODUCER_BINDING = {
          name = "some-queue"
        }
      }
      r2_buckets = {
        R2_BINDING = {
          name = "some-bucket"
          jurisdiction = "eu"
        }
      }
      services = {
        SERVICE_BINDING = {
          service = "example-worker"
          entrypoint = "MyHandler"
          environment = "production"
        }
      }
      usage_model = "standard"
      vectorize_bindings = {
        VECTORIZE = {
          index_name = "my_index"
        }
      }
      wrangler_config_hash = "abc123def456"
    }
    production = {
      ai_bindings = {
        AI_BINDING = {
          project_id = "some-project-id"
        }
      }
      always_use_latest_compatibility_date = false
      analytics_engine_datasets = {
        ANALYTICS_ENGINE_BINDING = {
          dataset = "api_analytics"
        }
      }
      browsers = {
        BROWSER = {

        }
      }
      build_image_major_version = 3
      compatibility_date = "2025-01-01"
      compatibility_flags = ["url_standard"]
      d1_databases = {
        D1_BINDING = {
          id = "445e2955-951a-43f8-a35b-a4d0c8138f63"
        }
      }
      durable_object_namespaces = {
        DO_BINDING = {
          namespace_id = "5eb63bbbe01eeed093cb22bb8f5acdc3"
        }
      }
      env_vars = {
        foo = {
          type = "plain_text"
          value = "hello world"
        }
      }
      fail_open = true
      hyperdrive_bindings = {
        HYPERDRIVE = {
          id = "a76a99bc342644deb02c38d66082262a"
        }
      }
      kv_namespaces = {
        KV_BINDING = {
          namespace_id = "5eb63bbbe01eeed093cb22bb8f5acdc3"
        }
      }
      limits = {
        cpu_ms = 100
      }
      mtls_certificates = {
        MTLS = {
          certificate_id = "d7cdd17c-916f-4cb7-aabe-585eb382ec4e"
        }
      }
      placement = {
        mode = "smart"
      }
      queue_producers = {
        QUEUE_PRODUCER_BINDING = {
          name = "some-queue"
        }
      }
      r2_buckets = {
        R2_BINDING = {
          name = "some-bucket"
          jurisdiction = "eu"
        }
      }
      services = {
        SERVICE_BINDING = {
          service = "example-worker"
          entrypoint = "MyHandler"
          environment = "production"
        }
      }
      usage_model = "standard"
      vectorize_bindings = {
        VECTORIZE = {
          index_name = "my_index"
        }
      }
      wrangler_config_hash = "abc123def456"
    }
  }
  source = {
    config = {
      deployments_enabled = true
      owner = "my-org"
      owner_id = "12345678"
      path_excludes = ["string"]
      path_includes = ["string"]
      pr_comments_enabled = true
      preview_branch_excludes = ["string"]
      preview_branch_includes = ["string"]
      preview_deployment_setting = "all"
      production_branch = "main"
      production_deployments_enabled = true
      repo_id = "12345678"
      repo_name = "my-repo"
    }
    type = "github"
  }
}
```
<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `account_id` (String) Identifier.
- `name` (String) Name of the project.
- `production_branch` (String) Production branch of the project. Used to identify production deployments.

### Optional

- `build_config` (Attributes) Configs for the project build process. (see [below for nested schema](#nestedatt--build_config))
- `deployment_configs` (Attributes) Configs for deployments in a project. (see [below for nested schema](#nestedatt--deployment_configs))
- `source` (Attributes) Configs for the project source control. (see [below for nested schema](#nestedatt--source))

### Read-Only

- `canonical_deployment` (Attributes) Most recent production deployment of the project. (see [below for nested schema](#nestedatt--canonical_deployment))
- `created_on` (String) When the project was created.
- `domains` (List of String) A list of associated custom domains for the project.
- `framework` (String) Framework the project is using.
- `framework_version` (String) Version of the framework the project is using.
- `id` (String) Name of the project.
- `latest_deployment` (Attributes) Most recent deployment of the project. (see [below for nested schema](#nestedatt--latest_deployment))
- `preview_script_name` (String) Name of the preview script.
- `production_script_name` (String) Name of the production script.
- `subdomain` (String) The Cloudflare subdomain associated with the project.
- `uses_functions` (Boolean) Whether the project uses functions.

<a id="nestedatt--build_config"></a>
### Nested Schema for `build_config`

Optional:

- `build_caching` (Boolean) Enable build caching for the project.
- `build_command` (String) Command used to build project.
- `destination_dir` (String) Output directory of the build.
- `root_dir` (String) Directory to run the command.
- `web_analytics_tag` (String) The classifying tag for analytics.
- `web_analytics_token` (String, Sensitive) The auth token for analytics.


<a id="nestedatt--deployment_configs"></a>
### Nested Schema for `deployment_configs`

Optional:

- `preview` (Attributes) Configs for preview deploys. (see [below for nested schema](#nestedatt--deployment_configs--preview))
- `production` (Attributes) Configs for production deploys. (see [below for nested schema](#nestedatt--deployment_configs--production))

<a id="nestedatt--deployment_configs--preview"></a>
### Nested Schema for `deployment_configs.preview`

Optional:

- `ai_bindings` (Attributes Map) Constellation bindings used for Pages Functions. (see [below for nested schema](#nestedatt--deployment_configs--preview--ai_bindings))
- `always_use_latest_compatibility_date` (Boolean) Whether to always use the latest compatibility date for Pages Functions.
- `analytics_engine_datasets` (Attributes Map) Analytics Engine bindings used for Pages Functions. (see [below for nested schema](#nestedatt--deployment_configs--preview--analytics_engine_datasets))
- `browsers` (Attributes Map) Browser bindings used for Pages Functions. (see [below for nested schema](#nestedatt--deployment_configs--preview--browsers))
- `build_image_major_version` (Number) The major version of the build image to use for Pages Functions.
- `compatibility_date` (String) Compatibility date used for Pages Functions.
- `compatibility_flags` (List of String) Compatibility flags used for Pages Functions.
- `d1_databases` (Attributes Map) D1 databases used for Pages Functions. (see [below for nested schema](#nestedatt--deployment_configs--preview--d1_databases))
- `durable_object_namespaces` (Attributes Map) Durable Object namespaces used for Pages Functions. (see [below for nested schema](#nestedatt--deployment_configs--preview--durable_object_namespaces))
- `env_vars` (Attributes Map) Environment variables used for builds and Pages Functions. (see [below for nested schema](#nestedatt--deployment_configs--preview--env_vars))
- `fail_open` (Boolean) Whether to fail open when the deployment config cannot be applied.
- `hyperdrive_bindings` (Attributes Map) Hyperdrive bindings used for Pages Functions. (see [below for nested schema](#nestedatt--deployment_configs--preview--hyperdrive_bindings))
- `kv_namespaces` (Attributes Map) KV namespaces used for Pages Functions. (see [below for nested schema](#nestedatt--deployment_configs--preview--kv_namespaces))
- `limits` (Attributes) Limits for Pages Functions. (see [below for nested schema](#nestedatt--deployment_configs--preview--limits))
- `mtls_certificates` (Attributes Map) mTLS bindings used for Pages Functions. (see [below for nested schema](#nestedatt--deployment_configs--preview--mtls_certificates))
- `placement` (Attributes) Placement setting used for Pages Functions. (see [below for nested schema](#nestedatt--deployment_configs--preview--placement))
- `queue_producers` (Attributes Map) Queue Producer bindings used for Pages Functions. (see [below for nested schema](#nestedatt--deployment_configs--preview--queue_producers))
- `r2_buckets` (Attributes Map) R2 buckets used for Pages Functions. (see [below for nested schema](#nestedatt--deployment_configs--preview--r2_buckets))
- `services` (Attributes Map) Services used for Pages Functions. (see [below for nested schema](#nestedatt--deployment_configs--preview--services))
- `usage_model` (String, Deprecated) The usage model for Pages Functions.
Available values: "standard", "bundled", "unbound".
- `vectorize_bindings` (Attributes Map) Vectorize bindings used for Pages Functions. (see [below for nested schema](#nestedatt--deployment_configs--preview--vectorize_bindings))
- `wrangler_config_hash` (String) Hash of the Wrangler configuration used for the deployment.

<a id="nestedatt--deployment_configs--preview--ai_bindings"></a>
### Nested Schema for `deployment_configs.preview.ai_bindings`

Required:

- `project_id` (String)


<a id="nestedatt--deployment_configs--preview--analytics_engine_datasets"></a>
### Nested Schema for `deployment_configs.preview.analytics_engine_datasets`

Required:

- `dataset` (String) Name of the dataset.


<a id="nestedatt--deployment_configs--preview--browsers"></a>
### Nested Schema for `deployment_configs.preview.browsers`


<a id="nestedatt--deployment_configs--preview--d1_databases"></a>
### Nested Schema for `deployment_configs.preview.d1_databases`

Required:

- `id` (String) UUID of the D1 database.


<a id="nestedatt--deployment_configs--preview--durable_object_namespaces"></a>
### Nested Schema for `deployment_configs.preview.durable_object_namespaces`

Required:

- `namespace_id` (String) ID of the Durable Object namespace.


<a id="nestedatt--deployment_configs--preview--env_vars"></a>
### Nested Schema for `deployment_configs.preview.env_vars`

Required:

- `type` (String) Available values: "plain_text", "secret_text".
- `value` (String, Sensitive) Environment variable value.


<a id="nestedatt--deployment_configs--preview--hyperdrive_bindings"></a>
### Nested Schema for `deployment_configs.preview.hyperdrive_bindings`

Required:

- `id` (String)


<a id="nestedatt--deployment_configs--preview--kv_namespaces"></a>
### Nested Schema for `deployment_configs.preview.kv_namespaces`

Required:

- `namespace_id` (String) ID of the KV namespace.


<a id="nestedatt--deployment_configs--preview--limits"></a>
### Nested Schema for `deployment_configs.preview.limits`

Required:

- `cpu_ms` (Number) CPU time limit in milliseconds.


<a id="nestedatt--deployment_configs--preview--mtls_certificates"></a>
### Nested Schema for `deployment_configs.preview.mtls_certificates`

Required:

- `certificate_id` (String)


<a id="nestedatt--deployment_configs--preview--placement"></a>
### Nested Schema for `deployment_configs.preview.placement`

Optional:

- `mode` (String) Placement mode.


<a id="nestedatt--deployment_configs--preview--queue_producers"></a>
### Nested Schema for `deployment_configs.preview.queue_producers`

Required:

- `name` (String) Name of the Queue.


<a id="nestedatt--deployment_configs--preview--r2_buckets"></a>
### Nested Schema for `deployment_configs.preview.r2_buckets`

Required:

- `name` (String) Name of the R2 bucket.

Optional:

- `jurisdiction` (String) Jurisdiction of the R2 bucket.


<a id="nestedatt--deployment_configs--preview--services"></a>
### Nested Schema for `deployment_configs.preview.services`

Required:

- `service` (String) The Service name.

Optional:

- `entrypoint` (String) The entrypoint to bind to.
- `environment` (String) The Service environment.


<a id="nestedatt--deployment_configs--preview--vectorize_bindings"></a>
### Nested Schema for `deployment_configs.preview.vectorize_bindings`

Required:

- `index_name` (String)



<a id="nestedatt--deployment_configs--production"></a>
### Nested Schema for `deployment_configs.production`

Optional:

- `ai_bindings` (Attributes Map) Constellation bindings used for Pages Functions. (see [below for nested schema](#nestedatt--deployment_configs--production--ai_bindings))
- `always_use_latest_compatibility_date` (Boolean) Whether to always use the latest compatibility date for Pages Functions.
- `analytics_engine_datasets` (Attributes Map) Analytics Engine bindings used for Pages Functions. (see [below for nested schema](#nestedatt--deployment_configs--production--analytics_engine_datasets))
- `browsers` (Attributes Map) Browser bindings used for Pages Functions. (see [below for nested schema](#nestedatt--deployment_configs--production--browsers))
- `build_image_major_version` (Number) The major version of the build image to use for Pages Functions.
- `compatibility_date` (String) Compatibility date used for Pages Functions.
- `compatibility_flags` (List of String) Compatibility flags used for Pages Functions.
- `d1_databases` (Attributes Map) D1 databases used for Pages Functions. (see [below for nested schema](#nestedatt--deployment_configs--production--d1_databases))
- `durable_object_namespaces` (Attributes Map) Durable Object namespaces used for Pages Functions. (see [below for nested schema](#nestedatt--deployment_configs--production--durable_object_namespaces))
- `env_vars` (Attributes Map) Environment variables used for builds and Pages Functions. (see [below for nested schema](#nestedatt--deployment_configs--production--env_vars))
- `fail_open` (Boolean) Whether to fail open when the deployment config cannot be applied.
- `hyperdrive_bindings` (Attributes Map) Hyperdrive bindings used for Pages Functions. (see [below for nested schema](#nestedatt--deployment_configs--production--hyperdrive_bindings))
- `kv_namespaces` (Attributes Map) KV namespaces used for Pages Functions. (see [below for nested schema](#nestedatt--deployment_configs--production--kv_namespaces))
- `limits` (Attributes) Limits for Pages Functions. (see [below for nested schema](#nestedatt--deployment_configs--production--limits))
- `mtls_certificates` (Attributes Map) mTLS bindings used for Pages Functions. (see [below for nested schema](#nestedatt--deployment_configs--production--mtls_certificates))
- `placement` (Attributes) Placement setting used for Pages Functions. (see [below for nested schema](#nestedatt--deployment_configs--production--placement))
- `queue_producers` (Attributes Map) Queue Producer bindings used for Pages Functions. (see [below for nested schema](#nestedatt--deployment_configs--production--queue_producers))
- `r2_buckets` (Attributes Map) R2 buckets used for Pages Functions. (see [below for nested schema](#nestedatt--deployment_configs--production--r2_buckets))
- `services` (Attributes Map) Services used for Pages Functions. (see [below for nested schema](#nestedatt--deployment_configs--production--services))
- `usage_model` (String, Deprecated) The usage model for Pages Functions.
Available values: "standard", "bundled", "unbound".
- `vectorize_bindings` (Attributes Map) Vectorize bindings used for Pages Functions. (see [below for nested schema](#nestedatt--deployment_configs--production--vectorize_bindings))
- `wrangler_config_hash` (String) Hash of the Wrangler configuration used for the deployment.

<a id="nestedatt--deployment_configs--production--ai_bindings"></a>
### Nested Schema for `deployment_configs.production.ai_bindings`

Required:

- `project_id` (String)


<a id="nestedatt--deployment_configs--production--analytics_engine_datasets"></a>
### Nested Schema for `deployment_configs.production.analytics_engine_datasets`

Required:

- `dataset` (String) Name of the dataset.


<a id="nestedatt--deployment_configs--production--browsers"></a>
### Nested Schema for `deployment_configs.production.browsers`


<a id="nestedatt--deployment_configs--production--d1_databases"></a>
### Nested Schema for `deployment_configs.production.d1_databases`

Required:

- `id` (String) UUID of the D1 database.


<a id="nestedatt--deployment_configs--production--durable_object_namespaces"></a>
### Nested Schema for `deployment_configs.production.durable_object_namespaces`

Required:

- `namespace_id` (String) ID of the Durable Object namespace.


<a id="nestedatt--deployment_configs--production--env_vars"></a>
### Nested Schema for `deployment_configs.production.env_vars`

Required:

- `type` (String) Available values: "plain_text", "secret_text".
- `value` (String, Sensitive) Environment variable value.


<a id="nestedatt--deployment_configs--production--hyperdrive_bindings"></a>
### Nested Schema for `deployment_configs.production.hyperdrive_bindings`

Required:

- `id` (String)


<a id="nestedatt--deployment_configs--production--kv_namespaces"></a>
### Nested Schema for `deployment_configs.production.kv_namespaces`

Required:

- `namespace_id` (String) ID of the KV namespace.


<a id="nestedatt--deployment_configs--production--limits"></a>
### Nested Schema for `deployment_configs.production.limits`

Required:

- `cpu_ms` (Number) CPU time limit in milliseconds.


<a id="nestedatt--deployment_configs--production--mtls_certificates"></a>
### Nested Schema for `deployment_configs.production.mtls_certificates`

Required:

- `certificate_id` (String)


<a id="nestedatt--deployment_configs--production--placement"></a>
### Nested Schema for `deployment_configs.production.placement`

Optional:

- `mode` (String) Placement mode.


<a id="nestedatt--deployment_configs--production--queue_producers"></a>
### Nested Schema for `deployment_configs.production.queue_producers`

Required:

- `name` (String) Name of the Queue.


<a id="nestedatt--deployment_configs--production--r2_buckets"></a>
### Nested Schema for `deployment_configs.production.r2_buckets`

Required:

- `name` (String) Name of the R2 bucket.

Optional:

- `jurisdiction` (String) Jurisdiction of the R2 bucket.


<a id="nestedatt--deployment_configs--production--services"></a>
### Nested Schema for `deployment_configs.production.services`

Required:

- `service` (String) The Service name.

Optional:

- `entrypoint` (String) The entrypoint to bind to.
- `environment` (String) The Service environment.


<a id="nestedatt--deployment_configs--production--vectorize_bindings"></a>
### Nested Schema for `deployment_configs.production.vectorize_bindings`

Required:

- `index_name` (String)




<a id="nestedatt--source"></a>
### Nested Schema for `source`

Required:

- `config` (Attributes) (see [below for nested schema](#nestedatt--source--config))
- `type` (String) The source control management provider.
Available values: "github", "gitlab".

<a id="nestedatt--source--config"></a>
### Nested Schema for `source.config`

Optional:

- `deployments_enabled` (Boolean, Deprecated) Whether to enable automatic deployments when pushing to the source repository.
When disabled, no deployments (production or preview) will be triggered automatically.
- `owner` (String) The owner of the repository.
- `owner_id` (String) The owner ID of the repository.
- `path_excludes` (List of String) A list of paths that should be excluded from triggering a preview deployment. Wildcard syntax (`*`) is supported.
- `path_includes` (List of String) A list of paths that should be watched to trigger a preview deployment. Wildcard syntax (`*`) is supported.
- `pr_comments_enabled` (Boolean) Whether to enable PR comments.
- `preview_branch_excludes` (List of String) A list of branches that should not trigger a preview deployment. Wildcard syntax (`*`) is supported. Must be used with `preview_deployment_setting` set to `custom`.
- `preview_branch_includes` (List of String) A list of branches that should trigger a preview deployment. Wildcard syntax (`*`) is supported. Must be used with `preview_deployment_setting` set to `custom`.
- `preview_deployment_setting` (String) Controls whether commits to preview branches trigger a preview deployment.
Available values: "all", "none", "custom".
- `production_branch` (String) The production branch of the repository.
- `production_deployments_enabled` (Boolean) Whether to trigger a production deployment on commits to the production branch.
- `repo_id` (String) The ID of the repository.
- `repo_name` (String) The name of the repository.



<a id="nestedatt--canonical_deployment"></a>
### Nested Schema for `canonical_deployment`

Read-Only:

- `aliases` (List of String) A list of alias URLs pointing to this deployment.
- `build_config` (Attributes) Configs for the project build process. (see [below for nested schema](#nestedatt--canonical_deployment--build_config))
- `created_on` (String) When the deployment was created.
- `deployment_trigger` (Attributes) Info about what caused the deployment. (see [below for nested schema](#nestedatt--canonical_deployment--deployment_trigger))
- `env_vars` (Attributes Map) Environment variables used for builds and Pages Functions. (see [below for nested schema](#nestedatt--canonical_deployment--env_vars))
- `environment` (String) Type of deploy.
Available values: "preview", "production".
- `id` (String) Id of the deployment.
- `is_skipped` (Boolean) If the deployment has been skipped.
- `latest_stage` (Attributes) The status of the deployment. (see [below for nested schema](#nestedatt--canonical_deployment--latest_stage))
- `modified_on` (String) When the deployment was last modified.
- `project_id` (String) Id of the project.
- `project_name` (String) Name of the project.
- `short_id` (String) Short Id (8 character) of the deployment.
- `source` (Attributes) Configs for the project source control. (see [below for nested schema](#nestedatt--canonical_deployment--source))
- `stages` (Attributes List) List of past stages. (see [below for nested schema](#nestedatt--canonical_deployment--stages))
- `url` (String) The live URL to view this deployment.
- `uses_functions` (Boolean) Whether the deployment uses functions.

<a id="nestedatt--canonical_deployment--build_config"></a>
### Nested Schema for `canonical_deployment.build_config`

Read-Only:

- `build_caching` (Boolean) Enable build caching for the project.
- `build_command` (String) Command used to build project.
- `destination_dir` (String) Assets output directory of the build.
- `root_dir` (String) Directory to run the command.
- `web_analytics_tag` (String) The classifying tag for analytics.
- `web_analytics_token` (String, Sensitive) The auth token for analytics.


<a id="nestedatt--canonical_deployment--deployment_trigger"></a>
### Nested Schema for `canonical_deployment.deployment_trigger`

Read-Only:

- `metadata` (Attributes) Additional info about the trigger. (see [below for nested schema](#nestedatt--canonical_deployment--deployment_trigger--metadata))
- `type` (String) What caused the deployment.
Available values: "github:push", "ad_hoc", "deploy_hook".

<a id="nestedatt--canonical_deployment--deployment_trigger--metadata"></a>
### Nested Schema for `canonical_deployment.deployment_trigger.metadata`

Read-Only:

- `branch` (String) Where the trigger happened.
- `commit_dirty` (Boolean) Whether the deployment trigger commit was dirty.
- `commit_hash` (String) Hash of the deployment trigger commit.
- `commit_message` (String) Message of the deployment trigger commit.



<a id="nestedatt--canonical_deployment--env_vars"></a>
### Nested Schema for `canonical_deployment.env_vars`

Read-Only:

- `type` (String) Available values: "plain_text", "secret_text".
- `value` (String, Sensitive) Environment variable value.


<a id="nestedatt--canonical_deployment--latest_stage"></a>
### Nested Schema for `canonical_deployment.latest_stage`

Read-Only:

- `ended_on` (String) When the stage ended.
- `name` (String) The current build stage.
Available values: "queued", "initialize", "clone_repo", "build", "deploy".
- `started_on` (String) When the stage started.
- `status` (String) State of the current stage.
Available values: "success", "idle", "active", "failure", "canceled".


<a id="nestedatt--canonical_deployment--source"></a>
### Nested Schema for `canonical_deployment.source`

Read-Only:

- `config` (Attributes) (see [below for nested schema](#nestedatt--canonical_deployment--source--config))
- `type` (String) The source control management provider.
Available values: "github", "gitlab".

<a id="nestedatt--canonical_deployment--source--config"></a>
### Nested Schema for `canonical_deployment.source.config`

Read-Only:

- `deployments_enabled` (Boolean, Deprecated) Whether to enable automatic deployments when pushing to the source repository.
When disabled, no deployments (production or preview) will be triggered automatically.
- `owner` (String) The owner of the repository.
- `owner_id` (String) The owner ID of the repository.
- `path_excludes` (List of String) A list of paths that should be excluded from triggering a preview deployment. Wildcard syntax (`*`) is supported.
- `path_includes` (List of String) A list of paths that should be watched to trigger a preview deployment. Wildcard syntax (`*`) is supported.
- `pr_comments_enabled` (Boolean) Whether to enable PR comments.
- `preview_branch_excludes` (List of String) A list of branches that should not trigger a preview deployment. Wildcard syntax (`*`) is supported. Must be used with `preview_deployment_setting` set to `custom`.
- `preview_branch_includes` (List of String) A list of branches that should trigger a preview deployment. Wildcard syntax (`*`) is supported. Must be used with `preview_deployment_setting` set to `custom`.
- `preview_deployment_setting` (String) Controls whether commits to preview branches trigger a preview deployment.
Available values: "all", "none", "custom".
- `production_branch` (String) The production branch of the repository.
- `production_deployments_enabled` (Boolean) Whether to trigger a production deployment on commits to the production branch.
- `repo_id` (String) The ID of the repository.
- `repo_name` (String) The name of the repository.



<a id="nestedatt--canonical_deployment--stages"></a>
### Nested Schema for `canonical_deployment.stages`

Read-Only:

- `ended_on` (String) When the stage ended.
- `name` (String) The current build stage.
Available values: "queued", "initialize", "clone_repo", "build", "deploy".
- `started_on` (String) When the stage started.
- `status` (String) State of the current stage.
Available values: "success", "idle", "active", "failure", "canceled".



<a id="nestedatt--latest_deployment"></a>
### Nested Schema for `latest_deployment`

Read-Only:

- `aliases` (List of String) A list of alias URLs pointing to this deployment.
- `build_config` (Attributes) Configs for the project build process. (see [below for nested schema](#nestedatt--latest_deployment--build_config))
- `created_on` (String) When the deployment was created.
- `deployment_trigger` (Attributes) Info about what caused the deployment. (see [below for nested schema](#nestedatt--latest_deployment--deployment_trigger))
- `env_vars` (Attributes Map) Environment variables used for builds and Pages Functions. (see [below for nested schema](#nestedatt--latest_deployment--env_vars))
- `environment` (String) Type of deploy.
Available values: "preview", "production".
- `id` (String) Id of the deployment.
- `is_skipped` (Boolean) If the deployment has been skipped.
- `latest_stage` (Attributes) The status of the deployment. (see [below for nested schema](#nestedatt--latest_deployment--latest_stage))
- `modified_on` (String) When the deployment was last modified.
- `project_id` (String) Id of the project.
- `project_name` (String) Name of the project.
- `short_id` (String) Short Id (8 character) of the deployment.
- `source` (Attributes) Configs for the project source control. (see [below for nested schema](#nestedatt--latest_deployment--source))
- `stages` (Attributes List) List of past stages. (see [below for nested schema](#nestedatt--latest_deployment--stages))
- `url` (String) The live URL to view this deployment.
- `uses_functions` (Boolean) Whether the deployment uses functions.

<a id="nestedatt--latest_deployment--build_config"></a>
### Nested Schema for `latest_deployment.build_config`

Read-Only:

- `build_caching` (Boolean) Enable build caching for the project.
- `build_command` (String) Command used to build project.
- `destination_dir` (String) Assets output directory of the build.
- `root_dir` (String) Directory to run the command.
- `web_analytics_tag` (String) The classifying tag for analytics.
- `web_analytics_token` (String, Sensitive) The auth token for analytics.


<a id="nestedatt--latest_deployment--deployment_trigger"></a>
### Nested Schema for `latest_deployment.deployment_trigger`

Read-Only:

- `metadata` (Attributes) Additional info about the trigger. (see [below for nested schema](#nestedatt--latest_deployment--deployment_trigger--metadata))
- `type` (String) What caused the deployment.
Available values: "github:push", "ad_hoc", "deploy_hook".

<a id="nestedatt--latest_deployment--deployment_trigger--metadata"></a>
### Nested Schema for `latest_deployment.deployment_trigger.metadata`

Read-Only:

- `branch` (String) Where the trigger happened.
- `commit_dirty` (Boolean) Whether the deployment trigger commit was dirty.
- `commit_hash` (String) Hash of the deployment trigger commit.
- `commit_message` (String) Message of the deployment trigger commit.



<a id="nestedatt--latest_deployment--env_vars"></a>
### Nested Schema for `latest_deployment.env_vars`

Read-Only:

- `type` (String) Available values: "plain_text", "secret_text".
- `value` (String, Sensitive) Environment variable value.


<a id="nestedatt--latest_deployment--latest_stage"></a>
### Nested Schema for `latest_deployment.latest_stage`

Read-Only:

- `ended_on` (String) When the stage ended.
- `name` (String) The current build stage.
Available values: "queued", "initialize", "clone_repo", "build", "deploy".
- `started_on` (String) When the stage started.
- `status` (String) State of the current stage.
Available values: "success", "idle", "active", "failure", "canceled".


<a id="nestedatt--latest_deployment--source"></a>
### Nested Schema for `latest_deployment.source`

Read-Only:

- `config` (Attributes) (see [below for nested schema](#nestedatt--latest_deployment--source--config))
- `type` (String) The source control management provider.
Available values: "github", "gitlab".

<a id="nestedatt--latest_deployment--source--config"></a>
### Nested Schema for `latest_deployment.source.config`

Read-Only:

- `deployments_enabled` (Boolean, Deprecated) Whether to enable automatic deployments when pushing to the source repository.
When disabled, no deployments (production or preview) will be triggered automatically.
- `owner` (String) The owner of the repository.
- `owner_id` (String) The owner ID of the repository.
- `path_excludes` (List of String) A list of paths that should be excluded from triggering a preview deployment. Wildcard syntax (`*`) is supported.
- `path_includes` (List of String) A list of paths that should be watched to trigger a preview deployment. Wildcard syntax (`*`) is supported.
- `pr_comments_enabled` (Boolean) Whether to enable PR comments.
- `preview_branch_excludes` (List of String) A list of branches that should not trigger a preview deployment. Wildcard syntax (`*`) is supported. Must be used with `preview_deployment_setting` set to `custom`.
- `preview_branch_includes` (List of String) A list of branches that should trigger a preview deployment. Wildcard syntax (`*`) is supported. Must be used with `preview_deployment_setting` set to `custom`.
- `preview_deployment_setting` (String) Controls whether commits to preview branches trigger a preview deployment.
Available values: "all", "none", "custom".
- `production_branch` (String) The production branch of the repository.
- `production_deployments_enabled` (Boolean) Whether to trigger a production deployment on commits to the production branch.
- `repo_id` (String) The ID of the repository.
- `repo_name` (String) The name of the repository.



<a id="nestedatt--latest_deployment--stages"></a>
### Nested Schema for `latest_deployment.stages`

Read-Only:

- `ended_on` (String) When the stage ended.
- `name` (String) The current build stage.
Available values: "queued", "initialize", "clone_repo", "build", "deploy".
- `started_on` (String) When the stage started.
- `status` (String) State of the current stage.
Available values: "success", "idle", "active", "failure", "canceled".

## Import

!> It is not possible to import a pages project with secret environment variables. If you have a secret environment variable, you must remove it from your project before importing it.

Import is supported using the following syntax:

```shell
$ terraform import cloudflare_pages_project.example '<account_id>/<project_name>'
```
