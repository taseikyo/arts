#!/bin/bash
# @Date    : 2020-12-23 15:57:04
# @Author  : Lewis Tian (taseikyo@gmail.com)
# @Link    : github.com/taseikyo
# @Version : python3.8

# 将所有 Markdown 生成离线 epub 电子书
# 完成了基本所有连接（锚点）的跳转
# 由于每周的 Weekly 中 ARTS（algorithm、review、tip 和 share）的格式转化导致没有锚点可跳

# 如果 pandoc 没有安装则先安装
if ! type pandoc >/dev/null 2>&1; then
	echo "Install pandoc"
	sudo apt install pandoc -y
fi

echo "Generate title.txt"
echo -e "---\ntitle: \"arts: algorithm, review, tip and share\"\n\
author: 九卿（Lewis Tian）\nlanguage: zh-CN\nrights: MIT License\nindent: true\nlinestretch: 1.25\n\
abstract: |\n\
    所谓 arts，即 algorithm, review, tip and share 四个单词首字母合集\n\
    algorithm 是指每周至少一道算法题，目前主要是 leetcode 和牛客的题目；\n\
    review 是指阅读一篇英文文章，或可再加以点评\n\
    tip 则是学到的技术技巧\n\
    share 则是分享自己的观点和思考，算是比较有看点的部分\n\
    这是一个每周的任务，按理说并不难完成，由于之前曾做过后面由于时间原因删库跑路了\n\
    为了避免类似事情发生，决定只要完成至少一件即可。\n\
    链接：https://github.com/taseikyo/arts\n..." > title.txt

echo "Generate temporary folder"
if [ -d "build" ]; then
	rm -rf build
fi
cp -r weekly build

echo "Delete head comment (Date, Author and Link)"
sed -i "1,6d" README.md
sed -i "1,6d" `grep -rl "Author" ./build`

echo "Modify the path of images"
sed -i "s/..\/images/images/g" `grep -rl "images" ./build`

echo "Modify the anchor of Weeklys"
# 替换每篇 weekly 中 readme 的跳转
sed -i "s/..\/README.md/#calendar/g" `grep -rl "readme" ./build`

echo "Modify the anchor of README.md & Weeklys"
# 替换 README 中 Calendar/List 的跳转
# 替换 Weekly 中 preview/next 的跳转
num=1
while :
do
	match=`grep -E "第 ${num} 期" README.md`
	if [[ ! $match ]]; then
		break
	fi
	weeklyno=`echo $match | cut -d" " -f 3`
	# weekly/202011W1.md
	raw_weekly=`echo $match | cut -d "(" -f 2 | cut -d ")" -f 1`
	# weekly-8
	new_weekly="#weekly-${weeklyno}"
	# 需要用 "|" 作为分隔符，因为 raw_weekly 中含有 "/"
	sed -i "s|${raw_weekly}|${new_weekly}|g" README.md
	# 修改 weekly 中的跳转
	# 202011W1.md
	raw_weekly=`echo $raw_weekly | cut -d "/" -f 2`
	sed -i "s|${raw_weekly}|${new_weekly}|g" `grep -rl $raw_weekly ./build`
	((num++))
done

echo "Generate epub file using pandoc"
pandoc -o arts.epub title.txt README.md build/*.md --epub-cover-image=images/header.png

if [ -d "build" ]; then
	echo "Remove temporary folder"
	rm -rf build
fi

if [ -f "title.txt" ]; then
	echo "Remove title.txt"
	rm -rf title.txt
fi

echo "Reset README.md"
git checkout README.md
