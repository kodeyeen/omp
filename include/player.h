#include "gomp.h"

#ifdef __cplusplus
extern "C"
{
#endif

int player_getID(void* player);
String player_getName(void* player);
int player_setName(void* player, String name);
Vector3 player_getPosition(void* player);
void player_sendClientMessage(void* player, int colour, String message);
void* player_getVehicle(void* player);

#ifdef __cplusplus
}
#endif
