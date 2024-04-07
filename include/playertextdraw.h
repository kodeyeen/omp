#include "gomp.h"

#ifdef __cplusplus
extern "C"
{
#endif

    void* playerTextDraw_create(void* player, float posX, float posY, String text);
    void playerTextDraw_release(void* textdraw, void* player);
    void playerTextDraw_setPosition(void* textdraw, float posX, float posY);
    Vector2 playerTextDraw_getPosition(void* textdraw);
    void playerTextDraw_setText(void* textdraw, String text);
    String playerTextDraw_getText(void* textdraw);
    void playerTextDraw_setLetterSize(void* textdraw, float sizeX, float sizeY);
    Vector2 playerTextDraw_getLetterSize(void* textdraw);
    void playerTextDraw_setTextSize(void* textdraw, float sizeX, float sizeY);
    Vector2 playerTextDraw_getTextSize(void* textdraw);
    void playerTextDraw_setAlignment(void* textdraw, int alignment);
    int playerTextDraw_getAlignment(void* textdraw);
    void playerTextDraw_setColour(void* textdraw, uint32_t colour);
    int playerTextDraw_getLetterColour(void* textdraw);
    void playerTextDraw_useBox(void* textdraw, int use);
    int playerTextDraw_hasBox(void* textdraw);
    void playerTextDraw_setBoxColour(void* textdraw, uint32_t colour);
    int playerTextDraw_getBoxColour(void* textdraw);
    void playerTextDraw_setShadow(void* textdraw, int size);
    int playerTextDraw_getShadow(void* textdraw);
    void playerTextDraw_setOutline(void* textdraw, int size);
    int playerTextDraw_getOutline(void* textdraw);
    void playerTextDraw_setBackgroundColour(void* textdraw, uint32_t colour);
    int playerTextDraw_getBackgroundColour(void* textdraw);
    void playerTextDraw_setStyle(void* textdraw, int style);
    int playerTextDraw_getStyle(void* textdraw);
    void playerTextDraw_setProportional(void* textdraw, int set);
    int playerTextDraw_isProportional(void* textdraw);
    void playerTextDraw_setSelectable(void* textdraw, int set);
    int playerTextDraw_isSelectable(void* textdraw);
    void playerTextDraw_setPreviewModel(void* textdraw, int model);
    int playerTextDraw_getPreviewModel(void* textdraw);
    void playerTextDraw_setPreviewRotation(void* textdraw, float rotX, float rotY, float rotZ);
    Vector3 playerTextDraw_getPreviewRotation(void* textdraw);
    void playerTextDraw_setPreviewVehicleColour(void* textdraw, int col1, int col2);
    VehicleColour playerTextDraw_getPreviewVehicleColour(void* textdraw);
    void playerTextDraw_setPreviewZoom(void* textdraw, float zoom);
    float playerTextDraw_getPreviewZoom(void* textdraw);
    void playerTextDraw_show(void* textdraw);
    void playerTextDraw_hide(void* textdraw);
    int playerTextDraw_isShown(void* textdraw);

#ifdef __cplusplus
}
#endif
