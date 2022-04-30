package square

import "testing"

func TestCalcSquare(t *testing.T) {
	type args struct {
		sideLen  float64
		sidesNum intCustomType
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{name: "square", args: struct {
			sideLen  float64
			sidesNum intCustomType
		}{sideLen: 2, sidesNum: 4}, want: 4},
		{name: "wrong", args: struct {
			sideLen  float64
			sidesNum intCustomType
		}{sideLen: 0, sidesNum: 0}, want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CalcSquare(tt.args.sideLen, tt.args.sidesNum); got != tt.want {
				t.Errorf("CalcSquare() = %v, want %v", got, tt.want)
			}
		})
	}
}
