#pragma once

#include <string>

class User;

enum class CommandType {
    FindMatch
};

class ServerCommand
{
public:
    virtual ~ServerCommand() = default;

    virtual std::string toJSON() const = 0;
    virtual CommandType type() const = 0;
};

class CmdFindMatch : public ServerCommand
{
public:
    CmdFindMatch(const User& user);
    std::string toJSON() const final;
    CommandType type() const final;
private:
    const User& user;
};
