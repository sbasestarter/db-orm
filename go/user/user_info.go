package user

import (
	"time"
)

type UserInfo struct {
	Avatar   string    `json:"avatar" xorm:"'avatar' not null comment('头像') TEXT"`
	CreateAt time.Time `json:"create_at" xorm:"'create_at' not null comment('创建时间') DATETIME"`
	NickName string    `json:"nick_name" xorm:"'nick_name' not null comment('昵称') unique VARCHAR(100)"`
	UserId   int64     `json:"user_id" xorm:"'user_id' not null pk autoincr comment('用户ID') BIGINT"`
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
