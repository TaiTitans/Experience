

---

## Types of Indexes

- **Clustered Index  
    It determines the physical order of records in the table. A table can have only one Clustered Index, and it is often created on the primary key column.  
    When a table has a clustered index, the data rows are stored on the disk in the same order as the clustered index.


```SQL
CREATE CLUSTERED INDEX index_name  
ON table_name (column1, column2, ...);
```

- **Non-Clustered Index  
    A non-clustered index is a separate data structure from the table and contains a copy of selected columns in a sorted format.  
    It allows for quick access to data based on the indexed columns without affecting the physical order of the data.  
    We can create multiple non-clustered indexes on a table.

```SQL
CREATE NONCLUSTERED INDEX index_name  
ON table_name (column1, column2, ...);
```


![](https://miro.medium.com/v2/resize:fit:828/format:webp/1*tzwl1lQTtkAsUJ9wdbOOrQ.png)


- **Unique Index  
    A unique index enforces the uniqueness of values in the column. Often used on columns representing primary keys or other unique identifiers.  
    It ensures that no two rows in the table can have the same value in the indexed column.

```SQL
CREATE UNIQUE INDEX index_name  
ON table_name (column1, column2, ...);
```

**Filtered Index / Partial Index**
When we create an index on a subset of rows of a table or on rows with specified condition then it is called Filtered Index.  
In some databases, Filtered Index and Partial Index have the same meaning, but some have slightly different meaning.

```SQL
CREATE INDEX index_name  
ON table_name (column1, column2, ...)  
WHERE filter_condition;
```

- **Covering Index  
    When an index includes all the columns required fulfilling a query, called Covering Index.  
    For example, we have created a non-clustered index on 3 columns, and we are retrieving these only column then this non-clustered column called covering index.  
    Covering index eliminates the need to access the table.

