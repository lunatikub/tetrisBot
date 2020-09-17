package player

var tetriminoJ = tetrimino{
	"J",
	[]piece{

		// +---+
		// |..X|
		// |XXX|
		// +---+
		{
			[][]int{
				{0, 0, 1},
				{1, 1, 1},
			},
			[]int{0, 0, 0},
			[]int{1, 1, 2},
			[]int{1, 3},
			2,
			3,
		},

		// +--+
		// |X.|
		// |X.|
		// |XX|
		// +--+
		{
			[][]int{
				{1, 0},
				{1, 0},
				{1, 1},
			},
			[]int{0, 0},
			[]int{3, 1},
			[]int{1, 1, 2},
			3,
			2,
		},

		// +---+
		// |XXX|
		// |X..|
		// +---+
		{
			[][]int{
				{1, 1, 1},
				{1, 0, 0},
			},
			[]int{0, 1, 1},
			[]int{2, 2, 2},
			[]int{3, 1},
			2,
			3,
		},

		// +--+
		// |XX|
		// |.X|
		// |.X|
		// +--+
		{
			[][]int{
				{1, 1},
				{0, 1},
				{0, 1},
			},
			[]int{2, 0},
			[]int{3, 3},
			[]int{2, 1, 1},
			3,
			2,
		},
	},
}
