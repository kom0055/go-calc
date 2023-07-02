package main

import (
	"reflect"
	"testing"
)

func TestEvaluate(t *testing.T) {
	type args struct {
		d     string
		debug bool
	}
	tests := []struct {
		name    string
		args    args
		want    interface{}
		wantErr bool
	}{
		{
			name: "",
			args: args{
				d:     "(1)",
				debug: false,
			},
			want:    1,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Evaluate(tt.args.d, tt.args.debug)
			if (err != nil) != tt.wantErr {
				t.Errorf("Evaluate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Evaluate() got = %v, want %v", got, tt.want)
			}
		})
	}
}
