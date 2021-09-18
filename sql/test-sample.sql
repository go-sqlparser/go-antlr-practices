SELECT LASTNAME, FIRSTNAME
FROM PERSON.PERSON C
INNER JOIN HUMANRESOURCES.EMPLOYEE E
ON C.BUSINESSENTITYID = E.BUSINESSENTITYID
JOIN SALES.SALESPERSON S 
ON E.BUSINESSENTITYID = S.BUSINESSENTITYID;

/*

-- https://docs.microsoft.com/en-us/sql/relational-databases/performance/subqueries?view=sql-server-ver15

SELECT LastName, FirstName
FROM Person.Person c
INNER JOIN HumanResources.Employee e
ON c.BusinessEntityID = e.BusinessEntityID
JOIN Sales.SalesPerson s 
ON e.BusinessEntityID = s.BusinessEntityID;


SELECT Ord.SalesOrderID, Ord.OrderDate,
    (SELECT MAX(OrdDet.UnitPrice)
     FROM Sales.SalesOrderDetail AS OrdDet
     WHERE Ord.SalesOrderID = OrdDet.SalesOrderID) AS MaxUnitPrice
FROM Sales.SalesOrderHeader AS Ord;

SELECT [Name]
FROM Production.Product
WHERE ListPrice =
    (SELECT ListPrice
     FROM Production.Product
     WHERE [Name] = 'Chainring Bolts' );

SELECT LastName, FirstName
FROM Person.Person
WHERE BusinessEntityID IN
    (SELECT BusinessEntityID
     FROM HumanResources.Employee
     WHERE BusinessEntityID IN
        (SELECT BusinessEntityID
         FROM Sales.SalesPerson)
    );

SELECT DISTINCT c.LastName, c.FirstName, e.BusinessEntityID 
FROM Person.Person AS c JOIN HumanResources.Employee AS e
ON e.BusinessEntityID = c.BusinessEntityID 
WHERE 5000.00 IN
    (SELECT Bonus
    FROM Sales.SalesPerson sp
    WHERE e.BusinessEntityID = sp.BusinessEntityID) ;

-- https://www.sqlservertutorial.net/sql-server-basics/sql-server-subquery/

SELECT
    product_name,
    list_price
FROM
    production.products
WHERE
    list_price > (
        SELECT
            AVG (list_price)
        FROM
            production.products
        WHERE
            brand_id IN (
                SELECT
                    brand_id
                FROM
                    production.brands
                WHERE
                    brand_name = 'Strider'
                OR brand_name = 'Trek'
            )
    )
ORDER BY
    list_price;

-- https://www.mssqltips.com/sqlservertip/6036/sql-server-subquery-example/

SELECT
  cat.ProductCategoryID,
  cat.Name cat_name,
  subcat.Name subcat_name
FROM [AdventureWorks2014].[Production].[ProductCategory] cat
INNER JOIN [AdventureWorks2014].[Production].[ProductSubcategory] subcat
   ON cat.ProductCategoryID = subcat.ProductCategoryID
WHERE cat.ProductCategoryID =
   -- the code in parentheses is the subquery
   ( SELECT cat.ProductCategoryID
     FROM [AdventureWorks2014].[Production].[ProductCategory] cat
     WHERE cat.ProductCategoryID = 1
   )

*/
