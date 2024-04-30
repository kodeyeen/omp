#include "gomp.h"

#ifdef __cplusplus
extern "C"
{
#endif

    void* textDraw_create(float posX, float posY, String text);
    void textDraw_release(void* textdraw);
    int textDraw_getID(void* textdraw);
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

    // PlayerTextDraw
    void* playerTextDraw_create(void* player, float posX, float posY, String text);
    void playerTextDraw_release(void* textdraw, void* player);
    int playerTextDraw_getID(void* textdraw);
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
