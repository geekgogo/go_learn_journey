# 🐱 如何使用 go mod

在项目文件夹中使用命令 `go mod init 项目名称`

🐶 比如这个项目文件夹名称为 39gomod

```
go mod init 39gomod
```

然后文件夹下会产生一个 go.mod 的文件

# 🍙 go mod 引入第三方包

## 有三种方法

1.  go get -u 包名称

    `go get -u github.com/shopspring/decimal`

2.  go mod download

    此方法不用先下载包，在代码中引入完成后，执行`go mod download`，系统会将需要的包下 载到`GOPATH`目
    录中

3.  go mod vendor

    与`go mod download`用法一致，不同的是这个命令会在当前项目中生成一个 vender 目录，然 后将引入的第
    三方

    包复制到此目录中

> 总结：使用第三方包时，推荐第三种方法，一来省去手动使用`go get`命令下载包的时间，二来依赖包都在项目
> 中比较清晰。还可以什么都不操作直接引入，运行时系统会自动将包下载到`GOPATH`目录中
