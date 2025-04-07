package tool

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"
)

type jsonTool struct {
}

var JsonTool = (*jsonTool)(nil)

func (r *jsonTool) JsonContentCompare(a []byte, b []byte) bool {
	if len(a) == 0 && len(b) == 0 {
		return true
	}
	if len(a) == 0 || len(b) == 0 {
		return false
	}
	dst1 := &bytes.Buffer{}
	if err := json.Compact(dst1, a); err != nil {
		fmt.Printf("JsonContentCompare error: %s", err.Error())
		return false
	}
	dst2 := &bytes.Buffer{}
	if err := json.Compact(dst2, b); err != nil {
		fmt.Printf("JsonContentCompare error: %s", err.Error())
		return false
	}
	return strings.Compare(dst1.String(), dst2.String()) == 0
}

func (r *jsonTool) JsonEncoding(data any, escapeHTML bool, pretty bool) string {
	str := &strings.Builder{}
	encoder := json.NewEncoder(str)
	// avoid special character parse
	encoder.SetEscapeHTML(escapeHTML)
	if pretty {
		encoder.SetIndent("", "    ")
	}
	encoder.Encode(data)
	return str.String()
}
