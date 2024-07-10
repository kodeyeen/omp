package omp

// #include "include/wrappers.h"
import "C"
import (
	"strings"
	"unsafe"

	"github.com/kodeyeen/event"
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

const dialogID = 999

var activeDialogs = make(map[int]Dialog, 1000)

type Dialog interface {
	ShowFor(player *Player)
	HideFor(player *Player)
}

func showDialog(plr *Player, style dialogStyle, title, body, button1, button2 string) {
	cTitle := C.CString(title)
	defer C.free(unsafe.Pointer(cTitle))

	cBody := C.CString(body)
	defer C.free(unsafe.Pointer(cBody))

	cButton1 := C.CString(button1)
	defer C.free(unsafe.Pointer(cButton1))

	cButton2 := C.CString(button2)
	defer C.free(unsafe.Pointer(cButton2))

	C.Dialog_Show(plr.handle, dialogID, C.int(style), cTitle, cBody, cButton1, cButton2)
}

func hideDialog(plr *Player) {
	C.Dialog_Hide(plr.handle)
}

type MessageDialog struct {
	*event.Dispatcher
	Title, Body, Button1, Button2 string
}

func NewMessageDialog(title, body, button1, button2 string) *MessageDialog {
	return &MessageDialog{
		Dispatcher: event.NewDispatcher(),
		Title:      title,
		Body:       body,
		Button1:    button1,
		Button2:    button2,
	}
}

func (d *MessageDialog) ShowFor(player *Player) {
	activeDialogs[player.ID()] = d

	event.Dispatch(d.Dispatcher, EventTypeDialogShow, &DialogShowEvent{
		Player: player,
	})

	showDialog(player, dialogStyleMsgBox, d.Title, d.Body, d.Button1, d.Button2)
}

func (d *MessageDialog) HideFor(player *Player) {
	event.Dispatch(d.Dispatcher, EventTypeDialogHide, &DialogHideEvent{
		Player: player,
	})

	delete(activeDialogs, player.ID())

	hideDialog(player)
}

type InputDialog struct {
	*event.Dispatcher
	Title, Body, Button1, Button2 string
	isPassword                    bool
}

func NewInputDialog(title, body, button1, button2 string) *InputDialog {
	return &InputDialog{
		Dispatcher: event.NewDispatcher(),
		Title:      title,
		Body:       body,
		Button1:    button1,
		Button2:    button2,
		isPassword: false,
	}
}

func NewPasswordDialog(title, body, button1, button2 string) *InputDialog {
	return &InputDialog{
		Dispatcher: event.NewDispatcher(),
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

	event.Dispatch(d.Dispatcher, EventTypeDialogShow, &DialogShowEvent{
		Player: player,
	})

	showDialog(player, style, d.Title, d.Body, d.Button1, d.Button2)
}

func (d *InputDialog) HideFor(player *Player) {
	event.Dispatch(d.Dispatcher, EventTypeDialogHide, &DialogHideEvent{
		Player: player,
	})

	delete(activeDialogs, player.ID())

	hideDialog(player)
}

type ListDialog struct {
	*event.Dispatcher
	Title, Button1, Button2 string
	items                   []string
}

func NewListDialog(title, button1, button2 string) *ListDialog {
	return &ListDialog{
		Dispatcher: event.NewDispatcher(),
		Title:      title,
		Button1:    button1,
		Button2:    button2,
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

	event.Dispatch(d.Dispatcher, EventTypeDialogShow, &DialogShowEvent{
		Player: player,
	})

	showDialog(player, dialogStyleList, d.Title, body, d.Button1, d.Button2)
}

func (d *ListDialog) HideFor(player *Player) {
	event.Dispatch(d.Dispatcher, EventTypeDialogHide, &DialogHideEvent{
		Player: player,
	})

	delete(activeDialogs, player.ID())

	hideDialog(player)
}

type TabListItem []string

type TabListDialog struct {
	*event.Dispatcher
	Title, Button1, Button2 string
	header                  TabListItem
	items                   []TabListItem
}

func NewTabListDialog(title, button1, button2 string) *TabListDialog {
	return &TabListDialog{
		Dispatcher: event.NewDispatcher(),
		Title:      title,
		Button1:    button1,
		Button2:    button2,
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

	event.Dispatch(d.Dispatcher, EventTypeDialogShow, &DialogShowEvent{
		Player: player,
	})

	showDialog(player, style, d.Title, body, d.Button1, d.Button2)
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
	event.Dispatch(d.Dispatcher, EventTypeDialogHide, &DialogHideEvent{
		Player: player,
	})

	delete(activeDialogs, player.ID())

	hideDialog(player)
}
