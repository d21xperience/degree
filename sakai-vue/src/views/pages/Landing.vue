<script setup lang="ts">
import AboutUsWidget from '@/components/AboutUsWidget.vue';
import CallToActionWidget from '@/components/landing/CallToActionWidget.vue';
import ContactUsWidget from '@/components/landing/ContactUsWidget.vue';
import FooterWidget from '@/components/landing/FooterWidget.vue';
import HowItWorksWidget from '@/components/landing/HowItWorksWidget.vue';
import MemberHaveJoined from '@/components/landing/MemberHaveJoined.vue';
import NavbarWidget from '@/components/landing/NavbarWidget.vue';
import TestimonialsWidget from '@/components/landing/TestimonialsWidget.vue';
import WhatsAppFloat from '@/components/landing/WhatsAppFloat.vue';
import WhyTrustUsWidget from '@/components/landing/WhyTrustUsWidget.vue';
</script>

<template>
    <div class="bg-surface-0 dark:bg-surface-900">
        <div id="home" class="landing-wrapper overflow-hidden">
            <NavbarWidget />
            <HeroSection />
            <!-- <HerosWidget /> -->
            <!-- <FeaturesWidget class="z-50" /> -->
            <HowItWorksWidget class="my-5" />
            <AboutUsWidget />
            <WhyTrustUsWidget />
            <MemberHaveJoined />
            <TestimonialsWidget />
            <ContactUsWidget />
            <CallToActionWidget />
            <WhatsAppFloat phone="6289602897265" />
            <FooterWidget />
        </div>
    </div>
</template>
