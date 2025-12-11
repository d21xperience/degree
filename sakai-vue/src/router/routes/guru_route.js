const guruRoute = [
    {
        path: 'info-guru',
        name: 'infoGuru',
        meta: { title: 'Info Guru', namaRoute: 'Guru' },
        component: () => import('@/views/pages/dapodik/data_guru/GuruListView.vue')
    },
    {
        path: 'input-guru',
        name: 'inputGuru',
        props: true,
        meta: { disableSelect: true, title: 'Tambah Guru' },
        component: () => import('@/views/pages/dapodik/data_guru/GuruCreateView.vue')
    },
    {
        path: 'edit-guru',
        name: 'editGuru',
        props: true,
        meta: { disableSelect: true, title: 'Edit Guru', namaRoute: 'Guru' },
        component: () => import('@/views/pages/dapodik/data_guru/GuruCreateView.vue')
    }
];

export default guruRoute;
