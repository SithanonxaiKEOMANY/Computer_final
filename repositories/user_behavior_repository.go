package repositories

import (
	"errors"
	"go_starter/models"
	"go_starter/requests"
	"go_starter/trails"
	"gorm.io/gorm"
	"time"
)

type UserBehaviorRepository interface {
	GetUserBehaviorWithPaginationRepository(request requests.UserBehaviorWithPaginationRequest, paginateRequest trails.PaginateRequest) (*trails.PaginatedResponse, []models.UserBehavior, error)
	UpdateUserBehaviorRepository(request []models.UserBehavior) error
	InsertUserBehaviorRepository(request []models.UserBehavior) error

	GetAllUserBehaviorRepository() ([]models.UserBehavior, error)
	GetByIdUserBehaviorRepository(id uint) (*models.UserBehavior, error)
	GetByClassroomIdRepository(classRoomID uint) (*models.UserBehavior, error)
	GetByUserIDRepository(userID uint) (*models.UserBehavior, error)
	CreateUserBehaviorRepository(request *models.UserBehavior) error

	DeleteUserBehaviorByUserIDRepository(userID uint) error
}

type userBehaviorRepository struct{ db *gorm.DB }

func (ub *userBehaviorRepository) GetUserBehaviorWithPaginationRepository(request requests.UserBehaviorWithPaginationRequest, paginateRequest trails.PaginateRequest) (*trails.PaginatedResponse, []models.UserBehavior, error) {
	var model []models.UserBehavior
	query := ub.db.Model(&models.UserBehavior{}).
		Preload("User").
		Preload("UserClassroom")

	// Apply conditions based on the request
	switch {
	case request.UserType != "" && request.ClassroomID != 0 && request.SubjectID != 0:
		query = query.Joins("LEFT JOIN users ON user_behaviors.user_id = users.id").
			Joins("LEFT JOIN user_classrooms ON user_behaviors.user_classroom_id = user_classrooms.id").
			Where("users.user_type = ?", request.UserType)
		userClassroomID, err := ub.getDataFromUserClassroomID(request.ClassroomID, request.SubjectID)
		if err != nil {
			return nil, nil, err
		}
		query = query.Where("user_classrooms.id = ?", userClassroomID)
		switch paginateRequest.Sorting {
		case "max":
			query = query.Order("user_behaviors.id DESC")
		case "min":
			query = query.Order("user_behaviors.id ASC")
		default:
			query = query.Order("user_behaviors.id DESC")
		}
	case request.UserType != "":
		query = query.Joins("LEFT JOIN users ON user_behaviors.user_id = users.id").
			Where("users.user_type = ?", request.UserType)
		switch paginateRequest.Sorting {
		case "max":
			query = query.Order("users.id DESC")
		case "min":
			query = query.Order("users.id ASC")
		default:
			query = query.Order("users.id DESC")
		}
	case request.ClassroomID != 0 && request.SubjectID != 0:
		userClassroomID, err := ub.getDataFromUserClassroomID(request.ClassroomID, request.SubjectID)
		if err != nil {
			return nil, nil, err
		}
		query = query.Joins("LEFT JOIN user_classrooms ON user_behaviors.user_classroom_id = user_classrooms.id").
			Where("user_classrooms.id = ?", userClassroomID)
		switch paginateRequest.Sorting {
		case "max":
			query = query.Order("user_behaviors.id DESC")
		case "min":
			query = query.Order("user_behaviors.id ASC")
		default:
			query = query.Order("user_behaviors.id DESC")
		}
	}

	// Handle pagination using PaginationUserBehavior function
	pagination, err := trails.PaginationUserBehavior(query, &model, paginateRequest, true)
	if err != nil {
		return nil, nil, err
	}

	// Ensure empty slice if no data found
	if len(model) == 0 {
		pagination.Items = []models.UserBehavior{}
	}

	return pagination, model, nil
}

func (ub *userBehaviorRepository) getDataFromUserClassroomID(classroomID, subjectID uint) (uint, error) {
	var userClassroom models.UserClassroom
	err := ub.db.Select("id").Where("classroom_id = ? AND subject_id = ?", classroomID, subjectID).First(&userClassroom).Error
	if err != nil {
		return 0, err
	}
	return userClassroom.ID, nil
}

