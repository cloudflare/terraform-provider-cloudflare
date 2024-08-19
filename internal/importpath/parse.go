package importpath

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/diag"
)

func ParseImportID(str string, format string, args ...any) (diag diag.Diagnostics) {
	path_spec, path := strings.Split(format, "/"), strings.Split(str, "/")

	if len(args) != len(path_spec) {
		panic(fmt.Sprintf("argument count does not match format: %q", format))
	}
	if len(path) != len(path_spec) {
		diag.AddError("invalid ID", fmt.Sprintf("expected urlencoded segments %q, got %q", format, str))
		return
	}

	var err error
	for i, arg := range args {
		spec, segment := path_spec[i], path[i]
		switch ptr := arg.(type) {
		case *bool:
			*ptr, err = strconv.ParseBool(segment)
			if err != nil {
				diag.AddError("invalid bool segment", fmt.Sprintf("for %q : %s -> %q", spec, err.Error(), segment))
			}
		case *int64:
			*ptr, err = strconv.ParseInt(segment, 10, 64)
			if err != nil {
				diag.AddError("invalid int segment", fmt.Sprintf("for %q : %s -> %q", spec, err.Error(), segment))
			}
		case *float64:
			*ptr, err = strconv.ParseFloat(segment, 64)
			if err != nil {
				diag.AddError("invalid float segment", fmt.Sprintf("for %q : %s -> %q", spec, err.Error(), segment))
			}
		case *string:
			*ptr, err = url.PathUnescape(segment)
			if err != nil {
				diag.AddError("invalid urlencoded segment", fmt.Sprintf("for %q : %s -> %q", spec, err.Error(), segment))
			}
		default:
			panic(fmt.Sprintf("invalid argument type for segment: %q", segment))
		}
	}
	return
}
