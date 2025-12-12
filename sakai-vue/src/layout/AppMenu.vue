<script setup>
import DialogSignOut from '@/components/DialogSignOut.vue';
import LoadingOverlay from '@/components/LoadingOverlay.vue';
import router from '@/router';
import { useAuth } from '@/views/pages/auth/composables/auth';
import { computed, ref } from 'vue';
import { useRoute } from 'vue-router';
import AppMenuItem from './AppMenuItem.vue';
// import { Network as NetworkText } from 'lucide-vue-next';
// import { shallowRef } from 'vue';
// const icons = shallowRef({
//     NetworkText
// });
const route = useRoute();
const sekolah = computed(() => route.params.sekolah);
const { user, onLogout } = useAuth(); // pastikan 'user' ada dan reactive
const isDialogSignOut = ref(false);

const role = computed(() => user.role); //|| 'guest');

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
                    label: 'Dapodik',
                    icon: 'pi pi-fw pi-building-columns',
                    items: [
                        {
                            label: 'Data Semester',
                            icon: 'pi pi-fw pi-bookmark',
                            to: `/${sekolahPath}/data-dapodik/semester`
                        },
                        {
                            label: 'Data Sekolah',
                            icon: 'pi pi-fw pi-building-columns',
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
                            to: `/${sekolahPath}/data-dapodik/kelas`
                        },
                        // {
                        //     label: 'Data Guru',
                        //     icon: 'pi pi-fw pi-users',
                        //     to: `/${sekolahPath}/data-dapodik/info-guru`
                        // },
                        {
                            label: 'Data Siswa',
                            icon: 'pi pi-fw pi-users',
                            to: `/${sekolahPath}/data-dapodik/siswa`
                        },
                        {
                            label: 'Data Nilai',
                            icon: 'pi pi-fw pi-book',
                            to: `/${sekolahPath}/data-dapodik/nilai`
                        },
                        {
                            label: 'Kenaikan & Kelulusan',
                            icon: 'pi pi-fw pi-building-columns',
                            items: [
                                {
                                    label: 'Kenaikan',
                                    icon: 'pi pi-fw pi-building-columns',
                                    to: `/${sekolahPath}/data-dapodik/status-kenaikan`
                                },
                                {
                                    label: 'Kelulusan',
                                    icon: 'pi pi-fw pi-building-columns',
                                    to: `/${sekolahPath}/data-dapodik/status-kenaikan`
                                }
                            ]
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
                        }
                        // {
                        //     label: 'Transkrip Nilai',
                        //     icon: 'pi pi-fw pi-bookmark',
                        //     to: `/${sekolahPath}/data-penerima/transkrip`
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
                },
                {
                    label: 'User',
                    icon: 'pi pi-fw pi-users',
                    items: [
                        {
                            label: 'Manajemen user',
                            icon: 'pi pi-fw pi-user',
                            to: `/${sekolahPath}/manajemen-user`
                        }
                        // {
                        //     label: 'Transkrip Nilai',
                        //     icon: 'pi pi-fw pi-bookmark',
                        //     to: `/${sekolahPath}/data-penerima/transkrip`
                        // }
                    ]
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
                    label: 'Master Data Akademik',
                    icon: 'pi pi-fw pi-bookmark',
                    items: [
                        {
                            label: 'Semester',
                            icon: 'pi pi-fw pi-bookmark',
                            to: `/su/info-semester`
                        },
                        {
                            label: 'Kurikulum',
                            icon: 'pi pi-fw pi-bookmark',
                            to: `/su/info-kurikulum`
                        },
                        {
                            label: 'Kelompok Mapel',
                            icon: 'pi pi-fw pi-bookmark',
                            to: '/su/daftar-sekolah'
                        },
                        {
                            label: 'Daftar sekolah',
                            icon: 'pi pi-fw pi-building-columns',
                            to: '/su/daftar-sekolah'
                        }
                    ]
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
    // console.log(newValue);
    switch (newValue) {
        case 'superadmin':
            router.push({ name: 'ownProfile' });
            break;
        case 'admin':
            router.push({ name: 'userProfile' });
            break;
        default:
            break;
    }
};

const isLoading = ref(false);
const logout = async () => {
    isLoading.value = true;
    try {
        isDialogSignOut.value = false;
        router.push({ name: 'landing' });
        await onLogout();
    } catch (error) {
        console.log(error);
    } finally {
        isLoading.value = false;
    }
};
</script>

<template>
    <ul class="layout-menu">
        <template v-for="(item, i) in model" :key="i">
            <app-menu-item v-if="!item.separator" :item="item" :index="i" />
            <li v-if="item.separator" class="menu-separator"></li>
        </template>
    </ul>

    <DialogSignOut v-model:visible="isDialogSignOut" @confirm="logout" />
    <LoadingOverlay :visible="isLoading"> Memuat data, harap tunggu... </LoadingOverlay>
</template>
