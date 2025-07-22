<template>
    <section class="testimonials-section">
        <h2 class="testimonials-title">Apa Kata Mereka?</h2>

        <div class="testimonials-container">
            <div v-for="(testimonial, index) in testimonials" :key="index" class="testimonial-card">
                <div class="testimonial-photo">
                    <img :src="testimonial.photo" :alt="testimonial.name" />
                </div>
                <div class="testimonial-quote">
                    <p>"{{ testimonial.quote }}"</p>
                </div>
                <div class="testimonial-name">{{ testimonial.name }}</div>
                <div class="testimonial-role">{{ testimonial.role }}</div>
            </div>
        </div>
    </section>
</template>

<script setup>
const testimonials = [
    {
        photo: 'https://images.unsplash.com/photo-1573497019940-1c28c88b4f3e?w=200',
        name: 'Budi Santoso',
        role: 'Admin Sekolah',
        quote: 'Sistem ini memudahkan kami memverifikasi ijazah palsu.'
    },
    {
        photo: 'https://images.unsplash.com/photo-1507003211169-0a1dd7228f2d?w=200',
        name: 'Anita Rahayu',
        role: 'Siswa',
        quote: 'Lamar kerja lebih percaya diri dengan ijazah terverifikasi!'
    }
];
</script>

<style scoped>
.testimonials-section {
    padding: 4rem 1rem;
    max-width: 1200px;
    margin: 0 auto;
}

.testimonials-title {
    text-align: center;
    color: #2563eb;
    font-size: 2rem;
    margin-bottom: 3rem;
}

.testimonials-container {
    display: grid;
    grid-template-columns: repeat(auto-fit, minmax(300px, 1fr));
    gap: 2rem;
    padding: 0 1rem;
}

.testimonial-card {
    background: white;
    border-radius: 12px;
    padding: 2rem;
    text-align: center;
    box-shadow: 0 4px 6px rgba(0, 0, 0, 0.05);
    transition: transform 0.3s ease;
}

.testimonial-card:hover {
    transform: translateY(-5px);
}

.testimonial-photo {
    width: 80px;
    height: 80px;
    border-radius: 50%;
    overflow: hidden;
    margin: 0 auto 1rem;
    border: 3px solid #ffd700;
}

.testimonial-photo img {
    width: 100%;
    height: 100%;
    object-fit: cover;
}

.testimonial-quote {
    position: relative;
    padding: 1.5rem;
    margin: 1rem 0;
    background: #f8fafc;
    border-radius: 8px;
    border-left: 4px solid #4caf50;
}

.testimonial-quote::before {
    content: '"';
    font-size: 3rem;
    color: #ffd700;
    position: absolute;
    left: 10px;
    top: -10px;
    opacity: 0.3;
}

.testimonial-quote p {
    font-style: italic;
    color: #4b5563;
    line-height: 1.6;
}

.testimonial-name {
    font-weight: bold;
    color: #1e40af;
    margin-top: 1rem;
}

.testimonial-role {
    color: #64748b;
    font-size: 0.9rem;
}

@media (max-width: 640px) {
    .testimonials-container {
        grid-template-columns: 1fr;
    }

    .testimonial-card {
        padding: 1.5rem;
    }
}
</style>
