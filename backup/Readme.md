# A simple script to backup MySQL Database

## Required Environment Variables

- We need to configure the enviroment with the following values.

```shell
${db_host}
${MYSQL_USER}
${MYSQL_PASS}
```

## To run

```cmd
sh create-mysql-dump.sh
```

After running the above command a db dump will be created in the same folder with file name `e-work-book.sql.YYYYMMDD` format.
