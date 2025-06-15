<template>
  <div class="min-h-screen flex flex-col bg-amber-50 text-gray-900 antialiased">
    <AppNavbar />

    <section class="relative flex-grow flex items-center justify-center overflow-hidden px-4 py-24 sm:py-32">
      <div aria-hidden class="absolute inset-0">
        <div class="absolute -top-20 -left-20 w-72 h-72 bg-amber-200 rounded-full mix-blend-multiply filter blur-3xl opacity-30 animate-blob"></div>
        <div class="absolute bottom-0 -right-10 w-96 h-96 bg-green-200 rounded-full mix-blend-multiply filter blur-2xl opacity-20 animate-blob animation-delay-400"></div>
      </div>

      <div class="relative z-10 max-w-4xl text-center space-y-8">
        <h1 class="text-5xl sm:text-6xl lg:text-7xl font-extrabold leading-tight animate-fade-in-up">
          Elevate Your <span class="text-green-600 drop-shadow-md">E-commerce</span> Game
        </h1>
        <p class="text-xl sm:text-2xl text-gray-700 max-w-3xl mx-auto animate-fade-in-up animation-delay-200">
          Focus on your passion—**Seller’s Dash** handles the heavy lifting so you can effortlessly **build, manage, and scale** your online business.
        </p>
        <div class="flex flex-col sm:flex-row justify-center gap-4 animate-fade-in-up animation-delay-400">
          <router-link
            to="/register"
            class="inline-flex items-center justify-center px-8 py-3 bg-green-600 text-white text-lg font-bold rounded-full shadow-xl hover:bg-green-700 transition transform hover:-translate-y-1"
          >
            Start Free Trial
          </router-link>
          <router-link
            to="#features"
            class="inline-flex items-center justify-center px-8 py-3 bg-white text-green-700 text-lg font-bold rounded-full border border-green-300 shadow-lg hover:bg-green-50 transition transform hover:scale-105"
          >
            View Features
          </router-link>
        </div>
      </div>
    </section>

    <section id="features" class="py-20 sm:py-28 lg:py-36 px-4 bg-white">
      <div class="max-w-7xl mx-auto">
        <div class="text-center mb-16 animate-fade-in">
          <h2 class="text-4xl sm:text-5xl font-extrabold text-gray-900">
            Everything You Need to <span class="text-green-600">Thrive</span> Online
          </h2>
          <p class="mt-4 text-lg text-gray-600 max-w-2xl mx-auto">
            From dazzling storefronts to seamless analytics, we’ve got your entire sales journey covered.
          </p>
        </div>
        <div class="grid gap-10 md:grid-cols-2 lg:grid-cols-3">
          <FeatureCard
            icon="M4 5a1 1 0 011-1h14a1 1 0 011 1v2a1 1 0 01-1 1H5a1 1 0 01-1-1V5zM4 13a1 1 0 011-1h6a1 1 0 011 1v6a1 1 0 01-1 1H5a1 1 0 01-1-1v-6zM16 13a1 1 0 011-1h2a1 1 0 011 1v6a1 1 0 01-1 1h-2a1 1 0 01-1-1v-6z"
            title="Stunning Storefronts"
            text="Design a captivating online presence with fully customizable themes that truly represent your brand’s unique identity."
            color="green"
            :delay="0"
          />
          <FeatureCard
            icon="M3 10h18M7 15h1m4 0h1m-7 4h12a3 3 0 003-3V8a3 3 0 00-3-3H6a3 3 0 00-3 3v8a3 3 0 003 3z"
            title="Seamless Payments"
            text="Effortlessly accept secure payments worldwide, supporting all major credit cards and popular methods."
            color="amber"
            :delay="200"
          />
          <FeatureCard
            icon="M9 19v-6a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2a2 2 0 002-2zm0 0V9a2 2 0 012-2h2a2 2 0 012 2v10m-6 0a2 2 0 002 2h2a2 2 0 002-2m0 0V5a2 2 0 012-2h2a2 2 0 012 2v14a2 2 0 01-2 2h-2a2 2 0 01-2-2z"
            title="Actionable Analytics"
            text="Unlock powerful insights into your sales, customer behavior, and traffic to make data-driven decisions that fuel growth."
            color="green"
            :delay="400"
          />
        </div>
      </div>
    </section>

    <AppFooter />
  </div>
</template>

<script setup>
import FeatureCard from '@/components/FeatureCard.vue'
import AppNavbar   from '@/components/AppNavbar.vue'
import AppFooter   from '@/components/AppFooter.vue'
</script>

<style>
body {
  font-family: 'Inter', sans-serif;
}

@keyframes fadeIn {
  from { opacity: 0; transform: translateY(20px); }
  to   { opacity: 1; transform: translateY(0); }
}

@keyframes fadeInUp {
  from { opacity: 0; transform: translateY(40px); }
  to   { opacity: 1; transform: translateY(0); }
}

.animate-fade-in {
  animation: fadeIn 0.8s ease-out forwards;
}

.animate-fade-in-up {
  animation: fadeInUp 0.8s ease-out forwards;
}

.animation-delay-200 { animation-delay: 0.2s; }
.animation-delay-400 { animation-delay: 0.4s; }

@keyframes blob {
  0%, 100% { transform: scale(1) translate(0,0); }
  33% { transform: scale(1.2) translate(30px,-20px); }
  66% { transform: scale(0.8) translate(-20px,30px); }
}
.animate-blob {
  animation: blob 8s infinite ease-in-out;
}
</style>