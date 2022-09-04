package gee

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type H map[string]interface{}

type Context struct {
	//	origin object
	Req    *http.Request
	Writer http.ResponseWriter

	//	request info
	Path   string
	Method string

	//	response info
	StatusCode int
}

func newContext(req *http.Request, w http.ResponseWriter) *Context {
	return &Context{
		Req:    req,
		Writer: w,
		Path:   req.URL.Path,
		Method: req.Method,
	}
}

// PostForm get parameter from form
func (c *Context) PostForm(key string) string {
	return c.Req.FormValue(key)
}

// Query query parameter from url
func (c *Context) Query(key string) string {
	return c.Req.URL.Query().Get(key)
}

// Status set response code
func (c *Context) Status(code int) {
	c.StatusCode = code
	c.Writer.WriteHeader(code)
}

// SetHeader set response header
func (c *Context) SetHeader(key, value string) {
	c.Writer.Header().Set(key, value)
}

//response return string
func (c *Context) String(code int, format string, values ...interface{}) {
	c.SetHeader("Content-Type", "text/plain")
	c.Status(code)
	c.Writer.Write([]byte(fmt.Sprintf(format, values...)))
}

// JSON response return json
func (c *Context) JSON(code int, obj interface{}) {
	c.SetHeader("Content-Type", "application/json")
	c.Status(code)
	encoder := json.NewEncoder(c.Writer)
	if err := encoder.Encode(obj); err != nil {
		panic(err)
	}
}

// Data response add data
func (c *Context) Data(code int, data []byte) {
	c.Status(code)
	c.Writer.Write(data)
}

// HTML response return html
func (c *Context) HTML(code int, html string) {
	c.Status(code)
	c.Writer.Write([]byte(html))
}
