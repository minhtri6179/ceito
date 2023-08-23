package models

type Answer struct {
	ID         uint   `json:"id" gorm:"primaryKey"`
	QuestionID uint   `json:"question_id"`
	Answer     string `json:"answer"`
	IsCorrect  bool   `json:"is_correct"`
	CreatedAt  int64  `json:"created_at"`
	UpdatedAt  int64  `json:"updated_at"`
}
