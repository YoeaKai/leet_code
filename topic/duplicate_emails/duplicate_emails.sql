-- Method 1
SELECT Email
FROM Person
GROUP BY Email
HAVING count(*) > 1

-- HAVING
/*
Use it when group by condition.
*/

-- Count duplicate times.
/*
SELECT count(1)
FROM Person
GROUP BY Email
*/

