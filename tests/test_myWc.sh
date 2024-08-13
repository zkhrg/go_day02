#!/bin/bash

# Файлы для хранения вывода команд
my_command_output="my_command_output.txt"
wc_command_output="wc_command_output.txt"

# Запускаем вашу команду и сохраняем её вывод
# Замените "your_command" на вашу команду
../bin/myWc -w ../file.log >> "$my_command_output"
../bin/myWc -l ../file.log >> "$my_command_output"
../bin/myWc -m ../file.log >> "$my_command_output"

# Запускаем команду wc с нужными флагами и сохраняем её вывод
# Например, используем флаги -lwc
wc -w ../file.log >> "$wc_command_output"
wc -l ../file.log >> "$wc_command_output"
wc -m ../file.log >> "$wc_command_output"

# Сортируем строки в обоих файлах (sort) и сохраняем результат в отдельные файлы
sorted_my_command_output="sorted_my_command_output.txt"
sorted_wc_command_output="sorted_wc_command_output.txt"
sort "$my_command_output" > "$sorted_my_command_output"
sort "$wc_command_output" > "$sorted_wc_command_output"

# Сравниваем отсортированные файлы
if diff "$sorted_my_command_output" "$sorted_wc_command_output" > /dev/null; then
    echo "Выводы совпадают (независимо от порядка строк)"
else
    echo "Выводы различаются"
fi

# Удаляем временные файлы
# rm "$my_command_output" "$wc_command_output" "$sorted_my_command_output" "$sorted_wc_command_output"
