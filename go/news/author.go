package news

type Author struct {
	Id       int64  `json:"id" xorm:"pk BIGINT"`
	NickName string `json:"nick_name" xorm:"VARCHAR(255)"`
	Avatar   string `json:"avatar" xorm:"VARCHAR(1024)"`
}

type AuthorHelper struct {
}

func (cols AuthorHelper) Id() string {
	return "id"
}

func (cols AuthorHelper) IdWT() string {
	return "author.id"
}

func (cols AuthorHelper) EqId() string {
	return "id = ?"
}

func (cols AuthorHelper) EqIdWT() string {
	return "author.id = ?"
}

func (cols AuthorHelper) NeId() string {
	return "id != ?"
}

func (cols AuthorHelper) NeIdWT() string {
	return "author.id != ?"
}

func (cols AuthorHelper) GtId() string {
	return "id > ?"
}

func (cols AuthorHelper) GtIdWT() string {
	return "author.id > ?"
}

func (cols AuthorHelper) GteId() string {
	return "id >= ?"
}

func (cols AuthorHelper) GteIdWT() string {
	return "author.id >= ?"
}

func (cols AuthorHelper) LtId() string {
	return "id < ?"
}

func (cols AuthorHelper) LtIdWT() string {
	return "author.id < ?"
}

func (cols AuthorHelper) LteId() string {
	return "id <= ?"
}

func (cols AuthorHelper) LteIdWT() string {
	return "author.id <= ?"
}

func (cols AuthorHelper) NickName() string {
	return "nick_name"
}

func (cols AuthorHelper) NickNameWT() string {
	return "author.nick_name"
}

func (cols AuthorHelper) EqNickName() string {
	return "nick_name = ?"
}

func (cols AuthorHelper) EqNickNameWT() string {
	return "author.nick_name = ?"
}

func (cols AuthorHelper) NeNickName() string {
	return "nick_name != ?"
}

func (cols AuthorHelper) NeNickNameWT() string {
	return "author.nick_name != ?"
}

func (cols AuthorHelper) GtNickName() string {
	return "nick_name > ?"
}

func (cols AuthorHelper) GtNickNameWT() string {
	return "author.nick_name > ?"
}

func (cols AuthorHelper) GteNickName() string {
	return "nick_name >= ?"
}

func (cols AuthorHelper) GteNickNameWT() string {
	return "author.nick_name >= ?"
}

func (cols AuthorHelper) LtNickName() string {
	return "nick_name < ?"
}

func (cols AuthorHelper) LtNickNameWT() string {
	return "author.nick_name < ?"
}

func (cols AuthorHelper) LteNickName() string {
	return "nick_name <= ?"
}

func (cols AuthorHelper) LteNickNameWT() string {
	return "author.nick_name <= ?"
}

func (cols AuthorHelper) Avatar() string {
	return "avatar"
}

func (cols AuthorHelper) AvatarWT() string {
	return "author.avatar"
}

func (cols AuthorHelper) EqAvatar() string {
	return "avatar = ?"
}

func (cols AuthorHelper) EqAvatarWT() string {
	return "author.avatar = ?"
}

func (cols AuthorHelper) NeAvatar() string {
	return "avatar != ?"
}

func (cols AuthorHelper) NeAvatarWT() string {
	return "author.avatar != ?"
}

func (cols AuthorHelper) GtAvatar() string {
	return "avatar > ?"
}

func (cols AuthorHelper) GtAvatarWT() string {
	return "author.avatar > ?"
}

func (cols AuthorHelper) GteAvatar() string {
	return "avatar >= ?"
}

func (cols AuthorHelper) GteAvatarWT() string {
	return "author.avatar >= ?"
}

func (cols AuthorHelper) LtAvatar() string {
	return "avatar < ?"
}

func (cols AuthorHelper) LtAvatarWT() string {
	return "author.avatar < ?"
}

func (cols AuthorHelper) LteAvatar() string {
	return "avatar <= ?"
}

func (cols AuthorHelper) LteAvatarWT() string {
	return "author.avatar <= ?"
}

func (cols AuthorHelper) TableName() string {
	return "author"
}

// OAuthor .
var OAuthor = AuthorHelper{}
