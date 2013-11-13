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

(def img-rock (fn [] 
  [(-> (Toolkit/getDefaultToolkit) (.getImage "img/rock.png"))])) 

(def img-tree (fn [] 
  [(-> (Toolkit/getDefaultToolkit) (.getImage "img/tree-000.png"))
   (-> (Toolkit/getDefaultToolkit) (.getImage "img/tree-001.png"))
   (-> (Toolkit/getDefaultToolkit) (.getImage "img/tree-002.png"))]))

(def img-grass (fn []  
  [(-> (Toolkit/getDefaultToolkit) (.getImage "img/grass-000.png"))
    (-> (Toolkit/getDefaultToolkit) (.getImage "img/grass-001.png"))
    (-> (Toolkit/getDefaultToolkit) (.getImage "img/grass-002.png"))
    (-> (Toolkit/getDefaultToolkit) (.getImage "img/grass-003.png"))
    (-> (Toolkit/getDefaultToolkit) (.getImage "img/grass-004.png"))
    (-> (Toolkit/getDefaultToolkit) (.getImage "img/grass-005.png"))]))

(def img-sprite (fn []
  [(-> (Toolkit/getDefaultToolkit) (.getImage "img/hero.png"))]))

(defn player []
  (rand-nth (img-sprite))) 

(defn grass [] 
  (rand-nth (img-grass))) 

(defn rock [] 
  (rand-nth (img-rock))) 

(defn tree [] 
  (rand-nth (img-tree))) 

(defn grassland [] 
  (map #(%) (for [x (range 0 800 16) y (range 0 800 16)]
             (fn [] {:x x :y y :sprite (grass)}) )))

(defn rocks [] 
  (repeatedly 10 (fn [] {:x (rand-int 800) :y (rand-int 800) :sprite (rock)}) ))

(defn trees [] 
  (repeatedly 100 (fn [] {:x (rand-int 800) :y (rand-int 800) :sprite (tree)}) ))

(defn creatures [] )

(defn -world [] 
  (apply concat (grassland) (rocks) (trees) (creatures))) 

(def world (memoize -world))

(defn draw [px, py, spt, g]
  (let [i  0 
        j  0 
        sw 16 
        sh 16 
        ]
  (. g drawImage 
      spt 
      px py 
      (+ px sw) (+ py sh) 
      (* i sw) (* j sh) 
      (* (+ i 1) sw) (* (+ j 1) sh) nil ))) 

(defn draw-sprite [spt g] 
  (draw (:x spt) (:y spt) (:sprite spt) g))


(defn beat [g] 
  (dorun (map #(draw-sprite % g) (world))) 
  (draw @agent-x @agent-y (player) g))

(defn game-panel []
  (proxy [JPanel ActionListener KeyListener] []
    (paintComponent [g] 
      (proxy-super paintComponent g)
      (beat g))
    (actionPerformed [e]
      (.repaint this))
    (getPreferredSize [] (Dimension. width height))))

(def panel (game-panel))

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
