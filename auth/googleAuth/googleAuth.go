package googleauth

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/url"

	"github.com/gofiber/fiber/v2"
	mlog "github.com/mike504110403/goutils/log"
	"github.com/tidwall/gjson"
	"github.com/valyala/fasthttp"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var cfg = Config{}
var client = &fasthttp.Client{}

func GetUserInfoByCode(code string) (ScopeEmailData, error) {
	var resData ScopeEmailData
	oauthConfig := &oauth2.Config{
		RedirectURL:  cfg.RedirectURL,
		ClientID:     cfg.ClientID,
		ClientSecret: cfg.SecretKey,
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email", "https://www.googleapis.com/auth/userinfo.profile"},
		Endpoint:     google.Endpoint,
	}
	token, err := oauthConfig.Exchange(context.Background(), code)
	if err != nil {
		return resData, err
	}
	return GetGoogleUserInfo(token.AccessToken)
}

func Init(newCfg Config) {
	cfg = newCfg
}

func LoginInit(redirectUri string) string {
	return fmt.Sprintf("%s?client_id=%s&response_type=%s&scope=%s/%s&redirect_uri=%s",
		UrlLogin,
		cfg.ClientID,
		ResponseTypeCode,
		UrlScope,
		ScopeEmail,
		redirectUri,
	)
}

func GetAccessToken(code string) (token string, err error) {
	req, res := fasthttp.AcquireRequest(), fasthttp.AcquireResponse()

	data := url.Values{
		"code":          {code},
		"client_id":     {cfg.ClientID},
		"client_secret": {cfg.SecretKey},
		"grant_type":    {GrantTypeAuthCode},
		"redirect_uri":  {cfg.RedirectURL},
	}

	req.SetRequestURI(UrlToken)
	req.Header.SetMethod(fiber.MethodPost)
	req.Header.SetContentType(fiber.MIMEApplicationForm)
	req.AppendBody([]byte(data.Encode()))

	if err := client.Do(req, res); err != nil {
		return token, err
	} else {
		fmt.Println(string(res.Body()))
		token = gjson.GetBytes(res.Body(), "access_token").String()
		return token, nil
	}
}

// GetGoogleUserInfo : 以Token取得會員資訊
func GetGoogleUserInfo(accessToken string) (ScopeEmailData, error) {
	req, res, resData := fasthttp.AcquireRequest(), fasthttp.AcquireResponse(), ScopeEmailData{}

	req.SetRequestURI(UrlUserInfo)
	req.Header.SetMethod(fiber.MethodGet)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", accessToken))

	if err := client.Do(req, res); err != nil {
		return resData, err
	} else if res.StatusCode() != fiber.StatusOK {
		mlog.Error(fmt.Sprintf("status: %d, body: %v", res.StatusCode(), res.Body()))
		return resData, errors.New("google oauth回覆錯誤")
	}

	if err := json.Unmarshal(res.Body(), &resData); err != nil {
		return resData, err
	} else {
		if resData.Error.Code == 401 {
			return resData, errors.New(resData.Error.Message)
		}
		return resData, nil
	}
}
