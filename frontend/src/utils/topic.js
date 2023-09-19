import topics from '../data/topics.json';

export function getAllTopics() {
   return Object.entries(topics)
      .map(function (_topic) {
         Reflect.set(_topic[1], 'id', _topic[0]);
         return _topic[1];
      })
      .sort((a, b) => a.name.localeCompare(b.name));
}

export function getPublicTopics() {
   return getAllTopics().filter((_topic) => _topic.public);
}

export function getPrivateTopics() {
   return getAllTopics().filter((_topic) => !_topic.public);
}

/**
 * Returns name of the topic by given `id`
 * @param {string} id - Topic Id
 */
export function getTopicName(id) {
   if (topics[id]) {
      return topics[id].name;
   } else {
      return 'Other';
   }
}
