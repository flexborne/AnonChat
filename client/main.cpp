// qt
#include <QObject>
#include <QQmlApplicationEngine>
#include <QCoreApplication>
#include <QGuiApplication>
#include <QQmlContext>
#include <QDebug>

// poco
#include <Poco/Net/WebSocket.h>
#include <Poco/Net/HTTPClientSession.h>
#include <Poco/Net/HTTPSClientSession.h>
#include <Poco/Net/HTTPRequest.h>
#include <Poco/Net/HTTPResponse.h>
#include <Poco/Net/Context.h>
#include <Poco/Buffer.h>

// std
#include <memory>
#include <iostream>
#include "WebSocketHandler.h"
#include "ReceiveRunnable.h"
#include "User.h"
#include "ServerCommand.h"

#include "enum_AgeCategory.h"
#include "enum_ConversationCategory.h"
#include "enum_Sex.h"

#include "u_defer.h"


int main(int argc, char *argv[]) try
{
    std::unique_ptr<Poco::Net::WebSocket> ws;
    std::unique_ptr<User> currentUser = std::make_unique<User>();

    QGuiApplication app(argc, argv);

    qmlRegisterUncreatableMetaObject(
      AgeCategory::staticMetaObject,
      "enum_ageCategory",
      1, 0,
      "AgeCategory",
      "Error: only enums"
    );

    qmlRegisterUncreatableMetaObject(
      ConversationCategory::staticMetaObject,
      "enum_conversationCategory",
      1, 0,
      "ConversationCategory",
      "Error: only enums"
    );

    qmlRegisterUncreatableMetaObject(
      Sex::staticMetaObject,
      "enum_sex",
      1, 0,
      "Sex",
      "Error: only enums"
    );


    const std::string serverAddress = "localhost";
    const int serverPort = 8080;

    // Create a HTTPClientSession, request and response
    auto session = Poco::Net::HTTPClientSession{serverAddress, serverPort};
    auto request = Poco::Net::HTTPRequest{Poco::Net::HTTPRequest::HTTP_GET, "/ws", Poco::Net::HTTPMessage::HTTP_1_1};
    auto response = Poco::Net::HTTPResponse{};
    ws = std::make_unique<Poco::Net::WebSocket>(session, request, response);
    defer([&ws] { if (ws) ws->close(); });

    WebSocketHandler webSocketHandler(ws.get());

    QObject::connect(currentUser.get(), &User::findMatch, &webSocketHandler, &WebSocketHandler::sendCommand);

    ReceiveRunnable receiveRunnable{std::ref(*ws)};
    Poco::Thread receiveFiber;
    receiveFiber.start(receiveRunnable);

    QObject::connect(&receiveRunnable, &ReceiveRunnable::idGenerated, currentUser.get(), &User::setId);

    QQmlApplicationEngine engine;

    engine.rootContext()->setContextProperty("wsHandler", &webSocketHandler);
    engine.rootContext()->setContextProperty("receiveRunnable", &receiveRunnable);
    engine.rootContext()->setContextProperty("currentUser", currentUser.get());

    const auto url = QUrl{u"qrc:///qml/main.qml"_qs};
        QObject::connect(&engine, &QQmlApplicationEngine::objectCreated,
                     &app, [url](QObject *obj, const QUrl &objUrl) {
        if (!obj && url == objUrl)
            QCoreApplication::exit(-1);
    }, Qt::QueuedConnection);
    engine.load(url);

    if (engine.rootObjects().isEmpty())
        return -1;


    return app.exec();
} catch (std::exception &e) {
    qDebug() << "Exception " << e.what();
}
