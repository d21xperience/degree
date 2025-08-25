<!-- <template>
    <div class="p-4">
        <div class="max-w-5xl max-lg:max-w-3xl max-sm:max-w-sm mx-auto">
            <h2 class="text-3xl font-bold text-slate-900 mb-8">Latest Blog Posts</h2>
            <div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-6 max-sm:gap-8">
                <div class="bg-gray-100 rounded-md overflow-hidden" v-for="item in data" :key="item.id">
                    <div class="bg-gray-50 aspect-[23/16]">
                        <img :src="item.photo_url" alt="Blog Post 1" class="w-full h-full object-cover object-top" />
                    </div>
                    <div class="p-6">
                        <h3 class="text-lg font-semibold text-slate-900 mb-3">{{ item.title }}</h3>
                        <p class="text-slate-600 text-[15px] leading-relaxed line-clamp-3">{{ item.review }}</p>
                        <a href="javascript:void(0);" class="mt-6 inline-block px-3 py-1.5 rounded-md tracking-wider bg-indigo-600 hover:bg-indigo-700 text-white text-sm font-medium">Read More</a>
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>

<script setup>
import { onMounted, ref } from 'vue';

const data = ref([
    {
        id: 1,
        photo_url: 'https://localhost/static/blog/img/hacks-watch.webp',
        title: 'Time Management Hacks',
        review: 'Master your schedule with proven strategies that save hours each week and reduce stress.',
        article_url: ''
    },
    {
        id: 2,
        photo_url: 'https://localhost/static/blog/img/Imagination.webp',
        title: 'The Power of Creativity',
        review: 'Uncover how creative thinking fuels innovation and helps brands stay competitive in any market.',
        article_url: ''
    },
    {
        id: 3,
        photo_url: 'https://localhost/static/blog/img/hotel-img.webp',
        title: 'The Rise of Boutique Hotels',
        review: 'Explore how boutique hotels are redefining luxury and guest experience in the travel industry.',
        article_url: ''
    },
    {
        id: 4,
        photo_url: 'https://localhost/static/blog/img/cardImg.webp',
        title: 'Ecommerce Trends to Watch',
        review: 'Stay ahead with insights into what is driving the future of online shopping and digital retail.',
        article_url: ''
    },
    {
        id: 5,
        photo_url: 'https://localhost/static/blog/img/prediction.webp',
        title: 'Creative Design Trends',
        review: 'Stay ahead of the curve with the latest creative design trends shaping the digital world.',
        article_url: ''
    },
    {
        id: 6,
        photo_url: 'https://localhost/static/blog/img/team-image.webp',
        title: 'Boost Team Productivity',
        review: 'Discover powerful techniques to help your team collaborate better and get more done.',
        article_url: ''
    }
]);
onMounted(() => {});
</script> -->

<script setup></script>

<template>
    <div class="min-h-screen bg-gray-50">
        <!-- Header -->
        <header class="bg-white shadow-sm">
            <div class="container mx-auto px-4 py-6 flex justify-between items-center">
                <h1 class="text-2xl font-bold text-indigo-600">My Blog</h1>
                <nav>
                    <ul class="flex space-x-6">
                        <li><a href="#" class="text-gray-600 hover:text-indigo-600">Home</a></li>
                        <li><a href="#" class="text-gray-600 hover:text-indigo-600">About</a></li>
                        <li><a href="#" class="text-gray-600 hover:text-indigo-600">Contact</a></li>
                    </ul>
                </nav>
            </div>
        </header>

        <!-- Main Content -->
        <main class="container mx-auto px-4 py-8">
            <router-view></router-view>
        </main>

        <!-- Footer -->
        <footer class="bg-gray-800 text-white py-8">
            <div class="container mx-auto px-4">
                <div class="flex flex-col md:flex-row justify-between">
                    <div class="mb-6 md:mb-0">
                        <h3 class="text-xl font-bold mb-4">My Blog</h3>
                        <p class="text-gray-400">Sharing knowledge and ideas about web development, design, and more.</p>
                    </div>
                    <div class="grid grid-cols-2 gap-8">
                        <div>
                            <h4 class="text-lg font-medium mb-4">Quick Links</h4>
                            <ul class="space-y-2">
                                <li><a href="#" class="text-gray-400 hover:text-white">Home</a></li>
                                <li><a href="#" class="text-gray-400 hover:text-white">About</a></li>
                                <li><a href="#" class="text-gray-400 hover:text-white">Blog</a></li>
                                <li><a href="#" class="text-gray-400 hover:text-white">Contact</a></li>
                            </ul>
                        </div>
                        <div>
                            <h4 class="text-lg font-medium mb-4">Connect</h4>
                            <ul class="space-y-2">
                                <li><a href="#" class="text-gray-400 hover:text-white">Twitter</a></li>
                                <li><a href="#" class="text-gray-400 hover:text-white">GitHub</a></li>
                                <li><a href="#" class="text-gray-400 hover:text-white">LinkedIn</a></li>
                                <li><a href="#" class="text-gray-400 hover:text-white">Instagram</a></li>
                            </ul>
                        </div>
                    </div>
                </div>
                <div class="border-t border-gray-700 mt-8 pt-8 text-center text-gray-400">
                    <p>&copy; {{ new Date().getFullYear() }} My Blog. All rights reserved.</p>
                </div>
            </div>
        </footer>
    </div>
</template>

<style>
/* You can add custom styles here if needed */
</style>
