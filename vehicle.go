package gomp

// #include "include/vehicle.h"
import "C"
import (
	"time"
	"unsafe"
)

type VehicleColor struct {
	Primary   int
	Secondary int
}

func RandomVehicleColor() *VehicleColor {
	panic("not implemented")
}

type VehicleDamageStatus struct {
	Panels int
	Doors  int
	Lights int
	Tyres  int
}

type VehicleDoorsState struct {
	FrontLeft  bool
	FrontRight bool
	RearLeft   bool
	RearRight  bool
}

type VehicleWindowsState struct {
	FrontLeft  bool
	FrontRight bool
	RearLeft   bool
	RearRight  bool
}

type Vehicle struct {
	handle unsafe.Pointer
}

func NewVehicle(model VehicleModel, pos Vector3, angle float32) (*Vehicle, error) {
	// TODO: error handling (invalid modelID and trains)
	handle := C.vehicle_create(C.int(0), C.int(model), C.float(pos.X), C.float(pos.Y), C.float(pos.Z), C.float(angle), C.int(0), C.int(0), C.int(-1), C.int(0))

	veh := &Vehicle{
		handle: handle,
	}

	return veh, nil
}

func FreeVehicle(veh *Vehicle) error {
	panic("not implemented")
}

// Gets an array variable of the IDs of the created vehicles on the server.
func Vehicles() []*Vehicle {
	panic("not implemented")
}

func (v *Vehicle) SetColor(primary int, secondary int) {
	C.vehicle_setColour(v.handle, C.int(primary), C.int(secondary))
}

func (v *Vehicle) Color() *VehicleColor {
	colour := C.vehicle_getColour(v.handle)

	return &VehicleColor{
		Primary:   int(colour.primary),
		Secondary: int(colour.secondary),
	}
}

func (v *Vehicle) SetPaintjob(paintjob int) {
	panic("not implemented")
}

func (v *Vehicle) Paintjob() int {
	panic("not implemented")
}

// Adds a 'component' (often referred to as a 'mod' (modification)) to a vehicle.
// Valid components can be found here: https://www.open.mp/docs/scripting/resources/carcomponentid
func (v *Vehicle) AddComponent(componentID int) {
	// TODO: component constants
	panic("not implemented")
}

// Remove a component from a vehicle.
func (v *Vehicle) RemoveComponent(componentID int) {
	panic("not implemented")
}

func (v *Vehicle) AttachTrailer(trailer *Vehicle) {
	panic("not implemented")
}

func (v *Vehicle) DetachTrailer() {
	panic("not implemented")
}

func (v *Vehicle) Trailer() *Vehicle {
	panic("not implemented")
}

func (v *Vehicle) ComponentInSlot(slot int) {
	// TODO: slot constants
	panic("not implemented")
}

func (v *Vehicle) DamageStatus() *VehicleDamageStatus {
	panic("not implemented")
}

// Sets the various visual damage statuses of a vehicle, such as popped tires, broken lights and damaged panels.
func (v *Vehicle) SetDamageStatus(damageStatus *VehicleDamageStatus) {
	panic("not implemented")
}

func (v *Vehicle) DistanceFromPoint(x, y, z float32) float32 {
	panic("not implemented")
}

// Get the player driving the vehicle.
func (v *Vehicle) Driver() (*Player, error) {
	panic("not implemented")
}

func (v *Vehicle) Health() float32 {
	panic("not implemented")
}

// Set a vehicle's health. When a vehicle's health decreases the engine will produce smoke, and finally fire when it decreases to less than 250 (25%).
func (v *Vehicle) SetHealth(health float32) {
	panic("not implemented")
}

func (v *Vehicle) HydraReactorAngle() float32 {
	panic("not implemented")
}

func (v *Vehicle) InteriorID() int {
	panic("not implemented")
}

func (v *Vehicle) SetInteriorID(interiorID int) {
	panic("not implemented")
}

// Gets the current vehicle landing gear state from the latest driver.
func (v *Vehicle) LandingGearState() int {
	panic("not implemented")
}

// Get the last driver of a vehicle.
func (v *Vehicle) LastDriver() (*Player, error) {
	panic("not implemented")
}

func (v *Vehicle) Matrix() {
	panic("not implemented")
}

// Gets the model ID of a vehicle.
func (v *Vehicle) Model() VehicleModel {
	panic("not implemented")
}

func (v *Vehicle) NumberPlate() string {
	panic("not implemented")
}

