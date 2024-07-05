package services

import (
	"go_starter/models"
	"go_starter/repositories"
	"go_starter/requests"
	"go_starter/responses"
	"time"

	"github.com/pkg/errors"
)

type SubjectService interface {
	CreateSubjectService(request requests.InsertSubjectRequest) (responses.SubjectMessageResponse, error)
	FilterSubjectBySubjectCodeService(request requests.SubjectRequest) (*responses.SubjectResponse, error)
	GetSubjectService() ([]responses.SubjectResponse, error)
	UpdateSubjectService(request requests.UpdateSubjectRequest) (*responses.SubjectMessageResponse, error)
	DeleteSubjectService(request requests.SubjectCodeRequest) (*responses.SubjectMessageResponse, error)
}

type subjectService struct {
	repositorySubject repositories.SubjectRepository
}

func (s *subjectService) CreateSubjectService(request requests.InsertSubjectRequest) (responses.SubjectMessageResponse, error) {
	subjectModel := models.Subject{
		SubjectCode: request.SubjectCode,
		SubjectName: request.SubjectName,
	}
	err := s.repositorySubject.CreateSubjectRepository(subjectModel)
	if err != nil {
		return responses.SubjectMessageResponse{
			Message: "failed to create subject ",
			Status:  false,
		}, err
	}
	return responses.SubjectMessageResponse{
		Message: "success",
		Status:  true,
	}, nil
}

func (s *subjectService) FilterSubjectBySubjectCodeService(request requests.SubjectRequest) (*responses.SubjectResponse, error) {
	getSubjectData, err := s.repositorySubject.FilterSubjectBySubjectCodeRepository(request.SubjectCode)
	if err != nil {
		return nil, err
	}
	response := responses.SubjectResponse{
		ID:          getSubjectData.ID,
		SubjectCode: getSubjectData.SubjectCode,
		SubjectName: getSubjectData.SubjectName,
	}
	return &response, err
}

func (s *subjectService) GetSubjectService() ([]responses.SubjectResponse, error) {
	getSubjectData, err := s.repositorySubject.GetSubjectRepository()
	if err != nil {
		return nil, err
	}
	if getSubjectData == nil {
		return []responses.SubjectResponse{}, nil
	}
	var response []responses.SubjectResponse
	for _, data := range getSubjectData {
		response = append(response, responses.SubjectResponse{
			ID:          data.ID,
			SubjectCode: data.SubjectCode,
			SubjectName: data.SubjectName,
		})
	}
	return response, err
}

// UpdateSubjectService implements SubjectService.
func (s *subjectService) UpdateSubjectService(request requests.UpdateSubjectRequest) (*responses.SubjectMessageResponse, error) {

	model := models.Subject{
		SubjectCode: request.SubjectCode,
		SubjectName: request.SubjectName,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	if err := s.repositorySubject.UpdateSubjectRepository(model); err != nil {

		return nil, err
	}
	response := &responses.SubjectMessageResponse{Message: "Success"}

	return response, nil
}

// DeleteSubjectService implements SubjectService.
func (s *subjectService) DeleteSubjectService(request requests.SubjectCodeRequest) (*responses.SubjectMessageResponse, error) {

	if request.Id == 0 {
		return nil, errors.New(" ID cannot be empty")
	}
	err := s.repositorySubject.DeleteSubjectRepository(request.Id)
	if err != nil {
		return nil, err
	}

	// If successful, return a success message response
	response := &responses.SubjectMessageResponse{Message: "success"}
	return response, nil
}

func NewSubjectService(
	repositorySubject repositories.SubjectRepository,
) SubjectService {
	return &subjectService{
		repositorySubject: repositorySubject,
	}
}
