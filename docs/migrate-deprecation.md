## Deprecation: cmd/migrate

The `cmd/migrate` CLI included in the Cloudflare Terraform Provider is now **deprecated** and will be removed in a future release.

### What changed

Historically, `cmd/migrate` was used to assist with Terraform configuration and state migrations between provider versions.

This functionality has been replaced by a dedicated and actively maintained tool:

- https://github.com/cloudflare/tf-migrate

### What you should do

All users should migrate to `tf-migrate` for any Terraform v4 → v5 (or future) migrations.

`tf-migrate` provides:

- Full configuration transformation support
- State-safe migrations via `moved {}` / `import {}` blocks
- Cross-file reference rewriting
- Ongoing support and improvements

### Timeline

- `cmd/migrate` is deprecated as of this release
- It will be removed in a future release (date TBD)

### Notes

- No new features or fixes will be added to `cmd/migrate`
- Bugs in `cmd/migrate` may not be addressed

If you are currently using `cmd/migrate`, plan your migration to `tf-migrate` as soon as possible.
