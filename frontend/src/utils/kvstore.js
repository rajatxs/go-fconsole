/** Returns default key prefix */
export function prefix() {
   return 'fconsole'
}

/**
 * Returns raw key including `prefix` and `suffix`
 * @param {string} suffix - Key suffix
 * @returns {string}
 */
export function key(suffix) {
   return `${prefix()}:${suffix}`
}

/**
 * Returns value by given key
 * @param {string} keySuffix - Key suffix
 * @param {string} [fallback] - Default value
 * @returns {string}
 */
export function getString(keySuffix, fallback = '') {
   return localStorage.getItem(key(keySuffix)) || fallback
}

/**
 * Writes specified `val` to storage
 * @param {string} keySuffix - Key suffix
 * @param {string} val - Value
 */
export function setString(keySuffix, val) {
   localStorage.setItem(key(keySuffix), val)
}
