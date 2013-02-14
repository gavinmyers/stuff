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
  (println "panel loaded")
  (proxy [JPanel ActionListener KeyListener] []
    (paintComponent [g] 
      (proxy-super paintComponent g)
      (. g drawImage canvas nil nil))
    (actionPerformed [e]
      (println e)
      (.repaint this))
    (getPreferredSize [] (Dimension. width height))))

 
(defn main []
  (def panel (cyan-panel))
  (def timer (Timer. 500 panel))
  (def spriteMap (ImageIO/read (File. "1bitcharanim.png")))
  (def hero (JLabel. (ImageIcon. (.getSubimage spriteMap 8 8 8 8))))
  (doto (JFrame. "Clojure Testing") 
    (.addKeyListener (proxy [java.awt.event.KeyListener] []
      (actionPerformed [e] (println e ) (.repaint this))
      (keyPressed [e] (println (.getKeyChar e) " key pressed") (System/exit 0))
      (keyReleased [e] (println (.getKeyChar e) " key released"))
      (keyTyped [e] (println (.getKeyChar e) " key typed"))))
    (. add panel)
    (. add hero)
    (. pack)
    (. setVisible true)
    (. setResizable true)
    (. setSize width height)
    (. setDefaultCloseOperation (JFrame/EXIT_ON_CLOSE)))
  (.start timer))

(main)
