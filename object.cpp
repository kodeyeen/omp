#include "include/gomp.h"
#include "include/object.h"

#ifdef __cplusplus
extern "C"
{
#endif

    void* object_create(int modelId, float posX, float posY, float posZ, float rotX, float rotY, float rotZ, float drawDistance)
    {
        return call<void*>("object_create", modelId, posX, posY, posZ, rotX, rotY, rotZ, drawDistance);
    }

    void object_release(void* object)
    {
        return call<void>("object_release", object);
    }

    void* object_getByID(int id)
    {
        return call<void*>("object_getByID", id);
    }

    void object_setDefaultCameraCollision(unsigned char set)
    {
        return call<void>("object_setDefaultCameraCollision", set);
    }

    void object_setDrawDistance(void* object, float drawDistance)
    {
        return call<void>("object_setDrawDistance", object, drawDistance);
    }

    float object_getDrawDistance(void* object)
    {
        return call<float>("object_getDrawDistance", object);
    }

    void object_setModel(void* object, int model)
    {
        return call<void>("object_setModel", object, model);
    }

    int object_getModel(void* object)
    {
        return call<int>("object_getModel", object);
    }

    void object_setCameraCollision(void* object, unsigned char set)
    {
        return call<void>("object_setCameraCollision", object, set);
    }

    unsigned char object_getCameraCollision(void* object)
    {
        return call<unsigned char>("object_getCameraCollision", object);
    }

    int object_move(void* object, float posX, float posY, float posZ, float rotX, float rotY, float rotZ, float speed)
    {
        return call<int>("object_move", object, posX, posY, posZ, rotX, rotY, rotZ, speed);
    }

    unsigned char object_isMoving(void* object)
    {
        return call<unsigned char>("object_isMoving", object);
    }

    void object_stop(void* object)
    {
        return call<void>("object_stop", object);
    }

    ObjectMoveData object_getMovingData(void* object)
    {
        return call<ObjectMoveData>("object_getMovingData", object);
    }

    void object_attachToVehicle(void* object, void* vehicle, float offsetX, float offsetY, float offsetZ, float rotX, float rotY, float rotZ)
    {
        return call<void>("object_attachToVehicle", object, vehicle, offsetX, offsetY, offsetZ, rotX, rotY, rotZ);
    }

    void object_resetAttachment(void* object)
    {
        return call<void>("object_resetAttachment", object);
    }

    ObjectAttachmentData object_getAttachmentData(void* object)
    {
        return call<ObjectAttachmentData>("object_getAttachmentData", object);
    }

    unsigned char object_isMaterialSlotUsed(void* object, uint32_t materialIndex)
    {
        return call<unsigned char>("object_isMaterialSlotUsed", object, materialIndex);
    }

    unsigned char object_getMaterial(void* object, uint32_t materialIndex, const ObjectMaterial* out)
    {
        return call<unsigned char>("object_getMaterial", object, materialIndex, out);
    }

    unsigned char object_getMaterialText(void* object, uint32_t materialIndex, const ObjectMaterialText* out)
    {
        return call<unsigned char>("object_getMaterialText", object, materialIndex, out);
    }

    void object_setMaterial(void* object, uint32_t materialIndex, int model, String textureLibrary, String textureName, uint32_t colour)
    {
        return call<void>("object_setMaterial", object, materialIndex, model, textureLibrary, textureName, colour);
    }

    void object_setMaterialText(void* object, uint32_t materialIndex, String text, ObjectMaterialSize materialSize, String fontFace, int fontSize, unsigned char bold, uint32_t fontColour, uint32_t backgroundColour, ObjectMaterialTextAlign align)
    {
        return call<void>("object_setMaterialText", object, materialIndex, text, materialSize, fontFace, fontSize, bold, fontColour, backgroundColour, align);
    }

    void object_attachToPlayer(void* object, void* player, float offsetX, float offsetY, float offsetZ, float rotX, float rotY, float rotZ)
    {
        return call<void>("object_attachToPlayer", object, player, offsetX, offsetY, offsetZ, rotX, rotY, rotZ);
    }

    void object_attachToObject(void* object, void* other, float offsetX, float offsetY, float offsetZ, float rotX, float rotY, float rotZ, unsigned char syncRotation)
    {
        return call<void>("object_attachToObject", object, other, offsetX, offsetY, offsetZ, rotX, rotY, rotZ, syncRotation);
    }

    // entity

    void object_setPosition(void* object, float posX, float posY, float posZ)
    {
        return call<void>("object_setPosition", object, posX, posY, posZ);
    }

    Vector3 object_getPosition(void* object)
    {
        return call<Vector3>("object_getPosition", object);
    }

    void object_setRotation(void* object, float rotX, float rotY, float rotZ)
    {
        return call<void>("object_setRotation", object, rotX, rotY, rotZ);
    }

    Vector3 object_getRotation(void* object)
    {
        return call<Vector3>("object_getRotation", object);
    }

#ifdef __cplusplus
}
#endif
