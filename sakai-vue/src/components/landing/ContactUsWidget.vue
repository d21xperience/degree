<script setup>
import { Facebook, Instagram, Linkedin, Mail, MapPin, Phone, Twitter } from 'lucide-vue-next';
import { computed, onMounted, reactive, ref } from 'vue';
// Mock icons - in real app, import from lucide-vue-next or similar

// Emits
const emit = defineEmits(['form-submit', 'form-success', 'form-error']);

// Reactive data
const sectionData = reactive({
    title: 'Hubungi Kami',
    subtitle: 'Ada pertanyaan? Jangan ragu untuk menghubungi kami.'
});

const formData = reactive({
    name: '',
    email: '',
    subject: '',
    message: ''
});

const formErrors = reactive({
    name: '',
    email: '',
    subject: '',
    message: ''
});

const subjectOptions = reactive([
    { value: 'general', label: 'Pertanyaan Umum' },
    { value: 'technical', label: 'Bantuan Teknis' },
    { value: 'partnership', label: 'Kemitraan' },
    { value: 'feedback', label: 'Saran & Masukan' },
    { value: 'other', label: 'Lainnya' }
]);

const contactInfo = reactive([
    {
        icon: Mail,
        label: 'Email',
        value: 'info@verifikasiijazah.id'
    },
    {
        icon: Phone,
        label: 'Telepon',
        value: '+62 21 1234 5678'
    },
    {
        icon: MapPin,
        label: 'Alamat',
        value: 'Jakarta, Indonesia'
    }
]);

const socialLinks = reactive([
    { name: 'Facebook', icon: Facebook, url: 'https://facebook.com' },
    { name: 'Twitter', icon: Twitter, url: 'https://twitter.com' },
    { name: 'Instagram', icon: Instagram, url: 'https://instagram.com' },
    { name: 'LinkedIn', icon: Linkedin, url: 'https://linkedin.com' }
]);

// Component state
const isSubmitting = ref(false);
const showSuccessMessage = ref(false);

// Computed
const messageLength = computed(() => formData.message.length);

const isFormValid = computed(() => {
    return formData.name.trim() && formData.email.trim() && formData.subject && formData.message.trim() && !Object.values(formErrors).some((error) => error);
});

// Methods
const validateField = (fieldName) => {
    formErrors[fieldName] = '';

    switch (fieldName) {
        case 'name':
            if (!formData.name.trim()) {
                formErrors.name = 'Nama wajib diisi';
            } else if (formData.name.trim().length < 2) {
                formErrors.name = 'Nama minimal 2 karakter';
            }
            break;

        case 'email':
            if (!formData.email.trim()) {
                formErrors.email = 'Email wajib diisi';
            } else if (!/^[^\s@]+@[^\s@]+\.[^\s@]+$/.test(formData.email)) {
                formErrors.email = 'Format email tidak valid';
            }
            break;

        case 'subject':
            if (!formData.subject) {
                formErrors.subject = 'Subjek wajib dipilih';
            }
            break;

        case 'message':
            if (!formData.message.trim()) {
                formErrors.message = 'Pesan wajib diisi';
            } else if (formData.message.trim().length < 10) {
                formErrors.message = 'Pesan minimal 10 karakter';
            } else if (formData.message.length > 500) {
                formErrors.message = 'Pesan maksimal 500 karakter';
            }
            break;
    }
};

const validateForm = () => {
    Object.keys(formData).forEach((field) => {
        validateField(field);
    });

    return isFormValid.value;
};

const resetForm = () => {
    Object.keys(formData).forEach((key) => {
        formData[key] = '';
    });
    Object.keys(formErrors).forEach((key) => {
        formErrors[key] = '';
    });
};

const handleSubmit = async () => {
    if (!validateForm()) {
        return;
    }

    isSubmitting.value = true;

    try {
        // Simulate API call
        await new Promise((resolve) => setTimeout(resolve, 2000));

        emit('form-submit', { ...formData });

        showSuccessMessage.value = true;
        resetForm();

        // Hide success message after 5 seconds
        setTimeout(() => {
            showSuccessMessage.value = false;
        }, 5000);

        emit('form-success');
    } catch (error) {
        console.error('Form submission error:', error);
        emit('form-error', error);
    } finally {
        isSubmitting.value = false;
    }
};

