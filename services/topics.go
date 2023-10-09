package services

import "github.com/rajatxs/go-fconsole/types"

var topics types.Topics = map[string]types.Topic{
	"lifestyle": {
		Name:      "Lifestyle",
		ThumbId:   "c25decb56feccf32b56f3bbce1d0649e",
		ThumbPath: "fivemin-prod/topic-thumb/d6kqr2vxy7siuchz3ugj",
		Public:    false,
	},
	"food-and-cooking": {
		Name:      "Food and Cooking",
		ThumbId:   "a38d1f899a322069b57f93a3a2aa1a0b",
		ThumbPath: "fivemin-prod/topic-thumb/ufafc9pcz80k9obim4ly",
		Public:    false,
	},
	"technology": {
		Name:      "Technology",
		ThumbId:   "d7032d16343fe1a6330db626bf5feeaf",
		ThumbPath: "fivemin-prod/topic-thumb/n6rztxn2ujb5b1ccwhbn",
		Public:    true,
	},
	"finance": {
		Name:      "Finance",
		ThumbId:   "b386564e5f7bb336c6cda43377f8f21d",
		ThumbPath: "fivemin-prod/topic-thumb/ewjibbyowrykv0zjuqq7",
		Public:    true,
	},
	"parenting": {
		Name:      "Parenting",
		ThumbId:   "122e62251f9fb3d3e90091a47af983a8",
		ThumbPath: "fivemin-prod/topic-thumb/xkn9wwkwgpdybn0ynqds",
		Public:    false,
	},
	"sports": {
		Name:      "Sports",
		ThumbId:   "4710819168f02f59326504082b132ba7",
		ThumbPath: "fivemin-prod/topic-thumb/cheancimace34cd4jmib",
		Public:    false,
	},
	"beauty-and-skincare": {
		Name:      "Beauty and Skincare",
		ThumbId:   "391d0fbf180c1864e125ae2de2056f49",
		ThumbPath: "fivemin-prod/topic-thumb/vpfyyuihzwn1xsrcmu9x",
		Public:    false,
	},
	"home-improvement": {
		Name:      "Home Improvement",
		ThumbId:   "545a8568b81b52b21d323e54139f060a",
		ThumbPath: "fivemin-prod/topic-thumb/cxcp7mxut7g9qqdemxpp",
		Public:    false,
	},
	"education": {
		Name:      "Education",
		ThumbId:   "ea4a1994e29888ee339a83ba875b2498",
		ThumbPath: "fivemin-prod/topic-thumb/ezuokyiqie8apwdpunli",
		Public:    false,
	},
	"entertainment": {
		Name:      "Entertainment",
		ThumbId:   "5455004f1dc1f498d5a4c8ee9f00ecf9",
		ThumbPath: "fivemin-prod/topic-thumb/vubqex8sbphmqhgxyeb3",
		Public:    false,
	},
	"business": {
		Name:      "Business",
		ThumbId:   "62180034fffab4bd0b9e2e11c0418876",
		ThumbPath: "fivemin-prod/topic-thumb/nxzh4ihpddkqm0nputm4",
		Public:    true,
	},
	"travel": {
		Name:      "Travel",
		ThumbId:   "9406a84d081a8e72bfba2451c6900872",
		ThumbPath: "fivemin-prod/topic-thumb/rjpwglx0iavcdhukkhuy",
		Public:    false,
	},
	"health": {
		Name:      "Health",
		ThumbId:   "7124b21e721cc4477af37edd67ac2549",
		ThumbPath: "fivemin-prod/topic-thumb/cdagmwqtctswknvqzkxu",
		Public:    false,
	},
	"social": {
		Name:      "Social",
		ThumbId:   "9fb66dab59f3e0ff79ca57f3c7e1fbd1",
		ThumbPath: "fivemin-prod/topic-thumb/ti3kpjhcd155t4h5toom",
		Public:    false,
	},
	"relationships": {
		Name:      "Relationships",
		ThumbId:   "36dc7b9c0f4aaf4f56dd0165b7859d16",
		ThumbPath: "fivemin-prod/topic-thumb/zcdz5monp2zqeipuqzdu",
		Public:    false,
	},
	"science": {
		Name:      "Science",
		ThumbId:   "f41cd486d292306f036ab3b54cb9749d",
		ThumbPath: "fivemin-prod/topic-thumb/xtcbet9ywsfif7uiiayu",
		Public:    true,
	},
	"programming": {
		Name:      "Programming",
		ThumbId:   "e7b77657550f169f0a6d6f679283b222",
		ThumbPath: "fivemin-prod/topic-thumb/zmlt6ft5tyivzcfeejci",
		Public:    true,
	},
}

// GetAllTopics returns object of all available topics
func GetAllTopics() *types.Topics {
	return &topics
}

// GetPublicTopics returns object of publicly available topics
func GetPublicTopics() (pubtops types.Topics) {
	pubtops = make(map[string]types.Topic)

	for topicId, topic := range topics {
		if topic.Public {
			pubtops[topicId] = topic
		}
	}

	return pubtops
}

// GetPrivateTopics returns object of private topics
func GetPrivateTopics() (pubtops types.Topics) {
	pubtops = make(map[string]types.Topic)

	for topicId, topic := range topics {
		if !topic.Public {
			pubtops[topicId] = topic
		}
	}

	return pubtops
}

// GetTopicById returns single topic object by given id
func GetTopicById(id string) types.Topic {
	return topics[id]
}

// GetTopicNameById returns topic name of by given id
func GetTopicNameById(id string) string {
	if topic, ok := topics[id]; ok {
		return topic.Name
	} else {
		return ""
	}
}
