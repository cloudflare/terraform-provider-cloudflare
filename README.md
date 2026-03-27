# trafgals/cloudflare Terraform Provider

A **community-maintained fork** of the [official Cloudflare Terraform provider](https://registry.terraform.io/providers/cloudflare/cloudflare/latest) with **experimental features**, **faster bug fixes**, and **additional resources**.

## Why This Fork Exists

This fork exists to:

1. **Ship features faster** - No release lag between upstream fixes and this fork
2. **Experimental resources** - Access new Cloudflare features before they're merged upstream
3. **Community-driven fixes** - Direct collaboration without the formal review process
4. **Live-tested changes** - All changes verified against the actual Cloudflare API

## Differences from Official Provider

| Aspect | Official Provider | trafgals/cloudflare |
|--------|------------------|---------------------|
| **Registry** | `cloudflare/cloudflare` | `trafgals/cloudflare` |
| **Release Speed** | Beta releases, slower | Faster iterations |
| **Experimental Features** | Limited | More aggressive inclusion |
| **Bug Fixes** | Through upstream PRs | Direct patches |
| **Acceptance Tests** | Abstracted | Live API testing |

## Current Experimental Features

### `cloudflare_worker` with `builds` Block

The `builds` block enables Workers Builds Git integration for automatic deployments from GitHub repositories.

```hcl
resource "cloudflare_worker" "example" {
  account_id = var.cloudflare_account_id
  name      = "my-worker"
  script     = file("./dist/worker.js")

  # NEW: Workers Builds Git Integration
  builds {
    name       = "default"
    main      = "./src/index.ts"
    directory  = "./"
    build_command = "npx wrangler deploy --dry-run --env production"
    env_vars {
      name  = "ENVIRONMENT"
      value = "production"
    }
    kv_namespace_bindings {
      namespace_id = cloudflare_workers_kv_namespace.example.id
      binding       = "KV"
    }
  }
}
```

### Resources with Live API Testing

All resources include acceptance tests run against the **live Cloudflare API** to ensure reliability:

- `cloudflare_ai_gateway`
- `cloudflare_secrets_store`
- `cloudflare_secrets_store_secret`

## Usage

```hcl
terraform {
  required_providers {
    cloudflare = {
      source  = "trafgals/cloudflare"
      version = "~> 1.0"
    }
  }
}

provider "cloudflare" {
  api_token = var.cloudflare_api_token
}
```

Or use environment variables:

```bash
export CLOUDFLARE_API_TOKEN="your-api-token"
export CLOUDFLARE_ACCOUNT_ID="your-account-id"
```

## Installation

Terraform automatically downloads providers from the Terraform Registry. No additional installation steps required.

## Status

**v1.0.59** - Includes:
- ✅ `builds` block for `cloudflare_worker` (fixes #6924)
- ✅ Worker auth fix for API tokens
- ✅ `ai_gateway` resource with live-tested acceptance tests
- ✅ `secrets_store` resource with live-tested acceptance tests  
- ✅ `secrets_store_secret` resource with live-tested acceptance tests

## Relationship to Official Provider

This is **not** an official Cloudflare product. It's a community fork that:

- **Tracks** the official upstream repository
- **Patches** experimental features and urgent fixes
- **Rebases** on official releases when appropriate

For production use of standard features, consider the [official provider](https://registry.terraform.io/providers/cloudflare/cloudflare/latest).

## Reporting Issues

- **This fork**: [GitHub Issues](https://github.com/trafgals/terraform-provider-cloudflare/issues)
- **Official provider**: [cloudflare/terraform-provider-cloudflare](https://github.com/cloudflare/terraform-provider-cloudflare/issues)

## Building from Source

```bash
# Clone the repository
git clone https://github.com/trafgals/terraform-provider-cloudflare.git
cd terraform-provider-cloudflare

# Install dependencies
./scripts/bootstrap

# Build
go build -o terraform-provider-cloudflare

# Run tests (requires Cloudflare credentials)
export CLOUDFLARE_API_TOKEN="your-token"
export CLOUDFLARE_ACCOUNT_ID="your-account-id"
go test ./internal/services/... -v -timeout 30m
```

## License

Same as the official provider - MPL 2.0