func (v *Vehicle) SetNumberPlate(numberPlate string) {
	panic("not implemented")
}

// Get the occupied tick of a vehicle.
func (v *Vehicle) OccupiedTick() int {
	panic("not implemented")
}

// Set the occupied tick of a vehicle.
func (v *Vehicle) SetOccupiedTick(occupiedTick int) {
	panic("not implemented")
}

// Allows you to retrieve the current state of a vehicle's doors.
func (v *Vehicle) DoorsState() *VehicleDoorsState {
	panic("not implemented")
}

// Allows you to open and close the doors of a vehicle.
func (v *Vehicle) SetDoorsState(state *VehicleDoorsState) {
	panic("not implemented")
}

// Allows you to retrieve the current state of a vehicle's windows.
func (v *Vehicle) WindowsState() *VehicleWindowsState {
	panic("not implemented")
}

// Allows you to open and close the windows of a vehicle.
func (v *Vehicle) SetWindowsState(state *VehicleWindowsState) {
	panic("not implemented")
}

func (v *Vehicle) StartEngine() {
	params := C.vehicle_getParams(v.handle)
	params.engine = C.schar(1)

	C.vehicle_setParams(v.handle, &params)
}

func (v *Vehicle) StopEngine() {
	params := C.vehicle_getParams(v.handle)
	params.engine = C.schar(0)

	C.vehicle_setParams(v.handle, &params)
}

func (v *Vehicle) IsEngineStarted() bool {
	params := C.vehicle_getParams(v.handle)

	return params.engine == C.schar(1)
}

func (v *Vehicle) TurnOnLights() {
	params := C.vehicle_getParams(v.handle)
	params.lights = C.schar(1)

	C.vehicle_setParams(v.handle, &params)
}

func (v *Vehicle) TurnOffLights() {
	params := C.vehicle_getParams(v.handle)
	params.lights = C.schar(0)

	C.vehicle_setParams(v.handle, &params)
}

func (v *Vehicle) AreLightsTurnedOn() bool {
	params := C.vehicle_getParams(v.handle)

	return params.lights == C.schar(1)
}

func (v *Vehicle) TurnOnAlarm() {
	params := C.vehicle_getParams(v.handle)
	params.alarm = C.schar(1)

	C.vehicle_setParams(v.handle, &params)
}

func (v *Vehicle) TurnOffAlarm() {
	params := C.vehicle_getParams(v.handle)
	params.alarm = C.schar(0)

	C.vehicle_setParams(v.handle, &params)
}

func (v *Vehicle) IsAlarmTurnedOn() bool {
	params := C.vehicle_getParams(v.handle)

	return params.alarm == C.schar(1)
}

func (v *Vehicle) LockDoors() {
	params := C.vehicle_getParams(v.handle)
	params.doors = C.schar(1)

	C.vehicle_setParams(v.handle, &params)
}

func (v *Vehicle) UnlockDoors() {
	params := C.vehicle_getParams(v.handle)
	params.doors = C.schar(0)

	C.vehicle_setParams(v.handle, &params)
}

func (v *Vehicle) AreDoorsLocked() bool {
	params := C.vehicle_getParams(v.handle)

	return params.doors == C.schar(1)
}

func (v *Vehicle) OpenHood() {
	params := C.vehicle_getParams(v.handle)
	params.bonnet = C.schar(1)

	C.vehicle_setParams(v.handle, &params)
}

func (v *Vehicle) CloseHood() {
	params := C.vehicle_getParams(v.handle)
	params.bonnet = C.schar(0)

	C.vehicle_setParams(v.handle, &params)
}

func (v *Vehicle) IsHoodOpen() bool {
	params := C.vehicle_getParams(v.handle)

	return params.bonnet == C.schar(1)
}

func (v *Vehicle) OpenTrunk() {
	params := C.vehicle_getParams(v.handle)
	params.boot = C.schar(1)

	C.vehicle_setParams(v.handle, &params)
}

func (v *Vehicle) CloseTrunk() {
	params := C.vehicle_getParams(v.handle)
	params.boot = C.schar(0)

	C.vehicle_setParams(v.handle, &params)
}

func (v *Vehicle) IsTrunkOpen() bool {
	params := C.vehicle_getParams(v.handle)

	return params.boot == C.schar(1)
}

// Returns a vehicle's siren state (on/off).
func (v *Vehicle) SirenState() bool {
	panic("not implemented")
}

// Turn the siren for a vehicle on or off.
func (v *Vehicle) SetSirenState(state int) {
	panic("not implemented")
}

