package gomock

import (
	"fmt"
	"testing"

	"github.com/budougumi0617/til/go/gomock/mock"
	"github.com/golang/mock/gomock"
)

func TestDoAndReturn(t *testing.T) {
	tests := []struct {
		name      string
		setClient func(*mock.MockClient, string)
		in        string
		want      string
		wantErr   bool
	}{
		{
			name: "Return",
			setClient: func(mc *mock.MockClient, in string) {
				mc.EXPECT().Do(in).Return("hoge", nil)
			},
			in:      "any value",
			want:    "hoge",
			wantErr: false,
		},
		{
			name: "DoAndReturn",
			setClient: func(mc *mock.MockClient, in string) {
				mc.EXPECT().Do(in).DoAndReturn(
					func(in string) (string, error) {
						return in, nil
					})
			},
			in:      "input value",
			want:    "input value",
			wantErr: false,
		},
		{
			name: "DoAndReturn2",
			setClient: func(mc *mock.MockClient, in string) {
				mc.EXPECT().Do(in).DoAndReturn(
					func(in string) (string, error) {
						return "", fmt.Errorf("%s", in)
					})
			},
			in:      "any  value",
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mc := mock.NewMockClient(ctrl)
		tt.setClient(mc, tt.in)
		got, err := mc.Do(tt.in)
		if !tt.wantErr && err != nil {
			t.Fatal(err)
		}

		if got != tt.want {
			t.Errorf("want %#v, but got %#v\n", tt.want, got)
		}
	}

}
