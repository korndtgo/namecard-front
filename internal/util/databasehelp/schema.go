package databasehelp

import (
	"fmt"
	"regexp"
	"strings"
)

type SchemaType int

const (
	CampaignDB SchemaType = iota
)

func (dh *DatabaseHelp) GetSchemaName(key SchemaType) (string, bool) {
	title := "database="
	regex, _ := regexp.Compile(fmt.Sprintf(`%s[A-Za-z0-9_-]+`, title))

	switch key {
	case CampaignDB:
		str := regex.FindString(dh.config.DB)
		return fmt.Sprintf("[%s].dbo", str[strings.Index(str, title)+len(title):]), true
	default:
		return "", false
	}
}
