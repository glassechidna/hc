package accessory

type Bridge struct {
	*Accessory
	*Container
}

// NewBridge returns a bridge which implements model.Bridge.
func NewBridge(info Info, container *Container) *Bridge {
	acc := Bridge{}
	acc.Accessory = New(info, TypeBridge)

	acc.Container = container
	container.AddAccessory(acc.Accessory)

	return &acc
}
