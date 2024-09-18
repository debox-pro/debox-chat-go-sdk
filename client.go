package debox_chat_go_sdk

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"sync"
	"time"
)

// GlobalForceUsingHTTP if GlobalForceUsingHTTP is true, then all request will use HTTP(ignore LogProject's UsingHTTP flag)
var GlobalForceUsingHTTP = false

// RetryOnServerErrorEnabled if RetryOnServerErrorEnabled is false, then all error requests will not be retried
var RetryOnServerErrorEnabled = true

var GlobalDebugLevel = 0

var MaxCompletedRetryCount = 20

var MaxCompletedRetryLatency = 5 * time.Minute

var InvalidCompressError = errors.New("Invalid Compress Type")

const DefaultLogUserAgent = "golang-sdk-v0.1.0"

// AuthVersionType the version of auth
type AuthVersionType string

const (
	// AuthV1 v1
	AuthV1 AuthVersionType = "v1"
)

// Error defines sls error
type Error struct {
	HTTPCode  int32  `json:"httpCode"`
	Code      string `json:"errorCode"`
	Message   string `json:"errorMessage"`
	RequestID string `json:"requestID"`
}

func IsDebugLevelMatched(level int) bool {
	return level <= GlobalDebugLevel
}

// NewClientError new client error
func NewClientError(err error) *Error {
	if err == nil {
		return nil
	}
	if clientError, ok := err.(*Error); ok {
		return clientError
	}
	clientError := new(Error)
	clientError.HTTPCode = -1
	clientError.Code = "ClientError"
	clientError.Message = err.Error()
	return clientError
}

func (e Error) String() string {
	b, err := json.MarshalIndent(e, "", "    ")
	if err != nil {
		return ""
	}
	return string(b)
}

func (e Error) Error() string {
	return e.String()
}

func IsTokenError(err error) bool {
	if clientErr, ok := err.(*Error); ok {
		if clientErr.HTTPCode == 401 {
			return true
		}
	}
	return false
}

// Client ...
type Client struct {
	Endpoint       string // IP or hostname of SLS endpoint
	XApiKey        string
	UserAgent      string // default defaultLogUserAgent
	RequestTimeOut time.Duration
	RetryTimeOut   time.Duration
	HTTPClient     *http.Client
	//Region         string
	AuthVersion   AuthVersionType
	accessKeyLock sync.RWMutex
}

// SetUserAgent set a custom userAgent
func (c *Client) SetUserAgent(userAgent string) {
	c.UserAgent = userAgent
}

// SetHTTPClient set a custom http client, all request will send to sls by this client
func (c *Client) SetHTTPClient(client *http.Client) {
	c.HTTPClient = client
}

// SetAuthVersion set signature version that the client used
func (c *Client) SetAuthVersion(version AuthVersionType) {
	c.accessKeyLock.Lock()
	c.AuthVersion = version
	c.accessKeyLock.Unlock()
}

// ResetAccessKeyToken reset client's access key token
func (c *Client) ResetAccessKeyToken(xApiKey string) {
	c.accessKeyLock.Lock()
	c.XApiKey = xApiKey
	c.accessKeyLock.Unlock()
}

// SendChatMsg send recall message.
func (c *Client) SendChatMsg(toUserId, groupId, message, opreate string) (*ChatProject, error) {
	type Body struct {
		ToUserId string `json:"to_user_id"`
		GroupId  string `json:"group_id"`
		Message  string `json:"message"`
	}
	body, err := json.Marshal(Body{
		ToUserId: toUserId,
		GroupId:  groupId,
		Message:  message,
	})
	if err != nil {
		return nil, err
	}

	h := map[string]string{
		"x-chat-bodyrawsize": fmt.Sprintf("%d", len(body)),
		"Content-Type":       "application/json",
		"Accept-Encoding":    "deflate",
		"X-API-KEY":          c.XApiKey,
	}

	uri := "/openapi/send_chat_message"
	proj := convert(c, opreate)
	resp, err := c.request("", "POST", uri, h, body)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	return proj, nil
}

