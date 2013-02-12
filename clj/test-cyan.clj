(ns cyan-test-2
  (:use clojure.test))

(defn something [x] 
  (= true x))

(deftest test-something
  (is (= false (something false))))
