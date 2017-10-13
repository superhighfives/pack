# Pack

![An example of Pack](https://user-images.githubusercontent.com/449385/31547153-322cd1b8-b01e-11e7-8810-4c0caaa55401.png)


A tool for viewing the contents of a `package.json` file.

## Installing ([via Brew](https://brew.sh/))

```
brew install superhighfives/tools/pack
```

## Run

Then just run `pack` in any directory with a `package.json` file:

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
go build pack.go
tar -czf pack-0.0.1.tar.gz pack
shasum -a 256 pack-0.0.1.tar.gz
```
