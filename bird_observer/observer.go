package bird_observer

import "github.com/abhi1kush/mockgen_golang_testing/bird"

func Observer(weight int, measure_err int) int {
	s := bird.Swan{
		Color:  "white",
		Weight: weight,
	}

	return s.FlyingSpeed() - measure_err
}
