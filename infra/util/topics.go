package util

import (
	"fmt"
	"regexp"
	"strings"
)

// visited -> {"event.create.": false, "event.create.slack": false, "event.create.google": false, "event.process.": false}
// [[event.create. event.create.slack event.create.google] [event.process.]]
func GetSubTopicsFromTopics(visited map[string]bool, sep string) [][]string {
	subTopics := make([][]string, 0)

	for topic, vis := range visited {
		if vis {
			continue
		}
		subTopic := make([]string, 0)

		// Determining the top level subtopic
		componentsOfTopics := strings.Split(topic, sep)
		var currentTopic, topLevelTopic string
		for _, comp := range componentsOfTopics {
			if currentTopic == "" {
				currentTopic = fmt.Sprintf("%s.", comp)
			} else {
				currentTopic = fmt.Sprintf("%s%s.", currentTopic, comp)
			}
			if _, ok := visited[currentTopic]; !ok {
				continue
			}

			visited[currentTopic] = true
			topLevelTopic = currentTopic
			subTopic = append(subTopic, currentTopic)
			break
		}

		r := regexp.MustCompile(fmt.Sprintf("^%s", topLevelTopic))

		for top, vis := range visited {
			if vis {
				continue
			}
			if r.MatchString(top) {
				visited[top] = true
				subTopic = append(subTopic, top)
			}
		}

		subTopics = append(subTopics, subTopic)
	}

	return subTopics
}