// Lifecycle
onMounted(() => {
    console.log('ContactSection mounted');
});
</script>

<template>
    <section id="contact" class="contact-section" role="region" aria-labelledby="contact-title">
        <div class="container">
            <header class="section-header">
                <h2 id="contact-title" class="section-title">
                    {{ sectionData.title }}
                </h2>
                <p class="section-subtitle">
                    {{ sectionData.subtitle }}
                </p>
            </header>

            <div class="contact-content">
                <div class="contact-form-container">
                    <form @submit.prevent="handleSubmit" class="contact-form" novalidate>
                        <div class="form-group">
                            <label for="contact-name" class="form-label"> Nama Lengkap </label>
                            <input
                                id="contact-name"
                                v-model="formData.name"
                                type="text"
                                class="form-input"
                                :class="{ error: formErrors.name }"
                                placeholder="Masukkan nama lengkap Anda"
                                :disabled="isSubmitting"
                                @input="validateField('name')"
                                @blur="validateField('name')"
                                aria-describedby="name-error"
                                required
                            />
                            <div id="name-error" class="form-error" v-if="formErrors.name" role="alert">
                                {{ formErrors.name }}
                            </div>
                        </div>

                        <div class="form-group">
                            <label for="contact-email" class="form-label"> Email </label>
                            <input
                                id="contact-email"
                                v-model="formData.email"
                                type="email"
                                class="form-input"
                                :class="{ error: formErrors.email }"
                                placeholder="Masukkan alamat email Anda"
                                :disabled="isSubmitting"
                                @input="validateField('email')"
                                @blur="validateField('email')"
                                aria-describedby="email-error"
                                required
                            />
                            <div id="email-error" class="form-error" v-if="formErrors.email" role="alert">
                                {{ formErrors.email }}
                            </div>
                        </div>

                        <div class="form-group">
                            <label for="contact-subject" class="form-label"> Subjek </label>
                            <select id="contact-subject" v-model="formData.subject" class="form-select" :class="{ error: formErrors.subject }" :disabled="isSubmitting" @change="validateField('subject')" aria-describedby="subject-error" required>
                                <option value="">Pilih subjek pesan</option>
                                <option v-for="(option, index) in subjectOptions" :key="index" :value="option.value">
                                    {{ option.label }}
                                </option>
                            </select>
                            <div id="subject-error" class="form-error" v-if="formErrors.subject" role="alert">
                                {{ formErrors.subject }}
                            </div>
                        </div>

                        <div class="form-group">
                            <label for="contact-message" class="form-label"> Pesan </label>
                            <textarea
                                id="contact-message"
                                v-model="formData.message"
                                class="form-textarea"
                                :class="{ error: formErrors.message }"
                                placeholder="Tulis pesan Anda di sini..."
                                rows="5"
                                :disabled="isSubmitting"
                                @input="validateField('message')"
                                @blur="validateField('message')"
                                aria-describedby="message-error message-counter"
                                required
                            ></textarea>
                            <div class="form-meta">
                                <div id="message-counter" class="character-counter" :class="{ warning: messageLength > 450 }">{{ messageLength }}/500 karakter</div>
                                <div id="message-error" class="form-error" v-if="formErrors.message" role="alert">
                                    {{ formErrors.message }}
                                </div>
                            </div>
                        </div>

                        <div class="form-actions">
                            <button type="submit" class="btn-submit" :disabled="!isFormValid || isSubmitting">
                                <span v-if="!isSubmitting">Kirim Pesan</span>
                                <span v-else class="loading-content">
                                    <svg class="loading-spinner" viewBox="0 0 24 24">
                                        <circle cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4" fill="none" opacity="0.25" />
                                        <path fill="currentColor" opacity="0.75" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z" />
                                    </svg>
                                    Mengirim...
                                </span>
                            </button>
                        </div>
                    </form>

                    <!-- Success Message -->
                    <div v-if="showSuccessMessage" class="success-message" role="alert" aria-live="polite">
                        <svg class="success-icon" viewBox="0 0 24 24" fill="none" stroke="currentColor">
                            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7"></path>
                        </svg>
                        <div>
                            <h3>Pesan Berhasil Dikirim!</h3>
                            <p>Terima kasih atas pesan Anda. Tim kami akan merespons dalam 1-2 hari kerja.</p>
                        </div>
                    </div>
                </div>

                <!-- Contact Information -->
                <div class="contact-info">
                    <h3 class="info-title">Informasi Kontak</h3>
                    <div class="info-items">
                        <div v-for="(info, index) in contactInfo" :key="index" class="info-item">
                            <div class="info-icon" aria-hidden="true">
                                <component :is="info.icon" />
                            </div>
                            <div class="info-content">
                                <h4 class="info-label">{{ info.label }}</h4>
                                <p class="info-value">{{ info.value }}</p>
                            </div>
                        </div>
                    </div>

                    <div class="social-links">
                        <h4 class="social-title">Ikuti Kami</h4>
                        <div class="social-icons">
                            <a v-for="(social, index) in socialLinks" :key="index" :href="social.url" class="social-link" :aria-label="`Ikuti kami di ${social.name}`" target="_blank" rel="noopener noreferrer">
                                <component :is="social.icon" />
                            </a>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    </section>
