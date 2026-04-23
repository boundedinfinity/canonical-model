package authenticator_application

import "github.com/boundedinfinity/rfc3339date"

type BackupCode struct {
	Code   string                      `json:"code"`
	UsedOn rfc3339date.Rfc3339DateTime `json:"used-on"`
}
