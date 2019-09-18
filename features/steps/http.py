#!/usr/bin/env python3


from behave import *
from hamcrest import *
import requests
import json


@when('请求 http')
def step_impl(context):
    obj = json.loads(context.text)
    res = requests.get(
        "{}/{}".format(context.config["url"], obj["path"]),
        params=obj["params"]
    )
    print(res.text)
    if "res" not in obj:
        return
    if "status" in obj['res']:
        assert_that(res.status_code, equal_to(obj["res"]["status"]))
    if "json" in obj['res']:
        assert_that(json.loads(res.text), equal_to(obj["res"]["json"]))
    # context.status = res.status_code
    # context.body = res.text
    # if context.status == 200:
    #     context.res = json.loads(res.text)
    # print({
    #     "status": context.status,
    #     "body": context.body,
    #     "cookies": context.cookies,
    # })
