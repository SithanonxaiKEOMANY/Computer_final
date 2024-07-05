package repositories

import (
	"errors"
	"fmt"
	"go_starter/models"

	"gorm.io/gorm"
)

type ClassRoomRepository interface {
	GetAllClassRoomRepositories() ([]models.Classroom, error)
	GetClassRoomByIdRepository(id uint) (*models.Classroom, error)
	GetClassRoomByClassYearRepository(classYear string) (*models.Classroom, error)
	CreateClassRoomRepository(request *models.Classroom) error
	UpdateClassRoomRepository(request *models.Classroom) error
	DeleteClassRoomRepository(id uint) error
	
}

type classroomRepository struct{ db *gorm.DB }

func (c *classroomRepository) CheckClassRoomCodeAlreadyHas(classRoomCode string) (bool, error) {
	var count int64
	query := c.db.Model(&models.Classroom{}).Where("class_room_code = ?", classRoomCode).Count(&count)
	if query.Error != nil {
		return false, query.Error
	}
	return count > 0, nil
}

func (c *classroomRepository) checkClassNameRepository(className, classYear int, major string) (bool, error) {
	var count int64
	err := c.db.Model(&models.Classroom{}).Where("class_name = ? AND class_year = ? AND major = ?", className, classYear, major).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func (c *classroomRepository) CreateClassRoomRepository(request *models.Classroom) error {
	// Check if class_name already exists
	exists, err := c.checkClassNameRepository(request.ClassName, request.ClassYear, request.Major)
	if err != nil {
		return err
	}

	if exists {
		return fmt.Errorf("class_name and class_year already exists")
	}
	if err = c.db.Create(request).Error; err != nil {
		return err
	}
	return nil
}

func (c *classroomRepository) DeleteClassRoomRepository(id uint) error {
	query := models.Classroom{ID: id}
	if err := c.db.Where("id = ?", id).Delete(&query).Error; err != nil {
		return err
	}
	return nil
}

func (c *classroomRepository) GetAllClassRoomRepositories() ([]models.Classroom, error) {
	var model []models.Classroom
	query := c.db.Find(&model).Error
	if query != nil {
		return nil, query
	}
	return model, nil
}

func (c *classroomRepository) GetClassRoomByIdRepository(id uint) (*models.Classroom, error) {
	var model models.Classroom
	query := c.db.Raw("SELECT *FROM classrooms WHERE id = ?", id).Scan(&model).Error

	if query != nil {
		return nil, query
	}
	return &model, nil
}

func (c *classroomRepository) GetClassRoomByClassYearRepository(classYear string) (*models.Classroom, error) {
	var model models.Classroom
	query := c.db.Raw("SELECT *FROM classrooms WHERE class_year = ?", classYear).Scan(&model).Error

	if query != nil {
		return nil, query
	}
	return &model, nil
}

// UpdateClassRoomRepository updates a classroom record in the database.
func (c *classroomRepository) UpdateClassRoomRepository(request *models.Classroom) error {
	query := c.db.Model(&models.Classroom{}).
		Where("id = ?", request.ID).Updates(request)

	if query.Error != nil {
		return fmt.Errorf("error updating classroom: %w", query.Error)
	}
	if query.RowsAffected == 0 {
		return errors.New("no classroom found to update")
	}
	return nil
}


func NewRoomRepository(db *gorm.DB) ClassRoomRepository {
	//db.Migrator().DropTable(models.ClassRoom{})
	//db.AutoMigrate(models.ClassRoom{})
	return &classroomRepository{db: db}
}
