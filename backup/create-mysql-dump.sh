#!/bin/sh

#----------------------------------------------------------
# A simple MySQL database backup script.
#----------------------------------------------------------

# set up all the mysqldump variables
FILE=e-work-book.sql.`date +"%Y%m%d"`
DBSERVER=${db_host}
DATABASE="e_work_book"
USER=${MYSQL_USER}
PASS=${MYSQL_PASS}

# in case you run this more than once a day, remove the previous version of the file
unalias rm     2> /dev/null
rm backup/${FILE}     2> /dev/null
rm backup/${FILE}.gz  2> /dev/null

# use this command for a database server on localhost. add other options if need be.
mysqldump --opt --user=${USER} --password=${PASS} ${DATABASE} > backup/${FILE}

# gzip the mysql database dump file
#gzip $FILE

# how the user the result
echo "backup/${FILE}.gz was created:"
ls -l backup/${FILE}
