-- Method 1
SELECT w1.Id 
FROM Weather w1, Weather w2
WHERE w1.Temperature > w2.Temperature 
AND TO_DAYS(w1.recordDate)-TO_DAYS(w2.recordDate)=1;
-- TO_DAYS(w1.recordDate) return the number of days from year 0 to date recordDate

-- Method 2
SELECT w1.id
FROM Weather AS w1 , Weather AS w2
WHERE w1.Temperature > w2.Temperature 
AND DATEDIFF(w1.recordDate , w2.recordDate) = 1;