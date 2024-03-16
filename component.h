#ifndef COMPONENT_H
#define COMPONENT_H

#ifdef __cplusplus
extern "C"
{
#endif

void* openLib(const char* path);
void* findFunc(void* handle, const char* name);
void initFuncs(void* handle);

const char* player_getName(void* player);
void player_sendClientMessage(void* player, int colour, const char* message);
void* vehicle_create(int isStatic, int modelId, float x, float y, float z, float angle, int colour1, int colour2, int respawnDelay, int addSiren);
void* pickup_create(int modelId, unsigned char type, float x, float y, float z, unsigned int virtualWorld, int isStatic, void* player);

#ifdef __cplusplus
}
#endif

#endif // COMPONENT_H