// RegisterCallbakUrl create a new event.
func (c *Client) RegisterCallbakUrl(url, method, operate string) (*ChatProject, error) {
	type Body struct {
		Url        string `json:"url"`
		HttpMethod string `json:"http_method"`
	}
	body, err := json.Marshal(Body{
		Url:        url,
		HttpMethod: method,
	})
	if err != nil {
		return nil, err
	}

	h := map[string]string{
		"x-chat-bodyrawsize": fmt.Sprintf("%d", len(body)),
		"Content-Type":       "application/json",
		"Accept-Encoding":    "deflate",
		"X-API-KEY":          c.XApiKey,
	}

	uri := "/openapi/register_callbak_url"
	proj := convert(c, operate)
	resp, err := c.request("", "POST", uri, h, body)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	return proj, nil
}

// SendRobotMsg send recall message.
func (c *Client) SendRobotMsg(toUserId, message, objectName, opreate string) (*ChatProject, error) {
	type Body struct {
		ToUserId   string `json:"to_user_id"`
		FromUserId string `json:"from_user_id"`
		ObjectName string `json:"object_name"`
		Message    string `json:"message"`
	}
	body, err := json.Marshal(Body{
		ToUserId:   toUserId,
		ObjectName: objectName,
		Message:    message,
	})
	if err != nil {
		return nil, err
	}

	h := map[string]string{
		"x-chat-bodyrawsize": fmt.Sprintf("%d", len(body)),
		"Content-Type":       "application/json",
		"Accept-Encoding":    "deflate",
		"X-API-KEY":          c.XApiKey,
	}

	uri := "/openapi/send_robot_message"
	proj := convert(c, opreate)
	resp, err := c.request("", "POST", uri, h, body)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	return proj, nil
}

// SendRobotMsg send recall message.
func (c *Client) SendRobotGroupMsg(toUserId, groupId, title, content, message, objectName, operate, href string) (*ChatProject, error) {
	type Body struct {
		ToUserId   string `json:"to_user_id"`
		GroupId    string `json:"group_id"`
		Title      string `json:"title"`
		Content    string `json:"content"`
		Message    string `json:"message"`
		ObjectName string `json:"object_name"`
		Href       string `json:"href"`
	}
	body, err := json.Marshal(Body{
		ToUserId:   toUserId,
		GroupId:    groupId,
		Title:      title,
		Content:    content,
		ObjectName: objectName,
		Message:    message,
		Href:       href,
	})
	if err != nil {
		return nil, err
	}

	h := map[string]string{
		"x-chat-bodyrawsize": fmt.Sprintf("%d", len(body)),
		"Content-Type":       "application/json",
		"Accept-Encoding":    "deflate",
		"X-API-KEY":          c.XApiKey,
	}

	uri := "/openapi/send_robot_group_message"
	proj := convert(c, operate)
	resp, err := c.request("", "POST", uri, h, body)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	return proj, nil
}

func convert(c *Client, projName string) *ChatProject {
	c.accessKeyLock.RLock()
	defer c.accessKeyLock.RUnlock()
	return convertLocked(c, projName)
}

func convertLocked(c *Client, projName string) *ChatProject {
	p, _ := NewChatProject(projName, c.Endpoint, c.XApiKey)
	p.UserAgent = c.UserAgent
	p.AuthVersion = c.AuthVersion
	if c.HTTPClient != nil {
		p.httpClient = c.HTTPClient
	}
	if c.RequestTimeOut != time.Duration(0) {
		p.WithRequestTimeout(c.RequestTimeOut)
	}
	if c.RetryTimeOut != time.Duration(0) {
		p.WithRetryTimeout(c.RetryTimeOut)
	}

	return p
}

// Close the client
func (c *Client) Close() error {
	return nil
}
