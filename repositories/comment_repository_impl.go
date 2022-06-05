package repositories

import (
	"errors"
	"final/models"

	"gorm.io/gorm"
)

type CommentRepositoryImpl struct {
}

func NewCommentRepository() CommentRepository {
	return &CommentRepositoryImpl{}
}

func (commentRepository *CommentRepositoryImpl) CreateComment(db *gorm.DB, comment models.Comment) (models.Comment, error) {
	err := db.Create(&comment).Error
	if err != nil {
		return comment, errors.New(err.Error())
	}

	commentCreated := models.Comment{
		ID:        comment.ID,
		Message:   comment.Message,
		PhotoID:   comment.PhotoID,
		UserID:    comment.UserID,
		CreatedAt: comment.CreatedAt,
	}

	return commentCreated, nil
}
func (commentRepository *CommentRepositoryImpl) GetComments(db *gorm.DB) ([]models.Comment, error) {
	comments := []models.Comment{}

	result := db.Table("comments").Scan(&comments)
	if result.RowsAffected == 0 {
		return comments, errors.New("comments not found")
	}

	for i, comment := range comments {
		user := models.User{}
		db.Table("users").Select([]string{"id", "email", "username"}).Where("id = ?", comment.UserID).Scan(&user)
		comments[i].User = &user

		photo := models.Photo{}
		db.Table("photos").Select([]string{"id", "title", "caption", "photo_url", "user_id"}).Where("id = ?", comment.PhotoID).Scan(&photo)
		comments[i].Photo = &photo
	}

	return comments, nil
}
func (commentRepository *CommentRepositoryImpl) UpdateComment(db *gorm.DB, comment models.Comment, commentId int) (models.Comment, error) {
	panic("implement me")
}
func (commentRepository *CommentRepositoryImpl) DeleteComment(db *gorm.DB, comment models.Comment) error {
	panic("implement me")
}