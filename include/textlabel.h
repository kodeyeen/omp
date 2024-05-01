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
    void textLabel_detachFromPlayer(void* textLabel, float posX, float posY, float posZ);
    void textLabel_detachFromVehicle(void* textLabel, float posX, float posY, float posZ);
    void textLabel_setTestLOS(void* textLabel, unsigned char status);
    unsigned char textLabel_getTestLOS(void* textLabel);
    unsigned char textLabel_isStreamedInForPlayer(void* textLabel, void* player);
    void textLabel_setPosition(void* textLabel, float posX, float posY, float posZ);
    Vector3 textLabel_getPosition(void* textLabel);
    void textLabel_setVirtualWorld(void* textLabel, int vw);
    int textLabel_getVirtualWorld(void* textLabel);

    // PlayerTextLabel

    void* playerTextLabel_create(String text, uint32_t colour, float posX, float posY, float posZ, float drawDistance, unsigned char los);
    void playerTextLabel_release(void* textLabel, void* player);
    void playerTextLabel_setText(void* textLabel, String text);
    String playerTextLabel_getText(void* textLabel);
    void playerTextLabel_setColour(void* textLabel, uint32_t colour);
    uint32_t playerTextLabel_getColour(void* textLabel);
    void playerTextLabel_setDrawDistance(void* textLabel, float drawDist);
    float playerTextLabel_getDrawDistance(void* textLabel);
    void playerTextLabel_attachToPlayer(void* textLabel, void* player, float offsetX, float offsetY, float offsetZ);
    void playerTextLabel_attachToVehicle(void* textLabel, void* vehicle, float offsetX, float offsetY, float offsetZ);
    TextLabelAttachmentData playerTextLabel_getAttachmentData(void* textLabel);
    void playerTextLabel_detachFromPlayer(void* textLabel, float posX, float posY, float posZ);
    void playerTextLabel_detachFromVehicle(void* textLabel, float posX, float posY, float posZ);
    void playerTextLabel_setTestLOS(void* textLabel, unsigned char status);
    unsigned char playerTextLabel_getTestLOS(void* textLabel);

    // entity

    void playerTextLabel_setPosition(void* textLabel, float posX, float posY, float posZ);
    Vector3 playerTextLabel_getPosition(void* textLabel);
    void playerTextLabel_setVirtualWorld(void* textLabel, int vw);
    int playerTextLabel_getVirtualWorld(void* textLabel);

#ifdef __cplusplus
}
#endif
