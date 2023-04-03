package simple_http

import "net/http"

type Context struct {
	response http.ResponseWriter
	request  *http.Request

	Method string
}

// 创建新的Context
func NewContext(w http.ResponseWriter, r *http.Request) *Context {
	return &Context{
		response: w,
		request:  r,

		Method: r.Method,
	}
}

func (c *Context) Cookie(name string) (*http.Cookie, error) {
	return c.request.Cookie(name)
}

func (c *Context) Cookies() []*http.Cookie {
	return c.request.Cookies()
}

func (c *Context) SetCookie(cookie *http.Cookie) {
	http.SetCookie(c.response, cookie)
}

// 设置响应状态
func (c *Context) SetStatus(code int) {
	c.response.WriteHeader(code)
}

// 设置响应内容
func (c *Context) SetBody(content string) {
	c.SetBodyForByte([]byte(content))
}

// 设置二进制响应内容
func (c *Context) SetBodyForByte(content []byte) {
	c.response.Write(content)
}

// 获取查询字符串
func (c *Context) Query(key string) string {
	return c.getQuery(key)
}

// 获取查询字符串，如果没找到就返回默认值
func (c *Context) DefaultQuery(key string, defaultVal string) string {
	if val := c.getQuery(key); val != "" {
		return val
	}

	return defaultVal
}

func (c *Context) getQuery(key string) string {
	return c.request.URL.Query().Get(key)
}
