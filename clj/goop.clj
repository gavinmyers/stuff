(import javax.swing.JFrame)
(import javax.swing.JPanel)
(import java.awt.event.ActionEvent)
(import java.awt.Dimension)
(import java.awt.Color)
(import javax.swing.Timer)
(import java.awt.event.ActionListener)
(import java.awt.event.KeyListener)
(import java.awt.image.BufferedImage)

(def width 1024)
(def height 768)
(def canvas (BufferedImage. width height BufferedImage/TYPE_INT_RGB))

(defn goop-panel []
  (proxy [JPanel ActionListener KeyListener] []
    (paintComponent [g] 
      (proxy-super paintComponent g)
      (. g drawImage canvas nil nil))
    (actionPerformed [e]
      (doseq [x (range width) y (range height)]
        (.setRGB canvas x y (. (new Color (rand-int 255) (rand-int 255) (rand-int 255)) (getRGB))))
      (.repaint this))
    (keyReleased [e])
    (keyTyped [e])
    (getPreferredSize [] (Dimension. width height))))

 
(defn main []
  (def panel (goop-panel))
  (def timer (Timer. 0 panel))
  (def frame (JFrame. "Goop"))
  (doseq [x (range width) y (range height)]
    (.setRGB canvas x y (. (new Color (rand-int 255) (rand-int 255) (rand-int 255)) (getRGB))))
  (doto frame
    (. add panel)
    (. pack)
    (. setVisible true)
    (. setResizable true)
    (. setSize width height)
    (. setDefaultCloseOperation (JFrame/EXIT_ON_CLOSE)))
  (.start timer))

(main)
