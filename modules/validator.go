package modules

import (
	"errors"
	"net/mail"

	"github.com/nyaruka/phonenumbers"
)

func ValidateEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

func ValidatePhone(phone string, countryCode string) (string, error) {
	num, err := phonenumbers.Parse(phone, countryCode)
	if err != nil {
		return "", err
	}

	if IsValid := phonenumbers.IsValidNumber(num); !IsValid {
		err := errors.New("phone_number_is_invalid")
		return "", err
	}

	res := phonenumbers.Format(num, phonenumbers.E164)

	return res, nil
}
