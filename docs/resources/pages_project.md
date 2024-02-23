---
page_title: "cloudflare_pages_project Resource - Cloudflare"
subcategory: ""
description: |-
  Provides a resource which manages Cloudflare Pages projects.
---

# cloudflare_pages_project (Resource)

Provides a resource which manages Cloudflare Pages projects.

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
  build_config {
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
  source {
    type = "github"
    config {
      owner                         = "cloudflare"
      repo_name                     = "ninjakittens"
      production_branch             = "main"
      pr_comments_enabled           = true
      deployments_enabled           = true
      production_deployment_enabled = true
      preview_deployment_setting    = "custom"
      preview_branch_includes       = ["dev", "preview"]
      preview_branch_excludes       = ["main", "prod"]
    }
  }
}

# Pages project managing deployment configs
resource "cloudflare_pages_project" "deployment_configs" {
  account_id        = "f037e56e89293a057740de681ac9abbe"
  name              = "this-is-my-project-01"
  production_branch = "main"
  deployment_configs {
    preview {
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
    }
    production {
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
    }
  }
}

# Pages project managing all configs
resource "cloudflare_pages_project" "deployment_configs" {
  account_id        = "f037e56e89293a057740de681ac9abbe"
  name              = "this-is-my-project-01"
  production_branch = "main"

  source {
    type = "github"
    config {
      owner                         = "cloudflare"
      repo_name                     = "ninjakittens"
      production_branch             = "main"
      pr_comments_enabled           = true
      deployments_enabled           = true
      production_deployment_enabled = true
      preview_deployment_setting    = "custom"
      preview_branch_includes       = ["dev", "preview"]
      preview_branch_excludes       = ["main", "prod"]
    }
  }

  build_config {
    build_command       = "npm run build"
    destination_dir     = "build"
    root_dir            = ""
    web_analytics_tag   = "cee1c73f6e4743d0b5e6bb1a0bcaabcc"
    web_analytics_token = "021e1057c18547eca7b79f2516f06o7x"
  }

  deployment_configs {
    preview {
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
    }
    production {
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
    }
  }
}
```
<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `account_id` (String) The account identifier to target for the resource.
- `name` (String) Name of the project.
- `production_branch` (String) The name of the branch that is used for the production environment.

### Optional

- `build_config` (Block List, Max: 1) Configuration for the project build process. Read more about the build configuration in the [developer documentation](https://developers.cloudflare.com/pages/platform/build-configuration). (see [below for nested schema](#nestedblock--build_config))
- `deployment_configs` (Block List, Max: 1) Configuration for deployments in a project. (see [below for nested schema](#nestedblock--deployment_configs))
- `source` (Block List, Max: 1) Configuration for the project source. Read more about the source configuration in the [developer documentation](https://developers.cloudflare.com/pages/platform/branch-build-controls/). (see [below for nested schema](#nestedblock--source))

### Read-Only

- `created_on` (String) When the project was created.
- `domains` (List of String) A list of associated custom domains for the project.
- `id` (String) The ID of this resource.
- `subdomain` (String) The Cloudflare subdomain associated with the project.

<a id="nestedblock--build_config"></a>
### Nested Schema for `build_config`

Optional:

- `build_caching` (Boolean) Enable build caching for the project.
- `build_command` (String) Command used to build project.
- `destination_dir` (String) Output directory of the build.
- `root_dir` (String) Your project's root directory, where Cloudflare runs the build command. If your site is not in a subdirectory, leave this path value empty.
- `web_analytics_tag` (String) The classifying tag for analytics.
- `web_analytics_token` (String) The auth token for analytics.


<a id="nestedblock--deployment_configs"></a>
### Nested Schema for `deployment_configs`

Optional:

- `preview` (Block List, Max: 1) Configuration for preview deploys. (see [below for nested schema](#nestedblock--deployment_configs--preview))
- `production` (Block List, Max: 1) Configuration for production deploys. (see [below for nested schema](#nestedblock--deployment_configs--production))

<a id="nestedblock--deployment_configs--preview"></a>
### Nested Schema for `deployment_configs.preview`

Optional:

- `always_use_latest_compatibility_date` (Boolean) Use latest compatibility date for Pages Functions. Defaults to `false`.
- `compatibility_date` (String) Compatibility date used for Pages Functions.
- `compatibility_flags` (List of String) Compatibility flags used for Pages Functions.
- `d1_databases` (Map of String) D1 Databases used for Pages Functions. Defaults to `map[]`.
- `durable_object_namespaces` (Map of String) Durable Object namespaces used for Pages Functions. Defaults to `map[]`.
- `environment_variables` (Map of String) Environment variables for Pages Functions. Defaults to `map[]`.
- `fail_open` (Boolean) Fail open used for Pages Functions. Defaults to `false`.
- `kv_namespaces` (Map of String) KV namespaces used for Pages Functions. Defaults to `map[]`.
- `placement` (Block List, Max: 1) Configuration for placement in the Cloudflare Pages project. (see [below for nested schema](#nestedblock--deployment_configs--preview--placement))
- `r2_buckets` (Map of String) R2 Buckets used for Pages Functions. Defaults to `map[]`.
- `secrets` (Map of String, Sensitive) Encrypted environment variables for Pages Functions. Defaults to `map[]`.
- `service_binding` (Block Set) Services used for Pages Functions. (see [below for nested schema](#nestedblock--deployment_configs--preview--service_binding))
- `usage_model` (String) Usage model used for Pages Functions. Available values: `unbound`, `bundled`, `standard`. Defaults to `bundled`.

<a id="nestedblock--deployment_configs--preview--placement"></a>
### Nested Schema for `deployment_configs.preview.placement`

Optional:

- `mode` (String) Placement Mode for the Pages Function.


<a id="nestedblock--deployment_configs--preview--service_binding"></a>
### Nested Schema for `deployment_configs.preview.service_binding`

Required:

- `name` (String) The global variable for the binding in your Worker code.
- `service` (String) The name of the Worker to bind to.

Optional:

- `environment` (String) The name of the Worker environment to bind to.



<a id="nestedblock--deployment_configs--production"></a>
### Nested Schema for `deployment_configs.production`

Optional:

- `always_use_latest_compatibility_date` (Boolean) Use latest compatibility date for Pages Functions. Defaults to `false`.
- `compatibility_date` (String) Compatibility date used for Pages Functions.
- `compatibility_flags` (List of String) Compatibility flags used for Pages Functions.
- `d1_databases` (Map of String) D1 Databases used for Pages Functions. Defaults to `map[]`.
- `durable_object_namespaces` (Map of String) Durable Object namespaces used for Pages Functions. Defaults to `map[]`.
- `environment_variables` (Map of String) Environment variables for Pages Functions. Defaults to `map[]`.
- `fail_open` (Boolean) Fail open used for Pages Functions. Defaults to `false`.
- `kv_namespaces` (Map of String) KV namespaces used for Pages Functions. Defaults to `map[]`.
- `placement` (Block List, Max: 1) Configuration for placement in the Cloudflare Pages project. (see [below for nested schema](#nestedblock--deployment_configs--production--placement))
- `r2_buckets` (Map of String) R2 Buckets used for Pages Functions. Defaults to `map[]`.
- `secrets` (Map of String, Sensitive) Encrypted environment variables for Pages Functions. Defaults to `map[]`.
- `service_binding` (Block Set) Services used for Pages Functions. (see [below for nested schema](#nestedblock--deployment_configs--production--service_binding))
- `usage_model` (String) Usage model used for Pages Functions. Available values: `unbound`, `bundled`, `standard`. Defaults to `bundled`.

<a id="nestedblock--deployment_configs--production--placement"></a>
### Nested Schema for `deployment_configs.production.placement`

Optional:

- `mode` (String) Placement Mode for the Pages Function.


<a id="nestedblock--deployment_configs--production--service_binding"></a>
### Nested Schema for `deployment_configs.production.service_binding`

Required:

- `name` (String) The global variable for the binding in your Worker code.
- `service` (String) The name of the Worker to bind to.

Optional:

- `environment` (String) The name of the Worker environment to bind to.




<a id="nestedblock--source"></a>
### Nested Schema for `source`

Optional:

- `config` (Block List, Max: 1) Configuration for the source of the Cloudflare Pages project. (see [below for nested schema](#nestedblock--source--config))
- `type` (String) Project host type.

<a id="nestedblock--source--config"></a>
### Nested Schema for `source.config`

Required:

- `production_branch` (String) Project production branch name.

Optional:

- `deployments_enabled` (Boolean) Toggle deployments on this repo. Defaults to `true`.
- `owner` (String) Project owner username. **Modifying this attribute will force creation of a new resource.**
- `pr_comments_enabled` (Boolean) Enable Pages to comment on Pull Requests. Defaults to `true`.
- `preview_branch_excludes` (List of String) Branches will be excluded from automatic deployment.
- `preview_branch_includes` (List of String) Branches will be included for automatic deployment.
- `preview_deployment_setting` (String) Preview Deployment Setting. Available values: `custom`, `all`, `none`. Defaults to `all`.
- `production_deployment_enabled` (Boolean) Enable production deployments. Defaults to `true`.
- `repo_name` (String) Project repository name. **Modifying this attribute will force creation of a new resource.**

## Import

!> It is not possible to import a pages project with secret environment variables. If you have a secret environment variable, you must remove it from your project before importing it.

Import is supported using the following syntax:

```shell
$ terraform import cloudflare_pages_project.example <account_id>/<project_name>
```
