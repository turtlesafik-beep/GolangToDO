package user_service

import (
	"context"

	"github.com/turtlesafik-beep/GolangToDO/internal/core/domain"
)

func (s *UsersService) CreateUser(
	ctx context.Context,
	user domain.User,
) (domain.User, error) {

}
