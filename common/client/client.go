// Copyright 2020 yhyzgn gtc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// author : 颜洪毅
// e-mail : yhyzgn@gmail.com
// time   : 2020-11-25 16:44
// version: 1.0.0
// desc   : 

package client

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/yhyzgn/gog"
	"github.com/yhyzgn/gtc/common"
	"github.com/yhyzgn/gtc/common/errors"
	"github.com/yhyzgn/gtc/common/profile"
	"math/rand"
	"net/http"
	"net/http/httputil"
	"strconv"
	"time"
)

type Client struct {
	host       string
	profile    *profile.Profile
	credential *common.Credential
	request    *Request
	response   *Response
	doer       *http.Client
}

func New(region common.Region) *Client {
	return NewWithRequest(NewRequest(region))
}

func NewWithRequest(req *Request) *Client {
	return &Client{
		profile:  profile.New(),
		request:  req,
		response: NewResponse(),
		doer:     http.DefaultClient,
	}
}

func (c *Client) Profile(prof *profile.Profile) *Client {
	c.profile = prof
	return c
}

func (c *Client) Host(host string) *Client {
	c.host = host
	return c
}

func (c *Client) Secret(secretId, secretKey string) *Client {
	c.credential = common.NewCredential(secretId, secretKey)
	return c
}

func (c *Client) Credential(credential *common.Credential) *Client {
	c.credential = credential
	return c
}

func (c *Client) Request(request *Request) *Client {
	c.request = request
	return c
}

func (c *Client) Option(opts ...Optional) *Client {
	if c.request != nil {
		c.request.Option(opts...)
	}
	return c
}

func (c *Client) Response(response *Response) *Client {
	c.response = response
	return c
}

//
// 支持的 HTTP 请求方法:
//
// POST（推荐）
// GET
// POST 请求支持的 Content-Type 类型：
//
// application/json（推荐），必须使用签名方法 v3（TC3-HMAC-SHA256）。
// application/x-www-form-urlencoded，必须使用签名方法 v1（HmacSHA1 或 HmacSHA256）。
// multipart/form-data（仅部分接口支持），必须使用签名方法 v3（TC3-HMAC-SHA256）。
// GET 请求的请求包大小不得超过32KB。POST 请求使用签名方法 v1（HmacSHA1、HmacSHA256）时不得超过1MB。POST 请求使用签名方法 v3（TC3-HMAC-SHA256）时支持10MB。
//
func (c *Client) Do(result interface{}) error {
	if c.request == nil {
		return errors.New("GTC.RequestError", common.InternalRequestID, "Request of Client can not be nil.")
	}
	if c.response == nil {
		return errors.New("GTC.ResponseError", common.InternalRequestID, "Response of Client can not be nil.")
	}

	if c.request.region != "" {
		c.request.Option(Region(c.request.region))
	}
	if c.host != "" {
		c.request.Option(Domain(c.host))
	}

	if c.request.IsContentType(common.ContentType.XForm) && c.isSignV3() {
		return errors.New("GTC.SignMethodError", common.InternalRequestID, "Must use sign method 'HmacSHA1' or 'HmacSHA256' when Content-Type is 'application/x-www-form-urlencoded'.")
	}

	if (c.request.IsContentType(common.ContentType.JSON) || c.request.IsContentType(common.ContentType.FormData)) && !c.isSignV3() {
		return errors.New("GTC.SignMethodError", common.InternalRequestID, "Must use sign method 'TC3-HMAC-SHA256' when Content-Type is 'application/json' or 'multipart/form-data'.")
	}

	if c.isSignV3() {
		return c.doWithSignV3(result)
	}
	return c.doWithSignV1(result)
}

func (c *Client) completeCommonParams() {
	params := c.request.GetParams()
	params["Region"] = c.request.GetRegion()
	if c.request.GetVersion() != "" {
		params["Version"] = c.request.GetVersion()
	}
	params["Action"] = c.request.GetAction()
	params["Timestamp"] = strconv.FormatInt(time.Now().Unix(), 10)
	params["Nonce"] = strconv.Itoa(rand.Int())
	params["RequestClient"] = common.SDK.Version
}

