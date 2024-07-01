package debox_chat_go_sdk

import (
	"net/http"
	"time"
)

// CreateNormalInterface create a normal client
func CreateNormalInterface(endpoint, xApiKey string) ClientInterface {
	return &Client{
		Endpoint: endpoint,
		XApiKey:  xApiKey,
	}
}

type UpdateTokenFunction func() (xApiKey string, expireTime time.Time, err error)

// CreateTokenAutoUpdateClient crate a AutoRetryClient
// this client will auto fetch security token and retry when operation is `Unauthorized`
// @note AutoRetryClient will destroy when shutdown channel is closed
//func CreateTokenAutoUpdateClient(endpoint string, tokenUpdateFunc UpdateTokenFunction, shutdown <-chan struct{}) (client ClientInterface, err error) {
//	xApiKey, expireTime, err := tokenUpdateFunc()
//	if err != nil {
//		return nil, err
//	}
//	tauc := &AutoRetryClient{
//		chatClient:             CreateNormalInterface(endpoint, xApiKey),
//		shutdown:               shutdown,
//		tokenUpdateFunc:        tokenUpdateFunc,
//		maxTryTimes:            3,
//		waitIntervalMin:        time.Duration(time.Second * 1),
//		waitIntervalMax:        time.Duration(time.Second * 60),
//		updateTokenIntervalMin: time.Duration(time.Second * 1),
//		nextExpire:             expireTime,
//	}
//	go tauc.flushSTSToken()
//	return tauc, nil
//}

// ClientInterface for all chat's open api
type ClientInterface interface {
	SetUserAgent(userAgent string)
	SetHTTPClient(client *http.Client)
	// #################### Client Operations #####################
	// ResetAccessKeyToken reset client's access key token
	ResetAccessKeyToken(xApiKey string)
	// SetAuthVersion Set signature version
	SetAuthVersion(version AuthVersionType)
	// Close the client
	Close() error

	// #################### Chat Operations #####################
	SendChatMsg(toUserId, groupId, message, operate string) (*ChatProject, error) //old
	RegisterCallbakUrl(url, method, operate string) (*ChatProject, error)
	SendRobotMsg(toUserId, message, objectName, operate string) (*ChatProject, error)
	SendRobotGroupMsg(toUserId, groupId, title, content, message, objectName, operate, href string) (*ChatProject, error)
}
