# gzip过滤器源码分析



## 压缩的必要性

在通信过程中压缩可以减少传输所用的码元数目，但是又不会丢失掉其中的信息，对于有限的传输带宽来说是由必要的



1. const

```go

const (
	encodingGzip = "gzip"
 
	headerAcceptEncoding  = "Accept-Encoding"
	headerContentEncoding = "Content-Encoding"
	headerContentLength   = "Content-Length"
	headerContentType     = "Content-Type"
	headerVary            = "Vary"
	headerSecWebSocketKey = "Sec-WebSocket-Key"
 
	BestCompression    = gzip.BestCompression
	BestSpeed          = gzip.BestSpeed
	DefaultCompression = gzip.DefaultCompression
	NoCompression      = gzip.NoCompression
)

```

NoCompression = 0

BestSpeed = 1

BestCompression = 9

DefaultCompression = -1

代表压缩的level，不能超过BestCompression





2. 

   ```go
   type gzipResponseWriter struct {
   	w *gzip.Writer
   	negroni.ResponseWriter
   	wroteHeader bool
   }
   ```
   wroteHeader代表response是否已经编码


3. 
```go
func (grw *gzipResponseWriter) Write(b []byte) (int, error) {
	if !grw.wroteHeader {
		grw.WriteHeader(http.StatusOK)
	}
	if grw.w == nil {
		return grw.ResponseWriter.Write(b)
	}
	if len(grw.Header().Get(headerContentType)) == 0 {
		grw.Header().Set(headerContentType, http.DetectContentType(b))
	}
	return grw.w.Write(b)
}
```
 写内容的函数，1.报头未写 -> WriteHeader() 2.gzipWriter没有，说明不gzip压缩 -> ResponseWriter写，返回 3.报头未设置 -> 通过net/http库函数自动检测内容类型设置 4.gzipWriter写


4. 
```go

func (grw *gzipResponseWriter) WriteHeader(code int) {
	headers := grw.ResponseWriter.Header()
	if headers.Get(headerContentEncoding) == "" {
		headers.Set(headerContentEncoding, encodingGzip)
		headers.Add(headerVary, headerAcceptEncoding)
	} else {
		grw.w.Reset(ioutil.Discard)
		grw.w = nil
	}
	grw.ResponseWriter.WriteHeader(code)
	grw.wroteHeader = true
}

```
如果目标页面的响应内容未预编码，采用gzip压缩方式压缩后再发送到客户端，同时设置Content-Encoding实体报头值为gzip，否则在写之前令gzipWriter失效（var Discard io.Writer = devNull(0)，使得它对任何写调用无条件成功)



5. 
```go

func (h *handler) ServeHTTP(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	// Skip compression if the client doesn't accept gzip encoding.
	if !strings.Contains(r.Header.Get(headerAcceptEncoding), encodingGzip) {
		next(w, r)
		return
	}
 
	// Skip compression if client attempt WebSocket connection
	if len(r.Header.Get(headerSecWebSocketKey)) > 0 {
		next(w, r)
		return
	}
 
	// Retrieve gzip writer from the pool. Reset it to use the ResponseWriter.
	// This allows us to re-use an already allocated buffer rather than
	// allocating a new buffer for every request.
	// We defer g.pool.Put here so that the gz writer is returned to the
	// pool if any thing after here fails for some reason (functions in
	// next could potentially panic, etc)
	gz := h.pool.Get().(*gzip.Writer)
	defer h.pool.Put(gz)
	gz.Reset(w)
 
	// Wrap the original http.ResponseWriter with negroni.ResponseWriter
	// and create the gzipResponseWriter.
	nrw := negroni.NewResponseWriter(w)
	grw := gzipResponseWriter{gz, nrw, false}
 
	// Call the next handler supplying the gzipResponseWriter instead of
	// the original.
	next(&grw, r)
 
	// Delete the content length after we know we have been written to.
	grw.Header().Del(headerContentLength)
 
	gz.Close()


```

处理handler中压缩请求的函数，1.浏览器不接受gzip编码 -> 跳过 2.浏览器尝试进行socket连接 -> 跳过

3.复用gzipWriter,创建gzipResponseWriter，传给下一个handler,删除报头“Content-Length”字段，关掉 gzipwriter

 

