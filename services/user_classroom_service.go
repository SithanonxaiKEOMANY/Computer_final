package services

import (
	"fmt"
	"github.com/pkg/errors"
	"go_starter/models"
	"go_starter/repositories"
	"go_starter/requests"
	"go_starter/responses"
	"strconv"
)

type UserClassRoomService interface {
	GetUserClassroomsService(request requests.GetUserClassroomRequest) ([]responses.GetUserClassroomResponse, error)
	GetUserClassroomByStudentTypeService(request requests.UserClassroomRequest) (*responses.StudentClassroomResponse, error)
	CreateUserClassRoomAndUserBehaviorService(request requests.CreateUserClassroomANDUserBehaviorRoomRequest) (*responses.MessageUserClassRoomResponse, error)

	//CreateUserClassRoomService(request requests.CreateUserClassroomRequest) (*responses.MessageUserClassRoomResponse, error)
	//GetUserClassroomByStudentTypeService(request requests.UserClassroomRequest) (*responses.StudentClassroomResponse, error)

	GetClassroomByTeacherService(request requests.TeacherIdRequest) ([]responses.UserClassroomResponse, error)
	GetAllUserClassRoomServices() ([]responses.UserClassRoomResponse, error)
	GetByIdUserClassRoomService(id uint) (*responses.UserClassRoomResponse, error)
	GetByUserIDService(userID uint) (*responses.UserClassRoomResponse, error)
	GetByUserClassRoomIdService(classRoomID uint) (*responses.UserClassRoomResponse, error)

	UpdateUserClassRoomService(request requests.UpdateUserClassRoomRequest) (*responses.MessageUserClassRoomResponse, error)
	DeleteUserClassRoomService(request requests.DeleteUserClassRoomByUserIDRequest) (*responses.MessageUserClassRoomResponse, error)
}

type userClassRoomService struct {
	repositoryUserClassRoom repositories.UserClassRoomRepository
	userBehaviorService     UserBehaviorService
}

func formatClassName(classroom models.Classroom) string {
	classYearString := strconv.Itoa(classroom.ClassYear)
	classNameString := strconv.Itoa(classroom.ClassName)
	return classYearString + classroom.Major + classNameString
}

func (uc *userClassRoomService) GetUserClassroomsService(request requests.GetUserClassroomRequest) ([]responses.GetUserClassroomResponse, error) {
	userClassroomsData, err := uc.repositoryUserClassRoom.GetUserClassroomsRepository(request.ID, request.UserType)
	if err != nil {
		return nil, err
	}
	var response []responses.GetUserClassroomResponse
	for _, item := range userClassroomsData {
		responseData := responses.GetUserClassroomResponse{
			ClassroomID:   item.ClassroomID,
			ClassroomName: formatClassName(item.Classroom),
			SubjectID:     item.SubjectID,
			SubjectName:   item.Subject.SubjectName,
		}
		response = append(response, responseData)
	}
	return response, nil
}

//func (uc *userClassRoomService) createUserClassRooms(request requests.CreateUserClassroomANDUserBehaviorRoomRequest) error {
//	var model []models.UserClassroom
//	for _, userID := range request.UserIDs {
//		userClassroom := models.UserClassroom{
//			UserID:      userID,
//			ClassroomID: request.ClassroomID,
//			SubjectID:   request.SubjectID, // Assuming SubjectID is available in request
//		}
//		model = append(model, userClassroom)
//	}
//
//	if err := uc.repositoryUserClassRoom.CreateUserClassroomRepository(model); err != nil {
//		return err
//	}
//	return nil
//}
//
//func (uc *userClassRoomService) insertStudentBehaviors(request requests.CreateUserClassroomANDUserBehaviorRoomRequest) error {
//	behaviorRequest := requests.StudentBehaviorRequest{
//		UserClassroomID: request.UserClassroomID,
//		UserIDs:         request.UserIDs,
//	}
//	_ = uc.userBehaviorService.InsertStudentBehaviorByStudentIdAndClassroomIdService(behaviorRequest)
//
//	return nil
//}

func (uc *userClassRoomService) CreateUserClassRoomAndUserBehaviorService(request requests.CreateUserClassroomANDUserBehaviorRoomRequest) (*responses.MessageUserClassRoomResponse, error) {
	if len(request.UserIDs) == 0 {
		return nil, errors.New("no user IDs provided")
	}
	// Step 1: Check if there's already an existing association for any user in the request
	for _, userID := range request.UserIDs {
		exists, err := uc.repositoryUserClassRoom.CheckUserClassRoomExistsRepository(userID, request.ClassroomID, request.SubjectID)
		if err != nil {
			return nil, err
		}
		if exists {
			return nil, errors.New("user already assigned to this classroom")
		}
	}

	// Step 2: Create user-classroom associations
	var model []models.UserClassroom
	for _, userID := range request.UserIDs {
		userClassroom := models.UserClassroom{
			UserID:      userID,
			ClassroomID: request.ClassroomID,
			SubjectID:   request.SubjectID, // Assuming SubjectID is available in request
		}
		model = append(model, userClassroom)
	}

	if err := uc.repositoryUserClassRoom.CreateUserClassroomRepository(model); err != nil {
		return nil, err
	}

	// Step 3: Retrieve the user-classroom association to get the UserClassroomID
	userClassroom, err := uc.repositoryUserClassRoom.GetUserClassroomForInsertUserBehaviorRepository(request.ClassroomID, request.SubjectID)
	if err != nil {
		return nil, err
	}
	fmt.Println("user_classroom_id:", userClassroom.ID)

	// Step 4: Insert student behaviors using the retrieved UserClassroomID
	behaviorRequest := requests.StudentBehaviorRequest{
		UserClassroomID: userClassroom.ID,
		UserIDs:         request.UserIDs,
	}
	if err = uc.userBehaviorService.InsertStudentBehaviorByStudentIdAndClassroomIdService(behaviorRequest); err != nil {
		return nil, err
	}

	response := &responses.MessageUserClassRoomResponse{
		Message: "Success",
		Status:  true,
	}
	return response, nil

}

