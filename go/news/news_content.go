package news

type NewsContent struct {
	Content string `json:"content" xorm:"MEDIUMTEXT"`
	Id      int64  `json:"id" xorm:"pk BIGINT"`
}

type NewsContentHelper struct {
}

func (cols NewsContentHelper) Content() string {
	return "content"
}

func (cols NewsContentHelper) ContentWT() string {
	return "news_content.content"
}

func (cols NewsContentHelper) EqContent() string {
	return "content = ?"
}

func (cols NewsContentHelper) EqContentWT() string {
	return "news_content.content = ?"
}

func (cols NewsContentHelper) InContent() string {
	return "content in ?"
}

func (cols NewsContentHelper) InContentWT() string {
	return "news_content.content in ?"
}

func (cols NewsContentHelper) NeContent() string {
	return "content != ?"
}

func (cols NewsContentHelper) NeContentWT() string {
	return "news_content.content != ?"
}

func (cols NewsContentHelper) GtContent() string {
	return "content > ?"
}

func (cols NewsContentHelper) GtContentWT() string {
	return "news_content.content > ?"
}

func (cols NewsContentHelper) GteContent() string {
	return "content >= ?"
}

func (cols NewsContentHelper) GteContentWT() string {
	return "news_content.content >= ?"
}

func (cols NewsContentHelper) LtContent() string {
	return "content < ?"
}

func (cols NewsContentHelper) LtContentWT() string {
	return "news_content.content < ?"
}

func (cols NewsContentHelper) LteContent() string {
	return "content <= ?"
}

func (cols NewsContentHelper) LteContentWT() string {
	return "news_content.content <= ?"
}

func (cols NewsContentHelper) Id() string {
	return "id"
}

func (cols NewsContentHelper) IdWT() string {
	return "news_content.id"
}

func (cols NewsContentHelper) EqId() string {
	return "id = ?"
}

func (cols NewsContentHelper) EqIdWT() string {
	return "news_content.id = ?"
}

func (cols NewsContentHelper) InId() string {
	return "id in ?"
}

func (cols NewsContentHelper) InIdWT() string {
	return "news_content.id in ?"
}

func (cols NewsContentHelper) NeId() string {
	return "id != ?"
}

func (cols NewsContentHelper) NeIdWT() string {
	return "news_content.id != ?"
}

func (cols NewsContentHelper) GtId() string {
	return "id > ?"
}

func (cols NewsContentHelper) GtIdWT() string {
	return "news_content.id > ?"
}

func (cols NewsContentHelper) GteId() string {
	return "id >= ?"
}

func (cols NewsContentHelper) GteIdWT() string {
	return "news_content.id >= ?"
}

func (cols NewsContentHelper) LtId() string {
	return "id < ?"
}

func (cols NewsContentHelper) LtIdWT() string {
	return "news_content.id < ?"
}

func (cols NewsContentHelper) LteId() string {
	return "id <= ?"
}

func (cols NewsContentHelper) LteIdWT() string {
	return "news_content.id <= ?"
}

func (cols NewsContentHelper) TableName() string {
	return "news_content"
}

// ONewsContent .
var ONewsContent = NewsContentHelper{}
