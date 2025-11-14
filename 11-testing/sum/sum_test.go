package sum

import "testing"

// create file with _test.go to register them as test file

//function name should start with Test to register them as test
// and it must have a pointer to testing.T as its first argument

// go test -v ./... // run all the tests for the project

func TestSumInt(t *testing.T) {
	input := []int{1, 2, 3, 4, 5}
	want := 15
	got := SumInt(input)
	if got != want {
		// test would continue on if test case fail in case of Errorf
		t.Errorf("sum of 1 to 5 should be %d, got %d", want, got)

		// test would stop on if test case fail in case of Fatalf
		//t.Fatalf("sum of 1 to 5 should be %d, go %d", want, got)
	}

	input = []int{1, -1}
	want = 0
	got = SumInt(input)
	if got != want {
		t.Errorf("sum of 1 to 5 should be %d, got %d", want, got)
	}

	want = 0
	got = SumInt(nil)
	if got != want {
		t.Errorf("sum of nil should be %d, got %d", want, got)
	}

}

func TestSumIntV2(t *testing.T) {
	// Table Driven Testing
	tt := []struct {
		name  string
		input []int
		want  int
	}{
		{
			name:  "sum of 1 to 5",
			input: []int{1, 2, 3, 4, 5},
			want:  15,
		},
		{
			name:  "nil slice",
			input: nil,
			want:  0,
		},
		{
			name:  "1 -1",
			input: []int{1, -1},
			want:  0,
		},
	}

	for _, tc := range tt {
		// t.Run creates a sub test
		// each test case would be running in a subtest
		// failing one subtest would not affect other subtests
		t.Run(tc.name, func(t *testing.T) {
			got := SumInt(tc.input)
			if got != tc.want {
				t.Fatalf("sum of %v should be %d, got %d", tc.input, tc.want, got)
			}
		})

	}
}
