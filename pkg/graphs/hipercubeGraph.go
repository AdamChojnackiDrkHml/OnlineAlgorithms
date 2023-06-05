package graphs

type Hypercube struct {
	serverLocation uint8
}

func Hypercube_newHypercube() *Hypercube {
	return &Hypercube{serverLocation: 0}
}

func (h *Hypercube) Distance(v1, v2 uint8) uint8 {
	dist := uint8(0)

	for i := 0; i < 6; i++ {
		dist += (v1&0b1 ^ v2&0b1)
		v1 = v1 >> 1
		v2 = v2 >> 1
	}

	return dist
}

func (h *Hypercube) ResourceLocation() uint8 {
	return h.serverLocation
}

func (h *Hypercube) SetResourceLocation(newLoc uint8) {
	h.serverLocation = newLoc
}

func (h *Hypercube) Request(req uint8) uint8 {
	return h.Distance(h.serverLocation, req)
}

func (h *Hypercube) MoveResource(loc uint8) uint8 {
	prevLoc := h.serverLocation
	h.SetResourceLocation(loc)
	return h.Distance(prevLoc, loc)
}
