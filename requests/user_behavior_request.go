package requests

type CreateUserBehaviorRequest struct {
	UserClassroomID uint   `json:"user_classroom_id"`
	UserIDs         []uint `json:"user_ids"`
}
type StudentBehaviorRequest struct {
	UserClassroomID uint   `json:"user_classroom_id"`
	UserIDs         []uint `json:"user_ids"`
}

type UserBehaviorWithPaginationRequest struct {
	PerPage     int    `json:"per_page"`
	CurrentPage int    `json:"current_page"`
	Sorting     string `json:"sorting"`
	UserType    string `json:"user_type"`
	ClassroomID uint   `json:"classroom_id"`
	SubjectID   uint   `json:"subject_id"`
}

type UserBehaviorRequest struct {
	UserClassroomID uint           `json:"user_classroom_id"`
	UserBehaviors   []UserBehavior `json:"user_behaviors" validate:"required"`
}

type UserBehavior struct {
	UserID           uint `json:"user_id" validate:"required"`
	StudentCheck     bool `json:"student_check" validate:"required"`
	StudentAbsent    bool `json:"student_absent" validate:"required"`
	StudentVacation  bool `json:"student_vacation" validate:"required"`
	StudentBreakRule bool `json:"student_break_rule" validate:"required"`
}

//type CreateUserBehaviorRoomRequest struct {
//	UserID      uint `json:"user_id"`
//	ClassRoomID uint `json:"class_room_id"`
//}
//
//type UpdateUserBehaviorRequest struct {
//	UserID      uint `json:"User_id"`
//	ClassRoomID uint `json:"class_room_id"`
//}
//type DeleteUserBehaviorRequest struct {
//	Id int `json:"id" validate:"required"`
//}
//
//type UserBehaviorIDRequest struct {
//	Id int `json:"id"`
//}
//type UserBehaviorClassRoomByIDRequest struct {
//	ClassRoomID uint `json:"class_room_id"`
//}
//type UserBehaviorClassRoomByUserIDRequest struct {
//	UserID uint `json:"User_id"`
//}
