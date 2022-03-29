-- Method 1
SELECT max(salary) as SecondHighestSalary from Employee 
WHERE salary < (SELECT max(Salary) from Employee);

-- Method 2
select IFNULL(
	(SELECT DISTINCT salary from Employee order by salary desc limit 1,1),null
) as SecondHighestSalary
