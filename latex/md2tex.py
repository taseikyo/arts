#!/usr/bin/python3
# -*- coding: utf-8 -*-
# @Date    : 2021-03-17 15:43:52
# @Author  : Lewis Tian (taseikyo@gmail.com)
# @Link    : github.com/taseikyo
# @Version : python3.8

"""
尽量将 md 转化为 latex 格式，减少手动修改的工作量
"""


import os

IMG_TEMPLATE = """
\\begin{{figure}}[htbp]
  \\centering
    \\includegraphics[width=\\textwidth]{{{src}}}
  \\caption{{\\textit{{{title}}}}}
\\end{{figure}}
"""


def get_topics(readme_path: str = "../README.md") -> list:
    """
    从 README.md 中获取每周的 topic

    @readme_path: README 相对于本脚本的路径
    """
    start_read = False
    topics = []
    with open(readme_path, encoding="utf-8") as f:
        for line in f:
            if start_read:
                if line[:3] == "- [":
                    topic = line.split("：")[-1].split("](")[0]
                    topics.append(topic)
            elif line[:-1] == "## List":
                start_read = True
    return topics[::-1]


def format_image(line: str, default_title: str = "") -> str:
    """
    格式化图片
    """
    if line.find(" ") > 0:
        src = line.split("](")[-1].split(" ")[0]
        title = line.split(" \"")[-1][:-3]
    else:
        src = line.split("](")[-1][:-2]
        title = src.split("/")[-1]
    if default_title:
        if default_title.find("](") > 0:
            foo = default_title.split("](")
            name = foo[0].split("[")[-1]
            site = foo[1].split("[")[-1]
            title = f"Photo by {name} on {site}"
        else:
            title = default_title[3:-2]
    return IMG_TEMPLATE.format(src=src, title=title)


def convert(topics: list, path: str = "../weekly") -> None:
    """
    将 MD 转化为 LATEX 格式
    """
    for weekly_id, file in enumerate(os.listdir(path)):
        format_lines = []
        if weekly_id == len(topics):
            break
        if os.path.exists(f"w{weekly_id+1}.tex"):
            continue
        # 章节名（每个 WEEKLY 的 topic）
        chapter_name = f"\\chapter{{{topics[weekly_id]}}}\\label{{chap:w{weekly_id+1}}}"
        with open(f"{path}/{file}", encoding="utf-8") as f:
            lines = f.readlines()
        # 字数信息
        total_numbers = lines[10]
        # 头图信息 & 来源
        head_image = lines[12]
        head_image_source = lines[14]
        head_image_txt = format_image(head_image, head_image_source)
        total_numbers_txt = f"\\textit{{{total_numbers[:-1]}}}"
        format_lines.append(chapter_name)
        format_lines.append(head_image_txt)
        format_lines.append(total_numbers_txt)
        # 找到 algorithm
        idx = 15
        while not lines[idx].startswith("## algorithm"):
            idx += 1
        while idx < len(lines):
            line = lines[idx]
            # 大节（algorithm，review，tip，share）
            if line[:3] == "## ":
                name = line.split(" [")[0][3:]
                section = f"\\section{{{name} \\hyperref[chap:w{weekly_id+1}]{{top}}}}\\label{{w{weekly_id+1}:{name}}}"
                txt = section
            # 小节（每一个 algorithm）
            elif line[:4] == "### ":
                name = line.split(" [")[-1].split("](")[0]
                link = line.split("](")[-1][:-2]
                # 没有链接
                if name[:4] == "### ":
                    name = line.split(". ")[-1][:-1]
                    subsection = f"\\subsection{{{name}}}"
                else:
                    subsection = f"\\subsection{{\\href{{{link}}}{{{name}}}}}"
                txt = subsection
            # 代码，内联代码用 \lstinline{print(hello)}
            elif line[:3] == "```":
                language = line[3:-1]
                if not language:
                    language = "Bash"
                idx += 1
                start_idx = idx
                while lines[idx][:3] != "```":
                    idx += 1
                code = ''.join([line for line in lines[start_idx: idx]])
                txt = f"\\begin{{lstlisting}}[language={language}]\n{code}\\end{{lstlisting}}"
            # 斜体
            # ！！！有问题是对于句子中含有斜体的，不太好修改，只能手动修改
            elif line[0] == "*" and line[-2:] == "*\n":
                txt = f"\\textit{{{line[1:-2]}}}"
            # 图片
            elif line[:2] == "![":
                txt = format_image(line)
                # print(txt)
            # 引用，使用我自己加的 \myquote 命令，见 elegantbook.cls
            elif line[:2] == "> " and line[2] != '[':
                txt = f"\\begin{{myquote}}\n{line[2:]}\\end{{myquote}}"
            # 无序列表
            elif line[:2] == "- ":
                start_idx = idx
                while lines[idx][:2] == "- ":
                    idx += 1
                code = ''.join(['\\item ' + line[2:]
                                for line in lines[start_idx: idx]])
                txt = f"\\begin{{itemize}}\n{code}\\end{{itemize}}"
                idx -= 1
                # print(txt)
            else:
                txt = ""
                pass
            if txt:
                format_lines.append(txt if txt[-1] != '\n' else txt[:-1])
            idx += 1

        with open(f"w{weekly_id+1}.tex", "w", encoding="utf-8") as f:
            f.write("\n\n".join(format_lines))


if __name__ == '__main__':
    convert(topics=get_topics())
