package sync

import (
	"context"

	"github.com/spo-iitk/Magicsheet-backend/internal/database"
	"gorm.io/gorm"
)

type RASRepository struct {
	db *gorm.DB
}

func NewRASrepository(db *gorm.DB) *RASRepository {
	return &RASRepository{
		db: db,
	}
}

func (r *RASRepository) GetActiveRecuritmentCycle(ctx context.Context) (*database.RecruitmentCycle, error) {
	var rc database.RecruitmentCycle

	err := r.db.WithContext(ctx).Where("is_active = ?", true).First(&rc).Error

	if err != nil {
		return nil, err
	}

	return &rc, nil
}
