# Sweeper Linter CI Integration

## Current Status: ✅ Linter Runs in CI (Warning Mode)

The sweeper linter is **already integrated** into your CI pipeline and runs on every push and pull request.

### How It Works

```
GitHub PR/Push
    ↓
CI Workflow (.github/workflows/ci.yml)
    ↓
Runs: ./scripts/lint
    ↓
Runs: go build ./...              (fails on compile errors)
    ↓
Runs: ./scripts/lint-sweepers     (currently warnings only)
    ↓
CI Result: ✅ PASS
```

### What You'll See in CI

```bash
==> Running sweeper linter
Found 146 violation(s) in 158 file(s):

⚠ internal/services/account/resource_test.go:32: Sweeper should use utils.ShouldSweepResource()
⚠ internal/services/access_rule/resource_test.go:1: Missing sweeper registration
...

Summary: 146 warning(s) in 158 file(s)
```

**CI Status:** ✅ Passes (warnings don't fail the build)

## Making It Blocking (When Ready)

When you've migrated enough resources and want the linter to **fail CI on violations**, change one line:

**File:** `internal/tools/sweeper-lint/main.go:63`

**Change:**
```go
// FROM (current):
os.Exit(0) // Exit 0 for now (warnings only)

// TO (blocking):
os.Exit(1) // Fail CI on violations
```

Then:
```bash
# Rebuild and commit
git add internal/tools/sweeper-lint/main.go
git commit -m "Enable blocking mode for sweeper linter"
```

**After this change:**
- ❌ CI will **fail** if any violations are found
- ✅ CI will **pass** only when all sweepers are compliant

## Migration Strategy

### Phase 1: Warning Mode (Current) ✅
**Status:** Active
**Behavior:** Linter runs, shows violations, CI passes
**Goal:** Give visibility into what needs fixing

```
PR merges → CI shows warnings → Team sees violations → No blocking
```

### Phase 2: Gradual Migration
**Status:** In Progress
**Goal:** Fix violations incrementally

Recommended approach:
1. Fix 5-10 resources per PR
2. Run linter locally: `./scripts/lint-sweepers`
3. Verify violations decrease
4. Track progress: Currently **146 violations** across **158 files**

### Phase 3: Blocking Mode (Future)
**Status:** Not Yet Active
**Trigger:** When violations drop to acceptable level (suggest <20)
**Action:** Change `os.Exit(0)` → `os.Exit(1)`

```
PR with violation → CI fails ❌ → Must fix before merge
```

## Checking Linter Locally

Before pushing code, run the linter locally:

```bash
# Run just the sweeper linter
./scripts/lint-sweepers

# Run all lints (includes sweeper linter)
./scripts/lint

# Check specific file
go run ./internal/tools/sweeper-lint ./internal/services/account_token/
```

## Monitoring Progress

Track violation count over time:

```bash
# Current violations
./scripts/lint-sweepers 2>&1 | grep "Summary:"
# Output: Summary: 146 warning(s) in 158 file(s)

# After migration
./scripts/lint-sweepers 2>&1 | grep "Summary:"
# Target: Summary: 0 warning(s) in 158 file(s)
```

## CI Configuration

### Current Setup

**.github/workflows/ci.yml:**
```yaml
jobs:
  lint:
    runs-on: ...
    steps:
      - uses: actions/checkout@v4
      - name: Setup go
        uses: actions/setup-go@v5
        with:
          go-version-file: ./go.mod
      - name: Bootstrap
        run: ./scripts/bootstrap
      - name: Run lints              # ← This runs the sweeper linter
        run: ./scripts/lint
```

**scripts/lint:**
```bash
#!/usr/bin/env bash
set -e
cd "$(dirname "$0")/.."

echo "==> Running Go build"
go build ./...

echo ""
echo "==> Running sweeper linter"
./scripts/lint-sweepers              # ← Sweeper linter runs here
```

**scripts/lint-sweepers:**
```bash
#!/usr/bin/env bash
set -e
cd "$(dirname "$0")/.."

echo "==> Building sweeper linter"
go build -o /tmp/sweeper-lint ./internal/tools/sweeper-lint

echo "==> Running sweeper linter"
/tmp/sweeper-lint ./internal/services

rm -f /tmp/sweeper-lint
```

## FAQ

### Q: Will this break my CI right now?
**A:** No. The linter is in warning mode (`os.Exit(0)`), so CI passes even with violations.

### Q: When should I enable blocking mode?
**A:** When you're comfortable blocking PRs on violations. Suggested threshold: <20 violations remaining.

### Q: Can I disable the linter temporarily?
**A:** Yes, comment out line 12 in `scripts/lint`:
```bash
# echo "==> Running sweeper linter"
# ./scripts/lint-sweepers
```

### Q: How do I see what violations are in my PR?
**A:** Check the CI logs in your PR. The linter output will show all violations found.

### Q: Can I exclude specific files from linting?
**A:** Currently no, but you can modify `internal/tools/sweeper-lint/main.go` to skip certain paths.

### Q: Does the linter slow down CI?
**A:** Minimal impact. The linter runs in ~3-5 seconds for the entire codebase.

## Example CI Output

### Successful Run (Current - Warning Mode)
```
Run ./scripts/lint
==> Running Go build
==> Running sweeper linter
==> Building sweeper linter
==> Running sweeper linter
Found 146 violation(s) in 158 file(s):

⚠ internal/services/account/resource_test.go:32: Sweeper should use utils.ShouldSweepResource()
⚠ internal/services/access_rule/resource_test.go:1: Missing sweeper registration
...

Summary: 146 warning(s) in 158 file(s)
✅ CI Passed
```

### After Enabling Blocking Mode
```
Run ./scripts/lint
==> Running Go build
==> Running sweeper linter
==> Building sweeper linter
==> Running sweeper linter
Found 5 violation(s) in 5 file(s):

⚠ internal/services/foo/resource_test.go:42: Sweeper should use utils.ShouldSweepResource()
...

Summary: 5 error(s) in 5 file(s)
❌ Error: Process completed with exit code 1.
```

## Next Steps

1. **Keep migrating resources** - Use the linter output as a todo list
2. **Monitor progress** - Track violations decreasing over time
3. **Enable blocking** - When ready, flip to `os.Exit(1)`
4. **Enjoy enforcement** - All future sweepers must be compliant

## Summary

✅ Linter is already running in CI (warning mode)
✅ No action needed for it to run
✅ Shows violations without blocking PRs
✅ One-line change to enable blocking when ready
✅ Helps enforce standards as team migrates resources
