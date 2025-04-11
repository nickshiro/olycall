package rest

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
)

// TODO: refactor errors names
var (
	ErrSyntax                        = errors.New("body contains badly-formed JSON")
	ErrUnexpectedEOF                 = errors.New("body contains badly-formed JSON")
	ErrUnmarshalType                 = errors.New("body contains incorrect JSON type")
	ErrEmptyBody                     = errors.New("body must not be empty")
	ErrUnknownField                  = errors.New("body contains unknown key")
	ErrInvalidUnmarshal              = errors.New("invalid arg passed")
	ErrBodyMustContainOnlySingleJSON = errors.New("body must only contain a single JSON value")
)

func WriteJSON(w http.ResponseWriter, status int, body any) error {
	jsonBytes, err := json.Marshal(body)
	if err != nil {
		return fmt.Errorf("failed to marshal body: %w", err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	if _, err := w.Write(jsonBytes); err != nil {
		return fmt.Errorf("failed to write response: %w", err)
	}

	return nil
}

func ReadJSON(r *http.Request, dst any) error {
	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()

	err := dec.Decode(dst)
	if err != nil {
		var (
			syntaxError           *json.SyntaxError
			unmarshalTypeError    *json.UnmarshalTypeError
			invalidUnmarshalError *json.InvalidUnmarshalError
		)

		switch {
		case errors.As(err, &syntaxError):
			return fmt.Errorf("%w (at character %d)", ErrSyntax, syntaxError.Offset)

		case errors.Is(err, io.ErrUnexpectedEOF):
			return ErrUnexpectedEOF

		case errors.As(err, &unmarshalTypeError):
			if unmarshalTypeError.Field != "" {
				return fmt.Errorf("%w for field %q", ErrUnmarshalType, unmarshalTypeError.Field)
			}

			return fmt.Errorf("%w (at character %d)", ErrUnmarshalType, unmarshalTypeError.Offset)

		case errors.Is(err, io.EOF):
			return ErrEmptyBody

		case strings.HasPrefix(err.Error(), "json: unknown field "):
			fieldName := strings.TrimPrefix(err.Error(), "json: unknown field ")

			return fmt.Errorf("%w %s", ErrUnknownField, fieldName)

		case errors.As(err, &invalidUnmarshalError):
			return fmt.Errorf("%w: %w", ErrInvalidUnmarshal, err)

		default:
			return fmt.Errorf("unknown error: %w", err)
		}
	}

	err = dec.Decode(&struct{}{})
	if err != io.EOF {
		return ErrBodyMustContainOnlySingleJSON
	}

	return nil
}
