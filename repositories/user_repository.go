package repositories

import (
	"errors"
	"go_starter/logs"
	"go_starter/models"
	"go_starter/requests"
	"go_starter/trails"
	"strings"

	"gorm.io/gorm"
)

type UserRepository interface {
	GetUserTypeWithPaginationRepository(request requests.UserWithPaginationRequest, paginateRequest trails.PaginateRequest) (*trails.PaginatedResponse, []models.User, error)
	GetUserByUserTypeRepository(userType string) ([]models.User, error)
	CountTotalSubjectsRepository() (int64, error)
	CountTotalTeachersRepository() (int64, error)
	CountTotalStudentsRepository() (int64, error)

	//Login Users
	UpdateUserPasswordByPhoneRepository(phone string, newPassword string) error
	SignUpForUserRepository(request models.User) (*models.User, error)

	//CRUD User
	GetUserByIDRepository(id uint) (*models.User, error)
	GetAllUserRepository() ([]models.User, error)
	GetByIdUserRepository(id uint) (*models.User, error)
	GetByPhoneUserRepository(phone string) (*models.User, error)
	GetByCodeIDUserRepository(code_ID string) (*models.User, error)
	CreateUserRepository(request *models.User) error
	UpdateUserRepository(request *models.User) error
	DeleteUserRepository(codeID string) error

	//Check Phone and Check CodeID
	CheckUserPhoneAlreadyHas(phone string) (bool, error)
	CheckUserCodeIDAlreadyHas(codeID string) (bool, error)
}

type userRepository struct{ db *gorm.DB }

func (u *userRepository) GetUserByIDRepository(id uint) (*models.User, error) {
	var model models.User
	// Execute raw SQL query
	query := u.db.Raw("SELECT * FROM users WHERE id = ?", id).Scan(&model).Error
	if query != nil {
		return nil, query
	}
	return &model, nil
}

func (u *userRepository) UpdateUserPasswordByPhoneRepository(phone string, newPassword string) error {
	var user models.User

	// Find the user by phone
	if err := u.db.Where("phone = ?", phone).First(&user).Error; err != nil {
		logs.Error(err)
		return err
	}

	// Update the user's password
	user.Password = newPassword
	if err := u.db.Save(&user).Error; err != nil {
		logs.Error(err)
		return err
	}

	return nil
}

func (u *userRepository) GetUserTypeWithPaginationRepository(request requests.UserWithPaginationRequest, paginateRequest trails.PaginateRequest) (*trails.PaginatedResponse, []models.User, error) {
	var model []models.User
	query := u.db.Model(&models.User{})

	// Apply filtering based on user type if provided
	if request.UserType != "" {
		query = query.Where("user_type = ?", request.UserType)
	}

	pagination, err := trails.PaginationData(query, &model, paginateRequest, true)
	if err != nil {
		return nil, nil, err
	}

	if len(model) == 0 {
		pagination.Items = []models.User{}
	}

	return pagination, model, nil
}

//func (u *userRepository) GetUserTypeWithPaginationRepository(request requests.UserWithPaginationRequest, paginateRequest trails.PaginateRequest) (*trails.PaginatedResponse, []models.User, error) {
//	var model []models.User
//	query := u.db.Model(&models.User{})
//	// Apply filtering based on phone if provided
//	if request.UserType != "" {
//		query = query.Where("user_type = ?", request.UserType)
//	}
//	// Check if user type exists
//	var count int64
//	query.Count(&count)
//	if count == 0 {
//		return &trails.PaginatedResponse{
//			TotalPages:  1,
//			PerPage:     paginateRequest.Item,
//			CurrentPage: paginateRequest.Page,
//			Sorting:     paginateRequest.Sorting,
//			UserType:    request.UserType,
//			Items:       []models.User{},
//		}, nil, nil
//	}
//
//	// Sorting
//	switch request.Sorting {
//	case "max":
//		query = query.Order("id DESC")
//	case "min":
//		query = query.Order("id ASC")
//	default:
//		query = query.Order("id DESC")
//	}
//	pagination, err := trails.PaginationData(query, &model, paginateRequest)
//	if err != nil {
//		return nil, nil, err
//	}
//	pagination.Items = model
//	pagination.UserType = request.UserType
//	return pagination, model, nil
//}

