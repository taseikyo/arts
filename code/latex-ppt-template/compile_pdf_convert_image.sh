#!/bin/bash

# 2021-04-10 22:50:30

for file in $(ls *demo.tex)
do
	xelatex.exe $file

	echo ${file%.*}.pdf

	python3.exe <<-EOF
	pdf_path="${file%.*}.pdf"
	img_path="${file%.*}.jpg"
	from pdf2image import convert_from_path

	pages = convert_from_path(pdf_path, 500)
	pages[0].save(img_path, "JPEG")
	print(pdf_path)
	EOF
done

rm -rf *.aux *.log *.nav *.out *.snm *.toc *.pdf

