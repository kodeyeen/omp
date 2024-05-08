#include "omp.h"

#ifdef __cplusplus
extern "C" {
#endif

    typedef struct {
        Vector3 targetPos;
        Vector3 targetRot;
        float speed;
    } ObjectMoveData;

    typedef struct {
        uint8_t type;
        unsigned char syncRotation;
        int ID;
        Vector3 offset;
        Vector3 rotation;
    } ObjectAttachmentData;

    typedef enum {
        ObjectMaterialSize_32x32 = 10,
        ObjectMaterialSize_64x32 = 20,
        ObjectMaterialSize_64x64 = 30,
        ObjectMaterialSize_128x32 = 40,
        ObjectMaterialSize_128x64 = 50,
        ObjectMaterialSize_128x128 = 60,
        ObjectMaterialSize_256x32 = 70,
        ObjectMaterialSize_256x64 = 80,
        ObjectMaterialSize_256x128 = 90,
        ObjectMaterialSize_256x256 = 100,
        ObjectMaterialSize_512x64 = 110,
        ObjectMaterialSize_512x128 = 120,
        ObjectMaterialSize_512x256 = 130,
        ObjectMaterialSize_512x512 = 140
    } ObjectMaterialSize;

    typedef enum {
        ObjectMaterialTextAlign_Left,
        ObjectMaterialTextAlign_Center,
        ObjectMaterialTextAlign_Right
    } ObjectMaterialTextAlign;

    typedef struct {
        int model;
        String textureLibrary;
        String textureName;
        uint32_t colour;
    } ObjectMaterial;

    typedef struct {
        String text;
        uint8_t materialSize;
        String fontFace;
        uint8_t fontSize;
        unsigned char bold;
        uint32_t fontColour;
        uint32_t backgroundColour;
        uint8_t alignment;
    } ObjectMaterialText;

    void* object_create(int modelId, float posX, float posY, float posZ, float rotX, float rotY, float rotZ, float drawDistance);
    void object_release(void* object);
    void* object_getByID(int id);
    void object_setDefaultCameraCollision(unsigned char set);
    void object_setDrawDistance(void* object, float drawDistance);
    float object_getDrawDistance(void* object);
    void object_setModel(void* object, int model);
    int object_getModel(void* object);
    void object_setCameraCollision(void* object, unsigned char set);
    unsigned char object_getCameraCollision(void* object);
    int object_move(void* object, float posX, float posY, float posZ, float rotX, float rotY, float rotZ, float speed);
    unsigned char object_isMoving(void* object);
    void object_stop(void* object);
    ObjectMoveData object_getMovingData(void* object);
    void object_attachToVehicle(void* object, void* vehicle, float offsetX, float offsetY, float offsetZ, float rotX, float rotY, float rotZ);
    void object_resetAttachment(void* object);
    ObjectAttachmentData object_getAttachmentData(void* object);
    unsigned char object_isMaterialSlotUsed(void* object, uint32_t materialIndex);
    unsigned char object_getMaterial(void* object, uint32_t materialIndex, const ObjectMaterial* out);
    unsigned char object_getMaterialText(void* object, uint32_t materialIndex, const ObjectMaterialText* out);
    void object_setMaterial(void* object, uint32_t materialIndex, int model, String textureLibrary, String textureName, uint32_t colour);
    void object_setMaterialText(void* object, uint32_t materialIndex, String text, ObjectMaterialSize materialSize, String fontFace, int fontSize, unsigned char bold, uint32_t fontColour, uint32_t backgroundColour, ObjectMaterialTextAlign align);
    void object_attachToPlayer(void* object, void* player, float offsetX, float offsetY, float offsetZ, float rotX, float rotY, float rotZ);
    void object_attachToObject(void* object, void* other, float offsetX, float offsetY, float offsetZ, float rotX, float rotY, float rotZ, unsigned char syncRotation);

    // entity
    void object_setPosition(void* object, float posX, float posY, float posZ);
    Vector3 object_getPosition(void* object);
    void object_setRotation(void* object, float rotX, float rotY, float rotZ);
    Vector3 object_getRotation(void* object);

    // PlayerObject
    void* playerObject_create(void* player, int modelId, float posX, float posY, float posZ, float rotX, float rotY, float rotZ, float drawDistance);
    void playerObject_release(void* object, void* player);
    void* playerObject_getByID(void* player, int id);
    void playerObject_setDrawDistance(void* object, float drawDistance);
    float playerObject_getDrawDistance(void* object);
    void playerObject_setModel(void* object, int model);
    int playerObject_getModel(void* object);
    void playerObject_setCameraCollision(void* object, unsigned char set);
    unsigned char playerObject_getCameraCollision(void* object);
    int playerObject_move(void* object, float posX, float posY, float posZ, float rotX, float rotY, float rotZ, float speed);
    unsigned char playerObject_isMoving(void* object);
    void playerObject_stop(void* object);
    ObjectMoveData playerObject_getMovingData(void* object);
    void playerObject_attachToVehicle(void* object, void* vehicle, float offsetX, float offsetY, float offsetZ, float rotX, float rotY, float rotZ);
    void playerObject_resetAttachment(void* object);
    ObjectAttachmentData playerObject_getAttachmentData(void* object);
    unsigned char playerObject_isMaterialSlotUsed(void* object, uint32_t materialIndex);
    unsigned char playerObject_getMaterial(void* object, uint32_t materialIndex, ObjectMaterial* out);
    unsigned char playerObject_getMaterialText(void* object, uint32_t materialIndex, ObjectMaterialText* out);
    void playerObject_setMaterial(void* object, uint32_t materialIndex, int model, String textureLibrary, String textureName, uint32_t colour);
    void playerObject_setMaterialText(void* object, uint32_t materialIndex, String text, ObjectMaterialSize materialSize, String fontFace, int fontSize, unsigned char bold, uint32_t fontColour, uint32_t backgroundColour, ObjectMaterialTextAlign align);
    void playerObject_attachToPlayer(void* object, void* player, float offsetX, float offsetY, float offsetZ, float rotX, float rotY, float rotZ);
    void playerObject_attachToObject(void* object, void* other, float offsetX, float offsetY, float offsetZ, float rotX, float rotY, float rotZ);

    // entity

    void playerObject_setPosition(void* object, float posX, float posY, float posZ);
    Vector3 playerObject_getPosition(void* object);
    void playerObject_setRotation(void* object, float rotX, float rotY, float rotZ);
    Vector3 playerObject_getRotation(void* object);

#ifdef __cplusplus
}
#endif
