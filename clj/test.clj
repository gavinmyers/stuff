(ns cyan-test
  (:use clojure.test))

(defn something [x] 
  (println x)
  (= true x))

(deftest test-something
  (is (= true (something true)))
  (is (= false (something nil)))
  (is (= false (something false))))

(defn world [] 
  (println "world"))

(deftest test-world
  (world))

(defn wolf []
  (println "you just made a wolf"))

(defn rabbit []
  (println "you just made a wabbit"))

(run-all-tests)
