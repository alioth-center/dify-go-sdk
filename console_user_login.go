package dify

import (
	"context"
	"encoding/json"
	"github.com/pkg/errors"
)

type UserLoginParams struct {
	Email      string `json:"email"`
	Password   string `json:"password"`
	RememberMe bool   `json:"remember_me"`
}

type UserLoginResponse struct {
	Result string `json:"result"`
	Data   string `json:"data"`
}

func (cl *Client) UserLogin(ctx context.Context, email string, password string) (result UserLoginResponse, err error) {
	var payload = UserLoginParams{
		Email:      email,
		Password:   password,
		RememberMe: true,
	}

	api := cl.GetConsoleAPI(ConsoleApiLogin)

	code, body, err := cl.sendPostRequestToConsole(ctx, api, payload)

	err = CommonRiskForSendRequest(code, err)
	if err != nil {
		return result, err
	}

	err = json.Unmarshal(body, &result)
	if err != nil {
		return result, errors.Wrap(err, "failed to unmarshal the response")
	}

	return result, nil
}
