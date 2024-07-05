package requests

type GetUserClassroomRequest struct {
	ID       uint   `json:"id" validate:"required"`
	UserType string `json:"user_type" validate:"required"`
}

type UserClassroomRequest struct {
	ClassroomID uint   `json:"classroom_id" validate:"required"`
	SubjectID   uint   `json:"subject_id" validate:"required"`
	UserType    string `json:"user_type"`
}

type TeacherIdRequest struct {
	UserID   int    `json:"user_id"`
	UserType string `json:"user_type"`
}
type CreateUserClassroomRequest struct {
	ClassroomID uint   `json:"classroom_id"`
	UserIDs     []uint `json:"user_ids"`
}

type CreateUserClassroomANDUserBehaviorRoomRequest struct {
	ClassroomID uint   `json:"classroom_id"`
	SubjectID   uint   `json:"subject_id"`
	UserIDs     []uint `json:"user_ids"`
}

type UpdateUserClassRoomRequest struct {
	UserID      uint `json:"User_id"`
	ClassRoomID uint `json:"class_room_id"`
}

type UserClassRoomByIDRequest struct {
	ClassRoomID uint `json:"class_room_id"`
}


type UserClassRoomByUserIDRequest struct {
	UserID      uint `json:"User_id"`
}

type DeleteUserClassRoomByUserIDRequest struct {
	Id uint `json:"id" validate:"required"`
}
