#!/usr/bin/env python3

from behave import *
from hamcrest import *


@given('执行 sql')
def step_impl(context):
    print(context.text)
    with context.mysql_conn.cursor() as cursor:
        cursor.execute(context.text)
    context.mysql_conn.commit()
