#!/usr/bin/env python3


from behave import *
from hamcrest import *
import requests
import json


@when('http 请求 {method:str} {path:str}')
def step_impl(context, method, path):
    if context.text:
        obj = json.loads(context.text)
    else:
        obj = {}
    if "params" not in obj:
        obj["params"] = {}
    if "json" not in obj:
        obj["json"] = {}
    if method == "GET":
        context.res = requests.get(
            "{}{}".format(context.config["url"], path),
            params=obj["params"], json=obj["json"]
        )
    if method == "PUT":
        context.res = requests.put(
            "{}{}".format(context.config["url"], path),
            params=obj["params"], json=obj["json"]
        )
    if method == "POST":
        if "file" in obj:
            context.res = requests.post(
                "{}{}".format(context.config["url"], path),
                params=obj["params"], json=obj["json"],
                files={
                    'file': open(obj["file"], 'rb')
                }
            )
        else:
            context.res = requests.post(
                "{}{}".format(context.config["url"], path),
                params=obj["params"], json=obj["json"]
            )


@then('http 检查 {status:int}')
def step_impl(context, status):
    res = context.res
    if context.text:
        obj = json.loads(context.text)
    else:
        obj = {}
    assert_that(res.status_code, equal_to(status))

    if "json" in obj:
        result = json.loads(res.text)
        print(type(obj["json"]))
        if isinstance(obj["json"], dict):
            for key in obj["json"]:
                assert_that(result[key], equal_to(obj["json"][key]))
        elif isinstance(obj["json"], list):
            for idx, val in enumerate(obj["json"]):
                assert_that(result[idx], equal_to(val))
        else:
            assert_that(result, equal_to(obj["json"]))

    if "text" in obj:
        assert_that(res.text, equal_to(obj["text"].strip()))

    if "cookies" in obj:
        for key in obj["cookies"]:
            cookies = res.cookies
            if obj["cookies"][key] == "exist":
                assert_that(cookies[key], is_not(None))
            else:
                assert_that(cookies[key], equal_to(obj["cookies"][key]))
