#include "omp.h"

#ifdef __cplusplus
extern "C" {
#endif

    void* menu_create(String title, float posX, float posY, uint8_t columns, float col1Width, float col2Width);
    void menu_release(void* menu);
    void menu_setColumnHeader(void* menu, String header, uint8_t column);
    void menu_addCell(void* menu, String itemText, uint8_t column);
    void menu_disableRow(void* menu, uint8_t row);
    unsigned char menu_isRowEnabled(void* menu, uint8_t row);
    void menu_disable(void* menu);
    unsigned char menu_isEnabled(void* menu);
    Vector2 menu_getPosition(void* menu);
    int menu_getRowCount(void* menu, uint8_t column);
    int menu_getColumnCount(void* menu);
    Vector2 menu_getColumnWidths(void* menu);
    String menu_getColumnHeader(void* menu, uint8_t column);
    String menu_getCell(void* menu, uint8_t column, uint8_t row);
    void menu_showForPlayer(void* menu, void* player);
    void menu_hideForPlayer(void* menu, void* player);

#ifdef __cplusplus
}
#endif
