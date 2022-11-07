package interactor

import (
	"src/domain"
	"src/domain/entity"
	"strconv"
)

// User Userビジネスロジック
type Member struct {
	MemberRepository         entity.MemberRepository
	OriginalMemberRepository domain.OriginalMemberRepository
}

func (interactor *Member) GetMember(id string) (entity.Member, error) {
	mid, err := strconv.Atoi(id)
	mem := interactor.MemberRepository.AdminDataByID(mid)
	return mem, err
}

func (interactor *Member) GetOriginalMember(id string) (domain.OriginalMember, error) {
	mid, err := strconv.Atoi(id)
	mem := interactor.OriginalMemberRepository.OriginalDataByID(mid)
	return mem, err
}
