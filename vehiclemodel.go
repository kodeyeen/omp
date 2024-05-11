package omp

// #include "include/omp.h"
import "C"

const (
	VehicleModelLandstalker VehicleModel = iota + 400
	VehicleModelBravura
	VehicleModelBuffalo
	VehicleModelLinerunner
	VehicleModelPerrenial
	VehicleModelSentinel
	VehicleModelDumper
	VehicleModelFiretruck
	VehicleModelTrashmaster
	VehicleModelStretch
	VehicleModelManana
	VehicleModelInfernus
	VehicleModelVoodoo
	VehicleModelPony
	VehicleModelMule
	VehicleModelCheetah
	VehicleModelAmbulance
	VehicleModelLeviathan
	VehicleModelMoonbeam
	VehicleModelEsperanto
	VehicleModelTaxi
	VehicleModelWashington
	VehicleModelBobcat
	VehicleModelMrWhoopee
	VehicleModelBFInjection
	VehicleModelHunter
	VehicleModelPremier
	VehicleModelEnforcer
	VehicleModelSecuricar
	VehicleModelBanshee
	VehicleModelPredator
	VehicleModelBus
	VehicleModelRhino
	VehicleModelBarracks
	VehicleModelHotknife
	VehicleModelArticleTrailer1
	VehicleModelPrevion
	VehicleModelCoach
	VehicleModelCabbie
	VehicleModelStallion
	VehicleModelRumpo
	VehicleModelRCBandit
	VehicleModelRomero
	VehicleModelPacker
	VehicleModelMonster
	VehicleModelAdmiral
	VehicleModelSqualo
	VehicleModelSeasparrow
	VehicleModelPizzaBoy
	VehicleModelTram
	VehicleModelArticleTrailer2
	VehicleModelTurismo
	VehicleModelSpeeder
	VehicleModelReefer
	VehicleModelTropic
	VehicleModelFlatbed
	VehicleModelYankee
	VehicleModelCaddy
	VehicleModelSolair
	VehicleModelRCVan
	VehicleModelSkimmer
	VehicleModelPCJ600
	VehicleModelFaggio
	VehicleModelFreeway
	VehicleModelRCBaron
	VehicleModelRCRaider
	VehicleModelGlendale
	VehicleModelOceanic
	VehicleModelSanchez
	VehicleModelSparrow
	VehicleModelPatriot
	VehicleModelQuad
	VehicleModelCoastguard
	VehicleModelDinghy
	VehicleModelHermes
	VehicleModelSabre
	VehicleModelRustler
	VehicleModelZR350
	VehicleModelWalton
	VehicleModelRegina
	VehicleModelComet
	VehicleModelBMX
	VehicleModelBurrito
	VehicleModelCamper
	VehicleModelMarquis
	VehicleModelBaggage
	VehicleModelDozer
	VehicleModelMaverick
	VehicleModelNewsChopper
	VehicleModelRancher
	VehicleModelFBIRancher
	VehicleModelVirgo
	VehicleModelGreenwood
	VehicleModelJetmax
	VehicleModelHotringC
	VehicleModelSandking
	VehicleModelBlistaCompact
	VehicleModelPoliceMaverick
	VehicleModelBoxville
	VehicleModelBenson
	VehicleModelMesa
	VehicleModelRCGoblin
	VehicleModelHotringA
	VehicleModelHotringB
	VehicleModelBloodringBanger
	VehicleModelRancherLure
	VehicleModelSuperGT
	VehicleModelElegant
	VehicleModelJourney
	VehicleModelBike
	VehicleModelMountainBike
	VehicleModelBeagle
	VehicleModelCropdust
	VehicleModelStuntplane
	VehicleModelTanker
	VehicleModelRoadTrain
	VehicleModelNebula
	VehicleModelMajestic
	VehicleModelBuccaneer
	VehicleModelShamal
	VehicleModelHydra
	VehicleModelFCR900
	VehicleModelNRG500
	VehicleModelHPV1000
	VehicleModelCementTruck
	VehicleModelTowtruck
	VehicleModelFortune
	VehicleModelCadrona
	VehicleModelFBITruck
	VehicleModelWillard
	VehicleModelForklift
	VehicleModelTractor
	VehicleModelCombineHarvester
	VehicleModelFeltzer
	VehicleModelRemington
	VehicleModelSlamvan
	VehicleModelBlade
	VehicleModelFreight
	VehicleModelBrownstreak
	VehicleModelVortex
	VehicleModelVincent
	VehicleModelBullet
	VehicleModelClover
	VehicleModelSadler
	VehicleModelFiretruckLA
	VehicleModelHustler
	VehicleModelIntruder
	VehicleModelPrimo
	VehicleModelCargobob
	VehicleModelTampa
	VehicleModelSunrise
	VehicleModelMerit
	VehicleModelUtilityVan
	VehicleModelNevada
	VehicleModelYosemite
	VehicleModelWindsor
	VehicleModelMonsterA
	VehicleModelMonsterB
	VehicleModelUranus
	VehicleModelJester
	VehicleModelSultan
	VehicleModelStratum
	VehicleModelElegy
	VehicleModelRaindance
	VehicleModelRCTiger
	VehicleModelFlash
	VehicleModelTahoma
	VehicleModelSavanna
	VehicleModelBandito
	VehicleModelFreightTrailer
	VehicleModelStreakTrailer
	VehicleModelKart
	VehicleModelMower
	VehicleModelDuneride
	VehicleModelSweeper
	VehicleModelBroadway
	VehicleModelTornado
	VehicleModelAT400
	VehicleModelDFT30
	VehicleModelHuntley
	VehicleModelStafford
	VehicleModelBF400
	VehicleModelNewsvan
	VehicleModelTug
	VehicleModelPetrolTrailer
	VehicleModelEmperor
	VehicleModelWayfarer
	VehicleModelEuros
	VehicleModelHotdog
	VehicleModelClub
	VehicleModelFreightBoxTrailer
	VehicleModelArticleTrailer3
	VehicleModelAndromada
	VehicleModelDodo
	VehicleModelRCCam
	VehicleModelLaunch
	VehicleModelLSPDCar
	VehicleModelSFPDCar
	VehicleModelLVPDCar
	VehicleModelRanger
	VehicleModelPicador
	VehicleModelSWAT
	VehicleModelAlpha
	VehicleModelPhoenix
	VehicleModelGlendaleDamaged
	VehicleModelSadlerDamaged
	VehicleModelBaggageTrailerA
	VehicleModelBaggageTrailerB
	VehicleModelStairsTrailer
	VehicleModelBoxburg
	VehicleModelFarmTrailer
	VehicleModelUtilityTrailer
)

