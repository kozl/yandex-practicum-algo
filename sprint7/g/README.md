# G. Банкомат

[Яндекс.Контест](https://contest.yandex.ru/contest/25596/problems/G/)

Тимофей пошёл снять деньги в банкомат. Ему нужно m франков. В банкомате в бесконечном количестве имеются купюры различных достоинств. Всего различных достоинств n. Купюр каждого достоинства можно взять бесконечно много. Нужно определить число способов, которыми Тимофей сможет набрать нужную сумму.

Пояснения к примерам:

### Пример 1
5 франков можно набрать следующими способами:
1 + 1 + 1 + 1 + 1
1 + 1 + 1 + 2
1 + 1 + 3
1 + 2 + 2
2 + 3

### Пример 2
Во втором примере всего две возможности набрать в сумме 3:
1 + 2
1 + 1 + 1

### Пример 3
Набрать ровно 8 франков купюрами по 5 франков невозможно. Ответ равен нулю.

## Формат ввода
В первой строке записано целое число m — сумма, которую нужно набрать. Во второй строке n — количество монет в банкомате. Оба числа не превосходят 300. Далее в третьей строке записано n уникальных натуральных чисел, каждое в диапазоне от 1 до 1000 –– достоинства купюр.

## Формат вывода
Нужно вывести число способов, которым Тимофей сможет набрать нужную сумму.