package player

var tetrominoL = tetromino{
	"L",
	[]piece{

		// +---+
		// |X..|
		// |XXX|
		// +---+
		{
			[][]int{
				{1, 0, 0},
				{1, 1, 1},
			},
			[]int{0, 0, 0},
			2,
			3,
		},

		// +--+
		// |XX|
		// |X.|
		// |X.|
		// +--+
		{
			[][]int{
				{1, 1},
				{1, 0},
				{1, 0},
			},
			[]int{0, 2},
			3,
			2,
		},

		// +---+
		// |XXX|
		// |..X|
		// +---+
		{
			[][]int{
				{1, 1, 1},
				{0, 0, 1},
			},
			[]int{1, 1, 0},
			2,
			3,
		},

		// +--+
		// |.X|
		// |.X|
		// |XX|
		// +--+
		{
			[][]int{
				{0, 1},
				{0, 1},
				{1, 1},
			},
			[]int{0, 0},
			3,
			2,
		},
	},
}