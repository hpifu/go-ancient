Feature: authors 测试

    Scenario: authors
        Given mysql 执行
            """
            DELETE FROM ancients WHERE id IN (1,2)
            """
        Given mysql 执行
            """
            INSERT INTO ancients (id, title, author, dynasty, content)
            VALUES (1, "test 静夜思", "李白", "唐", "床前明月光，疑是地上霜。举头望明月，低头思故乡。")
            """
        Given mysql 执行
            """
            INSERT INTO ancients (id, title, author, dynasty, content)
            VALUES (2, "test 绝句", "杜甫", "唐", "两个黄鹂鸣翠柳，一行白鹭上青天。 窗含西岭千秋雪，门泊东吴万里船。")
            """
        When http 请求 GET /author
            """
            {
                "params": {
                    "offset": 0,
                    "limit": 2
                }
            }
            """
        Then http 检查 200
            """
            {
                "json": [
                    "李白",
                    "杜甫"
                ]
            }
            """
        Given mysql 执行
            """
            DELETE FROM ancients WHERE id IN (1,2)
            """