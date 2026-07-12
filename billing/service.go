package billing

import (
	"context"

	"github.com/samz/billing/domain"
)

type BillRepository interface {
	GetOneBill(c context.Context, id int64) (domain.BillEntity, error)
	Save(c context.Context, bill *domain.BillEntity) error
}
type BillService struct {
	billRepository BillRepository
}

func (s *BillService) Save(c context.Context, bill *domain.BillEntity) error {
	return s.billRepository.Save(c, bill)
}

func NewService(billRepository BillRepository) *BillService {
	return &BillService{billRepository: billRepository}
}

func (s *BillService) GetByIdSvc(c context.Context, id int64) (domain.BillEntity, error) {
	return s.billRepository.GetOneBill(c, id)
}
