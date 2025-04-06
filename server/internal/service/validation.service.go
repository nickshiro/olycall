package service

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

var NoteTitleRule = []validation.Rule{
	validation.Length(0, 32),
}
