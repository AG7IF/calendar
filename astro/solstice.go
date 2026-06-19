package astro

type Solstice int

const (
	NoSolstice Solstice = iota
	VernalEquinox
	SummerSolstice
	AutumnalEquinox
	WinterSolstice
)

func (s Solstice) LaTeX() string {
	var solstice string
	switch s {
	case VernalEquinox:
		solstice = `\Aries`
	case SummerSolstice:
		solstice = `\Cancer`
	case AutumnalEquinox:
		solstice = `\Libra`
	case WinterSolstice:
		solstice = `\Capricorn`
	default:
		solstice = ""
	}

	return solstice
}
