// eslint-rules/no-primevue-uppercase-import.js
module.exports = {
    meta: {
        type: 'problem',
        docs: {
            description: 'Disallow uppercase letters in PrimeVue import paths (e.g. primevue/Select → use primevue/select)',
            category: 'Possible Errors',
            recommended: true
        },
        fixable: null, // atau 'code' jika mau auto-fix
        schema: [],
        messages: {
            uppercasePath: "PrimeVue import path '{{path}}' should be lowercase: '{{suggestion}}'"
        }
    },

    create(context) {
        return {
            ImportDeclaration(node) {
                const source = node.source.value;

                // Cek apakah ini import dari 'primevue/...'
                if (typeof source === 'string' && source.startsWith('primevue/')) {
                    const pathPart = source.slice('primevue/'.length); // e.g. "Select", "DataTable"

                    // Jika ada huruf besar di path → error
                    if (/[A-Z]/.test(pathPart)) {
                        const suggestion = `primevue/${pathPart.toLowerCase()}`;
                        context.report({
                            node: node.source,
                            messageId: 'uppercasePath',
                            data: {
                                path: source,
                                suggestion
                            }
                        });
                    }
                }
            }
        };
    }
};
