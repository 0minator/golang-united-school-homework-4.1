package reverse_string

import "testing"

func TestReverseString(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name       string
		args       args
		wantOutput string
	}{
		{name: "", args: struct{ input string }{input: "I am the Α and the Ω,\nthe First and the Last,\nthe Beginning and the End"}, wantOutput: "dnE eht dna gninnigeB eht\n,tsaL eht dna tsriF eht\n,Ω eht dna Α eht ma I"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotOutput := ReverseString(tt.args.input); gotOutput != tt.wantOutput {
				t.Errorf("ReverseString() = %v, want %v", gotOutput, tt.wantOutput)
			}
		})
	}
}
