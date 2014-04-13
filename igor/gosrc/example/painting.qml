import QtQuick 2.0
import GoExtensions 1.0

Rectangle {
    id: root

    width: 1260 
    height: 960 
    color: "blue"

    Tile {
    	x: 0; y: 0; width: 12; height: 12

        SequentialAnimation on x {
            loops: Animation.Infinite
            NumberAnimation { from: 0; to: 320; duration: 4000; easing.type: Easing.InOutQuad }
            NumberAnimation { from: 320; to: 0; duration: 4000; easing.type: Easing.InOutQuad }
        }
    }
}
