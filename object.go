package gomp

// #include "include/object.h"
// #include "include/player.h"
// #include "include/vehicle.h"
import "C"
import (
	"errors"
	"time"
	"unsafe"
)

type ObjectMoveData struct {
	TargetPos Vector3
	TargetRot Vector3
	Speed     float32
}

type ObjectAttachmentTarget interface {
	Vehicle | Object | Player
}

type ObjectAttachmentData[T ObjectAttachmentTarget] struct {
	SyncRotation bool
	Target       T
	Offset       Vector3
	Rotation     Vector3
}

type ObjectMaterialSize int

const (
	ObjectMaterialSize32x32   = 10
	ObjectMaterialSize64x32   = 20
	ObjectMaterialSize64x64   = 30
	ObjectMaterialSize128x32  = 40
	ObjectMaterialSize128x64  = 50
	ObjectMaterialSize128x128 = 60
	ObjectMaterialSize256x32  = 70
	ObjectMaterialSize256x64  = 80
	ObjectMaterialSize256x128 = 90
	ObjectMaterialSize256x256 = 100
	ObjectMaterialSize512x64  = 110
	ObjectMaterialSize512x128 = 120
	ObjectMaterialSize512x256 = 130
	ObjectMaterialSize512x512 = 140
)

type ObjectMaterialTextAlign int

const (
	ObjectMaterialTextAlignLeft = iota
	ObjectMaterialTextAlignCenter
	ObjectMaterialTextAlignRight
)

type ObjectMaterial struct {
	ModelID     int
	TextureLib  string
	TextureName string
	Color       Color
}

type ObjectMaterialText struct {
	Text            string
	MaterialSize    int
	FontFace        string
	FontSize        int
	IsBold          bool
	FontColor       Color
	BackgroundColor Color
	Alignment       int
}

type Object struct {
	handle unsafe.Pointer
}

func NewObject(modelID int, pos Vector3, rot Vector3, drawDist float32) (*Object, error) {
	cObject := C.object_create(C.int(modelID), C.float(pos.X), C.float(pos.Y), C.float(pos.Z), C.float(rot.X), C.float(rot.Y), C.float(rot.Z), C.float(drawDist))
	if cObject == nil {
		return nil, errors.New("object limit was reached")
	}

	return &Object{handle: cObject}, nil
}

func FreeObject(obj *Object) {
	C.object_release(obj.handle)
}

func EnableObjectsCameraCollision() {
	C.object_setDefaultCameraCollision(C.uchar(1))
}

func DisableObjectsCameraCollision() {
	C.object_setDefaultCameraCollision(C.uchar(0))
}

func (o *Object) SetDrawDistance(drawDist float32) {
	C.object_setDrawDistance(o.handle, C.float(drawDist))
}

func (o *Object) DrawDistance() float32 {
	return float32(C.object_getDrawDistance(o.handle))
}

func (o *Object) SetModel(model int) {
	C.object_setModel(o.handle, C.int(model))
}

func (o *Object) Model() int {
	return int(C.object_getModel(o.handle))
}

func (o *Object) EnableCameraCollision() {
	C.object_setCameraCollision(o.handle, C.uchar(1))
}

func (o *Object) DisableCameraCollision() {
	C.object_setCameraCollision(o.handle, C.uchar(0))
}

func (o *Object) IsCameraCollisionEnabled() bool {
	return C.object_getCameraCollision(o.handle) != 0
}

func (o *Object) Move(pos Vector3, rot Vector3, speed float32) time.Duration {
	millis := C.object_move(o.handle, C.float(pos.X), C.float(pos.Y), C.float(pos.Z), C.float(rot.X), C.float(rot.Y), C.float(rot.Z), C.float(speed))

	return time.Duration(millis) * time.Millisecond
}

func (o *Object) IsMoving() bool {
	return C.object_isMoving(o.handle) != 0
}

func (o *Object) Stop() {
	C.object_stop(o.handle)
}

func (o *Object) MovingData() ObjectMoveData {
	data := C.object_getMovingData(o.handle)

	return ObjectMoveData{
		TargetPos: Vector3{
			X: float32(data.targetPos.x),
			Y: float32(data.targetPos.y),
			Z: float32(data.targetPos.z),
		},
		TargetRot: Vector3{
			X: float32(data.targetPos.x),
			Y: float32(data.targetPos.y),
			Z: float32(data.targetPos.z),
		},
		Speed: float32(data.speed),
	}
}

func (o *Object) AttachToVehicle(veh *Vehicle, offset Vector3, rot Vector3) {
	C.object_attachToVehicle(o.handle, veh.handle, C.float(offset.X), C.float(offset.Y), C.float(offset.Z), C.float(rot.X), C.float(rot.Y), C.float(rot.Z))
}

func (o *Object) ResetAttachment() {
	C.object_resetAttachment(o.handle)
}

