-- Lib.hs
module Lib where

foreign export ccall add :: Int -> Int -> Int

add :: Int -> Int -> Int
add x y = x + y