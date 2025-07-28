<template>
    <!-- Featured Post -->
    <section class="mb-12">
        <div class="bg-white rounded-lg shadow-md overflow-hidden">
            <img :src="featuredPost.image" alt="Featured Post" class="w-full h-64 object-cover" />
            <div class="p-6">
                <span class="inline-block px-3 py-1 bg-indigo-100 text-indigo-600 rounded-full text-sm font-medium mb-2">
                    {{ featuredPost.category }}
                </span>
                <h2 class="text-2xl font-bold text-gray-800 mb-2">{{ featuredPost.title }}</h2>
                <p class="text-gray-600 mb-4">
                    {{ featuredPost.excerpt }}
                    <span v-if="!expandedPosts[featuredPost.id]">...</span>
                    <span v-if="expandedPosts[featuredPost.id]">{{ featuredPost.fullContent }}</span>
                </p>
                <div class="flex items-center">
                    <img :src="featuredPost.author.avatar" alt="Author" class="w-10 h-10 rounded-full mr-3" />
                    <div>
                        <p class="text-sm font-medium text-gray-900">{{ featuredPost.author.name }}</p>
                        <p class="text-sm text-gray-500">{{ formatDate(featuredPost.date) }}</p>
                    </div>
                    <button @click="toggleExpand(featuredPost.id)" class="ml-auto px-4 py-2 bg-indigo-600 text-white rounded-md hover:bg-indigo-700 transition">
                        {{ expandedPosts[featuredPost.id] ? 'Read Less' : 'Read More' }}
                    </button>
                </div>
            </div>
        </div>
    </section>
</template>

<script setup>
import { reactive, ref } from 'vue';

// Data
const email = ref('');
const expandedPosts = reactive({});

// Post data dengan konten lengkap
const featuredPost = {
    id: 1,
    title: 'Getting Started with Vue.js 3 and Composition API',
    excerpt: 'Learn how to build modern web applications with Vue.js 3 and the new Composition API in this comprehensive guide.',
    fullContent:
        'Vue.js 3 introduces the Composition API as an alternative to the Options API. This new API allows for better code organization and reuse through composable functions. In this article, we will explore the basics of setting up a Vue 3 project, creating components with the Composition API, and how to leverage reactive references and computed properties. We will also cover lifecycle hooks in the context of the Composition API and demonstrate how to extract reusable logic into composable functions. By the end of this guide, you will have a solid foundation to start building applications with Vue 3 and the Composition API.',
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
        fullContent:
            'TailwindCSS is a utility-first CSS framework that enables rapid UI development without leaving your HTML. Unlike traditional CSS frameworks, Tailwind provides low-level utility classes that let you build completely custom designs. In this article, we will cover the core concepts of TailwindCSS, including its utility classes, responsive design approach, and customization options. We will also explore best practices for organizing your Tailwind projects and how to integrate it with popular frameworks like Vue.js and React. By the end, you will understand how to leverage TailwindCSS to create maintainable, responsive designs quickly.',
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
        fullContent:
            'Node.js with Express is a powerful combination for building RESTful APIs. In this guide, we will walk through the process of setting up a Node.js project, installing Express, and creating a complete REST API with CRUD operations. We will cover important topics like middleware, routing, error handling, and data validation. Additionally, we will explore how to connect to a database (MongoDB), implement authentication with JWT, and document your API with Swagger. By following this tutorial, you will have a production-ready API that follows REST conventions and best practices.',
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
        fullContent:
            'Responsive web design has evolved significantly since its inception. In 2023, we have new challenges like foldable devices, variable viewport sizes, and performance considerations. This article covers modern responsive design techniques including CSS Grid, Flexbox, container queries, and the clamp() function for fluid typography. We will also discuss mobile-first vs. desktop-first approaches, responsive images with srcset, and how to test your designs across various devices. With these up-to-date techniques, you can create websites that provide an optimal viewing experience across all devices and screen sizes.',
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

const toggleExpand = (postId) => {
    expandedPosts[postId] = !expandedPosts[postId];
};
</script>
