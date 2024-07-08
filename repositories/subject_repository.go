package repositories

import (
	"errors"
	"fmt"
	"go_starter/models"

	"gorm.io/gorm"
)

type SubjectRepository interface {
	FilterSubjectBySubjectCodeRepository(subjectCode string) (*models.Subject, error)
	CreateSubjectRepository(model models.Subject) error
	GetSubjectRepository() ([]models.Subject, error)
	UpdateSubjectRepository(model models.Subject) error
	DeleteSubjectRepository(id uint) error
	CheckSubjectCodeAlreadyHas(subjectCode string) (bool, error)
}

type subjectRepository struct {
	db *gorm.DB
}

// CheckSubjectCodeAlreadyHas implements SubjectRepository.
func (s *subjectRepository) CheckSubjectCodeAlreadyHas(subjectCode string) (bool, error) {

	var count int64
	query := s.db.Model(&models.Subject{}).Where("subject_code = ?", subjectCode).Count(&count)
	if query.Error != nil {
		return false, query.Error
	}
	return count > 0, nil
}

func (s *subjectRepository) FilterSubjectBySubjectCodeRepository(subjectCode string) (*models.Subject, error) {
	var model models.Subject
	if err := s.db.Where("subject_code=?", subjectCode).First(&model).Error; err != nil {
		return nil, err
	}
	return &model, nil
}

func (s *subjectRepository) DeleteSubjectRepository(id uint) error {

	query := models.Subject{ID: id}

	if err := s.db.Where("id = ?", id).Delete(&query).Error; err != nil {
		return err
	}

	return nil
}

func (s *subjectRepository) UpdateSubjectRepository(model models.Subject) error {

	query := s.db.Model(&models.Subject{}).Where("subject_code = ?", model.SubjectCode).Updates(model)

	if query.Error != nil {
		return query.Error
	}
	if query.RowsAffected == 0 {
		return errors.New("no subject found")
	}
	return nil
}

func (s *subjectRepository) checkSubjectRepository(subjectCode string) (bool, error) {
	var count int64
	err := s.db.Model(&models.Subject{}).Where("subject_code = ?", subjectCode).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func (s *subjectRepository) CreateSubjectRepository(model models.Subject) error {
	checkSubjectCode, err := s.checkSubjectRepository(model.SubjectCode)
	if err != nil {
		return err
	}

	if checkSubjectCode {
		return fmt.Errorf("subject code already had been created")
	}

	if err = s.db.Create(&model).Error; err != nil {
		return err
	}
	return nil
}

func (s *subjectRepository) GetSubjectRepository() ([]models.Subject, error) {
	var model []models.Subject
	query := s.db.Find(&model).Error
	if query != nil {
		return nil, query
	}
	return model, nil
}

func NewSubjectRepository(db *gorm.DB) SubjectRepository {
	//db.Migrator().DropTable(models.Subject{})
	//db.AutoMigrate(models.Subject{})
	return &subjectRepository{db: db}
}
