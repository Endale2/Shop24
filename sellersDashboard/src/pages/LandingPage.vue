<template>
  <div class="min-h-screen flex flex-col bg-gradient-to-br from-green-50 to-green-100 text-gray-900 antialiased">
    <header class="w-full sticky top-0 z-50 bg-white/70 backdrop-blur-md border-b border-green-200 shadow-sm">
      <div class="max-w-7xl mx-auto flex items-center justify-between p-4 lg:px-8">
        <router-link to="/" class="flex items-center space-x-3 group">
          <div class="p-2 bg-amber-500 rounded-full group-hover:rotate-12 transition-transform duration-500 shadow-md">
            <svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6 text-white" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 10V3L4 14h7v7l9-11h-7z"/>
            </svg>
          </div>
          <span class="text-2xl font-extrabold text-gray-800">Seller’s <span class="text-green-600">Dash</span></span>
        </router-link>
        <nav class="flex items-center space-x-6">
          <router-link to="/login" class="text-gray-700 font-medium hover:text-green-600 transition-colors duration-200">
            Log In
          </router-link>
          <router-link to="/register"
                       class="px-5 py-2 bg-green-600 text-white font-semibold rounded-full shadow-lg hover:bg-green-700 transition-all duration-300 transform hover:-translate-y-0.5">
            Get Started
          </router-link>
        </nav>
      </div>
    </header>

    <section class="relative flex-grow flex items-center justify-center overflow-hidden px-4 py-24 sm:py-32">
      <div aria-hidden class="absolute inset-0 z-0">
        <div class="absolute -top-20 -left-20 w-72 h-72 bg-green-200 rounded-full mix-blend-multiply filter blur-3xl opacity-30 animate-blob"></div>
        <div class="absolute bottom-0 -right-10 w-96 h-96 bg-amber-200 rounded-full mix-blend-multiply filter blur-2xl opacity-20 animate-blob animation-delay-400"></div>
      </div>

      <div class="relative z-10 max-w-4xl text-center space-y-8">
        <h1 class="text-5xl sm:text-6xl lg:text-7xl font-extrabold leading-tight text-gray-900 animate-fade-in-up">
          Elevate Your <span class="text-green-600 drop-shadow-md">E-commerce</span> Game
        </h1>
        <p class="text-xl sm:text-2xl text-gray-700 max-w-3xl mx-auto animate-fade-in-up animation-delay-200">
          Focus on your passion—**Seller’s Dash** handles the heavy lifting so you can effortlessly **build, manage, and scale** your online business.
        </p>
        <div class="flex flex-col sm:flex-row justify-center gap-4 animate-fade-in-up animation-delay-400">
          <router-link to="/register"
                       class="inline-flex items-center justify-center px-8 py-3 bg-green-600 text-white text-lg font-bold rounded-full shadow-xl hover:bg-green-700 hover:shadow-2xl transition-all duration-300 transform hover:-translate-y-1">
            Start Free Trial
          </router-link>
          <router-link to="/features"
                       class="inline-flex items-center justify-center px-8 py-3 bg-white text-green-700 text-lg font-bold rounded-full border border-green-300 shadow-lg hover:bg-green-50 transition-all duration-300 transform hover:scale-105">
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
            text="Effortlessly accept secure payments from customers worldwide, supporting all major credit cards and popular methods."
            color="amber"
            :delay="200"
          />
          <FeatureCard
            icon="M9 19v-6a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2a2 2 0 002-2zm0 0V9a2 2 0 012-2h2a2 2 0 012 2v10m-6 0a2 2 0 002 2h2a2 2 0 002-2m0 0V5a2 2 0 012-2h2a2 2 0 012 2v14a2 2 0 01-2 2h-2a2 2 0 01-2-2z"
            title="Actionable Analytics"
            text="Unlock powerful insights into your sales, customer behavior, and traffic to make data-driven decisions that fuel growth."
            color="lime"
            :delay="400"
          />
        </div>
      </div>
    </section>

    <footer class="bg-gray-900 text-gray-400 py-12">
      <div class="max-w-7xl mx-auto px-4 lg:px-8 text-center space-y-4">
        <p class="text-lg">&copy; {{ new Date().getFullYear() }} Seller’s Dash. All rights reserved.</p>
        <div class="flex justify-center space-x-8 text-md">
          <a href="#" class="hover:text-white transition-colors duration-200">Privacy Policy</a>
          <a href="#" class="hover:text-white transition-colors duration-200">Terms of Service</a>
          <a href="#" class="hover:text-white transition-colors duration-200">Support</a>
        </div>
      </div>
    </footer>
  </div>
</template>

<script setup>
import FeatureCard from '../components/FeatureCard.vue' 
</script>

<style>
/* Base font family for a modern feel */
body {
  font-family: 'Inter', sans-serif; /* Or any other clean sans-serif font */
}

/* Animations (from previous version, ensuring they still apply) */
@keyframes fadeIn {
  from {
    opacity: 0;
    transform: translateY(20px);
  }
  to {
    opacity: 1;
    transform: translateY(0);
  }
}

@keyframes fadeInUp {
  from {
    opacity: 0;
    transform: translateY(40px);
  }
  to {
    opacity: 1;
    transform: translateY(0);
  }
}

.animate-fade-in {
  animation: fadeIn 0.8s ease-out forwards;
}

.animate-fade-in-up {
  animation: fadeInUp 0.8s ease-out forwards;
}

/* Staggered animations */
.animation-delay-200 {
  animation-delay: 0.2s;
}
.animation-delay-400 {
  animation-delay: 0.4s;
}
.animation-delay-600 {
  animation-delay: 0.6s;
}

/* Blob animation specific for the hero section */
@keyframes blob {
  0%, 100% {
    transform: scale(1) translate(0,0);
  }
  33% {
    transform: scale(1.2) translate(30px,-20px);
  }
  66% {
    transform: scale(0.8) translate(-20px,30px);
  }
}
.animate-blob {
  animation: blob 8s infinite ease-in-out; /* Added ease-in-out for smoother motion */
}
</style>