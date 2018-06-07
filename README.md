# snowflakeIDServe

A simple go http server to generate a snowflakeID

# use

### 1. install golang, set $GOPATH, $GOBIN

refer to [Golang Doc - How to Write Go Code](https://golang.org/doc/code.html)

### 2. modify $GOPATH/src/server/http.go :

```golang
// change your port
err = http.ListenAndServe(":port", nil)
```
### 3. build

cd $GOPATH/src/server/

go install

### 4. run server

$GOPATH/bin/server
