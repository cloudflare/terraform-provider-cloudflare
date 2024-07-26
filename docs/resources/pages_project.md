---
page_title: "cloudflare_pages_project Resource - Cloudflare"
subcategory: ""
description: |-
  
---

# cloudflare_pages_project (Resource)



-> If you are using a `source` block configuration, you must first have a
   connected GitHub or GitLab account connected to Cloudflare. See the
   [Getting Started with Pages] documentation on how to link your accounts.

## Example Usage

```terraform
# Direct upload Pages project
resource "cloudflare_pages_project" "basic_project" {
  account_id        = "f037e56e89293a057740de681ac9abbe"
  name              = "this-is-my-project-01"
  production_branch = "main"
}

# Pages project with managing build config
resource "cloudflare_pages_project" "build_config" {
  account_id        = "f037e56e89293a057740de681ac9abbe"
  name              = "this-is-my-project-01"
  production_branch = "main"
  build_config = {
    build_command       = "npm run build"
    destination_dir     = "build"
    root_dir            = ""
    web_analytics_tag   = "cee1c73f6e4743d0b5e6bb1a0bcaabcc"
    web_analytics_token = "021e1057c18547eca7b79f2516f06o7x"
  }
}

# Pages project managing project source
resource "cloudflare_pages_project" "source_config" {
  account_id        = "f037e56e89293a057740de681ac9abbe"
  name              = "this-is-my-project-01"
  production_branch = "main"
  source = [{
    type = "github"
    config = [{
      owner                         = "cloudflare"
      repo_name                     = "ninjakittens"
      production_branch             = "main"
      pr_comments_enabled           = true
      deployments_enabled           = true
      production_deployment_enabled = true
      preview_deployment_setting    = "custom"
      preview_branch_includes       = ["dev", "preview"]
      preview_branch_excludes       = ["main", "prod"]
    }]
  }]
}

# Pages project managing deployment configs
resource "cloudflare_pages_project" "deployment_configs" {
  account_id        = "f037e56e89293a057740de681ac9abbe"
  name              = "this-is-my-project-01"
  production_branch = "main"
  deployment_configs = {
    preview = [{
      environment_variables = {
        ENVIRONMENT = "preview"
      }
      secrets = {
        TURNSTILE_SECRET = "1x0000000000000000000000000000000AA"
      }
      kv_namespaces = {
        KV_BINDING = "5eb63bbbe01eeed093cb22bb8f5acdc3"
      }
      durable_object_namespaces = {
        DO_BINDING = "5eb63bbbe01eeed093cb22bb8f5acdc3"
      }
      r2_buckets = {
        R2_BINDING = "some-bucket"
      }
      d1_databases = {
        D1_BINDING = "445e2955-951a-4358-a35b-a4d0c813f63"
      }
      compatibility_date  = "2022-08-15"
      compatibility_flags = ["nodejs_compat"]
    }]
    production = [{
      environment_variables = {
        ENVIRONMENT = "production"
        OTHER_VALUE = "other value"
      }
      secrets = {
        TURNSTILE_SECRET       = "1x0000000000000000000000000000000AA"
        TURNSTILE_INVIS_SECRET = "2x0000000000000000000000000000000AA"
      }
      kv_namespaces = {
        KV_BINDING_1 = "5eb63bbbe01eeed093cb22bb8f5acdc3"
        KV_BINDING_2 = "3cdca5f8bb22bc390deee10ebbb36be5"
      }
      durable_object_namespaces = {
        DO_BINDING_1 = "5eb63bbbe01eeed093cb22bb8f5acdc3"
        DO_BINDING_2 = "3cdca5f8bb22bc390deee10ebbb36be5"
      }
      r2_buckets = {
        R2_BINDING_1 = "some-bucket"
        R2_BINDING_2 = "other-bucket"
      }
      d1_databases = {
        D1_BINDING_1 = "445e2955-951a-4358-a35b-a4d0c813f63"
        D1_BINDING_2 = "a399414b-c697-409a-a688-377db6433cd9"
      }
      compatibility_date  = "2022-08-16"
      compatibility_flags = ["nodejs_compat", "streams_enable_constructors"]
    }]
  }
}

# Pages project managing all configs
resource "cloudflare_pages_project" "deployment_configs" {
  account_id        = "f037e56e89293a057740de681ac9abbe"
  name              = "this-is-my-project-01"
  production_branch = "main"

  source = [{
    type = "github"
    config = [{
      owner                         = "cloudflare"
      repo_name                     = "ninjakittens"
      production_branch             = "main"
      pr_comments_enabled           = true
      deployments_enabled           = true
      production_deployment_enabled = true
      preview_deployment_setting    = "custom"
      preview_branch_includes       = ["dev", "preview"]
      preview_branch_excludes       = ["main", "prod"]
    }]
  }]

  build_config = {
    build_command       = "npm run build"
    destination_dir     = "build"
    root_dir            = ""
    web_analytics_tag   = "cee1c73f6e4743d0b5e6bb1a0bcaabcc"
    web_analytics_token = "021e1057c18547eca7b79f2516f06o7x"
  }

  deployment_configs = {
    preview = [{
      environment_variables = {
        ENVIRONMENT = "preview"
      }
      secrets = {
        TURNSTILE_SECRET = "1x0000000000000000000000000000000AA"
      }
      kv_namespaces = {
        KV_BINDING = "5eb63bbbe01eeed093cb22bb8f5acdc3"
      }
      durable_object_namespaces = {
        DO_BINDING = "5eb63bbbe01eeed093cb22bb8f5acdc3"
      }
      r2_buckets = {
        R2_BINDING = "some-bucket"
      }
      d1_databases = {
        D1_BINDING = "445e2955-951a-4358-a35b-a4d0c813f63"
      }
      compatibility_date  = "2022-08-15"
      compatibility_flags = ["nodejs_compat"]
    }]
    production = [{
      environment_variables = {
        ENVIRONMENT = "production"
        OTHER_VALUE = "other value"
      }
      secrets = {
        TURNSTILE_SECRET       = "1x0000000000000000000000000000000AA"
        TURNSTILE_INVIS_SECRET = "2x0000000000000000000000000000000AA"
      }
      kv_namespaces = {
        KV_BINDING_1 = "5eb63bbbe01eeed093cb22bb8f5acdc3"
        KV_BINDING_2 = "3cdca5f8bb22bc390deee10ebbb36be5"
      }
      durable_object_namespaces = {
        DO_BINDING_1 = "5eb63bbbe01eeed093cb22bb8f5acdc3"
        DO_BINDING_2 = "3cdca5f8bb22bc390deee10ebbb36be5"
      }
      r2_buckets = {
        R2_BINDING_1 = "some-bucket"
        R2_BINDING_2 = "other-bucket"
      }
      d1_databases = {
        D1_BINDING_1 = "445e2955-951a-4358-a35b-a4d0c813f63"
        D1_BINDING_2 = "a399414b-c697-409a-a688-377db6433cd9"
      }
      compatibility_date  = "2022-08-16"
      compatibility_flags = ["nodejs_compat", "streams_enable_constructors"]
    }]
  }
}
```
<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `account_id` (String) Identifier