func (v *Vehicle) Passengers() []*Player {
	cpassengerArr := C.vehicle_getPassengers(v.handle)
	defer C.freeArray(cpassengerArr)

	passengerHandles := unsafe.Slice(cpassengerArr.buf, int(cpassengerArr.length))
	passengers := make([]*Player, 0, len(passengerHandles))

	for _, handle := range passengerHandles {
		passengers = append(passengers, &Player{handle})
	}

	return passengers
}

func (v *Vehicle) Position() *Position {
	panic("not implemented")
}

func (v *Vehicle) PutPlayer(plr *Player, seatID int) {
	panic("not implemented")
}

// Set a vehicle's position.
func (v *Vehicle) SetPosition(pos *Position) {
	panic("not implemented")
}

// Get the respawn delay of a vehicle in seconds.
func (v *Vehicle) RespawnDelay() time.Duration {
	return 1 * time.Second
}

// Set the respawn delay of a vehicle.
func (v *Vehicle) SetRespawnDelay(respawnDelay time.Duration) {
	panic("not implemented")
}

// Returns respawn tick in milliseconds.
func (v *Vehicle) RespawnTick() time.Duration {
	panic("not implemented")
}

// Set the respawn tick of a vehicle in milliseconds.
func (v *Vehicle) SetRespawnTick(respawnTick time.Duration) {
	panic("not implemented")
}

// Returns a vehicle's rotation on all axes as a quaternion.
func (v *Vehicle) RotationQuat() {
	panic("not implemented")
}

func (v *Vehicle) TrainSpeed() int {
	panic("not implemented")
}

// Get the velocity of a vehicle on the X, Y and Z axes.
func (v *Vehicle) Velocity() *Position {
	panic("not implemented")
}

// Sets the X, Y and Z velocity of a vehicle.
func (v *Vehicle) SetVelocity(velocity *Position) {
	panic("not implemented")
}

// Get the virtual world of a vehicle.
func (v *Vehicle) VirtualWorld() int {
	panic("not implemented")
}

// Sets the 'virtual world' of a vehicle. Players will only be able to see vehicles in their own virtual world.
func (v *Vehicle) SetVirtualWorld(virtualWorld int) {
	panic("not implemented")
}

// Get the rotation of a vehicle on the Z axis (yaw).
func (v *Vehicle) ZAngle() float32 {
	panic("not implemented")
}

// Set the Z rotation (yaw) of a vehicle.
func (v *Vehicle) SetZAngle(zAngle float32) {
	panic("not implemented")
}

// Check if the vehicle was occupied since last spawn.
func (v *Vehicle) HasBeenOccupied() bool {
	panic("not implemented")
}

// Check if vehicle is occupied.
func (v *Vehicle) IsOccupied() bool {
	panic("not implemented")
}

// Sets the vehicle's occupancy.
func (v *Vehicle) Occupy(occupied bool) {
	panic("not implemented")
}

// Hides a vehicle from the game.
func (v *Vehicle) Hide() {
	panic("not implemented")
}

// Shows the hidden vehicle.
func (v *Vehicle) Show() {
	panic("not implemented")
}

// Checks if a vehicle is hidden.
func (v *Vehicle) IsHidden() bool {
	panic("not implemented")
}

// Checks if a vehicle has a trailer attached to it.
func (v *Vehicle) HasTrailer() bool {
	panic("not implemented")
}

// Check if a vehicle is dead.
func (v *Vehicle) IsDead() bool {
	panic("not implemented")
}

// Sets the vehicle to dead.
func (v *Vehicle) SetDead(dead bool) {
	panic("not implemented")
}

// Checks if a vehicle siren is on or off.
func (v *Vehicle) IsSirenEnabled() bool {
	panic("not implemented")
}

// Checks if a vehicle is streamed in for a player. Only nearby vehicles are streamed in (visible) for a player.
func (v *Vehicle) IsStreamedIn(forPlayer *Player) bool {
	panic("not implemented")
}

// Fully repairs a vehicle, including visual damage (bumps, dents, scratches, popped tires etc.).
func (v *Vehicle) Repair() {
	panic("not implemented")
}

// Sets the angular X, Y and Z velocity of a vehicle.
func (v *Vehicle) SetAngularVelocity(velocity *Position) {
	panic("not implemented")
}

// Sets a vehicle back to the position at where it was created.
func (v *Vehicle) SetToRespawn() {
	panic("not implemented")
}
