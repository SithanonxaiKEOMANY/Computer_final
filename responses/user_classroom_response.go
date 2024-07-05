package responses

type GetUserClassroomResponse struct {
	ClassroomID   uint   `json:"classroom_id"`
	ClassroomName string `json:"classroom_name"`
	SubjectID     uint   `json:"subject_id"`
	SubjectName   string `json:"subject_name"`
}
type UserClassRoomResponse struct {
	ID          uint   `json:"id" gorm:"primaryKey"`
	UserID      uint   `json:"user_id"`
	ClassRoomID uint   `json:"class_room_id"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

type MessageUserClassRoomResponse struct {
	Message string `json:"message"`
	Status  bool   `json:"status"`
}

type StudentClassroomResponse struct {
	ID           uint   `json:"id"`
	ClassName    string `json:"class_name"`
	SubjectName  string `json:"subject_name"`
	UserStudents []UserStudents
}

type UserStudents struct {
	ID        uint   `json:"id"`
	CodeID    string `json:"code_id"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Gender    string `json:"gender"`
}
