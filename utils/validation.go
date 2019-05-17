package utils

import (
	"bytes"
	"github.com/astaxie/beego/validation"
)

func ValidParams(params interface{}) (string, error) {
	valid := validation.Validation{}
	b, err := valid.Valid(params)
	if err != nil {
		return "", err
	}
	var buffer bytes.Buffer
	if !b {
		// validation does not pass
		// blabla...
		for _, err := range valid.Errors {
			//fmt.Println(err.Key, err.Message)
			buffer.WriteString(err.Key)
			buffer.WriteString(",")
			buffer.WriteString(err.Message)
			buffer.WriteString("|")
		}
	}
	return buffer.String(), nil
}
