# Pack

A tool for viewing the contents of a `package.json` file.

## Build steps

```
go build pack.go
tar -czf pack-0.0.1.tar.gz pack
shasum -a 256 pack-0.0.1.tar.gz
```