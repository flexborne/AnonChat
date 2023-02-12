import QtQuick 2.15
import QtQuick.Controls 2.15
import QtQuick.Layouts 1.14

ApplicationWindow {
    width: 700
    height: 800


    visible: true

    StackView {
        id: stackView
        initialItem: RandomSearch {
            // Pass stackView to RandomSearch so it can change the current item
            //stackView: stackView
        }
    }
}

