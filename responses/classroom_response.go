package responses

//type ClassRoomResponse struct {
//	ID            uint   `json:"id" gorm:"primaryKey"`
//	ClassName     string `json:"class_name"`
//	Major         string `json:"major"`
//	ClassYear     string `json:"class_year"`
//	SubjectName   string `json:"subject_name"`
//	ClassRoomCode string `json:"class_room_code"`
//	CreatedAt     string `json:"created_at"`
//	UpdatedAt     string `json:"updated_at"`
//}

type MessageClassRoomResponse struct {
	Message string `json:"message"`
}

type ClassroomResponse struct {
	ID        uint   `json:"id"`
	ClassName string `json:"class_name"`
}
