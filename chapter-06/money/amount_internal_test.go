package money

import (
	"testing"
)

func TestNewAmount(t *testing.T) {
	tt := map[string]struct {
		quantity Decimal
		currency Currency
		want     Amount
		err      error
	}{
		"1.50 €": {
			quantity: Decimal{subunits: 150, precision: 2},
			currency: Currency{code: "EUR", precision: 2},
			want: Amount{
				quantity: Decimal{subunits: 150, precision: 2},
				currency: Currency{code: "EUR", precision: 2},
			},
		},
		"1.500 €": {
			quantity: Decimal{subunits: 1500, precision: 3},
			currency: Currency{code: "EUR", precision: 2},
			err:      ErrTooPrecise,
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			got, err := NewAmount(tc.quantity, tc.currency)
			if err != tc.err {
				t.Errorf("got %v, want %v", err, tc.err)
			}
			if got != tc.want {
				t.Errorf("got %v, want %v", got, tc.want)
			}
		})
	}
}

func TestValidate(t *testing.T) {
	tt := map[string]struct {
		amount Amount
		err    error
	}{
		"too large": {
			amount: Amount{
				quantity: Decimal{subunits: maxDecimal + 1, precision: 2},
				currency: Currency{code: "EUR", precision: 2},
			},
			err: ErrTooLarge,
		},
		"too precise": {
			amount: Amount{
				quantity: Decimal{subunits: 150, precision: 4},
				currency: Currency{code: "XYZ", precision: 2},
			},
			err: ErrTooPrecise,
		},
		"valid": {
			amount: Amount{
				quantity: Decimal{subunits: 150, precision: 2},
				currency: Currency{code: "XYZ", precision: 2},
			},
			err: nil,
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			err := tc.amount.validate()
			if err != tc.err {
				t.Errorf("got %v, want %v", err, tc.err)
			}
		})
	}
}
