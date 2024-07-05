package services

import (
	"fmt"
	"go_starter/errs"
	"go_starter/models"
	"go_starter/repositories"
	"go_starter/requests"
	"go_starter/responses"
	"go_starter/security"
	"go_starter/trails"
	"strings"

	"github.com/pkg/errors"
)

type UserService interface {
	GetUserInfoService(request requests.UserIdRequest) (*responses.UserResponse, error)
	GetUserTypeWithPaginationService(request requests.UserWithPaginationRequest, paginateRequest trails.PaginateRequest) (*responses.UserPaginatedResponse, error)
	GetUserByUserTypeService(request requests.UserTypeRequest) ([]responses.UserResponse, error)

	//admin
	CountTotalSubjectsService() (*responses.CountSubjectsResponse, error)
	CountTotalTeachersService() (*responses.CountTeachersResponse, error)
	CountTotalStudentService() (*responses.CountStudentResponse, error)

	//Login
	ResetUserPasswordService(request requests.ResetPasswordRequest) (string, error)
	AdminSignInUserService(request requests.SignInUserRequest) (*responses.SignInResponse, error)
	SignInUserService(request requests.SignInUserRequest) (*responses.SignInResponse, error)
	SignUpUserService(request requests.SigUpUserRequest) (*responses.SignUpResponse, error)

	//CRUD
	GetAllUserService() ([]responses.UserResponse, error)
	GetByIdUserService(id uint) (*responses.UserResponse, error)
	GetByCodeIdUserService(request requests.UserCodeIdRequest) (*responses.UserResponse, error)
	GetByPhoneUserService(request requests.UserPhoneRequest) (*responses.UserResponse, error)
	CreateUserService(request requests.UserRequest) (*responses.MessageUserResponse, error)
	UpdateUserService(request requests.UserRequest) (*responses.MessageUserResponse, error)
	DeleteUserService(request requests.UserCodeIdRequest) (*responses.MessageUserResponse, error)
}

type userService struct {
	repositoryUserRepository repositories.UserRepository
}

func (u *userService) GetUserInfoService(request requests.UserIdRequest) (*responses.UserResponse, error) {
	data, err := u.repositoryUserRepository.GetUserByIDRepository(request.ID)

	if err != nil {
		return nil, err
	}
	response := &responses.UserResponse{
		ID:        data.ID,
		CodeID:    data.CodeID,
		Firstname: data.Firstname,
		Lastname:  data.Lastname,
		Phone:     data.Phone,
		Gender:    data.Gender,
		Degree:    data.Degree,
		Skill:     data.Skill,
		UserType:  data.UserType,
		CreatedAt: data.CreatedAt.Format("02-01-2006 15:01:05"),
		UpdatedAt: data.UpdatedAt.Format("02-01-2006 15:01:05"),
	}
	return response, err
}

func (u *userService) AdminSignInUserService(request requests.SignInUserRequest) (*responses.SignInResponse, error) {
	if request.Phone == "" {
		return nil, errs.ErrorBadRequest("PHONE_CANT_BY_EMPTY")
	}
	if len(request.Phone) > 10 || len(request.Phone) < 9 {

		return nil, errs.ErrorBadRequest("PHONE_INVALID")

	}
	trimSpacePassword := strings.TrimSpace(request.Password)
	if trimSpacePassword == "" {
		return nil, errs.ErrorBadRequest("Password_CANT_BE_EMPTY")
	}

	userData, err := u.repositoryUserRepository.GetByPhoneUserRepository(request.Phone)
	if err != nil {
		return nil, errs.ErrorBadRequest("id not found teacher ")
	}
	if userData.UserType != "admin" {
		return nil, errs.ErrorBadRequest("CANNOT_ACCESS")
	}
	err = security.VerifyPassword(userData.Password, request.Password)
	if err != nil {
		return nil, fmt.Errorf("password doesn't match")
	}
	response := responses.SignInResponse{
		Message:  "success",
		Phone:    userData.Phone,
		ID:       userData.ID,
		UserType: userData.UserType,
	}
	return &response, err
}

