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
(def frame (JFrame. "Goop"))

(defn rand-color [] 
  (int (. (new Color (rand-int 255) (rand-int 255) (rand-int 255)) (getRGB))))

(defn mutate-color [old_color] 
  (int (rand-int old_color))) 

(defn image-randomize []
  (int-array (repeatedly (* width height) #(rand-color))))

(defn image-mutate [old]
  (int-array (for [n (range 0 (* width height))] (mutate-color n)) ))

(defn randomize [^BufferedImage image]
  (let [w (.getWidth image)
        h (.getHeight image)
        out (image-mutate (image-randomize))]
    (. image setRGB 0 0 w h out 0 w)))

(defn goop-panel []
  (proxy [JPanel ActionListener KeyListener] []
    (paintComponent [g] 
      (proxy-super paintComponent g)
      (. g drawImage canvas nil nil))
    (actionPerformed [e]
      (randomize canvas) 
      (.repaint this))
    (keyReleased [e])
    (keyTyped [e])
    (getPreferredSize [] (Dimension. width height))))

(def panel (goop-panel))

(defn main []
  (randomize canvas) 
  (doto frame
    (. add panel)
    (. pack)
    (. setVisible true)
    (. setResizable true)
    (. setSize width height)
    (. setDefaultCloseOperation (JFrame/EXIT_ON_CLOSE)))
  (.start (Timer. 0 panel)))

(main)
