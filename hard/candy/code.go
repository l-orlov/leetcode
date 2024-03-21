package candy

import "math"

/*
Алгоритм:
- ищем минимальное число и от него идем заполнять все остальные
- когда встречаем 2 одинаковых числа рядом, то останавливаемся
и идем искать снова минимум на этом оставшемся промежутке. и начинаем заполнять от него.

Если больше 2 одинаковых чисел подряд, то между ними у всех будет 1.

каверзные кейсы:
- все числа одинаковые => всем дадим 1
- много повторяющихся чисел => будем много раз искать минимумы
*/

func candy(ratings []int) int {
	if len(ratings) == 0 {
		return -1
	}

	candies := make([]int, len(ratings))
	candyRecursion(ratings, candies, 0, len(ratings))

	sum := 0
	for i := range candies {
		sum += candies[i]
	}

	return sum
}

func candyRecursion(ratings, candies []int, startIdx, endIdx int) {
	minIdx := -1
	minVal := math.MaxInt
	for i, rating := range ratings[startIdx:endIdx] {
		if rating < minVal {
			minVal = rating
			minIdx = startIdx + i
		}
	}

	candies[minIdx] = 1

	// Идем заполнять вправо от min
	for i := minIdx + 1; i < endIdx; i++ {
		if ratings[i] > ratings[i-1] {
			candies[i] = candies[i-1] + 1
		} else if ratings[i] < ratings[i-1] {
			// Нужно заново искать min, поэтому делаем новый вызов candy
			candyRecursion(ratings, candies, i, endIdx)
			// Доп проверка
			if candies[i-1] <= candies[i] {
				candies[i-1] = candies[i] + 1
			}
			break
		} else {
			// Текущий элемент равен предыдущему.
			// Если следующий элемент равен текущему или текущий последний, то проставить у текущего 1.
			if i+1 < endIdx && ratings[i+1] == ratings[i] || i == endIdx-1 {
				candies[i] = 1
			} else {
				// Иначе нужно заново искать min, поэтому делаем новый вызов candy
				candyRecursion(ratings, candies, i, endIdx)
				break
			}
		}
	}
	// Идем заполнять влево от min
	for i := minIdx - 1; i >= startIdx; i-- {
		if ratings[i] > ratings[i+1] {
			candies[i] = candies[i+1] + 1
		} else if ratings[i] < ratings[i+1] {
			// Нужно заново искать min, поэтому делаем новый вызов candy
			candyRecursion(ratings, candies, startIdx, i+1)
			// Доп проверка
			if candies[i+1] <= candies[i] {
				candies[i+1] = candies[i] + 1
			}
			break
		} else {
			// Текущий элемент равен предыдущему.
			// Если следующий элемент равен текущему или текущий последний, то проставить у текущего 1.
			if i-1 >= startIdx && ratings[i-1] == ratings[i] || i == 0 {
				candies[i] = 1
			} else {
				// Иначе нужно заново искать min, поэтому делаем новый вызов candy
				candyRecursion(ratings, candies, startIdx, i+1)
				break
			}
		}
	}
}