func (u *userService) ResetUserPasswordService(request requests.ResetPasswordRequest) (string, error) {
	// Encrypt the new password
	encryptedPassword, err := security.EncryptPassword(request.NewPassword)
	if err != nil {
		return "", err
	}

	// Reset the user's password in the repository
	if err = u.repositoryUserRepository.UpdateUserPasswordByPhoneRepository(request.Phone, encryptedPassword); err != nil {
		return "", err
	}

	return "reset", nil
}

func (u *userService) GetUserTypeWithPaginationService(request requests.UserWithPaginationRequest, paginateRequest trails.PaginateRequest) (*responses.UserPaginatedResponse, error) {
	response, userData, err := u.repositoryUserRepository.GetUserTypeWithPaginationRepository(request, paginateRequest)
	if err != nil {
		return nil, err
	}

	var responseData []responses.UserResponse
	for _, item := range userData {
		responseData = append(responseData, responses.UserResponse{
			ID:        item.ID,
			CodeID:    item.CodeID,
			Firstname: item.Firstname,
			Lastname:  item.Lastname,
			Phone:     item.Phone,
			Gender:    item.Gender,
			Degree:    item.Degree,
			Skill:     item.Skill,
			UserType:  item.UserType,
			CreatedAt: item.CreatedAt.Format("02-01-2006 15:01:05"),
			UpdatedAt: item.UpdatedAt.Format("02-01-2006 15:01:05"),
		})
	}
	// Check if userData is empty, if so, return an empty slice for Items
	if len(userData) == 0 {
		responseData = []responses.UserResponse{}
	}

	paginatedResponse := &responses.UserPaginatedResponse{
		TotalPages:  response.TotalPages,
		PerPage:     response.PerPage,
		CurrentPage: response.CurrentPage,
		Sorting:     response.Sorting,
		Items:       responseData,
	}
	return paginatedResponse, nil
}

func (u *userService) GetUserByUserTypeService(request requests.UserTypeRequest) ([]responses.UserResponse, error) {
	users, err := u.repositoryUserRepository.GetUserByUserTypeRepository(request.UserType)
	if err != nil {
		return nil, err
	}

	var userResponses []responses.UserResponse
	for _, user := range users {
		userResponses = append(userResponses, responses.UserResponse{
			ID:        user.ID,
			CodeID:    user.CodeID,
			Firstname: user.Firstname,
			Lastname:  user.Lastname,
			Phone:     user.Phone,
			Gender:    user.Gender,
			Degree:    user.Degree,
			Skill:     user.Skill,
			UserType:  user.UserType,
			CreatedAt: user.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt: user.UpdatedAt.Format("2006-01-02 15:04:05"),
		})
	}

	if len(userResponses) == 0 {
		return nil, errors.New("no users found with the specified user type")
	}

	return userResponses, nil
}

func (u *userService) CountTotalSubjectsService() (*responses.CountSubjectsResponse, error) {
	getData, err := u.repositoryUserRepository.CountTotalSubjectsRepository()
	if err != nil {
		return nil, err
	}
	response := responses.CountSubjectsResponse{
		TotalSubjects: getData,
	}
	return &response, err
}

func (u *userService) CountTotalTeachersService() (*responses.CountTeachersResponse, error) {
	getData, err := u.repositoryUserRepository.CountTotalTeachersRepository()
	if err != nil {
		return nil, err
	}
	response := responses.CountTeachersResponse{
		TotalTeachers: getData,
	}
	return &response, err
}

func (u *userService) CountTotalStudentService() (*responses.CountStudentResponse, error) {
	getData, err := u.repositoryUserRepository.CountTotalStudentsRepository()
	if err != nil {
		return nil, err
	}
	response := responses.CountStudentResponse{
		TotalStudents: getData,
	}
	return &response, err
}