//func (ub *userBehaviorRepository) GetUserBehaviorWithPaginationRepository(request requests.UserBehaviorWithPaginationRequest, paginateRequest trails.PaginateRequest) (*trails.PaginatedResponse, []models.UserBehavior, error) {
//	var model []models.UserBehavior
//	query := ub.db.Model(&models.UserBehavior{}).
//		Joins("LEFT JOIN users ON user_behaviors.user_id = users.id").
//		Joins("LEFT JOIN user_classrooms ON user_behaviors.user_classroom_id = user_classrooms.id")
//
//	// Apply conditions based on the request
//	switch {
//	case request.UserType != "" && request.ClassroomID != 0 && request.SubjectID != 0:
//		query = query.Where("users.user_type = ?", request.UserType)
//		userClassroomID, err := ub.getDataFromUserClassroomID(request.ClassroomID, request.SubjectID)
//		if err != nil {
//			return nil, nil, err
//		}
//		query = query.Where("user_classrooms.id = ?", userClassroomID)
//		switch paginateRequest.Sorting {
//		case "max":
//			query = query.Order("user_behaviors.id DESC")
//		case "min":
//			query = query.Order("user_behaviors.id ASC")
//		default:
//			query = query.Order("user_behaviors.id DESC")
//		}
//	case request.UserType != "":
//		query = query.Where("users.user_type = ?", request.UserType)
//		switch paginateRequest.Sorting {
//		case "max":
//			query = query.Order("users.id DESC")
//		case "min":
//			query = query.Order("users.id ASC")
//		default:
//			query = query.Order("users.id DESC")
//		}
//	case request.ClassroomID != 0 && request.SubjectID != 0:
//		userClassroomID, err := ub.getDataFromUserClassroomID(request.ClassroomID, request.SubjectID)
//		if err != nil {
//			return nil, nil, err
//		}
//		query = query.Where("user_classrooms.id = ?", userClassroomID)
//		switch paginateRequest.Sorting {
//		case "max":
//			query = query.Order("user_behaviors.id DESC")
//		case "min":
//			query = query.Order("user_behaviors.id ASC")
//		default:
//			query = query.Order("user_behaviors.id DESC")
//		}
//	}
//
//	// Print the generated SQL query for debugging
//	sqlQuery := query.Statement.SQL.String()
//	fmt.Println("Generated SQL Query:", sqlQuery)
//
//	// Handle pagination using PaginationUserBehavior function
//	pagination, err := trails.PaginationUserBehavior(query, &model, paginateRequest, true)
//	if err != nil {
//		return nil, nil, err
//	}
//
//	// Ensure empty slice if no data found
//	if len(model) == 0 {
//		pagination.Items = []models.UserBehavior{}
//	}
//
//	return pagination, model, nil
//}
//
//func (ub *userBehaviorRepository) getDataFromUserClassroomID(classroomID, subjectID uint) (uint, error) {
//	var userClassroom models.UserClassroom
//	err := ub.db.Select("id").Where("classroom_id = ? AND subject_id = ?", classroomID, subjectID).First(&userClassroom).Error
//	if err != nil {
//		return 0, err
//	}
//	fmt.Println("userClassroom ID:", userClassroom.ID)
//	return userClassroom.ID, nil
//}

func (ub *userBehaviorRepository) UpdateUserBehaviorRepository(requests []models.UserBehavior) error {
	return ub.db.Transaction(func(tx *gorm.DB) error {
		for _, request := range requests {
			var model models.UserBehavior

			if err := tx.Where("user_id = ? AND user_classroom_id = ?", request.UserID, request.UserClassroomID).
				First(&model).Error; err != nil {
				return err
			}

			// Increment count fields based on request
			if request.StudentCheck {
				model.CountCheck++
			}
			if request.StudentAbsent {
				model.CountAbsent++
			}
			if request.StudentVacation {
				model.CountVacation++
			}
			if request.StudentBreakRule {
				model.CountBreakRule++
			}
			model.UpdatedAt = time.Now()

			// Save the updated user behavior record into the database
			if err := tx.Save(&model).Error; err != nil {
				return err
			}
		}
		return nil
	})
}

func (ub *userBehaviorRepository) InsertUserBehaviorRepository(requests []models.UserBehavior) error {
	// Ensure ID is not manually set to avoid duplicate primary key issues
	for i := range requests {
		requests[i].ID = 0
		// Default values for the other fields are already set in the service
	}

	// Use bulk insert for better performance
	if err := ub.db.Create(&requests).Error; err != nil {
		return err
	}
	return nil
}

func (ub *userBehaviorRepository) CreateUserBehaviorRepository(request *models.UserBehavior) error {

	if err := ub.db.Create(request).Error; err != nil {
		return err
	}
	return nil
}

func (ub *userBehaviorRepository) DeleteUserBehaviorByUserIDRepository(userID uint) error {

	var count int64
	if err := ub.db.Model(&models.UserClassroom{}).Where("user_id = ?", userID).Count(&count).Error; err != nil {
		return err
	}
	if count == 0 {
		return errors.New("user class room not found")
	}

	// Delete the UserClass
	if err := ub.db.Where("user = ?", userID).Delete(&models.UserClassroom{}).Error; err != nil {
		return err
	}
	return nil
}

func (ub *userBehaviorRepository) GetAllUserBehaviorRepository() ([]models.UserBehavior, error) {

	var model []models.UserBehavior
	query := ub.db.Find(&model).Error

	if query != nil {
		return nil, query
	}
	return model, nil
}

func (ub *userBehaviorRepository) GetByClassroomIdRepository(classRoomID uint) (*models.UserBehavior, error) {

	var model models.UserBehavior

	// Execute raw SQL query
	query := ub.db.Raw("SELECT * FROM user_behaviors WHERE classroom_id = ?", classRoomID).Scan(&model).Error
	if query != nil {
		return nil, query
	}
	return &model, nil
}

func (ub *userBehaviorRepository) GetByIdUserBehaviorRepository(id uint) (*models.UserBehavior, error) {

	var model models.UserBehavior

	// Execute raw SQL query
	query := ub.db.Raw("SELECT * FROM user_behaviors WHERE id = ?", id).Scan(&model).Error
	if query != nil {
		return nil, query
	}
	return &model, nil
}

func (ub *userBehaviorRepository) GetByUserIDRepository(userID uint) (*models.UserBehavior, error) {

	var model models.UserBehavior

	// Execute raw SQL query
	query := ub.db.Raw("SELECT * FROM user_behaviors WHERE user_id = ?", userID).Scan(&model).Error
	if query != nil {
		return nil, query
	}
	return &model, nil
}

func NewUserBehaviorRepository(db *gorm.DB) UserBehaviorRepository {
	// db.Migrator().DropTable(models.UserBehavior{})
	// db.AutoMigrate(models.UserBehavior{})
	return &userBehaviorRepository{db: db}
}
