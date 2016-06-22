package num2word

import "testing"

var samples = []struct {
	amount   float64
	upper    bool
	expected string
}{
	{1, true, "Один рубль 00 копеек"},
	{100.21, false, "сто рублей 21 копейка"},
}

func Test_RuMoney(t *testing.T) {
	for _, tt := range samples {
		res := RuMoney(tt.amount, tt.upper)
		if res != tt.expected {
			t.Errorf("RuMoney(%.2f): expected '%s', got '%s'", tt.amount, tt.expected, res)
		}
	}
}
