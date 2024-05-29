package crypto

import (
	"testing"
)

func TestEncrypt(t *testing.T) {
	type args struct {
		input string
	}

	cases := []struct {
		name    string
		input   string
		want    string
		wantErr bool
	}{
		{name: "case1", input: "[!@#$%^&*()]"},
		{name: "case1", input: "`1qazZSE$<>?"},
		{name: "case1", input: "@WSXXDR%"},
		{name: "case1", input: "#EDCCFT^"},
		{name: "case1", input: "1234567890"},
		{name: "case1", input: "abcdefghigklmnopqrstuvwxyz"},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			decPass, err := Encrypt(tt.input)
			if err != nil {
				t.Errorf("err:%+v", err)
			}

			tt.want, err = Decrypt(decPass)
			if err != nil {
				t.Errorf("err:%v", err)
			}

			if tt.input != tt.want {
				t.Errorf("input:%s, want:%s, decPass:%s, err:%v", tt.input, tt.want, decPass, err)
			} else {
				t.Logf("decPass:%s", decPass)
				t.Logf("input:%s", tt.input)
				t.Logf("want :%s", tt.want)
			}
		})
	}
}
