#include "gomp.h"

#ifdef __cplusplus
extern "C"
{
#endif

    void* textDraw_create(float posX, float posY, String text);
    void textDraw_release(void* textdraw);
    void textDraw_setPosition(void* textdraw, float posX, float posY);
    Vector2 textDraw_getPosition(void* textdraw);
    void textDraw_setText(void* textdraw, String text);
    String textDraw_getText(void* textdraw);
    void textDraw_setLetterSize(void* textdraw, float sizeX, float sizeY);
    Vector2 textDraw_getLetterSize(void* textdraw);
    void textDraw_setTextSize(void* textdraw, float sizeX, float sizeY);
    Vector2 textDraw_getTextSize(void* textdraw);
    void textDraw_setAlignment(void* textdraw, int alignment);
    int textDraw_getAlignment(void* textdraw);
    void textDraw_setColour(void* textdraw, uint32_t colour);
    int textDraw_getLetterColour(void* textdraw);
    void textDraw_useBox(void* textdraw, int use);
    int textDraw_hasBox(void* textdraw);
    void textDraw_setBoxColour(void* textdraw, uint32_t colour);
    int textDraw_getBoxColour(void* textdraw);
    void textDraw_setShadow(void* textdraw, int size);
    int textDraw_getShadow(void* textdraw);
    void textDraw_setOutline(void* textdraw, int size);
    int textDraw_getOutline(void* textdraw);
    void textDraw_setBackgroundColour(void* textdraw, uint32_t colour);
    int textDraw_getBackgroundColour(void* textdraw);
    void textDraw_setStyle(void* textdraw, int style);
    int textDraw_getStyle(void* textdraw);
    void textDraw_setProportional(void* textdraw, int set);
    int textDraw_isProportional(void* textdraw);
    void textDraw_setSelectable(void* textdraw, int set);
    int textDraw_isSelectable(void* textdraw);
    void textDraw_setPreviewModel(void* textdraw, int model);
    int textDraw_getPreviewModel(void* textdraw);
    void textDraw_setPreviewRotation(void* textdraw, float rotX, float rotY, float rotZ);
    Vector3 textDraw_getPreviewRotation(void* textdraw);
    void textDraw_setPreviewVehicleColour(void* textdraw, int col1, int col2);
    VehicleColour textDraw_getPreviewVehicleColour(void* textdraw);
    void textDraw_setPreviewZoom(void* textdraw, float zoom);
    float textDraw_getPreviewZoom(void* textdraw);
    void textDraw_showForPlayer(void* textdraw, void* player);
    void textDraw_hideForPlayer(void* textdraw, void* player);
    int textDraw_isShownForPlayer(void* textdraw, void* player);
    void textDraw_setTextForPlayer(void* textdraw, void* player, String text);

#ifdef __cplusplus
}
#endif
