package decimal

import (
	"fmt"
	"math"
	"strings"

	"github.com/dustin/go-humanize"
)

func RoundTo(n float64, decimals uint32) float64 {
	return math.Round(n*math.Pow(10, float64(decimals))) / math.Pow(10, float64(decimals))
}

func NumbericCommaWithDigits(value float64, digit int) string {
	c := humanize.CommafWithDigits(math.Round(value*math.Pow10(digit))/math.Pow10(digit), digit)
	var r string
	spliteds := strings.Split(c, ".")
	if len(spliteds) > 1 {
		if len(spliteds[1]) < digit {
			d := fmt.Sprintf("%s", spliteds[1]+strings.Repeat("0", digit))
			r = spliteds[0] + "." + d[:digit]
		} else {
			r = c
		}
	} else if digit > 0 {
		r = c + "." + strings.Repeat("0", digit)
	} else {
		r = c
	}
	return r
}
