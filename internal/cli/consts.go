package cli

const (
	CliName              = "Gocambridge"
	ExitOption           = "Exit"
	ReturnOption         = "Go back"
	ExitMessage          = "Bye!"
	InvalidOptionMessage = "Invalid option"
	InvalidURLMessage    = "invalid URL"
	UrlRegex             = `([^/]+)/class/([^/]+)/product/([^/]+)`
)

const (
	SelectUnitLabel      = "Select Unit"
	SelectSectionLabel   = "Select Section"
	SelectLessonLabel    = "Select Lesson"
	CorrectAnswersLabel  = "Correct Answer"
	CorrectAnswersSymbol = "#"
	NextActionLabel      = "Go back, Initial Menu, or Exit? (b/i/e)"
	SelectLimit          = 10
)

const (
	BackKey    = "b"
	InitialKey = "i"
	ExitKey    = "e"
)
