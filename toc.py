#!/usr/bin/env python
# -*- coding: utf-8 -*-

import re
import sys
from pathlib import Path

def strip_markdown_links(text: str) -> str:
    """
    去除 Markdown 链接：[文本](链接) → 文本
    """
    return re.sub(r'\[([^\]]+)\]\([^)]+\)', r'\1', text)

def slugify(title: str) -> str:
    """
    GitHub 风格锚点生成器
    """
    title = strip_markdown_links(title)
    slug = title.strip().lower()
    slug = re.sub(r"[^\w\- ]+", "", slug)
    slug = slug.replace(" ", "-")
    return slug

def generate_toc(md_text: str, min_level=1, max_level=6) -> str:
    """
    生成 TOC，并添加 `## Table of Contents` 标题
    """
    toc_lines = ["## Table of Contents"]
    for line in md_text.splitlines():
        match = re.match(r"^(#{1,6})\s+(.*)", line)
        if match:
            level = len(match.group(1))
            raw_title = match.group(2).strip()
            clean_title = strip_markdown_links(raw_title)
            if min_level <= level <= max_level:
                indent = " " * 4 * (level - min_level)
                anchor = slugify(raw_title)
                toc_lines.append(f"{indent}- [{clean_title}](#{anchor})")
    return "\n".join(toc_lines)

def replace_toc_marker(content: str, toc: str) -> str:
    """
    替换 -[toc]/-[TOC] 占位符为生成的 TOC
    """
    pattern = re.compile(r"-\[toc\]", re.IGNORECASE)
    return pattern.sub(toc, content)

def main(md_file: str):
    path = Path(md_file)
    if not path.exists():
        print(f"❌ 文件未找到: {md_file}")
        sys.exit(1)

    content = path.read_text(encoding="utf-8")
    toc = generate_toc(content)

    if not re.search(r"-\[toc\]", content, re.IGNORECASE):
        print("⚠️ 未找到 -[toc] 或 -[TOC] 标记，未作修改。")
        return

    new_content = replace_toc_marker(content, toc)
    path.write_text(new_content, encoding="utf-8")
    print(f"✅ TOC 已成功写入文件：{md_file}")

if __name__ == "__main__":
    if len(sys.argv) < 2:
        print("用法: python generate_toc.py README.md")
    else:
        main(sys.argv[1])
