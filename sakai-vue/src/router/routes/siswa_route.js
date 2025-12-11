const siswaRoute = [
    {
        path: 'siswa',
        children: [
            {
                path: '',
                name: 'infoSiswa',
                props: true,
                meta: { title: 'Info Siswa', namaRoute: 'Siswa' },
                component: () => import('@/views/pages/dapodik/data_siswa/ReadSiswa.vue')
            },
            {
                path: 'tambah-siswa',
                name: 'createSiswa',
                props: true,
                meta: { disableSelect: true, disableSemester: true, title: 'Tambah Siswa', namaRoute: 'Tambah Siswa' },
                component: () => import('@/views/pages/dapodik/data_siswa/SiswaCreateView.vue')
            },
            {
                path: ':id/edit',
                name: 'editSiswa',
                props: true,
                meta: { disableSelect: true, title: 'Edit Siswa' },
                component: () => import('@/views/pages/dapodik/data_siswa/AddSiswa.vue')
            }
        ]
    }
];

export default siswaRoute;
