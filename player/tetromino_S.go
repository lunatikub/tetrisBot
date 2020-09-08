package player

var tetrominoS = tetromino{
	"S",
	[]piece{

		// +---+
		// |.XX|
		// |XX.|
		// +---+
		{
			[][]int{
				{0, 1, 1},
				{1, 1, 0},
			},
			[]int{0, 0, 1},
			2,
			3,
		},

		// +--+
		// |X.|
		// |XX|
		// |.X|
		// +--+
		{
			[][]int{
				{1, 0},
				{1, 1},
				{0, 1},
			},
			[]int{1, 0},
			3,
			2,
		},
	},
}