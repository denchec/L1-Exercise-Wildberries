1) Какой самый эффективный способ конкатенации строк?
Ответ:
Сложение строк через знак "+"

2) Что такое интерфейсы, как они применяются в Go?
Ответ:
Оператор interface{} в основном используется пустым.
Используется, когда не понятно какой тип переменных будет передоваться т.к. он может вмещать в себя значения любого типа.

3) Чем отличаются RWMutex от Mutex?
Ответ:
У RWMutex появляется функция RLock и RUnlock, с помощью которых мы так же как и Lock ограничиваем запись в переменную, исключительно для данной горутины, но если другая горутина, так же, использует метод RLock, то она сможет считывать из переменной одновременно с первой горутиной.
Таких образом обеспечивается параллельное чтение объекта несколькими горутинами, что улучшает производительность.

4) Чем отличаются буферизированные и не буферизированные каналы?
Ответ:
Чем отличаются буферизированные и не буферизированные каналы? Небуферизированный канал всегда блокирует горутину, до чтения\записи Буферизированный канал не всегда блокирует горутину. В буферизированный канал можно писать до конца буфера, без блокировки. А блокировка будет вызвана только записью сверх буфера.

5) Какой размер у структуры struct{}{}?
Ответ: 0 бит

6) Есть ли в Go перегрузка методов или операторов?
Ответ: Нету
Диспетчеризация методов упрощается, если не требуется выполнять сопоставление типов. Опыт работы с другими языками показал нам, что наличие различных методов с одинаковыми именами, но разными сигнатурами иногда полезно, но на практике это также может быть запутанным и хрупким. Совпадение только по имени и требование согласованности типов было основным упрощающим решением в системе типов Go.

7) В какой последовательности будут выведены элементы map[int]int?
Ответ:
По величине ключа, от меньшего к большему

10. Что выведет данная программа и почему? 
Ответ:
Вывод: каждый принт выведет по "1"
В памяти компьютера создается 2 переменные "p", одна находится в функции, а другая в main
Когда мы выводим переменную "p" - берем всегда только ту, что находится в main

11) Что выведет данная программа и почему? 
Ответ:
Вывод: 0, 1, 2, 3, 4 - в конкурирующем порядке, а так же ошибку - fatal error: all goroutines are asleep - deadlock!
Надо передать указатель на wg, потому что мы хотим сказать, что мы обращаемся к прошлой wg и не определяем новую
wg можно не передавать в качестве аргумента, горутина и так имеет доступ к нему
Выведет: 0, 1, 2, 3, 4 - в порядке конкуренции, exit - в конце

12) Что выведет данная программа и почему?
Ответ:
Выведет: 0
Почему:
n := 1, n++  работают внутри оператора if, но не взаимодействуют с мэйном

13) Что выведет данная программа и почему?
Ответ:
Выведет: 100, 2, 3, 4, 5
Почему:
Если использовать оператор append и выйти за пределы массива - создасться новый массив и новый срез, который будет указывать на него. При операции v = append(v, b), так и выходит

14) Что выведет данная программа и почему?
Ответ:
Вывод: [b b a][a a]
При операции slice = append(slice, "a") мы выходим за пределы массива и создаем новый массив и срез, указывающий на него.
