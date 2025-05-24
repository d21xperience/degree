<script setup>
import router from '@/router';
import { useAuth } from '@/views/pages/auth/composables/auth';
import { computed, ref } from 'vue';
import { useRoute } from 'vue-router';
import AppMenuItem from './AppMenuItem.vue';
const route = useRoute();
const sekolah = computed(() => route.params.sekolah);
const { onLogout } = useAuth();
const isDialogSignOut = ref(false);
const model = ref([
    {
        label: 'Home',
        items: [{ label: 'Dashboard', icon: 'pi pi-fw pi-home', to: `/${sekolah.value}/dashboard` }]
    },
    // {
    //     label: 'Data Akademik',
    //     items: [
    //         {
    //             label: 'Transkrip',
    //             icon: 'pi pi-fw pi-file'
    //         },
    //         {
    //             label: 'Ijazah',
    //             icon: 'pi pi-fw pi-file'
    //         }
    //     ]
    // },
    {
        label: 'Data User',
        icon: 'pi pi-sign-out',
        items: [
            {
                label: 'Profile',
                icon: 'pi pi-fw pi-user',
                to: `/${sekolah.value}/user-profile`
            }
        ]
    },
    // {
    //     label: 'UI Components',
    //     items: [
    //         { label: 'Form Layout', icon: 'pi pi-fw pi-id-card', to: '/uikit/formlayout' },
    //         { label: 'Input', icon: 'pi pi-fw pi-check-square', to: '/uikit/input' },
    //         { label: 'Button', icon: 'pi pi-fw pi-mobile', to: '/uikit/button', class: 'rotated-icon' },
    //         { label: 'Table', icon: 'pi pi-fw pi-table', to: '/uikit/table' },
    //         { label: 'List', icon: 'pi pi-fw pi-list', to: '/uikit/list' },
    //         { label: 'Tree', icon: 'pi pi-fw pi-share-alt', to: '/uikit/tree' },
    //         { label: 'Panel', icon: 'pi pi-fw pi-tablet', to: '/uikit/panel' },
    //         { label: 'Overlay', icon: 'pi pi-fw pi-clone', to: '/uikit/overlay' },
    //         { label: 'Media', icon: 'pi pi-fw pi-image', to: '/uikit/media' },
    //         { label: 'Menu', icon: 'pi pi-fw pi-bars', to: '/uikit/menu' },
    //         { label: 'Message', icon: 'pi pi-fw pi-comment', to: '/uikit/message' },
    //         { label: 'File', icon: 'pi pi-fw pi-file', to: '/uikit/file' },
    //         { label: 'Chart', icon: 'pi pi-fw pi-chart-bar', to: '/uikit/charts' },
    //         { label: 'Timeline', icon: 'pi pi-fw pi-calendar', to: '/uikit/timeline' },
    //         { label: 'Misc', icon: 'pi pi-fw pi-circle', to: '/uikit/misc' }
    //     ]
    // },
    // {
    //     label: 'Pages',
    //     icon: 'pi pi-fw pi-briefcase',
    //     to: '/pages',
    //     items: [
    //         {
    //             label: 'Landing',
    //             icon: 'pi pi-fw pi-globe',
    //             to: '/landing'
    //         },
    //         {
    //             label: 'Auth',
    //             icon: 'pi pi-fw pi-user',
    //             items: [
    //                 {
    //                     label: 'Login',
    //                     icon: 'pi pi-fw pi-sign-in',
    //                     to: '/auth/login'
    //                 },
    //                 {
    //                     label: 'Error',
    //                     icon: 'pi pi-fw pi-times-circle',
    //                     to: '/auth/error'
    //                 },
    //                 {
    //                     label: 'Access Denied',
    //                     icon: 'pi pi-fw pi-lock',
    //                     to: '/auth/access'
    //                 }
    //             ]
    //         },
    //         {
    //             label: 'Crud',
    //             icon: 'pi pi-fw pi-pencil',
    //             to: '/pages/crud'
    //         },
    //         {
    //             label: 'Not Found',
    //             icon: 'pi pi-fw pi-exclamation-circle',
    //             to: '/pages/notfound'
    //         },
    //         {
    //             label: 'Empty',
    //             icon: 'pi pi-fw pi-circle-off',
    //             to: '/pages/empty'
    //         }
    //     ]
    // },
    {
        label: 'Master',
        items: [
            {
                label: 'Dapodik',
                icon: 'pi pi-fw pi-building-columns',
                items: [
                    {
                        label: 'Data Sekolah',
                        icon: 'pi pi-fw pi-bookmark',
                        to: `/${sekolah.value}/data-dapodik/info-sekolah`
                    },
                    {
                        label: 'Data Kelas',
                        icon: 'pi pi-fw pi-building-columns',
                        to: `/${sekolah.value}/data-dapodik/info-kelas`
                    },
                    {
                        label: 'Data Guru',
                        icon: 'pi pi-fw pi-users',
                        to: `/${sekolah.value}/data-dapodik/info-guru`
                    },
                    {
                        label: 'Data Siswa',
                        icon: 'pi pi-fw pi-users',
                        to: `/${sekolah.value}/data-dapodik/info-siswa`
                    },
                    {
                        label: 'Data Mapel',
                        icon: 'pi pi-fw pi-book',
                        to: `/${sekolah.value}/data-dapodik/info-mapel`
                    }
                ]
            },
            {
                label: 'DNS',
                icon: 'pi pi-fw pi-bookmark',
                items: [
                    {
                        label: 'Ijazah',
                        icon: 'pi pi-fw pi-bookmark',
                        to: `/${sekolah.value}/data-penerima/ijazah`
                    }
                    // {
                    //     label: 'Transkrip Nilai',
                    //     icon: 'pi pi-fw pi-bookmark'
                    // }
                ]
            },
            {
                label: 'Seting',
                icon: 'pi pi-fw pi-cog',
                items: [
                    {
                        label: 'Ijazah & Transkrip',
                        icon: 'pi pi-fw pi-file',
                        to: `/${sekolah.value}/settings/ijazah`
                    },
                    {
                        label: 'Blockahain',
                        icon: 'pi pi-fw pi-globe',
                        to: `/${sekolah.value}/settings/blockchain`
                    },
                    {
                        label: 'IPFS',
                        icon: 'pi pi-fw pi-desktop',
                        to: `/${sekolah.value}/settings/ipfs`
                    }
                ]
            },
            {
                label: 'SC-Ijazah',
                icon: 'pi pi-fw pi-file-check',
                to: `/${sekolah.value}/blockhain/sc-ijazah`
                // command: router.push({ name: 'scIjazah' })
            },
            {
                label: 'Transaksi',
                icon: 'pi pi-chart-line',
                command: () => {
                    router.push({ name: 'daftarTrx' });
                }
            }
        ]
    },

    {
        items: [
            {
                label: 'Sign Out',
                icon: 'pi pi-fw pi-sign-out',
                command: () => {
                    isDialogSignOut.value = true;
                }
            }
        ]
    }
    // {
    //     label: 'Get Started',
    //     items: [
    //         {
    //             label: 'Documentation',
    //             icon: 'pi pi-fw pi-book',
    //             to: '/documentation'
    //         },
    //         {
    //             label: 'View Source',
    //             icon: 'pi pi-fw pi-github',
    //             url: 'https://github.com/primefaces/sakai-vue',
    //             target: '_blank'
    //         }
    //     ]
    // }
]);
// const dialogSignOut = ref(false);
</script>

<template>
    <ul class="layout-menu">
        <template v-for="(item, i) in model" :key="item">
            <app-menu-item v-if="!item.separator" :item="item" :index="i"></app-menu-item>
            <li v-if="item.separator" class="menu-separator"></li>
        </template>
    </ul>

    <!-- Dialog start -->
    <DialogSignOut v-model:visible="isDialogSignOut" @confirm="onLogout" />
    <!-- Dialog end -->
</template>

<style lang="scss" scoped></style>
