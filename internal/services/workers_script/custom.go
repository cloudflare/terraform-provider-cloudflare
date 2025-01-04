package workers_script

import (
	"fmt"
	"io"
	"mime/multipart"
	"net/textproto"
	"strings"
)

func writeFileBytes(partName string, filename string, contentType string, content io.Reader, writer *multipart.Writer) error {
	h := make(textproto.MIMEHeader)
	header := "form-data"

	if escapeQuotes(partName) != "" {
		header += fmt.Sprintf(`; name="%s"`, escapeQuotes(partName))
	}

	if escapeQuotes(filename) != "" {
		header += fmt.Sprintf(`; filename="%s"`, escapeQuotes(filename))
	}

	h.Set("Content-Disposition", header)
	h.Set("Content-Type", contentType)
	filewriter, err := writer.CreatePart(h)
	if err != nil {
		return err
	}
	_, err = io.Copy(filewriter, content)
	return err
}

var quoteEscaper = strings.NewReplacer("\\", "\\\\", `"`, "\\\"")

func escapeQuotes(s string) string {
	return quoteEscaper.Replace(s)
}
