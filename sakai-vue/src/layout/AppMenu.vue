<script setup>
import DialogSignOut from '@/components/DialogSignOut.vue';
import router from '@/router';
import { useAuth } from '@/views/pages/auth/composables/auth';
import { computed, ref } from 'vue';
import { useRoute } from 'vue-router';
import AppMenuItem from './AppMenuItem.vue';

import { Network as NetworkText } from 'lucide-vue-next';
import { shallowRef } from 'vue';
const icons = shallowRef({
    NetworkText
});
const route = useRoute();
const sekolah = computed(() => route.params.sekolah);
const { user, onLogout } = useAuth(); // pastikan 'user' ada dan reactive
const isDialogSignOut = ref(false);

const role = computed(() => user); //|| 'guest');

const model = computed(() => {
    const sekolahPath = sekolah.value;
    const menu = [];

    // Menu: Home
    // menu.push({
    //     label: 'Home',
    //     items: [
    //         {
    //             label: 'Dashboard',
    //             icon: 'pi pi-fw pi-home',
    //             to: `/${sekolahPath}/dashboard`
    //         }
    //     ]
    // });

    // Menu untuk ADMIN
    if (role.value === 'admin') {
        menu.push({
            label: 'Master',
            items: [
                {
                    label: 'Dashboard',
                    icon: 'pi pi-fw pi-home',
                    to: `/${sekolahPath}/dashboard`
                },
                {
                    label: 'User',
                    icon: 'pi pi-fw pi-users',
                    items: [
                        {
                            label: 'Manajemen user',
                            icon: 'pi pi-fw pi-user',
                            to: `/${sekolahPath}/data-penerima/ijazah`
                        }
                        // {
                        //     label: 'Transkrip Nilai',
                        //     icon: 'pi pi-fw pi-bookmark',
                        //     to: `/${sekolahPath}/data-penerima/transkrip`
                        // }
                    ]
                },
                {
                    label: 'Dapodik',
                    icon: 'pi pi-fw pi-building-columns',
                    items: [
                        {
                            label: 'Semester',
                            icon: 'pi pi-fw pi-bookmark',
                            to: `/${sekolahPath}/data-dapodik/info-semester`
                        },
                        {
                            label: 'Data Sekolah',
                            icon: 'pi pi-fw pi-bookmark',
                            to: `/${sekolahPath}/data-dapodik/info-sekolah`
                        },
                        {
                            label: 'Data Mapel',
                            icon: 'pi pi-fw pi-book',
                            to: `/${sekolahPath}/data-dapodik/info-mapel`
                        },
                        {
                            label: 'Data Kelas',
                            icon: 'pi pi-fw pi-building-columns',
                            to: `/${sekolahPath}/data-dapodik/info-kelas`
                        },
                        {
                            label: 'Data Guru',
                            icon: 'pi pi-fw pi-users',
                            to: `/${sekolahPath}/data-dapodik/info-guru`
                        },
                        {
                            label: 'Data Siswa',
                            icon: 'pi pi-fw pi-users',
                            to: `/${sekolahPath}/data-dapodik/info-siswa`
                        },
                        {
                            label: 'Data Nilai',
                            icon: 'pi pi-fw pi-book',
                            to: `/${sekolahPath}/data-dapodik/info-nilai`
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
                            to: `/${sekolahPath}/data-penerima/ijazah`
                        },
                        {
                            label: 'Transkrip Nilai',
                            icon: 'pi pi-fw pi-bookmark',
                            to: `/${sekolahPath}/data-penerima/transkrip`
                        }
                    ]
                },
                {
                    label: 'Seting',
                    icon: 'pi pi-fw pi-cog',
                    items: [
                        {
                            label: 'Ijazah & Transkrip',
                            icon: 'pi pi-fw pi-file',
                            to: `/${sekolahPath}/settings/ijazah`
                        },
                        {
                            label: 'Blockahain',
                            icon: 'pi pi-fw pi-globe',
                            to: `/${sekolahPath}/settings/blockchain`
                        },
                        {
                            label: 'IPFS',
                            icon: 'pi pi-fw pi-desktop',
                            to: `/${sekolahPath}/settings/ipfs`
                        }
                    ]
                },
                {
                    label: 'SC-Ijazah',
                    icon: 'pi pi-fw pi-file-check',
                    to: `/${sekolahPath}/blockhain/sc-ijazah`
                },
                {
                    label: 'Transaksi',
                    icon: 'pi pi-chart-line',
                    command: () => {
                        router.push({ name: 'daftarTrx' });
                    }
                }
            ]
        });
    }

    // Menu: Super Admin
    if (role.value === 'superadmin') {
        menu.push({
            label: 'Super Admin',
            items: [
                {
                    label: 'Dashboard',
                    icon: 'pi pi-fw pi-home',
                    to: '/su/dashboard'
                },
                {
                    label: 'Daftar sekolah',
                    icon: 'pi pi-fw pi-building-columns',
                    to: '/su/daftar-sekolah'
                },
                {
                    label: 'Blockchain',
                    icon: 'pi pi-fw pi-box',
                    items: [
                        {
                            label: 'Network',
                            icon: 'pi pi-network',
                            to: '/su/blockchain/bc-network'
                        },

                        {
                            label: 'Wallet',
                            to: '/su/blockchain/wallet'
                        },
                        {
                            label: 'Smart Contract',
                            to: '/su/blockchain/smart-contract'
                        },
                        {
                            label: 'Environment',
                            to: '/su/blockchain/networks'
                        }
                    ]
                },
                {
                    label: 'IPFS',
                    icon: 'pi pi-fw pi-globe',
                    to: '/su/ipfs'
                },
                {
                    label: 'Transaksi',
                    icon: 'pi pi-fw pi-chart-bar',
                    to: '/su/transaksi'
                },
                {
                    label: 'CMS',
                    icon: 'pi pi-fw pi-pen-to-square',
                    to: '/su/cms'
                }
            ]
        });
    }

    // Menu untuk Siswa
    if (role.value === 'siswa') {
        menu.push({
            label: 'Siswa',
            item: [
                {
                    label: 'Ijazah',
                    icon: 'pi pi-fw pi-file',
                    command: () => {
                        isDialogSignOut.value = true;
                    }
                }
            ]
        });
    }
    // Menu: Sign Out (selalu ditampilkan)
    menu.push({
        items: [
            {
                label: 'Profile',
                icon: 'pi pi-fw pi-user',
                command: () => getProfile(role.value)
            },
            {
                label: 'Sign Out',
                icon: 'pi pi-fw pi-sign-out',
                command: () => {
                    isDialogSignOut.value = true;
                }
            }
        ]
    });

    return menu;
});

const getProfile = (newValue) => {
    console.log(newValue);
    switch (newValue) {
        case 'superadmin':
            router.push({ name: 'ownProfile' });
            break;

        default:
            break;
    }
};

// onMounted(() => {
//     console.log(role.value);
// });
</script>

<template>
    <ul class="layout-menu">
        <template v-for="(item, i) in model" :key="i">
            <app-menu-item v-if="!item.separator" :item="item" :index="i" />
            <li v-if="item.separator" class="menu-separator"></li>
        </template>
    </ul>

    <DialogSignOut v-model:visible="isDialogSignOut" @confirm="onLogout" />
</template>
