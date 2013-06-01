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
(def frame (JFrame. "Goop"))

(def image-start 
  (int-array (repeatedly (* width height) #(rand-int 1000000))))

(defn img-avg [^long i]
  (aset-int image-start i (/ (+ 1
    (if (= true (< (+ i width) size)) 
      (aget ^ints image-start (+ i ^long width))
      (aget ^ints image-start (+ i 0)))
    (aget ^ints image-start (- i 2))
    (aget ^ints image-start (- i 1))
    (aget ^ints image-start (- i 0))
    (aget ^ints image-start (+ i 1)) 
    (aget ^ints image-start (+ i 2))) 6)) )


(defn image-mutate [img]
  (dotimes 
    [n size] 
    (if (= true (< n (- size 5)) (> n 5))
        (img-avg n) 
        0) ))

(defn randomize [^BufferedImage image]
  (let [w (.getWidth image)
        h (.getHeight image)
        out (image-mutate image)]
  (. image setRGB 0 0 w h image-start 0 w)))

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