func (c *Client) doWithSignV1(result interface{}) error {
	params := c.genCommonParams(false)
	c.request.SetParams(params)

	// 签名
	c.signRequest()

	ctx, cancel := context.WithTimeout(context.Background(), c.profile.Timeout)
	defer cancel()

	realRequest, err := http.NewRequestWithContext(ctx, c.request.GetHttpMethod(), c.request.GetEndpoint(), c.request.GetReader())
	if err != nil {
		return err
	}

	realRequest.Header.Set("Host", c.request.GetDomain())
	realRequest.Header.Set("Content-Type", c.request.GetContentType())

	if c.profile.Debug {
		if bs, err := httputil.DumpRequest(realRequest, true); err != nil {
			gog.ErrorF("dump request failed because {}", err)
		} else {
			gog.DebugF("[http request] <==>\n{}", string(bs))
		}
	}

	realResponse, err := c.doer.Do(realRequest)
	if err != nil {
		return errors.New("GTC.NetworkError", common.InternalRequestID, fmt.Sprintf("Fail to get response because %s", err))
	}

	return c.decodeResponse(realResponse, result)
}

func (c *Client) doWithSignV3(result interface{}) error {
	params := c.genCommonParams(true)
	header := make(map[string]string, len(params))
	for name, value := range params {
		header[name] = fmt.Sprintf("%v", value)
	}
	if c.request.GetHttpMethod() == http.MethodGet {
		header["Content-Type"] = common.ContentType.XForm
	} else {
		header["Content-Type"] = c.request.GetContentType()
	}

	httpRequestMethod := c.request.GetHttpMethod()
	canonicalURI := "/"
	canonicalQueryString := ""
	if httpRequestMethod == http.MethodGet {
		params := make(map[string]interface{})
		for key, value := range c.request.GetParams() {
			params[key] = value
		}
		canonicalQueryString = QueryParams(params)
	}
	canonicalHeader := fmt.Sprintf("content-type:%s\nhost:%s\n", header["Content-Type"], header["Host"])
	signedHeader := "content-type;host"
	requestPayload := ""
	if httpRequestMethod == http.MethodPost && c.request.GetBody() != nil {
		bs, err := json.Marshal(c.request.GetBody())
		if err != nil {
			return err
		}
		requestPayload = string(bs)
	}

	hashedRequestPayload := ""
	if c.profile.UnsignedPayload {
		hashedRequestPayload = common.Sha256Hex("UNSIGNED-PAYLOAD")
		header["X-TC-Content-SHA256"] = "UNSIGNED-PAYLOAD"
	} else {
		hashedRequestPayload = common.Sha256Hex(requestPayload)
	}

	// 1. 拼接规范请求串
	// https://cloud.tencent.com/document/product/406/42617#1.-.E6.8B.BC.E6.8E.A5.E8.A7.84.E8.8C.83.E8.AF.B7.E6.B1.82.E4.B8.B2
	canonicalRequest := fmt.Sprintf("%s\n%s\n%s\n%s\n%s\n%s",
		httpRequestMethod,
		canonicalURI,
		canonicalQueryString,
		canonicalHeader,
		signedHeader,
		hashedRequestPayload)
	c.debug(canonicalRequest)

	// 2. 拼接待签名字符串
	// https://cloud.tencent.com/document/product/406/42617#2.-.E6.8B.BC.E6.8E.A5.E5.BE.85.E7.AD.BE.E5.90.8D.E5.AD.97.E7.AC.A6.E4.B8.B2
	algorithm := common.SignMethod.TC3HmacSha256
	requestTimestamp := header["X-TC-Timestamp"]
	timestamp, _ := strconv.ParseInt(requestTimestamp, 10, 64)
	date := time.Unix(timestamp, 0).UTC().Format("2006-01-02")
	credentialScope := fmt.Sprintf("%s/%s/tc3_request", date, c.request.GetService())
	hashedCanonicalRequest := common.Sha256Hex(canonicalRequest)
	stringToSign := fmt.Sprintf("%s\n%s\n%s\n%s", algorithm, requestTimestamp, credentialScope, hashedCanonicalRequest)
	c.debug(stringToSign)

	// 3. 计算签名
	// https://cloud.tencent.com/document/product/406/42617#3.-.E8.AE.A1.E7.AE.97.E7.AD.BE.E5.90.8D
	secretDate := common.HmacSha256(date, "TC3"+c.credential.SecretKey)
	secretService := common.HmacSha256(c.request.GetService(), secretDate)
	secretSigning := common.HmacSha256("tc3_request", secretService)
	signature := hex.EncodeToString([]byte(common.HmacSha256(stringToSign, secretSigning)))
	c.debug(signature)

	// 4. 拼接 Authorization
	// https://cloud.tencent.com/document/product/406/42617#4.-.E6.8B.BC.E6.8E.A5-Authorization
	authorization := fmt.Sprintf("%s Credential=%s/%s, SignedHeaders=%s, Signature=%s",
		algorithm,
		c.credential.SecretID,
		credentialScope,
		signedHeader,
		signature)
	c.debug(authorization)
	header["Authorization"] = authorization

	// 5. 请求
	ctx, cancel := context.WithTimeout(context.Background(), c.profile.Timeout)
	defer cancel()

	fmt.Println(c.request.GetEndpoint())
	realRequest, err := http.NewRequestWithContext(ctx, c.request.GetHttpMethod(), c.request.GetEndpoint(), c.request.GetReader())
	if err != nil {
		return err
	}

	for key, value := range header {
		realRequest.Header.Set(key, value)
	}

	if c.profile.Debug {
		if bs, err := httputil.DumpRequest(realRequest, true); err != nil {
			gog.ErrorF("dump request failed because {}", err)
		} else {
			gog.DebugF("[http request] <==>\n{}", string(bs))
		}
	}

	realResponse, err := c.doer.Do(realRequest)
	if err != nil {
		return errors.New("GTC.NetworkError", common.InternalRequestID, fmt.Sprintf("Fail to get response because %s", err))
	}

	return c.decodeResponse(realResponse, result)
}

