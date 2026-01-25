package main

import (
	"slices"
	"strings"
)

func checkBadWord(msg string) string {

	msgList := strings.Split(msg, " ")

	if slices.Contains(msgList, "kerfuffle") || slices.Contains(msgList, "sharbert") || slices.Contains(msgList, "fornax") {
		return cleanBadWord(msgList)
	}

	return msg

}

func cleanBadWord(msg []string) string {

	badWordlist := []string{"kerfuffle", "sharbert", "fornax"}

	for i, item := range msg {

		if slices.Contains(badWordlist, strings.ToLower(item)) {
			msg[i] = "****"
		}

	}

	return strings.Join(msg, " ")
}
