package omp

// #include "include/menu.h"
import "C"
import (
	"errors"
	"unsafe"
)

type Menu struct {
	handle unsafe.Pointer
}

func NewMenu(title string, pos Vector2, columns int, col1Width, col2Width float32) (*Menu, error) {
	cTitle := newCString(title)
	freeCString(cTitle)

	cMenu := C.menu_create(cTitle, C.float(pos.X), C.float(pos.Y), C.uchar(columns), C.float(col1Width), C.float(col2Width))
	if cMenu == nil {
		return nil, errors.New("menu limit was reached")
	}

	return &Menu{handle: cMenu}, nil
}

func FreeMenu(menu *Menu) {
	C.menu_release(menu.handle)
}

func (m *Menu) SetColumnHeader(column int, header string) {
	cHeader := newCString(header)
	defer freeCString(cHeader)

	C.menu_setColumnHeader(m.handle, cHeader, C.uchar(column))
}

func (m *Menu) ColumnHeader(column int) string {
	header := C.menu_getColumnHeader(m.handle, C.uchar(column))

	return C.GoStringN(header.buf, C.int(header.length))
}

func (m *Menu) AddItem(column int, text string) {
	cText := newCString(text)
	defer freeCString(cText)

	C.menu_addCell(m.handle, cText, C.uchar(column))
}

func (m *Menu) DisableRow(row int) {
	C.menu_disableRow(m.handle, C.uchar(row))
}

func (m *Menu) IsRowEnabled(row int) bool {
	return C.menu_isRowEnabled(m.handle, C.uchar(row)) != 0
}

func (m *Menu) Disable() {
	C.menu_disable(m.handle)
}

func (m *Menu) IsEnabled() bool {
	return C.menu_isEnabled(m.handle) != 0
}

func (m *Menu) Position() Vector2 {
	pos := C.menu_getPosition(m.handle)

	return Vector2{
		X: float32(pos.x),
		Y: float32(pos.y),
	}
}

func (m *Menu) RowCount(column int) int {
	return int(C.menu_getRowCount(m.handle, C.uchar(column)))
}

func (m *Menu) ColumnCount() int {
	return int(C.menu_getColumnCount(m.handle))
}

func (m *Menu) ColumnWidths() Vector2 {
	widths := C.menu_getColumnWidths(m.handle)

	return Vector2{
		X: float32(widths.x),
		Y: float32(widths.y),
	}
}

func (m *Menu) Item(column, row int) string {
	item := C.menu_getCell(m.handle, C.uchar(column), C.uchar(row))

	return C.GoStringN(item.buf, C.int(item.length))
}

func (m *Menu) ShowFor(plr *Player) {
	C.menu_showForPlayer(m.handle, plr.handle)
}

func (m *Menu) HideFor(plr *Player) {
	C.menu_hideForPlayer(m.handle, plr.handle)
}
