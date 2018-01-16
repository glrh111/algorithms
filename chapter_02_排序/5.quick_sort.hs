
quicksort :: [Integer] -> [Integer]
quicksort [] = []
quicksort (x:[]) = [x]
quicksort (x:xs) = quicksort (filter (< x) xs) ++ [x] ++ quicksort (filter (>= x) xs)


main =
--    putStrLn ("hello, world")
    print (quicksort [1,4,3,5,6])
--    putStrLn $ show [1,2,3]