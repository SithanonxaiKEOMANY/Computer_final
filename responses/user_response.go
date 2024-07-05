package responses

type UserPaginatedResponse struct {
	TotalPages  int            `json:"total_pages"`
	PerPage     int            `json:"per_page"`
	CurrentPage int            `json:"current_page"`
	Sorting     string         `json:"sorting"`
	Items       []UserResponse `json:"items"`
}

type UserResponse struct {
	ID        uint   `json:"id"`
	CodeID    string `json:"code_id"`
	Firstname string `json:"firstname" `
	Lastname  string `json:"lastname"`
	Phone     string `json:"phone"`
	Gender    string `json:"gender"`
	Degree    string `json:"degree"`
	Skill     string `json:"skill"`
	UserType  string `json:"user_type"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type UserClassroomResponse struct {
	ID          uint   `json:"id"`
	ClassName   int    `json:"class_name"`
	Major       string `json:"major"`
	ClassYear   int    `json:"class_year"`
	SubjectName string `json:"subject_name"`
}

type User struct {
	ID        uint   `json:"id"`
	CodeID    string `json:"code_id"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

type MessageUserResponse struct {
	Message string `json:"message"`
}

type SignUpResponse struct {
	Message  string `json:"message"`
	Phone    string `json:"phone"`
	ID       uint   `json:"id"`
	UserType string `json:"user_type"`
	//AccessToken string `json:"access_token"`
}

type SignInResponse struct {
	Message  string `json:"message"`
	Phone    string `json:"phone"`
	ID       uint   `json:"id"`
	UserType string `json:"user_type"`
	//AccessToken string `json:"access_token"`
}

type CountSubjectsResponse struct {
	TotalSubjects int64 `json:"total_subjects"`
}

type CountTeachersResponse struct {
	TotalTeachers int64 `json:"total_teachers"`
}
type CountStudentResponse struct {
	TotalStudents int64 `json:"total_students"`
}
