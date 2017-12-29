package event

// DevicePaired is emitted when transport paired with a device (e.g. iOS client successfully paired with the accessory)
type DevicePaired struct{}

// DeviceUnpaired is emitted when pairing with a device is removed (e.g. iOS client removed the accessory from HomeKit)
type DeviceUnpaired struct{}

// ContainerAccessoryAdded is emitted when an accessory is added to an accessory.Container
type ContainerAccessoryAdded struct{
	//Accessory *accessory.Accessory
	Accessory interface{}
}

// ContainerAccessoryAdded is emitted when an accessory is removed from an accessory.Container
type ContainerAccessoryRemoved struct{
	Accessory interface{}
}