func (c *Client) decodeResponse(realResponse *http.Response, result interface{}) error {
	err := c.response.Decode(realResponse.Body, result)
	if err != nil {
		return errors.New("GTC.JSONDecoderError", common.InternalRequestID, fmt.Sprintf("Fail to decode json from response body because %s", err))
	}
	return nil
}

func (c *Client) isSignV3() bool {
	return c.request.GetHttpMethod() == http.MethodPost && c.profile.SignMethod == common.SignMethod.TC3HmacSha256
}

func (c *Client) genCommonParams(isSignV3 bool) (params map[string]interface{}) {
	if isSignV3 {
		// v3 签名
		params = map[string]interface{}{
			"Host":               c.request.GetDomain(),
			"X-TC-Action":        c.request.GetAction(),
			"X-TC-Region":        c.request.GetRegion(),
			"X-TC-Timestamp":     strconv.FormatInt(time.Now().Unix(), 10),
			"X-TC-Version":       c.request.GetVersion(),
			"X-TC-RequestClient": common.SDK.Version,
			"X-TC-Language":      c.profile.Language,
		}
		if c.credential.Token != "" {
			params["X-TC-Token"] = c.credential.Token
		}
		// Authorization 参数需签名后再设置
	} else {
		// v1 签名
		params = map[string]interface{}{
			"Action":          c.request.GetAction(),
			"Region":          c.request.GetRegion(),
			"Timestamp":       strconv.FormatInt(time.Now().Unix(), 10),
			"Nonce":           strconv.Itoa(rand.Int()),
			"SecretId":        c.credential.SecretID,
			"Version":         c.request.GetVersion(),
			"SignatureMethod": c.profile.SignMethod,
			"RequestClient":   common.SDK.Version,
			"Language":        c.profile.Language,
		}
		if c.credential.Token != "" {
			params["Token"] = c.credential.Token
		}
		// Signature 参数需签名后再设置
	}
	return
}

func (c *Client) signRequest() {
	if c.profile.SignMethod != common.SignMethod.HmacSHA256 {
		c.profile.SignMethod = common.SignMethod.HmacSHA1
	}
	pre := common.PreSignString(c.request.GetHttpMethod(), c.request.GetDomain(), c.request.GetPath(), c.request.GetParams())
	signature := common.Sign(c.profile.SignMethod, c.credential.SecretKey, pre)
	c.request.SetParam("Signature", signature)
}

func (c *Client) debug(args ...interface{}) {
	if c.profile.Debug {
		gog.Debug(args...)
	}
}

/*
curl -X POST https://cvm.ap-guangzhou.tencentcloudapi.com
 -H "Authorization: TC3-HMAC-SHA256 Credential=xxxxxxxxxxxxxxxxxxxxxxxxxxxx/2020-12-04/cvm/tc3_request, SignedHeaders=content-type;host, Signature=f64dd58bf6f695bbcebc9cac6add02ee79f15854a3fd99a2840ed7c470d774c2"
 -H "Content-Type: application/json"
 -H "Host: cvm.ap-guangzhou.tencentcloudapi.com"
 -H "X-TC-Action: DescribeInstances"xxxxxxxxxxxxxxxxxxxxxxxxxxxx
 -H "X-TC-Timestamp: 1607056961"
 -H "X-TC-Version: 2017-03-12"
 -H "X-TC-Region: ap-guangzhou"
 -d '{"InstanceIds.0":"ins-09dx96dg","Limit":20,"Offset":0}'
*/
