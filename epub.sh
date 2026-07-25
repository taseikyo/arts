#!/bin/bash
# @Date    : 2020-12-23 15:57:04
# @Author  : Lewis Tian (taseikyo@gmail.com)
# @Link    : github.com/taseikyo

# 2026/07/15 更新，GitHub Actions 中任意步骤失败时立即退出，方便定位错误
set -e

# 将所有 Markdown 生成离线 epub 电子书
# 完成了基本所有连接（锚点）的跳转
# 由于每周的 Weekly 中 ARTS（algorithm、review、tip 和 share）的格式转化导致没有锚点可跳

# 2020/12/30 更新
# 有趣的是当给 pandoc 加上 -f gfm -t epub 参数时，ARTS 锚点正常跳转
# 但是会 title.txt 就会失去作用，有下列警告
# [WARNING] This document format requires a nonempty <title> element.
# 生成的 epub 电子书的标题、作者、License 确实都没了，title.txt 的内容堆在首页

# 2026/07/12 更新，安装 latex 报错，先 update 一下，报错信息贴在下面
# 顺便把 pandoc 更新到最新版本好了
# Not building database; man-db/auto-update is not 'true'.
# Install latex
# E: Failed to fetch https://security.ubuntu.com/ubuntu/pool/main/libi/libinput/libinput-bin_1.25.0-1ubuntu3.4_amd64.deb  404  Not Found [IP: 40.81.13.82 80]
# E: Failed to fetch https://security.ubuntu.com/ubuntu/pool/main/libi/libinput/libinput10_1.25.0-1ubuntu3.4_amd64.deb  404  Not Found [IP: 40.81.13.82 80]
# E: Unable to fetch some archives, maybe run apt-get update or try with --fix-missing?
sudo apt-get update

# 如果 pandoc 没有安装则先安装
if ! type pandoc >/dev/null 2>&1; then
	echo "Install pandoc"
	# sudo apt-get install pandoc -y
	# 2025/3/1 更新，通过 apt 安装的 pandoc 是老版本的（3.1.3）
	# 而使用的模板有个更新需要使用 3.2.1+ 版本的 pandoc，因此需要手动安装
	# https://github.com/Wandmalfarbe/pandoc-latex-template/issues/405
	sudo apt install -y curl
	tag=3.10
	wget https://github.com/jgm/pandoc/releases/download/$tag/pandoc-$tag-1-amd64.deb
	sudo dpkg -i pandoc-$tag-1-amd64.deb >/dev/null
fi

if ! type xelatex >/dev/null 2>&1 || \
   ! kpsewhich sourcesans.sty >/dev/null 2>&1; then

	echo "Install latex"

	# 2026/07/15 更新，GitHub Actions runner 更新后 eisvogel 模板缺少 sourcesans.sty
	# texlive-full 安装结果不可控且隐藏错误，因此明确安装 PDF 生成需要的 latex 依赖

	# 2026/7/23 更新完依然报错，再次让 ChatGPT 帮我修改脚本，增加了 fonts-source-sans-pro
	# 并执行 mktexlsr 刷新 tex

	sudo apt-get update

	sudo apt-get install -y \
	    texlive-xetex \
	    texlive-latex-extra \
	    texlive-fonts-extra \
	    texlive-fonts-recommended \
	    fonts-source-sans-pro \
	    fonts-noto-cjk

	sudo mktexlsr
fi

# 2026/07/15 更新，检查 eisvogel 依赖是否安装成功
if ! kpsewhich sourcesans.sty >/dev/null; then
	echo "Missing sourcesans.sty"
	exit 1
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
# 替换每篇 weekly 中 algorithm review tip share 的跳转
num=0
# back2top='$^{top}$'
for file in `ls build/*`; do
    if [[ num -eq 0 ]]; then
        anchor=""
    else
        anchor="-${num}"
    fi
    sed -i "s/#algorithm-/#algorithm${anchor}/g" $file
    sed -i "s/#review-/#review${anchor}/g" $file
    sed -i "s/#tip-/#tip${anchor}/g" $file
    sed -i "s/#share-/#share${anchor}/g" $file
    # 将返回 README 的 emoji 修改为 top
    # sed -i "s/#algorithm-/#algorithm-top${anchor}/g" $file
    # sed -i "s/#review-/#review-top${anchor}/g" $file
    # sed -i "s/#tip-/#tip-top${anchor}/g" $file
    # sed -i "s/#share-/#share-top${anchor}/g" $file
    # sed -i "s/⬆/${back2top}/g" $file
    # sed -i "s/🔝/${back2top}/g" $file
    ((num++))
done

echo "Generate epub file using pandoc"
pandoc -o arts.epub title.txt README.md build/*.md --epub-cover-image=images/header.png

# 替换添加 pdf 模板的 mate 信息
echo "Add meta data"
metafile=code/meta.txt
curtime=$(date "+%Y-%m-%d")
newstr="date: \"${curtime}\""
line=$(grep -n "date" $metafile | head -n 1 | cut -d":"  -f1)
sed -n "${line}p" $metafile
sed -i "${line}d" $metafile
sed -i "$((line-1))a $newstr" $metafile > /dev/null

# 添加标题信息，由于添加了目录所以另起一页
echo "\\pagebreak" >> $metafile
echo "# ARTS: Algorithm, Review, Tip and Share"  >> $metafile

# 将 meta 信息添加到 README 中
cat $metafile README.md > README.tmp.md
mv README.tmp.md README.md

# 安装楷体
echo "Install font (KaiTi)"
tar xf code/simkai.tar.gz
sudo cp simkai.ttf /usr/share/fonts/
sudo mkfontscale
sudo mkfontdir
sudo fc-cache
# if [ $(fc-list :lang=zh | grep kai | wc -l) -eq 0 ]; then
# 	echo "Install font (KaiTi)"
# 	tar xf code/simkai.tar.gz
# 	sudo cp simkai.ttf /usr/share/fonts/
# 	sudo mkfontscale
# 	sudo mkfontdir
# 	sudo fc-cache
# fi

# 每个章节从新页开始
echo "\\pagebreak" >> README.md
for file in $(ls build/*.md)
do
    echo "\\pagebreak" >> $file
done

# 替换每篇 weekly 中 readme 的跳转
sed -i "s/#calendar/#arts-algorithm-review-tip-and-share/g" `grep -rl "readme" ./build`

# 替换 webp 为 jpg 格式
# !! 注意后面只要写有 webp 格式的图片，应该再保存一份 jpg 格式
sed -i "s/webp)/jpg)/g" `grep -rl "webp)" ./build`

echo "Generate pdf file using pandoc"
# 利用 eisvogel 模板（Wandmalfarbe/pandoc-latex-template）直接生成 PDF
pandoc README.md build/*.md -o arts.pdf --from markdown --template code/eisvogel --listings --pdf-engine=xelatex -V CJKmainfont="KaiTi" -V colorlinks -V urlcolor=NavyBlue --toc

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
git checkout code/meta.txt
