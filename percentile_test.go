package percentile

import (
	"testing"
)

func TestPercentile(t *testing.T) {
	// 1 から 100 までの値を持つ配列を作成
	var numbers []float64
	for i := 1; i <= 100; i++ {
		numbers = append(numbers, float64(i))
	}

	// numbers について、25, 50, 75, 90, 95, 99 パーセンタイルを取得する
	for _, v := range []struct {
		percent int
		expect  float64
	}{
		{percent: 0, expect: 1.0},
		{percent: 25, expect: 25.5},
		{percent: 50, expect: 50.5},
		{percent: 75, expect: 75.5},
		{percent: 90, expect: 90.5},
		{percent: 95, expect: 95.5},
		{percent: 99, expect: 99.5},
		{percent: 100, expect: 100.0},
	} {
		actual := Calculate(numbers, v.percent)
		if actual != v.expect {
			t.Errorf("percentile(%d): Output = %f; want %f", v.percent, actual, v.expect)
		}
	}
}
