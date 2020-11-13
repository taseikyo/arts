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
import random

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


def has_read_this_url(url: str = ""):
    """
    判断是否已经添加了 @url 对应的博客
    """
    assert url.strip() != ""
    files = [f for f in os.listdir("weekly") if f.endswith("md")]
    for file in files:
        with open(f"weekly/{file}", encoding="utf-8") as f:
            for idx, line in enumerate(f):
                if line.find(url) >= 0:
                    print(f"{url} has been found: {file}:{idx}!")
                    return
    print(f"{url} has not been read!")


def update_readme():
    """
    尝试使用 GitHub Action 自动更新 readme
    """
    with open("README.md", encoding="utf-8") as f:
        lines = f.readlines()

    for i, x in enumerate(lines):
        if x.find("Weeklys") > 0:
            start = i + 2
            break

    # GitHub Action 会出现乱序的情况，很奇怪
    files = os.listdir("weekly")
    files.sort()
    print(files)

    word = random.choice(
        [
            "OK",
            "Wow",
            "Nice",
            "Amazing",
            "Excellent",
            "Well done",
            "Outstanding",
        ]
    )
    lines[start] = f"{word}! {len(files)} posts in total. Keep going!\n"

    # 表格每行 5 个，不足 5 个补空
    cnt = 0
    one_line = []
    all_line = []
    for file in files:
        one_line.append(f" [{file.split('.')[0]}](weekly/{file}) ")
        cnt += 1
        if cnt == 5:
            all_line.append(one_line)
            one_line = []
            cnt = 0

    if 0 < cnt < 5:
        one_line += [" " for _ in range(cnt, 6)]
        all_line.append(one_line)

    table = "\n".join([f"|{'|'.join(x)}|" for x in all_line])

    license = "\n\n## License\n\nCopyright (c) 2020 Lewis Tian. Licensed under the MIT license.\n"

    with open("README.md", "w", encoding="utf-8") as f:
        f.write("".join(lines[: start + 4]))
        f.write(table)
        f.write(license)


if __name__ == "__main__":
    update_readme()
