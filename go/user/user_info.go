package user

import (
	"time"
)

type UserInfo struct {
	Avatar     string    `json:"avatar" xorm:"not null comment('头像') TEXT"`
	CreateAt   time.Time `json:"create_at" xorm:"not null comment('创建时间') DATETIME"`
	NickName   string    `json:"nick_name" xorm:"not null comment('昵称') unique VARCHAR(100)"`
	Privileges int       `json:"privileges" xorm:"not null comment('权限位') INT"`
	UserId     int64     `json:"user_id" xorm:"not null pk autoincr comment('用户ID') BIGINT"`
}

type UserInfoHelper struct {
}

func (cols UserInfoHelper) Avatar() string {
	return "avatar"
}

func (cols UserInfoHelper) AvatarWT() string {
	return "user_info.avatar"
}

func (cols UserInfoHelper) EqAvatar() string {
	return "avatar = ?"
}

func (cols UserInfoHelper) EqAvatarWT() string {
	return "user_info.avatar = ?"
}

func (cols UserInfoHelper) InAvatar() string {
	return "avatar in ?"
}

func (cols UserInfoHelper) InAvatarWT() string {
	return "user_info.avatar in ?"
}

func (cols UserInfoHelper) NeAvatar() string {
	return "avatar != ?"
}

func (cols UserInfoHelper) NeAvatarWT() string {
	return "user_info.avatar != ?"
}

func (cols UserInfoHelper) GtAvatar() string {
	return "avatar > ?"
}

func (cols UserInfoHelper) GtAvatarWT() string {
	return "user_info.avatar > ?"
}

func (cols UserInfoHelper) GteAvatar() string {
	return "avatar >= ?"
}

func (cols UserInfoHelper) GteAvatarWT() string {
	return "user_info.avatar >= ?"
}

func (cols UserInfoHelper) LtAvatar() string {
	return "avatar < ?"
}

func (cols UserInfoHelper) LtAvatarWT() string {
	return "user_info.avatar < ?"
}

func (cols UserInfoHelper) LteAvatar() string {
	return "avatar <= ?"
}

func (cols UserInfoHelper) LteAvatarWT() string {
	return "user_info.avatar <= ?"
}

func (cols UserInfoHelper) CreateAt() string {
	return "create_at"
}

func (cols UserInfoHelper) CreateAtWT() string {
	return "user_info.create_at"
}

func (cols UserInfoHelper) EqCreateAt() string {
	return "create_at = ?"
}

func (cols UserInfoHelper) EqCreateAtWT() string {
	return "user_info.create_at = ?"
}

func (cols UserInfoHelper) InCreateAt() string {
	return "create_at in ?"
}

func (cols UserInfoHelper) InCreateAtWT() string {
	return "user_info.create_at in ?"
}

func (cols UserInfoHelper) NeCreateAt() string {
	return "create_at != ?"
}

func (cols UserInfoHelper) NeCreateAtWT() string {
	return "user_info.create_at != ?"
}

func (cols UserInfoHelper) GtCreateAt() string {
	return "create_at > ?"
}

func (cols UserInfoHelper) GtCreateAtWT() string {
	return "user_info.create_at > ?"
}

func (cols UserInfoHelper) GteCreateAt() string {
	return "create_at >= ?"
}

func (cols UserInfoHelper) GteCreateAtWT() string {
	return "user_info.create_at >= ?"
}

func (cols UserInfoHelper) LtCreateAt() string {
	return "create_at < ?"
}

func (cols UserInfoHelper) LtCreateAtWT() string {
	return "user_info.create_at < ?"
}

func (cols UserInfoHelper) LteCreateAt() string {
	return "create_at <= ?"
}

func (cols UserInfoHelper) LteCreateAtWT() string {
	return "user_info.create_at <= ?"
}

func (cols UserInfoHelper) NickName() string {
	return "nick_name"
}

func (cols UserInfoHelper) NickNameWT() string {
	return "user_info.nick_name"
}

func (cols UserInfoHelper) EqNickName() string {
	return "nick_name = ?"
}

func (cols UserInfoHelper) EqNickNameWT() string {
	return "user_info.nick_name = ?"
}

func (cols UserInfoHelper) InNickName() string {
	return "nick_name in ?"
}

func (cols UserInfoHelper) InNickNameWT() string {
	return "user_info.nick_name in ?"
}

