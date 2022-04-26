-- Method 1
SELECT D.Name AS Department, E.Name AS Employee, E.Salary 
FROM 
	Employee E,
	Department D   
WHERE E.DepartmentId = D.Id 
  AND NOT EXISTS 
  (SELECT 1 FROM Employee B WHERE B.Salary > E.Salary AND E.DepartmentId = B.DepartmentId) 

-- Method 2
SELECT D.Name AS Department ,E.Name AS Employee ,E.Salary 
FROM
	Employee E,
	(SELECT DepartmentId,max(Salary) as max FROM Employee GROUP BY DepartmentId) T,
	Department D
WHERE E.DepartmentId = T.DepartmentId 
  AND E.Salary = T.max
  AND E.DepartmentId = D.id

-- Method 3
SELECT D.Name AS Department ,E.Name AS Employee ,E.Salary 
from 
	Employee E,
	Department D 
WHERE E.DepartmentId = D.id 
  AND (DepartmentId,Salary) in 
  (SELECT DepartmentId, max(Salary) as max FROM Employee GROUP BY DepartmentId)