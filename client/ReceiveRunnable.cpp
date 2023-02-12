#include "ReceiveRunnable.h"

#include <Poco/Net/WebSocket.h>

#include "utils_JSON.h"
#include <QDebug>


ReceiveRunnable::ReceiveRunnable(Poco::Net::WebSocket& ws, QObject* parent) : QObject{parent}, ws{ws} {}

void ReceiveRunnable::run()
{
    while (true) {
        char buffer[1024];
        int flags;

        if (ws.poll(std::chrono::milliseconds(500), Poco::Net::Socket::SELECT_READ)){
            int received = ws.receiveFrame(buffer, sizeof(buffer), flags);
            auto stringBuffer = std::string(buffer, received);

            std::optional msgType = getValueFromJSON<std::string>(stringBuffer, "type");
            if (!msgType.has_value()) {
                continue;
            }

            if (*msgType == "id") {
                std::optional id = getValueFromJSON<std::string>(stringBuffer, "id");
                if (id.has_value()) {
                    emit idGenerated(std::move(*id));
                }
            }

            if (*msgType == "chat") {
                std::optional senderId = getValueFromJSON<std::string>(stringBuffer, "senderId");
                std::optional text = getValueFromJSON<std::string>(stringBuffer, "text");
                if (senderId.has_value() && text.has_value()) {
                    emit messageReceived(QString::fromStdString(std::move(*senderId)), QString::fromStdString(std::move(*text)));
                }
            }

            if (*msgType == "event") {
                std::optional eventType = getValueFromJSON<std::string>(stringBuffer, "event");
                if (!eventType.has_value()) {
                    continue;
                }
            }
        }
    }
}
