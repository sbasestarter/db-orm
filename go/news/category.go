package news

type Category struct {
	Id   int    `json:"id" xorm:"not null pk INT"`
	Name string `json:"name" xorm:"not null VARCHAR(255)"`
}

type CategoryHelper struct {
}

func (cols CategoryHelper) Id() string {
	return "id"
}

func (cols CategoryHelper) IdWT() string {
	return "category.id"
}

func (cols CategoryHelper) EqId() string {
	return "id = ?"
}

func (cols CategoryHelper) EqIdWT() string {
	return "category.id = ?"
}

func (cols CategoryHelper) InId() string {
	return "id in ?"
}

func (cols CategoryHelper) InIdWT() string {
	return "category.id in ?"
}

func (cols CategoryHelper) NeId() string {
	return "id != ?"
}

func (cols CategoryHelper) NeIdWT() string {
	return "category.id != ?"
}

func (cols CategoryHelper) GtId() string {
	return "id > ?"
}

func (cols CategoryHelper) GtIdWT() string {
	return "category.id > ?"
}

func (cols CategoryHelper) GteId() string {
	return "id >= ?"
}

func (cols CategoryHelper) GteIdWT() string {
	return "category.id >= ?"
}

func (cols CategoryHelper) LtId() string {
	return "id < ?"
}

func (cols CategoryHelper) LtIdWT() string {
	return "category.id < ?"
}

func (cols CategoryHelper) LteId() string {
	return "id <= ?"
}

func (cols CategoryHelper) LteIdWT() string {
	return "category.id <= ?"
}

func (cols CategoryHelper) Name() string {
	return "name"
}

func (cols CategoryHelper) NameWT() string {
	return "category.name"
}

func (cols CategoryHelper) EqName() string {
	return "name = ?"
}

func (cols CategoryHelper) EqNameWT() string {
	return "category.name = ?"
}

func (cols CategoryHelper) InName() string {
	return "name in ?"
}

func (cols CategoryHelper) InNameWT() string {
	return "category.name in ?"
}

func (cols CategoryHelper) NeName() string {
	return "name != ?"
}

func (cols CategoryHelper) NeNameWT() string {
	return "category.name != ?"
}

func (cols CategoryHelper) GtName() string {
	return "name > ?"
}

func (cols CategoryHelper) GtNameWT() string {
	return "category.name > ?"
}

func (cols CategoryHelper) GteName() string {
	return "name >= ?"
}

func (cols CategoryHelper) GteNameWT() string {
	return "category.name >= ?"
}

func (cols CategoryHelper) LtName() string {
	return "name < ?"
}

func (cols CategoryHelper) LtNameWT() string {
	return "category.name < ?"
}

func (cols CategoryHelper) LteName() string {
	return "name <= ?"
}

func (cols CategoryHelper) LteNameWT() string {
	return "category.name <= ?"
}

func (cols CategoryHelper) TableName() string {
	return "category"
}

// OCategory .
var OCategory = CategoryHelper{}
