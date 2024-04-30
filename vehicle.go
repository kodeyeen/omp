package gomp

// #include <stdlib.h>
// #include <string.h>
// #include "include/player.h"
// #include "include/vehicle.h"
import "C"
import (
	"errors"
	"time"
	"unsafe"
)

type VehicleColor struct {
	Primary   Color
	Secondary Color
}

func NewRandomVehicleColor() *VehicleColor {
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
	cVeh := C.vehicle_create(0, C.int(model), C.float(pos.X), C.float(pos.Y), C.float(pos.Z), C.float(angle), 0, 0, -1, 0)
	if cVeh == nil {
		return nil, errors.New("vehicle limit reached")
	}

	return &Vehicle{handle: cVeh}, nil
}

func NewStaticVehicle(model VehicleModel, pos Vector3, angle float32) (*Vehicle, error) {
	vehicle := C.vehicle_create(1, C.int(model), C.float(pos.X), C.float(pos.Y), C.float(pos.Z), C.float(angle), 0, 0, -1, 0)

	return &Vehicle{handle: vehicle}, nil
}

func FreeVehicle(veh *Vehicle) {
	C.vehicle_release(veh.handle)
}

// Gets an array variable of the IDs of the created vehicles on the server.
func Vehicles() []*Vehicle {
	panic("not implemented")
}

// Checks if a vehicle is streamed in for a player. Only nearby vehicles are streamed in (visible) for a player.
func (v *Vehicle) IsStreamedInFor(player *Player) bool {
	return C.vehicle_isStreamedInForPlayer(v.handle, player.handle) != 0
}

func (v *Vehicle) SetColor(color VehicleColor) {
	C.vehicle_setColour(v.handle, C.int(color.Primary), C.int(color.Secondary))
}

func (v *Vehicle) Color() VehicleColor {
	colour := C.vehicle_getColour(v.handle)

	return VehicleColor{
		Primary:   Color(colour.primary),
		Secondary: Color(colour.secondary),
	}
}

// Set a vehicle's health. When a vehicle's health decreases the engine will produce smoke, and finally fire when it decreases to less than 250 (25%).
func (v *Vehicle) SetHealth(health float32) {
	C.vehicle_setHealth(v.handle, C.float(health))
}

func (v *Vehicle) Health() float32 {
	return float32(C.vehicle_getHealth(v.handle))
}

// Get the player driving the vehicle.
func (v *Vehicle) Driver() (*Player, error) {
	driver := C.vehicle_getDriver(v.handle)

	if driver == nil {
		return nil, errors.New("vehicle has no driver")
	}

	return &Player{handle: driver}, nil
}

func (v *Vehicle) Passengers() []*Player {
	cPassengerArr := C.vehicle_getPassengers(v.handle)
	defer C.freeArray(cPassengerArr)

	passengerHandles := unsafe.Slice(cPassengerArr.buf, int(cPassengerArr.length))
	passengers := make([]*Player, 0, len(passengerHandles))

	for _, handle := range passengerHandles {
		passengers = append(passengers, &Player{handle: handle})
	}

	return passengers
}

func (v *Vehicle) SetNumberPlate(numberPlate string) {
	cNumberPlate := C.CString(numberPlate)
	defer C.free(unsafe.Pointer(cNumberPlate))

	C.vehicle_setPlate(v.handle, C.String{
		buf:    cNumberPlate,
		length: C.strlen(cNumberPlate),
	})
}

func (v *Vehicle) NumberPlate() string {
	numberPlate := C.vehicle_getPlate(v.handle)

	return C.GoStringN(numberPlate.buf, C.int(numberPlate.length))
}

// Sets the various visual damage statuses of a vehicle, such as popped tires, broken lights and damaged panels.
func (v *Vehicle) SetDamageStatus(damageStatus VehicleDamageStatus) {
	panic("not implemented")
}

func (v *Vehicle) DamageStatus() VehicleDamageStatus {
	panic("not implemented")
}

func (v *Vehicle) SetPaintjob(paintjob int) {
	C.vehicle_setPaintjob(v.handle, C.int(paintjob))
}

func (v *Vehicle) Paintjob() int {
	return int(C.vehicle_getPaintjob(v.handle))
}

// Adds a 'component' (often referred to as a 'mod' (modification)) to a vehicle.
// Valid components can be found here: https://www.open.mp/docs/scripting/resources/carcomponentid
func (v *Vehicle) AddComponent(componentID int) {
	// TODO: component constants
	C.vehicle_addComponent(v.handle, C.int(componentID))
}

