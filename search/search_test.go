package search

type (
	Args struct {
		tab     []int
		element int
	}
	Result struct {
		expected bool
		index    int
	}
	Case struct {
		args   Args
		result Result
	}
)

var tests = []Case{
	{
		Args{
			[]int{1, 2, 3},
			1,
		},
		Result{
			true,
			0,
		},
	},
	{
		Args{
			[]int{1, 2, 3},
			3,
		},
		Result{
			true,
			2,
		},
	},
	{
		Args{
			[]int{1, 2, 3},
			4,
		},
		Result{
			false,
			3,
		},
	},
}