func (cols UserInfoHelper) NeNickName() string {
	return "nick_name != ?"
}

func (cols UserInfoHelper) NeNickNameWT() string {
	return "user_info.nick_name != ?"
}

func (cols UserInfoHelper) GtNickName() string {
	return "nick_name > ?"
}

func (cols UserInfoHelper) GtNickNameWT() string {
	return "user_info.nick_name > ?"
}

func (cols UserInfoHelper) GteNickName() string {
	return "nick_name >= ?"
}

func (cols UserInfoHelper) GteNickNameWT() string {
	return "user_info.nick_name >= ?"
}

func (cols UserInfoHelper) LtNickName() string {
	return "nick_name < ?"
}

func (cols UserInfoHelper) LtNickNameWT() string {
	return "user_info.nick_name < ?"
}

func (cols UserInfoHelper) LteNickName() string {
	return "nick_name <= ?"
}

func (cols UserInfoHelper) LteNickNameWT() string {
	return "user_info.nick_name <= ?"
}

func (cols UserInfoHelper) Privileges() string {
	return "privileges"
}

func (cols UserInfoHelper) PrivilegesWT() string {
	return "user_info.privileges"
}

func (cols UserInfoHelper) EqPrivileges() string {
	return "privileges = ?"
}

func (cols UserInfoHelper) EqPrivilegesWT() string {
	return "user_info.privileges = ?"
}

func (cols UserInfoHelper) InPrivileges() string {
	return "privileges in ?"
}

func (cols UserInfoHelper) InPrivilegesWT() string {
	return "user_info.privileges in ?"
}

func (cols UserInfoHelper) NePrivileges() string {
	return "privileges != ?"
}

func (cols UserInfoHelper) NePrivilegesWT() string {
	return "user_info.privileges != ?"
}

func (cols UserInfoHelper) GtPrivileges() string {
	return "privileges > ?"
}

func (cols UserInfoHelper) GtPrivilegesWT() string {
	return "user_info.privileges > ?"
}

func (cols UserInfoHelper) GtePrivileges() string {
	return "privileges >= ?"
}

func (cols UserInfoHelper) GtePrivilegesWT() string {
	return "user_info.privileges >= ?"
}

func (cols UserInfoHelper) LtPrivileges() string {
	return "privileges < ?"
}

func (cols UserInfoHelper) LtPrivilegesWT() string {
	return "user_info.privileges < ?"
}

func (cols UserInfoHelper) LtePrivileges() string {
	return "privileges <= ?"
}

func (cols UserInfoHelper) LtePrivilegesWT() string {
	return "user_info.privileges <= ?"
}

func (cols UserInfoHelper) UserId() string {
	return "user_id"
}

func (cols UserInfoHelper) UserIdWT() string {
	return "user_info.user_id"
}

func (cols UserInfoHelper) EqUserId() string {
	return "user_id = ?"
}

func (cols UserInfoHelper) EqUserIdWT() string {
	return "user_info.user_id = ?"
}

func (cols UserInfoHelper) InUserId() string {
	return "user_id in ?"
}

func (cols UserInfoHelper) InUserIdWT() string {
	return "user_info.user_id in ?"
}

func (cols UserInfoHelper) NeUserId() string {
	return "user_id != ?"
}

func (cols UserInfoHelper) NeUserIdWT() string {
	return "user_info.user_id != ?"
}

func (cols UserInfoHelper) GtUserId() string {
	return "user_id > ?"
}

func (cols UserInfoHelper) GtUserIdWT() string {
	return "user_info.user_id > ?"
}

func (cols UserInfoHelper) GteUserId() string {
	return "user_id >= ?"
}

func (cols UserInfoHelper) GteUserIdWT() string {
	return "user_info.user_id >= ?"
}

func (cols UserInfoHelper) LtUserId() string {
	return "user_id < ?"
}

func (cols UserInfoHelper) LtUserIdWT() string {
	return "user_info.user_id < ?"
}

func (cols UserInfoHelper) LteUserId() string {
	return "user_id <= ?"
}

func (cols UserInfoHelper) LteUserIdWT() string {
	return "user_info.user_id <= ?"
}

func (cols UserInfoHelper) TableName() string {
	return "user_info"
}

// OUserInfo .
var OUserInfo = UserInfoHelper{}
