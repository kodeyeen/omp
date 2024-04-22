#include "gomp.h"

#ifdef __cplusplus
extern "C"
{
#endif

    typedef struct
    {
        int playerID;
        int vehicleID;
    } TextLabelAttachmentData;

    void* textLabel_create(String text, uint32_t colour, float posX, float posY, float posZ, float drawDist, int vw, unsigned char los);
    void textLabel_release(void* textLabel);
    void textLabel_setText(void* textLabel, String text);
    String textLabel_getText(void* textLabel);
    void textLabel_setColour(void* textLabel, uint32_t colour);
    uint32_t textLabel_getColour(void* textLabel);
    void textLabel_setDrawDistance(void* textLabel, float drawDist);
    float textLabel_getDrawDistance(void* textLabel);
    void textLabel_attachToPlayer(void* textLabel, void* player, float offsetX, float offsetY, float offsetZ);
    void textLabel_attachToVehicle(void* textLabel, void* vehicle, float offsetX, float offsetY, float offsetZ);
    TextLabelAttachmentData textLabel_getAttachmentData(void* textLabel);
    void textLabel_detachFromPlayer(void* textLabel, void* player, float posX, float posY, float posZ);
    void textLabel_detachFromVehicle(void* textLabel, void* vehicle, float posX, float posY, float posZ);
    void textLabel_setTestLOS(void* textLabel, unsigned char status);
    unsigned char textLabel_getTestLOS(void* textLabel);
    unsigned char textLabel_isStreamedInForPlayer(void* textLabel, void* player);

#ifdef __cplusplus
}
#endif
