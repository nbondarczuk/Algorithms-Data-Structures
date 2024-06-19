package power

type (
	Args struct {
		value float64
		exp   int
	}
	Results struct {
		positive bool
		expected float64
	}
	Case struct {
		args   Args
		result Results
	}
)

var tests = []Case{
	{
		Args{1.0, 1}, Results{true, 1.0},
	},
	{
		Args{1.0, 2}, Results{true, 1.0},
	},
	{
		Args{2.0, 1}, Results{true, 2.0},
	},
	{
		Args{2.0, 2}, Results{true, 4.0},
	},
	{
		Args{3.0, 3}, Results{true, 27.0},
	},
}
