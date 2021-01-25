package bird_observer

import (
	"testing"

	"github.com/abhi1kush/mockgen_golang_testing/bird/mocks"
	"github.com/golang/mock/gomock"
)

func TestObserver(t *testing.T) {
	type args struct {
		weight      int
		measure_err int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "tc1",
			args: args{
				weight:      10,
				measure_err: 2,
			},
			want: 100/10 - 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewMockController(t)
			defer ctrl.Finish()
			mockBird := mocks.NewMockBird(ctrl)
			mockBird.EXPECT().FlyingSpeed().Return(100)
			if got := Observer(tt.args.weight, tt.args.measure_err); got != tt.want {
				t.Errorf("Observer() = %v, want %v", got, tt.want)
			}
		})
	}
}
