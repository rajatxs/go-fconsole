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
