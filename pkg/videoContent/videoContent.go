package videoContent

type VideoContent struct {
	imageUri     string
	second       int
	isBackground bool
	layer        int
	scale        float32
	position     Position
}

type Position struct {
	x float32
	y float32
}
