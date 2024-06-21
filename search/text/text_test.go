package text

type (
	Args struct {
		text string
		word string
	}
	Result struct {
		index int
	}
	Case struct {
		args   Args
		result Result
	}
)

var tests = []Case{
	{
		Args{
			"alfa beta gamma delta iota kappa",
			"alfa",
		},
		Result{
			0,
		},
	},
	{
		Args{
			"alfa beta gamma delta iota kappa",
			"beta",
		},
		Result{
			5,
		},
	},
	{
		Args{
			"alfa beta gamma delta iota kappa",
			"kappa",
		},
		Result{
			27,
		},
	},
	{
		Args{
			"alfa beta gamma delta iota kappa",
			"ta",
		},
		Result{
			7,
		},
	},
	{
		Args{
			"alfa beta gamma delta iota kappa",
			"tau",
		},
		Result{
			-1,
		},
	},
}
