> @Date    : 2020-11-01 10:53:07
>
> @Author  : Lewis Tian (taseikyo@gmail.com)
>
> @Link    : github.com/taseikyo

[![](images/header.png "ARTS")](#calendar)

建建删删，一个强迫症的自我修养，真是有毒，希望这是最后一次。

arts 的想法由 [陈浩](https://github.com/haoel) 提出：

- algorithm
- review
- tip
- share

四点分别是*解决一道算法题*，*点评一篇英文文章*，*学习一项技术技巧*，*分享一个观点或思考*。这是个每周计划，也许有时候由于某些原因做到不全部，但是能至少完成一项吧，毕竟比如技术技巧还是比较轻松的。

## Offline Reading

Download [arts.epub](https://github.com/taseikyo/arts/releases)

<details>
<summary></summary>

除了在线的两种方式（GitHub & GitBook），我闲得无聊又写了个脚本，使得可以离线看，主要就是用 Shell 脚本将这些 Markdown 整合生成一个 epub 文档，然后用 Calibre 转成 mobi 格式，发送到 Kindle 看了下效果，还不错。

为了修改其中的跳转链接（锚点）还看了挺多博客的，结果还是得靠自己发散思维来解决，主要用到了 `pandoc`，`grep`，`cut` 和 `sed`，最后跟我说：`sed` 真牛逼！

食用方法（Linux 环境，或者 Windows 下使用 WSL）：

```Bash
tian@ubuntu:/mnt/f/GitHub$ git clone https://github.com/taseikyo/arts.git
tian@ubuntu:/mnt/f/GitHub$ cd arts
tian@ubuntu:/mnt/f/GitHub/arts$ chmod +x ./epub.sh
tian@ubuntu:/mnt/f/GitHub/arts$ ./epub.sh
Generate title.txt
Generate temporary folder
Modify the path of images
Modify the anchor of Weeklys
Modify the anchor of README.md & Weeklys
Generate epub file using pandoc
Remove temporary folder
Remove title.txt
Reset README.md
Updated 1 path from the index
```

于是一个 "arts.epub" 文档就生成了，可能会报下面错误，用 vim 打开，设置文件格式为 unix 类型（`:set ff=unix`），然后运行。

```Bash
tian@ubuntu:/mnt/f/GitHub/arts$ ./epub.sh
./epub.sh: line 6: $'\r': command not found
```

再更新：除了 epub 格式电子书，又用 LaTex 重写了 Makrdown，在 latex/output 文件夹有一份生成的 [pdf](latex/output/arts.pdf)。

再再更新：直接利用 eisvogel 模板（Wandmalfarbe/pandoc-latex-template）直接生成 PDF，集成到 epub.sh 中，由于需要指定中文字体（楷体）又得在 GitHub Action 中，所以花了不少功夫，好在最后搞定了，这下每次 push 新 tag 时，会自动 release 两种电子书。

</details>

## Calendar

Amazing! 26 posts in total. Keep going!

|            :shipit:            |        :jack_o_lantern:        |             :beer:             |           :fish_cake:          |            :octocat:           |
|:------------------------------:|:------------------------------:|:------------------------------:|:------------------------------:|:------------------------------:|
| [202011W1](weekly/202011W1.md) | [202011W2](weekly/202011W2.md) | [202011W3](weekly/202011W3.md) | [202011W4](weekly/202011W4.md) | [202012W1](weekly/202012W1.md) |
| [202012W2](weekly/202012W2.md) | [202012W3](weekly/202012W3.md) | [202012W4](weekly/202012W4.md) | [202012W5](weekly/202012W5.md) | [202101W1](weekly/202101W1.md) |
| [202101W2](weekly/202101W2.md) | [202101W3](weekly/202101W3.md) | [202101W4](weekly/202101W4.md) | [202102W1](weekly/202102W1.md) | [202102W2](weekly/202102W2.md) |
| [202102W3](weekly/202102W3.md) | [202102W4](weekly/202102W4.md) | [202103W1](weekly/202103W1.md) | [202103W2](weekly/202103W2.md) | [202103W3](weekly/202103W3.md) |
| [202103W4](weekly/202103W4.md) | [202103W5](weekly/202103W5.md) | [202104W1](weekly/202104W1.md) | [202104W2](weekly/202104W2.md) | [202104W3](weekly/202104W3.md) |
| [202104W4](weekly/202104W4.md) | | | | | |

## List

### 2021

#### 四月

- [第 26 期：任何人答应你的事都不算数，只有你自己能做主的事才算数。](weekly/202104W4.md)
- [第 25 期：没有消息就是消息](weekly/202104W3.md)
- [第 24 期：真正浪费时间和金钱的方式](weekly/202104W2.md)
- [第 23 期：一辈子单身的话怎么样？](weekly/202104W1.md)

#### 三月

- [第 22 期：先学会做人，再来做事](weekly/202103W5.md)
- [第 21 期：选择困难症是优柔寡断的体现](weekly/202103W4.md)
- [第 20 期：想与命运赛跑，就必须走过生命中最崎岖的路。](weekly/202103W3.md)
- [第 19 期：生活是为你工作，而不是与你为敌](weekly/202103W2.md)
- [第 18 期：你讨厌的工作是最好的老师](weekly/202103W1.md)

#### 二月

- [第 17 期：高校"唯论文"导向的后果](weekly/202102W4.md)
- [第 16 期：你花几十万看虚拟主播真的值得吗？](weekly/202102W3.md)
- [第 15 期：为何年味越来越淡？](weekly/202102W2.md)
- [第 14 期：树大必有枯枝，人多必有白痴](weekly/202102W1.md)

#### 一月

- [第 13 期：形式主义何时休？](weekly/202101W4.md)
- [第 12 期：程序员未来怎么转型？](weekly/202101W3.md)
- [第 11 期：有不同的意见才是真正的尊重](weekly/202101W2.md)
- [第 10 期：你为什么不更新你的个人网站？](weekly/202101W1.md)

### 2020

#### 十二月

- [第 9 期：选择合适的尺度](weekly/202012W5.md)
- [第 8 期：停止追求完美 —— 过你想要的生活](weekly/202012W4.md)
- [第 7 期：我们应该掌握的三门编程语言](weekly/202012W3.md)
- [第 6 期：你也可以做大事，前提是你有时间](weekly/202012W2.md)
- [第 5 期：在你感觉准备好之前采取行动](weekly/202012W1.md)

#### 十一月

- [第 4 期：有时候，精英主义是可以接受的](weekly/202011W4.md)
- [第 3 期：自由并不简单](weekly/202011W3.md)
- [第 2 期：忘记业余项目，专注于工作](weekly/202011W2.md)
- [第 1 期：你的头脑是二值逻辑，还是三值逻辑？](weekly/202011W1.md)


## License

Copyright (c) 2020 Lewis Tian. Licensed under the MIT license.
