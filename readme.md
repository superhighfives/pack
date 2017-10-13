# Pack

![An example of Pack](https://user-images.githubusercontent.com/449385/31547153-322cd1b8-b01e-11e7-8810-4c0caaa55401.png)


A tool for viewing the contents of a `package.json` file.

## Build steps

```
go build pack.go
tar -czf pack-0.0.1.tar.gz pack
shasum -a 256 pack-0.0.1.tar.gz
```
