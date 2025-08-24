package part3

import (
	"gorm.io/gorm"
)

// GetUserWithPostsAndComments 查询指定用户的所有文章及其评论
func GetUserWithPostsAndComments(db *gorm.DB, userID uint) (*User, error) {
	var user User
	// 嵌套预加载：加载用户→文章→评论（避免N+1查询）
	err := db.Preload("Posts.Comments").First(&user, userID).Error
	return &user, err
}

// GetPostWithMostComments 查询评论数量最多的文章信息
func GetPostWithMostComments(db *gorm.DB) (*Post, error) {
	// 步骤1：子查询统计每篇文章的评论数，按评论数降序取第一条
	var commentCountResult struct {
		PostID     uint
		CommentNum int64
	}
	subQuery := db.Model(&Comment{}).
		Select("post_id, COUNT(*) as comment_num").
		Group("post_id").
		Order("comment_num DESC").
		Limit(1)

	if err := db.Raw("?", subQuery).Scan(&commentCountResult).Error; err != nil {
		return nil, err
	}

	// 步骤2：根据文章ID查询完整信息（含作者、评论）
	var topPost Post
	err := db.Preload("User").Preload("Comments").First(&topPost, commentCountResult.PostID).Error
	return &topPost, err
}
