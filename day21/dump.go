package day21

func topLevelMove(from, to int) string {
	switch from {
	case 0:
		switch to {
		case 0:
			return ""

		case 1:
			return ">"

		case 2:
			return ">>"

		case 3:
			return ">^"

		case 10:
			return ">>^"
		}

	case 1:
		switch to {
		case 0:
			return "<"

		case 1:
			return ""

		case 2:
			return ">"

		case 3:
			return "^"

		case 10:
			return ">^"
		}

	case 2:
		switch to {
		case 0:
			return "<<"

		case 1:
			return "<"

		case 2:
			return ""

		case 3:
			return "<^"

		case 10:
			return "^"
		}

	case 3:
		switch to {
		case 0:
			return "v<"

		case 1:
			return "v"

		case 2:
			return "v>"

		case 3:
			return ""

		case 10:
			return ">"
		}

	case 10:
		switch to {
		case 0:
			return "v<<"

		case 1:
			return "v<"

		case 2:
			return "v"

		case 3:
			return "<"

		case 10:
			return ""
		}
	}

	return ""
}
