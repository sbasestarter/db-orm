package news

import (
	"time"
)

type NewsInfo struct {
	CategoryId  int       `json:"category_id" xorm:"not null INT"`
	CreatedAt   time.Time `json:"created_at" xorm:"not null created DATETIME"`
	Id          int64     `json:"id" xorm:"pk autoincr BIGINT"`
	ImagesJson  string    `json:"images_json" xorm:"not null TEXT"`
	PublishedAt time.Time `json:"published_at" xorm:"not null DATETIME"`
	Summary     string    `json:"summary" xorm:"not null TEXT"`
	Title       string    `json:"title" xorm:"not null VARCHAR(512)"`
}

type NewsInfoHelper struct {
}

func (cols NewsInfoHelper) CategoryId() string {
	return "category_id"
}

func (cols NewsInfoHelper) CategoryIdWT() string {
	return "news_info.category_id"
}

func (cols NewsInfoHelper) EqCategoryId() string {
	return "category_id = ?"
}

func (cols NewsInfoHelper) EqCategoryIdWT() string {
	return "news_info.category_id = ?"
}

func (cols NewsInfoHelper) InCategoryId() string {
	return "category_id in ?"
}

func (cols NewsInfoHelper) InCategoryIdWT() string {
	return "news_info.category_id in ?"
}

func (cols NewsInfoHelper) NeCategoryId() string {
	return "category_id != ?"
}

func (cols NewsInfoHelper) NeCategoryIdWT() string {
	return "news_info.category_id != ?"
}

func (cols NewsInfoHelper) GtCategoryId() string {
	return "category_id > ?"
}

func (cols NewsInfoHelper) GtCategoryIdWT() string {
	return "news_info.category_id > ?"
}

func (cols NewsInfoHelper) GteCategoryId() string {
	return "category_id >= ?"
}

func (cols NewsInfoHelper) GteCategoryIdWT() string {
	return "news_info.category_id >= ?"
}

func (cols NewsInfoHelper) LtCategoryId() string {
	return "category_id < ?"
}

func (cols NewsInfoHelper) LtCategoryIdWT() string {
	return "news_info.category_id < ?"
}

func (cols NewsInfoHelper) LteCategoryId() string {
	return "category_id <= ?"
}

func (cols NewsInfoHelper) LteCategoryIdWT() string {
	return "news_info.category_id <= ?"
}

func (cols NewsInfoHelper) CreatedAt() string {
	return "created_at"
}

func (cols NewsInfoHelper) CreatedAtWT() string {
	return "news_info.created_at"
}

func (cols NewsInfoHelper) EqCreatedAt() string {
	return "created_at = ?"
}

func (cols NewsInfoHelper) EqCreatedAtWT() string {
	return "news_info.created_at = ?"
}

func (cols NewsInfoHelper) InCreatedAt() string {
	return "created_at in ?"
}

func (cols NewsInfoHelper) InCreatedAtWT() string {
	return "news_info.created_at in ?"
}

func (cols NewsInfoHelper) NeCreatedAt() string {
	return "created_at != ?"
}

func (cols NewsInfoHelper) NeCreatedAtWT() string {
	return "news_info.created_at != ?"
}

func (cols NewsInfoHelper) GtCreatedAt() string {
	return "created_at > ?"
}

func (cols NewsInfoHelper) GtCreatedAtWT() string {
	return "news_info.created_at > ?"
}

func (cols NewsInfoHelper) GteCreatedAt() string {
	return "created_at >= ?"
}

func (cols NewsInfoHelper) GteCreatedAtWT() string {
	return "news_info.created_at >= ?"
}

func (cols NewsInfoHelper) LtCreatedAt() string {
	return "created_at < ?"
}

func (cols NewsInfoHelper) LtCreatedAtWT() string {
	return "news_info.created_at < ?"
}

func (cols NewsInfoHelper) LteCreatedAt() string {
	return "created_at <= ?"
}

func (cols NewsInfoHelper) LteCreatedAtWT() string {
	return "news_info.created_at <= ?"
}

