package trails

import (
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type PaginateRequest struct {
	Item    int    `json:"item" validate:"required"`
	Page    int    `json:"page" validate:"required"`
	Sorting string `json:"sorting" validate:"required"`
}

type PaginatedResponse struct {
	TotalPages  int         `json:"total_pages"`
	PerPage     int         `json:"per_page"`
	CurrentPage int         `json:"current_page"`
	Sorting     string      `json:"sorting"`
	Items       interface{} `json:"items"`
}

func PaginationUserBehavior(db *gorm.DB, model interface{}, paginate PaginateRequest, preloadAssociations bool) (*PaginatedResponse, error) {
	if paginate.Item <= 0 || paginate.Page <= 0 {
		return nil, errors.New("item and page must be positive integers")
	}

	var total int64
	db.Model(model).Count(&total)
	totalPages := (int(total) + paginate.Item - 1) / paginate.Item
	offset := (paginate.Page - 1) * paginate.Item

	if paginate.Page > totalPages {
		return &PaginatedResponse{
			TotalPages:  1,
			PerPage:     0,
			CurrentPage: 1,
			Sorting:     paginate.Sorting,
			Items:       []interface{}{},
		}, nil
	}

	query := db.Model(model).Limit(paginate.Item).Offset(offset)

	if preloadAssociations {
		query = query.Preload(clause.Associations)
	}

	//switch paginate.Sorting {
	//case "max":
	//	query = query.Order("user_behaviors.id DESC")
	//case "min":
	//	query = query.Order("user_behaviors.id ASC")
	//default:
	//	query = query.Order("user_behaviors.id DESC")
	//}

	result := query.Find(model)
	if result.Error != nil {
		return nil, result.Error
	}

	pagination := &PaginatedResponse{
		TotalPages:  totalPages,
		PerPage:     paginate.Item,
		CurrentPage: paginate.Page,
		Sorting:     paginate.Sorting,
		Items:       model,
	}
	return pagination, nil
}

func PaginationData(db *gorm.DB, model interface{}, paginate PaginateRequest, preloadAssociations bool) (*PaginatedResponse, error) {
	if paginate.Item <= 0 || paginate.Page <= 0 {
		return nil, errors.New("item and page must be positive integers")
	}

	var total int64
	db.Model(model).Count(&total)
	totalPages := (int(total) + paginate.Item - 1) / paginate.Item
	offset := (paginate.Page - 1) * paginate.Item

	//if paginate.Page > totalPages {
	//	return nil, errors.New("current_page exceeds total_page")
	//}

	if paginate.Page > totalPages {
		return &PaginatedResponse{
			TotalPages:  1,
			PerPage:     0,
			CurrentPage: 1,
			Sorting:     paginate.Sorting,
			Items:       []interface{}{},
		}, nil
	}

	query := db.Model(model).Limit(paginate.Item).Offset(offset)

	if preloadAssociations {
		query = query.Preload(clause.Associations)
	}

	switch paginate.Sorting {
	case "max":
		query = query.Order("id DESC")
	case "min":
		query = query.Order("id ASC")
	default:
		query = query.Order("id DESC")
	}

	result := query.Find(model)
	if result.Error != nil {
		return nil, result.Error
	}

	pagination := &PaginatedResponse{
		TotalPages:  totalPages,
		PerPage:     paginate.Item,
		CurrentPage: paginate.Page,
		Sorting:     paginate.Sorting,
		Items:       model,
	}
	return pagination, nil
}

//type PaginateRequest struct {
//	Item     int    `json:"item" validate:"required"`
//	Page     int    `json:"page" validate:"required"`
//	Sorting  string `json:"sorting" validate:"required"`
//	UserType string `json:"user_type" validate:"required"`
//}
//
//type PaginatedResponse struct {
//	TotalPages  int         `json:"total_pages"`
//	PerPage     int         `json:"per_page"`
//	CurrentPage int         `json:"current_page"`
//	Sorting     string      `json:"sorting"`
//	UserType    string      `json:"user_type"`
//	Items       interface{} `json:"items"`
//}
//
//func PaginationData(db *gorm.DB, model interface{}, paginate PaginateRequest) (*PaginatedResponse, error) {
//	if paginate.Item <= 0 || paginate.Page <= 0 {
//		return nil, errors.New("per_page and current_page must be positive integers")
//	}
//
//	var total int64
//	db.Model(model).Count(&total)
//	totalPages := (int(total) + paginate.Item - 1) / paginate.Item
//	offset := (paginate.Page - 1) * paginate.Item
//
//	if paginate.Page > totalPages {
//		return nil, errors.New("current_page exceeds total_page")
//	}
//
//	// Check if the current page is beyond the total pages
//	if paginate.Page > totalPages {
//		pagination := &PaginatedResponse{
//			TotalPages:  totalPages,
//			PerPage:     paginate.Item,
//			CurrentPage: paginate.Page,
//			Sorting:     paginate.Sorting,
//			UserType:    paginate.UserType,
//			Items:       []interface{}{}, // Return an empty list instead of null
//		}
//		return pagination, nil
//	}
//
//	query := db.Preload(clause.Associations).Limit(paginate.Item).Offset(offset)
//
//	result := query.Find(model)
//	if result.Error != nil {
//		return nil, result.Error
//	}
//
//	pagination := &PaginatedResponse{
//		TotalPages:  totalPages,
//		PerPage:     paginate.Item,
//		CurrentPage: paginate.Page,
//		Sorting:     paginate.Sorting,
//		UserType:    paginate.UserType,
//		Items:       model,
//	}
//	return pagination, nil
//}
