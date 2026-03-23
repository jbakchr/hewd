package formatter

const (
	Reset = "\033[0m"
	Bold  = "\033[1m"

	// Colors
	Red             = "\033[31m"
	RedBold         = "\033[31;1m"
	Green           = "\033[32m"
	Yellow          = "\033[33m"
	Blue            = "\033[34m"
	cyan            = "\033[36m"
	cyanBold        = "\033[36;1m"
	CyanUnderline   = "\033[4;36m"
	CyanItalic      = "\033[3;36m"
	whiteBold       = "\033[37;1m"
	whiteItalic     = "\033[3;37m"
	whiteBoldItalic = "\033[1;3;37m"
)

// DisableColor allows global disabling of ANSI formatting.
// You could later toggle this for CI or when stdout is not a TTY.
var DisableColor = false

// apply wraps text in a style escape sequence unless color is disabled.
func apply(style, s string) string {
	if DisableColor {
		return s
	}
	return style + s + Reset
}

// Cyan applies cyan coloring to the given string.
func Cyan(s string) string {
	return apply(cyan, s)
}

func CyanBold(s string) string {
	return apply(cyanBold, s)
}

// WhiteBold prints white bold text.
func WhiteBold(s string) string {
	return apply(whiteBold, s)
}

// WhiteItalic prints white italic text.
func WhiteItalic(s string) string {
	return apply(whiteItalic, s)
}

// WhiteBoldItalic prints white bold italic text.
func WhiteBoldItalic(s string) string {
	return apply(whiteBoldItalic, s)
}