</template>

<style scoped>
.contact-section {
    padding: 5rem 0;
    background-color: #f9fafb;
}

.container {
    max-width: 1200px;
    margin: 0 auto;
    padding: 0 2rem;
}

.section-header {
    text-align: center;
    margin-bottom: 4rem;
}

.section-title {
    color: #1e3a8a;
    font-family: 'Poppins', sans-serif;
    font-weight: 700;
    font-size: 2.5rem;
    margin-bottom: 1rem;
    line-height: 1.2;
}

.section-subtitle {
    color: #6b7280;
    font-family: 'Inter', sans-serif;
    font-size: 1.25rem;
    max-width: 600px;
    margin: 0 auto;
    line-height: 1.6;
}

.contact-content {
    display: grid;
    grid-template-columns: 2fr 1fr;
    gap: 4rem;
    align-items: start;
}

.contact-form-container {
    background: white;
    padding: 2.5rem;
    border-radius: 12px;
    box-shadow: 0 4px 6px rgba(0, 0, 0, 0.05);
    border: 1px solid #e5e7eb;
}

.contact-form {
    display: flex;
    flex-direction: column;
    gap: 1.5rem;
}

.form-group {
    display: flex;
    flex-direction: column;
    gap: 0.5rem;
}

.form-label {
    font-weight: 500;
    color: #374151;
    font-family: 'Inter', sans-serif;
    font-size: 0.875rem;
}

.form-input,
.form-select,
.form-textarea {
    padding: 0.75rem 1rem;
    border: 2px solid #d1d5db;
    border-radius: 8px;
    font-size: 1rem;
    transition:
        border-color 0.2s ease,
        box-shadow 0.2s ease;
    font-family: 'Inter', sans-serif;
    background: white;
}

.form-input:focus,
.form-select:focus,
.form-textarea:focus {
    outline: none;
    border-color: #1e3a8a;
    box-shadow: 0 0 0 3px rgba(30, 58, 138, 0.1);
}

.form-input.error,
.form-select.error,
.form-textarea.error {
    border-color: #ef4444;
}

.form-input:disabled,
.form-select:disabled,
.form-textarea:disabled {
    background: #f9fafb;
    cursor: not-allowed;
    opacity: 0.6;
}

.form-textarea {
    resize: vertical;
    min-height: 120px;
}

.form-meta {
    display: flex;
    justify-content: space-between;
    align-items: center;
    gap: 1rem;
}

.character-counter {
    font-size: 0.75rem;
    color: #6b7280;
    font-family: 'Inter', sans-serif;
}

.character-counter.warning {
    color: #f59e0b;
}

.form-error {
    font-size: 0.875rem;
    color: #ef4444;
    font-family: 'Inter', sans-serif;
}

.form-actions {
    margin-top: 1rem;
}

