(ns cyan.application 
  (:use [clojure.tools.logging :only (info error)])
  (:use [seesaw.core :as s :only (listen)]))

(import javax.swing.JFrame)
(import javax.swing.JPanel)
(import java.awt.event.ActionEvent)
(import java.awt.Dimension)
(import java.awt.Color)
(import javax.swing.Timer)
(import java.awt.event.ActionListener)
(import java.awt.event.KeyListener)
(import java.awt.image.BufferedImage)
(import java.awt.Toolkit)
(set! *warn-on-reflection* true)

(def width 1024)
(def height 768)
(def size (* width height))
(def canvas (BufferedImage. width height BufferedImage/TYPE_INT_RGB))
(def frame (JFrame. "Another work in progress"))
(def agent-x (atom 150)) 
(def agent-y (atom 150)) 


(def sprite 
  (ref (-> (Toolkit/getDefaultToolkit) (.getImage "img/hero.png"))))
(def tree 
  (ref (-> (Toolkit/getDefaultToolkit) (.getImage "img/tree.png"))))
(def rock 
  (ref (-> (Toolkit/getDefaultToolkit) (.getImage "img/rock.png"))))

(def grass001 
  (ref (-> (Toolkit/getDefaultToolkit) (.getImage "img/grass-001.png"))))
(def grass002
  (ref (-> (Toolkit/getDefaultToolkit) (.getImage "img/grass-002.png"))))
(def grass003 
  (ref (-> (Toolkit/getDefaultToolkit) (.getImage "img/grass-003.png"))))
(def grass004 
  (ref (-> (Toolkit/getDefaultToolkit) (.getImage "img/grass-004.png"))))
(def grass005 
  (ref (-> (Toolkit/getDefaultToolkit) (.getImage "img/grass-005.png"))))

(defn img-rock [] @rock) 
(defn img-tree [] @tree) 
(defn img-sprite [] @sprite) 

(defn img-grass-001 [] @grass001) 
(defn img-grass-002 [] @grass002) 
(defn img-grass-003 [] @grass003) 
(defn img-grass-004 [] @grass004) 
(defn img-grass-005 [] @grass005) 

(defn draw [px, py, spt, g]
  (let [i  0 
        j  0 
        sw 16 
        sh 16 
        ]
  (. g drawImage 
      (spt) 
      px py 
      (+ px sw) (+ py sh) 
      (* i sw) (* j sh) 
      (* (+ i 1) sw) (* (+ j 1) sh) nil ))) 

(defn -scene []
  (repeatedly 100 (fn [] {:x (rand-int 800) :y (rand-int 800)}) ))

(def scene (-scene)) 

(defn goop-panel []
  (proxy [JPanel ActionListener KeyListener] []
    (paintComponent [g] 
      (proxy-super paintComponent g)
      (dorun (for [x (range 0 800 16) y (range 0 800 16)]
           (draw x y img-grass-002 g)))
      (dorun (map #(draw (:x %) (:y %) img-tree g) scene))
      (draw 50 50 img-grass-004 g)
      (draw 100 100 img-rock g)
      (draw 150 100 img-rock g)
      (draw @agent-x @agent-y img-sprite g))
    (actionPerformed [e]
      (.repaint this))
    (getPreferredSize [] (Dimension. width height))))

(def panel (goop-panel))

(def key-actions 
  {:76 #(swap! agent-x (fn [n] (+ n 5)))
   :72 #(swap! agent-x (fn [n] (- n 5)))
   :75 #(swap! agent-y (fn [n] (- n 5)))
   :74 #(swap! agent-y (fn [n] (+ n 5)))
   :32 #(info "SPACE")
   :16 #(info "SHIFT")
   :18 #(info "OPTION")
   :17 #(info "CONTROL")
   :157 #(info "COMMAND")})

(defn key-press [e] 
  (let [keyvar (keyword (str (.getKeyCode e)))
        keyfunc (keyvar key-actions)]
    (if (nil? keyfunc) 
      (info (.getKeyCode e))
      (keyfunc) )))

(defn main []
  (doto frame
    (. add panel)
    (. pack)
    (. setVisible true)
    (. setResizable true)
    (. setSize width height)
    (. setDefaultCloseOperation (JFrame/EXIT_ON_CLOSE)))
  (.start (Timer. 0 panel)))

  (listen frame :key-pressed key-press)
(main)
