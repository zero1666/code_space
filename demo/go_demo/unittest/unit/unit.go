package unit

import (
	"errors"
	"fmt"
	"git.code.oa.com/trpc-go/trpc-go"
	thttp "git.code.oa.com/trpc-go/trpc-go/http"
	"regexp"
)

type Detail struct {
	Username string
	Email    string
}

type GetPersonDetailRsp struct {
	RetCode int32   `json:"retCode"`
	Result  *Detail `json:"result"`
	RetMsg  string  `json:"retMsg"`
}

// 检查用户名是否非法
func checkUsername(username string) bool {
	const pattern = `^[a-z0-9_-]{3,16}$`

	reg := regexp.MustCompile(pattern)
	return reg.MatchString(username)
}

// 检查用户邮箱是否非法
func checkEmail(email string) bool {
	const pattern = `^[a-zA-Z0-9_-]+@[a-zA-Z0-9_-]+(\.[a-zA-Z0-9_-]+)+$`

	reg := regexp.MustCompile(pattern)
	return reg.MatchString(email)
}

// 通过 http 拉取对应用户的资料信息
func getPersonDetailHttp(username string) (*Detail, error) {
	ctx := trpc.BackgroundContext()

	proxy := thttp.NewClientProxy("trpc.person.svr.cgi")
	rsp := &GetPersonDetailRsp{}
	err := proxy.Get(ctx, fmt.Sprintf("/getPersonDetail?username=%v", username), rsp)
	if err != nil {
		return nil, err
	}

	return rsp.Result, nil
}

// 拉取用户资料信息并校验
func GetPersonDetail(username string) (*Detail, error) {
	// 检查用户名是否有效
	if ok := checkUsername(username); !ok {
		return nil, errors.New("invalid username")
	}

	// 从 http 接口获取信息
	detail, err := getPersonDetailHttp(username)
	if err != nil {
		return nil, err
	}

	// 校验
	if ok := checkEmail(detail.Email); !ok {
		return nil, errors.New("invalid email")
	}

	return detail, nil
}
