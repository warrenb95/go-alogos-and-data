package questions

import (
	"testing"
)

func Test_isUnique(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "is unique",
			args: args{
				input: "abcdef",
			},
			want: true,
		},
		{
			name: "is not unique",
			args: args{
				input: "abcddef",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isUnique(tt.args.input); got != tt.want {
				t.Errorf("isUnique() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_checkPurmutations(t *testing.T) {
	type args struct {
		input1 string
		input2 string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "is purmutation",
			args: args{
				input1: "abc",
				input2: "cba",
			},
			want: true,
		},
		{
			name: "is not purmutation",
			args: args{
				input1: " abc",
				input2: "cba",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkPurmutations(tt.args.input1, tt.args.input2); got != tt.want {
				t.Errorf("checkPurmutations() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_urlify(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "book example",
			args: args{
				input: "Mr John Smith",
			},
			want: "Mr%20John%20Smith",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := urlify(tt.args.input); got != tt.want {
				t.Errorf("urlify() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_palindromePermutation(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "book example, should return true",
			args: args{
				input: "tact coa",
			},
			want: true,
		},
		{
			name: "my false example",
			args: args{
				input: "hello",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := palindromePermutation(tt.args.input); got != tt.want {
				t.Errorf("palindromePermutation() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_oneAway(t *testing.T) {
	type args struct {
		input  string
		output string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "pale -> ple == true",
			args: args{
				input:  "pale",
				output: "ple",
			},
			want: true,
		},
		{
			name: "pales -> pale == true",
			args: args{
				input:  "pales",
				output: "pale",
			},
			want: true,
		},
		{
			name: "pale -> bale == true",
			args: args{
				input:  "pale",
				output: "bale",
			},
			want: true,
		},
		{
			name: "pale -> bake == false",
			args: args{
				input:  "pale",
				output: "bake",
			},
			want: false,
		},
		{
			name: "pale -> palers",
			args: args{
				input:  "pale",
				output: "palers",
			},
			want: false,
		},
		{
			name: "palers -> pale ",
			args: args{
				input:  "palers",
				output: "pale",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := oneAway(tt.args.input, tt.args.output); got != tt.want {
				t.Errorf("oneAway() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_stringCompression(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "book exaple",
			args: args{
				input: "aabcccccaaa",
			},
			want: "a2b1c5a3",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := stringCompression(tt.args.input); got != tt.want {
				t.Errorf("stringCompression() = %v, want %v", got, tt.want)
			}
		})
	}
}
