package user

type UserExt struct {
	UserId int64  `json:"user_id" xorm:"not null pk comment('用户ID') BIGINT"`
	Phone  string `json:"phone" xorm:"not null comment('手机号码') VARCHAR(50)"`
	Email  string `json:"email" xorm:"not null comment('邮箱') VARCHAR(50)"`
	Wechat string `json:"wechat" xorm:"not null comment('微信号') VARCHAR(50)"`
}

type UserExtHelper struct {
}

func (cols UserExtHelper) UserId() string {
	return "user_id"
}

func (cols UserExtHelper) UserIdWT() string {
	return "user_ext.user_id"
}

func (cols UserExtHelper) EqUserId() string {
	return "user_id = ?"
}

func (cols UserExtHelper) EqUserIdWT() string {
	return "user_ext.user_id = ?"
}

func (cols UserExtHelper) NeUserId() string {
	return "user_id != ?"
}

func (cols UserExtHelper) NeUserIdWT() string {
	return "user_ext.user_id != ?"
}

func (cols UserExtHelper) GtUserId() string {
	return "user_id > ?"
}

func (cols UserExtHelper) GtUserIdWT() string {
	return "user_ext.user_id > ?"
}

func (cols UserExtHelper) GteUserId() string {
	return "user_id >= ?"
}

func (cols UserExtHelper) GteUserIdWT() string {
	return "user_ext.user_id >= ?"
}

func (cols UserExtHelper) LtUserId() string {
	return "user_id < ?"
}

func (cols UserExtHelper) LtUserIdWT() string {
	return "user_ext.user_id < ?"
}

func (cols UserExtHelper) LteUserId() string {
	return "user_id <= ?"
}

func (cols UserExtHelper) LteUserIdWT() string {
	return "user_ext.user_id <= ?"
}

func (cols UserExtHelper) Phone() string {
	return "phone"
}

func (cols UserExtHelper) PhoneWT() string {
	return "user_ext.phone"
}

func (cols UserExtHelper) EqPhone() string {
	return "phone = ?"
}

func (cols UserExtHelper) EqPhoneWT() string {
	return "user_ext.phone = ?"
}

func (cols UserExtHelper) NePhone() string {
	return "phone != ?"
}

func (cols UserExtHelper) NePhoneWT() string {
	return "user_ext.phone != ?"
}

func (cols UserExtHelper) GtPhone() string {
	return "phone > ?"
}

func (cols UserExtHelper) GtPhoneWT() string {
	return "user_ext.phone > ?"
}

func (cols UserExtHelper) GtePhone() string {
	return "phone >= ?"
}

func (cols UserExtHelper) GtePhoneWT() string {
	return "user_ext.phone >= ?"
}

func (cols UserExtHelper) LtPhone() string {
	return "phone < ?"
}

func (cols UserExtHelper) LtPhoneWT() string {
	return "user_ext.phone < ?"
}

func (cols UserExtHelper) LtePhone() string {
	return "phone <= ?"
}

func (cols UserExtHelper) LtePhoneWT() string {
	return "user_ext.phone <= ?"
}

func (cols UserExtHelper) Email() string {
	return "email"
}

func (cols UserExtHelper) EmailWT() string {
	return "user_ext.email"
}

func (cols UserExtHelper) EqEmail() string {
	return "email = ?"
}

func (cols UserExtHelper) EqEmailWT() string {
	return "user_ext.email = ?"
}

func (cols UserExtHelper) NeEmail() string {
	return "email != ?"
}

func (cols UserExtHelper) NeEmailWT() string {
	return "user_ext.email != ?"
}

func (cols UserExtHelper) GtEmail() string {
	return "email > ?"
}

func (cols UserExtHelper) GtEmailWT() string {
	return "user_ext.email > ?"
}

func (cols UserExtHelper) GteEmail() string {
	return "email >= ?"
}

func (cols UserExtHelper) GteEmailWT() string {
	return "user_ext.email >= ?"
}

func (cols UserExtHelper) LtEmail() string {
	return "email < ?"
}

func (cols UserExtHelper) LtEmailWT() string {
	return "user_ext.email < ?"
}

func (cols UserExtHelper) LteEmail() string {
	return "email <= ?"
}

func (cols UserExtHelper) LteEmailWT() string {
	return "user_ext.email <= ?"
}

func (cols UserExtHelper) Wechat() string {
	return "wechat"
}

func (cols UserExtHelper) WechatWT() string {
	return "user_ext.wechat"
}

func (cols UserExtHelper) EqWechat() string {
	return "wechat = ?"
}

func (cols UserExtHelper) EqWechatWT() string {
	return "user_ext.wechat = ?"
}

func (cols UserExtHelper) NeWechat() string {
	return "wechat != ?"
}

func (cols UserExtHelper) NeWechatWT() string {
	return "user_ext.wechat != ?"
}

func (cols UserExtHelper) GtWechat() string {
	return "wechat > ?"
}

func (cols UserExtHelper) GtWechatWT() string {
	return "user_ext.wechat > ?"
}

func (cols UserExtHelper) GteWechat() string {
	return "wechat >= ?"
}

func (cols UserExtHelper) GteWechatWT() string {
	return "user_ext.wechat >= ?"
}

func (cols UserExtHelper) LtWechat() string {
	return "wechat < ?"
}

func (cols UserExtHelper) LtWechatWT() string {
	return "user_ext.wechat < ?"
}

func (cols UserExtHelper) LteWechat() string {
	return "wechat <= ?"
}

func (cols UserExtHelper) LteWechatWT() string {
	return "user_ext.wechat <= ?"
}

func (cols UserExtHelper) TableName() string {
	return "user_ext"
}

// OUserExt .
var OUserExt = UserExtHelper{}
