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
	NORMAL_QUESTION_REGEX   = `<p>Incorrect!<br />\s*Correct answer:<br />\s*(.+)</p>`
	MULTI_QUESTION_REGEX    = `Correct Answer:&lt;br /> (.*?)&lt;/p>`
	CHOICE_QUESTION_REGEX   = `(?s)Correct answers:<br />(.*?)</p>`
	POSSIBLE_QUESTION_REGEX = `(?s)Possible answers:<br />(.*?)</p>`
)