func (cols NewsInfoHelper) Id() string {
	return "id"
}

func (cols NewsInfoHelper) IdWT() string {
	return "news_info.id"
}

func (cols NewsInfoHelper) EqId() string {
	return "id = ?"
}

func (cols NewsInfoHelper) EqIdWT() string {
	return "news_info.id = ?"
}

func (cols NewsInfoHelper) InId() string {
	return "id in ?"
}

func (cols NewsInfoHelper) InIdWT() string {
	return "news_info.id in ?"
}

func (cols NewsInfoHelper) NeId() string {
	return "id != ?"
}

func (cols NewsInfoHelper) NeIdWT() string {
	return "news_info.id != ?"
}

func (cols NewsInfoHelper) GtId() string {
	return "id > ?"
}

func (cols NewsInfoHelper) GtIdWT() string {
	return "news_info.id > ?"
}

func (cols NewsInfoHelper) GteId() string {
	return "id >= ?"
}

func (cols NewsInfoHelper) GteIdWT() string {
	return "news_info.id >= ?"
}

func (cols NewsInfoHelper) LtId() string {
	return "id < ?"
}

func (cols NewsInfoHelper) LtIdWT() string {
	return "news_info.id < ?"
}

func (cols NewsInfoHelper) LteId() string {
	return "id <= ?"
}

func (cols NewsInfoHelper) LteIdWT() string {
	return "news_info.id <= ?"
}

func (cols NewsInfoHelper) ImagesJson() string {
	return "images_json"
}

func (cols NewsInfoHelper) ImagesJsonWT() string {
	return "news_info.images_json"
}

func (cols NewsInfoHelper) EqImagesJson() string {
	return "images_json = ?"
}

func (cols NewsInfoHelper) EqImagesJsonWT() string {
	return "news_info.images_json = ?"
}

func (cols NewsInfoHelper) InImagesJson() string {
	return "images_json in ?"
}

func (cols NewsInfoHelper) InImagesJsonWT() string {
	return "news_info.images_json in ?"
}

func (cols NewsInfoHelper) NeImagesJson() string {
	return "images_json != ?"
}

func (cols NewsInfoHelper) NeImagesJsonWT() string {
	return "news_info.images_json != ?"
}

func (cols NewsInfoHelper) GtImagesJson() string {
	return "images_json > ?"
}

func (cols NewsInfoHelper) GtImagesJsonWT() string {
	return "news_info.images_json > ?"
}

func (cols NewsInfoHelper) GteImagesJson() string {
	return "images_json >= ?"
}

func (cols NewsInfoHelper) GteImagesJsonWT() string {
	return "news_info.images_json >= ?"
}

func (cols NewsInfoHelper) LtImagesJson() string {
	return "images_json < ?"
}

func (cols NewsInfoHelper) LtImagesJsonWT() string {
	return "news_info.images_json < ?"
}

func (cols NewsInfoHelper) LteImagesJson() string {
	return "images_json <= ?"
}

func (cols NewsInfoHelper) LteImagesJsonWT() string {
	return "news_info.images_json <= ?"
}

func (cols NewsInfoHelper) PublishedAt() string {
	return "published_at"
}

func (cols NewsInfoHelper) PublishedAtWT() string {
	return "news_info.published_at"
}

func (cols NewsInfoHelper) EqPublishedAt() string {
	return "published_at = ?"
}

func (cols NewsInfoHelper) EqPublishedAtWT() string {
	return "news_info.published_at = ?"
}

func (cols NewsInfoHelper) InPublishedAt() string {
	return "published_at in ?"
}

func (cols NewsInfoHelper) InPublishedAtWT() string {
	return "news_info.published_at in ?"
}

func (cols NewsInfoHelper) NePublishedAt() string {
	return "published_at != ?"
}

func (cols NewsInfoHelper) NePublishedAtWT() string {
	return "news_info.published_at != ?"
}

