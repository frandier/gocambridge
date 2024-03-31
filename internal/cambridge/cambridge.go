package cambridge

import (
	"fmt"
	"net/http"
	"regexp"
	"sort"
	"strings"

	"github.com/dop251/goja"
	"github.com/go-resty/resty/v2"
	"golang.org/x/net/html"
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
	genericQuestion := regexp.MustCompile(GENERIC_QUESTION_REGEX)
	multiQuestion := regexp.MustCompile(MULTI_QUESTION_REGEX)
	choiceQuestion := regexp.MustCompile(CHOICE_QUESTION_REGEX)
	possibleAnswersQuestion := regexp.MustCompile(POSSIBLE_QUESTION_REGEX)
	valueResponse := regexp.MustCompile(VALUE_RESPONSE_REGEX)

	objectKeys := ajaxDataObject.Keys()

	sort.Slice(objectKeys, func(i, j int) bool {
		return compare(objectKeys[i], objectKeys[j])
	})

	for _, key := range objectKeys {
		dataValue := ajaxDataObject.Get(key)
		dataValueStr := dataValue.String()

		questionType := getQuestionType(dataValueStr)

		if questionType == "" || questionType == INVALID_QUESTION_TYPE {
			continue
		}

		var matches [][]string

		// TODO: Refactor this switch statement
		switch questionType {
		case "Order:Match:Text gap":
			matches = append(matches, genericQuestion.FindAllStringSubmatch(dataValueStr, -1)...)
			if len(matches) == 0 {
				matches = append(matches, multiQuestion.FindAllStringSubmatch(dataValueStr, -1)...)
			}
		case "Identify:Select:Dropdown":
			matches = append(matches, genericQuestion.FindAllStringSubmatch(dataValueStr, -1)...)
		case "Input:Completion:Text gap":
			matches = append(matches, normalQuestion.FindAllStringSubmatch(dataValueStr, -1)...)
		case "Identify:Select:Radiobutton":
			matches = append(matches, genericQuestion.FindAllStringSubmatch(dataValueStr, -1)...)
		case "Identify:Select:Checkbox":
			matches = choiceQuestion.FindAllStringSubmatch(dataValueStr, -1)
			if len(matches) == 0 {
				doc, _ := html.Parse(strings.NewReader(dataValueStr))
				var correctAnswersText string
				traverse(doc, &correctAnswersText)
				return []string{correctAnswersText}, nil
			}
		default:
			matches = append(matches, possibleAnswersQuestion.FindAllStringSubmatch(dataValueStr, -1)...)
		}

		for _, match := range matches {
			if len(match) > 1 {
				valueMatches := valueResponse.FindAllStringSubmatch(match[1], -1)
				if len(valueMatches) > 0 {
					cleanedString := cleanSimpleQuestionValue(valueMatches)
					results = append(results, cleanedString)
					continue
				}

				result := html.UnescapeString(match[1])
				result = strings.ReplaceAll(result, "&apos;", "'")
				results = append(results, result)
			}
		}
	}

	return results, nil
}
