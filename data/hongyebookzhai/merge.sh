#!/bin/bash
# @Date    : 2024-12-11 19:45:30
# @Author  : Lewis Tian (taseikyo@gmail.com)
# @Link    : github.com/taseikyo

# 检查参数
if [ "$#" -lt 2 ]; then
    echo "Usage: $0 <input_directory> <output_file>"
    exit 1
fi

# 获取输入参数
input_directory="$1"
output_file="$2"

# 检查输入目录是否存在
if [ ! -d "$input_directory" ]; then
    echo "Error: Directory '$input_directory' does not exist."
    exit 1
fi

# 如果输出文件已存在，先删除它
[ -f "$output_file" ] && rm "$output_file"

# 遍历指定目录中的所有 .txt 文件
for file in "$input_directory"/*.txt; do
    # 检查是否有匹配的文件
    [ -f "$file" ] || continue
    
    # 提取文件名（去除路径）
    filename=$(basename "$file")
    
    # 去除 index 部分和 .txt 后缀，保留实际标题
    title=$(echo "$filename" | sed -E 's/^[0-9]+_//; s/\.txt$//')
    
    # 添加标题作为分隔符
    echo "$title" >> "$output_file"
    # 添加一个空行作为间隔
    echo "" >> "$output_file"
    
    # 追加文件内容
    cat "$file" >> "$output_file"
    
    # 添加一个空行作为间隔
    echo "" >> "$output_file"
done

echo "合并完成，输出文件为 '$output_file'"
