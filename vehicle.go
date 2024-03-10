package main

// #include "component.h"
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
	return nil
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

type VehicleParams struct {
	Engine    bool
	Lights    bool
	Alarm     bool
	Doors     bool
	Bonnet    bool
	Boot      bool
	Objective bool
}

type VehicleSpawnInfo struct {
}

type Vehicle struct {
	handle unsafe.Pointer
}

func NewVehicle(modelID int, pos Position, angle float32) (*Vehicle, error) {
	// TODO: error handling (invalid modelID and trains)
	handle := C.vehicle_create(0, C.int(modelID), C.float(pos.X), C.float(pos.Y), C.float(pos.Z), C.float(angle), C.int(0), C.int(0), C.int(-1), 0)

	veh := &Vehicle{
		handle: handle,
	}

	return veh, nil
}

func DestroyVehicle(veh *Vehicle) error {
	return nil
}

// Gets the model count of a vehicle model.
func VehicleModelCount(modelID int) int {
	return 0
}

// Retrieve information about a specific vehicle model such as the size or position of seats.
func VehicleModelInfo(modelID int, infotype int) *Position {
	return nil
}

// Get the number of used vehicle models on the server.
func VehicleModelsUsed() int {
	return 0
}

// Gets an array variable of the IDs of the created vehicles on the server.
func Vehicles() []*Vehicle {
	return nil
}

func (v *Vehicle) SetColor(color *VehicleColor) {

}

func (v *Vehicle) Color() *VehicleColor {
	return nil
}

func (v *Vehicle) SetPaintjob(paintjob int) {

}

func (v *Vehicle) Paintjob() int {
	return 0
}

func (v *Vehicle) AddComponent(componentID int) {
	// TODO: component constants
}

// Remove a component from a vehicle.
func (v *Vehicle) RemoveComponent(componentID int) {

}

func (v *Vehicle) AttachTrailer(trailer *Vehicle) {

}

func (v *Vehicle) DetachTrailer() {

}

func (v *Vehicle) Trailer() *Vehicle {
	return nil
}

func (v *Vehicle) ComponentInSlot(slot int) {
	// TODO: slot constants
}

func (v *Vehicle) DamageStatus() *VehicleDamageStatus {
	return nil
}

// Sets the various visual damage statuses of a vehicle, such as popped tires, broken lights and damaged panels.
func (v *Vehicle) SetDamageStatus(damageStatus *VehicleDamageStatus) {

}

func (v *Vehicle) DistanceFromPoint(point Position) float32 {
	return 0.0
}

// Get the player driving the vehicle.
func (v *Vehicle) Driver() (*Player, error) {
	return nil, nil
}

func (v *Vehicle) Health() float32 {
	return 0.0
}

// Set a vehicle's health. When a vehicle's health decreases the engine will produce smoke, and finally fire when it decreases to less than 250 (25%).
func (v *Vehicle) SetHealth(health float32) {

}

func (v *Vehicle) HydraReactorAngle() float32 {
	return 0.0
}

func (v *Vehicle) InteriorID() int {
	return 0
}

func (v *Vehicle) SetInteriorID(interiorID int) {

}

// Gets the current vehicle landing gear state from the latest driver.
func (v *Vehicle) LandingGearState() int {
	return 0
}

// Get the last driver of a vehicle.
func (v *Vehicle) LastDriver() (*Player, error) {
	return nil, nil
}

func (v *Vehicle) Matrix() {

}

// Gets the model ID of a vehicle.
func (v *Vehicle) ModelID() int {
	// TODO: modelID constants
	return 0
}

func (v *Vehicle) NumberPlate() string {
	return ""
}

func (v *Vehicle) SetNumberPlate(numberPlate string) {

}

// Get the occupied tick of a vehicle.
func (v *Vehicle) OccupiedTick() int {
	return 0
}

// Set the occupied tick of a vehicle.
func (v *Vehicle) SetOccupiedTick(occupiedTick int) {

}

// Allows you to retrieve the current state of a vehicle's doors.
func (v *Vehicle) DoorsState() *VehicleDoorsState {
	return nil
}

// Allows you to open and close the doors of a vehicle.
func (v *Vehicle) SetDoorsState(state *VehicleDoorsState) {

}

