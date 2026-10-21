#!/usr/bin/env python
# -*- coding: utf-8 -*-
# @Date    : 2021-05-25 12:52:20
# @Author  : Lewis Tian (taseikyo@gmail.com)
# @Link    : github.com/taseikyo

import os
import exifread
import requests


def process_num(x):
    """
    将 [26, 5, 10243/2000] 转化为 26°5'5.125"
    """
    x_last = eval(str(x[-1]))
    new_x = x[0].num + x[1].num / 60 + x_last / 3600

    return "{:.13f}".format(new_x)


def extract_image(pic_path):
    GPS = {}
    date = ""
    with open(pic_path, "rb") as f:
        tags = exifread.process_file(f)
        date = tags.get("Image DateTime", "0").values
        # 纬度标志
        GPS["GPSLatitudeRef"] = tags.get("GPS GPSLatitudeRef", "0").values
        # 纬度
        GPS["GPSLatitude"] = process_num(tags.get("GPS GPSLatitude", "0").values)
        # 经度标志
        GPS["GPSLongitudeRef"] = tags.get("GPS GPSLongitudeRef", "0").values
        # 经度
        GPS["GPSLongitude"] = process_num(tags.get("GPS GPSLongitude", "0").values)

    return {"GPS_information": GPS, "date_information": date}


def find_address_from_bd(GPS):
    secret_key = "wLyevcXk5QY36hTKmvV5350F"
    if not GPS:
        return "该照片无GPS信息"
    lat, lng = (
        GPS["GPSLatitude"],
        GPS["GPSLongitude"],
    )
    baidu_map_api = (
        f"http://api.map.baidu.com/geocoder/v2/?ak={secret_key}&"
        f"callback=&location={lat},{lng}s&output=json&pois=0"
    )
    r = requests.get(baidu_map_api)
    info = r.json()["result"]

    formatted_address = info["formatted_address"]
    province = info["addressComponent"]["province"]
    city = info["addressComponent"]["city"]
    district = info["addressComponent"]["district"]
    location = info["sematic_description"]
    return formatted_address, province, city, district, location


if __name__ == "__main__":
    data = extract_image("a.jpg")
    info = find_address_from_bd(data["GPS_information"])
    print(f"拍摄时间：{data['date_information']}\n照片拍摄地址：{info}")
