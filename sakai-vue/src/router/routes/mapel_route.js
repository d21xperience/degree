const mapelRoute = [
    {
        path: 'info-mapel',
        name: 'infoMapel',
        meta: { disableSemester: true, title: 'Info Mapel', namaRoute: 'Mata Pelajaran' },
        component: () => import('@/views/pages/dapodik/data_matapelajaran/ReadMapel.vue')
    }
];

export default mapelRoute;
