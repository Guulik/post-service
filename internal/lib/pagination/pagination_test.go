package pagination

import (
	"errors"
	"github.com/stretchr/testify/require"
	"posts/internal/constants"
	"testing"
)

func TestGetLimitAndOffset(t *testing.T) {
	type args struct {
		page     *int
		pageSize *int
	}
	tests := []struct {
		name       string
		args       args
		wantOffset int
		wantLimit  int
		wantErr    error
	}{
		{
			name:       "Both nil",
			args:       args{nil, nil},
			wantOffset: 0,
			wantLimit:  -1,
		},
		{
			name:       "Nil page",
			args:       args{nil, ptr(10)},
			wantOffset: 0,
			wantLimit:  -1,
		},
		{
			name:       "Nil pageSize",
			args:       args{ptr(2), nil},
			wantOffset: 0,
			wantLimit:  -1,
		},
		{
			name:       "First page",
			args:       args{ptr(1), ptr(10)},
			wantOffset: 0,
			wantLimit:  10,
		},
		{
			name:       "Second page",
			args:       args{ptr(2), ptr(10)},
			wantOffset: 10,
			wantLimit:  10,
		},
		{
			name:       "Third page",
			args:       args{ptr(3), ptr(5)},
			wantOffset: 10,
			wantLimit:  5,
		},
		{
			name:       "Zero page",
			args:       args{ptr(0), ptr(10)},
			wantOffset: 0,
			wantLimit:  -1,
			wantErr:    errors.New(constants.WrongPageError),
		},
		{
			name:       "Negative page",
			args:       args{ptr(-1), ptr(10)},
			wantOffset: 0,
			wantLimit:  -1,
			wantErr:    errors.New(constants.WrongPageError),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			limit, offset, err := GetLimitAndOffset(tt.args.page, tt.args.pageSize)
			require.Equal(t, tt.wantLimit, limit)
			require.Equal(t, tt.wantOffset, offset)
			if err != nil {
				require.Equal(t, err.Error(), tt.wantErr.Error())
			}
		})
	}
}

func ptr(i int) *int {
	return &i
}
