import QtQuick 2.14
import QtQuick.Controls 2.14
import QtQuick.Layouts 1.14

Item {
    id: chat
    width: 700
    height: 800

    function splitMessage(message, maxLength) {
        var splitMessage = [];
        var currentPart = "";
        var words = message.split(" ");

        for (var i = 0; i < words.length; i++) {
            var word = words[i];
            if (currentPart.length + word.length <= maxLength) {
                currentPart += word + " ";
            } else {
                splitMessage.push(currentPart);
                currentPart = word + " ";
            }
        }

        splitMessage.push(currentPart);
        return splitMessage;
    }


    ColumnLayout {
        anchors.fill: parent

        ListView {
            id: messageList
            Layout.fillWidth: true
            Layout.fillHeight: true
            model: messageModel
            delegate: ItemDelegate {
                id: messageDelegate
                width: childrenRect.width
                height: childrenRect.height

                property bool isFromCurrentUser: currentUser.id === model.senderId

                Component.onCompleted: if (isFromCurrentUser) { anchors.right = parent.right }

                Rectangle {
                    id: wrapper
                    width: childrenRect.width
                    height: childrenRect.height
                    radius: 10
                    color: isFromCurrentUser ? "#6BC1C3" : "#F3E0E0"
                    opacity: 0.9

                    property string message: model.message

                    ColumnLayout {
                        Repeater {
                            id: content
                            model: splitMessage(message, 40)
                            delegate: Text {
                                padding: 5
                                text: modelData
                                font.pointSize: 18
                                color: "black"
                                wrapMode: Text.Wrap
                            }
                        }

                        RowLayout {
                            id: senderLayout
                            Layout.fillWidth: true
                            Layout.alignment: Qt.AlignBottom | (isFromCurrentUser ? Qt.AlignRight : Qt.AlignLeft)
                            Text {
                                id: sender
                                text: model.senderName
                                font.pointSize: 10
                                color: "grey"
                            }
                        }
                    }
                }
            }


            Connections {
                target: receiveRunnable
                function onMessageReceived(senderId, message) {
                    messageModel.append({"senderId": senderId, "senderName": "anon", "message": message});
                }
            }
        }

        RowLayout {
            Layout.fillWidth: true
            Layout.alignment: Qt.AlignBottom

            TextField {
                id: inputField
                Layout.fillWidth: true
                Layout.preferredHeight: 50
                placeholderText: "Type your message here"
                onAccepted: {
                    wsHandler.sendMessage(inputField.text)
                    inputField.text = ""
                }
            }

            Button {
                text: "Send"
                onClicked: {
                    wsHandler.sendMessage(inputField.text)
                    inputField.text = ""
                }
            }
        }
    }

    ListModel {
        id: messageModel
    }
}
