# Release Instructions for makeqr v1.0.0

## CI Fixes Applied

The following issues in the CI workflow have been fixed:

### 1. `.golangci.yml` Configuration
- **Issue**: `version: 2` was a number instead of string, causing golangci-lint v2.6.2 validation failure
- **Fix**: Changed to `version: "2"`
- **Issue**: `exclude-use-default: false` is deprecated in golangci-lint v2.6.2
- **Fix**: Removed the deprecated option

### 2. `.github/workflows/release.yml` Updates
- Updated `actions/setup-go` from v4 to v6
- Quoted Go version for consistency: `go-version: "1.25"`
- Updated `goreleaser/goreleaser-action` from v4 to v6
- Updated deprecated flag from `--rm-dist` to `release --clean`

### 3. Added `.goreleaser.yml` Configuration
Created a complete goreleaser configuration supporting:
- Cross-platform builds (Linux, macOS, Windows)
- Multi-architecture (amd64, arm64)
- Proper archive formats (tar.gz for Unix, zip for Windows)
- Checksum generation
- Automatic changelog creation

## Testing the CI Fix

### Option 1: Approve the PR Workflow
1. Go to the GitHub Actions tab
2. Find the CI workflow run for this PR
3. Approve the workflow run (required for workflows from bot accounts)
4. Verify all steps pass: golangci-lint, tests, vet, and build

### Option 2: Merge to Main
1. Merge this PR to the main branch
2. The CI workflow will run automatically on the push to main
3. Verify the workflow completes successfully

### Local Verification (Already Completed)
```bash
# Tests pass
go test ./...
# Output: ok makeqr/internal/cli, ok makeqr/internal/qr

# Linting passes
golangci-lint run --timeout=5m
# Output: 0 issues

# Vet passes
go vet ./...
# Output: (no errors)

# Build succeeds
go build ./...
# Output: (success)
```

## Creating the v1.0.0 Release

Once the CI is verified working on main:

### Step 1: Create and Push the Tag

```bash
# Ensure you're on main with latest changes
git checkout main
git pull origin main

# Create annotated tag for v1.0.0
git tag -a v1.0.0 -m "Release v1.0.0: Initial stable release of makeqr CLI"

# Push the tag to trigger release workflow
git push origin v1.0.0
```

### Step 2: Monitor the Release Workflow

1. Go to GitHub Actions tab
2. Watch for the "Release" workflow to start
3. The workflow will:
   - Check out the code
   - Set up Go 1.25
   - Run goreleaser to:
     - Build binaries for all platforms/architectures
     - Create release archives
     - Generate checksums
     - Create a GitHub Release with changelog
     - Attach all build artifacts

### Step 3: Verify the Release

1. Go to the Releases page on GitHub
2. Verify v1.0.0 release is created
3. Check that all platform binaries are attached:
   - `makeqr_Linux_x86_64.tar.gz`
   - `makeqr_Linux_arm64.tar.gz`
   - `makeqr_Darwin_x86_64.tar.gz`
   - `makeqr_Darwin_arm64.tar.gz`
   - `makeqr_Windows_x86_64.zip`
   - `makeqr_Windows_arm64.zip`
   - `checksums.txt`

## Expected Release Artifacts

The v1.0.0 release will include:
- Multi-platform binaries (Linux, macOS, Windows)
- Multiple architectures (amd64, arm64)
- SHA256 checksums
- Automatic changelog from commit history
- Source code archives (zip and tar.gz)

## Troubleshooting

### If CI Still Fails
- Check the golangci-lint version matches v2.6.2
- Verify Go version 1.25.4 is available in the runner
- Review the workflow logs for specific errors

### If Release Fails
- Ensure GITHUB_TOKEN has release permissions
- Check goreleaser.yml syntax
- Verify the tag follows semantic versioning (v*.*.*)