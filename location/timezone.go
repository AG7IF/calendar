package location

import (
	"fmt"
	"strings"
	"time"
)

const SecondsInDay = 86400

type TZ int

const (
	AcpZ TZ = iota
	AcpA
	AcpB
	AcpC
	AcpD
	AcpE
	AcpF
	AcpG
	AcpH
	AcpI
	AcpJ
	AcpK
	AcpM
	AcpN
	AcpO
	AcpP
	AcpQ
	AcpR
	AcpS
	AcpT
	AcpU
	AcpV
	AcpW
	AcpX
	AcpY
)

func ParseTZ(s string) (TZ, error) {
	switch strings.TrimSpace(strings.ToUpper(s)) {
	case "Z":
		return AcpZ, nil
	case "A":
		return AcpA, nil
	case "B":
		return AcpB, nil
	case "C":
		return AcpC, nil
	case "D":
		return AcpD, nil
	case "E":
		return AcpE, nil
	case "F":
		return AcpF, nil
	case "G":
		return AcpG, nil
	case "H":
		return AcpH, nil
	case "I":
		return AcpI, nil
	case "J":
		return AcpJ, nil
	case "K":
		return AcpK, nil
	case "M":
		return AcpM, nil
	case "N":
		return AcpN, nil
	case "O":
		return AcpO, nil
	case "P":
		return AcpP, nil
	case "Q":
		return AcpQ, nil
	case "R":
		return AcpR, nil
	case "S":
		return AcpS, nil
	case "T":
		return AcpT, nil
	case "U":
		return AcpU, nil
	case "V":
		return AcpV, nil
	case "W":
		return AcpW, nil
	case "X":
		return AcpX, nil
	case "Y":
		return AcpY, nil
	default:
		return AcpZ, fmt.Errorf("unrecognized ACP timezone code: %s", s)
	}
}

func (tz TZ) HourOffset() int {
	switch tz {
	case AcpZ:
		return 0
	case AcpA:
		return 1
	case AcpB:
		return 2
	case AcpC:
		return 3
	case AcpD:
		return 4
	case AcpE:
		return 5
	case AcpF:
		return 6
	case AcpG:
		return 7
	case AcpH:
		return 8
	case AcpI:
		return 9
	case AcpJ:
		return 10
	case AcpK:
		return 11
	case AcpM:
		return 12
	case AcpN:
		return -1
	case AcpO:
		return -2
	case AcpP:
		return -3
	case AcpQ:
		return -4
	case AcpR:
		return -5
	case AcpS:
		return -6
	case AcpT:
		return -7
	case AcpU:
		return -8
	case AcpV:
		return -9
	case AcpW:
		return -10
	case AcpX:
		return -11
	case AcpY:
		return -12
	default:
		panic(fmt.Sprintf("invalid ACP timezone value: %d", tz))
	}
}

func (tz TZ) Location() *time.Location {
	offset := tz.HourOffset()
	return time.FixedZone(tz.String(), offset*SecondsInDay)
}

func (tz TZ) String() string {
	switch tz {
	case AcpZ:
		return "Z"
	case AcpA:
		return "A"
	case AcpB:
		return "B"
	case AcpC:
		return "C"
	case AcpD:
		return "D"
	case AcpE:
		return "E"
	case AcpF:
		return "F"
	case AcpG:
		return "G"
	case AcpH:
		return "H"
	case AcpI:
		return "I"
	case AcpJ:
		return "J"
	case AcpK:
		return "K"
	case AcpM:
		return "M"
	case AcpN:
		return "N"
	case AcpO:
		return "O"
	case AcpP:
		return "P"
	case AcpQ:
		return "Q"
	case AcpR:
		return "R"
	case AcpS:
		return "S"
	case AcpT:
		return "T"
	case AcpU:
		return "U"
	case AcpV:
		return "V"
	case AcpW:
		return "W"
	case AcpX:
		return "X"
	case AcpY:
		return "Y"
	default:
		panic(fmt.Sprintf("invalid ACP timezone value: %d", tz))
	}
}
