package cli

import (
	"fmt"
	"regexp"

	"github.com/frandier/gocambridge/internal/cambridge"
)

func Run(url, cookie string) error {
	re := regexp.MustCompile(UrlRegex)
	matches := re.FindStringSubmatch(url)

	if len(matches) != 4 {
		return fmt.Errorf(InvalidURLMessage)
	}

	org := matches[1]
	idClass := matches[2]
	product := matches[3]
	cambridgeClient := cambridge.NewCambridge(org, product, cookie)
	resp, err := cambridgeClient.GetUnits(idClass)
	if err != nil {
		return err
	}

	return selectUnit(resp, cambridgeClient, product)
}
