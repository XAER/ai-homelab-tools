package clients

import (
	"encoding/base64"
)

func fiberutilsBasic(user, pass string) string {
	return base64.StdEncoding.EncodeToString([]byte(user + ":" + pass))
}

