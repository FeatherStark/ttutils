# TITAN-Utils

个人用于快速开发的依赖库，封装了常用的函数，避免多项目代码重写。

旨在让 Golang 变得像 Python 一样简单易用。

## 方法说明（部分展示）

具体使用看方法注释即可。

### 快速使用

`go get -u github.com/FeatherStark/ttutils`

### cmd 执行系统命令

---

`func CommandExecute(command string) CommandExecResult{...}`

Description: CommandExecute 执行系统命令并打印结果

Args: command string 命令字符串

Returns: CommandExecResult 命令执行结果

---

```go
execResult := ttutils.CommandExecute("ipconfig")
fmt.Println(execResult.Success)
fmt.Println(execResult.CommandString)
fmt.Println(execResult.Output)
fmt.Println(execResult.Error)
```


### jwt 解析JWT参数

---

`func JwtParseClaims(jwtTokenString, jwtSecret string) (jwt.MapClaims, error){...}`

Description: JwtParseClaims 传入JWT字符串，解析并返回其中的claims（明文字段内容），如claims['username']获取username的值

Args: jwtTokenString JWT字符串, jwtSecret JWT密钥

Returns: jwt.MapClaims, error

---

```go
jwtToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwidXNlcm5hbWUiOiJKb2huIERvZSIsImlhdCI6MTUxNjIzOTAyMn0.E1iIYCH3fs06Z9aEmDNaHOx9G-zeiqO2xtugPqaQHrQ"
jwtSecret := "mysecret"
claims, err := ttutils.JwtParseClaims(jwtToken, jwtSecret)
if err != nil {
    fmt.Println(err.Error())
    return
}
if claims["username"] == "John Doe" {
    fmt.Println("username is John Doe")
}
```

### request HTTP请求的配置和发送

---

`func RequestGetConfig(uri string) *HttpRequestConfig{...}`

Description: RequestGetConfig GET 请求。

Args: uri string 请求地址。

Returns: *HttpRequestConfig HTTP 请求配置。

---

`func RequestPostConfig(uri string) *HttpRequestConfig{...}`

Description: RequestPostConfig POST 请求  默认的 Content-Type 是 application/x-www-form-urlencoded。

Args: uri string 请求地址, data string 请求数据。

Returns: *HttpRequestConfig HTTP 请求配置。

---

```go
cfg := ttutils.RequestPostConfig("/post")
cfg.Header.Store("Cookie", "user=admin;")
cfg.Header.Store("Content-Type", "application/x-www-form-urlencoded")
cfg.FollowRedirect = false
cfg.VerifyTls = false
cfg.Data = `hello`
resp, err := ttutils.DoHttpRequest("https://httpbin.org/", cfg)
if err != nil {
    fmt.Println(err)
    return
}
fmt.Println(resp.Utf8Html)
```