package omp

// #include "include/player.h"
import "C"
import (
	"strings"
)

type dialogStyle int

const (
	dialogStyleMsgBox dialogStyle = iota
	dialogStyleInput
	dialogStyleList
	dialogStylePassword
	dialogStyleTablist
	dialogStyleTablistHeaders
)

var activeDialogs = make(map[int]dialog, 1000)

type dialog interface {
	ShowFor(player *Player)
	HideFor(player *Player)
}

type MessageDialog struct {
	Events                        *Dispatcher
	Title, Body, Button1, Button2 string
}

func NewMessageDialog(title, body, button1, button2 string) *MessageDialog {
	return &MessageDialog{
		Events:  NewDispatcher(),
		Title:   title,
		Body:    body,
		Button1: button1,
		Button2: button2,
	}
}

func (d *MessageDialog) ShowFor(player *Player) {
	activeDialogs[player.ID()] = d

	player.showDialog(dialogStyleMsgBox, d.Title, d.Body, d.Button1, d.Button2)
}

func (d *MessageDialog) HideFor(player *Player) {
	delete(activeDialogs, player.ID())

	player.hideDialog()
}

type InputDialog struct {
	Events                        *Dispatcher
	Title, Body, Button1, Button2 string
	isPassword                    bool
}

func NewInputDialog(title, body, button1, button2 string) *InputDialog {
	return &InputDialog{
		Events:     NewDispatcher(),
		Title:      title,
		Body:       body,
		Button1:    button1,
		Button2:    button2,
		isPassword: false,
	}
}

func NewPasswordDialog(title, body, button1, button2 string) *InputDialog {
	return &InputDialog{
		Events:     NewDispatcher(),
		Title:      title,
		Body:       body,
		Button1:    button1,
		Button2:    button2,
		isPassword: true,
	}
}

func (d *InputDialog) ShowFor(player *Player) {
	style := dialogStyleInput
	if d.isPassword {
		style = dialogStylePassword
	}

	activeDialogs[player.ID()] = d

	player.showDialog(style, d.Title, d.Body, d.Button1, d.Button2)
}

func (d *InputDialog) HideFor(player *Player) {
	delete(activeDialogs, player.ID())

	player.hideDialog()
}

type ListDialog struct {
	Events                  *Dispatcher
	Title, Button1, Button2 string
	items                   []string
}

func NewListDialog(title, button1, button2 string) *ListDialog {
	return &ListDialog{
		Events:  NewDispatcher(),
		Title:   title,
		Button1: button1,
		Button2: button2,
	}
}

func (d *ListDialog) SetItems(items []string) {
	d.items = items
}

func (d *ListDialog) Add(items ...string) {
	d.items = append(d.items, items...)
}

func (d *ListDialog) ShowFor(player *Player) {
	body := strings.Join(d.items, "\n")

	activeDialogs[player.ID()] = d

	player.showDialog(dialogStyleList, d.Title, body, d.Button1, d.Button2)
}

func (d *ListDialog) HideFor(player *Player) {
	delete(activeDialogs, player.ID())

	player.hideDialog()
}

type TabListItem []string

type TabListDialog struct {
	Events                  *Dispatcher
	Title, Button1, Button2 string
	header                  TabListItem
	items                   []TabListItem
}

func NewTabListDialog(title, button1, button2 string) *TabListDialog {
	return &TabListDialog{
		Events:  NewDispatcher(),
		Title:   title,
		Button1: button1,
		Button2: button2,
	}
}

func (d *TabListDialog) SetHeader(header TabListItem) {
	d.header = header
}

func (d *TabListDialog) SetItems(items []TabListItem) {
	d.items = items
}

func (d *TabListDialog) Add(items ...TabListItem) {
	d.items = append(d.items, items...)
}

func (d *TabListDialog) ShowFor(player *Player) {
	style := dialogStyleTablist
	if d.header != nil {
		style = dialogStyleTablistHeaders
	}

	body := d.makeBody(style)

	activeDialogs[player.ID()] = d

	player.showDialog(style, d.Title, body, d.Button1, d.Button2)
}

func (d *TabListDialog) makeBody(style dialogStyle) string {
	var rawItems []string

	if style == dialogStyleTablistHeaders {
		rawItems = append(rawItems, strings.Join(d.header, "\t"))
	}

	for _, item := range d.items {
		rawItems = append(rawItems, strings.Join(item, "\t"))
	}

	return strings.Join(rawItems, "\n")
}

func (d *TabListDialog) HideFor(player *Player) {
	delete(activeDialogs, player.ID())

	player.hideDialog()
}
