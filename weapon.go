package gomp

const (
	WeaponFist Weapon = iota
	WeaponBrassKnuckle
	WeaponGolfClub
	WeaponNiteStick
	WeaponKnife
	WeaponBat
	WeaponShovel
	WeaponPoolStick
	WeaponKatana
	WeaponChainsaw
	WeaponDildo
	WeaponDildo2
	WeaponVibrator
	WeaponVibrator2
	WeaponFlower
	WeaponCane
	WeaponGrenade
	WeaponTeargas
	WeaponMoltov
)

const (
	WeaponColt45 Weapon = iota + 22
	WeaponSilenced
	WeaponDeagle
	WeaponShotgun
	WeaponSawedoff
	WeaponShotgspa
	WeaponUZI
	WeaponMP5
	WeaponAK47
	WeaponM4
	WeaponTEC9
	WeaponRifle
	WeaponSniper
	WeaponRocketLauncher
	WeaponHeatSeeker
	WeaponFlameThrower
	WeaponMinigun
	WeaponSatchel
	WeaponBomb
	WeaponSprayCan
	WeaponFireExtinguisher
	WeaponCamera
	WeaponNight_Vis_Goggles
	WeaponThermal_Goggles
	WeaponParachute
)

const (
	WeaponVehicle Weapon = iota + 49
	WeaponHeliblades
	WeaponExplosion
)

const (
	WeaponDrown Weapon = iota + 53
	WeaponCollision
	WeaponEnd
)

type Weapon int

func (w Weapon) SlotType() WeaponSlotType {
	panic("not implemented")
}

const (
	WeaponSlotTypeUnknown WeaponSlotType = iota - 1
	WeaponSlotTypeUnarmed
	WeaponSlotTypeMelee
	WeaponSlotTypePistol
	WeaponSlotTypeShotgun
	WeaponSlotTypeMachineGun
	WeaponSlotTypeAssaultRifle
	WeaponSlotTypeLongRifle
	WeaponSlotTypeArtillery
	WeaponSlotTypeExplosives
	WeaponSlotTypeEquipment
	WeaponSlotTypeGift
	WeaponSlotTypeGadget
	WeaponSlotTypeDetonator
)

type WeaponSlotType int

type WeaponSlot struct {
	Weapon Weapon
	Ammo   int
}
