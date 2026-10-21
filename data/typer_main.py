#!/usr/bin/env python
# -*- coding: utf-8 -*-
# @Date    : 2021-01-25 13:17:37
# @Author  : Lewis Tian (taseikyo@gmail.com)
# @Link    : github.com/taseikyo
# @Version : python3.8

import typer


def main(name: str):
    typer.echo(f"Hello {name}")


if __name__ == "__main__":
    typer.run(main)
