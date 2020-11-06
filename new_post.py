#!/usr/bin/env python3
# -*- coding: utf-8 -*-
# @Date    : 2020-11-06 17:34:32
# @Author  : Lewis Tian (taseikyo@gmail.com)
# @Link    : github.com/taseikyo
# @Version : python3.8

"""
根据模板生成一个 weekly 基本内容
"""

import os
import time

TEMPLATE = """
> @Author  : Lewis Tian (taseikyo@gmail.com)
>
> @Link    : github.com/taseikyo
>
> @Range   : {begin_date} - {begin_date}

# Weekly #{current_id}

[readme](../README.md) | [previous]({previous_post}) | [next]({next_post})

![](../images/1.jpg "Weekly #{current_id}")

\\**Photo by Lewis Tian on Unsplash*

## Table of Contents

- [algorithm](#algorithm-)
- [review](#review-)
- [tip](#tip-)
- [share](#share-)

## algorithm [⬆](#weekly-{current_id})

## review [⬆](#weekly-{current_id})

## tip [⬆](#weekly-{current_id})

## share [⬆](#weekly-{current_id})

[readme](../README.md) | [previous]({previous_post}) | [next]({next_post})

"""


def new_post():
    """
    根据当前已有的 weekly 生成新的
    """
    files = os.listdir("weekly")
    current_id = len(files) + 1
    current_date = time.strftime("%Y-%m-%d", time.localtime())
    current_month = time.strftime("%Y%m", time.localtime())
    # 202011W1.md
    current_name = f"{current_month}W{int(files[-1][-4])+1}.md"
    next_name = f"{current_month}W{int(files[-1][-4])+2}.md"
    with open(f"weekly/{current_name}", "w", encoding="utf-8") as f:
        f.write(
            TEMPLATE.format(
                begin_date=current_date,
                current_id=current_id,
                previous_post=files[-1],
                next_post=next_name,
            )[1:]
        )


if __name__ == "__main__":
    new_post()
