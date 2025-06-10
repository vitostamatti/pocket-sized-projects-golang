package money

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

const (
	ErrInvalidDecimal = Error("unable to convert the decimal")
	ErrTooLarge       = Error("quantity over 10^12 is too large")
)

// const maxDecimal = 1e12
const maxDecimal = 1e16

// Decimal can represent a floating-point number with a fixed precision.
// Example:
// 1.52 = 152 * 10^-2 will be stored as {152, 2}
type Decimal struct {
	subunits  int64
	precision byte
}

// ExchangeRate represents a rate to convert from one currency to another.
type ExchangeRate Decimal

func ParseDecimal(value string) (Decimal, error) {
	intPart, fracPart, _ := strings.Cut(value, ".")

	subunits, err := strconv.ParseInt(intPart+fracPart, 10, 64)
	if err != nil {
		return Decimal{}, fmt.Errorf("%w: %s", ErrInvalidDecimal, err.Error())
	}

	if subunits > maxDecimal {
		return Decimal{}, ErrTooLarge
	}
	precision := byte(len(fracPart))
	decimal := Decimal{subunits, precision}
	decimal.simplify()
	return decimal, nil
}

// Using %10 returns the last digit inbase 10 of a number.
// If precision is positive, that digit belongs to the right side
// of the decimal point. Thus we can remove it by dividing by 10.
func (d *Decimal) simplify() {
	for d.subunits%10 == 0 && d.precision > 0 {
		d.precision--
		d.subunits /= 10
	}
}

// pow10 returns 10 raised to the power of the given byte.
// its optimized for small powers.
func pow10(power byte) int64 {
	switch power {
	case 0:
		return 1
	case 1:
		return 10
	case 2:
		return 100
	case 3:
		return 1000
	default:
		return int64(math.Pow(10, float64(power)))
	}
}

func (d *Decimal) String() string {
	if d.precision == 0 {
		return fmt.Sprintf("%d", d.subunits)
	}
	centsPerUnit := pow10(d.precision)
	frac := d.subunits % centsPerUnit
	integer := d.subunits / centsPerUnit

	decimalFormat := "%d.%0" + strconv.Itoa(int(d.precision)) + "d"

	return fmt.Sprintf(decimalFormat, integer, frac)

}
