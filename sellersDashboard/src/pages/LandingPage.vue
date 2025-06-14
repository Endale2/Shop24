<template>
  <div class="bg-green-50 text-gray-800 antialiased font-sans">
    <header class="w-full bg-white/90 backdrop-blur-lg sticky top-0 z-50 border-b border-green-100 shadow-sm">
      <div class="max-w-7xl mx-auto flex justify-between items-center p-4 lg:px-8">
        <router-link to="/" class="text-3xl font-extrabold text-green-700 tracking-tight flex items-center group">
          <svg class="h-8 w-8 mr-2 text-amber-500 group-hover:rotate-6 transition-transform duration-300" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 10V3L4 14h7v7l9-11h-7z" />
          </svg>
          Seller's <span class="text-amber-600">Dash</span>
        </router-link>
        <nav class="flex items-center space-x-6">
          <router-link to="/login" class="text-gray-600 font-medium hover:text-green-700 transition-all duration-300 ease-in-out transform hover:scale-105">
            Log In
          </router-link>
          <router-link to="/register" class="px-6 py-2 bg-green-500 text-white font-semibold rounded-full shadow-lg hover:shadow-xl transition-all duration-300 ease-in-out transform hover:-translate-y-0.5">
            Get Started
          </router-link>
        </nav>
      </div>
    </header>

    <main class="relative overflow-hidden">
      <section class="relative py-24 sm:py-32 lg:py-40 px-4 text-center">
        <div class="absolute inset-0 bg-pattern-light opacity-20 z-0"></div>
        <div class="max-w-5xl mx-auto relative z-10">
          <h2 class="text-5xl sm:text-6xl lg:text-7xl font-extrabold text-gray-900 leading-tight mb-8 animate-fade-in-up">
            Elevate Your <span class="text-green-600 drop-shadow-md">E-commerce</span> Game
          </h2>
          <p class="text-xl sm:text-2xl text-gray-700 max-w-3xl mx-auto mb-12 animate-fade-in-up animation-delay-200">
            **Seller's Dash** empowers you with intuitive tools to effortlessly **build, manage, and scale** your online business. Focus on your passion, we'll handle the platform.
          </p>
          <div class="flex justify-center space-x-6 animate-fade-in-up animation-delay-400">
            <router-link to="/register" class="px-10 py-4 bg-green-500 text-white text-xl font-bold rounded-full shadow-lg hover:shadow-xl transition-all duration-300 ease-in-out transform hover:scale-105 hover:-translate-y-1">
              Start Your Free Journey
            </router-link>
            <router-link to="/features" class="px-10 py-4 bg-white text-green-700 text-xl font-bold rounded-full border border-green-200 shadow-md hover:bg-green-50 transition-all duration-300 ease-in-out transform hover:scale-105">
              Discover Features
            </router-link>
          </div>
        </div>
      </section>

      <section id="features" class="py-20 sm:py-28 lg:py-36 px-4 bg-white">
        <div class="max-w-7xl mx-auto">
          <div class="text-center mb-16 animate-fade-in">
            <h3 class="text-4xl sm:text-5xl font-extrabold text-gray-900">
              Everything You Need to <span class="text-green-600">Thrive</span> Online
            </h3>
            <p class="mt-6 text-xl text-gray-600 max-w-2xl mx-auto">
              From dazzling storefronts to seamless analytics, we've got your entire sales journey covered.
            </p>
          </div>
          <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-10 text-center">
            <div class="p-8 bg-white rounded-2xl border border-green-100 shadow-lg hover:shadow-2xl transition-all duration-300 ease-in-out transform hover:-translate-y-2 group animate-fade-in animation-delay-200">
              <div class="flex items-center justify-center h-16 w-16 rounded-full bg-green-100 text-green-600 mx-auto mb-6 group-hover:bg-green-200 transition-colors duration-300">
                <svg xmlns="http://www.w3.org/2000/svg" class="h-9 w-9" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 5a1 1 0 011-1h14a1 1 0 011 1v2a1 1 0 01-1 1H5a1 1 0 01-1-1V5zM4 13a1 1 0 011-1h6a1 1 0 011 1v6a1 1 0 01-1 1H5a1 1 0 01-1-1v-6zM16 13a1 1 0 011-1h2a1 1 0 011 1v6a1 1 0 01-1 1h-2a1 1 0 01-1-1v-6z" />
                </svg>
              </div>
              <h4 class="text-2xl font-bold text-gray-900 mb-3">Stunning Storefronts</h4>
              <p class="text-gray-600 leading-relaxed">
                Design a captivating online presence with fully customizable themes that truly represent your brand's unique identity.
              </p>
            </div>
            <div class="p-8 bg-white rounded-2xl border border-green-100 shadow-lg hover:shadow-2xl transition-all duration-300 ease-in-out transform hover:-translate-y-2 group animate-fade-in animation-delay-400">
              <div class="flex items-center justify-center h-16 w-16 rounded-full bg-amber-100 text-amber-600 mx-auto mb-6 group-hover:bg-amber-200 transition-colors duration-300">
                <svg xmlns="http://www.w3.org/2000/svg" class="h-9 w-9" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 10h18M7 15h1m4 0h1m-7 4h12a3 3 0 003-3V8a3 3 0 00-3-3H6a3 3 0 00-3 3v8a3 3 0 003 3z" />
                </svg>
              </div>
              <h4 class="text-2xl font-bold text-gray-900 mb-3">Seamless Payments</h4>
              <p class="text-gray-600 leading-relaxed">
                Effortlessly accept secure payments from customers worldwide, supporting all major credit cards and popular methods.
              </p>
            </div>
            <div class="p-8 bg-white rounded-2xl border border-green-100 shadow-lg hover:shadow-2xl transition-all duration-300 ease-in-out transform hover:-translate-y-2 group animate-fade-in animation-delay-600">
              <div class="flex items-center justify-center h-16 w-16 rounded-full bg-lime-100 text-lime-600 mx-auto mb-6 group-hover:bg-lime-200 transition-colors duration-300">
                <svg xmlns="http://www.w3.org/2000/svg" class="h-9 w-9" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 19v-6a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2a2 2 0 002-2zm0 0V9a2 2 0 012-2h2a2 2 0 012 2v10m-6 0a2 2 0 002 2h2a2 2 0 002-2m0 0V5a2 2 0 012-2h2a2 2 0 012 2v14a2 2 0 01-2 2h-2a2 2 0 01-2-2z" />
                </svg>
              </div>
              <h4 class="text-2xl font-bold text-gray-900 mb-3">Actionable Analytics</h4>
              <p class="text-gray-600 leading-relaxed">
                Unlock powerful insights into your sales, customer behavior, and traffic to make data-driven decisions that fuel growth.
              </p>
            </div>
          </div>
        </div>
      </section>
    </main>

    <footer class="bg-gray-900 text-gray-400">
      <div class="max-w-7xl mx-auto py-12 px-4 lg:px-8 text-center">
        <p class="text-lg">&copy; {{ new Date().getFullYear() }} Seller's Dash. All rights reserved.</p>
        <div class="mt-6 flex justify-center space-x-8 text-md">
          <a href="#" class="hover:text-white transition-colors duration-300">Privacy Policy</a>
          <a href="#" class="hover:text-white transition-colors duration-300">Terms of Service</a>
          <a href="#" class="hover:text-white transition-colors duration-300">Support</a>
        </div>
      </div>
    </footer>
  </div>
</template>

<script setup>
// No script logic needed for this static landing page.
</script>

<style>
/* Base font family for a modern feel */
body {
  font-family: 'Inter', sans-serif; /* Or any other clean sans-serif font */
}

/* Custom utility for patterned background - adjust for new theme */
.bg-pattern-light {
  background-image: url("data:image/svg+xml,%3Csvg width='60' height='60' viewBox='0 0 60 60' xmlns='http://www.w3.org/2000/svg'%3E%3Cg fill='none' fill-rule='evenodd'%3E%3Cg fill='%239C92AC' fill-opacity='0.05'%3E%3Cpath d='M36 34v-4h-2v4h-4v2h4v4h2v-4h4v-2h-4zm0 0v-4h-2v4h-4v2h4v4h2v-4h4v-2h-4zM24 22v-4h-2v4h-4v2h4v4h2v-4h4v-2h-4zm0 0v-4h-2v4h-4v2h4v4h2v-4h4v-2h-4zM12 10v-4h-2v4h-4v2h4v4h2v-4h4v-2h-4zm0 0v-4h-2v4h-4v2h4v4h2v-4h4v-2h-4z'/%3E%3C/g%3E%3C/g%3E%3C/svg%3E");
}

/* Animations */
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
</style>