(import javax.imageio.ImageIO)
(import java.io.File)
(import javax.swing.JFrame)
(import javax.swing.JLabel)
(import javax.swing.JPanel)
(import javax.swing.ImageIcon)
(import java.awt.event.ActionEvent)
(import java.awt.Dimension)
(import java.awt.Color)
(import javax.swing.Timer)
(import java.awt.image.PixelGrabber)
(import java.awt.event.ActionListener)
(import java.awt.event.KeyListener)
(import java.awt.image.BufferedImage)

(def width 1024)
(def height 768)
(def canvas (BufferedImage. width height BufferedImage/TYPE_INT_RGB))

(defn cyan-panel []
  (proxy [JPanel ActionListener KeyListener] []
    (paintComponent [g] 
      (proxy-super paintComponent g)
      (. g drawImage canvas nil nil))
    (actionPerformed [e]
      (.repaint this))
    (keyReleased [e])
    (keyTyped [e])
    (getPreferredSize [] (Dimension. width height))))

 
(defn main []
  (def panel (cyan-panel))
  (def timer (Timer. 5000 panel))
  (def spriteMap (ImageIO/read (File. "1bitcharanim.png")))
  (def hero (JLabel. (ImageIcon. (.getSubimage spriteMap 8 8 8 8))))
  (def frame (JFrame. "Goop"))
  (doto frame
    (. add panel)
    (. add hero)
    (. pack)
    (. setVisible true)
    (. setResizable true)
    (. setSize width height)
    (. setDefaultCloseOperation (JFrame/EXIT_ON_CLOSE)))
  (.start timer))

(main)