type VehicleModel int

// Get the number of used vehicle models on the server.
func UsedVehicleModelCount() int {
	panic("not implemented")
}

func VehicleCountByModel(model VehicleModel) int {
	panic("not implemented")
}

func (m VehicleModel) Size() Vector3 {
	var cSize C.Vector3
	C.getVehicleModelInfo(C.int(m), 1, &cSize)

	return Vector3{
		X: float32(cSize.x),
		Y: float32(cSize.y),
		Z: float32(cSize.z),
	}
}

func (m VehicleModel) FrontSeatPosition() Vector3 {
	var cPos C.Vector3
	C.getVehicleModelInfo(C.int(m), 2, &cPos)

	return Vector3{
		X: float32(cPos.x),
		Y: float32(cPos.y),
		Z: float32(cPos.z),
	}
}

func (m VehicleModel) RearSeatPosition() Vector3 {
	var cPos C.Vector3
	C.getVehicleModelInfo(C.int(m), 3, &cPos)

	return Vector3{
		X: float32(cPos.x),
		Y: float32(cPos.y),
		Z: float32(cPos.z),
	}
}

func (m VehicleModel) PetrolCapPosition() Vector3 {
	var cPos C.Vector3
	C.getVehicleModelInfo(C.int(m), 4, &cPos)

	return Vector3{
		X: float32(cPos.x),
		Y: float32(cPos.y),
		Z: float32(cPos.z),
	}
}

func (m VehicleModel) FrontWheelsPosition() Vector3 {
	var cPos C.Vector3
	C.getVehicleModelInfo(C.int(m), 5, &cPos)

	return Vector3{
		X: float32(cPos.x),
		Y: float32(cPos.y),
		Z: float32(cPos.z),
	}
}

func (m VehicleModel) RearWheelsPosition() Vector3 {
	var cPos C.Vector3
	C.getVehicleModelInfo(C.int(m), 6, &cPos)

	return Vector3{
		X: float32(cPos.x),
		Y: float32(cPos.y),
		Z: float32(cPos.z),
	}
}

func (m VehicleModel) MiddleWheelsPosition() Vector3 {
	var cPos C.Vector3
	C.getVehicleModelInfo(C.int(m), 7, &cPos)

	return Vector3{
		X: float32(cPos.x),
		Y: float32(cPos.y),
		Z: float32(cPos.z),
	}
}

func (m VehicleModel) FrontBumperZ() float32 {
	var cPos C.Vector3
	C.getVehicleModelInfo(C.int(m), 8, &cPos)

	return float32(cPos.z)
}

func (m VehicleModel) RearBumperHeight() float32 {
	var cPos C.Vector3
	C.getVehicleModelInfo(C.int(m), 9, &cPos)

	return float32(cPos.z)
}
