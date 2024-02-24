package videoContent

type VideoContent struct {
	contentUri   string
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