//func (uc *userClassRoomService) rollbackUserClassRooms(request requests.UserClassRoomRequest) error {
//	// Delete user-classroom associations that were created
//	for _, userID := range request.UserIDs {
//		err := uc.repositoryUserClassRoom.DeleteUserClassRoomAssociation(userID, request.ClassroomID)
//		if err != nil {
//			// Log or handle the error accordingly
//			return fmt.Errorf("failed to delete user-classroom association: %v", err)
//		}
//	}
//	return nil
//}

//func (uc *userClassRoomService) CreateUserClassRoomService(request requests.UserClassRoomRequest) (*responses.MessageUserClassRoomResponse, error) {
//	//// Check if there's already an existing association for any user in the request
//	//for _, userID := range request.UserIDs {
//	//	exists, err := uc.repositoryUserClassRoom.CheckUserClassRoomExistsRepository(userID, request.ClassRoomID)
//	//	if err != nil {
//	//		return nil, err
//	//	}
//	//	if exists {
//	//		return nil, errors.New("user already assigned to this classroom")
//	//	}
//	//}
//
//	if len(request.UserIDs) == 0 {
//		return nil, errors.New("no user IDs provided")
//	}
//	var model []models.UserClassroom
//	for _, userID := range request.UserIDs {
//		userClassroom := models.UserClassroom{
//			UserID:      userID,
//			ClassroomID: request.ClassroomID,
//		}
//		model = append(model, userClassroom)
//	}
//
//	if err := uc.repositoryUserClassRoom.CreateUserClassroomRepository(model); err != nil {
//		return nil, err
//	}
//
//	response := &responses.MessageUserClassRoomResponse{
//		Message: "Success",
//		Status:  true,
//	}
//	return response, nil
//}

func (uc *userClassRoomService) GetUserClassroomByStudentTypeService(request requests.UserClassroomRequest) (*responses.StudentClassroomResponse, error) {

	userClassrooms, err := uc.repositoryUserClassRoom.GetUserClassroomByStudentTypeRepository(request.ClassroomID, request.UserType, request.SubjectID)
	if err != nil {
		return nil, err
	}
	/*fmt.Println("response:", userClassrooms)*/
	if len(userClassrooms) == 0 {
		return nil, errors.New("no student data found ")
	}
	response := &responses.StudentClassroomResponse{
		ID:           userClassrooms[0].ID,
		ClassName:    formatClassName(userClassrooms[0].Classroom),
		SubjectName:  userClassrooms[0].Subject.SubjectName,
		UserStudents: make([]responses.UserStudents, len(userClassrooms)),
	}

	// Map user data to response struct
	for i, userClassroom := range userClassrooms {
		response.UserStudents[i] = responses.UserStudents{
			ID:        userClassroom.User.ID,
			CodeID:    userClassroom.User.CodeID,
			Firstname: userClassroom.User.Firstname,
			Lastname:  userClassroom.User.Lastname,
			Gender:    userClassroom.User.Gender,
		}
	}

	return response, nil
}

//func (uc *userClassRoomService) GetUserClassroomByStudentTypeService(request requests.UserClassroomRequest) (*responses.StudentClassroomResponse, error) {
//	userClassrooms, err := uc.repositoryUserClassRoom.GetUserClassroomByStudentTypeRepository(int(request.ClassroomID), request.UserType)
//	if err != nil {
//		return nil, err
//	}
//	var response responses.StudentClassroomResponse
//	response.ID = request.ClassroomID
//	response.ClassroomID = request.ClassroomID
//
//	if len(userClassrooms) == 0 {
//		response.ClassName = strconv.Itoa(int(userClassrooms[0].Classroom.ClassName))
//	}
//
//	var userStudents []responses.UserStudent
//	for _, data := range userClassrooms {
//		user := responses.UserStudent{
//			ID:        data.UserID,
//			CodeID:    data.User.CodeID,
//			Firstname: data.User.Firstname,
//			Lastname:  data.User.Lastname,
//			Gender:    data.User.Gender,
//		}
//		userStudents = append(userStudents, user)
//	}
//	response.UserStudent = userStudents
//	return &response, nil
//
//}

