package authenticationclient

import (
	"fmt"
	"log"
	gohttp "net/http"
	"net/url"
	"strings"
)

const (
	PasswordKey        = "password"
	ContentTypeKey     = "Content-Type"
	XWwwFormUrlEncoded = "application/x-www-form-urlencoded"
	ResourceUrlFormat  = "%s/users/%s/authentication"
)

type Client struct {
	BaseUrl    string
	HttpClient *gohttp.Client
}

func New(baseUrl string) *Client {
	c := Client{
		BaseUrl:    baseUrl,
		HttpClient: gohttp.DefaultClient,
	}

	return &c
}

func (c *Client) AuthenticateUser(username, password string) bool {
	resourceUrl := fmt.Sprintf(ResourceUrlFormat, c.BaseUrl, username)
	data := url.Values{}
	data.Add(PasswordKey, password)

	request, _ := gohttp.NewRequest(gohttp.MethodPost, resourceUrl, strings.NewReader(data.Encode()))
	request.Header.Add(ContentTypeKey, XWwwFormUrlEncoded)

	response, err := c.HttpClient.Do(request)

	if err != nil {
		log.Println(err)
		return false
	}

	return response.StatusCode == gohttp.StatusNoContent
}
