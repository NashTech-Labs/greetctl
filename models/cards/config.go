package cards

// config contains the parameters of greeting.
type Config struct {
	Occasion string
	Language string
	UserName string
	CardID   string
}

const (
	langFR             = "fr"
	langEN             = "en"
	bDay               = "birthday"
	newYear            = "newyear"
	txGiving           = "thanksgiving"
	newYearGreetingENG = "Happy New Year !! "
	newYearGreetingFR  = "Bonne année !! "
	txDayGreetingENG   = "Happy Thanks Giving !! "
	txDayGreetingFR    = "Joyeux Action de Graces !! "
	bDayGreetingENG    = "Happy Birthday !! "
	bDayGreetingFR     = "Bon anniversaire !! "
)
