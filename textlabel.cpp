#include "include/textlabel.h"

extern "C" {
    void* textLabel_create(String text, uint32_t colour, float posX, float posY, float posZ, float drawDist, int vw, unsigned char los) {
        return call<void*>("textLabel_create", text, colour, posX, posY, posZ, drawDist, vw, los);
    }

    void textLabel_release(void* textLabel) {
        return call<void>("textLabel_release", textLabel);
    }

    void textLabel_setText(void* textLabel, String text) {
        return call<void>("textLabel_setText", textLabel, text);
    }

    String textLabel_getText(void* textLabel) {
        return call<String>("textLabel_getText", textLabel);
    }

    void textLabel_setColour(void* textLabel, uint32_t colour) {
        return call<void>("textLabel_setColour", textLabel, colour);
    }

    uint32_t textLabel_getColour(void* textLabel) {
        return call<uint32_t>("textLabel_getColour", textLabel);
    }

    void textLabel_setDrawDistance(void* textLabel, float drawDist) {
        return call<void>("textLabel_setDrawDistance", textLabel, drawDist);
    }

    float textLabel_getDrawDistance(void* textLabel) {
        return call<float>("textLabel_getDrawDistance", textLabel);
    }

    void textLabel_attachToPlayer(void* textLabel, void* player, float offsetX, float offsetY, float offsetZ) {
        return call<void>("textLabel_attachToPlayer", textLabel, player, offsetX, offsetY, offsetZ);
    }

    void textLabel_attachToVehicle(void* textLabel, void* vehicle, float offsetX, float offsetY, float offsetZ) {
        return call<void>("textLabel_attachToVehicle", textLabel, vehicle, offsetX, offsetY, offsetZ);
    }

    TextLabelAttachmentData textLabel_getAttachmentData(void* textLabel) {
        return call<TextLabelAttachmentData>("textLabel_getAttachmentData", textLabel);
    }

    void textLabel_detachFromPlayer(void* textLabel, float posX, float posY, float posZ) {
        return call<void>("textLabel_detachFromPlayer", textLabel, posX, posY, posZ);
    }

    void textLabel_detachFromVehicle(void* textLabel, float posX, float posY, float posZ) {
        return call<void>("textLabel_detachFromVehicle", textLabel, posX, posY, posZ);
    }

    void textLabel_setTestLOS(void* textLabel, unsigned char status) {
        return call<void>("textLabel_setTestLOS", textLabel, status);
    }

    unsigned char textLabel_getTestLOS(void* textLabel) {
        return call<unsigned char>("textLabel_getTestLOS", textLabel);
    }

    unsigned char textLabel_isStreamedInForPlayer(void* textLabel, void* player) {
        return call<unsigned char>("textLabel_isStreamedInForPlayer", textLabel, player);
    }

    void textLabel_setPosition(void* textLabel, float posX, float posY, float posZ) {
        return call<void>("textLabel_setPosition", textLabel, posX, posY, posZ);
    }

    Vector3 textLabel_getPosition(void* textLabel) {
        return call<Vector3>("textLabel_getPosition", textLabel);
    }

    void textLabel_setVirtualWorld(void* textLabel, int vw) {
        return call<void>("textLabel_setVirtualWorld", textLabel, vw);
    }

    int textLabel_getVirtualWorld(void* textLabel) {
        return call<int>("textLabel_getVirtualWorld", textLabel);
    }

    void* playerTextLabel_create(void* player, String text, uint32_t colour, float posX, float posY, float posZ, float drawDistance, unsigned char los) {
        return call<void*>("playerTextLabel_create", player, text, colour, posX, posY, posZ, drawDistance, los);
    }

    void playerTextLabel_release(void* textLabel, void* player) {
        return call<void>("playerTextLabel_release", textLabel, player);
    }

    void playerTextLabel_setText(void* textLabel, String text) {
        return call<void>("playerTextLabel_setText", textLabel, text);
    }

    String playerTextLabel_getText(void* textLabel) {
        return call<String>("playerTextLabel_getText", textLabel);
    }

    void playerTextLabel_setColour(void* textLabel, uint32_t colour) {
        return call<void>("playerTextLabel_setColour", textLabel, colour);
    }

    uint32_t playerTextLabel_getColour(void* textLabel) {
        return call<uint32_t>("playerTextLabel_getColour", textLabel);
    }

    void playerTextLabel_setDrawDistance(void* textLabel, float drawDist) {
        return call<void>("playerTextLabel_setDrawDistance", textLabel, drawDist);
    }

    float playerTextLabel_getDrawDistance(void* textLabel) {
        return call<float>("playerTextLabel_getDrawDistance", textLabel);
    }

    void playerTextLabel_attachToPlayer(void* textLabel, void* player, float offsetX, float offsetY, float offsetZ) {
        return call<void>("playerTextLabel_attachToPlayer", textLabel, player, offsetX, offsetY, offsetZ);
    }

    void playerTextLabel_attachToVehicle(void* textLabel, void* vehicle, float offsetX, float offsetY, float offsetZ) {
        return call<void>("playerTextLabel_attachToVehicle", textLabel, vehicle, offsetX, offsetY, offsetZ);
    }

    TextLabelAttachmentData playerTextLabel_getAttachmentData(void* textLabel) {
        return call<TextLabelAttachmentData>("playerTextLabel_getAttachmentData", textLabel);
    }

    void playerTextLabel_detachFromPlayer(void* textLabel, float posX, float posY, float posZ) {
        return call<void>("playerTextLabel_detachFromPlayer", textLabel, posX, posY, posZ);
    }

    void playerTextLabel_detachFromVehicle(void* textLabel, float posX, float posY, float posZ) {
        return call<void>("playerTextLabel_detachFromVehicle", textLabel, posX, posY, posZ);
    }

    void playerTextLabel_setTestLOS(void* textLabel, unsigned char status) {
        return call<void>("playerTextLabel_setTestLOS", textLabel, status);
    }

    unsigned char playerTextLabel_getTestLOS(void* textLabel) {
        return call<unsigned char>("playerTextLabel_getTestLOS", textLabel);
    }

    void playerTextLabel_setPosition(void* textLabel, float posX, float posY, float posZ) {
        return call<void>("playerTextLabel_setPosition", textLabel, posX, posY, posZ);
    }

    Vector3 playerTextLabel_getPosition(void* textLabel) {
        return call<Vector3>("playerTextLabel_getPosition", textLabel);
    }

    void playerTextLabel_setVirtualWorld(void* textLabel, int vw) {
        return call<void>("playerTextLabel_setVirtualWorld", textLabel, vw);
    }

    int playerTextLabel_getVirtualWorld(void* textLabel) {
        return call<int>("playerTextLabel_getVirtualWorld", textLabel);
    }
}
