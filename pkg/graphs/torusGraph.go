package graphs

type Torus struct {
	serverLocation uint8
}

func Torus_newTorus() *Torus {
	return &Torus{serverLocation: 0}
}

func (t *Torus) Distance(v1, v2 uint8) uint8 {
	v1Rep := t.verticeIndexToCoordinates(v1)
	v2Rep := t.verticeIndexToCoordinates(v2)

	dist := uint8(0)

	for i := range v1Rep {
		diff := absDiff(v1Rep[i], v2Rep[i])
		diffAlt := 4 - diff
		dist += min(diff, diffAlt)
	}

	return dist
}

func (t *Torus) ResourceLocation() uint8 {
	return t.serverLocation
}

func (t *Torus) SetResourceLocation(newLoc uint8) {
	t.serverLocation = newLoc
}

func (t *Torus) Request(req uint8) uint8 {
	return t.Distance(t.serverLocation, req)
}

func (t *Torus) MoveResource(loc uint8) uint8 {
	prevLoc := t.serverLocation
	t.SetResourceLocation(loc)
	return t.Distance(prevLoc, loc)
}

func (t *Torus) verticeIndexToCoordinates(v uint8) []uint8 {
	fst := uint8(0b00110000)
	snd := uint8(0b00001100)
	trd := uint8(0b00000011)

	return []uint8{
		(fst & v) >> 4,
		(snd & v) >> 2,
		(trd & v),
	}
}
