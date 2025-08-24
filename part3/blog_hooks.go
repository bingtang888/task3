package part3

import (
	"gorm.io/gorm"
)

// User_hook 用户模型（带文章数量统计）
// 明确指定表名：user_hooks（与旧表 users 完全独立）
type User_hook struct {
	gorm.Model
	Username  string      `gorm:"type:varchar(64);not null;unique"` // 修复：仅保留唯一约束，无需重复索引标签
	Posts     []Post_hook `gorm:"foreignKey:UserID"`                // 关联新的 Post_hook 模型
	PostCount int         `gorm:"not null;default:0"`               // 文章数量统计
}

// Post_hook 文章模型（带评论状态）
// 明确指定表名：post_hooks（与旧表 posts 完全独立）
type Post_hook struct {
	gorm.Model
	Title         string          `gorm:"type:varchar(255);not null"`
	Content       string          `gorm:"type:text;not null"`
	UserID        uint            `gorm:"not null"`                                // 关联 User_hook 的 ID
	User          User_hook       `gorm:"foreignKey:UserID"`                       // 反向关联新的 User_hook 模型
	Comments      []Comment_hooks `gorm:"foreignKey:PostID"`                       // 关联 Comment 模型
	CommentStatus string          `gorm:"type:varchar(20);not null;default:'有评论'"` // 评论状态
}

// Comment 评论模型（复用，但明确关联 Post_hook）
// 明确指定表名：comment_hooks（与旧表 comments 完全独立）
type Comment_hooks struct {
	gorm.Model
	Content string    `gorm:"type:text;not null"`
	PostID  uint      `gorm:"not null"`          // 关联 Post_hook 的 ID
	Post    Post_hook `gorm:"foreignKey:PostID"` // 反向关联新的 Post_hook 模型
}

// 修复1：Post_hook 的 BeforeCreate 钩子（关联 User_hook 模型）
func (p *Post_hook) BeforeCreate(tx *gorm.DB) error {
	// 更新 User_hook 表的 post_count（而非旧的 User 表）
	return tx.Model(&User_hook{}).
		Where("id = ?", p.UserID).
		Update("post_count", gorm.Expr("post_count + 1")).Error
}

// 修复2：Comment 的 AfterDelete 钩子（关联 Post_hook 模型）
func (c *Comment_hooks) AfterDelete(tx *gorm.DB) error {
	// 1. 查询 Comment 表（新表 comment_hooks）中该文章的剩余评论数
	var count int64
	if err := tx.Model(&Comment_hooks{}).
		Where("post_id = ?", c.PostID).
		Count(&count).Error; err != nil {
		return err
	}

	// 2. 更新 Post_hook 表的 comment_status（而非旧的 Post 表）
	if count == 0 {
		return tx.Model(&Post_hook{}).
			Where("id = ?", c.PostID).
			Update("comment_status", "无评论").Error
	}
	return nil
}