// Allows you to retrieve the current state of a vehicle's windows.
func (v *Vehicle) WindowsState() *VehicleWindowsState {
	return nil
}

// Allows you to open and close the windows of a vehicle.
func (v *Vehicle) SetWindowsState(state *VehicleWindowsState) {

}

// Gets a vehicle's parameters.
func (v *Vehicle) Params() *VehicleParams {
	return nil
}

// Sets a vehicle's parameters for all players.
func (v *Vehicle) SetParams(params *VehicleParams) {

}

// Set the parameters of a vehicle for a player.
func (v *Vehicle) SetParamsForPlayer(player *Player, objective, areDoorsLoked bool) {

}

// Returns a vehicle's siren state (on/off).
func (v *Vehicle) SirenState() bool {
	return false
}

// Turn the siren for a vehicle on or off.
func (v *Vehicle) SetSirenState(state int) {

}

func (v *Vehicle) Position() *Position {
	return nil
}

// Set a vehicle's position.
func (v *Vehicle) SetPosition(position *Position) {

}

// Get the respawn delay of a vehicle in seconds.
func (v *Vehicle) RespawnDelay() time.Duration {
	return 1 * time.Second
}

// Set the respawn delay of a vehicle.
func (v *Vehicle) SetRespawnDelay(respawnDelay time.Duration) {

}

// Returns respawn tick in milliseconds.
func (v *Vehicle) RespawnTick() time.Duration {
	return 1 * time.Millisecond
}

// Set the respawn tick of a vehicle in milliseconds.
func (v *Vehicle) SetRespawnTick(respawnTick time.Duration) {

}

// Returns a vehicle's rotation on all axes as a quaternion.
func (v *Vehicle) RotationQuat() {

}

// Gets the vehicle spawn location and colors.
func (v *Vehicle) SpawnInfo() *VehicleSpawnInfo {
	return nil
}

func (v *Vehicle) SetSpawnInfo(spawnInfo *VehicleSpawnInfo) {

}

func (v *Vehicle) TrainSpeed() int {
	return 0
}

// Get the velocity of a vehicle on the X, Y and Z axes.
func (v *Vehicle) Velocity() *Position {
	return nil
}

// Sets the X, Y and Z velocity of a vehicle.
func (v *Vehicle) SetVelocity(velocity *Position) {

}

// Get the virtual world of a vehicle.
func (v *Vehicle) VirtualWorld() int {
	return 0
}

// Sets the 'virtual world' of a vehicle. Players will only be able to see vehicles in their own virtual world.
func (v *Vehicle) SetVirtualWorld(virtualWorld int) {

}

// Get the rotation of a vehicle on the Z axis (yaw).
func (v *Vehicle) ZAngle() float32 {
	return 0.0
}

// Set the Z rotation (yaw) of a vehicle.
func (v *Vehicle) SetZAngle(zAngle float32) {

}

// Check if a vehicle is occupied.
func (v *Vehicle) IsOccupied() bool {
	return false
}

// Sets the vehicle's occupancy.
func (v *Vehicle) Occupy(occupied bool) {

}

// Hides a vehicle from the game.
func (v *Vehicle) Hide() {

}

// Shows the hidden vehicle.
func (v *Vehicle) Show() {

}

// Checks if a vehicle is hidden.
func (v *Vehicle) IsHidden() bool {
	return false
}

// Checks if a vehicle has a trailer attached to it.
func (v *Vehicle) HasTrailer() bool {
	return false
}

// Check if a vehicle is dead.
func (v *Vehicle) IsDead() bool {
	return false
}

// Sets the vehicle to dead.
func (v *Vehicle) SetDead(dead bool) {

}

// Checks if a vehicle siren is on or off.
func (v *Vehicle) IsSirenEnabled() bool {
	return false
}

// Checks if a vehicle is streamed in for a player. Only nearby vehicles are streamed in (visible) for a player.
func (v *Vehicle) IsStreamedIn(forPlayer *Player) bool {
	return false
}

// Fully repairs a vehicle, including visual damage (bumps, dents, scratches, popped tires etc.).
func (v *Vehicle) Repair() {

}

// Sets the angular X, Y and Z velocity of a vehicle.
func (v *Vehicle) SetAngularVelocity(velocity *Position) {

}

// Sets a vehicle back to the position at where it was created.
func (v *Vehicle) SetToRespawn() {

}
