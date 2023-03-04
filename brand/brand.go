package brand

import (
	"fleet-backend/graph/model"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BrandRequest struct {
	ID          primitive.ObjectID `json:"Id" bson:"_id"`
	Name        string             `json:"name" bson:"name"`
	Description string             `json:"description" bson:"description"`
	URLName     string             `json:"urlName" bson:"urlName"`
	Image       string             `json:"image" bson:"image"`
	CreateAt    *time.Time         `json:"createAt" bson:"createAt"`
	UpdateAt    time.Time          `json:"updateAt" bson:"updateAt"`
	CreateBy    *string            `json:"createBy" bson:"createBy"`
	UpdateBy    *string            `json:"updateBy" bson:"updateBy"`
	Active      int                `json:"active" bson:"active"`
}

type Store interface {
	GetByID(primitive.ObjectID) (*model.Brands, error)
	GetAll() ([]*model.Brands, error)
	Create(BrandRequest) error
	Update(BrandRequest) error
}
