# Pack

![An example of Pack](https://user-images.githubusercontent.com/449385/31547153-322cd1b8-b01e-11e7-8810-4c0caaa55401.png)

A tool for viewing the available script commands in a `package.json` file.

## Installing ([via Brew](https://brew.sh/))

```
brew install superhighfives/tools/pack
```

Once installed, run `pack` in any directory with a `package.json` file:

```
Available script commands in package.json
-----------------------------------------
     start react-scripts start
start-prod yarn build && node app.js
     build react-scripts build
      test react-scripts test --env=jsdom
     eject react-scripts eject
      flow flow
```

🍻

## Development

Pack is a Go project. You'll need Go installed — the version is specified in `go.mod`.

### Setup

```
git clone https://github.com/superhighfives/pack.git
cd pack
go mod download
```

### Running locally

```
go run .
```

This will look for a `package.json` in the current directory. The repo includes one as a test fixture.

### Testing

```
go test ./...
```

### Building

```
go build -ldflags="-s -w" -o pack .
```

The `-ldflags="-s -w"` strips debug symbols to reduce binary size.

## Continuous integration

Every push and pull request runs the CI workflow (`.github/workflows/ci.yml`),
which checks formatting (`gofmt`), runs `go vet`, builds, and runs the tests.

## Releases

Releases are automated via GitHub Actions. The recommended way to cut a release
is to bump the version:

1. Update the `VERSION` file (e.g. `1.2.3`) in a pull request and merge it to
   `main`.

When CI passes on `main`, it checks whether a `v<VERSION>` tag already exists.
If it doesn't, CI calls the release workflow to build, publish, and update the
Homebrew formula.

You can also trigger a release manually by pushing a `v*` tag (e.g. `v1.2.3`):

```
git tag v1.2.3
git push origin v1.2.3
```

Either path runs the release workflow (`.github/workflows/release.yml`), which:

1. **Builds** the binary on macOS with stripped symbols
2. **Packages** it into a tarball (`pack-<version>.tar.gz`) with a SHA256 checksum file
3. **Creates a GitHub Release** with auto-generated release notes and the tarball attached
4. **Opens a PR** against the [homebrew-tools](https://github.com/superhighfives/homebrew-tools) tap to update the `pack.rb` formula with the new version, URL, and SHA256

The Homebrew PR requires a `HOMEBREW_TAP_TOKEN` secret configured in the repository settings.

Once the Homebrew PR is merged, users can upgrade via `brew upgrade pack`.
