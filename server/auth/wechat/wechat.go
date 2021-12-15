package wechat

import (
	"encoding/json"
	"fmt"

	"github.com/medivhzhan/weapp/v2"
)

const appid string = "wx995dffbc5040a95c"

const appSecret string = "ae878a7757dad39c15726849e4b6df4c"

type Service struct {
}

func (*Service) Resolver(code string) (string, error) {
	res, err := weapp.Login(appid, appSecret, code)

	if err != nil {
		return "", fmt.Errorf("weapp.Login: %v", err)
	}

	if err := res.GetResponseError(); err != nil {
		return "", fmt.Errorf("weapp response err: %v", err)
	}

	s, _ := json.Marshal(res)

	fmt.Printf("%s", s)

	return res.OpenID, nil
}
