import AppLayout from '@/layout/AppLayout.vue';
import store from '@/store';
import { ref } from 'vue';
import { createRouter, createWebHistory } from 'vue-router';
import { authGuard } from './guards/authGuard';
export const isLoading = ref(false); // state loading global

const router = createRouter({
    history: createWebHistory(import.meta.env.BASE_URL),
    routes: [
        {
            path: '/',
            name: 'landing',
            component: () => import('@/views/pages/Landing.vue')
        },
        {
            path: '/faq',
            name: 'faq',
            component: () => import('@/views/pages/faq/Faq.vue')
        },
        {
            path: '/blog',
            name: 'blog',
            component: () => import('@/views/pages/blog/Main.vue'),
            children: [
                {
                    path: '',
                    name: 'cek',
                    component: () => import('@/views/pages/blog/artikel/Cek.vue')
                }
            ]
        },
        {
            path: '/auth/login',
            name: 'login',
            meta: { title: 'Login' },
            component: () => import('@/views/pages/auth/Login.vue')
        },
        {
            path: '/auth/register',
            name: 'register',
            component: () => import('@/views/pages/auth/Register.vue'),
            meta: { title: 'Register' }
            // beforeEnter: redirectIfAuthenticated
        },
        {
            path: '/auth/register/success',
            name: 'registerSuccess',
            component: () => import('@/components/authentication_components/RegisterSuccess.vue'),
            meta: { title: 'Register' }
            // beforeEnter: redirectIfAuthenticated
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
                    path: 'manajemen-user',
                    name: 'manajemenUser',
                    meta: { requiresAuth: true, title: 'Manajemen User' },
                    component: () => import('@/views/pages/data_user/MajemenUser.vue'),
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
                            meta: { disableSelect: true, title: 'Edit Guru', namaRoute: 'Guru' },
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
                            meta: { disableSelect: false, title: 'Nilai', namaRoute: 'Info NIlai' },
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
                            path: 'status-kenaikan',
                            name: 'infoKenaikan',
                            meta: { title: 'Info Kenaikan', namaRoute: 'Kelas' },
                            component: () => import('@/views/pages/dapodik/KenaikanDanKelulusan.vue')
                        }
                    ]
                },
                {
                    path: 'settings',
                    name: 'settings',
                    meta: { role: 'admin' },
                    children: [
                        {
                            path: 'ijazah',
                            component: () => import('@/views/pages/sc_ijazah/settings/Ijazah_Setting.vue'),
                            name: 'ijazah'
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
                    path: 'info-semester',
                    name: 'infoSemester',
                    meta: { title: 'Info Semester', namaRoute: 'Semester' },
                    component: () => import('@/views/pages/dapodik/DataSemester.vue')
                },
                // {

                // },
                {
                    path: 'info-kurikulum',
                    name: 'infoKurikulum',
                    meta: { title: 'Info Kurikulum', namaRoute: 'Kurikulum' },
                    component: () => import('@/views/pages/dapodik/DataKurikulum.vue')
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
                            component: () => import('@/views/pages/super_admin/blockchain/SmartContract.vue')
                        },
                        {
                            path: 'networks',
                            name: 'networks',
                            component: () => import('@/views/pages/super_admin/blockchain/ListBCNetwork.vue')
                        },
                        {
                            path: 'wallet',
                            name: 'wallet',
                            meta: { title: 'Wallet Account' },
                            component: () => import('@/views/pages/super_admin/blockchain/ListWallet.vue')
                        },
                        {
                            path: 'bc-network',
                            name: 'bcNetwork',
                            meta: { title: 'Blockhain Network' },
                            component: () => import('@/views/pages/super_admin/blockchain/BCNetworkConfig.vue')
                        }
                    ]
                },
                {
                    path: 'ipfs',
                    name: 'ipfs',
                    component: () => import('@/views/pages/super_admin/blockchain/ipfs.vue')
                },
                {
                    path: 'transaksi',
                    name: 'transaksi',
                    component: () => import('@/views/pages/super_admin/blockchain/Transaksi.vue')
                },
                {
                    path: 'own-profile',
                    name: 'ownProfile',
                    component: () => import('@/views/pages/super_admin/blockchain/OwnProfile.vue')
                }
            ]
        }
    ]
});

// Saat mulai navigasi → tampilkan loading
// router.beforeEach((to, from, next) => {
//     isLoading.value = true;
//     try {
//         const token = localStorage.getItem('token');
//         if (to.meta.requiresAuth && (!token || isTokenExpired(token))) {
//             console.log(isTokenExpired(token));
//             localStorage.clear();
//             next('/auth/login');
//         } else {
//             next();
//         }
//     } catch (error) {
//         alert('error');
//     } finally {
//         isLoading.value = false;
//     }
// });

// 🔐 Global navigation guard
router.beforeEach(async (to, from, next) => {
    isLoading.value = true;
    console.log('berforeEach....', to);
    // Cek auth sekali saat app pertama kali load
    if (store.state.authService.isCheckingAuth && !store.state.authService.user) {
        await store.dispatch('authService/checkAuth');
    }

    const isAuthenticated = store.getters['authService/isAuthenticated'];
    // console.log('nilai isAuthenticated', isAuthenticated);
    if (to.meta.requiresAuth && !isAuthenticated) {
        // next({ name: 'login' });
        // ✅ Simpan query asli + tambahkan redirect jika perlu
        next({
            name: 'login',
            query: {
                ...to.query, // ← lestarikan semua query (termasuk `from=...`)
                redirect: to.fullPath // opsional: untuk redirect setelah login
            }
        });
    } else if (to.meta.guestOnly && isAuthenticated) {
        console.log('gues only');
        next({ name: 'Dashboard' });
    } else {
        next();
    }
});

// Tambahkan afterEach untuk update title
router.afterEach((to) => {
    const defaultTitle = 'Verfikasi Ijazah';
    document.title = to.meta.title ? `${to.meta.title}` : defaultTitle;
    isLoading.value = false;
});

// // 🧨 Tangani 401 dari interceptor → logout & redirect
// window.addEventListener('unauthorized', () => {
//     store.dispatch('authService/logout');
//     router.push({ name: 'login' });
// });

// src/router/index.js (atau tempat Anda daftarkan event listener)

window.addEventListener('unauthorized', () => {
    const currentRoute = router.currentRoute.value;

    // ✅ Logout dulu (hapus cookie & reset state)
    store.dispatch('authService/logout');

    // ✅ Cek: apakah route saat ini memerlukan autentikasi?
    if (currentRoute.meta.requiresAuth) {
        // 🔐 Protected route → redirect ke login + simpan redirect
        router.push({
            name: 'login',
            query: { redirect: currentRoute.fullPath }
        });
    }
    // 🌐 Public route (landing page, dll) → biarkan user di halaman ini
    //    (misal: ganti navbar jadi "Login/Register")
});

export default router;
