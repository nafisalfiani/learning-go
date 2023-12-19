// Ungraded challenge 1
package main

import (
	"errors"
	"fmt"
	"reflect"
	"regexp"
	"strconv"
)

// Error adalah tipe error yang digunakan untuk mengindikasikan kesalahan validasi.
type Error struct {
	Field string
	Err   error
}

func (e *Error) Error() string {
	return e.Field + ": " + e.Err.Error()
}

// StructTag adalah fungsi dari package validate yang digunakan untuk melakukan proses validasi object struct terhadap tag yang ada di struct tersebut.
func StructTag(data interface{}) error {
	value := reflect.ValueOf(data)
	if value.Kind() == reflect.Ptr {
		value = value.Elem()
	}

	if value.Kind() != reflect.Struct {
		return errors.New("Data harus berupa struct")
	}

	typ := value.Type()

	for i := 0; i < value.NumField(); i++ {
		field := value.Field(i)
		fieldType := typ.Field(i)
		tag := fieldType.Tag.Get("validate")

		if tag != "" {
			if tag == "required" {
				if isZero(field) {
					return &Error{fieldType.Name, errors.New("field wajib diisi")}
				}
			} else if tag == "email" {
				if field.Kind() == reflect.String && !isValidEmail(field.String()) {
					return &Error{fieldType.Name, errors.New("format email tidak valid")}
				}
			} else if tag == "min" || tag == "max" {
				if field.Kind() == reflect.Int {
					min, max := getMinMaxValues(tag, fieldType)
					val := field.Int()
					if tag == "min" && val < min {
						return &Error{fieldType.Name, errors.New("nilai terlalu kecil")}
					} else if tag == "max" && val > max {
						return &Error{fieldType.Name, errors.New("nilai terlalu besar")}
					}
				}
			} else if tag == "minLen" || tag == "maxLen" {
				if field.Kind() == reflect.String {
					minLen, maxLen := getMinMaxValues(tag, fieldType)
					val := int64(len(field.String()))
					if tag == "minLen" && val < minLen {
						return &Error{fieldType.Name, errors.New("panjang terlalu pendek")}
					} else if tag == "maxLen" && val > maxLen {
						return &Error{fieldType.Name, errors.New("panjang terlalu panjang")}
					}
				}
			}
		}
	}

	return nil
}

// isValidEmail adalah fungsi yang menggunakan regex untuk memeriksa validitas alamat email.
func isValidEmail(email string) bool {
	// Gunakan regex yang sesuai dengan format email yang diinginkan
	regex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	matched, err := regexp.MatchString(regex, email)
	if err != nil {
		return matched
	}

	return matched
}

// getMinMaxValues adalah fungsi yang mengambil nilai min atau max dari tag struct.
func getMinMaxValues(tag string, fieldType reflect.StructField) (int64, int64) {
	tagValue, ok := fieldType.Tag.Lookup(tag)
	if !ok {
		return 0, 0
	}

	value, err := strconv.ParseInt(tagValue, 10, 64)
	if err != nil {
		return 0, 0
	}

	return value, value
}

// isZero adalah fungsi yang memeriksa apakah nilai field adalah alias atau Go default value.
func isZero(field reflect.Value) bool {
	return reflect.DeepEqual(field.Interface(), reflect.Zero(field.Type()).Interface())
}

type Data struct {
	RequiredField string `validate:"required"`
	EmailField    string `validate:"email"`
	IntField      int    `validate:"min:10,max:100"`
	StringField   string `validate:"minLen:5,maxLen:20"`
}

func main() {
	data := Data{
		RequiredField: "nilainya",
		EmailField:    "contoh@email.com",
		IntField:      50,
		StringField:   "PanjangString",
	}

	err := StructTag(&data)
	if err != nil {
		if vErr, ok := err.(*Error); ok {
			fmt.Println("Validasi error:", vErr.Error())
		}
	}

	fmt.Println("data is successfully validated")
}
