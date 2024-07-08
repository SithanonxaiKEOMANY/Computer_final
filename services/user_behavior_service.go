package services

import (
	"go_starter/models"
	"go_starter/repositories"
	"go_starter/requests"
	"go_starter/responses"
	"go_starter/trails"
	"time"

	"github.com/pkg/errors"
)

type UserBehaviorService interface {
	ReportUserBehaviorWithPaginationService(request requests.UserBehaviorWithPaginationRequest, paginateRequest trails.PaginateRequest) (*responses.UserBehaviorPaginatedResponse, error)
	UpdateUserBehaviorService(request requests.UserBehaviorRequest) (*responses.MessageUserBehaviorResponse, error)
	InsertStudentBehaviorByStudentIdAndClassroomIdService(request requests.StudentBehaviorRequest) error

	//GetAllUserBehaviorServices() ([]responses.UserBehaviorResponse, error)
	//GetByIdUserBehaviorService(id uint) (*responses.UserBehaviorResponse, error)
	//GetByUserIDService(userID uint) (*responses.UserBehaviorResponse, error)
	//GetByUserBehaviorIdService(classRoomID uint) (*responses.UserBehaviorResponse, error)
	//CreateUserBehaviorService(request requests.UserBehaviorRequest) (*responses.MessageUserBehaviorResponse, error)
	//UpdateUserBehaviorService(request requests.UserBehaviorRequest) (*responses.MessageUserBehaviorResponse, error)

	//DeleteUserBehaviorService(request requests.UserBehaviorClassRoomByUserIDRequest) (*responses.MessageUserBehaviorResponse, error)
}

type userBehaviorService struct {
	repositoryUserBehavior repositories.UserBehaviorRepository
}

func (ub *userBehaviorService) ReportUserBehaviorWithPaginationService(request requests.UserBehaviorWithPaginationRequest, paginateRequest trails.PaginateRequest) (*responses.UserBehaviorPaginatedResponse, error) {
	response, userBehaviorData, err := ub.repositoryUserBehavior.GetUserBehaviorWithPaginationRepository(request, paginateRequest)
	if err != nil {
		return nil, err
	}
	var responseData []responses.UserBehaviorResponse
	for _, item := range userBehaviorData {
		responseData = append(responseData, responses.UserBehaviorResponse{
			ID:              item.ID,
			UserID:          item.UserID,
			UserClassroomID: item.UserClassroomID,
			CountCheck:      item.CountCheck,
			CountAbsent:     item.CountAbsent,
			CountVacation:   item.CountVacation,
			CountBreakRule:  item.CountBreakRule,
			//StudentCheck:     item.StudentCheck,
			//StudentAbsent:    item.StudentAbsent,
			//StudentVacation:  item.StudentVacation,
			//StudentBreakRule: item.StudentBreakRule,
			//CreatedAt:        item.CreatedAt.Format("02-01-2006 15:01:05"),
			//UpdatedAt:        item.UpdatedAt.Format("02-01-2006 15:01:05"),
		})
	}
	if len(userBehaviorData) == 0 {
		responseData = []responses.UserBehaviorResponse{}
	}
	paginatedResponse := responses.UserBehaviorPaginatedResponse{
		TotalPages:  response.TotalPages,
		PerPage:     response.PerPage,
		CurrentPage: response.CurrentPage, 
		Sorting:     response.Sorting,
		Items:       responseData,
	}
	return &paginatedResponse, nil
}

func (ub *userBehaviorService) UpdateUserBehaviorService(request requests.UserBehaviorRequest) (*responses.MessageUserBehaviorResponse, error) {
	if request.UserClassroomID == 0 {
		return nil, errors.New("user_class_room_id is required")
	}

	// Prepare a slice of UserBehavior models
	var userBehaviors []models.UserBehavior
	for _, userBehaviorRequest := range request.UserBehaviors {
		if userBehaviorRequest.UserID == 0 {
			return nil, errors.New("user_id is required for each user behavior")
		}

		userBehavior := models.UserBehavior{
			UserClassroomID:  request.UserClassroomID,
			UserID:           userBehaviorRequest.UserID,
			StudentCheck:     userBehaviorRequest.StudentCheck,
			StudentAbsent:    userBehaviorRequest.StudentAbsent,
			StudentVacation:  userBehaviorRequest.StudentVacation,
			StudentBreakRule: userBehaviorRequest.StudentBreakRule,
			UpdatedAt:        time.Now(),
		}

		userBehaviors = append(userBehaviors, userBehavior)
	}

	// Update the user behaviors in the database
	err := ub.repositoryUserBehavior.UpdateUserBehaviorRepository(userBehaviors)
	if err != nil {
		return nil, err
	}

	// Prepare the response
	response := &responses.MessageUserBehaviorResponse{
		Message: "User behaviors updated successfully",
		Status:  true,
	}

	return response, nil
}

func (ub *userBehaviorService) InsertStudentBehaviorByStudentIdAndClassroomIdService(request requests.StudentBehaviorRequest) error {
	if request.UserClassroomID == 0 {
		return errors.New("classroom_id is required")
	}
	if len(request.UserIDs) == 0 {
		return errors.New("at least one user_id is required")
	}

	behaviors := make([]models.UserBehavior, len(request.UserIDs))
	for i, userID := range request.UserIDs {
		behaviors[i] = models.UserBehavior{
			UserID:           userID,
			UserClassroomID:  request.UserClassroomID,
			StudentCheck:     false,
			StudentAbsent:    false,
			StudentVacation:  false,
			StudentBreakRule: false,
			CountCheck:       0,
			CountAbsent:      0,
			CountVacation:    0,
			CountBreakRule:   0,
			CreatedAt:        time.Now(),
			UpdatedAt:        time.Now(),
		}
	}

	err := ub.repositoryUserBehavior.InsertUserBehaviorRepository(behaviors)
	if err != nil {
		return err
	}

	return nil
}

