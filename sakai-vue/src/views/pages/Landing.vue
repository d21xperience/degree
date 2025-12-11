<script setup lang="ts">
import NavbarWidget from '@/components/landing/NavbarWidget.vue';
</script>

<template>
    <div class="bg-surface-0 dark:bg-surface-900">
        <div id="home" class="landing-wrapper overflow-hidden">
            <NavbarWidget />
            <!-- <HeroSection /> -->
            <!-- <HerosWidget /> -->
            <!-- <FeaturesWidget class="z-50" /> -->
            <HowItWorksWidget class="my-5" />
            <WhyTrustUsWidget />
            <MemberHaveJoined />
            <AboutUsWidget />
            <TestimonialsWidget />
            <ContactUsWidget />
            <CallToActionWidget />
            <WhatsAppFloat phone="6289602897265" />
            <FooterWidget />
        </div>
    </div>
</template>
