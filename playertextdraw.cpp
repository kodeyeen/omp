#include "include/gomp.h"
#include "include/playertextdraw.h"

#ifdef __cplusplus
extern "C"
{
#endif

    void* playerTextDraw_create(void* player, float posX, float posY, String text)
    {
        return call<void*>("playerTextDraw_create", player, posX, posY, text);
    }

    void playerTextDraw_release(void* textdraw, void* player)
    {
        return call<void>("playerTextDraw_release", textdraw, player);
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

