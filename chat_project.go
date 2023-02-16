package debox_chat_go_sdk

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"regexp"
	"strings"
	"time"
)

const (
	httpScheme  = "https://"
	httpsScheme = "https://"
	ipRegexStr  = `\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3}.*`
)

var (
	ipRegex = regexp.MustCompile(ipRegexStr)
)

// ChatProject defines chat project
type ChatProject struct {
	Name           string `json:"projectName"`    // Project name
	Description    string `json:"description"`    // Project description
	Status         string `json:"status"`         // Normal
	Owner          string `json:"owner"`          // empty
	Region         string `json:"region"`         // region
	CreateTime     string `json:"createTime"`     // unix time seconds, eg 1524539357
	LastModifyTime string `json:"lastModifyTime"` // unix time seconds, eg 1524539357

	Endpoint      string // IP or hostname of endpoint
	XApiKey       string
	SecurityToken string
	UsingHTTP     bool   // default https
	UserAgent     string // default defaultLogUserAgent
	AuthVersion   AuthVersionType
	baseURL       string
	retryTimeout  time.Duration
	httpClient    *http.Client
}

// NewLogProject creates a new SLS project.
func NewChatProject(name, endpoint, xApiKey string) (p *ChatProject, err error) {
	p = &ChatProject{
		Name:         name,
		Endpoint:     endpoint,
		XApiKey:      xApiKey,
		httpClient:   defaultHttpClient,
		retryTimeout: defaultRetryTimeout,
	}
	p.parseEndpoint()
	return p, nil
}

// WithToken add token parameter
func (p *ChatProject) WithToken(token string) (*ChatProject, error) {
	p.SecurityToken = token
	return p, nil
}

// WithRequestTimeout with custom timeout for a request
func (p *ChatProject) WithRequestTimeout(timeout time.Duration) *ChatProject {
	if p.httpClient == defaultHttpClient || p.httpClient == nil {
		p.httpClient = &http.Client{
			Timeout: timeout,
		}
	} else {
		p.httpClient.Timeout = timeout
	}
	return p
}

// WithRetryTimeout with custom timeout for a operation
// each operation may send one or more HTTP requests in case of retry required.
func (p *ChatProject) WithRetryTimeout(timeout time.Duration) *ChatProject {
	p.retryTimeout = timeout
	return p
}

// RawRequest send raw http request to LogService and return the raw http response
// @note you should call http.Response.Body.Close() to close body stream
func (p *ChatProject) RawRequest(method, uri string, headers map[string]string, body []byte) (*http.Response, error) {
	ctx := context.Background()
	return realRequest(ctx, p, method, uri, headers, body)
}

func (p *ChatProject) init() {
	if p.retryTimeout == time.Duration(0) {
		if p.httpClient == nil {
			p.httpClient = defaultHttpClient
		}
		p.retryTimeout = defaultRetryTimeout
		p.parseEndpoint()
	}
}

func (p *ChatProject) getBaseURL() string {
	if len(p.baseURL) > 0 {
		return p.baseURL
	}
	p.parseEndpoint()
	return p.baseURL
}

func (p *ChatProject) parseEndpoint() {
	scheme := httpScheme // default to http scheme
	host := p.Endpoint

	if strings.HasPrefix(p.Endpoint, httpScheme) {
		scheme = httpScheme
		host = strings.TrimPrefix(p.Endpoint, scheme)
	} else if strings.HasPrefix(p.Endpoint, httpsScheme) {
		scheme = httpsScheme
		host = strings.TrimPrefix(p.Endpoint, scheme)
	}

	if GlobalForceUsingHTTP || p.UsingHTTP {
		scheme = httpScheme
	}
	if ipRegex.MatchString(host) { // ip format
		// use direct ip proxy
		url, _ := url.Parse(fmt.Sprintf("%s%s", scheme, host))
		if p.httpClient == nil || p.httpClient == defaultHttpClient {
			p.httpClient = &http.Client{
				Transport: &http.Transport{
					Proxy: http.ProxyURL(url),
				},
				Timeout: defaultRequestTimeout,
			}
		} else {
			p.httpClient.Transport = &http.Transport{
				Proxy: http.ProxyURL(url),
			}
		}

	}
	if len(p.Name) == 0 {
		p.baseURL = fmt.Sprintf("%s%s", scheme, host)
	} else {
		p.baseURL = fmt.Sprintf("%s%s.%s", scheme, p.Name, host)
	}
}
