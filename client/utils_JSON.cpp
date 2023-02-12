#include "utils_JSON.h"

auto createSimpleChatMessageJSON(const std::string& message) -> std::string
{
    return createJSON("type", "chat", "text", message);
}
