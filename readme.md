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

## Build steps

```
go build -ldflags="-s -w" -o pack .
tar -czf pack-<version>.tar.gz pack
shasum -a 256 pack-<version>.tar.gz
```

Releases are automated via GitHub Actions — pushing a `v*` tag triggers a build and attaches the tarball to the release.
