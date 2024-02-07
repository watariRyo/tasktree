package model

type BaseTask struct {
	ID        int64
	UserID    int
	Title     *string
	Content   *string
	DayOfWeek string
}
