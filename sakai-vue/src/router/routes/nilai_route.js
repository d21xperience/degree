const nilaiRoute = [
    {
        path: 'nilai',
        children: [
            {
                path: '',
                name: 'infoNilai',
                props: true,
                meta: { disableSelect: false, title: 'Nilai', namaRoute: 'Info NIlai' },
                component: () => import('@/views/pages/dapodik/data_nilai/DataNilai.vue')
            },
            {
                path: ':id/edit',
                name: 'inputNilai',
                props: true,
                meta: { disableSelect: true, title: 'Tambah Nilai', namaRoute: 'Siswa' },
                component: () => import('@/views/pages/dapodik/data_siswa/AddSiswa.vue')
            }
        ]
    }
];

export default nilaiRoute;