func ObjectAttachedData[T ObjectAttachmentTarget](obj *Object) (ObjectAttachmentData[T], error) {
	data := C.object_getAttachmentData(obj.handle)

	t := any(new(T))

	var result ObjectAttachmentData[T]
	var target any

	switch t.(type) {
	case *Vehicle:
		veh := C.vehicle_getByID(data.ID)
		if veh == nil {
			return result, errors.New("object is not attached to a vehicle")
		}

		target = Vehicle{handle: veh}
	case *Object:
		obj := C.object_getByID(data.ID)
		if obj == nil {
			return result, errors.New("object is not attached to an object")
		}

		target = Object{handle: obj}
	case *Player:
		plr := C.player_getByID(data.ID)
		if plr == nil {
			return result, errors.New("object is not attached to a player")
		}

		target = Player{handle: plr}
	}

	result = ObjectAttachmentData[T]{
		SyncRotation: data.syncRotation != 0,
		Target:       target.(T),
		Offset: Vector3{
			X: float32(data.offset.x),
			Y: float32(data.offset.y),
			Z: float32(data.offset.z),
		},
		Rotation: Vector3{
			X: float32(data.rotation.x),
			Y: float32(data.rotation.y),
			Z: float32(data.rotation.z),
		},
	}

	return result, nil
}

func (o *Object) IsMaterialSlotUsed(slotIdx int) bool {
	return C.object_isMaterialSlotUsed(o.handle, C.uint(slotIdx)) != 0
}

func (o *Object) Material(slotIdx int) (ObjectMaterial, error) {
	var data C.ObjectMaterial
	res := C.object_getMaterial(o.handle, C.uint(slotIdx), &data) != 0
	if !res {
		return ObjectMaterial{}, errors.New("invalid slot index is specified")
	}

	material := ObjectMaterial{
		ModelID:     int(data.model),
		TextureLib:  C.GoStringN(data.textureLibrary.buf, C.int(data.textureLibrary.length)),
		TextureName: C.GoStringN(data.textureName.buf, C.int(data.textureName.length)),
		Color:       Color(data.colour),
	}

	return material, nil
}

func (o *Object) MaterialText(slotIdx int) (ObjectMaterialText, error) {
	var data C.ObjectMaterialText
	res := C.object_getMaterialText(o.handle, C.uint(slotIdx), &data) != 0
	if !res {
		return ObjectMaterialText{}, errors.New("invalid slot index is specified")
	}

	materialText := ObjectMaterialText{
		Text:            C.GoStringN(data.text.buf, C.int(data.text.length)),
		MaterialSize:    int(data.materialSize),
		FontFace:        C.GoStringN(data.fontFace.buf, C.int(data.fontFace.length)),
		FontSize:        int(data.fontSize),
		IsBold:          data.bold != 0,
		FontColor:       Color(data.fontColour),
		BackgroundColor: Color(data.backgroundColour),
		Alignment:       int(data.alignment),
	}

	return materialText, nil
}

func (o *Object) SetMaterial(idx, model int, textureLib, textureName string, color int) {
	cTextureLib := newCString(textureLib)
	defer freeCString(cTextureLib)

	cTextureName := newCString(textureName)
	defer freeCString(cTextureName)

	C.object_setMaterial(o.handle, C.uint(idx), C.int(model), cTextureLib, cTextureName, C.uint(color))
}

func (o *Object) SetMaterialText(
	slotIdx int,
	text string,
	size ObjectMaterialSize,
	fontFace string,
	fontSize int,
	bold bool,
	fontColor, bgColor Color,
	align ObjectMaterialTextAlign,
) {
	cText := newCString(text)
	defer freeCString(cText)

	cFontFace := newCString(fontFace)
	defer freeCString(cFontFace)

	C.object_setMaterialText(o.handle, C.uint(slotIdx), cText, C.ObjectMaterialSize(size), cFontFace, C.int(fontSize), newCUchar(bold), C.uint(fontColor), C.uint(bgColor), C.ObjectMaterialTextAlign(align))
}

func (o *Object) AttachToPlayer(plr *Player, offset Vector3, rot Vector3) {
	C.object_attachToPlayer(o.handle, plr.handle, C.float(offset.X), C.float(offset.Y), C.float(offset.Z), C.float(rot.X), C.float(rot.Y), C.float(rot.Z))
}

func (o *Object) SetPosition(pos Vector3) {
	C.object_setPosition(o.handle, C.float(pos.X), C.float(pos.Y), C.float(pos.Z))
}

func (o *Object) Position() Vector3 {
	pos := C.object_getPosition(o.handle)

	return Vector3{
		X: float32(pos.x),
		Y: float32(pos.y),
		Z: float32(pos.z),
	}
}

func (o *Object) SetRotation(rot Vector3) {
	C.object_setRotation(o.handle, C.float(rot.X), C.float(rot.Y), C.float(rot.Z))
}

func (o *Object) Rotation() Vector3 {
	rot := C.object_getRotation(o.handle)

	return Vector3{
		X: float32(rot.x),
		Y: float32(rot.y),
		Z: float32(rot.z),
	}
}
