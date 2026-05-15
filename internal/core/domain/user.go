package domain

import (
	"fmt"
	"regexp"

	core_errors "github.com/turtlesafik-beep/GolangToDO/internal/core/errors"
)

type User struct {
	ID      int
	Version int

	FullName    string
	PhoneNumber *string
}

func NewUser(
	id int,
	version int,
	fullName string,
	phoneNumber *string,
) User {
	return User{
		ID:          id,
		Version:     version,
		FullName:    fullName,
		PhoneNumber: phoneNumber,
	}
}

func NewUserUninitialized(
	fullName string,
	phoneNumber *string,
) User {
	return NewUser(
		UninitializedID,
		UninitializedVersion,
		fullName,
		phoneNumber,
	)
}

func (u *User) Validate() error {
	fullNameLength := len([]rune(u.FullName))
	if fullNameLength < 3 || fullNameLength > 100 {
		return fmt.Errorf(
			"invalid `FullName` len: %d: %w",
			fullNameLength,
			core_errors.ErrInavildArgument,
		)
	}

	if u.PhoneNumber != nil {
		phoneNumberLen := len([]rune(*u.PhoneNumber))
		if phoneNumberLen < 10 || phoneNumberLen > 15 {
			return fmt.Errorf(
				"invalid `PhoneNumber` len: %d, %w",
				phoneNumberLen,
				core_errors.ErrInavildArgument,
			)
		}

		re := regexp.MustCompile(`^\+[0-9]+$`)

		if !re.MatchString(*u.PhoneNumber) {
			return fmt.Errorf(
				"invalid `PhoneNumber` format: %w",
				core_errors.ErrInavildArgument,
			)
		}
	}

	return nil
}

type UserPath struct {
	FullName    Nullable[string]
	PhoneNumber Nullable[string]
}

func (p *UserPath) Validate() error {
	if p.FullName.Set && p.FullName.Value == nil {
		return fmt.Errorf(
			"FullName can't be pathed to null: %w",
			core_errors.ErrInavildArgument,
		)
	}

	return nil
}

func (u *User) ApplyPatch(path UserPath) error {
	if err := path.Validate(); err != nil {
		return fmt.Errorf("validate user path: %w", err)
	}

	tmp := *u

	if path.FullName.Set {
		tmp.FullName = *path.FullName.Value
	}

	if path.PhoneNumber.Set {
		tmp.PhoneNumber = path.PhoneNumber.Value
	}

	if err := tmp.Validate(); err != nil {
		return fmt.Errorf("validate pathed user: %w", err)
	}

	*u = tmp

	return nil
}
