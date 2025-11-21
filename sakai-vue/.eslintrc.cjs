// /* eslint-env node */
// require('@rushstack/eslint-patch/modern-module-resolution');

// module.exports = {
//     root: true,
//     env: {
//         browser: true,
//         es2021: true,
//         node: true
//     },
//     // env: {
//     //     node: true
//     // },
//     // extends: ['plugin:vue/vue3-essential', 'eslint:recommended', '@vue/eslint-config-prettier'],
//     extends: [
//         'eslint:recommended',
//         'plugin:vue/vue3-recommended', // ← lebih ketat dari 'essential', cocok untuk proyek production
//         '@vue/eslint-config-prettier' // ← menonaktifkan aturan yang bentrok dengan Prettier
//     ],
//     // parserOptions: {
//     //     ecmaVersion: 'latest'
//     // },
//     parserOptions: {
//         ecmaVersion: 'latest',
//         sourceType: 'module'
//     },
//     plugins: ['vue'],
//     // rules: {
//     //     'vue/multi-word-component-names': 'off',
//     //     'vue/no-reserved-component-names': 'off',
//     //     'vue/component-tags-order': [
//     //         'error',
//     //         {
//     //             order: ['script', 'template', 'style']
//     //         }
//     //     ]
//     // }
//     rules: {
//         // 🔥 Aturan kritis berdasarkan error yang pernah kamu alami:
//         'vue/no-lifecycle-after-await': 'error', // cegah onMounted() setelah await
//         'vue/multi-word-component-names': 'off', // izinkan AppTopbar.vue, dll.
//         'vue/require-default-prop': 'warn',
//         'vue/return-in-computed-property': 'error',

//         // 🔒 Keamanan & keandalan:
//         'no-undef': 'error',
//         'no-unused-vars': ['error', { argsIgnorePattern: '^_' }],
//         'no-unsafe-optional-chaining': ['error', { disallowArithmeticOperators: true }],

//         // 🧹 Gaya (opsional, disesuaikan dengan Prettier):
//         'vue/html-self-closing': [
//             'error',
//             {
//                 html: { void: 'always', normal: 'never', component: 'always' }
//             }
//         ],
//         'vue/no-unregistered-components': [
//             'error',
//             {
//                 ignorePatterns: ['Button', 'Menu', 'Card', 'Dialog', '*'] // ← sesuaikan dengan komponen PrimeVue yang sering dipakai
//             }
//         ]
//     }
// };
/* eslint-env node */
require('@rushstack/eslint-patch/modern-module-resolution');

module.exports = {
    root: true,
    env: {
        browser: true,
        es2021: true,
        node: true
    },
    extends: ['eslint:recommended', 'plugin:vue/vue3-recommended', '@vue/eslint-config-prettier'],
    parserOptions: {
        ecmaVersion: 'latest',
        sourceType: 'module'
    },
    plugins: ['vue'],
    globals: {
        // Vue 3 macro compiler globals
        defineProps: 'readonly',
        defineEmits: 'readonly',
        defineExpose: 'readonly',
        withDefaults: 'readonly'
    },
    rules: {
        // 🔥 Kritis
        'vue/no-lifecycle-after-await': 'error',
        'vue/return-in-computed-property': 'error',

        // 🛑 Nonaktifkan karena auto-import PrimeVue
        'vue/no-unregistered-components': 'off',

        // 🔒 Keamanan
        'no-undef': 'error',
        'no-unused-vars': [
            'error',
            {
                argsIgnorePattern: '^_',
                varsIgnorePattern: '^(ref|reactive|computed|shallowRef|readonly)$'
            }
        ],
        'no-unsafe-optional-chaining': ['error', { disallowArithmeticOperators: true }],

        // 🧹 Gaya
        'vue/multi-word-component-names': 'off',
        'vue/require-default-prop': 'warn',
        'vue/html-self-closing': [
            'error',
            {
                html: {
                    void: 'always',
                    normal: 'never',
                    component: 'always'
                }
            }
        ]
    },

    // ✅ Hindari false-positive di file konfigurasi
    overrides: [
        {
            files: ['*.config.js', 'vite.config.*', 'tailwind.config.*'],
            env: { node: true },
            rules: {
                'no-undef': 'off',
                'no-unused-vars': 'off'
            }
        }
    ]
};
