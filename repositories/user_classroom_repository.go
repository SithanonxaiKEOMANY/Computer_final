package repositories

import (
	"errors"
	"go_starter/models"
	"gorm.io/gorm/clause"

	"gorm.io/gorm"
)

type UserClassRoomRepository interface {
	GetUserClassroomsRepository(id uint, userType string) ([]models.UserClassroom, error)
	GetUserClassroomForInsertUserBehaviorRepository(classroomID uint, subjectID uint) (*models.UserClassroom, error)
	DeleteUserClassRoomAssociation(userID uint, classroomID uint) error
	CheckUserClassRoomExistsRepository(userID uint, classRoomID uint, subjectID uint) (bool, error)
	CreateUserClassroomRepository(request []models.UserClassroom) error
	GetUserClassroomByStudentTypeRepository(classroomID uint, userType string, subjectID uint) ([]models.UserClassroom, error)

	//GetUserClassroomByStudentTypeRepository(classroomID int, userType string) ([]models.UserClassroom, error)
	GetClassroomByTeacherRepository(userId int, userType string) ([]models.UserClassroom, error)

	GetAllUserClassRoomRepository() ([]models.UserClassroom, error)
	GetByIdUserClassRoomRepository(id uint) (*models.UserClassroom, error)
	GetByClassroomIdRepository(classRoomID uint) (*models.UserClassroom, error)
	GetByUserIDRepository(userID uint) (*models.UserClassroom, error)

	UpdateUserClassRoomRepository(request *models.UserClassroom) error
	DeleteUserClassRoomIDRepository(id uint) error
}

type userClassRoomRepository struct{ db *gorm.DB }

func (uc *userClassRoomRepository) GetUserClassroomsRepository(id uint, userType string) ([]models.UserClassroom, error) {
	var userClassrooms []models.UserClassroom

	err := uc.db.Preload("User", "id = ? AND user_type = ?", id, userType).
		Preload("Classroom", func(db *gorm.DB) *gorm.DB {
			return db.Select("id, class_year, major, class_name")
		}).
		Preload("Subject", func(db *gorm.DB) *gorm.DB {
			return db.Select("id, subject_code, subject_name")
		}).
		Where("user_id = ?", id).
		Find(&userClassrooms).Error

	if err != nil {
		return nil, err
	}

	return userClassrooms, nil
}

func (uc *userClassRoomRepository) GetUserClassroomForInsertUserBehaviorRepository(classroomID uint, subjectID uint) (*models.UserClassroom, error) {
	var userClassroom models.UserClassroom
	err := uc.db.Select("id").Where("classroom_id = ? AND subject_id = ?", classroomID, subjectID).First(&userClassroom).Error
	if err != nil {
		return nil, err
	}
	return &userClassroom, nil
}