.btn-submit {
    width: 100%;
    padding: 0.875rem 1.5rem;
    background: linear-gradient(135deg, #1e3a8a, #3b82f6);
    color: white;
    border: none;
    border-radius: 8px;
    font-weight: 600;
    cursor: pointer;
    transition: all 0.3s ease;
    font-family: 'Inter', sans-serif;
    font-size: 1rem;
    min-height: 50px;
    display: flex;
    align-items: center;
    justify-content: center;
    gap: 0.5rem;
}

.btn-submit:hover:not(:disabled) {
    transform: translateY(-1px);
    box-shadow: 0 4px 12px rgba(30, 58, 138, 0.3);
}

.btn-submit:disabled {
    background: #9ca3af;
    cursor: not-allowed;
    transform: none;
    box-shadow: none;
}

.loading-content {
    display: flex;
    align-items: center;
    gap: 0.5rem;
}

.loading-spinner {
    width: 20px;
    height: 20px;
    animation: spin 1s linear infinite;
}

.success-message {
    background: #f0fdf4;
    border: 1px solid #bbf7d0;
    border-radius: 8px;
    padding: 1.5rem;
    display: flex;
    gap: 1rem;
    align-items: flex-start;
    margin-top: 1rem;
}

.success-icon {
    width: 24px;
    height: 24px;
    color: #16a34a;
    flex-shrink: 0;
}

.success-message h3 {
    color: #166534;
    font-weight: 600;
    margin: 0 0 0.5rem 0;
    font-family: 'Inter', sans-serif;
}

.success-message p {
    color: #15803d;
    margin: 0;
    font-family: 'Inter', sans-serif;
    font-size: 0.875rem;
}

/* Contact Info Styles */
.contact-info {
    background: white;
    padding: 2rem;
    border-radius: 12px;
    box-shadow: 0 4px 6px rgba(0, 0, 0, 0.05);
    border: 1px solid #e5e7eb;
    height: fit-content;
}

.info-title {
    color: #1e3a8a;
    font-family: 'Poppins', sans-serif;
    font-weight: 600;
    font-size: 1.25rem;
    margin-bottom: 1.5rem;
}

.info-items {
    display: flex;
    flex-direction: column;
    gap: 1.5rem;
    margin-bottom: 2rem;
}

.info-item {
    display: flex;
    gap: 1rem;
    align-items: flex-start;
}

.info-icon {
    width: 20px;
    height: 20px;
    color: #1e3a8a;
    flex-shrink: 0;
    margin-top: 2px;
}

.info-content {
    flex: 1;
}

.info-label {
    font-weight: 600;
    color: #374151;
    margin: 0 0 0.25rem 0;
    font-family: 'Inter', sans-serif;
    font-size: 0.875rem;
}

.info-value {
    color: #6b7280;
    margin: 0;
    font-family: 'Inter', sans-serif;
    font-size: 0.875rem;
}

.social-links {
    border-top: 1px solid #e5e7eb;
    padding-top: 1.5rem;
}

.social-title {
    color: #374151;
    font-family: 'Inter', sans-serif;
    font-weight: 600;
    font-size: 1rem;
    margin-bottom: 1rem;
}

.social-icons {
    display: flex;
    gap: 1rem;
}

.social-link {
    width: 40px;
    height: 40px;
    background: #f3f4f6;
    border-radius: 8px;
    display: flex;
    align-items: center;
    justify-content: center;
    color: #6b7280;
    transition: all 0.3s ease;
    text-decoration: none;
}

.social-link:hover {
    background: #1e3a8a;
    color: white;
    transform: translateY(-2px);
}

.social-link svg {
    width: 20px;
    height: 20px;
}

/* Animations */
@keyframes spin {
    from {
        transform: rotate(0deg);
    }
    to {
        transform: rotate(360deg);
    }
}

/* Responsive Design */
@media (max-width: 1024px) {
    .contact-content {
        grid-template-columns: 1fr;
        gap: 2rem;
    }
}

@media (max-width: 768px) {
    .container {
        padding: 0 1rem;
    }

    .section-title {
        font-size: 2rem;
    }

    .section-subtitle {
        font-size: 1.1rem;
    }

    .contact-form-container,
    .contact-info {
        padding: 1.5rem;
    }

    .form-meta {
        flex-direction: column;
        align-items: flex-start;
    }
}

@media (max-width: 480px) {
    .section-title {
        font-size: 1.75rem;
    }

    .contact-form-container,
    .contact-info {
        padding: 1rem;
    }

    .social-icons {
        justify-content: center;
    }
}

/* Accessibility improvements */
@media (prefers-reduced-motion: reduce) {
    .btn-submit,
    .social-link,
    .loading-spinner {
        transition: none;
        animation: none;
    }
}

/* High contrast mode support */
@media (prefers-contrast: high) {
    .form-input,
    .form-select,
    .form-textarea {
        border-color: #000;
    }

    .btn-submit {
        background: #000;
    }

    .section-title {
        color: #000;
    }
}
</style>
