#!/usr/bin/env python3
# -*- coding: utf-8 -*-
# @Date    : 2024-12-05 23:15:45
# @Author  : Lewis Tian (taseikyo@gmail.com)
# @Link    : github.com/taseikyo
# @Version : python3.13

import os
import requests
from bs4 import BeautifulSoup as Soup
import threading

HOST = "https://www.hongyebookzhai.com"
SRC = "text"
THREAD_COUNT = 16
try:
    INDEX = len(os.listdir("text")) + 1
except:
    INDEX = 1
HEADER = {
    "user-agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/131.0.0.0 Safari/537.36 Edg/131.0.0.0"
}

if not os.path.exists(SRC):
    os.mkdir(SRC)


def pre_parse(book_id):
    """
    基于小说的 ID 拿到所有章节的链接
    """
    if not os.path.exists(book_id):
        os.mkdir(book_id)

    url = f"{HOST}/shuzhai/{book_id}"
    r = requests.get(url, headers=HEADER)
    r.encoding = r.apparent_encoding

    soup = Soup(r.text, "html5lib")
    main = soup.find("div", {"id": "list"})

    chapters = main.find_all("dd")
    chapter_map = {}
    for ch in chapters:
        title = ch.a.text
        link = ch.a["href"]
        chapter_id = int(link.split("/")[-1].split(".")[0])
        chapter_map[chapter_id] = title

    return chapter_map


def parse(book_id, chapter_map, thread_count, thread_id):
    """
    对线程数取模，若等于本线程 ID，则爬取此链接
    """
    for chapter_id, title in chapter_map.items():
        if chapter_id % thread_count != thread_id:
            continue

        new_title = sanitize_filename(title)
        filename = f"{book_id}/{chapter_id}_{new_title}.txt"
        if os.path.exists(filename):
            continue

        url = f"{HOST}/shuzhai/{book_id}/{chapter_id}.html"
        r = requests.get(url, headers=HEADER)
        r.encoding = r.apparent_encoding

        soup = Soup(r.text, "html5lib")
        main = soup.find("div", {"id": "content"})
        body = main.text

        with open(filename, "w", encoding="utf-8") as f:
            f.write(body)
        print(url, new_title)


def sanitize_filename(filename):
    """
    替换 Windows 文件名中非法字符为空。
    """
    # 定义非法字符
    illegal_chars = r'<>:"/\|?*'

    # 替换非法字符为空
    sanitized = "".join(char for char in filename if char not in illegal_chars)

    # 去除前后空格和结尾句点
    sanitized = sanitized.strip().rstrip(".")

    return sanitized


def parse_parallel(book_id):
    """
    并发查询该小说所有章节
    """
    threads = []
    map = pre_parse(book_id)
    for thread_id in range(THREAD_COUNT):
        thread = threading.Thread(
            target=parse, args=(book_id, map, THREAD_COUNT, thread_id)
        )
        threads.append(thread)
        thread.start()

    for thread in threads:
        thread.join()


def parse_series(url):
    """
    从当前页面一直往下一章爬
    """
    if len(url) == 0:
        return

    r = requests.get(url, headers=HEADER)
    r.encoding = r.apparent_encoding

    soup = Soup(r.text, "html5lib")

    top = soup.find("div", {"class": "bookname"})
    title = top.h1.text.replace("正文 ", "")
    links = top.find_all("a")
    next_url = ""
    for link in links:
        if link.text == "下一章":
            next_url = link["href"]

    main = soup.find("div", {"id": "content"})
    body = main.text

    # save text
    global INDEX
    new_title = sanitize_filename(title)
    with open(f"{SRC}/{INDEX:04}_{new_title}.txt", "w", encoding="utf-8") as f:
        f.write(body)
    INDEX += 1

    # parse next
    print(title, next_url)

    return parse_series(f"{HOST}/{next_url}")


if __name__ == "__main__":
    # parse_series(f"{HOST}/shuzhai/1/1.html")
    # parse_parallel("1")
    pass
