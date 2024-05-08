package main

import (
	"time"

	"github.com/kodeyeen/omp"
)

type Character struct {
	*omp.Player
	AreBuildingRemoved bool
	SpectateState      SpectateState
	LastDiedAt         time.Time
	LastResuppliedAt   time.Time
	LastKiller         *Character
}

func (c *Character) RemoveNeededBuildings() {
	if c.AreBuildingRemoved {
		return
	}

	c.RemoveDefaultObjects(9090, 10.0, omp.Vector3{X: 2317.0859, Y: 572.2656, Z: -20.9688})
	c.RemoveDefaultObjects(9091, 10.0, omp.Vector3{X: 2317.0859, Y: 572.2656, Z: -20.9688})
	c.RemoveDefaultObjects(13483, 0.25, omp.Vector3{X: 2113.5781, Y: -96.7344, Z: 0.9844})
	c.RemoveDefaultObjects(12990, 0.25, omp.Vector3{X: 2113.5781, Y: -96.7344, Z: 0.9844})
	c.RemoveDefaultObjects(935, 0.25, omp.Vector3{X: 2119.8203, Y: -84.4063, Z: -0.0703})
	c.RemoveDefaultObjects(1369, 0.25, omp.Vector3{X: 2104.0156, Y: -105.2656, Z: 1.7031})
	c.RemoveDefaultObjects(935, 0.25, omp.Vector3{X: 2122.3750, Y: -83.3828, Z: 0.4609})
	c.RemoveDefaultObjects(935, 0.25, omp.Vector3{X: 2119.5313, Y: -82.8906, Z: -0.1641})
	c.RemoveDefaultObjects(935, 0.25, omp.Vector3{X: 2120.5156, Y: -79.0859, Z: 0.2188})
	c.RemoveDefaultObjects(935, 0.25, omp.Vector3{X: 2119.4688, Y: -69.7344, Z: 0.2266})
	c.RemoveDefaultObjects(935, 0.25, omp.Vector3{X: 2119.4922, Y: -73.6172, Z: 0.1250})
	c.RemoveDefaultObjects(935, 0.25, omp.Vector3{X: 2117.8438, Y: -67.8359, Z: 0.1328})

	c.AreBuildingRemoved = true
}

func (c *Character) SetupForClassSelection() {
	// Set the player's orientation when they're selecting a class.
	c.SetPosition(omp.Vector3{X: 1984.4445, Y: 157.9501, Z: 55.9384})
	c.SetCameraPosition(omp.Vector3{X: 1984.4445, Y: 160.9501, Z: 55.9384})
	c.SetCameraLookAt(omp.Vector3{X: 1984.4445, Y: 157.9501, Z: 55.9384}, omp.PlayerCameraCutTypeCut)
	c.SetFacingAngle(0.0)
}

func (c *Character) SetTeamFromClass(cls *omp.Class) {
	// Set their team number based on the class they selected.
	clsID := cls.ID()

	if clsID == 0 || clsID == 1 {
		c.SetTeam(TeamGreen)
	} else if clsID == 2 || clsID == 3 {
		c.SetTeam(TeamBlue)
	}
}

func (c *Character) HandleSpectating() {
	if c.LastKiller != nil && (c.LastKiller.State() == omp.PlayerStateOnFoot ||
		c.LastKiller.State() == omp.PlayerStateDriver ||
		c.LastKiller.State() == omp.PlayerStatePassenger) {

		c.SpectateCharacter(c.LastKiller)
		c.SpectateState = SpectateStatePlayer
	} else {
		if c.SpectateState != SpectateStateFixed {
			c.SpectateFixedPosition()
			c.SpectateState = SpectateStateFixed
		}
	}
}

func (c *Character) SpectateCharacter(target *Character) {
	if target.State() == omp.PlayerStateOnFoot {
		_, err := c.SpectatingPlayer()
		if err != nil {
			c.SpectatePlayer(target.Player, omp.PlayerSpectateModeNormal)
		}
	} else if target.State() == omp.PlayerStateDriver || target.State() == omp.PlayerStatePassenger {
		_, err := c.SpectatingVehicle()
		if err != nil {
			targetVeh, err := target.Vehicle()
			if err == nil {
				c.SpectateVehicle(targetVeh, omp.PlayerSpectateModeNormal)
			}
		}
	}
}

func (c *Character) SpectateFixedPosition() {
	if c.Team() == TeamGreen {
		c.SetCameraPosition(omp.Vector3{X: 2221.5820, Y: -273.9985, Z: 61.7806})
		c.SetCameraLookAt(omp.Vector3{X: 2220.9978, Y: -273.1861, Z: 61.4606}, omp.PlayerCameraCutTypeCut)
	} else {
		c.SetCameraPosition(omp.Vector3{X: 2274.8467, Y: 591.3257, Z: 30.1311})
		c.SetCameraLookAt(omp.Vector3{X: 2275.0503, Y: 590.3463, Z: 29.9460}, omp.PlayerCameraCutTypeCut)
	}
}

func (c *Character) DoResupply() {
	if c.LastResuppliedAt.IsZero() || time.Since(c.LastResuppliedAt) > 30*time.Second {
		c.LastResuppliedAt = time.Now()
		c.ResetWeapons()
		c.GiveWeapon(omp.WeaponM4, 100)
		c.GiveWeapon(omp.WeaponMP5, 200)
		c.GiveWeapon(omp.WeaponSniper, 10)
		c.SetHealth(100.0)
		c.SetArmor(100.0)
		c.ShowGameText("Resupplied", 2*time.Second, 5)
		c.PlaySound(1150, omp.Vector3{})
	}
}

func (c *Character) SetColorFromTeam() {
	if c.Team() == TeamGreen {
		c.SetColor(ColorGreen)
	} else if c.Team() == TeamBlue {
		c.SetColor(ColorBlue)
	}
}
