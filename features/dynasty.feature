Feature: dynasty 测试

    Scenario: dynasty
        Given 执行 sql
            """
            DELETE FROM ancients WHERE id IN (1,2)
            """
        Given 执行 sql
            """
            INSERT INTO ancients (id, title, author, dynasty, content)
            VALUES (1, "test 静夜思", "李白", "唐", "床前明月光，疑是地上霜。举头望明月，低头思故乡。")
            """
        Given 执行 sql
            """
            INSERT INTO ancients (id, title, author, dynasty, content)
            VALUES (2, "test 绝句", "杜甫", "唐", "两个黄鹂鸣翠柳，一行白鹭上青天。 窗含西岭千秋雪，门泊东吴万里船。")
            """
        When 请求 http
            """
            {
                "method": "get",
                "path": "/dynasty/唐",
                "params": {
                    "offset": 0,
                    "limit": 2
                }
            }
            """
        Then 检查 http
            """
            {
                "status": 200,
                "json": [{
                    "id": 1,
                    "title": "test 静夜思",
                    "author": "李白",
                    "dynasty": "唐"
                }, {
                    "id": 2,
                    "title": "test 绝句",
                    "author": "杜甫",
                    "dynasty": "唐"
                }]
            }
            """
        Given 执行 sql
            """
            DELETE FROM ancients WHERE id IN (1,2)
            """