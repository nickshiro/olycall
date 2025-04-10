package core

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

var UserNameRule = []validation.Rule{
	validation.Length(1, 32),
}
