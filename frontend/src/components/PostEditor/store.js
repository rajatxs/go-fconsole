import { reactive } from 'vue';

export const state = reactive({
   /** @type {string} */
   title: '',

   /** @type {string} */
   slug: '',

   /** @type {string|null} */
   topic: null,

   /** @type {string} */
   desc: '',

   /** @type {string[]} */
   tags: [],

   /** @type {boolean} */
   publicScope: true,

   /** @type {string} */
   coverImageRefName: '',

   /** @type {string} */
   coverImageRefUrl: '',

   /** @type {string} */
   coverImagePublicId: '',

   /** @type {string} */
   coverImageAssetId: '',
});

/**
 * Sets state properties from given `data` payload
 * @param {import('../../../wailsjs/go/models').models.PostDocument} data 
 */
export function setMetadata(data) {
   state.title = data.title;
   state.slug = data.slug;
   state.topic = data.topic;
   state.desc = data.desc;
   state.tags = data.tags;
   state.publicScope = data.public;

   if (data.coverImage) {
      state.coverImageRefName = data.coverImage.refName;
      state.coverImageRefUrl = data.coverImage.refUrl;
      state.coverImagePublicId = data.coverImage.path;
      state.coverImageAssetId = data.coverImage.id;
   }
}

/** Sets default values to metadata reactive properties */
export function clearMetadata() {
   state.title = '';
   state.slug = '';
   state.topic = null;
   state.desc = '';
   state.tags = [];
   state.publicScope = true;
   state.coverImageRefName = '';
   state.coverImageRefUrl = '';
   state.coverImagePublicId = '';
   state.coverImageAssetId = '';
}
