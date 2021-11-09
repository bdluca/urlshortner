package sequence

import (
	b64 "encoding/base64"
	"strconv"
	"strings"
	"sync/atomic"
)

var id uint64

func Generate() string {
	newID := strconv.FormatUint(atomic.AddUint64(&id, 1), 10)
	encodedID := b64.StdEncoding.EncodeToString([]byte(newID))
	encodedID = strings.ReplaceAll(encodedID, "=", "")
	return encodedID
}