func (uc *userClassRoomService) GetClassroomByTeacherService(request requests.TeacherIdRequest) ([]responses.UserClassroomResponse, error) {
	getUserData, err := uc.repositoryUserClassRoom.GetClassroomByTeacherRepository(request.UserID, request.UserType)
	if err != nil {
		return nil, err
	}

	var response []responses.UserClassroomResponse
	for _, data := range getUserData {
		userClassroom := responses.UserClassroomResponse{
			ID:        data.Classroom.ID,
			ClassName: data.Classroom.ClassName,
			Major:     data.Classroom.Major,
			ClassYear: data.Classroom.ClassYear,
			//SubjectName: data.Classroom.ClassName,
		}
		response = append(response, userClassroom)
	}
	return response, nil
}

func (uc *userClassRoomService) DeleteUserClassRoomService(request requests.DeleteUserClassRoomByUserIDRequest) (*responses.MessageUserClassRoomResponse, error) {

	if request.Id == 0 {
		return nil, errors.New("User ID cannot be empty")
	}

	// Call the repository method to delete the student record by ID
	err := uc.repositoryUserClassRoom.DeleteUserClassRoomIDRepository(request.Id)
	if err != nil {
		return nil, err
	}

	// If successful, return a success message response
	response := &responses.MessageUserClassRoomResponse{Message: "success"}
	return response, nil
}

func (uc *userClassRoomService) GetAllUserClassRoomServices() ([]responses.UserClassRoomResponse, error) {

	getUser, err := uc.repositoryUserClassRoom.GetAllUserClassRoomRepository()

	if err != nil {
		return nil, err
	}
	if getUser == nil {
		return nil, errors.New("get User slice is nil")
	}

	var response []responses.UserClassRoomResponse

	for _, data := range getUser {
		userResponse := responses.UserClassRoomResponse{
			ID:          data.ID,
			UserID:      data.UserID,
			ClassRoomID: data.ClassroomID,
			CreatedAt:   data.CreatedAt.Format("02-01-2006 15:01:05"),
			UpdatedAt:   data.UpdatedAt.Format("02-01-2006 15:01:05"),
		}
		response = append(response, userResponse)
	}
	return response, err
}

func (uc *userClassRoomService) GetByIdUserClassRoomService(id uint) (*responses.UserClassRoomResponse, error) {

	data, err := uc.repositoryUserClassRoom.GetByUserIDRepository(uint(id))
	if err != nil {
		return nil, err
	}
	userResponse := responses.UserClassRoomResponse{
		ID:          data.ID,
		UserID:      data.UserID,
		ClassRoomID: data.ClassroomID,
		CreatedAt:   data.CreatedAt.Format("02-01-2006 15:01:05"),
		UpdatedAt:   data.UpdatedAt.Format("02-01-2006 15:01:05"),
	}
	return &userResponse, err
}

func (uc *userClassRoomService) GetByUserClassRoomIdService(classRoomID uint) (*responses.UserClassRoomResponse, error) {

	data, err := uc.repositoryUserClassRoom.GetByClassroomIdRepository(uint(classRoomID))
	if err != nil {
		return nil, err
	}
	userResponse := responses.UserClassRoomResponse{
		ID:          data.ID,
		UserID:      data.UserID,
		ClassRoomID: data.ClassroomID,
		CreatedAt:   data.CreatedAt.Format("02-01-2006 15:01:05"),
		UpdatedAt:   data.UpdatedAt.Format("02-01-2006 15:01:05"),
	}
	return &userResponse, err
}

func (uc *userClassRoomService) GetByUserIDService(userID uint) (*responses.UserClassRoomResponse, error) {

	data, err := uc.repositoryUserClassRoom.GetByIdUserClassRoomRepository(uint(userID))
	if err != nil {
		return nil, err
	}
	userResponse := responses.UserClassRoomResponse{
		ID:          data.ID,
		UserID:      data.UserID,
		ClassRoomID: data.ClassroomID,
		CreatedAt:   data.CreatedAt.Format("02-01-2006 15:01:05"),
		UpdatedAt:   data.UpdatedAt.Format("02-01-2006 15:01:05"),
	}
	return &userResponse, err
}

func (uc *userClassRoomService) UpdateUserClassRoomService(request requests.UpdateUserClassRoomRequest) (*responses.MessageUserClassRoomResponse, error) {

	model := models.UserClassroom{
		UserID:      request.UserID,
		ClassroomID: request.ClassRoomID,
	}
	if err := uc.repositoryUserClassRoom.UpdateUserClassRoomRepository(&model); err != nil {
		return nil, err
	}
	response := &responses.MessageUserClassRoomResponse{Message: "success"}
	return response, nil
}

func NewUserClassRoomService(
	repositoryUserClassRoom repositories.UserClassRoomRepository,
	userBehaviorService UserBehaviorService,
) UserClassRoomService {

	return &userClassRoomService{
		repositoryUserClassRoom: repositoryUserClassRoom,
		userBehaviorService:     userBehaviorService,
	}
}
