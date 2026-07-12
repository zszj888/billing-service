package billing

import (
	"context"

	"github.com/samz/billing/domain"
)

type BillRepository interface {
	GetOneBill(c context.Context, id int64) (domain.BillEntity, error)
}
type BillService struct {
	billRepository BillRepository
}

func NewService(billRepository BillRepository) *BillService {
	return &BillService{billRepository: billRepository}
}

func (s *BillService) GetByIdSvc(c context.Context, id int64) domain.BillEntity {
	bill, err := s.billRepository.GetOneBill(c, id)
	if err != nil {

	}
	return bill
}
