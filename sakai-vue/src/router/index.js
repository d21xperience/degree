import AppLayout from '@/layout/AppLayout.vue';
import { createRouter, createWebHistory } from 'vue-router';
import authGuard from './guards/authGuard';
let kondisi = false;
const router = createRouter({
    history: createWebHistory(import.meta.env.BASE_URL),
    routes: [
        {
            path: '/',
            name: 'landing',
            component: () => import('@/views/pages/Landing.vue')
        },
        {
            path: '/auth/login',
            name: 'login',
            component: () => import('@/views/pages/auth/Login.vue')
        },
        {
            path: '/auth/register',
            name: 'register',
            component: () => import('@/views/pages/auth/Register.vue'),
            meta: {
                title: 'Register'
            }
        },
        {
            path: '/auth/access',
            name: 'accessDenied',
            component: () => import('@/views/pages/auth/Access.vue')
        },
        {
            path: '/auth/error',
            name: 'error',
            component: () => import('@/views/pages/auth/Error.vue')
        },
        {
            path: '/:sekolah/',
            component: AppLayout,
            // name: 'main',
            props: true,
            beforeEnter: authGuard,
            meta: { requiresAuth: true },
            children: [
                {
                    path: 'dashboard',
                    name: 'dashboard',
                    meta: { requiresAuth: true, title: 'Dashboard' },
                    component: () => import('@/views/Dashboard.vue'),
                    props: true
                },
                {
                    path: 'user-profile',
                    name: 'userProfile',
                    meta: { requiresAuth: true, title: 'User Profile' },
                    component: () => import('@/views/pages/data_user/Profile.vue'),
                    props: true
                },
                {
                    path: 'data-dapodik',
                    component: () => import('@/views/pages/dapodik/Main.vue'),
                    meta: { role: 'admin' },
                    children: [
                        {
                            path: 'info-siswa',
                            name: 'infoSiswa',
                            props: true,
                            meta: { title: 'Info Siswa', namaRoute: 'Siswa' },
                            component: () => import('@/views/pages/dapodik/data_siswa/ReadSiswa.vue')
                        },
                        {
                            path: 'info-sekolah',
                            name: 'infoSekolah',
                            meta: { title: 'Info Sekolah', namaRoute: 'Sekolah' },
                            component: () => import('@/views/pages/dapodik/DataSekolah.vue')
                        },
                        {
                            path: 'info-guru',
                            name: 'infoGuru',
                            meta: { title: 'Info Guru', namaRoute: 'Guru' },
                            component: () => import('@/views/pages/dapodik/data_guru/ReadGuru.vue')
                        },
                        {
                            path: 'input-guru',
                            name: 'inputGuru',
                            props: true,
                            meta: { disableSelect: true, title: 'Tambah Guru' },
                            component: () => import('@/views/pages/dapodik/data_guru/AddGuru.vue')
                        },
                        {
                            path: 'edit-guru',
                            name: 'editGuru',
                            props: true,
                            meta: { disableSelect: true, title: 'Edit Guru' },
                            component: () => import('@/views/pages/dapodik/data_guru/AddGuru.vue')
                        },

                        {
                            path: 'info-kelas',
                            name: 'infoKelas',
                            meta: { title: 'Info Kelas', namaRoute: 'Kelas' },
                            component: () => import('@/views/pages/dapodik/data_kelas/ReadKelas.vue')
                        },

                        {
                            path: 'edit-kelas',
                            name: 'editKelas',
                            props: true,
                            meta: { disableSelect: true, title: 'Tambah Kelas', namaRoute: 'Kelas' },
                            component: () => import('@/views/pages/dapodik/data_kelas/AddKelas.vue')
                        },

                        {
                            path: 'input-kelas',
                            name: 'inputKelas',
                            props: true,
                            meta: { disableSelect: true, title: 'Tambah Kelas', namaRoute: 'Kelas' },
                            component: () => import('@/views/pages/dapodik/data_kelas/AddKelas.vue')
                        },
                        {
                            path: 'input-siswa',
                            name: 'inputSiswa',
                            props: true,
                            meta: { disableSelect: true, title: 'Tambah Siswa', namaRoute: 'Siswa' },
                            component: () => import('@/views/pages/dapodik/data_siswa/AddSiswa.vue')
                        },
                        {
                            path: 'edit-siswa',
                            name: 'editSiswa',
                            props: true,
                            meta: { disableSelect: true, title: 'Edit Siswa' },
                            component: () => import('@/views/pages/dapodik/data_siswa/AddSiswa.vue')
                        },
                        {
                            path: 'info-mapel',
                            name: 'infoMapel',
                            meta: { title: 'Info Mapel', namaRoute: 'Mata Pelajaran' },
                            component: () => import('@/views/pages/dapodik/data_matapelajaran/ReadMapel.vue')
                        },
                        {
                            path: 'info-nilai',
                            name: 'infoNilai',
                            props: true,
                            meta: { disableSelect: true, title: 'Nilai', namaRoute: 'Info NIlai' },
                            component: () => import('@/views/pages/dapodik/data_nilai/DataNilai.vue')
                        },
                        {
                            path: 'input-nilai',
                            name: 'inputNilai',
                            props: true,
                            meta: { disableSelect: true, title: 'Tambah Nilai', namaRoute: 'Siswa' },
                            component: () => import('@/views/pages/dapodik/data_siswa/AddSiswa.vue')
                        },
                        {
                            path: 'edit-siswa',
                            name: 'editSiswa',
                            props: true,
                            meta: { disableSelect: true, title: 'Edit Siswa' },
                            component: () => import('@/views/pages/dapodik/data_siswa/AddSiswa.vue')
                        }
                    ]
                },
                // {
                //     path: 'blockhain',
                //     name: 'blockhain',
                //     component: () => import('@/views/super_admin/Main.vue'),
                //     children: [
                //         {
                //             path: 'sc-ijazah',
                //             name: 'scIjazah',
                //             component: () => import('@/views/pages/sc_ijazah/SCIjazah.vue')
                //         },
                //         {
                //             path: 'daftar-trx',
                //             name: 'daftarTrx',
                //             component: () => import('@/views/pages/sc_ijazah/ListTrx.vue')
                //         }
                //     ]
                // },
                {
                    path: 'settings',
                    name: 'settings',
                    meta: { role: 'admin' },
                    children: [
                        {
                            path: 'blockchain',
                            component: () => import('@/views/pages/sc_ijazah/settings/BlockchainSetting.vue')
                        },
                        {
                            path: 'ipfs',
                            component: () => import('@/views/pages/sc_ijazah/settings/Client_IPFS.vue')
                        },
                        {
                            path: 'ijazah',
                            component: () => import('@/views/pages/sc_ijazah/settings/Ijazah_Setting.vue')
                        }
                    ]
                },
                {
                    path: 'data-penerima',
                    component: () => import('@/views/pages/data_penerima/Main.vue'),
                    meta: { role: 'admin' },
                    children: [
                        {
                            path: 'ijazah',
                            name: 'readIjazah',
                            component: () => import('@/views/pages/data_penerima/Ijazah.vue')
                        },
                        {
                            path: 'edit-ijazah',
                            name: 'editIjazah',
                            meta: { disableSelect: true, title: 'Tambah Kelas', namaRoute: 'Kelas' },
                            props: true,
                            component: () => import('@/views/pages/data_penerima/EditDataPenerima.vue')
                        },
                        {
                            path: 'transkrip',
                            name: 'readTranskrip',
                            component: () => import('@/views/pages/data_penerima/Transkrip.vue')
                        }
                    ]
                }
            ]
        },
        {
            path: '/pages/notfound',
            name: 'notfound',
            component: () => import('@/views/pages/NotFound.vue')
        },

        {
            path: '/not-found',
            name: 'not-found',
            component: () => import('@/views/pages/NotFound.vue')
        },
        {
            path: '/:pathMatch(.*)*',
            redirect: '/not-found'
        },
        {
            path: '/su',
            name: 'superAdmin',
            props: true,
            beforeEnter: authGuard,
            meta: { requiresAuth: true, role: 'superadmin' },
            component: AppLayout,
            children: [
                {
                    path: 'dashboard',
                    name: 'suDashboard',
                    component: () => import('@/views/pages/super_admin/Dashboard.vue')
                },
                {
                    path: 'daftar-sekolah',
                    name: 'daftarSekolah',
                    component: () => import('@/views/pages/super_admin/Dashboard.vue')
                },
                {
                    path: 'cms',
                    name: 'cms',
                    component: () => import('@/views/pages/super_admin/cms/Content.vue')
                },
                {
                    path: 'blockchain',
                    // name: 'smartContract',
                    component: () => import('@/views/pages/super_admin/blockchain/Main.vue'),
                    children: [
                        {
                            path: 'smart-contract',
                            name: 'smartcontract',
                            component: () => import('@/views/pages/super_admin/blockchain/BlockchainSettings.vue')
                        },
                        {
                            path: 'networks',
                            name: 'networks',

                            component: () => import('@/views/pages/super_admin/blockchain/ListBCNetwork.vue')
                        },
                        {
                            path: 'ipfs',
                            name: 'ipfs'

                            // component: () => import('@views/pages/super_admin/smart_contract/ListBCNetworks')
                        }
                    ]
                }
            ]
        }
    ]
});

export default router;
