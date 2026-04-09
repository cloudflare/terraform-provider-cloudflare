package schemata

import (
	"fmt"
	"strings"
)

type Description struct {
	Scopes              []string
	MarkdownDescription string
}

func (d Description) String() string {
	sb := strings.Builder{}
	if len(d.Scopes) > 0 {
		sb.WriteString("Accepted Permissions\n\n")
	}
	for _, scope := range d.Scopes {
		sb.WriteString("- `")
		sb.WriteString(strings.ReplaceAll(scope, "`", "\\`"))
		sb.WriteString("`\n")
	}
	if d.MarkdownDescription != "" {
		if sb.Len() > 0 {
			sb.WriteString("\n")
		}
		sb.WriteString(d.MarkdownDescription)
	}
	return sb.String()
}

var _ fmt.Stringer = (*Description)(nil)
