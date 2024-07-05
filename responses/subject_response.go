package responses

type SubjectResponse struct {
	ID          uint   `json:"id"`
	SubjectCode string `json:"subject_code"`
	SubjectName string `json:"subject_name"`
}

type SubjectMessageResponse struct {
	Message string `json:"message"`
	Status  bool   `json:"status"`
}
