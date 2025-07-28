<template>
    <div>
        Featured Post
        <section class="mb-12">
            <div class="bg-white rounded-lg shadow-md overflow-hidden">
                <img :src="featuredPost.image" alt="Featured Post" class="w-full h-64 object-cover" />
                <div class="p-6">
                    <span class="inline-block px-3 py-1 bg-indigo-100 text-indigo-600 rounded-full text-sm font-medium mb-2">
                        {{ featuredPost.category }}
                    </span>
                    <h2 class="text-2xl font-bold text-gray-800 mb-2">{{ featuredPost.title }}</h2>
                    <p class="text-gray-600 mb-4">{{ featuredPost.excerpt }}</p>
                    <div class="flex items-center">
                        <img :src="featuredPost.author.avatar" alt="Author" class="w-10 h-10 rounded-full mr-3" />
                        <div>
                            <p class="text-sm font-medium text-gray-900">{{ featuredPost.author.name }}</p>
                            <p class="text-sm text-gray-500">{{ formatDate(featuredPost.date) }}</p>
                        </div>
                        <a href="#" class="ml-auto px-4 py-2 bg-indigo-600 text-white rounded-md hover:bg-indigo-700 transition"> Read More </a>
                    </div>
                </div>
            </div>
        </section>

        <!-- Blog Posts Grid -->
        <section>
            <h2 class="text-2xl font-bold text-gray-800 mb-6">Latest Articles</h2>
            <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-8">
                <article v-for="post in posts" :key="post.id" class="bg-white rounded-lg shadow-md overflow-hidden hover:shadow-lg transition">
                    <img :src="post.image" :alt="post.title" class="w-full h-48 object-cover" />
                    <div class="p-6">
                        <span class="inline-block px-3 py-1 bg-indigo-100 text-indigo-600 rounded-full text-sm font-medium mb-2">
                            {{ post.category }}
                        </span>
                        <h3 class="text-xl font-bold text-gray-800 mb-2">{{ post.title }}</h3>
                        <p class="text-gray-600 mb-4">{{ post.excerpt }}</p>
                        <div class="flex items-center">
                            <img :src="post.author.avatar" alt="Author" class="w-8 h-8 rounded-full mr-2" />
                            <div>
                                <p class="text-sm font-medium text-gray-900">{{ post.author.name }}</p>
                                <p class="text-xs text-gray-500">{{ formatDate(post.date) }}</p>
                            </div>
                            <a href="#" class="ml-auto text-indigo-600 hover:text-indigo-800 text-sm font-medium"> Read → </a>
                        </div>
                    </div>
                </article>
            </div>
        </section>

        <!-- Newsletter Subscription -->
        <section class="mt-16 bg-indigo-50 rounded-lg p-8">
            <div class="max-w-2xl mx-auto text-center">
                <h2 class="text-2xl font-bold text-gray-800 mb-2">Subscribe to our Newsletter</h2>
                <p class="text-gray-600 mb-6">Get the latest articles and news delivered to your inbox.</p>
                <div class="flex flex-col sm:flex-row gap-2">
                    <input type="email" placeholder="Enter your email" class="flex-grow px-4 py-2 rounded-md border border-gray-300 focus:outline-none focus:ring-2 focus:ring-indigo-500" v-model="email" />
                    <button class="px-6 py-2 bg-indigo-600 text-white rounded-md hover:bg-indigo-700 transition" @click="subscribe">Subscribe</button>
                </div>
            </div>
        </section>
    </div>
</template>

<script setup>
import { ref } from 'vue';

// Data
const email = ref('');
const featuredPost = {
    id: 1,
    title: 'Getting Started with Vue.js 3 and Composition API',
    excerpt: 'Learn how to build modern web applications with Vue.js 3 and the new Composition API in this comprehensive guide.',
    category: 'Vue.js',
    date: '2023-05-15',
    image: 'https://images.unsplash.com/photo-1633356122544-f134324a6cee?ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D&auto=format&fit=crop&w=1470&q=80',
    author: {
        name: 'Jane Doe',
        avatar: 'https://randomuser.me/api/portraits/women/44.jpg'
    }
};

const posts = [
    {
        id: 2,
        title: 'Mastering TailwindCSS for Rapid UI Development',
        excerpt: 'Discover how TailwindCSS can speed up your development workflow and help you build beautiful interfaces.',
        category: 'CSS',
        date: '2023-05-10',
        image: 'https://images.unsplash.com/photo-1633356122102-3fe601e05bd2?ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D&auto=format&fit=crop&w=1470&q=80',
        author: {
            name: 'John Smith',
            avatar: 'https://randomuser.me/api/portraits/men/32.jpg'
        }
    },
    {
        id: 3,
        title: 'Building RESTful APIs with Node.js and Express',
        excerpt: 'A step-by-step guide to creating robust and scalable REST APIs using Node.js and Express framework.',
        category: 'Backend',
        date: '2023-05-05',
        image: 'https://images.unsplash.com/photo-1555066931-4365d14bab8c?ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D&auto=format&fit=crop&w=1470&q=80',
        author: {
            name: 'Alex Johnson',
            avatar: 'https://randomuser.me/api/portraits/men/75.jpg'
        }
    },
    {
        id: 4,
        title: 'Responsive Web Design Principles in 2023',
        excerpt: 'Learn the latest techniques and best practices for creating websites that work on any device.',
        category: 'Design',
        date: '2023-04-28',
        image: 'https://images.unsplash.com/photo-1547658719-da2b51169166?ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D&auto=format&fit=crop&w=1528&q=80',
        author: {
            name: 'Sarah Williams',
            avatar: 'https://randomuser.me/api/portraits/women/68.jpg'
        }
    }
];

// Methods
const formatDate = (dateString) => {
    const options = { year: 'numeric', month: 'long', day: 'numeric' };
    return new Date(dateString).toLocaleDateString('en-US', options);
};

const subscribe = () => {
    if (email.value) {
        alert(`Thank you for subscribing with ${email.value}!`);
        email.value = '';
    } else {
        alert('Please enter your email address');
    }
};
</script>
