package errs

import (
	"log"

	"github.com/davecgh/go-spew/spew"
)

func Spew(v ...any) {
	dumped := spew.Sdump(v...)
	log.Println(dumped)
}