func (uc *userClassRoomRepository) DeleteUserClassRoomAssociation(userID uint, classroomID uint) error {
	result := uc.db.Where("user_id = ? AND classroom_id = ?", userID, classroomID).Delete(&models.UserClassroom{})
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (uc *userClassRoomRepository) CheckUserClassRoomExistsRepository(userID uint, classRoomID uint, subjectID uint) (bool, error) {
	var count int64
	if err := uc.db.Model(&models.UserClassroom{}).
		Where("user_id = ? AND classroom_id =? AND subject_id = ?", userID, classRoomID, subjectID).
		Count(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}

func (uc *userClassRoomRepository) CreateUserClassroomRepository(request []models.UserClassroom) error {

	if err := uc.db.Clauses(clause.OnConflict{
		DoNothing: true,
	}).Create(&request).Error; err != nil {
		return err
	}
	return nil
}

func (uc *userClassRoomRepository) GetUserClassroomByStudentTypeRepository(classroomID uint, userType string, subjectID uint) ([]models.UserClassroom, error) {
	var userClassrooms []models.UserClassroom

	query := uc.db.Preload("Classroom").Preload("Subject")

	if userType == "" {
		// If userType is empty, list all data and order by user_type with 'teacher' first
		query = query.Preload("User").
			Joins("JOIN users ON users.id = user_classrooms.user_id").
			Where("user_classrooms.classroom_id = ? AND user_classrooms.subject_id = ?", classroomID, subjectID).
			Order("CASE WHEN users.user_type = 'teacher' THEN 1 ELSE 2 END")
	} else {
		// If userType is provided, filter by userType
		query = query.Preload("User").
			Joins("JOIN users ON users.id = user_classrooms.user_id").
			Where("user_classrooms.classroom_id = ? AND user_classrooms.subject_id = ? AND users.user_type = ?", classroomID, subjectID, userType)
	}

	if err := query.Find(&userClassrooms).Error; err != nil {
		return nil, err
	}

	return userClassrooms, nil

}

//func (uc *userClassRoomRepository) GetUserClassroomByStudentTypeRepository(classroomID int, userType string) ([]models.UserClassroom, error) {
//	var model []models.UserClassroom
//
//	// Construct the raw SQL query
//	sqlQuery := `
//        SELECT user_classrooms.*, classrooms.*, users.* FROM user_classrooms
//        JOIN classrooms ON user_classrooms.classroom_id = classrooms.id
//        JOIN users ON user_classrooms.user_id = users.id
//        WHERE user_classrooms.classroom_id = ? AND users.user_type = ?
//    `
//
//	// Execute the raw SQL query
//	query := uc.db.Raw(sqlQuery, classroomID, userType).Scan(&model)
//	if query.Error != nil {
//		return nil, query.Error
//	}
//
//	return model, nil
//}

//func (uc *userClassRoomRepository) GetUserClassroomByStudentTypeRepository(classroomID int, userType string) ([]models.UserClassroom, error) {
//	var model []models.UserClassroom
//	query := uc.db.Preload("Classroom").Preload("User").
//		Joins("JOIN users ON user_classrooms.user_id = users.id").
//		Where("user_classrooms.classroom_id = ? AND users.user_type = ?", classroomID, userType).
//		Find(&model).Error
//	if query != nil {
//		return nil, query
//	}
//	return model, nil
//}

//func (uc *userClassRoomRepository) GetUserClassroomByStudentTypeRepository(classroomID int, userType string) ([]models.UserClassroom, error) {
//	var model []models.UserClassroom
//	query := uc.db.Preload("Classroom").Preload("User").Where("classroom_id=? AND user_type=?", classroomID, userType).Find(&model).Error
//	if query != nil {
//		return nil, query
//	}
//	return model, nil
//}

func (uc *userClassRoomRepository) GetClassroomByTeacherRepository(userId int, userType string) ([]models.UserClassroom, error) {
	var model []models.UserClassroom
	query := uc.db.Preload("ClassRoom").Preload("User").Where("user_id=? AND user_type=?", userId, userType).Find(&model).Error
	if query != nil {
		return nil, query
	}
	return model, nil
}

func (uc *userClassRoomRepository) DeleteUserClassRoomByUserIDRepository(userID uint) error {

	var count int64
	if err := uc.db.Model(&models.UserClassroom{}).Where("user_id = ?", userID).Count(&count).Error; err != nil {
		return err
	}
	if count == 0 {
		return errors.New("user class room not found")
	}

	// Delete the UserClass
	if err := uc.db.Where("user = ?", userID).Delete(&models.UserClassroom{}).Error; err != nil {
		return err
	}
	return nil
}

func (uc *userClassRoomRepository) GetAllUserClassRoomRepository() ([]models.UserClassroom, error) {

	var model []models.UserClassroom
	query := uc.db.Find(&model).Error

	if query != nil {
		return nil, query
	}
	return model, nil
}

func (uc *userClassRoomRepository) GetByClassroomIdRepository(classRoomID uint) (*models.UserClassroom, error) {

	var model models.UserClassroom

	// Execute raw SQL query
	query := uc.db.Raw("SELECT * FROM user_class_rooms WHERE class_room_id = ?", classRoomID).Scan(&model).Error
	if query != nil {
		return nil, query
	}
	return &model, nil
}

func (uc *userClassRoomRepository) GetByUserIDRepository(userID uint) (*models.UserClassroom, error) {

	var model models.UserClassroom

	// Execute raw SQL query
	query := uc.db.Raw("SELECT * FROM user_class_rooms WHERE user_id = ?", userID).Scan(&model).Error
	if query != nil {
		return nil, query
	}
	return &model, nil
}

func (uc *userClassRoomRepository) GetByIdUserClassRoomRepository(id uint) (*models.UserClassroom, error) {

	var model models.UserClassroom

	// Execute raw SQL query
	query := uc.db.Raw("SELECT * FROM user_class_rooms WHERE id = ?", id).Scan(&model).Error
	if query != nil {
		return nil, query
	}
	return &model, nil
}

func (uc *userClassRoomRepository) UpdateUserClassRoomRepository(request *models.UserClassroom) error {

	query := uc.db.Model(&models.UserClassroom{}).Where("user_id = ?", request.UserID).Updates(request)
	if query.Error != nil {
		return query.Error
	}
	if query.RowsAffected == 0 {
		return errors.New("no user_id found")
	}
	return nil
}

func (uc *userClassRoomRepository) DeleteUserClassRoomIDRepository(id uint) error {

	query := models.UserClassroom{ID: id}

	if err := uc.db.Where("id = ?", id).Delete(&query).Error; err != nil {
		return err
	}

	return nil
}

func NewUserClassRoom(db *gorm.DB) UserClassRoomRepository {
	//db.Migrator().DropTable(models.UserClassroom{}, models.User{}, models.ClassRoom{}, models.UserBehavior{})
	//db.AutoMigrate(models.UserClassroom{})

	return &userClassRoomRepository{db: db}
}
