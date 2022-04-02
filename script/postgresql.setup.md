# PostgreSQL
---

*[How To Install and Use PostgreSQL on Ubuntu 18.04](https://www.digitalocean.com/community/tutorials/how-to-install-and-use-postgresql-on-ubuntu-18-04)*

## Installing PostgreSQL
---

```bash
bash postgresql.install.sh
```

## Using PostgreSQL Roles and Databases
---

### Switching Over to the postgres Account

```bash
sudo -i -u postgres
```
```bash
postgres@controller:~$
```

```bash
psql
```
```bash
postgres@controller:~$ psql
psql (14.2 (Ubuntu 14.2-1.pgdg20.04+1))
Type "help" for help.

postgres=# 
```

```bash
\q
```
```bash
postgres=# \q
postgres@controller:~$ 
```

### Accessing a Postgres Prompt Without Switching Accounts

```bash
sudo -u postgres psql
```
```bash
raktim@controller:~$ sudo -u postgres psql
psql (14.2 (Ubuntu 14.2-1.pgdg20.04+1))
Type "help" for help.

postgres=# 
```

```bash
\q
```
```bash
postgres=# \q
raktim@controller:~$ 
```

## Creating a New Role
---

```bash
raktim@controller:~$ createuser --interactive
Enter name of role to add: postgresrole
Shall the new role be a superuser? (y/n) y
createuser: error: connection to server on socket "/var/run/postgresql/.s.PGSQL.5432" failed: FATAL:  role "raktim" does not exist
```

```bash
postgres@controller:~$ createuser --interactive
Enter name of role to add: raktim
Shall the new role be a superuser? (y/n) y
postgres@controller:~$ 
```

## Basic privileges for Postgres 14 or later
---

```bash
movies=# ALTER USER raktim CREATEDB;
ALTER ROLE
movies=# GRANT pg_read_all_data TO raktim;
GRANT ROLE
movies=# GRANT pg_write_all_data TO raktim;
GRANT ROLE
movies=# \du
```
```bash
                                                 List of roles
 Role name |                         Attributes                         |              Member of               
-----------+------------------------------------------------------------+--------------------------------------
 postgres  | Superuser, Create role, Create DB, Replication, Bypass RLS | {}
 raktim    | Create DB                                                  | {pg_read_all_data,pg_write_all_data}
```

## Creating a New Database
---

- create database
```bash
postgres@controller:~$ createdb movies
```

- check if database has been created
```bash
postgres@controller:~$ psql
postgres=# \l
postgres=# \q
```

- drop database
```bash
postgres@controller:~$ dropdb movies
```

- check if database has been deleted
```bash
postgres@controller:~$ psql
postgres=# \l
postgres=# \q
```

## Opening a Postgres Prompt with the New Role
---

```bash
postgres=# create user raktim with encrypted password '12345678';
CREATE ROLE
postgres=# grant all privileges on database movies to raktim;
GRANT
postgres=# \du
```

## Creating and Deleting Tables
---

```bash
postgres=# \c movies
You are now connected to database "movies" as user "postgres".
movies=# \dt
Did not find any relations.
```

```bash
movies=# CREATE TABLE movies (
id SERIAL,
movieID varchar(50) NOT NULL,
movieName varchar(50) NOT NULL,
PRIMARY KEY (id)
);
CREATE TABLE
movies=# \dt
         List of relations
 Schema |  Name  | Type  |  Owner   
--------+--------+-------+----------
 public | movies | table | postgres
(1 row)

movies=# 
```

## Adding, Querying, and Deleting Data in a Table
---

```bash
movies=# INSERT INTO movies (
movies(# movieID,
movies(# movieName
movies(# )
movies-# VALUES
movies-# ('1', 'movie3'),
movies-# ('2', 'movie2'),
movies-# ('3', 'movie1');
INSERT 0 3
movies=# SELECT * FROM movies;
 id | movieid | moviename 
----+---------+-----------
  1 | 1       | movie3
  2 | 2       | movie2
  3 | 3       | movie1
(3 rows)

movies=# 
```

```bash
movies=# DELETE FROM movies
movies-# WHERE id = 2;
DELETE 1
movies=# SELECT * FROM movies;
 id | movieid | moviename 
----+---------+-----------
  1 | 1       | movie3
  3 | 3       | movie1
(2 rows)

movies=# 
```

## Adding and Deleting Columns from a Table
---

## Updating Data in a Table
---

