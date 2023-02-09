package models

import "github.com/jinzhu/gorm"

type Article struct {
	Model

	TagID int `json:"tag_id" gorm:"index"`
	Tag   Tag `json:"tag"`

	Title         string `json:"title"`
	Desc          string `json:"desc"`
	Content       string `json:"content"`
	CreatedBy     string `json:"created_by"`
	ModifiedBy    string `json:"modified_by"`
	CoverImageUrl string `json:"cover_image_url"`
	State         int    `json:"state"`
}

//gorm:index，用于声明这个字段为索引，如果你使用了自动迁移功能则会有所影响，不使用则无影响
//Tag字段，实际是一个嵌套的struct，它利用TagID与Tag模型相互关联，在执行查询的时候，能够达到Article、Tag关联查询的功能

//func (article *Article) BeforeCreate(scope *gorm.Scope) error {
//	scope.SetColumn("CreatedOn", time.Now().Unix())
//
//	return nil
//}
//
//func (article *Article) BeforeUpdate(scope *gorm.Scope) error {
//	scope.SetColumn("ModifiedOn", time.Now().Unix())
//
//	return nil
//}

// ExistArticleByID checks if an article exists based on ID
func ExistArticleByID(id int) (bool, error) {
	var article Article
	//未被软删除
	err := db.Select("id").Where("id = ? AND deleted_on = ? ", id, 0).First(&article).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	if article.ID > 0 {
		return true, nil
	}

	return false, nil
}

//func ExistArticleByID(id int) bool {
//	var article Article
//	// select id where id = ? from article order by id limit 1;
//	db.Select("id").Where("id = ?", id).First(&article)
//
//	if article.ID > 0 {
//		return true
//	}
//
//	return false
//}

func GetArticleTotal(maps interface{}) (int, error) {
	var count int
	//select count(1) from article where maps( name = ... and id = ... and state = ...);
	//Count 用于获取匹配的记录数
	err := db.Model(&Article{}).Where(maps).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}

func GetArticles(pageNum int, pageSize int, maps interface{}) ([]*Article, error) {
	var articles []*Article
	//预加载 联表查询
	err := db.Preload("Tag").Where(maps).Offset(pageNum).Limit(pageSize).Find(&articles).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return articles, nil
}

func GetArticle(id int) (*Article, error) {
	var article Article
	err := db.Where("id = ? AND deleted_on = ? ", id, 0).First(&article).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	err = db.Model(&article).Related(&article.Tag).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return &article, nil
	//db.Where("id = ?", id).First(&article)
	//db.Model(&article).Related(&article.Tag)
	//
	//return
}

func EditArticle(id int, data interface{}) error {

	err := db.Model(&Article{}).Where("id = ?", id).Updates(data).Error
	if err != nil {
		return err
	}
	return nil
}

func AddArticle(data map[string]interface{}) error {
	err := db.Create(&Article{
		TagID:         data["tag_id"].(int),
		Title:         data["title"].(string),
		Desc:          data["desc"].(string),
		Content:       data["desc"].(string),
		CreatedBy:     data["created_by"].(string),
		State:         data["state"].(int),
		CoverImageUrl: data["cover_image_url"].(string),
	}).Error
	if err != nil {
		return err
	}
	return nil
}

func DeleteArticle(id int) error {
	err := db.Where("id = ?", id).Delete(Article{}).Error
	if err != nil {
		return err
	}
	return nil
}

func CleanAllArticle() error {
	err := db.Unscoped().Where("deleted_on != ? ", 0).Delete(&Article{}).Error
	if err != nil {
		return err
	}

	return nil
}
