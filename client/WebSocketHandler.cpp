#include "WebSocketHandler.h"

#include <Poco/Net/WebSocket.h>

#include "utils_JSON.h"
#include "ServerCommand.h"

WebSocketHandler::WebSocketHandler(Poco::Net::WebSocket* ws, QObject *parent)
    : QObject{parent}, ws{ws}
{ }

void WebSocketHandler::sendCommand(const ServerCommand& serverCommand)
{
    const auto message = serverCommand.toJSON();
    ws->sendFrame(message.data(), message.size());
}


void WebSocketHandler::sendMessage(const QString& text)
{
    const auto message = createSimpleChatMessageJSON(text.toStdString());
    ws->sendFrame(message.data(), message.size());
}
