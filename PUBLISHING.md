# Publishing Guide

This guide walks you through publishing `goforge` so anyone can install it globally.

## Prerequisites

- GitHub account
- Git configured with your credentials
- Go 1.26.2+ installed

## Step-by-Step Publishing

### 1. Update Module Path

First, update the module path in these files to match your GitHub username:

**Files to update:**
- `go.mod` - Change `github.com/viveksharma/goforge` to `github.com/YOUR_USERNAME/goforge`
- `.goreleaser.yaml` - Update the `owner` field under `release.github`

```bash
# Quick find and replace (macOS/Linux):
find . -type f \( -name "*.go" -o -name "go.mod" -o -name ".goreleaser.yaml" \) \
  -exec sed -i '' 's/viveksharma/YOUR_USERNAME/g' {} +
```

### 2. Initialize Git Repository

```bash
cd /Users/viveksharma/Documents/goforge
git init
git add .
git commit -m "Initial commit: goforge v0.1.0"
```

### 3. Create GitHub Repository

**Option A: Using GitHub CLI**
```bash
gh repo create goforge --public --source=. --remote=origin --push
```

**Option B: Manual**
1. Go to https://github.com/new
2. Repository name: `goforge`
3. Make it **Public**
4. Don't initialize with README (we already have one)
5. Click "Create repository"

Then:
```bash
git remote add origin https://github.com/YOUR_USERNAME/goforge.git
git branch -M main
git push -u origin main
```

### 4. Create Your First Release

```bash
# Tag the release
git tag v0.1.0

# Push the tag (this triggers GitHub Actions)
git push origin v0.1.0
```

**GitHub Actions** will automatically:
- Build binaries for macOS, Linux, Windows (AMD64 & ARM64)
- Create a GitHub Release with all binaries
- Generate checksums

### 5. Test Installation

After the GitHub Action completes (~2 minutes):

```bash
# Anyone can now install:
go install github.com/YOUR_USERNAME/goforge@latest

# Verify:
goforge version
```

## Distribution Methods

### Method 1: Go Install (Built-in)
✅ **Already works!** Just share:
```bash
go install github.com/YOUR_USERNAME/goforge@latest
```

### Method 2: Direct Binary Download
Users can download from: `https://github.com/YOUR_USERNAME/goforge/releases`

### Method 3: Homebrew (macOS/Linux)

Create a separate tap repository:

```bash
# Create formula repository
gh repo create homebrew-tap --public

# Create formula file
cat > goforge.rb << 'EOF'
class Goforge < Formula
  desc "Forge production-ready Go applications"
  homepage "https://github.com/YOUR_USERNAME/goforge"
  url "https://github.com/YOUR_USERNAME/goforge/archive/v0.1.0.tar.gz"
  sha256 "REPLACE_WITH_ACTUAL_SHA256"
  license "MIT"

  depends_on "go" => :build

  def install
    system "go", "build", *std_go_args(ldflags: "-s -w"), "./cmd/goforge"
  end

  test do
    assert_match "goforge version", shell_output("#{bin}/goforge version")
  end
end
EOF

git add goforge.rb
git commit -m "Add goforge formula"
git push
```

Users install via:
```bash
brew tap YOUR_USERNAME/tap
brew install goforge
```

### Method 4: Docker Image (Optional)

```dockerfile
# Add to your repo
FROM golang:1.26.2-alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o goforge ./cmd/goforge

FROM alpine:latest
COPY --from=builder /app/goforge /usr/local/bin/
ENTRYPOINT ["goforge"]
```

Users run:
```bash
docker run YOUR_USERNAME/goforge create myapp
```

## Updating & Versioning

### Release a New Version

```bash
# Make your changes
git add .
git commit -m "feat: add new feature"

# Bump version
git tag v0.2.0
git push origin main
git push origin v0.2.0
```

### Versioning Strategy

Use semantic versioning:
- `v0.x.x` - Pre-1.0 (breaking changes allowed)
- `v1.0.0` - First stable release
- `v1.1.0` - New features (backward compatible)
- `v1.1.1` - Bug fixes

## Marketing Your Tool

### 1. Add Topics to GitHub
Go to your repo → Settings → Topics
Add: `go`, `cli`, `template`, `generator`, `golang`, `fiber`, `boilerplate`

### 2. Create a Good README
- Clear installation instructions
- GIF/screenshots of usage
- Feature list with emojis
- Badges (build status, go version, license)

### 3. Share It
- Reddit: r/golang
- Hacker News: news.ycombinator.com
- Dev.to: Write an article
- Twitter/X: Post with #golang hashtag
- awesome-go: Submit PR to https://github.com/avelino/awesome-go

### 4. Add Badges to README

```markdown
[![Go Report Card](https://goreportcard.com/badge/github.com/YOUR_USERNAME/go-template)](https://goreportcard.com/report/github.com/YOUR_USERNAME/go-template)
[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](https://opensource.org/licenses/MIT)
[![Release](https://img.shields.io/github/release/YOUR_USERNAME/go-template.svg)](https://github.com/YOUR_USERNAME/go-template/releases)
```

## Troubleshooting

### Go Install Shows "Not Found"
- Ensure repository is **public**
- Wait 5-10 minutes after pushing (Go proxy cache)
- Force update: `GOPROXY=direct go install github.com/YOUR_USERNAME/go-template@latest`

### GitHub Actions Failing
- Check you have GitHub Actions enabled in repo settings
- Verify `.github/workflows/release.yml` exists
- Check Actions tab for error logs

### Users Can't Install
Make sure:
1. Repository is public
2. `go.mod` has correct module path
3. You've pushed at least one tag (e.g., `v0.1.0`)

## Success Metrics

Track these to measure adoption:
- GitHub Stars ⭐
- Go Proxy Downloads: https://pkg.go.dev/github.com/YOUR_USERNAME/go-template
- GitHub Releases download count
- Issues/PRs (community engagement)

---

**You're ready to publish!** 🚀

Start with Step 1 and work through each section. In ~10 minutes, your tool will be installable worldwide.
