package forms

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"unicode/utf8"
)

type Form struct {
	url.Values
	ErrorsMap errors
}

func New(data url.Values) *Form {
	return &Form{
		data,
		map[string][]string{},
	}
}

func (f *Form) Required(fields ...string) {
	for _, field := range fields {
		value := f.Get(field)
		if strings.TrimSpace(value) == "" {
			f.ErrorsMap.Add(field, "This field cannot be blank")
		}
	}
}

func (f *Form) MaxLength(field string, d int) {
	value := f.Get(field)
	if value == "" {
		return
	}
	if utf8.RuneCountInString(value) > d {
		f.ErrorsMap.Add(field, fmt.Sprintf("This field is too long (maximum is %d characters)", d))
	}
}

func (f *Form) FileFormat(buff []byte) {
	filetype := http.DetectContentType(buff)
	if filetype != "image/jpeg" && filetype != "image/png" {
		f.ErrorsMap.Add("inputCover", "The provided file format is not allowed. "+
			"Please upload a JPEG or PNG image")
	}
}

func (f *Form) Valid() bool {
	return len(f.ErrorsMap) == 0
}
