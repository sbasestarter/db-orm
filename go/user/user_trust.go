package user

type UserTrust struct {
	Cnt    int    `json:"cnt" xorm:"not null comment('登录次数') INT"`
	Id     int64  `json:"id" xorm:"pk autoincr BIGINT"`
	Ip     string `json:"ip" xorm:"not null comment('IP地址') unique(id_ip) VARCHAR(50)"`
	UserId int64  `json:"user_id" xorm:"not null comment('用户ID') unique(id_ip) BIGINT"`
}

type UserTrustHelper struct {
}

func (cols UserTrustHelper) Cnt() string {
	return "cnt"
}

func (cols UserTrustHelper) CntWT() string {
	return "user_trust.cnt"
}

func (cols UserTrustHelper) EqCnt() string {
	return "cnt = ?"
}

func (cols UserTrustHelper) EqCntWT() string {
	return "user_trust.cnt = ?"
}

func (cols UserTrustHelper) InCnt() string {
	return "cnt in ?"
}

func (cols UserTrustHelper) InCntWT() string {
	return "user_trust.cnt in ?"
}

func (cols UserTrustHelper) NeCnt() string {
	return "cnt != ?"
}

func (cols UserTrustHelper) NeCntWT() string {
	return "user_trust.cnt != ?"
}

func (cols UserTrustHelper) GtCnt() string {
	return "cnt > ?"
}

func (cols UserTrustHelper) GtCntWT() string {
	return "user_trust.cnt > ?"
}

func (cols UserTrustHelper) GteCnt() string {
	return "cnt >= ?"
}

func (cols UserTrustHelper) GteCntWT() string {
	return "user_trust.cnt >= ?"
}

func (cols UserTrustHelper) LtCnt() string {
	return "cnt < ?"
}

func (cols UserTrustHelper) LtCntWT() string {
	return "user_trust.cnt < ?"
}

func (cols UserTrustHelper) LteCnt() string {
	return "cnt <= ?"
}

func (cols UserTrustHelper) LteCntWT() string {
	return "user_trust.cnt <= ?"
}

func (cols UserTrustHelper) Id() string {
	return "id"
}

func (cols UserTrustHelper) IdWT() string {
	return "user_trust.id"
}

func (cols UserTrustHelper) EqId() string {
	return "id = ?"
}

func (cols UserTrustHelper) EqIdWT() string {
	return "user_trust.id = ?"
}

func (cols UserTrustHelper) InId() string {
	return "id in ?"
}

func (cols UserTrustHelper) InIdWT() string {
	return "user_trust.id in ?"
}

func (cols UserTrustHelper) NeId() string {
	return "id != ?"
}

func (cols UserTrustHelper) NeIdWT() string {
	return "user_trust.id != ?"
}

func (cols UserTrustHelper) GtId() string {
	return "id > ?"
}

func (cols UserTrustHelper) GtIdWT() string {
	return "user_trust.id > ?"
}

func (cols UserTrustHelper) GteId() string {
	return "id >= ?"
}

func (cols UserTrustHelper) GteIdWT() string {
	return "user_trust.id >= ?"
}

func (cols UserTrustHelper) LtId() string {
	return "id < ?"
}

func (cols UserTrustHelper) LtIdWT() string {
	return "user_trust.id < ?"
}

func (cols UserTrustHelper) LteId() string {
	return "id <= ?"
}

func (cols UserTrustHelper) LteIdWT() string {
	return "user_trust.id <= ?"
}

func (cols UserTrustHelper) Ip() string {
	return "ip"
}

func (cols UserTrustHelper) IpWT() string {
	return "user_trust.ip"
}

func (cols UserTrustHelper) EqIp() string {
	return "ip = ?"
}

func (cols UserTrustHelper) EqIpWT() string {
	return "user_trust.ip = ?"
}

func (cols UserTrustHelper) InIp() string {
	return "ip in ?"
}

func (cols UserTrustHelper) InIpWT() string {
	return "user_trust.ip in ?"
}

func (cols UserTrustHelper) NeIp() string {
	return "ip != ?"
}

func (cols UserTrustHelper) NeIpWT() string {
	return "user_trust.ip != ?"
}

func (cols UserTrustHelper) GtIp() string {
	return "ip > ?"
}

func (cols UserTrustHelper) GtIpWT() string {
	return "user_trust.ip > ?"
}

func (cols UserTrustHelper) GteIp() string {
	return "ip >= ?"
}

func (cols UserTrustHelper) GteIpWT() string {
	return "user_trust.ip >= ?"
}

func (cols UserTrustHelper) LtIp() string {
	return "ip < ?"
}

func (cols UserTrustHelper) LtIpWT() string {
	return "user_trust.ip < ?"
}

func (cols UserTrustHelper) LteIp() string {
	return "ip <= ?"
}

func (cols UserTrustHelper) LteIpWT() string {
	return "user_trust.ip <= ?"
}

func (cols UserTrustHelper) UserId() string {
	return "user_id"
}

func (cols UserTrustHelper) UserIdWT() string {
	return "user_trust.user_id"
}

func (cols UserTrustHelper) EqUserId() string {
	return "user_id = ?"
}

func (cols UserTrustHelper) EqUserIdWT() string {
	return "user_trust.user_id = ?"
}

func (cols UserTrustHelper) InUserId() string {
	return "user_id in ?"
}

func (cols UserTrustHelper) InUserIdWT() string {
	return "user_trust.user_id in ?"
}

func (cols UserTrustHelper) NeUserId() string {
	return "user_id != ?"
}

func (cols UserTrustHelper) NeUserIdWT() string {
	return "user_trust.user_id != ?"
}

func (cols UserTrustHelper) GtUserId() string {
	return "user_id > ?"
}

func (cols UserTrustHelper) GtUserIdWT() string {
	return "user_trust.user_id > ?"
}

func (cols UserTrustHelper) GteUserId() string {
	return "user_id >= ?"
}

func (cols UserTrustHelper) GteUserIdWT() string {
	return "user_trust.user_id >= ?"
}

func (cols UserTrustHelper) LtUserId() string {
	return "user_id < ?"
}

func (cols UserTrustHelper) LtUserIdWT() string {
	return "user_trust.user_id < ?"
}

func (cols UserTrustHelper) LteUserId() string {
	return "user_id <= ?"
}

func (cols UserTrustHelper) LteUserIdWT() string {
	return "user_trust.user_id <= ?"
}

func (cols UserTrustHelper) TableName() string {
	return "user_trust"
}

// OUserTrust .
var OUserTrust = UserTrustHelper{}
