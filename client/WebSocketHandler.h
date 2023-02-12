#pragma once

#include <QObject>

namespace Poco::Net {
    class WebSocket;
}

class ServerCommand;


class WebSocketHandler : public QObject
{
    Q_OBJECT
public:
    explicit WebSocketHandler(Poco::Net::WebSocket* ws, QObject* parent = nullptr);

public slots:
    void sendCommand(const ServerCommand& serverCommand);
    void sendMessage(const QString& text);

private:
    Poco::Net::WebSocket* ws;
};
