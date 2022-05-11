-- Method 1
SELECT c.name as Customers FROM Customers c
WHERE NOT EXISTS (SELECT 1 FROM Orders o WHERE c.Id = o.CustomerId)

-- Method 2
SELECT name as Customers FROM Customers
WHERE Id NOT IN (SELECT CustomerId FROM Orders)

-- Method 3
SELECT c.Name as Customers FROM Customers c
LEFT JOIN Orders o on c.Id = o.CustomerId
WHERE o.CustomerId is NULL

