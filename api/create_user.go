package api

import (
	"github.com/ChainuGO/ChainuGO/request_define"
)

// CreateUser create user
// @param openId user open id
// @return data, timestamp, sign, clientSign, error
func (s *Sdk) CreateUser(openId string) ([]byte, string, string, string, error) {

	return s.signPack(
		request_define.RequestCreateUser{
			OpenID: openId,
		},
	)
}
