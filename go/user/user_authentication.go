package user

type UserAuthentication struct {
	UserId   int64  `json:"user_id" xorm:"'user_id' not null pk comment('用户ID') BIGINT"`
	Password string `json:"password" xorm:"'password' not null comment('密码') VARCHAR(100)"`
	Token2fa string `json:"token_2fa" xorm:"'token_2fa' not null comment('二次验证token') VARCHAR(256)"`
}

type UserAuthenticationHelper struct {
}

func (cols UserAuthenticationHelper) UserId() string {
	return "user_id"
}

func (cols UserAuthenticationHelper) UserIdWT() string {
	return "user_authentication.user_id"
}

func (cols UserAuthenticationHelper) EqUserId() string {
	return "user_id = ?"
}

func (cols UserAuthenticationHelper) EqUserIdWT() string {
	return "user_authentication.user_id = ?"
}

func (cols UserAuthenticationHelper) NeUserId() string {
	return "user_id != ?"
}

func (cols UserAuthenticationHelper) NeUserIdWT() string {
	return "user_authentication.user_id != ?"
}

func (cols UserAuthenticationHelper) GtUserId() string {
	return "user_id > ?"
}

func (cols UserAuthenticationHelper) GtUserIdWT() string {
	return "user_authentication.user_id > ?"
}

func (cols UserAuthenticationHelper) GteUserId() string {
	return "user_id >= ?"
}

func (cols UserAuthenticationHelper) GteUserIdWT() string {
	return "user_authentication.user_id >= ?"
}

func (cols UserAuthenticationHelper) LtUserId() string {
	return "user_id < ?"
}

func (cols UserAuthenticationHelper) LtUserIdWT() string {
	return "user_authentication.user_id < ?"
}

func (cols UserAuthenticationHelper) LteUserId() string {
	return "user_id <= ?"
}

func (cols UserAuthenticationHelper) LteUserIdWT() string {
	return "user_authentication.user_id <= ?"
}

func (cols UserAuthenticationHelper) Password() string {
	return "password"
}

func (cols UserAuthenticationHelper) PasswordWT() string {
	return "user_authentication.password"
}

func (cols UserAuthenticationHelper) EqPassword() string {
	return "password = ?"
}

func (cols UserAuthenticationHelper) EqPasswordWT() string {
	return "user_authentication.password = ?"
}

func (cols UserAuthenticationHelper) NePassword() string {
	return "password != ?"
}

func (cols UserAuthenticationHelper) NePasswordWT() string {
	return "user_authentication.password != ?"
}

func (cols UserAuthenticationHelper) GtPassword() string {
	return "password > ?"
}

func (cols UserAuthenticationHelper) GtPasswordWT() string {
	return "user_authentication.password > ?"
}

func (cols UserAuthenticationHelper) GtePassword() string {
	return "password >= ?"
}

func (cols UserAuthenticationHelper) GtePasswordWT() string {
	return "user_authentication.password >= ?"
}

func (cols UserAuthenticationHelper) LtPassword() string {
	return "password < ?"
}

func (cols UserAuthenticationHelper) LtPasswordWT() string {
	return "user_authentication.password < ?"
}

func (cols UserAuthenticationHelper) LtePassword() string {
	return "password <= ?"
}

func (cols UserAuthenticationHelper) LtePasswordWT() string {
	return "user_authentication.password <= ?"
}

func (cols UserAuthenticationHelper) Token2fa() string {
	return "token_2fa"
}

func (cols UserAuthenticationHelper) Token2faWT() string {
	return "user_authentication.token_2fa"
}

func (cols UserAuthenticationHelper) EqToken2fa() string {
	return "token_2fa = ?"
}

func (cols UserAuthenticationHelper) EqToken2faWT() string {
	return "user_authentication.token_2fa = ?"
}

func (cols UserAuthenticationHelper) NeToken2fa() string {
	return "token_2fa != ?"
}

func (cols UserAuthenticationHelper) NeToken2faWT() string {
	return "user_authentication.token_2fa != ?"
}

func (cols UserAuthenticationHelper) GtToken2fa() string {
	return "token_2fa > ?"
}

func (cols UserAuthenticationHelper) GtToken2faWT() string {
	return "user_authentication.token_2fa > ?"
}

func (cols UserAuthenticationHelper) GteToken2fa() string {
	return "token_2fa >= ?"
}

func (cols UserAuthenticationHelper) GteToken2faWT() string {
	return "user_authentication.token_2fa >= ?"
}

func (cols UserAuthenticationHelper) LtToken2fa() string {
	return "token_2fa < ?"
}

func (cols UserAuthenticationHelper) LtToken2faWT() string {
	return "user_authentication.token_2fa < ?"
}

func (cols UserAuthenticationHelper) LteToken2fa() string {
	return "token_2fa <= ?"
}

func (cols UserAuthenticationHelper) LteToken2faWT() string {
	return "user_authentication.token_2fa <= ?"
}

func (cols UserAuthenticationHelper) TableName() string {
	return "user_authentication"
}

// OUserAuthentication .
var OUserAuthentication = UserAuthenticationHelper{}
