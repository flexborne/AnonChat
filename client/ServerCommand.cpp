#include "ServerCommand.h"

#include "utils_JSON.h"
#include "User.h"

CmdFindMatch::CmdFindMatch(const User& user) : user{user} {}

auto CmdFindMatch::toJSON() const -> std::string
{
    return createJSON("type", "cmd",
                      "request", "findMatch",
                      "sex", static_cast<size_t>(user.sex),
                      "conversationCategory", static_cast<size_t>(user.conversationCategory),
                            "ageCategory", static_cast<size_t>(user.ageCategory),
                      "sexPreferences", static_cast<size_t>(user.preferences.sex),
                      "agePreferences", static_cast<size_t>(user.preferences.ageFlags));
}

auto CmdFindMatch::type() const -> CommandType
{
    return CommandType::FindMatch;
}
