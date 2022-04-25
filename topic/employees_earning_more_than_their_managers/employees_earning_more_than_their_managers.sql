-- Method 1
SELECT e1.name as Employee
FROM Employee as e1, Employee as e2 
WHERE e1.managerId is not NULL and e1.managerId = e2.id and e1.salary > e2.salary

-- Method 2
SELECT e1.name as Employee
FROM Employee as e1 inner join Employee as e2 on e1.managerId = e2.id
WHERE e1.salary > e2.salary

-- Method 3
SELECT e1.name as Employee
FROM Employee as e1
WHERE (SELECT salary FROM Employee as e2 WHERE e1.managerId = e2.id) < e1.salary

