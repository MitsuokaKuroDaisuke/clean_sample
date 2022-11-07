package presenters

import (
	"net/http"
	"src/domain"
	"src/domain/entity"

	"github.com/labstack/echo/v4"
)

type MemberOutputPort struct {
}

func (pr *MemberOutputPort) Member(c echo.Context, member entity.Member) error {

	data := struct {
		Member entity.Member
	}{
		Member: member,
	}
	return c.Render(http.StatusOK, "member", data)

}

func (pr *MemberOutputPort) OriginalMember(c echo.Context, member domain.OriginalMember) error {

	data := struct {
		OriginalMember domain.OriginalMember
	}{
		OriginalMember: member,
	}
	return c.Render(http.StatusOK, "original_member", data)

}
