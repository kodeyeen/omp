#include "include/menu.h"

extern "C" {
    void* menu_create(String title, float posX, float posY, uint8_t columns, float col1Width, float col2Width) {
        return call<void*>("menu_create", title, posX, posY, columns, col1Width, col2Width);
    }

    void menu_release(void* menu) {
        return call<void>("menu_release", menu);
    }

    void menu_setColumnHeader(void* menu, String header, uint8_t column) {
        return call<void>("menu_setColumnHeader", menu, header, column);
    }

    void menu_addCell(void* menu, String itemText, uint8_t column) {
        return call<void>("menu_addCell", menu, itemText, column);
    }

    void menu_disableRow(void* menu, uint8_t row) {
        return call<void>("menu_disableRow", menu, row);
    }

    unsigned char menu_isRowEnabled(void* menu, uint8_t row) {
        return call<unsigned char>("menu_isRowEnabled", menu, row);
    }

    void menu_disable(void* menu) {
        return call<void>("menu_disable", menu);
    }

    unsigned char menu_isEnabled(void* menu) {
        return call<unsigned char>("menu_isEnabled", menu);
    }

    Vector2 menu_getPosition(void* menu) {
        return call<Vector2>("menu_getPosition", menu);
    }

    int menu_getRowCount(void* menu, uint8_t column) {
        return call<int>("menu_getRowCount", menu, column);
    }

    int menu_getColumnCount(void* menu) {
        return call<int>("menu_getColumnCount", menu);
    }

    Vector2 menu_getColumnWidths(void* menu) {
        return call<Vector2>("menu_getColumnWidths", menu);
    }

    String menu_getColumnHeader(void* menu, uint8_t column) {
        return call<String>("menu_getColumnHeader", menu, column);
    }

    String menu_getCell(void* menu, uint8_t column, uint8_t row) {
        return call<String>("menu_getCell", menu, column, row);
    }

    void menu_showForPlayer(void* menu, void* player) {
        return call<void>("menu_showForPlayer", menu, player);
    }

    void menu_hideForPlayer(void* menu, void* player) {
        return call<void>("menu_hideForPlayer", menu, player);
    }
}
