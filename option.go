package cards

type option uint8

const (
	lowOptMask  = 0b01000000
	highOptMask = 0b10000000
)

func (c *Card) OptL() bool {
	return option(*c)&lowOptMask == lowOptMask
}

func (c *Card) OptH() bool {
	return option(*c)&highOptMask == highOptMask
}

func OptL(c *Card) bool {
	return c.OptL()
}

func OptH(c *Card) bool {
	return c.OptH()
}

func (c *Card) SetOptL(val bool) {
	c.setOpt(val, 6)
}

func (c *Card) SetOptH(val bool) {
	c.setOpt(val, 7)
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
