#include "include/console.h"

extern "C" {
    void console_send(String command) {
        return call<void>("console_send", command);
    }
}
