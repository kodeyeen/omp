package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
	"text/template"
)

func main() {
	api, err := parseAPIDocs("capi/apidocs/api.json")
	if err != nil {
		log.Fatalf("Failed to parse API docs file: %s\n", err.Error())
	}

	err = genFromTemplate("include/cwrappers.h", "cwrappers.h.go.tmpl", api)
	if err != nil {
		log.Fatalf("Failed to generate header file: %s\n", err.Error())
	}

	err = genFromTemplate("cwrappers.c", "cwrappers.c.go.tmpl", api)
	if err != nil {
		log.Fatalf("Failed to generate source file: %s\n", err.Error())
	}
}

func parseAPIDocs(path string) (*API, error) {
	inf, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("failed to open API docs file: %w", err)
	}

	var api *API

	err = json.NewDecoder(inf).Decode(&api)
	if err != nil {
		return nil, fmt.Errorf("failed to decode API docs file: %w", err)
	}

	return api, nil
}

func genFromTemplate(dstPath, tmplPath string, api *API) error {
	tmpl, err := os.ReadFile(tmplPath)
	if err != nil {
		return fmt.Errorf("failed to read template file: %w", err)
	}

	funcs := template.FuncMap{
		"split": strings.Split,
		"join":  strings.Join,
	}

	t := template.Must(template.New("func").Funcs(funcs).Parse(string(tmpl)))

	dstFile, err := os.Create(dstPath)
	if err != nil {
		return fmt.Errorf("failed to create destination file: %w", err)
	}

	err = t.Execute(dstFile, []*Component{
		{Name: "Actor", Funcs: api.Actor},
		{Name: "Checkpoint", Funcs: api.Checkpoint},
		{Name: "RaceCheckpoint", Funcs: api.RaceCheckpoint},
		{Name: "Class", Funcs: api.Class},
		{Name: "Player", Funcs: api.Player},
		{Name: "Config", Funcs: api.Config},
		{Name: "Core", Funcs: api.Core},
		{Name: "NPC", Funcs: api.NPC},
		{Name: "Dialog", Funcs: api.Dialog},
		{Name: "Event", Funcs: api.Event},
		{Name: "GangZone", Funcs: api.GangZone},
		{Name: "Menu", Funcs: api.Menu},
		{Name: "Object", Funcs: api.Object},
		{Name: "PlayerObject", Funcs: api.PlayerObject},
		{Name: "Pickup", Funcs: api.Pickup},
		{Name: "All", Funcs: api.All},
		{Name: "Recording", Funcs: api.Recording},
		{Name: "TextDraw", Funcs: api.TextDraw},
		{Name: "PlayerTextDraw", Funcs: api.PlayerTextDraw},
		{Name: "TextLabel", Funcs: api.TextLabel},
		{Name: "PlayerTextLabel", Funcs: api.PlayerTextLabel},
		{Name: "Vehicle", Funcs: api.Vehicle},
	})
	if err != nil {
		return fmt.Errorf("failed to execute template file: %w", err)
	}

	return nil
}
