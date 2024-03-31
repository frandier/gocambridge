package cambridge

const (
	USER_AGENT         = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/118.0.0.0 Safari/537.36"
	APP_VERSION_HEADER = "cambridgeone-app-version"
	APP_VERSION        = "v2"
	BASE_CAMBRIDGE_URL = "https://www.cambridgeone.org/nlp/apigateway/%s/product/%s/"
	LESSON_URL         = "https://content.cambridgeone.org/cup1/products/%s/2/assets/ext-cup-xapiscoreable/%s/data.js"
	COOKIE_NAME        = "c1_sid"
	AJAX_DATA          = "ajaxData"
	AJAX_DATA_ERROR    = "ajaxData not found"
)

const (
	GENERIC_QUESTION_REGEX    = `<p>Incorrect!<br />\s*Correct answer:<br />\s*(.+)</p>`
	NORMAL_QUESTION_REGEX     = `<correctResponse>(.*?)<\/correctResponse>`
	MULTI_QUESTION_REGEX      = `Correct Answer:&lt;br /> (.*?)&lt;/p>`
	CHOICE_QUESTION_REGEX     = `(?s)Correct answers:<br />(.*?)</p>`
	CHOICE_QUESTION_V2_REGEX  = `(?s)Correct answer:<br />(.*?)</p>`
	VALUE_RESPONSE_REGEX      = `<value>(.*?)</value>`
	POSSIBLE_QUESTION_REGEX   = `(?s)Possible answers:<br />(.*?)</p>`
	QUESTION_IDENTIFIER_REGEX = `identifier="([^"]+)"`
)

const (
	INVALID_QUESTION_TYPE   = "Present:Present:Present"
	LEARNING_OBJECTINFO_XML = "LearningObjectInfo.xml"
	CORRECT_ANSWERS_FINDER  = "Correct answers:"
)
