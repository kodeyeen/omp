#include "include/omp.h"
#include "include/textdraw.h"

#ifdef __cplusplus
extern "C"
{
#endif

    void* textDraw_create(float posX, float posY, String text)
    {
        return call<void*>("textDraw_create", posX, posY, text);
    }

    void textDraw_release(void* textdraw)
    {
        return call<void>("textDraw_release", textdraw);
    }

    int textDraw_getID(void* textdraw)
    {
        return call<int>("textDraw_getID", textdraw);
    }

    void textDraw_setPosition(void* textdraw, float posX, float posY)
    {
        return call<void>("textDraw_setPosition", textdraw, posX, posY);
    }

    Vector2 textDraw_getPosition(void* textdraw)
    {
        return call<Vector2>("textDraw_getPosition", textdraw);
    }

    void textDraw_setText(void* textdraw, String text)
    {
        return call<void>("textDraw_setText", textdraw, text);
    }

    String textDraw_getText(void* textdraw)
    {
        return call<String>("textDraw_getText", textdraw);
    }

    void textDraw_setLetterSize(void* textdraw, float sizeX, float sizeY)
    {
        return call<void>("textDraw_setLetterSize", textdraw, sizeX, sizeY);
    }

    Vector2 textDraw_getLetterSize(void* textdraw)
    {
        return call<Vector2>("textDraw_getLetterSize", textdraw);
    }

    void textDraw_setTextSize(void* textdraw, float sizeX, float sizeY)
    {
        return call<void>("textDraw_setTextSize", textdraw, sizeX, sizeY);
    }

    Vector2 textDraw_getTextSize(void* textdraw)
    {
        return call<Vector2>("textDraw_getTextSize", textdraw);
    }

    void textDraw_setAlignment(void* textdraw, int alignment)
    {
        return call<void>("textDraw_setAlignment", textdraw, alignment);
    }

    int textDraw_getAlignment(void* textdraw)
    {
        return call<int>("textDraw_getAlignment", textdraw);
    }

    void textDraw_setColour(void* textdraw, uint32_t colour)
    {
        return call<void>("textDraw_setColour", textdraw, colour);
    }

    int textDraw_getLetterColour(void* textdraw)
    {
        return call<int>("textDraw_getLetterColour", textdraw);
    }

    void textDraw_useBox(void* textdraw, int use)
    {
        return call<void>("textDraw_useBox", textdraw, use);
    }

    int textDraw_hasBox(void* textdraw)
    {
        return call<int>("textDraw_hasBox", textdraw);
    }

    void textDraw_setBoxColour(void* textdraw, uint32_t colour)
    {
        return call<void>("textDraw_setBoxColour", textdraw, colour);
    }

    int textDraw_getBoxColour(void* textdraw)
    {
        return call<int>("textDraw_getBoxColour", textdraw);
    }

    void textDraw_setShadow(void* textdraw, int size)
    {
        return call<void>("textDraw_setShadow", textdraw, size);
    }

    int textDraw_getShadow(void* textdraw)
    {
        return call<int>("textDraw_getShadow", textdraw);
    }

    void textDraw_setOutline(void* textdraw, int size)
    {
        return call<void>("textDraw_setOutline", textdraw, size);
    }

    int textDraw_getOutline(void* textdraw)
    {
        return call<int>("textDraw_getOutline", textdraw);
    }

    void textDraw_setBackgroundColour(void* textdraw, uint32_t colour)
    {
        return call<void>("textDraw_setBackgroundColour", textdraw, colour);
    }

    int textDraw_getBackgroundColour(void* textdraw)
    {
        return call<int>("textDraw_getBackgroundColour", textdraw);
    }

    void textDraw_setStyle(void* textdraw, int style)
    {
        return call<void>("textDraw_setStyle", textdraw, style);
    }

    int textDraw_getStyle(void* textdraw)
    {
        return call<int>("textDraw_getStyle", textdraw);
    }

    void textDraw_setProportional(void* textdraw, int set)
    {
        return call<void>("textDraw_setProportional", textdraw, set);
    }

    int textDraw_isProportional(void* textdraw)
    {
        return call<int>("textDraw_isProportional", textdraw);
    }

    void textDraw_setSelectable(void* textdraw, int set)
    {
        return call<void>("textDraw_setSelectable", textdraw, set);
    }

    int textDraw_isSelectable(void* textdraw)
    {
        return call<int>("textDraw_isSelectable", textdraw);
    }

    void textDraw_setPreviewModel(void* textdraw, int model)
    {
        return call<void>("textDraw_setPreviewModel", textdraw, model);
    }

    int textDraw_getPreviewModel(void* textdraw)
    {
        return call<int>("textDraw_getPreviewModel", textdraw);
    }

    void textDraw_setPreviewRotation(void* textdraw, float rotX, float rotY, float rotZ)
    {
        return call<void>("textDraw_setPreviewRotation", textdraw, rotX, rotY, rotZ);
    }

    Vector3 textDraw_getPreviewRotation(void* textdraw)
    {
        return call<Vector3>("textDraw_getPreviewRotation", textdraw);
    }

    void textDraw_setPreviewVehicleColour(void* textdraw, int col1, int col2)
    {
        return call<void>("textDraw_setPreviewVehicleColour", textdraw, col1, col2);
    }

    VehicleColour textDraw_getPreviewVehicleColour(void* textdraw)
    {
        return call<VehicleColour>("textDraw_getPreviewVehicleColour", textdraw);
    }

    void textDraw_setPreviewZoom(void* textdraw, float zoom)
    {
        return call<void>("textDraw_setPreviewZoom", textdraw, zoom);
    }

    float textDraw_getPreviewZoom(void* textdraw)
    {
        return call<float>("textDraw_getPreviewZoom", textdraw);
    }

    void textDraw_showForPlayer(void* textdraw, void* player)
    {
        return call<void>("textDraw_showForPlayer", textdraw, player);
    }

    void textDraw_hideForPlayer(void* textdraw, void* player)
    {
        return call<void>("textDraw_hideForPlayer", textdraw, player);
    }

    int textDraw_isShownForPlayer(void* textdraw, void* player)
    {
        return call<int>("textDraw_isShownForPlayer", textdraw, player);
    }

    void textDraw_setTextForPlayer(void* textdraw, void* player, String text)
    {
        return call<void>("textDraw_setTextForPlayer", textdraw, player, text);
    }

    // Player textdraw

    void* playerTextDraw_create(void* player, float posX, float posY, String text)
    {
        return call<void*>("playerTextDraw_create", player, posX, posY, text);
    }

    void playerTextDraw_release(void* textdraw, void* player)
    {
        return call<void>("playerTextDraw_release", player, textdraw);
    }

    int playerTextDraw_getID(void* textdraw)
    {
        return call<int>("playerTextDraw_getID", textdraw);
    }

    void playerTextDraw_setPosition(void* textdraw, float posX, float posY)
    {
        return call<void>("playerTextDraw_setPosition", textdraw, posX, posY);
    }

    Vector2 playerTextDraw_getPosition(void* textdraw)
    {
        return call<Vector2>("playerTextDraw_getPosition", textdraw);
    }

    void playerTextDraw_setText(void* textdraw, String text)
    {
        return call<void>("playerTextDraw_setText", textdraw, text);
    }

    String playerTextDraw_getText(void* textdraw)
    {
        return call<String>("playerTextDraw_getText", textdraw);
    }

    void playerTextDraw_setLetterSize(void* textdraw, float sizeX, float sizeY)
    {
        return call<void>("playerTextDraw_setLetterSize", textdraw, sizeX, sizeY);
    }

    Vector2 playerTextDraw_getLetterSize(void* textdraw)
    {
        return call<Vector2>("playerTextDraw_getLetterSize", textdraw);
    }

    void playerTextDraw_setTextSize(void* textdraw, float sizeX, float sizeY)
    {
        return call<void>("playerTextDraw_setTextSize", textdraw, sizeX, sizeY);
    }

    Vector2 playerTextDraw_getTextSize(void* textdraw)
    {
        return call<Vector2>("playerTextDraw_getTextSize", textdraw);
    }

    void playerTextDraw_setAlignment(void* textdraw, int alignment)
    {
        return call<void>("playerTextDraw_setAlignment", textdraw, alignment);
    }

    int playerTextDraw_getAlignment(void* textdraw)
    {
        return call<int>("playerTextDraw_getAlignment", textdraw);
    }

    void playerTextDraw_setColour(void* textdraw, uint32_t colour)
    {
        return call<void>("playerTextDraw_setColour", textdraw, colour);
    }

    int playerTextDraw_getLetterColour(void* textdraw)
    {
        return call<int>("playerTextDraw_getLetterColour", textdraw);
    }

    void playerTextDraw_useBox(void* textdraw, int use)
    {
        return call<void>("playerTextDraw_useBox", textdraw, use);
    }

    int playerTextDraw_hasBox(void* textdraw)
    {
        return call<int>("playerTextDraw_hasBox", textdraw);
    }

    void playerTextDraw_setBoxColour(void* textdraw, uint32_t colour)
    {
        return call<void>("playerTextDraw_setBoxColour", textdraw, colour);
    }

    int playerTextDraw_getBoxColour(void* textdraw)
    {
        return call<int>("playerTextDraw_getBoxColour", textdraw);
    }

    void playerTextDraw_setShadow(void* textdraw, int size)
    {
        return call<void>("playerTextDraw_setShadow", textdraw, size);
    }

    int playerTextDraw_getShadow(void* textdraw)
    {
        return call<int>("playerTextDraw_getShadow", textdraw);
    }

    void playerTextDraw_setOutline(void* textdraw, int size)
    {
        return call<void>("playerTextDraw_setOutline", textdraw, size);
    }

    int playerTextDraw_getOutline(void* textdraw)
    {
        return call<int>("playerTextDraw_getOutline", textdraw);
    }

    void playerTextDraw_setBackgroundColour(void* textdraw, uint32_t colour)
    {
        return call<void>("playerTextDraw_setBackgroundColour", textdraw, colour);
    }

    int playerTextDraw_getBackgroundColour(void* textdraw)
    {
        return call<int>("playerTextDraw_getBackgroundColour", textdraw);
    }

    void playerTextDraw_setStyle(void* textdraw, int style)
    {
        return call<void>("playerTextDraw_setStyle", textdraw, style);
    }

    int playerTextDraw_getStyle(void* textdraw)
    {
        return call<int>("playerTextDraw_getStyle", textdraw);
    }

    void playerTextDraw_setProportional(void* textdraw, int set)
    {
        return call<void>("playerTextDraw_setProportional", textdraw, set);
    }

    int playerTextDraw_isProportional(void* textdraw)
    {
        return call<int>("playerTextDraw_isProportional", textdraw);
    }

    void playerTextDraw_setSelectable(void* textdraw, int set)
    {
        return call<void>("playerTextDraw_setSelectable", textdraw, set);
    }

    int playerTextDraw_isSelectable(void* textdraw)
    {
        return call<int>("playerTextDraw_isSelectable", textdraw);
    }

    void playerTextDraw_setPreviewModel(void* textdraw, int model)
    {
        return call<void>("playerTextDraw_setPreviewModel", textdraw, model);
    }

    int playerTextDraw_getPreviewModel(void* textdraw)
    {
        return call<int>("playerTextDraw_getPreviewModel", textdraw);
    }

    void playerTextDraw_setPreviewRotation(void* textdraw, float rotX, float rotY, float rotZ)
    {
        return call<void>("playerTextDraw_setPreviewRotation", textdraw, rotX, rotY, rotZ);
    }

    Vector3 playerTextDraw_getPreviewRotation(void* textdraw)
    {
        return call<Vector3>("playerTextDraw_getPreviewRotation", textdraw);
    }

    void playerTextDraw_setPreviewVehicleColour(void* textdraw, int col1, int col2)
    {
        return call<void>("playerTextDraw_setPreviewVehicleColour", textdraw, col1, col2);
    }

    VehicleColour playerTextDraw_getPreviewVehicleColour(void* textdraw)
    {
        return call<VehicleColour>("playerTextDraw_getPreviewVehicleColour", textdraw);
    }

    void playerTextDraw_setPreviewZoom(void* textdraw, float zoom)
    {
        return call<void>("playerTextDraw_setPreviewZoom", textdraw, zoom);
    }

    float playerTextDraw_getPreviewZoom(void* textdraw)
    {
        return call<float>("playerTextDraw_getPreviewZoom", textdraw);
    }

    void playerTextDraw_show(void* textdraw)
    {
        return call<void>("playerTextDraw_show", textdraw);
    }

    void playerTextDraw_hide(void* textdraw)
    {
        return call<void>("playerTextDraw_hide", textdraw);
    }

    int playerTextDraw_isShown(void* textdraw)
    {
        return call<int>("playerTextDraw_isShown", textdraw);
    }

#ifdef __cplusplus
}
#endif