### Optional

- `build_config` (Attributes) Configs for the project build process. (see [below for nested schema](#nestedatt--build_config))
- `canonical_deployment` (Attributes) (see [below for nested schema](#nestedatt--canonical_deployment))
- `deployment_configs` (Attributes) Configs for deployments in a project. (see [below for nested schema](#nestedatt--deployment_configs))
- `latest_deployment` (Attributes) (see [below for nested schema](#nestedatt--latest_deployment))
- `name` (String) Name of the project.
- `production_branch` (String) Production branch of the project. Used to identify production deployments.
- `project_name` (String) Name of the project.

### Read-Only

- `created_on` (String) When the project was created.
- `domains` (List of String) A list of associated custom domains for the project.
- `id` (String) Id of the project.
- `source` (String)
- `subdomain` (String) The Cloudflare subdomain associated with the project.

<a id="nestedatt--build_config"></a>
### Nested Schema for `build_config`

Optional:

- `build_caching` (Boolean) Enable build caching for the project.
- `build_command` (String) Command used to build project.
- `destination_dir` (String) Output directory of the build.
- `root_dir` (String) Directory to run the command.
- `web_analytics_tag` (String) The classifying tag for analytics.
- `web_analytics_token` (String) The auth token for analytics.


<a id="nestedatt--canonical_deployment"></a>
### Nested Schema for `canonical_deployment`

Read-Only:

- `aliases` (List of String) A list of alias URLs pointing to this deployment.
- `build_config` (String)
- `created_on` (String) When the deployment was created.
- `env_vars` (String) A dict of env variables to build this deploy.
- `environment` (String) Type of deploy.
- `id` (String) Id of the deployment.
- `is_skipped` (Boolean) If the deployment has been skipped.
- `latest_stage` (String)
- `modified_on` (String) When the deployment was last modified.
- `project_id` (String) Id of the project.
- `project_name` (String) Name of the project.
- `short_id` (String) Short Id (8 character) of the deployment.
- `source` (String)
- `stages` (Attributes List) List of past stages. (see [below for nested schema](#nestedatt--canonical_deployment--stages))
- `url` (String) The live URL to view this deployment.

<a id="nestedatt--canonical_deployment--stages"></a>
### Nested Schema for `canonical_deployment.stages`

Optional:

- `name` (String) The current build stage.

Read-Only:

- `ended_on` (String) When the stage ended.
- `started_on` (String) When the stage started.
- `status` (String) State of the current stage.



<a id="nestedatt--deployment_configs"></a>
### Nested Schema for `deployment_configs`

Optional:

- `preview` (Attributes) Configs for preview deploys. (see [below for nested schema](#nestedatt--deployment_configs--preview))
- `production` (Attributes) Configs for production deploys. (see [below for nested schema](#nestedatt--deployment_configs--production))

<a id="nestedatt--deployment_configs--preview"></a>
### Nested Schema for `deployment_configs.preview`

Optional:

- `ai_bindings` (Attributes) Constellation bindings used for Pages Functions. (see [below for nested schema](#nestedatt--deployment_configs--preview--ai_bindings))
- `analytics_engine_datasets` (Attributes) Analytics Engine bindings used for Pages Functions. (see [below for nested schema](#nestedatt--deployment_configs--preview--analytics_engine_datasets))
- `browsers` (Attributes) Browser bindings used for Pages Functions. (see [below for nested schema](#nestedatt--deployment_configs--preview--browsers))
- `compatibility_date` (String) Compatibility date used for Pages Functions.
- `compatibility_flags` (List of String) Compatibility flags used for Pages Functions.
- `d1_databases` (Attributes) D1 databases used for Pages Functions. (see [below for nested schema](#nestedatt--deployment_configs--preview--d1_databases))
- `durable_object_namespaces` (Attributes) Durabble Object namespaces used for Pages Functions. (see [below for nested schema](#nestedatt--deployment_configs--preview--durable_object_namespaces))
- `env_vars` (Attributes) Environment variables for build configs. (see [below for nested schema](#nestedatt--deployment_configs--preview--env_vars))
- `hyperdrive_bindings` (Attributes) Hyperdrive bindings used for Pages Functions. (see [below for nested schema](#nestedatt--deployment_configs--preview--hyperdrive_bindings))
- `kv_namespaces` (Attributes) KV namespaces used for Pages Functions. (see [below for nested schema](#nestedatt--deployment_configs--preview--kv_namespaces))
- `mtls_certificates` (Attributes) mTLS bindings used for Pages Functions. (see [below for nested schema](#nestedatt--deployment_configs--preview--mtls_certificates))
- `placement` (Attributes) Placement setting used for Pages Functions. (see [below for nested schema](#nestedatt--deployment_configs--preview--placement))
- `queue_producers` (Attributes) Queue Producer bindings used for Pages Functions. (see [below for nested schema](#nestedatt--deployment_configs--preview--queue_producers))
- `r2_buckets` (Attributes) R2 buckets used for Pages Functions. (see [below for nested schema](#nestedatt--deployment_configs--preview--r2_buckets))
- `services` (Attributes) Services used for Pages Functions. (see [below for nested schema](#nestedatt--deployment_configs--preview--services))
- `vectorize_bindings` (Attributes) Vectorize bindings used for Pages Functions. (see [below for nested schema](#nestedatt--deployment_configs--preview--vectorize_bindings))

<a id="nestedatt--deployment_configs--preview--ai_bindings"></a>
### Nested Schema for `deployment_configs.preview.ai_bindings`

Optional:

- `ai_binding` (Attributes) AI binding. (see [below for nested schema](#nestedatt--deployment_configs--preview--ai_bindings--ai_binding))

<a id="nestedatt--deployment_configs--preview--ai_bindings--ai_binding"></a>
### Nested Schema for `deployment_configs.preview.ai_bindings.ai_binding`

Optional:

- `project_id` (String)



<a id="nestedatt--deployment_configs--preview--analytics_engine_datasets"></a>
### Nested Schema for `deployment_configs.preview.analytics_engine_datasets`

Optional:

- `analytics_engine_binding` (Attributes) Analytics Engine binding. (see [below for nested schema](#nestedatt--deployment_configs--preview--analytics_engine_datasets--analytics_engine_binding))

<a id="nestedatt--deployment_configs--preview--analytics_engine_datasets--analytics_engine_binding"></a>
### Nested Schema for `deployment_configs.preview.analytics_engine_datasets.analytics_engine_binding`

Optional:

- `dataset` (String) Name of the dataset.



<a id="nestedatt--deployment_configs--preview--browsers"></a>
### Nested Schema for `deployment_configs.preview.browsers`

Optional:

- `browser` (String) Browser binding.


<a id="nestedatt--deployment_configs--preview--d1_databases"></a>
### Nested Schema for `deployment_configs.preview.d1_databases`

Optional:

- `d1_binding` (Attributes) D1 binding. (see [below for nested schema](#nestedatt--deployment_configs--preview--d1_databases--d1_binding))

<a id="nestedatt--deployment_configs--preview--d1_databases--d1_binding"></a>
### Nested Schema for `deployment_configs.preview.d1_databases.d1_binding`

Optional:

- `id` (String) UUID of the D1 database.



<a id="nestedatt--deployment_configs--preview--durable_object_namespaces"></a>
### Nested Schema for `deployment_configs.preview.durable_object_namespaces`

Optional:

- `do_binding` (Attributes) Durabble Object binding. (see [below for nested schema](#nestedatt--deployment_configs--preview--durable_object_namespaces--do_binding))

<a id="nestedatt--deployment_configs--preview--durable_object_namespaces--do_binding"></a>
### Nested Schema for `deployment_configs.preview.durable_object_namespaces.do_binding`

Optional:

- `namespace_id` (String) ID of the Durabble Object namespace.



<a id="nestedatt--deployment_configs--preview--env_vars"></a>
### Nested Schema for `deployment_configs.preview.env_vars`

Optional:

- `environment_variable` (Attributes) Environment variable. (see [below for nested schema](#nestedatt--deployment_configs--preview--env_vars--environment_variable))

<a id="nestedatt--deployment_configs--preview--env_vars--environment_variable"></a>
### Nested Schema for `deployment_configs.preview.env_vars.environment_variable`

Optional:

- `type` (String) The type of environment variable (plain text or secret)
- `value` (String) Environment variable value.



<a id="nestedatt--deployment_configs--preview--hyperdrive_bindings"></a>
### Nested Schema for `deployment_configs.preview.hyperdrive_bindings`

Optional:

- `hyperdrive` (Attributes) Hyperdrive binding. (see [below for nested schema](#nestedatt--deployment_configs--preview--hyperdrive_bindings--hyperdrive))

<a id="nestedatt--deployment_configs--preview--hyperdrive_bindings--hyperdrive"></a>
### Nested Schema for `deployment_configs.preview.hyperdrive_bindings.hyperdrive`

Optional:

- `id` (String)



<a id="nestedatt--deployment_configs--preview--kv_namespaces"></a>
### Nested Schema for `deployment_configs.preview.kv_namespaces`

Optional:

- `kv_binding` (Attributes) KV binding. (see [below for nested schema](#nestedatt--deployment_configs--preview--kv_namespaces--kv_binding))

<a id="nestedatt--deployment_configs--preview--kv_namespaces--kv_binding"></a>
### Nested Schema for `deployment_configs.preview.kv_namespaces.kv_binding`

Optional:

- `namespace_id` (String) ID of the KV namespace.



<a id="nestedatt--deployment_configs--preview--mtls_certificates"></a>
### Nested Schema for `deployment_configs.preview.mtls_certificates`

Optional:

- `mtls` (Attributes) mTLS binding. (see [below for nested schema](#nestedatt--deployment_configs--preview--mtls_certificates--mtls))

<a id="nestedatt--deployment_configs--preview--mtls_certificates--mtls"></a>
### Nested Schema for `deployment_configs.preview.mtls_certificates.mtls`

Optional:

- `certificate_id` (String)



<a id="nestedatt--deployment_configs--preview--placement"></a>
### Nested Schema for `deployment_configs.preview.placement`

Optional:

- `mode` (String) Placement mode.


<a id="nestedatt--deployment_configs--preview--queue_producers"></a>
### Nested Schema for `deployment_configs.preview.queue_producers`

Optional:

- `queue_producer_binding` (Attributes) Queue Producer binding. (see [below for nested schema](#nestedatt--deployment_configs--preview--queue_producers--queue_producer_binding))

<a id="nestedatt--deployment_configs--preview--queue_producers--queue_producer_binding"></a>
### Nested Schema for `deployment_configs.preview.queue_producers.queue_producer_binding`

Optional:

- `name` (String) Name of the Queue.



<a id="nestedatt--deployment_configs--preview--r2_buckets"></a>
### Nested Schema for `deployment_configs.preview.r2_buckets`

Optional:

- `r2_binding` (Attributes) R2 binding. (see [below for nested schema](#nestedatt--deployment_configs--preview--r2_buckets--r2_binding))

<a id="nestedatt--deployment_configs--preview--r2_buckets--r2_binding"></a>
### Nested Schema for `deployment_configs.preview.r2_buckets.r2_binding`

Optional:

- `jurisdiction` (String) Jurisdiction of the R2 bucket.
- `name` (String) Name of the R2 bucket.



<a id="nestedatt--deployment_configs--preview--services"></a>
### Nested Schema for `deployment_configs.preview.services`

Optional:

- `service_binding` (Attributes) Service binding. (see [below for nested schema](#nestedatt--deployment_configs--preview--services--service_binding))

<a id="nestedatt--deployment_configs--preview--services--service_binding"></a>
### Nested Schema for `deployment_configs.preview.services.service_binding`

Optional:

- `entrypoint` (String) The entrypoint to bind to.
- `environment` (String) The Service environment.
- `service` (String) The Service name.



<a id="nestedatt--deployment_configs--preview--vectorize_bindings"></a>
### Nested Schema for `deployment_configs.preview.vectorize_bindings`

Optional:

- `vectorize` (Attributes) Vectorize binding. (see [below for nested schema](#nestedatt--deployment_configs--preview--vectorize_bindings--vectorize))

<a id="nestedatt--deployment_configs--preview--vectorize_bindings--vectorize"></a>
### Nested Schema for `deployment_configs.preview.vectorize_bindings.vectorize`

Optional:

- `index_name` (String)




<a id="nestedatt--deployment_configs--production"></a>
### Nested Schema for `deployment_configs.production`

Optional:

- `ai_bindings` (Attributes) Constellation bindings used for Pages Functions. (see [below for nested schema](#nestedatt--deployment_configs--production--ai_bindings))
- `analytics_engine_datasets` (Attributes) Analytics Engine bindings used for Pages Functions. (see [below for nested schema](#nestedatt--deployment_configs--production--analytics_engine_datasets))
- `browsers` (Attributes) Browser bindings used for Pages Functions. (see [below for nested schema](#nestedatt--deployment_configs--production--browsers))
- `compatibility_date` (String) Compatibility date used for Pages Functions.
- `compatibility_flags` (List of String) Compatibility flags used for Pages Functions.
- `d1_databases` (Attributes) D1 databases used for Pages Functions. (see [below for nested schema](#nestedatt--deployment_configs--production--d1_databases))
- `durable_object_namespaces` (Attributes) Durabble Object namespaces used for Pages Functions. (see [below for nested schema](#nestedatt--deployment_configs--production--durable_object_namespaces))
- `env_vars` (Attributes) Environment variables for build configs. (see [below for nested schema](#nestedatt--deployment_configs--production--env_vars))
- `hyperdrive_bindings` (Attributes) Hyperdrive bindings used for Pages Functions. (see [below for nested schema](#nestedatt--deployment_configs--production--hyperdrive_bindings))
- `kv_namespaces` (Attributes) KV namespaces used for Pages Functions. (see [below for nested schema](#nestedatt--deployment_configs--production--kv_namespaces))
- `mtls_certificates` (Attributes) mTLS bindings used for Pages Functions. (see [below for nested schema](#nestedatt--deployment_configs--production--mtls_certificates))
- `placement` (Attributes) Placement setting used for Pages Functions. (see [below for nested schema](#nestedatt--deployment_configs--production--placement))
- `queue_producers` (Attributes) Queue Producer bindings used for Pages Functions. (see [below for nested schema](#nestedatt--deployment_configs--production--queue_producers))
- `r2_buckets` (Attributes) R2 buckets used for Pages Functions. (see [below for nested schema](#nestedatt--deployment_configs--production--r2_buckets))
- `services` (Attributes) Services used for Pages Functions. (see [below for nested schema](#nestedatt--deployment_configs--production--services))
- `vectorize_bindings` (Attributes) Vectorize bindings used for Pages Functions. (see [below for nested schema](#nestedatt--deployment_configs--production--vectorize_bindings))

<a id="nestedatt--deployment_configs--production--ai_bindings"></a>
### Nested Schema for `deployment_configs.production.ai_bindings`

Optional:

- `ai_binding` (Attributes) AI binding. (see [below for nested schema](#nestedatt--deployment_configs--production--ai_bindings--ai_binding))

<a id="nestedatt--deployment_configs--production--ai_bindings--ai_binding"></a>
### Nested Schema for `deployment_configs.production.ai_bindings.ai_binding`

Optional:

- `project_id` (String)



<a id="nestedatt--deployment_configs--production--analytics_engine_datasets"></a>
### Nested Schema for `deployment_configs.production.analytics_engine_datasets`

Optional:

- `analytics_engine_binding` (Attributes) Analytics Engine binding. (see [below for nested schema](#nestedatt--deployment_configs--production--analytics_engine_datasets--analytics_engine_binding))

<a id="nestedatt--deployment_configs--production--analytics_engine_datasets--analytics_engine_binding"></a>
### Nested Schema for `deployment_configs.production.analytics_engine_datasets.analytics_engine_binding`

Optional:

- `dataset` (String) Name of the dataset.



<a id="nestedatt--deployment_configs--production--browsers"></a>
### Nested Schema for `deployment_configs.production.browsers`

Optional:

- `browser` (String) Browser binding.


<a id="nestedatt--deployment_configs--production--d1_databases"></a>
### Nested Schema for `deployment_configs.production.d1_databases`

Optional:

- `d1_binding` (Attributes) D1 binding. (see [below for nested schema](#nestedatt--deployment_configs--production--d1_databases--d1_binding))

<a id="nestedatt--deployment_configs--production--d1_databases--d1_binding"></a>
### Nested Schema for `deployment_configs.production.d1_databases.d1_binding`

Optional:

- `id` (String) UUID of the D1 database.



<a id="nestedatt--deployment_configs--production--durable_object_namespaces"></a>
### Nested Schema for `deployment_configs.production.durable_object_namespaces`

Optional:

- `do_binding` (Attributes) Durabble Object binding. (see [below for nested schema](#nestedatt--deployment_configs--production--durable_object_namespaces--do_binding))

<a id="nestedatt--deployment_configs--production--durable_object_namespaces--do_binding"></a>
### Nested Schema for `deployment_configs.production.durable_object_namespaces.do_binding`

Optional:

- `namespace_id` (String) ID of the Durabble Object namespace.



<a id="nestedatt--deployment_configs--production--env_vars"></a>
### Nested Schema for `deployment_configs.production.env_vars`

Optional:

- `environment_variable` (Attributes) Environment variable. (see [below for nested schema](#nestedatt--deployment_configs--production--env_vars--environment_variable))

<a id="nestedatt--deployment_configs--production--env_vars--environment_variable"></a>
### Nested Schema for `deployment_configs.production.env_vars.environment_variable`

Optional:

- `type` (String) The type of environment variable (plain text or secret)
- `value` (String) Environment variable value.



<a id="nestedatt--deployment_configs--production--hyperdrive_bindings"></a>
### Nested Schema for `deployment_configs.production.hyperdrive_bindings`

Optional:

- `hyperdrive` (Attributes) Hyperdrive binding. (see [below for nested schema](#nestedatt--deployment_configs--production--hyperdrive_bindings--hyperdrive))

<a id="nestedatt--deployment_configs--production--hyperdrive_bindings--hyperdrive"></a>
### Nested Schema for `deployment_configs.production.hyperdrive_bindings.hyperdrive`

Optional:

- `id` (String)



<a id="nestedatt--deployment_configs--production--kv_namespaces"></a>
### Nested Schema for `deployment_configs.production.kv_namespaces`

Optional:

- `kv_binding` (Attributes) KV binding. (see [below for nested schema](#nestedatt--deployment_configs--production--kv_namespaces--kv_binding))

<a id="nestedatt--deployment_configs--production--kv_namespaces--kv_binding"></a>
### Nested Schema for `deployment_configs.production.kv_namespaces.kv_binding`

Optional:

- `namespace_id` (String) ID of the KV namespace.



<a id="nestedatt--deployment_configs--production--mtls_certificates"></a>
### Nested Schema for `deployment_configs.production.mtls_certificates`

Optional:

- `mtls` (Attributes) mTLS binding. (see [below for nested schema](#nestedatt--deployment_configs--production--mtls_certificates--mtls))

<a id="nestedatt--deployment_configs--production--mtls_certificates--mtls"></a>
### Nested Schema for `deployment_configs.production.mtls_certificates.mtls`

Optional:

- `certificate_id` (String)



<a id="nestedatt--deployment_configs--production--placement"></a>
### Nested Schema for `deployment_configs.production.placement`

Optional:

- `mode` (String) Placement mode.


<a id="nestedatt--deployment_configs--production--queue_producers"></a>
### Nested Schema for `deployment_configs.production.queue_producers`

Optional:

- `queue_producer_binding` (Attributes) Queue Producer binding. (see [below for nested schema](#nestedatt--deployment_configs--production--queue_producers--queue_producer_binding))

<a id="nestedatt--deployment_configs--production--queue_producers--queue_producer_binding"></a>
### Nested Schema for `deployment_configs.production.queue_producers.queue_producer_binding`

Optional:

- `name` (String) Name of the Queue.



<a id="nestedatt--deployment_configs--production--r2_buckets"></a>
### Nested Schema for `deployment_configs.production.r2_buckets`

Optional:

- `r2_binding` (Attributes) R2 binding. (see [below for nested schema](#nestedatt--deployment_configs--production--r2_buckets--r2_binding))

<a id="nestedatt--deployment_configs--production--r2_buckets--r2_binding"></a>
### Nested Schema for `deployment_configs.production.r2_buckets.r2_binding`

Optional:

- `jurisdiction` (String) Jurisdiction of the R2 bucket.
- `name` (String) Name of the R2 bucket.



<a id="nestedatt--deployment_configs--production--services"></a>
### Nested Schema for `deployment_configs.production.services`

Optional:

- `service_binding` (Attributes) Service binding. (see [below for nested schema](#nestedatt--deployment_configs--production--services--service_binding))

<a id="nestedatt--deployment_configs--production--services--service_binding"></a>
### Nested Schema for `deployment_configs.production.services.service_binding`

Optional:

- `entrypoint` (String) The entrypoint to bind to.
- `environment` (String) The Service environment.
- `service` (String) The Service name.



<a id="nestedatt--deployment_configs--production--vectorize_bindings"></a>
### Nested Schema for `deployment_configs.production.vectorize_bindings`

Optional:

- `vectorize` (Attributes) Vectorize binding. (see [below for nested schema](#nestedatt--deployment_configs--production--vectorize_bindings--vectorize))

<a id="nestedatt--deployment_configs--production--vectorize_bindings--vectorize"></a>
### Nested Schema for `deployment_configs.production.vectorize_bindings.vectorize`

Optional:

- `index_name` (String)





<a id="nestedatt--latest_deployment"></a>
### Nested Schema for `latest_deployment`

Read-Only:

- `aliases` (List of String) A list of alias URLs pointing to this deployment.
- `build_config` (String)
- `created_on` (String) When the deployment was created.
- `env_vars` (String) A dict of env variables to build this deploy.
- `environment` (String) Type of deploy.
- `id` (String) Id of the deployment.
- `is_skipped` (Boolean) If the deployment has been skipped.
- `latest_stage` (String)
- `modified_on` (String) When the deployment was last modified.
- `project_id` (String) Id of the project.
- `project_name` (String) Name of the project.
- `short_id` (String) Short Id (8 character) of the deployment.
- `source` (String)
- `stages` (Attributes List) List of past stages. (see [below for nested schema](#nestedatt--latest_deployment--stages))
- `url` (String) The live URL to view this deployment.

<a id="nestedatt--latest_deployment--stages"></a>
### Nested Schema for `latest_deployment.stages`

Optional:

- `name` (String) The current build stage.

Read-Only:

- `ended_on` (String) When the stage ended.
- `started_on` (String) When the stage started.
- `status` (String) State of the current stage.

## Import

!> It is not possible to import a pages project with secret environment variables. If you have a secret environment variable, you must remove it from your project before importing it.

Import is supported using the following syntax:

```shell
$ terraform import cloudflare_pages_project.example <account_id>/<project_name>
```
