import QtQuick 2.0
import GoExtensions 1.0

Rectangle {
    id: root

    width: 1260 
    height: 960 
    color: "gray"

    Player {
      id: player
      x: 640
      y: 480
      Item {
        id: playerRect 
        width: 16
        height: 18
        clip: true
        Image {
            id: playerImg 
            x: 0 
            y: -16 
            source: "DawnLike_1/Characters/Player0.png" 
        }
      }
    }
    Image {
        id: playerImg2 
        x: 200 
        y: 200 
        source: "DawnLike_1/Characters/Player0.png" 
    }
    MouseArea {
        anchors.fill: parent
        onClicked: player.handleClick(mouse.x, mouse.y)
    }

}