func (u *userService) CreateUserService(request requests.UserRequest) (*responses.MessageUserResponse, error) {

	// Convert the student ID to uppercase
	codeID := strings.ToUpper(request.CodeID)

	// Check if the Code ID or phone number is already in use
	if checkCodeID, err := u.repositoryUserRepository.CheckUserCodeIDAlreadyHas(codeID); err != nil {
		return nil, err
	} else if checkCodeID {
		return nil, errors.New("code ID already in use")
	}

	if checkPhone, err := u.repositoryUserRepository.CheckUserPhoneAlreadyHas(request.Phone); err != nil {
		return nil, err
	} else if checkPhone {
		return nil, errors.New("phone number already in use")
	}
	// Create the student model
	model := models.User{
		CodeID:    request.CodeID,
		Firstname: request.Firstname,
		Lastname:  request.Lastname,
		Phone:     request.Phone,
		Gender:    request.Gender,
		Degree:    request.Degree,
		Skill:     request.Skill,
		UserType:  request.UserType,
	}

	// Call the repository method to create the student record
	if err := u.repositoryUserRepository.CreateUserRepository(&model); err != nil {
		return nil, err
	}

	// If successful, return a success message response
	response := &responses.MessageUserResponse{Message: "success"}
	return response, nil
}

func (u *userService) DeleteUserService(request requests.UserCodeIdRequest) (*responses.MessageUserResponse, error) {

	if request.CodeID == "" {
		return nil, errors.New("code ID cannot be empty")
	}

	// Call the repository method to delete the student record by ID
	err := u.repositoryUserRepository.DeleteUserRepository(request.CodeID)
	if err != nil {
		return nil, err
	}

	// If successful, return a success message response
	response := &responses.MessageUserResponse{Message: "success"}
	return response, nil
}

func (u *userService) GetAllUserService() ([]responses.UserResponse, error) {

	getUser, err := u.repositoryUserRepository.GetAllUserRepository()
	if err != nil {
		return nil, err
	}

	// Check if getStudent is nil
	if getUser == nil {
		return nil, errors.New("getUser slice is nil")
	}

	var response []responses.UserResponse
	for _, data := range getUser {
		userResponse := responses.UserResponse{

			ID:        data.ID,
			CodeID:    data.CodeID,
			Firstname: data.Firstname,
			Lastname:  data.Lastname,
			Phone:     data.Phone,
			Gender:    data.Gender,
			Degree:    data.Degree,
			Skill:     data.Skill,
			UserType:  data.UserType,
			CreatedAt: data.CreatedAt.Format("02-01-2006 15:01:05"),
			UpdatedAt: data.UpdatedAt.Format("02-01-2006 15:01:05"),
		}
		response = append(response, userResponse)
	}
	return response, err
}

func (u *userService) GetByCodeIdUserService(request requests.UserCodeIdRequest) (*responses.UserResponse, error) {

	data, err := u.repositoryUserRepository.GetByCodeIDUserRepository(request.CodeID)

	if err != nil {
		return nil, err
	}
	response := &responses.UserResponse{
		ID:        data.ID,
		CodeID:    data.CodeID,
		Firstname: data.Firstname,
		Lastname:  data.Lastname,
		Phone:     data.Phone,
		Gender:    data.Gender,
		Degree:    data.Degree,
		Skill:     data.Skill,
		UserType:  data.UserType,
		CreatedAt: data.CreatedAt.Format("02-01-2006 15:01:05"),
		UpdatedAt: data.UpdatedAt.Format("02-01-2006 15:01:05"),
	}
	return response, err
}

func (u *userService) GetByIdUserService(id uint) (*responses.UserResponse, error) {

	data, err := u.repositoryUserRepository.GetByIdUserRepository(uint(id))
	if err != nil {
		return nil, err
	}

	response := &responses.UserResponse{
		ID:        data.ID,
		CodeID:    data.CodeID,
		Firstname: data.Firstname,
		Lastname:  data.Lastname,
		Phone:     data.Phone,
		Gender:    data.Gender,
		Degree:    data.Degree,
		Skill:     data.Skill,
		UserType:  data.UserType,
		CreatedAt: data.CreatedAt.Format("02-01-2006 15:01:05"),
		UpdatedAt: data.UpdatedAt.Format("02-01-2006 15:01:05"),
	}
	return response, err
}

func (u *userService) GetByPhoneUserService(request requests.UserPhoneRequest) (*responses.UserResponse, error) {

	data, err := u.repositoryUserRepository.GetByPhoneUserRepository(request.Phone)

	if err != nil {
		return nil, err
	}
	response := &responses.UserResponse{
		ID:        data.ID,
		CodeID:    data.CodeID,
		Firstname: data.Firstname,
		Lastname:  data.Lastname,
		Phone:     data.Phone,
		Gender:    data.Gender,
		Degree:    data.Degree,
		Skill:     data.Skill,
		UserType:  data.UserType,
		CreatedAt: data.CreatedAt.Format("02-01-2006 15:01:05"),
		UpdatedAt: data.UpdatedAt.Format("02-01-2006 15:01:05"),
	}
	return response, err

}

