// eslint-rules/index.js
const noPrimevueUppercase = require('./no-primevue-uppercase-import');

/** @type {import('eslint').Rule.RuleModule} */
module.exports.rules = {
    'no-primevue-uppercase-import': noPrimevueUppercase
};

// Opsional: beri nama plugin
module.exports.name = 'primevue-safe';
