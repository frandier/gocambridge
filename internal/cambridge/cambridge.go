package cambridge

import (
	"fmt"
	"html"
	"net/http"
	"regexp"
	"strings"

	"github.com/dop251/goja"
	"github.com/go-resty/resty/v2"
)

type Cambridge struct {
	http *resty.Client
}

func NewCambridge(org, product, cookie string) *Cambridge {
	return &Cambridge{
		http: resty.New().
			SetBaseURL(fmt.Sprintf(BASE_CAMBRIDGE_URL, org, product)).
			SetHeaders(map[string]string{APP_VERSION_HEADER: APP_VERSION}).
			SetCookies([]*http.Cookie{{Name: COOKIE_NAME, Value: cookie}}),
	}
}

func (c *Cambridge) GetUnits(class string) (*UnitsResult, error) {
	result := &UnitsResult{}
	resp, err := c.http.R().SetResult(result).Get(class)
	if err != nil {
		return nil, err
	}
	if resp.IsError() {
		return nil, resp.Error().(error)
	}
	return result, nil
}

func (c *Cambridge) GetLessonResponse(productCode, lessonId string) ([]string, error) {
	resp, err := c.http.R().Get(fmt.Sprintf(LESSON_URL, productCode, lessonId))
	if err != nil {
		return nil, err
	}

	runtime := goja.New()
	_, err = runtime.RunString(string(resp.Body()))
	if err != nil {
		return nil, err
	}

	ajaxDataValue := runtime.Get(AJAX_DATA)
	if ajaxDataValue == nil {
		return nil, fmt.Errorf(AJAX_DATA_ERROR)
	}

	ajaxDataObject := ajaxDataValue.ToObject(runtime)
	var results []string

	normalQuestion := regexp.MustCompile(NORMAL_QUESTION_REGEX)
	multiQuestion := regexp.MustCompile(MULTI_QUESTION_REGEX)
	choiceQuestion := regexp.MustCompile(CHOICE_QUESTION_REGEX)
	possibleAnswersQuestion := regexp.MustCompile(POSSIBLE_QUESTION_REGEX)

	for _, key := range ajaxDataObject.Keys() {
		dataValue := ajaxDataObject.Get(key)
		dataValueStr := dataValue.String()

		var matches [][]string
		matches = choiceQuestion.FindAllStringSubmatch(dataValueStr, -1)
		matches = append(matches, normalQuestion.FindAllStringSubmatch(dataValueStr, -1)...)
		matches = append(matches, multiQuestion.FindAllStringSubmatch(dataValueStr, -1)...)
		matches = append(matches, possibleAnswersQuestion.FindAllStringSubmatch(dataValueStr, -1)...)

		for _, match := range matches {
			if len(match) > 1 {
				result := html.UnescapeString(match[1])
				result = strings.ReplaceAll(result, "&apos;", "'")
				results = append(results, result)
			}
		}
	}

	return results, nil
}
