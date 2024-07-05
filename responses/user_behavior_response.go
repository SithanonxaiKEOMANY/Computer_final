package responses

type UserBehaviorPaginatedResponse struct {
	TotalPages  int                    `json:"total_pages"`
	PerPage     int                    `json:"per_page"`
	CurrentPage int                    `json:"current_page"`
	Sorting     string                 `json:"sorting"`
	Items       []UserBehaviorResponse `json:"items"`
}

type UserBehaviorResponse struct {
	ID              uint `json:"id" gorm:"primaryKey"`
	UserID          uint `json:"user_id"`
	UserClassroomID uint `json:"user_classroom_id"`
	//StudentCheck     bool `json:"student_check"`
	//StudentAbsent    bool `json:"student_absent"`
	//StudentVacation  bool `json:"student_vacation"`
	//StudentBreakRule bool `json:"student_break_rule"`
	CountCheck     int `json:"count_check"`
	CountAbsent    int `json:"count_absent"`
	CountVacation  int `json:"count_vacation"`
	CountBreakRule int `json:"count_break_rule"`
}

type MessageUserBehaviorResponse struct {
	Message string `json:"message"`
	Status  bool   `json:"status"`
}
