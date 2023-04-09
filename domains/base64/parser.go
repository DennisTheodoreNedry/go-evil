package base64

import (
	"fmt"

	"github.com/TeamPhoneix/go-evil/domains/base64/decode"
	"github.com/TeamPhoneix/go-evil/domains/base64/encode"
	"github.com/s9rA16Bf4/notify_handler/go/notify"
)

func Parser(function string, value string, s_json string) ([]string, string) {
	call := []string{}

	switch function {
	case "encode":
		call, s_json = encode.Encode(value, s_json)

	case "decode":
		call, s_json = decode.Decode(value, s_json)

	default:
		notify.Error(fmt.Sprintf("Unknown function '%s'", function), "base64.Parser()")

	}

	return call, s_json
}
