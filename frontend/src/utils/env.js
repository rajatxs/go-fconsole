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
