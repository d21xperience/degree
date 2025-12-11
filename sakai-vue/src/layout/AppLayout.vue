<script setup>
defineProps(['sekolah']);
import { useLayout } from '@/layout/composables/layout';
import { computed, ref, watch } from 'vue';
import AppFooter from './AppFooter.vue';
import AppSidebar from './AppSidebar.vue';
import AppTopbar from './AppTopbar.vue';

const { layoutConfig, layoutState, isSidebarActive } = useLayout();

const outsideClickListener = ref(null);

watch(isSidebarActive, (newVal) => {
    if (newVal) {
        bindOutsideClickListener();
    } else {
        unbindOutsideClickListener();
    }
});

const containerClass = computed(() => {
    return {
        'layout-overlay': layoutConfig.menuMode === 'overlay',
        'layout-static': layoutConfig.menuMode === 'static',
        'layout-static-inactive': layoutState.staticMenuDesktopInactive && layoutConfig.menuMode === 'static',
        'layout-overlay-active': layoutState.overlayMenuActive,
        'layout-mobile-active': layoutState.staticMenuMobileActive
    };
});

function bindOutsideClickListener() {
    if (!outsideClickListener.value) {
        outsideClickListener.value = (event) => {
            if (isOutsideClicked(event)) {
                layoutState.overlayMenuActive = false;
                layoutState.staticMenuMobileActive = false;
                layoutState.menuHoverActive = false;
            }
        };
        document.addEventListener('click', outsideClickListener.value);
    }
}

function unbindOutsideClickListener() {
    if (outsideClickListener.value) {
        document.removeEventListener('click', outsideClickListener);
        outsideClickListener.value = null;
    }
}

function isOutsideClicked(event) {
    const sidebarEl = document.querySelector('.layout-sidebar');
    const topbarEl = document.querySelector('.layout-menu-button');

    return !(sidebarEl.isSameNode(event.target) || sidebarEl.contains(event.target) || topbarEl.isSameNode(event.target) || topbarEl.contains(event.target));
}

// onMounted(() => {
//     window.addEventListener('idle-configure', (e) => {
//         console.log('[IDLE] Config received:', e.detail);
//         // if (this.idle && this.idle.configure) {
//         //     this.idle.configure(e.detail.timeout); // ✅ kirim ke composable
//         // }
//     });
// });
// 🔑 State untuk UI warning
// const isIdleWarning = ref(false);
// const idleCountdown = ref(0);

// // 🛠 Inisialisasi composable
// const idle = useIdleDetection({
//     timeout: 25 * 60 * 1000, // default 25 menit
//     warningThreshold: 60 * 1000, // 60 detik warning
//     onWarning: (count) => {
//         isIdleWarning.value = true;
//         idleCountdown.value = count;
//     },
//     onIdle: () => {
//         isIdleWarning.value = false;
//         // Logout via store
//         store.dispatch('authService/logoutDueToIdle');
//     },
//     onActivity: () => {
//         isIdleWarning.value = false;
//     }
// });

// // 📦 Vuex Store
// const store = useStore();

// // 🔄 Reset idle detection saat user aktif
// const resetIdle = () => {
//     idle.reset();
//     isIdleWarning.value = false;
// };

// 🎯 Setup event listener & reaktivitas
// onMounted(() => {
//     // 1. Dengarkan event 'idle-configure' dari store
//     const handleIdleConfigure = (e) => {
//         if (e.detail?.timeout) {
//             idle.configure(e.detail.timeout); // ✅ kirim ke composable
//         }
//     };

//     // 2. Dengarkan event 'user-active' (misal: dari navbar)
//     const handleUserActive = () => resetIdle();

//     window.addEventListener('idle-configure', handleIdleConfigure);
//     window.addEventListener('user-active', handleUserActive);

//     // 3. Mulai idle detection saat user login
//     const stopWatching = watch(
//         () => store.state.authService.isAuthenticated,
//         (isAuth) => {
//             if (isAuth) {
//                 // Pastikan expiresAt tersedia
//                 const expiresAt = store.state.authService.expiresAt;
//                 if (expiresAt) {
//                     const timeLeft = expiresAt - Date.now();
//                     if (timeLeft > 0) {
//                         const timeout = Math.min(timeLeft * 0.9, 25 * 60 * 1000);
//                         idle.configure(timeout);
//                     }
//                 }
//             }
//         },
//         { immediate: true } // cek saat mount
//     );

//     // Cleanup
//     onUnmounted(() => {
//         window.removeEventListener('idle-configure', handleIdleConfigure);
//         window.removeEventListener('user-active', handleUserActive);
//         stopWatching();
//     });
// });
</script>

<template>
    <div class="layout-wrapper" :class="containerClass">
        <app-topbar />
        <app-sidebar />
        <!-- Modal Idle Warning (opsional) -->
        <!-- <Teleport to="body">
            <div v-if="isIdleWarning" class="idle-overlay" @click="resetIdle">
                <div class="idle-modal" @click.stop>
                    <i class="pi pi-info-circle" style="font-size: 2rem; color: #e0a800"></i>
                    <h3>Session Berakhir</h3>
                    <p>
                        Akan logout otomatis dalam <strong>{{ idleCountdown }} detik</strong> karena tidak ada aktivitas.
                    </p>
                    <button class="p-button p-button-warning" @click="resetIdle">Lanjutkan Sesi</button>
                </div>
            </div>
        </Teleport> -->

        <div class="layout-main-container">
            <div class="layout-main">
                <router-view />
            </div>
            <app-footer />
        </div>
        <div class="layout-mask animate-fadein"></div>
    </div>
    <Toast />
</template>
