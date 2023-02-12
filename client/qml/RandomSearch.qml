import QtQuick 2.15
import QtQuick.Controls 2.15
import QtQuick.Layouts 1.14

import enum_ageCategory 1.0
import enum_conversationCategory 1.0
import enum_sex 1.0

import "enumTranslations.js" as EnumTranslations


Item {
    width: 700
    height: 800

    //property var stackView

    id: randomSearchForm

    property int selectedConversationCategory: ConversationCategory.Casual
    property int selectedYourSex: Sex.Undefined
    property int selectedCompanionsSex: Sex.Undefined
    property int selectedYourAge: 1 << 0
    property int selectedCompanionsAge: 1 << 0

    ColumnLayout {
        anchors.horizontalCenter: parent.horizontalCenter

        spacing: 20

        Label {
            text: qsTr("Conversation topic")
            font.bold: true
        }

        RowLayout {
            Layout.alignment: Qt.AlignCenter
            Repeater {
                model: ConversationCategory.Flirtatious + 1
                Button {
                    Layout.fillWidth: true
                    text: EnumTranslations.conversationCatToString(index)
                    checkable: true
                    checked: randomSearchForm.selectedConversationCategory === index
                    onClicked: {
                        randomSearchForm.selectedConversationCategory = index
                    }
                }
            }
        }

        RowLayout {
            Layout.alignment: Qt.AlignCenter

            ColumnLayout {
                Layout.alignment: Qt.AlignLeft
                Label {
                    text: qsTr("Your sex")
                    font.bold: true
                }

                RowLayout {
                    Repeater {
                        model: Sex.Undefined + 1
                        Button {
                            Layout.fillWidth: true
                            text: EnumTranslations.sexToString(index)
                            checkable: true
                            checked: randomSearchForm.selectedYourSex === index
                            onClicked: {
                                randomSearchForm.selectedYourSex = index
                            }
                        }
                    }
                }
            }

            ColumnLayout {
                Layout.alignment: Qt.AlignRight
                Label {
                    text: qsTr("Companion's sex")
                    font.bold: true
                }

                RowLayout {
                    Repeater {
                        model: Sex.Undefined + 1
                        Button {
                            Layout.fillWidth: true
                            text: EnumTranslations.sexToString(index)
                            checkable: true
                            checked: randomSearchForm.selectedCompanionsSex === index
                            onClicked: {
                                randomSearchForm.selectedCompanionsSex = index
                            }
                        }
                    }
                }
            }
        }

        RowLayout {
            Layout.alignment: Qt.AlignCenter

            spacing: 10

            ColumnLayout {
                Layout.alignment: Qt.AlignLeft

                Label {
                    text: qsTr("Your age")
                    font.bold: true
                }

                Repeater {
                    model: 4
                    Button {
                        Layout.fillWidth: true
                        text: EnumTranslations.ageCatToString(1 << index)
                        checkable: true
                        checked: randomSearchForm.selectedYourAge === (1 << index)
                        onClicked: {
                            randomSearchForm.selectedYourAge = (1 << index)
                        }
                    }
                }

            }

            ColumnLayout {
                Layout.alignment: Qt.AlignRight

                Label {
                    text: qsTr("Companion's age")
                    font.bold: true
                }

                Repeater {
                    model: 4
                    Button {
                        Layout.fillWidth: true
                        text: EnumTranslations.ageCatToString(1 << index)
                        checkable: true
                        checked: randomSearchForm.selectedCompanionsAge === (1 << index)
                        onClicked: {
                            randomSearchForm.selectedCompanionsAge = (1 << index)
                        }
                    }
                }
            }
        }

        Button {
            Layout.alignment: Qt.AlignCenter
            text: qsTr("Start Chat")
            font.bold: true
            onClicked: {
                currentUser.set(randomSearchForm.selectedYourSex,
                                randomSearchForm.selectedYourAge,
                                randomSearchForm.selectedConversationCategory,
                                randomSearchForm.selectedCompanionsSex,
                                randomSearchForm.selectedCompanionsAge)
                stackView.push(Qt.resolvedUrl("Chat.qml"))
            }
        }
    }
}
