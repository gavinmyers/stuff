(ns cyan.application) 

(import javax.swing.JFrame)
(import javax.swing.JPanel)
(import java.awt.event.ActionEvent)
(import java.awt.Dimension)
(import java.awt.Color)
(import javax.swing.Timer)
(import java.awt.event.ActionListener)
(import java.awt.event.KeyListener)
(import java.awt.image.BufferedImage)
(set! *warn-on-reflection* true)

(def width 1024)
(def height 768)
(def size (* width height))
(def canvas (BufferedImage. width height BufferedImage/TYPE_INT_RGB))
(def frame (JFrame. "Another work in progress"))

(def image-start 
  (int-array (repeatedly (* width height) #(rand-int 10000))))

(def snake-start 
  (let [w (rand-int 1024)
        h (rand-int 768)]
    (int-array (repeatedly (* w h) #(rand-int 0)))))

(defn image-mutate [img]
  img)

(defn snake [] 
 snake-start)

(defn beat [image]
  (let [w (.getWidth image)
        h (.getHeight image)
        out (image-mutate image)
        res (. image setRGB 0 0 w h image-start 0 w)
        sn (. image setRGB (rand-int 100) 10 100 10 (snake) 0 100)]))

(defn goop-panel []
  (proxy [JPanel ActionListener KeyListener] []
    (paintComponent [g] 
      (proxy-super paintComponent g)
      (. g drawImage canvas nil nil))
    (actionPerformed [e]
      (beat canvas) 
      (.repaint this))
    (keyReleased [e])
    (keyTyped [e])
    (getPreferredSize [] (Dimension. width height))))

(def panel (goop-panel))

(defn main []
  (doto frame
    (. add panel)
    (. pack)
    (. setVisible true)
    (. setResizable true)
    (. setSize width height)
    (. setDefaultCloseOperation (JFrame/EXIT_ON_CLOSE)))
  (.start (Timer. 0 panel)))

(main)
