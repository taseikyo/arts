#!/usr/bin/env python3
# -*- coding: utf-8 -*-
# @Date    : 2020-11-09 21:38:25
# @Author  : Lewis Tian (taseikyo@gmail.com)
# @Link    : github.com/taseikyo
# @Version : Python 3.8

def reverse_words():
    return " ".join(input().split()[::-1])

print(reverse_words())
