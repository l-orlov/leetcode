package main

import "fmt"

// docs: https://yuminlee2.medium.com/manachers-algorithm-longest-palindromic-substring-cc97fa615f8b

func LongestPalindrome(str string) string {
	newStr := insertBogusChars(str, '|')
	palRad := make([]int, len(newStr))
	palCenter := 0

	center := 0
	radius := 0
	for center < len(newStr) {
		expanedRadius := radius + 1
		for center-expanedRadius >= 0 && center+expanedRadius < len(newStr) && newStr[center-expanedRadius] == newStr[center+expanedRadius] {
			radius = expanedRadius
			expanedRadius = radius + 1
		}

		palRad[center] = radius
		if palRad[center] > palRad[palCenter] {
			palCenter = center
		}

		oldCenter := center
		oldRadius := radius
		oldCenterRightBorder := oldCenter + oldRadius
		center += 1
		radius = 0
		for center <= oldCenterRightBorder {
			mirroredCenter := 2*oldCenter - center
			maxMirroredRadius := oldCenterRightBorder - center
			if palRad[mirroredCenter] < maxMirroredRadius {
				palRad[center] = palRad[mirroredCenter]
				center += 1
			} else if palRad[mirroredCenter] > maxMirroredRadius {
				palRad[center] = maxMirroredRadius
				center += 1
			} else {
				radius = maxMirroredRadius
				break
			}
		}
	}

	lognestPalindrome := newStr[palCenter-palRad[palCenter] : palCenter+palRad[palCenter]+1]

	return trimBogusChars(lognestPalindrome, '|')
}

func insertBogusChars(s string, bogusChar rune) string {
	newStr := make([]rune, 2*len(s)+1)
	newStr[0] = bogusChar
	i := 1
	for _, char := range s {
		newStr[i] = char
		newStr[i+1] = bogusChar
		i += 2
	}
	return string(newStr)
}

func trimBogusChars(s string, bogusChar rune) string {
	trimmedStr := []rune{}
	for _, char := range s {
		if char != bogusChar {
			trimmedStr = append(trimmedStr, char)
		}
	}
	return string(trimmedStr)
}

func main() {
	fmt.Println(LongestPalindrome("hellosannasmith"))
}
