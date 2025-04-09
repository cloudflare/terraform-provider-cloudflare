# Releasing

## GitHub (Recommended)

> [!NOTE]
> Releasing with GitHub Actions is currently not usable due to the Golang build cache being larger than the available resources for the free runners. There is work to remedy this but in the meantime, you must follow the manual steps.

- Merge release PR.
- If needed, trigger the resync Terraform registry releases CI job.

## Manual

> [!NOTE]
> Depending on your local Go build cache, you may hit "out of disk space" issues in $TMP" errors. To workaround this, run the release script multiple times while the cache is rebuilt. The script is idempotent and is fine to be run multiple times to get all the artifacts.

- Merge GitHub release PR.
- Load Terraform GPG key into local keychain.
- Set the GPG fingerprint.
  ```
  export GPG_FINGERPRINT="..."
  ```
- Ensure GoReleaser is installed.
- Locally checkout the release tag.
- Run `script/release`.
- Open GitHub release and edit it.
- Upload all binary archives, `terraform-provider-cloudflare_<version>_SHA256SUMS`, `terraform-provider-cloudflare_<version>_SHA256SUMS.sig` and `terraform-provider-cloudflare_<version>_manifest.json` to the GitHub release.
- Trigger the resync Terraform registry releases CI job.
