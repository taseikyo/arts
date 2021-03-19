> @Date    : 2021-03-17 09:51:45
>
> @Author  : Lewis Tian (taseikyo@gmail.com)
>
> @Link    : github.com/taseikyo

# LaTeX 迁移计划

突然发病，想用 LaTeX 重写一下 md，开始以为可以直接用 pandoc 转化，然而发现根本没法用，后来打算用 Python 脚本转化，然后发现并没有那么简单，就解析了几种简单的格式，其他还是得手动修改。

主要使用的模板为 [ElegantLaTeX/ElegantBook](https://github.com/ElegantLaTeX/ElegantBook)，自行加入了引用格式 `myquote`，具体查看 [elegantbook.cls](elegantbook.cls)，完全是仿照 `remark` 格式来的，其实也没做啥，在段落前加了 '>' 符号，然后将颜色改为灰色。

编译的话使用 MiKTeX 自带的 TeXWorks 就行，编译方式选择 XeLaTex+MakeIndex+BibTex 就行了。

项目结构如下所示：

```Bash
|-- docs
|    |-- w1.tex
|    |-- w2.tex
|    |-- w3.tex
|    |-- ...
|-- image
|    |-- 23132915.png
|    |-- cover.jpg
|    |-- ...
|-- output
|    |-- arts.pdf
|-- elegantbook.cls
|-- main.tex
|-- md2tex.py
|-- README.md
```

docs 文件夹下存放的是每个 weekly 的 tex 源文件，output 文件夹存放的是生成的 arts.pdf，main.tex 是主 tex 文件，elegantbook.cls 是模板文件。
