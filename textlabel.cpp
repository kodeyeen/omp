#include "include/gomp.h"
#include "include/textlabel.h"

#ifdef __cplusplus
extern "C"
{
#endif

    void* textLabel_create(String text, uint32_t colour, float posX, float posY, float posZ, float drawDist, int vw, unsigned char los)
    {
        return call<void*>("textLabel_create", text, colour, posX, posY, posZ, drawDist, vw, los);
    }

    void textLabel_release(void* textLabel)
    {
        return call<void>("textLabel_release", textLabel);
    }

    void textLabel_setText(void* textLabel, String text)
    {
        return call<void>("textLabel_setText", textLabel, text);
    }

    String textLabel_getText(void* textLabel)
    {
        return call<String>("textLabel_getText", textLabel);
    }

    void textLabel_setColour(void* textLabel, uint32_t colour)
    {
        return call<void>("textLabel_setColour", textLabel, colour);
    }

    uint32_t textLabel_getColour(void* textLabel)
    {
        return call<uint32_t>("textLabel_getColour", textLabel);
    }

    void textLabel_setDrawDistance(void* textLabel, float drawDist)
    {
        return call<void>("textLabel_setDrawDistance", textLabel, drawDist);
    }

    float textLabel_getDrawDistance(void* textLabel)
    {
        return call<float>("textLabel_getDrawDistance", textLabel);
    }

    void textLabel_attachToPlayer(void* textLabel, void* player, float offsetX, float offsetY, float offsetZ)
    {
        return call<void>("textLabel_attachToPlayer", textLabel, player, offsetX, offsetY, offsetZ);
    }

    void textLabel_attachToVehicle(void* textLabel, void* vehicle, float offsetX, float offsetY, float offsetZ)
    {
        return call<void>("textLabel_attachToVehicle", textLabel, vehicle, offsetX, offsetY, offsetZ);
    }

    TextLabelAttachmentData textLabel_getAttachmentData(void* textLabel)
    {
        return call<TextLabelAttachmentData>("textLabel_getAttachmentData", textLabel);
    }

    void textLabel_detachFromPlayer(void* textLabel, void* player, float posX, float posY, float posZ)
    {
        return call<void>("textLabel_detachFromPlayer", textLabel, player, posX, posY, posZ);
    }

    void textLabel_detachFromVehicle(void* textLabel, void* vehicle, float posX, float posY, float posZ)
    {
        return call<void>("textLabel_detachFromVehicle", textLabel, vehicle, posX, posY, posZ);
    }

    void textLabel_setTestLOS(void* textLabel, unsigned char status)
    {
        return call<void>("textLabel_setTestLOS", textLabel, status);
    }

    unsigned char textLabel_getTestLOS(void* textLabel)
    {
        return call<unsigned char>("textLabel_getTestLOS", textLabel);
    }

    unsigned char textLabel_isStreamedInForPlayer(void* textLabel, void* player)
    {
        return call<unsigned char>("textLabel_isStreamedInForPlayer", textLabel, player);
    }

#ifdef __cplusplus
}
#endif
