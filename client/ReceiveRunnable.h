#pragma once

#include <QObject>
#include <Poco/Runnable.h>

namespace Poco::Net {
    class WebSocket;
}


class ReceiveRunnable: public QObject,
                       public Poco::Runnable

{
    Q_OBJECT
public:
    ReceiveRunnable(Poco::Net::WebSocket& ws, QObject* parent = nullptr);

    void run() final;
signals:
    void idGenerated(std::string id);
    void messageReceived(QString senderId, QString message);

private:
    Poco::Net::WebSocket& ws;
};
