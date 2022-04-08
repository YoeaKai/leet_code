-- Method 1
SELECT
  score,
  @rank := @rank + (@prev <> (@prev := score)) "rank"
FROM
  Scores,
  (SELECT @rank := 0, @prev := -1) init
ORDER BY score DESC

-- Method 2_1
SELECT
  score,
  (SELECT count(distinct score) FROM Scores WHERE score >= s.Score) "rank"
FROM Scores s
ORDER BY score desc

-- Method 2_2, faster
SELECT
  score,
  (SELECT count(*) FROM (SELECT distinct score s FROM Scores) tmp WHERE s >= score) "rank"
FROM Scores
ORDER BY score desc

-- Method 3, can't use in leetcode
SELECT score, DENSE_RANK() OVER(ORDER BY score DESC) AS "rank" 
FROM Scores
ORDER BY "rank"