func (u *userRepository) GetUserByUserTypeRepository(userType string) ([]models.User, error) {
	var users []models.User
	result := u.db.Where("user_type = ?", userType).Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	if len(users) == 0 {
		return nil, errors.New("no users found with the specified user type")
	}
	return users, nil
}

func (u *userRepository) CountTotalSubjectsRepository() (int64, error) {
	var count int64
	err := u.db.Model(&models.Subject{}).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (u *userRepository) CountTotalTeachersRepository() (int64, error) {
	var count int64
	err := u.db.Model(&models.User{}).Where("user_type = ?", "teacher").Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (u *userRepository) CountTotalStudentsRepository() (int64, error) {
	var count int64
	err := u.db.Model(&models.User{}).Where("user_type = ?", "student").Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (u *userRepository) GetStudentClassroomByTeacher(userId int) ([]models.UserClassroom, error) {
	panic("unimplemented")
}

func (u *userRepository) CheckUserCodeIDAlreadyHas(codeID string) (bool, error) {

	var count int64
	// Convert studentID to uppercase for comparison
	upperCodeID := strings.ToUpper(codeID)
	// Perform a case-insensitive comparison
	query := u.db.Model(&models.User{}).Where("UPPER(code_id) = ?", upperCodeID).Count(&count)
	if query.Error != nil {
		return false, query.Error
	}
	return count > 0, nil
}

func (u *userRepository) CheckUserPhoneAlreadyHas(phone string) (bool, error) {

	var count int64
	query := u.db.Model(&models.User{}).Where("phone = ?", phone).Count(&count)
	if query.Error != nil {
		return false, query.Error
	}
	return count > 0, nil
}

func (u *userRepository) SignUpForUserRepository(request models.User) (*models.User, error) {

	create := u.db.Create(&request)
	if create.Error != nil {
		logs.Error(create.Error)
		return nil, create.Error
	}
	return &request, nil
}

func (u *userRepository) CreateUserRepository(request *models.User) error {

	if err := u.db.Create(request).Error; err != nil {
		return err
	}
	return nil
}

func (u *userRepository) DeleteUserRepository(codeID string) error {

	var count int64
	if err := u.db.Model(&models.User{}).Where("code_id = ?", codeID).Count(&count).Error; err != nil {
		return err
	}
	if count == 0 {
		return errors.New("code ID not found")
	}

	// Delete the student
	if err := u.db.Where("code_id = ?", codeID).Delete(&models.User{}).Error; err != nil {
		return err
	}
	return nil
}

func (u *userRepository) GetAllUserRepository() ([]models.User, error) {

	var model []models.User
	query := u.db.Find(&model).Error
	if query != nil {
		return nil, query
	}
	return model, nil
}

func (u *userRepository) GetByPhoneUserRepository(phone string) (*models.User, error) {

	var model models.User
	query := u.db.First(&model, "phone = ?", phone)

	if query.Error != nil {
		return nil, nil
	}
	return &model, nil
}

func (u *userRepository) GetByCodeIDUserRepository(code_ID string) (*models.User, error) {

	var model models.User
	// Execute raw SQL query
	query := u.db.Raw("SELECT * FROM users WHERE code_id = ?", code_ID).Scan(&model).Error
	if query != nil {
		return nil, query
	}
	return &model, nil
}

func (u *userRepository) GetByIdUserRepository(id uint) (*models.User, error) {

	var model models.User

	// Execute raw SQL query
	query := u.db.Raw("SELECT * FROM users WHERE id = ?", id).Scan(&model).Error
	if query != nil {
		return nil, query
	}
	return &model, nil
}

func (u *userRepository) UpdateUserRepository(request *models.User) error {

	query := u.db.Model(&models.User{}).Where("code_id = ?", request.CodeID).Updates(request)
	if query.Error != nil {
		return query.Error
	}
	if query.RowsAffected == 0 {
		return errors.New("no code_id found")
	}
	return nil
}

func NewUserRepository(db *gorm.DB) UserRepository {
	// db.Migrator().DropTable(&models.User{})
	//db.AutoMigrate(&models.User{})
	return &userRepository{db: db}
}
