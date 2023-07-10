package handlers

import (
	"github.com/golang-jwt/jwt/v5"
	"reflect"
	"testing"
)

func TestGetSigningMethod(t *testing.T) {
	type args struct {
		label string
	}
	tests := []struct {
		name    string
		args    args
		want    jwt.SigningMethod
		wantErr bool
	}{
		{name: "RS256", args: args{label: "RS256"}, want: jwt.SigningMethodRS256, wantErr: false},
		{name: "hs384", args: args{label: "hs384"}, want: jwt.SigningMethodHS384, wantErr: false},
		{name: "bogus123", args: args{label: "bogus123"}, want: nil, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetSigningMethod(tt.args.label)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetSigningMethod() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetSigningMethod() got = %v, want %v", got, tt.want)
			}
		})
	}
}