func (v *Vehicle) ComponentInSlot(slot int) int {
	return int(C.vehicle_getComponentInSlot(v.handle, C.int(slot)))
}

// Remove a component from a vehicle.
func (v *Vehicle) RemoveComponent(componentID int) {
	C.vehicle_removeComponent(v.handle, C.int(componentID))
}

func (v *Vehicle) PutPlayer(plr *Player, seatID int) {
	C.vehicle_putPlayer(v.handle, plr.handle, C.int(seatID))
}

// Set the Z rotation (yaw) of a vehicle.
func (v *Vehicle) SetZAngle(zAngle float32) {
	C.vehicle_setZAngle(v.handle, C.float(zAngle))
}

// Get the rotation of a vehicle on the Z axis (yaw).
func (v *Vehicle) ZAngle() float32 {
	return float32(C.vehicle_getZAngle(v.handle))
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

	return params.engine != 0
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

	return params.lights != 0
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

	return params.alarm != 0
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

	return params.doors != 0
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

	return params.bonnet != 0
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

	return params.boot != 0
}

// Allows you to open and close the doors of a vehicle.
func (v *Vehicle) SetDoorsState(state VehicleDoorsState) {
	panic("not implemented")
}

// Allows you to retrieve the current state of a vehicle's doors.
func (v *Vehicle) DoorsState() VehicleDoorsState {
	panic("not implemented")
}

// Allows you to open and close the windows of a vehicle.
func (v *Vehicle) SetWindowsState(state VehicleWindowsState) {
	panic("not implemented")
}

// Allows you to retrieve the current state of a vehicle's windows.
func (v *Vehicle) WindowsState() VehicleWindowsState {
	panic("not implemented")
}

// Check if a vehicle is dead.
func (v *Vehicle) IsDead() bool {
	return C.vehicle_isDead(v.handle) != 0
}

func (v *Vehicle) Respawn() {
	C.vehicle_respawn(v.handle)
}

// Set the respawn delay of a vehicle.
func (v *Vehicle) SetRespawnDelay(respawnDelay time.Duration) {
	C.vehicle_setRespawnDelay(v.handle, C.int(respawnDelay.Seconds()))
}

// Get the respawn delay of a vehicle in seconds.
func (v *Vehicle) RespawnDelay() time.Duration {
	return time.Duration(C.vehicle_getRespawnDelay(v.handle)) * time.Second
}

func (v *Vehicle) IsRespawning() bool {
	return C.vehicle_isRespawning(v.handle) != 0
}

func (v *Vehicle) SetInterior(interior int) {
	C.vehicle_setInterior(v.handle, C.int(interior))
}

func (v *Vehicle) Interior() int {
	return int(C.vehicle_getInterior(v.handle))
}

func (v *Vehicle) AttachTrailer(trailer *Vehicle) {
	C.vehicle_attachTrailer(v.handle, trailer.handle)
}

func (v *Vehicle) DetachTrailer() {
	C.vehicle_detachTrailer(v.handle)
}

func (v *Vehicle) IsTrailer() bool {
	return C.vehicle_isTrailer(v.handle) != 0
}

func (v *Vehicle) Trailer() (*Vehicle, error) {
	trailer := C.vehicle_getTrailer(v.handle)

	if trailer == nil {
		return nil, errors.New("vehicle has no trailer")
	}

	return &Vehicle{handle: trailer}, nil
}

func (v *Vehicle) Cab() *Vehicle {
	cab := C.vehicle_getCab(v.handle)

	return &Vehicle{handle: cab}
}

// Fully repairs a vehicle, including visual damage (bumps, dents, scratches, popped tires etc.).
func (v *Vehicle) Repair() {
	C.vehicle_repair(v.handle)
}

// Sets the X, Y and Z velocity of a vehicle.
func (v *Vehicle) SetVelocity(vel Vector3) {
	C.vehicle_setVelocity(v.handle, C.float(vel.X), C.float(vel.Y), C.float(vel.Z))
}

// Get the velocity of a vehicle on the X, Y and Z axes.
func (v *Vehicle) Velocity() Vector3 {
	vel := C.vehicle_getVelocity(v.handle)

	return Vector3{
		X: float32(vel.x),
		Y: float32(vel.y),
		Z: float32(vel.z),
	}
}

// Sets the angular X, Y and Z velocity of a vehicle.
func (v *Vehicle) SetAngularVelocity(vel Vector3) {
	C.vehicle_setAngularVelocity(v.handle, C.float(vel.X), C.float(vel.Y), C.float(vel.Z))
}

