# HonorableTodd_udemy_Web-Development-w-Google-s-Go-golang-Programming-Language_Practice
# Preferred To anybody :
***
***
- [x] **My Learning and Review :**

- This course emphasis myself, to better skill gained. Really That was an awesome experienced, never been    before. Discuss about every tiny intermediate fundamental things. I really love this course and preferred this course to everyone who want to learn or skilled Google's Golang.
- [x] **End of the course Encourage myself and sharing valuable resourse to take next steps**
- [x] **Special Thanks to Califonia State University Professor Tood Mcleod**:
- I'm really grateful to professor Todd Mcleod he gave me inspiration and free access oportunity this valuable course on Twitter.
- This course covered another most important topic How to Succeed in Life.
-                                                                   Regrads
                                                               Md. Nihad Hossain
***
***
# Course Outline:

- [x]  **This is Tood Mcleod course outline.**
- I'm Just share this outline my github ripo.
- [x] **Course Outline Links:**
- https://docs.google.com/document/d/1d6pLCMFyUfrxRlm7Wmh5LJ3yJEPBaygHHvROAwEA-10/edit?usp=sharing


# Completion and certified (Certificate):
- [x] **Persistently,patiently you are bound to succeed and growing mind set, my one of best takeway from this course.**

![](codesamples/40_certificate/UC-65ad2361-1240-480b-b225-be55d0b18daa.jpg)
***

