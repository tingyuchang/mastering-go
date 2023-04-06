# Web Services

- http.Response
- http.Request 
- http.Transport (low level, usually use http.Client)
- http.Client (high level)

although Go web server do many thing efficiently and securely, but Apache, Nginx, Caddy(written by Go) are better options.



## http handler

```go
// HandleFunc registers the handler function for the given pattern
// in the DefaultServeMux.
// The documentation for ServeMux explains how patterns are matched.
func HandleFunc(pattern string, handler func(ResponseWriter, *Request)) {
	DefaultServeMux.HandleFunc(pattern, handler)
}

```
如果使用 defaultServerMux 的話，就會用到 `http.HandleFunc("/", myHandler)` 這種方式來註冊 handler

不過如果是使用自定義的 ServerMux 則是會使用 `mux.Handle("/", http.HandlerFunc(defaultHandler))`

`http.HandlerFunc` 比較簡單，只是一個方便的轉化，讓一般的 func 轉型成 `http.Handler`

原始碼的部分也很好理解，就是替一般的 func 實現了 `ServeHTTP(ResponseWriter, *Request)` 這一個 interface func 

所以 `http.HandlerFunc` 滿足了 `http.Handler` 這一個 interface.

```go
// The HandlerFunc type is an adapter to allow the use of
// ordinary functions as HTTP handlers. If f is a function
// with the appropriate signature, HandlerFunc(f) is a
// Handler that calls f.
type HandlerFunc func(ResponseWriter, *Request)

// ServeHTTP calls f(w, r).
func (f HandlerFunc) ServeHTTP(w ResponseWriter, r *Request) {
	f(w, r)
}
```

因為名稱很類似，所以很容易搞混，不過 package http 的重點要回到 `http.Handler` 這一個 interface 上面

基本上一個 request 要怎麼處理會由 Server 呼叫 ServeHTTP 這一個 func 

```go
type Handler interface {
	ServeHTTP(ResponseWriter, *Request)
}
```

接下來來看看 ServerMux 的部分

1. Handler 
2. handler
3. ServeHTTP
4. Handle

從執行順序上應該是 4 -> 3 -> 1 -> 2

Handle 註冊了 path pattern string 與對應的 http.Handler

ServeHTTP 這裡有點意思，在 `func ListenAndServe(addr string, handler Handler) error ` 中，會把自定義的 ServerMux 放到 handler 中

因為 ServerMux 也實現了 ServeHTTP，接下來的處理 http request 的時候，server 會呼叫 ServerMux.ServeHTTP() 再由它轉發給其他註冊了 pattern string 的 handler.

接著 ServeHTTP 會呼叫 Handler 來取得對應的 http.Handler 並執行 ServeHTTP()

```go
func (mux *ServeMux) Handler(r *Request) (h Handler, pattern string)
// handler is the main implementation of Handler.
func (mux *ServeMux) handler(host, path string) (h Handler, pattern string)
// ServeHTTP dispatches the request to the handler whose
func (mux *ServeMux) ServeHTTP(w ResponseWriter, r *Request)
// Handle registers the handler for the given pattern.
func (mux *ServeMux) Handle(pattern string, handler Handler) 
```


## 

```go
func (srv *Server) Serve(l net.Listener) error {
	// ... skip
	for {
		// ... skip
        go c.serve(connCtx)
    }
}
```