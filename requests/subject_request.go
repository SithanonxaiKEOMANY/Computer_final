package requests

type SubjectRequest struct {
	SubjectCode string `json:"subject_code" validate:"required"`
}

type InsertSubjectRequest struct {
	SubjectCode string `json:"subject_code" validate:"required"`
	SubjectName string `json:"subject_name" validate:"required"`
}

type UpdateSubjectRequest struct {
	SubjectCode string `json:"subject_code" validate:"required"`
	SubjectName string `json:"subject_name" validate:"required"`
}

type DeleteSubjectRequest struct {
	Id          uint   `json:"id" validate:"required"`
	SubjectCode string `json:"subject_code" validate:"required"`
}
type SubjectCodeRequest struct {
	Id uint `json:"id" validate:"required"`
}
