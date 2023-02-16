package debox_chat_go_sdk

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"net/url"
	"sort"
	"strings"
	"time"
)

const (
	HTTPHeaderAuthorization   = "Authorization"
	HTTPHeaderContentMD5      = "Content-MD5"
	HTTPHeaderContentType     = "Content-Type"
	HTTPHeaderContentLength   = "Content-Length"
	HTTPHeaderDate            = "Date"
	HTTPHeaderHost            = "Host"
	HTTPHeaderUserAgent       = "User-Agent"
	HTTPHeaderAPIVersion      = "x-chat-apiversion"
	HTTPHeaderSignatureMethod = "x-chat-signaturemethod"
	HTTPHeaderBodyRawSize     = "x-chat-bodyrawsize"
	ISO8601                   = "20060102T150405Z"
)

type Signer interface {
	// Sign modifies @param headers only, adds signature and other http headers
	// that log services authorization requires.
	Sign(method, uriWithQuery string, headers map[string]string, body []byte) error
}

// GMT location
var gmtLoc = time.FixedZone("GMT", 0)

// NowRFC1123 returns now time in RFC1123 format with GMT timezone,
// eg, "Mon, 02 Jan 2006 15:04:05 GMT".
func nowRFC1123() string {
	return time.Now().In(gmtLoc).Format(time.RFC1123)
}

// SignerV1 version v1
type SignerV1 struct {
	xApiKey string
}

func NewSignerV1(xApiKey string) *SignerV1 {
	return &SignerV1{
		xApiKey: xApiKey,
	}
}

func (s *SignerV1) Sign(method, uri string, headers map[string]string, body []byte) error {
	var contentMD5, contentType, date, canoHeaders, canoResource string
	if body != nil {
		contentMD5 = fmt.Sprintf("%X", md5.Sum(body))
		headers[HTTPHeaderContentMD5] = contentMD5
	}

	if val, ok := headers[HTTPHeaderContentType]; ok {
		contentType = val
	}

	date, ok := headers[HTTPHeaderDate]
	if !ok {
		return fmt.Errorf("Can't find 'Date' header")
	}
	headers[HTTPHeaderSignatureMethod] = signatureMethod
	var slsHeaderKeys sort.StringSlice

	// Calc CanonicalizedSLSHeaders
	slsHeaders := make(map[string]string, len(headers))
	for k, v := range headers {
		l := strings.TrimSpace(strings.ToLower(k))
		if strings.HasPrefix(l, "x-log-") || strings.HasPrefix(l, "x-acs-") {
			slsHeaders[l] = strings.TrimSpace(v)
			slsHeaderKeys = append(slsHeaderKeys, l)
		}
	}

	sort.Sort(slsHeaderKeys)
	for i, k := range slsHeaderKeys {
		canoHeaders += k + ":" + slsHeaders[k]
		if i+1 < len(slsHeaderKeys) {
			canoHeaders += "\n"
		}
	}

	// Calc CanonicalizedResource
	u, err := url.Parse(uri)
	if err != nil {
		return err
	}

	canoResource += u.EscapedPath()
	if u.RawQuery != "" {
		var keys sort.StringSlice

		vals := u.Query()
		for k := range vals {
			keys = append(keys, k)
		}

		sort.Sort(keys)
		canoResource += "?"
		for i, k := range keys {
			if i > 0 {
				canoResource += "&"
			}

			for _, v := range vals[k] {
				canoResource += k + "=" + v
			}
		}
	}

	signStr := method + "\n" +
		contentMD5 + "\n" +
		contentType + "\n" +
		date + "\n" +
		canoHeaders + "\n" +
		canoResource

	// Signature = base64(hmac-sha1(UTF8-Encoding-Of(SignString)，AccessKeySecret))
	mac := hmac.New(sha1.New, []byte(s.xApiKey))
	_, err = mac.Write([]byte(signStr))
	if err != nil {
		return err
	}
	digest := base64.StdEncoding.EncodeToString(mac.Sum(nil))
	auth := fmt.Sprintf("SLS %s:%s", s.xApiKey, digest)
	headers[HTTPHeaderAuthorization] = auth
	return nil
}

func dateTimeISO8601() string {
	return time.Now().In(gmtLoc).Format(ISO8601)
}

// SignerV4 sign version v4, a non-empty region is required
type SignerV4 struct {
	xApiKey string
	region  string
}

func NewSignerV4(xApiKey, region string) *SignerV4 {
	return &SignerV4{
		xApiKey: xApiKey,
		region:  region,
	}
}

func (s *SignerV4) Sign(method, uri string, headers map[string]string, body []byte) error {
	//if s.region == "" {
	//	return errSignerV4MissingRegion
	//}
	//
	//uri, urlParams, err := s.parseUri(uri)
	//if err != nil {
	//	return err
	//}
	//
	//dateTime, ok := headers[HTTPHeaderLogDate]
	//if !ok {
	//	return fmt.Errorf("can't find '%s' header", HTTPHeaderLogDate)
	//}
	//date := dateTime[:8]
	//// Host should not contain schema here.
	//if host, ok := headers[HTTPHeaderHost]; ok {
	//	if strings.HasPrefix(host, "http://") {
	//		headers[HTTPHeaderHost] = host[len("http://"):]
	//	} else if strings.HasPrefix(host, "https://") {
	//		headers[HTTPHeaderHost] = host[len("https://"):]
	//	}
	//}
	//
	//contentLength := len(body)
	//var sha256Payload string
	//if contentLength != 0 {
	//	sha256Payload = fmt.Sprintf("%x", sha256.Sum256(body))
	//} else {
	//	sha256Payload = emptyStringSha256
	//}
	//headers[HTTPHeaderLogContentSha256] = sha256Payload
	//headers[HTTPHeaderContentLength] = strconv.Itoa(contentLength)
	//
	//// Canonical headers
	//signedHeadersStr, canonicalHeaderStr := s.buildCanonicalHeaders(headers)
	//
	//// CanonicalRequest
	//canonReq := s.buildCanonicalRequest(method, uri, sha256Payload, canonicalHeaderStr, signedHeadersStr, urlParams)
	//scope := s.buildScope(date, s.region)
	//
	//// SignKey + signMessage => signature
	//strToSign := s.buildSignMessage(canonReq, dateTime, scope)
	//key, err := s.buildSigningKey(s.accessKeySecret, s.region, date)
	//if err != nil {
	//	return err
	//}
	//hash, err := s.hmacSha256([]byte(strToSign), key)
	//if err != nil {
	//	return err
	//}
	//signature := hex.EncodeToString(hash)
	//headers[HTTPHeaderAuthorization] = s.buildAuthorization(s.accessKeyID, signature, scope)
	return nil
}
