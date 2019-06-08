package fAnalysis

import (
	"reflect"
	"sort"
	"testing"
)

func Test_fAnalysis(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "asd-3",
			args: args{
				s: "asd asd asd",
			},
			want: []string{
				"asd",
			},
		}, {
			name: "asd-3 with symbols",
			args: args{
				s: "asd, asd: asd ,.!",
			},
			want: []string{
				"asd",
			},
		}, {
			name: "asd-4 zxc-1 qwe-2",
			args: args{
				s: "asd asd asd zxc asd qwe qwe",
			},
			want: []string{
				"asd", "qwe", "zxc",
			},
		}, {
			name: "asd-3 a-5 b-4 c-3 d-1 v-2",
			args: args{
				s: "asd a a c a b v b d asd c a c v b asd a b",
			},
			want: []string{
				"a", "asd", "b", "c", "d", "v",
			},
		}, {
			name: "a-7 b-4 c-3 d-1 e-2 r-2 t-1 u-2 v-2 w-3 z-1 x-1 p-1",
			args: args{
				s: "a a c a b a v b a d c r r a c v w w w t u u e e b a b",
			},
			want: []string{
				"a", "b", "c", "d", "e", "r", "t", "u", "v", "w",
			},
		}, {
			name: "single word",
			args: args{
				s: "asd",
			},
			want: []string{
				"asd",
			},
		}, {
			name: "empty string",
			args: args{
				s: "",
			},
			want: []string{},
		}, {
			name: "string without letters",
			args: args{
				s: ",.()?!:''",
			},
			want: []string{},
		}, {
			name: "10 different words",
			args: args{
				s: "b h c j f e g a i d",
			},
			want: []string{
				"a", "b", "c", "d", "e", "f", "g", "h", "i", "j",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := fAnalysis(tt.args.s)
			sort.Strings(got)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("fAnalysis() = %v, want %v", got, tt.want)
			}
		})
	}
}
