package user

type UserSource struct {
	Id       int64  `json:"id" xorm:"'id' pk autoincr comment('用户ID') BIGINT"`
	UserName string `json:"user_name" xorm:"'user_name' not null comment('用户名') unique(name_ve) VARCHAR(100)"`
	UserVe   string `json:"user_ve" xorm:"'user_ve' not null comment('用户验证设备类型') unique(name_ve) VARCHAR(50)"`
	UserId   int64  `json:"user_id" xorm:"'user_id' not null default 0 comment('关联的实际用户ID') BIGINT"`
}

type UserSourceHelper struct {
}

func (cols UserSourceHelper) Id() string {
	return "id"
}

func (cols UserSourceHelper) IdWT() string {
	return "user_source.id"
}

func (cols UserSourceHelper) EqId() string {
	return "id = ?"
}

func (cols UserSourceHelper) EqIdWT() string {
	return "user_source.id = ?"
}

func (cols UserSourceHelper) NeId() string {
	return "id != ?"
}

func (cols UserSourceHelper) NeIdWT() string {
	return "user_source.id != ?"
}

func (cols UserSourceHelper) GtId() string {
	return "id > ?"
}

func (cols UserSourceHelper) GtIdWT() string {
	return "user_source.id > ?"
}

func (cols UserSourceHelper) GteId() string {
	return "id >= ?"
}

func (cols UserSourceHelper) GteIdWT() string {
	return "user_source.id >= ?"
}

func (cols UserSourceHelper) LtId() string {
	return "id < ?"
}

func (cols UserSourceHelper) LtIdWT() string {
	return "user_source.id < ?"
}

func (cols UserSourceHelper) LteId() string {
	return "id <= ?"
}

func (cols UserSourceHelper) LteIdWT() string {
	return "user_source.id <= ?"
}

func (cols UserSourceHelper) UserName() string {
	return "user_name"
}

func (cols UserSourceHelper) UserNameWT() string {
	return "user_source.user_name"
}

func (cols UserSourceHelper) EqUserName() string {
	return "user_name = ?"
}

func (cols UserSourceHelper) EqUserNameWT() string {
	return "user_source.user_name = ?"
}

func (cols UserSourceHelper) NeUserName() string {
	return "user_name != ?"
}

func (cols UserSourceHelper) NeUserNameWT() string {
	return "user_source.user_name != ?"
}

func (cols UserSourceHelper) GtUserName() string {
	return "user_name > ?"
}

func (cols UserSourceHelper) GtUserNameWT() string {
	return "user_source.user_name > ?"
}

func (cols UserSourceHelper) GteUserName() string {
	return "user_name >= ?"
}

func (cols UserSourceHelper) GteUserNameWT() string {
	return "user_source.user_name >= ?"
}

func (cols UserSourceHelper) LtUserName() string {
	return "user_name < ?"
}

func (cols UserSourceHelper) LtUserNameWT() string {
	return "user_source.user_name < ?"
}

func (cols UserSourceHelper) LteUserName() string {
	return "user_name <= ?"
}

func (cols UserSourceHelper) LteUserNameWT() string {
	return "user_source.user_name <= ?"
}

func (cols UserSourceHelper) UserVe() string {
	return "user_ve"
}

func (cols UserSourceHelper) UserVeWT() string {
	return "user_source.user_ve"
}

func (cols UserSourceHelper) EqUserVe() string {
	return "user_ve = ?"
}

func (cols UserSourceHelper) EqUserVeWT() string {
	return "user_source.user_ve = ?"
}

func (cols UserSourceHelper) NeUserVe() string {
	return "user_ve != ?"
}

func (cols UserSourceHelper) NeUserVeWT() string {
	return "user_source.user_ve != ?"
}

func (cols UserSourceHelper) GtUserVe() string {
	return "user_ve > ?"
}

func (cols UserSourceHelper) GtUserVeWT() string {
	return "user_source.user_ve > ?"
}

func (cols UserSourceHelper) GteUserVe() string {
	return "user_ve >= ?"
}

func (cols UserSourceHelper) GteUserVeWT() string {
	return "user_source.user_ve >= ?"
}

func (cols UserSourceHelper) LtUserVe() string {
	return "user_ve < ?"
}

func (cols UserSourceHelper) LtUserVeWT() string {
	return "user_source.user_ve < ?"
}

func (cols UserSourceHelper) LteUserVe() string {
	return "user_ve <= ?"
}

func (cols UserSourceHelper) LteUserVeWT() string {
	return "user_source.user_ve <= ?"
}

func (cols UserSourceHelper) UserId() string {
	return "user_id"
}

func (cols UserSourceHelper) UserIdWT() string {
	return "user_source.user_id"
}

func (cols UserSourceHelper) EqUserId() string {
	return "user_id = ?"
}

func (cols UserSourceHelper) EqUserIdWT() string {
	return "user_source.user_id = ?"
}

func (cols UserSourceHelper) NeUserId() string {
	return "user_id != ?"
}

func (cols UserSourceHelper) NeUserIdWT() string {
	return "user_source.user_id != ?"
}

func (cols UserSourceHelper) GtUserId() string {
	return "user_id > ?"
}

func (cols UserSourceHelper) GtUserIdWT() string {
	return "user_source.user_id > ?"
}

func (cols UserSourceHelper) GteUserId() string {
	return "user_id >= ?"
}

func (cols UserSourceHelper) GteUserIdWT() string {
	return "user_source.user_id >= ?"
}

func (cols UserSourceHelper) LtUserId() string {
	return "user_id < ?"
}

func (cols UserSourceHelper) LtUserIdWT() string {
	return "user_source.user_id < ?"
}

func (cols UserSourceHelper) LteUserId() string {
	return "user_id <= ?"
}

func (cols UserSourceHelper) LteUserIdWT() string {
	return "user_source.user_id <= ?"
}

func (cols UserSourceHelper) TableName() string {
	return "user_source"
}

// OUserSource .
var OUserSource = UserSourceHelper{}