//func (ub *userBehaviorService) DeleteUserBehaviorService(request requests.UserBehaviorClassRoomByUserIDRequest) (*responses.MessageUserBehaviorResponse, error) {
//
//	err := ub.repositoryUserBehavior.DeleteUserBehaviorByUserIDRepository(request.UserID)
//	if err != nil {
//		return nil, err
//	}
//
//	// If successful, return a success message response
//	response := &responses.MessageUserBehaviorResponse{Message: "success"}
//	return response, nil
//}

//func (ub *userBehaviorService) GetAllUserBehaviorServices() ([]responses.UserBehaviorResponse, error) {
//
//	getUser, err := ub.repositoryUserBehavior.GetAllUserBehaviorRepository()
//
//	if err != nil {
//		return nil, err
//	}
//	if getUser == nil {
//		return nil, errors.New("get User slice is nil")
//	}
//
//	var response []responses.UserBehaviorResponse
//
//	for _, data := range getUser {
//		userResponse := responses.UserBehaviorResponse{
//			ID:     data.ID,
//			UserID: data.UserID,
//			//ClassRoomID:         data.ClassRoomID,
//			StudentCheck:        data.StudentCheck,
//			StudentAbsent:       data.StudentAbsent,
//			StudentVacation:     data.StudentVacation,
//			StudentBreakingRule: data.StudentBreakRule,
//			CreatedAt:           data.CreatedAt.Format("02-01-2006 15:01:05"),
//			UpdatedAt:           data.UpdatedAt.Format("02-01-2006 15:01:05"),
//		}
//		response = append(response, userResponse)
//	}
//	return response, err
//}
//
//func (ub *userBehaviorService) GetByIdUserBehaviorService(id uint) (*responses.UserBehaviorResponse, error) {
//
//	data, err := ub.repositoryUserBehavior.GetByIdUserBehaviorRepository(uint(id))
//	if err != nil {
//		return nil, err
//	}
//	userResponse := responses.UserBehaviorResponse{
//		ID:     data.ID,
//		UserID: data.UserID,
//		//ClassRoomID:         data.ClassRoomID,
//		StudentCheck:        data.StudentCheck,
//		StudentAbsent:       data.StudentAbsent,
//		StudentVacation:     data.StudentVacation,
//		StudentBreakingRule: data.StudentBreakRule,
//		CreatedAt:           data.CreatedAt.Format("02-01-2006 15:01:05"),
//		UpdatedAt:           data.UpdatedAt.Format("02-01-2006 15:01:05"),
//	}
//	return &userResponse, err
//}
//
//func (ub *userBehaviorService) GetByUserBehaviorIdService(classRoomID uint) (*responses.UserBehaviorResponse, error) {
//
//	data, err := ub.repositoryUserBehavior.GetByClassroomIdRepository(uint(classRoomID))
//	if err != nil {
//		return nil, err
//	}
//	userResponse := responses.UserBehaviorResponse{
//		ID:     data.ID,
//		UserID: data.UserID,
//		//ClassRoomID:         data.ClassRoomID,
//		StudentCheck:        data.StudentCheck,
//		StudentAbsent:       data.StudentAbsent,
//		StudentVacation:     data.StudentVacation,
//		StudentBreakingRule: data.StudentBreakRule,
//		CreatedAt:           data.CreatedAt.Format("02-01-2006 15:01:05"),
//		UpdatedAt:           data.UpdatedAt.Format("02-01-2006 15:01:05"),
//	}
//	return &userResponse, err
//}
//
//func (ub *userBehaviorService) GetByUserIDService(userID uint) (*responses.UserBehaviorResponse, error) {
//
//	data, err := ub.repositoryUserBehavior.GetByIdUserBehaviorRepository(uint(userID))
//	if err != nil {
//		return nil, err
//	}
//	userResponse := responses.UserBehaviorResponse{
//		ID:     data.ID,
//		UserID: data.UserID,
//		//ClassRoomID:         data.ClassRoomID,
//		StudentCheck:        data.StudentCheck,
//		StudentAbsent:       data.StudentAbsent,
//		StudentVacation:     data.StudentVacation,
//		StudentBreakingRule: data.StudentBreakRule,
//		CreatedAt:           data.CreatedAt.Format("02-01-2006 15:01:05"),
//		UpdatedAt:           data.UpdatedAt.Format("02-01-2006 15:01:05"),
//	}
//	return &userResponse, err
//}

//func (u *userBehaviorService) UpdateUserBehaviorService(request requests.UserBehaviorRequest) (*responses.MessageUserBehaviorResponse, error) {
//
//	//model := models.UserBehavior{
//	//	UserID: request.UserID,
//	//	//ClassRoomID:      request.ClassRoomID,
//	//	StudentCheck:     request.StudentCheck,
//	//	StudentAbsent:    request.StudentAbsent,
//	//	StudentVacation:  request.StudentVacation,
//	//	StudentBreakRule: request.StudentBreakingRule,
//	//}
//	//if err := u.repositoryUserBehavior.UpdateUserBehaviorRepository(&model); err != nil {
//	//	return nil, err
//	//}
//	//response := &responses.MessageUserBehaviorResponse{Message: "success"}
//	//return response, nil
//	return nil, nil
//}

func NewUserBehaviorService(
	repositoryUserBehavior repositories.UserBehaviorRepository,
) UserBehaviorService {

	return &userBehaviorService{
		repositoryUserBehavior: repositoryUserBehavior,
	}
}