func (v *Vehicle) AngularVelocity() Vector3 {
	vel := C.vehicle_getAngularVelocity(v.handle)

	return Vector3{
		X: float32(vel.x),
		Y: float32(vel.y),
		Z: float32(vel.z),
	}
}

// Gets the model ID of a vehicle.
func (v *Vehicle) Model() VehicleModel {
	return VehicleModel(C.vehicle_getModel(v.handle))
}

// Gets the current vehicle landing gear state from the latest driver.
func (v *Vehicle) LandingGearState() int {
	return int(C.vehicle_getLandingGearState(v.handle))
}

// Check if the vehicle was occupied since last spawn.
func (v *Vehicle) HasBeenOccupied() bool {
	return C.vehicle_hasBeenOccupied(v.handle) != 0
}

func (v *Vehicle) LastOccupiedAt() time.Time {
	millis := C.vehicle_getLastOccupiedTime(v.handle)

	return time.Unix(0, int64(millis)*int64(time.Millisecond))
}

func (v *Vehicle) LastSpawnedAt() time.Time {
	millis := C.vehicle_getLastSpawnTime(v.handle)

	return time.Unix(0, int64(millis)*int64(time.Millisecond))
}

// Check if vehicle is occupied.
func (v *Vehicle) IsOccupied() bool {
	return C.vehicle_isOccupied(v.handle) != 0
}

// Turn the siren for a vehicle on.
func (v *Vehicle) EnableSiren() {
	C.vehicle_setSiren(v.handle, 1)
}

// Turn the siren for a vehicle off.
func (v *Vehicle) DisableSiren() {
	C.vehicle_setSiren(v.handle, 0)
}

func (v *Vehicle) IsSirenEnabled() bool {
	return C.vehicle_getSirenState(v.handle) != 0
}

func (v *Vehicle) HydraThrustAngle() int {
	return int(C.vehicle_getHydraThrustAngle(v.handle))
}

func (v *Vehicle) TrainSpeed() float32 {
	return float32(C.vehicle_getTrainSpeed(v.handle))
}

// Get the last driver of a vehicle.
func (v *Vehicle) LastDriver() (*Player, error) {
	lastDriverID := C.vehicle_getLastDriverPoolID(v.handle)
	lastDriver := C.player_getByID(lastDriverID)

	if lastDriver == nil {
		return nil, errors.New("vehicle has no last driver")
	}

	return &Player{handle: lastDriver}, nil
}

// Set a vehicle's position.
func (v *Vehicle) SetPosition(pos Vector3) {
	C.vehicle_setPosition(v.handle, C.float(pos.X), C.float(pos.Y), C.float(pos.Z))
}

func (v *Vehicle) Position() Vector3 {
	pos := C.vehicle_getPosition(v.handle)

	return Vector3{
		X: float32(pos.x),
		Y: float32(pos.y),
		Z: float32(pos.z),
	}
}

// Returns a vehicle's rotation on all axes as a quaternion.
func (v *Vehicle) Rotation() Vector4 {
	rquat := C.vehicle_getRotation(v.handle)

	return Vector4{
		X: float32(rquat.x),
		Y: float32(rquat.y),
		Z: float32(rquat.z),
		W: float32(rquat.w),
	}
}

// Sets the 'virtual world' of a vehicle. Players will only be able to see vehicles in their own virtual world.
func (v *Vehicle) SetVirtualWorld(virtualWorld int) {
	C.vehicle_setVirtualWorld(v.handle, C.int(virtualWorld))
}

// Get the virtual world of a vehicle.
func (v *Vehicle) VirtualWorld() int {
	return int(C.vehicle_getVirtualWorld(v.handle))
}

func (v *Vehicle) DistanceFrom(point Vector3) float32 {
	return float32(C.vehicle_getDistanceFromPoint(v.handle, C.float(point.X), C.float(point.Y), C.float(point.Z)))
}

func (v *Vehicle) IsInRangeOf(point Vector3, _range float32) bool {
	return C.vehicle_isInRangeOfPoint(v.handle, C.float(_range), C.float(point.X), C.float(point.Y), C.float(point.Z)) != 0
}

func (v *Vehicle) Matrix() {
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

// Sets the vehicle to dead.
func (v *Vehicle) SetDead(dead bool) {
	panic("not implemented")
}

// TODO
// SetParamsForPlayer
