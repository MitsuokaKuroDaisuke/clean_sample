package interactor

import (
	"context"
	"fmt"
	"src/domain"
	"src/domain/entity"
	"strconv"
)

// User Userビジネスロジック
type Member struct {
	MemberRepository         entity.MemberRepository
	MemberDetailRepository   entity.MemberDetailRepository
	MemberMatchRepository    entity.MemberMatchRepository
	OriginalMemberRepository domain.OriginalMemberRepository
	Transaction              domain.Transaction
}

func (interactor *Member) GetMember(id string) (entity.Member, error) {
	mid, err := strconv.Atoi(id)
	mem := interactor.MemberRepository.AdminDataByID(mid)

	if err := interactor.Transaction.DoInTx(func(ctx context.Context) error {
		interactor.MemberDetailRepository.Create(ctx)
		interactor.MemberMatchRepository.Create(ctx)
		return fmt.Errorf("何かのエラー発生でロールバック")
	}); err != nil {
		return mem, err
	}
	return mem, err
}

func (interactor *Member) GetOriginalMember(id string) (domain.OriginalMember, error) {
	mid, err := strconv.Atoi(id)
	mem := interactor.OriginalMemberRepository.OriginalDataByID(mid)
	return mem, err
}
