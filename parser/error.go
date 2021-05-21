package parser

import "fmt"

type FileParseError struct {
  Filename   string
  Additional error
}

func (e FileParseError) Error() string {
  return fmt.Sprintf("\"%s\": (%s)", e.Filename, e.Additional.Error());
}
