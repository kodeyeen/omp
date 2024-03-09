#include <Windows.h>

typedef const char *(*player_getNamePtr)(void *player);
typedef void (*player_sendClientMessagePtr)(void *player, int colour, const char *message);

player_getNamePtr player_getNameFunc;
player_sendClientMessagePtr player_sendClientMessageFunc;

const char *player_getName(void *player) {
    return player_getNameFunc(player);
}

void player_sendClientMessage(void *player, int colour, const char *message) {
    return player_sendClientMessageFunc(player, colour, message);
}

// 

void *loadLib(const char *name) {
    return LoadLibrary((LPCTSTR)name);
}

void unloadLib(void *handle) {
    FreeLibrary((HMODULE)handle);
}

void *findFunc(void *handle, const char *name) {
    return GetProcAddress((HMODULE)handle, (LPCSTR)name);
}

void initFuncs(void *handle) {
    player_getNameFunc = (player_getNamePtr)findFunc(handle, "player_getName");
    player_sendClientMessageFunc = (player_sendClientMessagePtr)findFunc(handle, "player_sendClientMessage");
}
