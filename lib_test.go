package stoverflow_data

import "testing"
import "fmt"

func TestGetData(t *testing.T) {
	type args struct {
		tags []string
	}
	tests := []struct {
		name string
		args args
	}{
		{"Python Web Dev", args{[]string{"python", "django"}}},
		{"No tags", args{[]string{}}},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ret := GetData(tt.args.tags)
			for _, item := range ret {
				fmt.Printf("%+v\n", item)
			}
		})
	}
}
