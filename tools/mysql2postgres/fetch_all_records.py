# prints all the records of all the tables

import sqlalchemy
import pymysql
import os
#
from sqlalchemy import create_engine
from sqlalchemy.ext.declarative import declarative_base
from sqlalchemy import Column, Integer, String, DateTime, ForeignKey
from sqlalchemy.orm import sessionmaker, relationship

from model import table_list

def main():
    # connect to the mysql database
    engine = create_engine(os.getenv("DATABASE_URL"))

    def convert_to_sql_expression(expr):
        if expr is None:
            return "NULL"
        else:
            expr = str(expr)
            expr = expr.replace("'", "''")
            return "'%s'" % (expr,)


    for t,table_model in table_list.items():
        print("-- " + t)
        #get all the records from each table_model as one dict per record
        records = engine.execute(table_model.select())
        #print the records
        for record in records:
            # generate PostgreSQL Insert statement
            print("INSERT INTO %s (" % (t,), end='')
            column_list = record.keys()
            for column in column_list:
                print(" %s," % (column,), end='')
            print(" ) VALUES (" +",".join(["%s" % (convert_to_sql_expression(field),) for field in record]) + ");")
        
            #for field,value in record.items():
            #    print("%s = %s" % (field,value))

            #print(record)

        


    # disconnect from the sqlalchemy engine
    engine.dispose()

if __name__ == "__main__":
    main()
    
