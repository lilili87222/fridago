# frida-go 



### frida golang binding(no cgo implement)

### Supports system(If necessary, he can support all platforms)
- windows x64
- linux
- macos
- posix

### install

```
go get -u github.com/a97077088/frida-go
```

### build
```
go build
```

### single exe build
```
go build -tags tempdll
```


#### The code style is based on [swift-frida](https://github.com/frida/frida-swift)

#### dylib from the project [ying32/dylib](https://github.com/ying32/dylib)

#### tempdll from the project [ying32/govcl](https://github.com/ying32/govcl)