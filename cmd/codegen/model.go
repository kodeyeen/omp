package main

type API struct {
	Actor           []*Func `json:"Actor"`
	Checkpoint      []*Func `json:"Checkpoint"`
	RaceCheckpoint  []*Func `json:"RaceCheckpoint"`
	Class           []*Func `json:"Class"`
	Player          []*Func `json:"Player"`
	Component       []*Func `json:"Component"`
	Config          []*Func `json:"Config"`
	Core            []*Func `json:"Core"`
	NPC             []*Func `json:"NPC"`
	CustomModel     []*Func `json:"CustomModel"`
	Dialog          []*Func `json:"Dialog"`
	Event           []*Func `json:"Event"`
	GangZone        []*Func `json:"GangZone"`
	Menu            []*Func `json:"Menu"`
	Object          []*Func `json:"Object"`
	PlayerObject    []*Func `json:"PlayerObject"`
	Pickup          []*Func `json:"Pickup"`
	All             []*Func `json:"All"`
	Recording       []*Func `json:"Recording"`
	TextDraw        []*Func `json:"TextDraw"`
	PlayerTextDraw  []*Func `json:"PlayerTextDraw"`
	TextLabel       []*Func `json:"TextLabel"`
	PlayerTextLabel []*Func `json:"PlayerTextLabel"`
	Vehicle         []*Func `json:"Vehicle"`
}

type FuncParam struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

type Func struct {
	Ret    string       `json:"ret"`
	Name   string       `json:"name"`
	Params []*FuncParam `json:"params"`
}

type Group struct {
	Name  string
	Funcs []*Func
}
