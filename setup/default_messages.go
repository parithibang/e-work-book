package setup

import "github.com/astaxie/beego/validation"

func setDefaultMessage() {
	validation.SetDefaultMessage(map[string]string{
		"Required":     "Field is required",
		"Min":          "Minimum allowed value %d",
		"Max":          "Maximum allowed value %d",
		"Range":        "Must be in the range of %d before %d",
		"MinSize":      "Minimum length is %d",
		"MaxSize":      "Maximum allowable length %d",
		"Length":       "The length must be equal %d",
		"Alpha":        "Must be valid alpha characters",
		"Numeric":      "Must be valid numeric characters",
		"AlphaNumeric": "Must be valid alpha or numeric characters",
		"Match":        "Must match %s",
		"NoMatch":      "Must not match %s",
		"AlphaDash":    "Must be valid alpha or numeric or dash(-_) characters",
		"Email":        "Must be a valid email address",
		"IP":           "Must be a valid ip address",
		"Base64":       "Must be valid base64 characters",
		"Mobile":       "Must be valid mobile number",
		"Tel":          "Must be valid telephone number",
		"Phone":        "Must be valid telephone or mobile phone number",
		"ZipCode":      "Must be valid zipcode",
	})
}
