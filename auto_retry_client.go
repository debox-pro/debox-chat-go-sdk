package debox_chat_go_sdk

import (
	"net/http"
	"sync"
	"time"
)

type AutoRetryClient struct {
	chatClient             ClientInterface
	shutdown               <-chan struct{}
	closeFlag              bool
	tokenUpdateFunc        UpdateTokenFunction
	maxTryTimes            int
	waitIntervalMin        time.Duration
	waitIntervalMax        time.Duration
	updateTokenIntervalMin time.Duration
	nextExpire             time.Time

	lock               sync.Mutex
	lastFetch          time.Time
	lastRetryFailCount int
	lastRetryInterval  time.Duration
}

func (c *AutoRetryClient) SetUserAgent(userAgent string) {
	c.chatClient.SetUserAgent(userAgent)
}

// SetHTTPClient set a custom http client, all request will send to sls by this client
func (c *AutoRetryClient) SetHTTPClient(client *http.Client) {
	c.chatClient.SetHTTPClient(client)
}

// SetAuthVersion set auth version that the client used
func (c *AutoRetryClient) SetAuthVersion(version AuthVersionType) {
	c.chatClient.SetAuthVersion(version)
}

func (c *AutoRetryClient) Close() error {
	c.closeFlag = true
	return nil
}

func (c *AutoRetryClient) ResetAccessKeyToken(xApiKey string) {
	c.chatClient.ResetAccessKeyToken(xApiKey)
}

func (c *AutoRetryClient) SendChatMsg(toUserId, groupId, message, operate string) (prj *ChatProject, err error) {
	for i := 0; i < c.maxTryTimes; i++ {
		prj, err = c.chatClient.SendChatMsg(toUserId, groupId, message, operate)
		if err != nil {
			return
		}
	}
	return
}

func (c *AutoRetryClient) RegisterCallbakUrl(url, method, operate string) (prj *ChatProject, err error) {
	for i := 0; i < c.maxTryTimes; i++ {
		prj, err = c.chatClient.RegisterCallbakUrl(url, method, operate)
		if err != nil {
			return
		}
	}
	return
}

func (c *AutoRetryClient) SendRobotMsg(toUserId, message, objectName, operate string) (prj *ChatProject, err error) {
	for i := 0; i < c.maxTryTimes; i++ {
		prj, err = c.chatClient.SendRobotMsg(toUserId, message, objectName, operate)
		if err != nil {
			return
		}
	}
	return
}

func (c *AutoRetryClient) SendRobotGroupMsg(toUserId, groupId, title, content, message, objectName, operate, href string) (prj *ChatProject, err error) {
	for i := 0; i < c.maxTryTimes; i++ {
		prj, err = c.chatClient.SendRobotGroupMsg(toUserId, groupId, title, content, message, objectName, operate, href)
		if err != nil {
			return
		}
	}
	return
}
