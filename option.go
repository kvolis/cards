package cards

type option uint8

const (
	lowOptMask  = 0b01000000
	highOptMask = 0b10000000
)

// LOpt returns the value of the 6th bit (the Lower of the two free bits)
func (c *Card) LOpt() bool {
	return option(*c)&lowOptMask == lowOptMask
}

// HOpt returns the value of the 7th bit (the Higher of the two free bits)
func (c *Card) HOpt() bool {
	return option(*c)&highOptMask == highOptMask
}

// LOpt returns the value of the 6th bit (the Lower of the two free bits)
func LOpt(c *Card) bool {
	return c.LOpt()
}

// HOpt returns the value of the 7th bit (the Higher of the two free bits)
func HOpt(c *Card) bool {
	return c.HOpt()
}

// SetLOpt sets the value of the 6th bit (the Lower of the two free bits)
func (c *Card) SetLOpt(val bool) {
	c.setOpt(val, 6)
}

// SetHOpt sets the value of the 7th bit (the Higher of the two free bits)
func (c *Card) SetHOpt(val bool) {
	c.setOpt(val, 7)
}

// SetLOpt sets the value of the 6th bit (the Lower of the two free bits)
func SetLOptTo(c *Card, val bool) {
	c.SetLOpt(val)
}

// SetHOpt sets the value of the 7th bit (the Higher of the two free bits)
func SetHOptTo(c *Card, val bool) {
	c.SetHOpt(val)
}

func (c *Card) setOpt(val bool, pos uint8) {
	if val {
		c.setBit(pos)
		return
	}
	c.clearBit(pos)
}

func (c *Card) setBit(pos uint8) {
	mask := 1 << pos
	*c |= Card(mask)
}

func (c *Card) clearBit(pos uint8) {
	mask := ^(1 << pos)
	*c &= Card(mask)
}
