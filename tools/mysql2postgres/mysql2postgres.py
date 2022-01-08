"""
The purpose of this program is to generate all the
sqlalchemy models required in order to convert a MySQL database
 to a PostgreSQL database.
"""

""" use sqlalchemy to connect to the mysql datbase"""
import sqlalchemy
import pymysql
import os
from sqlalchemy import create_engine
from sqlalchemy.ext.declarative import declarative_base
from sqlalchemy import Column, Integer, String, DateTime, ForeignKey
from sqlalchemy.orm import sessionmaker, relationship



def fix_type(type_):
    """ convert datatype from mysql to sqlalchemy """

    # take the left part of the string before the 'open-paranthesis'
    type_ = type_.split('(')[0]

    if type_ == 'int':
        return 'Integer'
    elif type_ == 'varchar':
        return 'String'
    elif type_ == 'datetime':
        return 'DateTime'
    else:
        return 'String'

def main():
        
    connection = sqlalchemy.create_engine(os.getenv("DATABASE_URL"))

    # get the list of tables from mysql into a list
    tables = connection.execute("SHOW TABLES;")

    tlist = []
    # print the list of tables
    for table in tables:
        print("# ", table)
        table_name = table[0]
        tlist.append(table_name)

        # get the columns from the table
        columns = connection.execute("describe %s;" % table_name)
        #print the column list
        
        
        print("%s = Table(" % (table_name,))
        print("    '%s', meta, " % (table_name,))

        for column in columns:
            column_name = column[0]
            column_type = column[1]
            column_type = fix_type(column_type)
            column_extra = ""
            comment = column[1]
            if column_name == "id":
                column_extra = ", primary_key=True"
            print("    Column('%s', %s %s), #%s" % (column_name, column_type, column_extra, comment))
        print(")")

    print("table_list = { " + ",".join(["'"+t+"':"+t for t in tlist])+"}")
    # disconnect from the db
    connection.dispose()

if __name__ == "__main__":
    main()
    