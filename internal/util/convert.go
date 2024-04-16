package util

import (
	"regexp"
	"strings"
	"time"
)

const (
	LayoutDateDDMMYTYY string = "02/01/2006"
)

func ConvertStringToTime(dateStr, layout string) (time.Time, error) {
	date, err := time.Parse(layout, dateStr)
	if err != nil {
		return time.Time{}, err
	}
	return date, nil
}

func SlugGeneration(input string) string {
	output := strings.TrimSpace(strings.ToLower(input))
	patterns := map[*regexp.Regexp]string{
		regexp.MustCompile(`(à|á|ạ|ả|ã|â|ầ|ấ|ậ|ẩ|ẫ|ă|ằ|ắ|ặ|ẳ|ẵ)`): "a",
		regexp.MustCompile(`(è|é|ẹ|ẻ|ẽ|ê|ề|ế|ệ|ể|ễ)`):             "e",
		regexp.MustCompile(`(ì|í|ị|ỉ|ĩ)`):                         "i",
		regexp.MustCompile(`(ò|ó|ọ|ỏ|õ|ô|ồ|ố|ộ|ổ|ỗ|ơ|ờ|ớ|ợ|ở|ỡ)`): "o",
		regexp.MustCompile(`(ù|ú|ụ|ủ|ũ|ư|ừ|ứ|ự|ử|ữ)`):             "u",
		regexp.MustCompile(`(ỳ|ý|ỵ|ỷ|ỹ)`):                         "y",
		regexp.MustCompile(`(đ)`):                                 "d",
		regexp.MustCompile(`([^0-9a-z-\s])`):                      "",
		regexp.MustCompile(`(\s+)`):                               "-",
		regexp.MustCompile(`^-+`):                                 "",
		regexp.MustCompile(`-+$`):                                 "",
	}
	for pattern, replace := range patterns {
		output = pattern.ReplaceAllString(output, replace)
	}
	return output
}
