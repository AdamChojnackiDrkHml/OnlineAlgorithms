package graphs

type Graph interface {
	Distance(uint8, uint8) uint8
	ResourceLocation() uint8
	SetResourceLocation(uint8)
	Request(uint8) uint8
	MoveResource(uint8) uint8
}

func min(a, b uint8) uint8 {
	if a < b {
		return a
	}

	return b
}

func absDiff(a, b uint8) uint8 {
	if a < b {
		return b - a
	}

	return a - b
}

func CreateGraph() Graph {
	return Torus_newTorus()
}
