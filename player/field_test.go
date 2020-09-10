package player

import (
	"testing"
)

func (f *field) testGetYPlay(test *testing.T, t *tetromino, r int, x int, expectedY int) {
	p := &t.pieces[r]
	if y := f.getY(p, x); y != expectedY {
		test.Errorf("tetromino [%s,r:%d,x:%d]: expected:%d, got:%d",
			t.name, r, x, expectedY, y)
	}
}

func TestGetYPlay(test *testing.T) {
	f := newField(fieldHeight, fieldWidth)

	f.setRow(18, []int{0, 1})
	f.setRow(19, []int{1, 1, 1})

	f.testGetYPlay(test, &tetrominoJ, 1, 0, 18)
	f.testGetYPlay(test, &tetrominoJ, 1, 1, 18)
	f.testGetYPlay(test, &tetrominoJ, 1, 2, 19)
	f.testGetYPlay(test, &tetrominoJ, 1, 3, 20)

	f.testGetYPlay(test, &tetrominoT, 2, 0, 18)
	f.testGetYPlay(test, &tetrominoT, 2, 1, 19)
	f.testGetYPlay(test, &tetrominoT, 2, 2, 20)
}

func TestGetYPlayWell(test *testing.T) {
	f := newField(fieldHeight, fieldWidth)

	f.setRow(16, []int{1, 0, 1})
	f.setRow(17, []int{1, 0, 1})
	f.setRow(18, []int{1, 0, 1})
	f.setRow(19, []int{1, 0, 1})

	f.testGetYPlay(test, &tetrominoI, 1, 0, 16)
	f.testGetYPlay(test, &tetrominoI, 1, 1, 20)
	f.testGetYPlay(test, &tetrominoI, 1, 2, 16)
}

func TestPush(test *testing.T) {
	f := newField(fieldHeight, fieldWidth)

	f.push(&tetrominoJ.pieces[1], 0)
	f.push(&tetrominoO.pieces[0], 0)

	if !f.eqRow(15, []int{1, 1}) ||
		!f.eqRow(16, []int{1, 1}) ||
		!f.eqRow(17, []int{1, 0}) ||
		!f.eqRow(18, []int{1, 0}) ||
		!f.eqRow(19, []int{1, 1}) {
		test.Errorf("TestPush row")
	}
	if f.col[0] != 5 || f.col[1] != 5 {
		test.Errorf("TestPush col")
	}
}

func TestInvalidPush(test *testing.T) {
	f := newField(fieldHeight, fieldWidth)

	for i := 0; i < 5; i++ {
		if !f.push(&tetrominoI.pieces[1], 0) {
			test.Errorf("Expected push succeed")
		}
	}

	if f.push(&tetrominoI.pieces[1], 0) {
		test.Errorf("Expected push failed")
	}

	if !f.push(&tetrominoI.pieces[0], 6) {
		test.Errorf("Expected push succeed")
	}
	if f.push(&tetrominoI.pieces[0], 7) {
		test.Errorf("Expected push failed")
	}
}

func (f *field) fieldIsEmpty(test *testing.T) bool {
	for x := 0; x < f.width; x++ {
		if f.col[x] != 0 {
			test.Errorf("col[%d] expected:0, got:%d", x, f.col[x])
			return false
		}
	}
	for y := 0; y < f.height; y++ {
		if f.row[y] != 0 {
			test.Errorf("row[%d] expected:0, got:%d", y, f.row[y])
			return false
		}
	}
	return true
}

func TestCompleteRow(test *testing.T) {
	f := newField(fieldHeight, fieldWidth)

	f.push(&tetrominoI.pieces[0], 0)
	f.push(&tetrominoI.pieces[0], 4)
	f.push(&tetrominoO.pieces[0], 8)
	f.push(&tetrominoI.pieces[0], 0)
	f.push(&tetrominoI.pieces[0], 4)
	if !f.fieldIsEmpty(test) {
		return
	}

	for x := 0; x < f.width; x++ {
		f.push(&tetrominoI.pieces[1], x)
	}
	if !f.fieldIsEmpty(test) {
		return
	}
}
