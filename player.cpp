#include "gomp.h"
#include "player.h"

#ifdef __cplusplus
extern "C"
{
#endif

    int player_getID(void* player)
    {
        return call<int>("player_getID", player);
    }

    String player_getName(void* player)
    {
        return call<String>("player_getName", player);
    }

    int player_setName(void* player, String name)
    {
        return call<int>("player_setName", player, name);
    }

    Vector3 player_getPosition(void* player)
    {
        return call<Vector3>("player_getPosition", player);
    }

    void player_sendClientMessage(void* player, int colour, String message)
    {
        return call<void>("player_sendClientMessage", player, colour, message);
    }

    void* player_getVehicle(void* player)
    {
        return call<void*>("player_getVehicle", player);
    }

#ifdef __cplusplus
}
#endif
