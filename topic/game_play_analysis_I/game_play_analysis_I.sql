-- Method 1
SELECT player_id, MIN(event_date) first_login 
FROM Activity
GROUP BY player_id

-- Method 2
SELECT player_id, first_login
FROM (
    SELECT player_id, event_date first_login,
    dense_rank() OVER(
        PARTITION BY player_id
        ORDER BY event_date
    ) poz
    FROM Activity
) src
WHERE poz = 1