package sort

type (
	Args struct {
		tab []int
	}
	Result struct {
		tab []int
	}
	Case struct {
		args     Args
		expected Result
	}
)

var tests = []Case{
	{
		Args{
			[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
		Result{
			[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
	},
	{
		Args{
			[]int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0},
		},
		Result{
			[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
	},
	{
		Args{
			[]int{8, 9, 6, 7, 4, 5, 2, 3, 0, 1},
		},
		Result{
			[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
	},
}