func (cols NewsInfoHelper) GtPublishedAt() string {
	return "published_at > ?"
}

func (cols NewsInfoHelper) GtPublishedAtWT() string {
	return "news_info.published_at > ?"
}

func (cols NewsInfoHelper) GtePublishedAt() string {
	return "published_at >= ?"
}

func (cols NewsInfoHelper) GtePublishedAtWT() string {
	return "news_info.published_at >= ?"
}

func (cols NewsInfoHelper) LtPublishedAt() string {
	return "published_at < ?"
}

func (cols NewsInfoHelper) LtPublishedAtWT() string {
	return "news_info.published_at < ?"
}

func (cols NewsInfoHelper) LtePublishedAt() string {
	return "published_at <= ?"
}

func (cols NewsInfoHelper) LtePublishedAtWT() string {
	return "news_info.published_at <= ?"
}

func (cols NewsInfoHelper) Summary() string {
	return "summary"
}

func (cols NewsInfoHelper) SummaryWT() string {
	return "news_info.summary"
}

func (cols NewsInfoHelper) EqSummary() string {
	return "summary = ?"
}

func (cols NewsInfoHelper) EqSummaryWT() string {
	return "news_info.summary = ?"
}

func (cols NewsInfoHelper) InSummary() string {
	return "summary in ?"
}

func (cols NewsInfoHelper) InSummaryWT() string {
	return "news_info.summary in ?"
}

func (cols NewsInfoHelper) NeSummary() string {
	return "summary != ?"
}

func (cols NewsInfoHelper) NeSummaryWT() string {
	return "news_info.summary != ?"
}

func (cols NewsInfoHelper) GtSummary() string {
	return "summary > ?"
}

func (cols NewsInfoHelper) GtSummaryWT() string {
	return "news_info.summary > ?"
}

func (cols NewsInfoHelper) GteSummary() string {
	return "summary >= ?"
}

func (cols NewsInfoHelper) GteSummaryWT() string {
	return "news_info.summary >= ?"
}

func (cols NewsInfoHelper) LtSummary() string {
	return "summary < ?"
}

func (cols NewsInfoHelper) LtSummaryWT() string {
	return "news_info.summary < ?"
}

func (cols NewsInfoHelper) LteSummary() string {
	return "summary <= ?"
}

func (cols NewsInfoHelper) LteSummaryWT() string {
	return "news_info.summary <= ?"
}

func (cols NewsInfoHelper) Title() string {
	return "title"
}

func (cols NewsInfoHelper) TitleWT() string {
	return "news_info.title"
}

func (cols NewsInfoHelper) EqTitle() string {
	return "title = ?"
}

func (cols NewsInfoHelper) EqTitleWT() string {
	return "news_info.title = ?"
}

func (cols NewsInfoHelper) InTitle() string {
	return "title in ?"
}

func (cols NewsInfoHelper) InTitleWT() string {
	return "news_info.title in ?"
}

func (cols NewsInfoHelper) NeTitle() string {
	return "title != ?"
}

func (cols NewsInfoHelper) NeTitleWT() string {
	return "news_info.title != ?"
}

func (cols NewsInfoHelper) GtTitle() string {
	return "title > ?"
}

func (cols NewsInfoHelper) GtTitleWT() string {
	return "news_info.title > ?"
}

func (cols NewsInfoHelper) GteTitle() string {
	return "title >= ?"
}

func (cols NewsInfoHelper) GteTitleWT() string {
	return "news_info.title >= ?"
}

func (cols NewsInfoHelper) LtTitle() string {
	return "title < ?"
}

func (cols NewsInfoHelper) LtTitleWT() string {
	return "news_info.title < ?"
}

func (cols NewsInfoHelper) LteTitle() string {
	return "title <= ?"
}

func (cols NewsInfoHelper) LteTitleWT() string {
	return "news_info.title <= ?"
}

func (cols NewsInfoHelper) TableName() string {
	return "news_info"
}

// ONewsInfo .
var ONewsInfo = NewsInfoHelper{}
