const kelasRoute = [
    {
        path: 'kelas',
        name: 'infoKelas',
        meta: { title: 'Info Kelas', namaRoute: 'Kelas' },
        component: () => import('@/views/pages/dapodik/data_kelas/KelasListView.vue')
    },

    {
        path: 'kelas/:id/edit',
        name: 'editKelas',
        props: true,
        meta: { disableSelect: true, title: 'Tambah Kelas', namaRoute: 'Kelas' },
        component: () => import('@/views/pages/dapodik/data_kelas/KelasEditView.vue')
    },

    {
        path: 'create',
        name: 'inputKelas',
        props: true,
        meta: { disableSelect: true, title: 'Tambah Kelas', namaRoute: 'Kelas' },
        component: () => import('@/views/pages/dapodik/data_kelas/KelasCreateView.vue')
    }
];

export default kelasRoute;
