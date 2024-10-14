package app

import (
	"shopping-clone/database"
)

func CreateComment(comment *database.Comment) error {
	return database.CreateComment(comment)
}

func ReadComment(commentID string) (*database.Comment, error) {
	return database.ReadComment(commentID)
}

func UpdateComment(comment *database.Comment) error {
	return database.UpdateComment(comment)
}

func DeleteComment(commentID string) error {
	return database.DeleteComment(commentID)
}

func GetAllCommentsByCustomerID(customerID string) (*[]database.Comment, error) {
	return database.GetAllCommentsByCustomerID(customerID)
}
