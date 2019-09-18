#!/usr/bin/env python3


from behave import *
from hamcrest import *
import requests
import json


@when('请求 http')
def step_impl(context):
    obj = json.loads(context.text)
    context.res = requests.get(
        "{}/{}".format(context.config["url"], obj["path"]),
        params=obj["params"]
    )

@then('检查 http')
def step_impl(context):
    res = context.res
    obj = json.loads(context.text)
    print(res.text)
    if "status" in obj:
        assert_that(res.status_code, equal_to(obj["status"]))
    if "json" in obj:
        assert_that(json.loads(res.text), equal_to(obj["json"]))