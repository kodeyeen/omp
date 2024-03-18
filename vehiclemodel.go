package gomp

const (
	Landstalker VehicleModel = iota + 400
	Bravura
	Buffalo
	// TODO
)

type VehicleModel int

// Get the number of used vehicle models on the server.
func VehicleModelsUsed() int {
	panic("not implemented")
}

func VehicleCountForModel(model VehicleModel) int {
	panic("not implemented")
}

func (m VehicleModel) Size() *Vector3 {
	// use GetVehicleModelInfo
	panic("not implemented")
}

func (m VehicleModel) FrontSeatPosition() *Vector3 {
	// use GetVehicleModelInfo
	panic("not implemented")
}

func (m VehicleModel) RearSeatPosition() *Vector3 {
	// use GetVehicleModelInfo
	panic("not implemented")
}

func (m VehicleModel) PetrolCapPosition() *Vector3 {
	// use GetVehicleModelInfo
	panic("not implemented")
}

func (m VehicleModel) FrontWheelsPosition() *Vector3 {
	// use GetVehicleModelInfo
	panic("not implemented")
}

func (m VehicleModel) RearWheelsPosition() *Vector3 {
	// use GetVehicleModelInfo
	panic("not implemented")
}

func (m VehicleModel) MiddleWheelsPosition() *Vector3 {
	// use GetVehicleModelInfo
	panic("not implemented")
}

func (m VehicleModel) FrontBumperHeight() float32 {
	// use GetVehicleModelInfo
	panic("not implemented")
}

func (m VehicleModel) RearBumperHeight() float32 {
	// use GetVehicleModelInfo
	panic("not implemented")
}
