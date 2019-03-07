package gomock

import (
	"fmt"
	"testing"

	"github.com/budougumi0617/til/go/gomock/mock"
	"github.com/golang/mock/gomock"
)

func TestReturn(t *testing.T) {
	in := "any value"
	want := "hoge"

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mc := mock.NewMockClient(ctrl)

	mc.EXPECT().Do(in).Return("hoge", nil)

	if got, _ := mc.Do(in); got != want {
		t.Errorf("want %#v, but got %#v\n", want, got)
	}

}

func TestDoAndReturn(t *testing.T) {
	in := "value"
	want := "value modified in mock"

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mc := mock.NewMockClient(ctrl)

	mc.EXPECT().Do(in).DoAndReturn(
		func(in string) (string, error) {
			return fmt.Sprint(in, " modified in mock"), nil
		})
	if got, _ := mc.Do(in); got != want {
		t.Errorf("want %#v, but got %#v\n", want, got)
	}

}

func TestDoAndReturn2(t *testing.T) {

	in := "any value"

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mc := mock.NewMockClient(ctrl)

	mc.EXPECT().Do(in).DoAndReturn(
		func(in string) (string, error) {
			return "", fmt.Errorf("%s", in)
		})

	if _, err := mc.Do(in); err == nil {
		t.Error("cannot get error")
	}

}

func TestDoAndReturn3(t *testing.T) {

	errin := "raise error"
	specify := "specify"

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mc := mock.NewMockClient(ctrl)

	mc.EXPECT().Do(gomock.Any()).DoAndReturn(
		func(in string) (string, error) {
			switch in {
			case errin:
				return "", fmt.Errorf("%s", in)
			case specify:
				return "!!!!!", nil
			}
			return in, nil
		})
	// ここから！！！
	if _, err := mc.Do(errin); err == nil {
		t.Error("cannot get error")
	}

}

func TestBySubTest(t *testing.T) {
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
