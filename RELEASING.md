# Releasing the v3 security rebuild

The `Release` workflow runs when a `v*` tag is pushed. It uses the workflow and
configuration at the tagged commit, builds with Go 1.27.3, signs the checksums,
and publishes the GitHub release. The Terraform Registry then ingests the release
through the repository's release webhook.

## Before tagging

1. Merge the release PR into `3.x`. Confirm the release commit includes the
   workflow, GoReleaser configuration, changelog, and `scripts/release-notes.md`.
2. Set the `3.35.1` changelog date to the actual release date before tagging.
   Confirm the release notes accurately describe the security rebuild and the
   removal of the Windows ARM32 artifact.
3. Validate with GoReleaser **2.18.3** and Go **1.27.3**:

   ```sh
   goreleaser check
   GOTOOLCHAIN=go1.27.3 goreleaser release --snapshot --clean --skip=publish,sign,announce --parallelism=2
   GOTOOLCHAIN=go1.27.3 go test -mod=readonly ./...
   git diff --exit-code -- go.mod go.sum
   ```

   The legacy `make test` invokes the removed `go test -i` flag; use `go test`
   directly. Investigate test and vet failures separately. Snapshot builds do
   not exercise production signing or publication.
4. Confirm 13 archive targets: Darwin AMD64/ARM64; FreeBSD and Linux
   386/AMD64/ARM/ARM64; Windows 386/AMD64/ARM64. Inspect representative binaries
   with `go version -m`; they must report Go 1.27.3. Load the native binary with
   Terraform and run `terraform providers schema -json`.
5. Confirm the repository GPG secrets are valid, the signing public key is
   registered with the Terraform Registry, and the Registry release webhook is
   active. `release.make_latest: false` must remain set for this older major.

## Tag the reviewed commit

Fetch the branch and tags, then inspect and record the exact release SHA:

```sh
git fetch origin 3.x --tags
git log -1 origin/3.x
git show <verified-release-commit-sha>:.github/workflows/release.yml
git tag -a v3.35.1 <verified-release-commit-sha> -m "Release v3.35.1"
git push origin refs/tags/v3.35.1
```

Use maintainer credentials to push the tag. A tag pushed using a workflow's
`GITHUB_TOKEN` generally does not trigger another workflow. Pushing the tag
starts publication; merging the PR alone does not.

## Watch the release

- Confirm the Actions run is for `v3.35.1` at the recorded SHA, and the compiler
  verification step reports Go 1.27.3.
- Confirm GPG import, all builds, signing, and uploads succeed. Inspect the
  release even if a later announcement step fails; it may already be published.
- Confirm the release is public, stable, has the prepared notes, and has not
  displaced the existing GitHub "Latest" release.
- Expect 13 ZIPs plus `terraform-provider-cloudflare_3.35.1_manifest.json`,
  `terraform-provider-cloudflare_3.35.1_SHA256SUMS`, and
  `terraform-provider-cloudflare_3.35.1_SHA256SUMS.sig`.
- Download the assets. Verify SHA-256 checksums for the ZIPs and manifest, and
  verify the detached checksum signature against the registered signing key.
  Inspect downloaded binaries with `go version -m` for Go 1.27.3 and the source
  SHA. The manifest must retain protocol version `6.0`.
- Check the Registry webhook delivery and that version `3.35.1` appears in the
  Registry. In a clean Terraform directory, require
  `cloudflare/cloudflare` at exactly `= 3.35.1`, run `terraform init`, and then
  `terraform providers schema -json`. Check the installed version, signature
  verification, and downloaded binary's compiler metadata.

If the job fails before publication, inspect partial drafts/assets before
retrying. Never move a published tag or replace published binaries: use a new
patch version for corrections to released artifacts. For Registry ingestion
failures, check webhook responses, asset names, and signing-key registration;
use Registry resync after correcting the cause when needed.
