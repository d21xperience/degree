const semesterRoute = [
    {
        path: 'semester',
        // name: 'semesterInfo',
        meta: { disableSelect: true, title: 'Semester' },
        children: [
            {
                path: '',
                name: 'semesterInfo',
                meta: { title: 'Semester', namaRoute: 'Semester' },
                component: () => import('@/views/pages/dapodik/data_semester/SemesterListView.vue')
            }
        ]
    }
];

export default semesterRoute;
