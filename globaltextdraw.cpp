#include "include/gomp.h"
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

#ifdef __cplusplus
}
#endif