func (u *userService) SignInUserService(request requests.SignInUserRequest) (*responses.SignInResponse, error) {

	if request.Phone == "" {
		return nil, errs.ErrorBadRequest("PHONE_CANT_BY_EMPTY")
	}
	if len(request.Phone) > 10 || len(request.Phone) < 9 {

		return nil, errs.ErrorBadRequest("PHONE_INVALID")

	}
	trimSpacePassword := strings.TrimSpace(request.Password)
	if trimSpacePassword == "" {
		return nil, errs.ErrorBadRequest("Password_CANT_BE_EMPTY")
	}

	userData, err := u.repositoryUserRepository.GetByPhoneUserRepository(request.Phone)
	if err != nil {
		return nil, errs.ErrorBadRequest("not found teacher id")
	}
	if userData.UserType != "teacher" && userData.UserType != "student" {
		return nil, errs.ErrorBadRequest("CANNOT_ACCESS")
	}
	err = security.VerifyPassword(userData.Password, request.Password)
	if err != nil {
		return nil, fmt.Errorf("password doesn't match")
	}
	response := responses.SignInResponse{
		Message:  "success",
		Phone:    userData.Phone,
		ID:       userData.ID,
		UserType: userData.UserType,
	}
	return &response, err

}

func (u *userService) SignUpUserService(request requests.SigUpUserRequest) (*responses.SignUpResponse, error) {
	// Validate phone number
	if request.Phone == "" {
		return nil, errs.ErrorBadRequest("PHONE_CANT_BE_EMPTY")
	}
	if len(request.Phone) > 10 || len(request.Phone) < 9 {
		return nil, errs.ErrorBadRequest("PHONE_INVALID")
	}
	if checkPhone, err := u.repositoryUserRepository.CheckUserPhoneAlreadyHas(request.Phone); err != nil {
		return nil, err
	} else if checkPhone {
		return nil, errors.New("phone number already in use")
	}
	trimSpaceUser := strings.TrimSpace(request.Password)
	if trimSpaceUser == "" {
		return nil, errs.ErrorBadRequest("Password_CANT_BE_EMPTY")
	}
	encryptPassword, err := security.EncryptPassword(request.Password)
	if err != nil {
		return nil, err
	}
	if request.UserType != "teacher" && request.UserType != "student" && request.UserType != "admin" {
		return nil, fmt.Errorf("user_type needs to be  'teacher' or 'student'")
		//return nil, fmt.Errorf("user_type needs to be  'teacher' or 'student' or 'admin' ")

	}

	user := models.User{
		Phone:     request.Phone,
		Password:  encryptPassword,
		UserType:  request.UserType,
		CodeID:    request.CodeID,
		Firstname: request.Firstname,
		Lastname:  request.Lastname,
		Gender:    request.Gender,
		Degree:    request.Degree,
	}
	signUpUser, err := u.repositoryUserRepository.SignUpForUserRepository(user)
	if err != nil {
		return nil, err
	}
	response := responses.SignUpResponse{
		Message:  "success",
		Phone:    signUpUser.Phone,
		ID:       signUpUser.ID,
		UserType: request.UserType,
	}
	return &response, nil

}

func (u *userService) UpdateUserService(request requests.UserRequest) (*responses.MessageUserResponse, error) {

	model := models.User{
		CodeID:    request.CodeID,
		Firstname: request.Firstname,
		Lastname:  request.Lastname,
		Phone:     request.Phone,
		Gender:    request.Gender,
		Degree:    request.Degree,
		Skill:     request.Skill,
		UserType:  request.UserType,
	}
	if err := u.repositoryUserRepository.UpdateUserRepository(&model); err != nil {

		return nil, err
	}
	response := &responses.MessageUserResponse{Message: "success"}
	return response, nil
}

func NewUserService(repositoryUserRepository repositories.UserRepository) UserService {
	return &userService{
		repositoryUserRepository: repositoryUserRepository,
	}
}
