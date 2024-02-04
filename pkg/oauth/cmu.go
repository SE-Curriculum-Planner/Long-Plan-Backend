package oauth

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/SE-Curriculum-Planner/Long-Plan-Backend/config"
	"github.com/SE-Curriculum-Planner/Long-Plan-Backend/pkg/errors"
	"github.com/SE-Curriculum-Planner/Long-Plan-Backend/pkg/lodash"
	"github.com/SE-Curriculum-Planner/Long-Plan-Backend/pkg/requestor"
	"github.com/golang-jwt/jwt/v4"
)

type UserClaims struct {
	User UserDto `json:"user"`
	jwt.RegisteredClaims
}

type accessTokenDto struct {
	AccessToken string `json:"access_token"`
}

type AccType string

const (
	Employee AccType = "MISEmpAcc"
	Student  AccType = "StdAcc"
)

type UserDto struct {
	CmuitaccountName   string  `json:"cmuitaccount_name"`
	Cmuitaccount       string  `json:"cmuitaccount"`
	StudentID          string  `json:"student_id"`
	PrenameID          string  `json:"prename_id"`
	PrenameTH          string  `json:"prename_TH"`
	PrenameEN          string  `json:"prename_EN"`
	FirstnameTH        string  `json:"firstname_TH"`
	FirstnameEN        string  `json:"firstname_EN"`
	LastnameTH         string  `json:"lastname_TH"`
	LastnameEN         string  `json:"lastname_EN"`
	OrganizationCode   string  `json:"organization_code"`
	OrganizationNameTH string  `json:"organization_name_TH"`
	OrganizationNameEN string  `json:"organization_name_EN"`
	ItaccounttypeID    AccType `json:"itaccounttype_id"`
	ItaccounttypeTH    string  `json:"itaccounttype_TH"`
	ItaccounttypeEN    string  `json:"itaccounttype_EN"`
}

func CmuOauthValidation(code string) (*UserDto, error) {
	accessToken, err := getAccessToken(code)
	if err != nil {
		return nil, err
	}
	user, err := getCmuBasicInfo(accessToken.AccessToken)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func getAccessToken(code string) (*accessTokenDto, error) {
	config := config.Config.CmuOauth
	url := config.CmuOauthToken
	header := map[string]string{
		"Content-Type": "application/x-www-form-urlencoded",
	}
	params := map[string]interface{}{
		"code":          code,
		"redirect_uri":  config.CmuOauthRedirectURL,
		"client_id":     config.CmuOauthClientID,
		"client_secret": config.CmuOauthClientSecret,
		"grant_type":    "authorization_code",
	}
	paramsEncode := requestor.BuildQueryParams(params)
	payload := strings.Replace(paramsEncode, "?", "", 1)
	res, statusCode, err := requestor.HttpPost[interface{}](url, header, payload)
	if err != nil {
		return nil, errors.InternalErr(err.Error())
	}

	statusCodeStr := strconv.Itoa(statusCode)
	if strings.HasPrefix(statusCodeStr, "2") {
		var result accessTokenDto
		lodash.Recast(res, &result)
		return &result, nil
	} else {
		return nil, errors.CmuOauthErr("can't get access_token")
	}
}

func getCmuBasicInfo(accessToken string) (*UserDto, error) {
	config := config.Config.CmuOauth
	url := config.CmuOauthInfo
	header := map[string]string{
		"Authorization": fmt.Sprintf("Bearer %v", accessToken),
	}
	res, statusCode, err := requestor.HttpGet[map[string]interface{}](url, header)
	if err != nil {
		return nil, errors.InternalErr(err.Error())
	}

	statusCodeStr := strconv.Itoa(statusCode)
	if strings.HasPrefix(statusCodeStr, "2") {
		var result UserDto
		lodash.Recast(res, &result)
		return &result, nil
	} else {
		return nil, errors.CmuOauthErr("can't get user info")
	}
}