***
# Growing Mindset and Learned  things from this course :
- [x] **Why golang choose for web debelopment**
- With every passing day, there is an increasing demand for speed, cost-effective and fewer resources for the website applications development. Along with it, everyone wants advanced usability as well as a smooth and continuous user experience. And to satisfy these requirements of web development advanced programming languages are needed like Golang.
- Key features {compiles into single binary,static typed,performance,most of the cases no need third praty library,IDE support and debugging,rapidly soppurt latest technologies like [ElasticSearch,TLS1.3,Couchbase,MondoDB etc]}
- Valuable and superb article for why choose web dev {https://go.dev/solutions/webdev/#key-benefits} {https://reemishirsath.medium.com/things-you-need-to-know-about-golang-for-web-development-9dd1ca1c883d}
- Take advantage of Multicore CPU's.
- Fastest growing and highest paying programming language.
- Invented by 3 geniuses [Rob Pike,Ken Thompson,Robert Geriesemer].
***

***
- [x] **How to Succeed in Life?**
- Grit and Perseverance -- Angela Lee Duckworth
- Growing mind set and Fixed mindset Theory.
- Focus and commitment.
*** 

***
- [x] **Most liked and followed quotes of my teachers**
- Begin with the End In Mind {set goals}
- Put first think first {time on task}
- Persistently,patiently you are bound to succeed.
- Drop by drop  the bucket gets filled.
- Complex is easy when presented by master.
- Take time to understand the wisdom of the master.
***

***
- [x] **Go Forums**
- https://forum.golangbridge.org/
- [x] **Official Documentation**
- golang spec
- effective go
***
***
- [x] **golang.org vs godoc.org**

- golang.org [{About language},{Standard library}]
- godoc.org [{Standard library and Third party packages}]  -[https://pkg.go.dev/]

***
***
- [x] **Concatenation :**
- HTML to TEXT file
***
***
- [x] **Parsing data with using text/tempalte :**

```go
- func Must(t *Template, err error) *Template
- func New(name string) *Template
- func ParseFS(fsys fs.FS, patterns ...string) (*Template, error)
- func ParseFiles(filenames ...string) (*Template, error)
- func ParseGlob(pattern string) (*Template, error)
- func (t *Template) Parse(text string) (*Template, error)
- func (t *Template) Execute(wr io.Writer, data interface{}) error
- func (t *Template) ExecuteTemplate(wr io.Writer, name string, data interface{}) error
- func New(name string) *Template
- func init()
- variable in golang template
- {{$name :=.}}
- function in golang template
- func (t *Template) Funcs(funcMap FuncMap) *Template
```
***
***
- [x] **Function in templates**
- [x] **Pipeline in templates**
- [x] **Predefined Global function in templatess**
- [x] **Predefined comperison operators**
- [x] **Comment in template**
- [x] **Define a templates**
- [x] **Use/invoke templates**
- [x] **Method Use/invoke in templates**
- [x] **Cross Site Scripting and html/template packagee**
***
***
- [x] **route/request router/multiplexer/mux/servemux/server/http router/http request router/http multiplexer/http mux/http servemux/http server**
- [x] **HTTP Request/Response**
- [x] **HTTP/2**
- [x] **RFCS[Request For Comments]**
- [x] **IETF[Internet Engineering Task Force]**
- [x] **W3C[World Wide Web Consortium]**
- [x] **[RFC 7230](https://datatracker.ietf.org/doc/html/rfc7230#section-3.1.1)**
- [x] **[HTTP Messages](https://developer.mozilla.org/en-US/docs/Web/HTTP/Messages)**
- [x] **[HTTP response status codes ](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status)**
- [x] **[HTTP request methods](https://developer.mozilla.org/en-US/docs/Web/HTTP/Methods)**
- [x] **[HTTP headers](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers)**
- [x] **[bufio](https://pkg.go.dev/bufio#pkg-index)**
***

- [x] **net Package**
```go
HTTP Server
 HTTP uses TCP.
 To create a server that works with HTTP, we just create a TCP server.
 To configure our code to handle request/response in an HTTP fashion which works with browsers, we - - need to adhere to HTTP standards.
 ```

 - [x] **TCP server essentials**
 - [x] **Listen**
 ```go
net.Listen

 func Listen(net, laddr string) (Listener, error)
 **Listener**
 net.Listener

 type Listener interface {
    // Accept waits for and returns the next connection to the listener.
    Accept() (Conn, error)

    // Close closes the listener.
    // Any blocked Accept operations will be unblocked and return errors.
    Close() error

    // Addr returns the listener's network address.
    Addr() Addr
}
```
- [x] **Connection**
```go
net.Conn

type Conn interface {
    // Read reads data from the connection.
    Read(b []byte) (n int, err error)

    // Write writes data to the connection.
    Write(b []byte) (n int, err error)

    // Close closes the connection.
    // Any blocked Read or Write operations will be unblocked and return errors.
    Close() error

    // LocalAddr returns the local network address.
    LocalAddr() Addr

    // RemoteAddr returns the remote network address.
    RemoteAddr() Addr

    SetDeadline(t time.Time) error

    SetReadDeadline(t time.Time) error

    SetWriteDeadline(t time.Time) error
}
```
- [x] **Dial**
```go
net.Dial

func Dial(network, address string) (Conn, error)
Write
io.WriteString

func WriteString(w Writer, s string) (n int, err error)
fmt.Fprintln

func Fprintln(w io.Writer, a ...interface{}) (n int, err error)
Read
ioutil.ReadAll
func ReadAll(r io.Reader) ([]byte, error)
bufio.NewScanner
func NewScanner(r io.Reader) *Scanner
bufio.Scan
func (s *Scanner) Scan() bool
bufio.Text
func (s *Scanner) Text() string
Read & Write
io.Copy
func Copy(dst Writer, src Reader) (written int64, err error)
```
- [x] **Multiplexer handle**
***

***
- [x] **net/http Package**
- [x] **Handler**
```go
type Handler interface {
    - ServeHTTP(ResponseWriter, *Request)
}
```
- [x] **Server**
```go
 http.ListenAndServe
```
- [x] **Request**
```go
 Method string
 URL *url.URL
	//	Header = map[string][]string{
	//		"Accept-Encoding": {"gzip, deflate"},
	//		"Accept-Language": {"en-us"},
	//		"Foo": {"Bar", "two"},
	//	}
    Header Header
    Body io.ReadCloser
    ContentLength int64
    Host string
    // This field is only available after ParseForm is called.
    Form url.Values
    // This field is only available after ParseForm is called.
    PostForm url.Values
    MultipartForm *multipart.Form
    // RemoteAddr allows HTTP servers and other software to record
	// the network address that sent the request, usually for
	// logging. 
    RemoteAddr string
}
```
- [x] **Retrieve URL & Form data**
```go
http.Request is a struct. It has the fields Form & PostForm. If we read the documentation on these, we see:

    // Form contains the parsed form data, including both the URL
    // field's query parameters and the POST or PUT form data.
    // This field is only available after **ParseForm** is called.
    // The HTTP client ignores Form and uses Body instead.
    Form url.Values

    // PostForm contains the parsed form data from POST, PATCH,
    // or PUT body parameters.
    // This field is only available after **ParseForm** is called.
    // The HTTP client ignores PostForm and uses Body instead.
    PostForm url.Values

 If we look at ParseForm,

 go func (r *Request) ParseForm() error

 we see that this is a method attached to a *http.Request.

 If we look at FormValue*

go func (r *Request) FormValue(key string) string

 we see that this is a method attached to a *http.Request. FormValue returns the first value for the named component of the query. POST and PUT body parameters take precedence over URL query string values. FormValue calls ParseMultipartForm and ParseForm if necessary and ignores any errors returned by these functions. If key is not present, FormValue returns the empty string. To access multiple values of the same key, call ParseForm and then inspect Request.Form directly.

See the HTTP Method
 The http.Request type is a struct which has a Method field.

 See URL values
The http.Request type is a struct which has a URL field. Notice that the type is a *url.URL

Take a look at type url.URL

 type URL struct {
    Scheme     string
    Opaque     string    // encoded opaque data
    User       *Userinfo // username and password information
    Host       string    // host or host:port
    Path       string
    RawPath    string // encoded path hint (Go 1.5 and later only; see EscapedPath method)
    ForceQuery bool   // append a query ('?') even if RawQuery is empty
    RawQuery   string // encoded query values, without '?'
    Fragment   string // fragment for references, without '#'
}
Work with the HTTP header
The http.Request type is a struct which has a Header field.

From the documentation about the http.Header struct, we see that:

 **type Header map[string][]string**
```
- [x] **Response**
```go
http.ResponseWriter

type ResponseWriter interface {
    // Header returns the header map that will be sent by
    // WriteHeader. Changing the header after a call to
    // WriteHeader (or Write) has no effect 
    Header() Header

    // Write writes the data to the connection as part of an HTTP reply.
    //
    // If WriteHeader has not yet been called, Write calls
    // WriteHeader(http.StatusOK) before writing the data. If the Header
    // does not contain a Content-Type line, Write adds a Content-Type set
    // to the result of passing the initial 512 bytes of written data to
    // DetectContentType.
    Write([]byte) (int, error)

    // WriteHeader sends an HTTP response header with status code.
    // If WriteHeader is not called explicitly, the first call to Write
    // will trigger an implicit WriteHeader(http.StatusOK).
    // Thus explicit calls to WriteHeader are mainly used to
    // send error codes.
    WriteHeader(int)
}
Setting a response header
An http.ResponseWriter has a method Header() which returns a http.Header.

Look at the documentation for http.Header

type Header map[string][]string**
Look at the methods which are attached to type http.Header

type Header
func (h Header) Add(key, value string)
func (h Header) Del(key string)
func (h Header) Get(key string) string
func (h Header) Set(key, value string)
func (h Header) Write(w io.Writer) error
func (h Header) WriteSubset(w io.Writer, exclude map[string]bool) error
We can set headers for a response like this:

res.Header().Set("Content-Type", "text/html; charset=utf-8")
```
***

***
- [x] **ServeMux**
```go
http.ServeMux
type ServeMux
	func NewServeMux() *ServeMux
	func (mux *ServeMux) Handle(pattern string, handler Handler)
	func (mux *ServeMux) HandleFunc(pattern string, handler func(ResponseWriter, *Request))
	func (mux *ServeMux) Handler(r *Request) (h Handler, pattern string)
	func (mux *ServeMux) ServeHTTP(w ResponseWriter, r *Request)

We can use Handle like this:
	mux := http.NewServeMux()
	mux.Handle("/", h)
	mux.Handle("/cat/", c)

DefaultServeMux
ListenAndServe starts an HTTP server with a given address and handler. The handler is usually nil, which means to use DefaultServeMux. Handle and HandleFunc add handlers to DefaultServeMux:

http.ListenAndServe(":8080", nil)

HandleFunc
http.HandleFunc
func HandleFunc(pattern string, handler func(ResponseWriter, *Request))

HandlerFunc
http.HandlerFunc
type HandlerFunc func(ResponseWriter, *Request)
func (f HandlerFunc) ServeHTTP(w ResponseWriter, r *Request)
```
***
