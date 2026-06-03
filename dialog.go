package omp

// #include "include/player.h"
import "C"
import (
	"context"
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
	Title, Body, Button1, Button2 string
	onResponse                    func(ctx context.Context, e *MessageDialogResponseEvent) error
}

func NewMessageDialog(title, body, button1, button2 string) *MessageDialog {
	return &MessageDialog{
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

func (d *MessageDialog) OnResponse(f func(ctx context.Context, e *MessageDialogResponseEvent) error) {
	d.onResponse = f
}

type InputDialog struct {
	Title, Body, Button1, Button2 string
	isPassword                    bool
	onResponse                    func(ctx context.Context, e *InputDialogResponseEvent) error
}

func NewInputDialog(title, body, button1, button2 string) *InputDialog {
	return &InputDialog{
		Title:      title,
		Body:       body,
		Button1:    button1,
		Button2:    button2,
		isPassword: false,
	}
}

func NewPasswordDialog(title, body, button1, button2 string) *InputDialog {
	return &InputDialog{
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

func (d *InputDialog) OnResponse(f func(ctx context.Context, e *InputDialogResponseEvent) error) {
	d.onResponse = f
}

type ListDialog struct {
	Title, Button1, Button2 string
	items                   []string
	onResponse              func(ctx context.Context, e *ListDialogResponseEvent) error
}

func NewListDialog(title, button1, button2 string) *ListDialog {
	return &ListDialog{
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

func (d *ListDialog) OnResponse(f func(ctx context.Context, e *ListDialogResponseEvent) error) {
	d.onResponse = f
}

type TabListItem []string

type TabListDialog struct {
	Title, Button1, Button2 string
	header                  TabListItem
	items                   []TabListItem
	onResponse              func(ctx context.Context, e *TabListDialogResponseEvent) error
}

func NewTabListDialog(title, button1, button2 string) *TabListDialog {
	return &TabListDialog{
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

func (d *TabListDialog) OnResponse(f func(ctx context.Context, e *TabListDialogResponseEvent) error) {
	d.onResponse = f
}
