var _env = {};

/**
 * Returns value of environment variable
 * @param {string} envname - Variable name
 * @param {string} [fallback] - Default value
 */
export function getVariable(envname, fallback = '') {
   return Reflect.get(_env, envname) || fallback;
}

export function setVariables(vars = {}) {
   _env = vars;
}

/**
 * Platform environment
 * @returns {string}
 */
export function getEnv() {
   return Reflect.get(_env, 'ENV') || '';
}

/**
 * Console Admin Id
 * @returns {string}
 */
export function getAdminId() {
   return Reflect.get(_env, 'ADMIN_ID') || '';
}

/**
 * Cloudinary Id
 * @returns {string}
 */
export function getCloudinaryId() {
   return Reflect.get(_env, 'CLOUDINARY_ID') || '';
}
