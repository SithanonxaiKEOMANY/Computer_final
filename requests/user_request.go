package requests

//CRUD RestAPI Request Data for Create, Update, Delete

type UserWithPaginationRequest struct {
	PerPage     int    `json:"per_page"`
	CurrentPage int    `json:"current_page"`
	Sorting     string `json:"sorting"`
	UserType    string `json:"user_type"`
}

type ClassroomIDRequest struct {
	ClassroomID int `json:"classroom_id" validate:"required"`
}

type ResetPasswordRequest struct {
	Phone       string `json:"phone" validate:"required"`
	NewPassword string `json:"new_password"`
}

type SigUpUserRequest struct {
	Phone     string `json:"phone" validate:"required,min=9,max=10"`
	Password  string `json:"password" validate:"required"`
	UserType  string `json:"user_type" validate:"required"`
	CodeID    string `json:"code_id" validate:"required"`
	Firstname string `json:"firstname" validate:"required"`
	Lastname  string `json:"lastname" validate:"required"`
	Gender    string `json:"gender" validate:"required"`
	Degree    string `json:"degree"`
}

type SignInUserRequest struct {
	Phone    string `json:"phone" validate:"required,min=9,max=10"`
	Password string `json:"password" validate:"required"`
	// UserType string `json:"user_type" validate:"required"`
	//Token    string `json:"token"`
}

type CreateUserRequest struct {
	Phone  string `json:"phone" validate:"required,min=9,max=10"`
	CodeID string `json:"code_id" validate:"required"`
}

type UpdateUserRequest struct {
	Id     int    `json:"id" validate:"required"`
	Phone  string `json:"phone" validate:"required,min=9,max=10"`
	CodeID string `json:"code_id" validate:"required"`
}

type DeleteUserRequest struct {
	Id     int    `json:"id" validate:"required"`
	Phone  string `json:"phone" validate:"required,min=9,max=10"`
	CodeID string `json:"code_id" validate:"required"`
}

type UserIdRequest struct {
	ID uint `json:"id" validate:"required"`
}

type UserCodeIdRequest struct {
	CodeID string `json:"code_id"`
}

type UserPhoneRequest struct {
	Phone string `json:"phone"`
}

type UserRequest struct {
	CodeID    string `json:"code_id" validate:"required"`
	Firstname string `json:"firstname" `
	Lastname  string `json:"lastname"`
	Password  string `json:"password"`
	Phone     string `json:"phone" validate:"required,min=9,max=10"`
	Gender    string `json:"gender"`
	Degree    string `json:"degree"`
	Skill     string `json:"skill"`
	UserType  string `json:"user_type"`
}

type UserImageRequest struct {
	StudentID string `json:"student_id" validate:"required"`
	Image     []byte `json:"image" validate:"required"`
}

type UserTypeRequest struct {
	UserType string `json:"user_type"`
}
