#include <Windows.h>

typedef const char *(*player_getNameFunc)(void *player);
typedef void (*player_sendClientMessageFunc)(void *player, int colour, const char *message);

player_getNameFunc player_getNamePtr;
player_sendClientMessageFunc player_sendClientMessagePtr;

const char *player_getName(void *player) {
    return player_getNamePtr(player);
}

void player_sendClientMessage(void *player, int colour, const char *message) {
    return player_sendClientMessagePtr(player, colour, message);
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
    player_getNamePtr = (player_getNameFunc)findFunc(handle, "player_getName");
    player_sendClientMessagePtr = (player_sendClientMessageFunc)findFunc(handle, "player_sendClientMessage");
}
