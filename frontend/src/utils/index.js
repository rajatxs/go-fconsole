import { getVariable } from './env';

/**
 * Returns grouped array with specific number of array elements
 * @param {any[]} arr
 * @param {number} count
 * @returns
 */
export function groupArray(arr, count = 3) {
   const result = [];

   for (let i = 0; i < arr.length; i += count) {
      result.push(arr.slice(i, i + count));
   }

   return result;
}

/**
 * Truncates long text
 * @param {string} text
 * @param {number} len
 */
export function truncateText(text, len = 36) {
   if (text.length > len) {
      return text.substring(0, len - 3) + '...';
   }
   return text;
}

/**
 * Returns simplified date string
 * @param {Date|string} time
 */
export function formatTime(time) {
   if (typeof time === 'string') {
      time = new Date(time);
   }

   const currentTime = new Date();
   const months = [
      'Jan',
      'Feb',
      'Mar',
      'Apr',
      'May',
      'Jun',
      'Jul',
      'Aug',
      'Sep',
      'Oct',
      'Nov',
      'Dec',
   ];

   const month = months[time.getMonth()];

   if (time.getFullYear() === currentTime.getFullYear()) {
      return `${month} ${time.getDate()}`;
   } else {
      return `${month} ${time.getDate()}, ${time.getFullYear()}`;
   }
}

/**
 * Returns absolute image url of cover image by given `imagePath`
 * @param {string} imagePath - Post cover image path
 */
export function getPostCoverImageURL(imagePath) {
   return `https://res.cloudinary.com/${getVariable('CLOUDINARY_ID')}/image/upload/c_scale,h_600/${imagePath}.webp`
}

/**
 * Returns absolute image url of embedded image inside post body
 * @param {string} imagePath - Image path
 */
export function getPostEmbeddedImageUrl(imagePath) {
   return `https://res.cloudinary.com/${getVariable('CLOUDINARY_ID')}/image/upload/c_scale,h_600/${imagePath}`
}

/**
 * Returns absolute image url of post topic thumb
 * @param {string} imagePath - Image path
 */
export function getPostTopicImageUrl(imagePath) {
   return `https://res.cloudinary.com/${getVariable('CLOUDINARY_ID')}/image/upload/c_scale,h_400/${imagePath}.webp`
}

/**
 * Returns computed slug from given `text`
 * @param {string} text 
 */
export function computeSlug(text) {
   return text
      .toLowerCase()
      .replace(/\s+/g, '-')
      .replace(/[^a-z0-9-]/g, '')
      .replace(/-+/g, '-')
      .replace(/^-|-$/g, '');
}
