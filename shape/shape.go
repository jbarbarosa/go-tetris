package shape

type Shape struct {
	Rows [4][4]byte
}

func O() Shape {
	return Shape{
		[4][4]byte{
			{0, 0, 0, 0},
			{0, 1, 1, 0},
			{0, 1, 1, 0},
			{0, 0, 0, 0},
		},
	}
}

func T() Shape {
	return Shape{
		[4][4]byte{
			{0, 0, 0, 0},
			{0, 0, 0, 0},
			{0, 1, 0, 0},
			{1, 1, 1, 0},
		},
	}
}

func I() Shape {
	return Shape{
		[4][4]byte{
			{0, 1, 0, 0},
			{0, 1, 0, 0},
			{0, 1, 0, 0},
			{0, 1, 0, 0},
		},
	}
}

func L() Shape {
	return Shape{
		[4][4]byte{
			{0, 0, 0, 0},
			{0, 1, 0, 0},
			{0, 1, 0, 0},
			{0, 1, 1, 0},
		},
	}
}

func J() Shape {
	return Shape{
		[4][4]byte{
			{0, 0, 0, 0},
			{0, 0, 1, 0},
			{0, 0, 1, 0},
			{0, 1, 1, 0},
		},
	}
}

func S() Shape {
	return Shape{
		[4][4]byte{
			{0, 0, 0, 0},
			{0, 0, 0, 0},
			{0, 0, 1, 1},
			{0, 1, 1, 0},
		},
	}
}

func Z() Shape {
	return Shape{
		[4][4]byte{
			{0, 0, 0, 0},
			{0, 0, 0, 0},
			{1, 1, 0, 0},
			{0, 1, 1, 0},
		},
	}
}
