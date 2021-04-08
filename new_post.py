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

## algorithm [🔝](#weekly-{current_id})

## review [🔝](#weekly-{current_id})

## tip [🔝](#weekly-{current_id})

## share [🔝](#weekly-{current_id})

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
    if current_month == files[-1][:6]:
        current_name = f"{current_month}W{int(files[-1][-4])+1}.md"
        next_name = f"{current_month}W{int(files[-1][-4])+2}.md"
    else:
        current_name = f"{files[-1][:6]}W{int(files[-1][-4])+1}.md"
        next_name = f"{files[-1][:6]}W{int(files[-1][-4])+2}.md"
    print(f"{current_name=} {next_name=}")
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
        if x.find("Calendar") > 0:
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

    # 日历样式+列表样式；前者显示时间，后者显示主题
    with open("SUMMARY.md", encoding="utf-8") as f:
        lists = f.readlines()[2:]
    lists[0] = "\n\n## List\n"
    for i, x in enumerate(lists):
        if x.startswith("## "):
            lists[i] = x.replace("## ", "### ")
        elif x.startswith("### "):
            lists[i] = x.replace("### ", "#### ")

    with open("README.md", "w", encoding="utf-8") as f:
        f.write("".join(lines[: start + 4]))
        f.write(table)
        f.write("".join(lists))
        f.write(license)


def update_summary():
    """
    由于 gitbook 展示列表乱序，听说是根据 SUMMARY 来的
    于是在更新 README 的时候顺便更新
    """
    files = os.listdir("weekly")
    files.sort(reverse=True)
    nums = len(files)
    # 按 【年 - 月 - 第 x 期：share 的主题】排列
    month_table = {
        "01": "一",
        "02": "二",
        "03": "三",
        "04": "四",
        "05": "五",
        "06": "六",
        "07": "七",
        "08": "八",
        "09": "九",
        "10": "十",
        "11": "十一",
        "12": "十二",
    }
    current_year = ""
    current_month = ""
    summary_list = ["# SUMMARY\n", "- [关于本书](README.md)\n"]
    for idx, file in enumerate(files):
        with open(f"weekly/{file}", encoding="utf-8") as f:
            lines = f.readlines()
            for x, line in enumerate(lines):
                if line.find("## share") == 0:
                    topic = lines[x + 2]
        # 去掉链接，前面的序号，以及中英的标记
        topic = topic.split("](")[0].split("1. ")[-1].split("（")[0].strip()
        if topic[0] == "[":
            topic = topic[1:]

        if current_year != file[:4]:
            current_year = file[:4]
            tmp = f"\n## {current_year}\n"
            summary_list.append(tmp)

        if current_month != file[4:6]:
            current_month = file[4:6]
            tmp = f"\n### {month_table[current_month]}月\n"
            summary_list.append(tmp)

        tmp = f"- [第 {nums-idx} 期：{topic}](weekly/{file})"
        summary_list.append(tmp)

    summary_list.append("")

    with open("SUMMARY.md", "w", encoding="utf-8") as f:
        f.write("\n".join(summary_list).replace("\n\n\n", "\n\n"))


if __name__ == "__main__":
    update_summary()
    update_readme()
