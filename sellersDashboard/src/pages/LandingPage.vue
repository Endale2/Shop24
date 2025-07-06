<template>
  <div class="min-h-screen flex flex-col bg-gradient-to-br from-amber-50 to-green-50 text-gray-900 antialiased font-inter">
    <AppNavbar />

    <!-- Hero Section -->
    <section class="relative flex-grow flex items-center justify-center overflow-hidden px-4 py-24 sm:py-32 lg:py-40">
      <!-- Background Blob Animations -->
      <div aria-hidden class="absolute inset-0 z-0">
        <div class="absolute -top-20 -left-20 w-80 h-80 bg-amber-200 rounded-full mix-blend-multiply filter blur-3xl opacity-30 animate-blob"></div>
        <div class="absolute bottom-0 right-0 w-96 h-96 bg-green-200 rounded-full mix-blend-multiply filter blur-2xl opacity-20 animate-blob animation-delay-400"></div>
        <div class="absolute top-1/4 left-1/4 w-64 h-64 bg-purple-100 rounded-full mix-blend-multiply filter blur-3xl opacity-20 animate-blob animation-delay-200"></div>
      </div>

      <div class="relative z-10 max-w-6xl text-center space-y-8 px-4">
        <div class="animate-fade-in-up">
          <h1 class="text-5xl sm:text-6xl lg:text-7xl font-extrabold leading-tight text-gray-800">
            Elevate Your <span class="text-green-700 drop-shadow-md">E-commerce</span> Game
          </h1>
        </div>
        
        <div class="animate-fade-in-up animation-delay-200">
          <p class="text-xl sm:text-2xl text-gray-700 max-w-4xl mx-auto leading-relaxed">
            Focus on your passion—<strong class="font-bold text-green-800">Seller's Dash</strong> handles the heavy lifting so you can effortlessly <strong class="font-bold text-green-800">build, manage, and scale</strong> your online business.
          </p>
        </div>
        
        <div class="flex flex-col sm:flex-row justify-center gap-6 animate-fade-in-up animation-delay-400">
          <router-link
            to="/register"
            class="inline-flex items-center justify-center px-10 py-4 bg-gradient-to-r from-green-600 to-green-700 text-white text-xl font-bold rounded-full shadow-xl hover:shadow-2xl transition-all duration-300 transform hover:-translate-y-1 focus:outline-none focus:ring-2 focus:ring-green-400 focus:ring-offset-2"
          >
            <svg class="w-6 h-6 mr-3" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 10V3L4 14h7v7l9-11h-7z" />
            </svg>
            Start Free Trial
          </router-link>
          
          <router-link
            to="#features"
            class="inline-flex items-center justify-center px-10 py-4 bg-white text-green-700 text-xl font-bold rounded-full border-2 border-green-300 shadow-lg hover:shadow-xl transition-all duration-300 transform hover:scale-105 focus:outline-none focus:ring-2 focus:ring-green-400 focus:ring-offset-2"
          >
            <svg class="w-6 h-6 mr-3" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 19v-6a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2a2 2 0 002-2zm0 0V9a2 2 0 012-2h2a2 2 0 012 2v10m-6 0a2 2 0 002 2h2a2 2 0 002-2m0 0V5a2 2 0 012-2h2a2 2 0 012 2v14a2 2 0 01-2 2h-2a2 2 0 01-2-2z" />
            </svg>
            View Features
          </router-link>
        </div>

        <!-- Trust Indicators -->
        <div class="animate-fade-in-up animation-delay-600">
          <div class="flex flex-col sm:flex-row items-center justify-center space-y-4 sm:space-y-0 sm:space-x-8 text-sm text-gray-600">
            <div class="flex items-center">
              <svg class="w-5 h-5 text-green-500 mr-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7" />
              </svg>
              <span>No setup fees</span>
            </div>
            <div class="flex items-center">
              <svg class="w-5 h-5 text-green-500 mr-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7" />
              </svg>
              <span>14-day free trial</span>
            </div>
            <div class="flex items-center">
              <svg class="w-5 h-5 text-green-500 mr-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7" />
              </svg>
              <span>Cancel anytime</span>
            </div>
          </div>
        </div>
      </div>
    </section>

    <!-- Features Section -->
    <section id="features" class="py-20 sm:py-28 lg:py-36 px-4 bg-white rounded-t-3xl shadow-inner-top">
      <div class="max-w-7xl mx-auto">
        <div class="text-center mb-16 animate-fade-in">
          <h2 class="text-4xl sm:text-5xl font-extrabold text-gray-900 mb-6">
            Everything You Need to <span class="text-green-700">Thrive</span> Online
          </h2>
          <p class="mt-4 text-xl text-gray-600 max-w-3xl mx-auto leading-relaxed">
            From dazzling storefronts to seamless analytics, we've got your entire sales journey covered with powerful tools designed for modern sellers.
          </p>
        </div>
        
        <div class="grid gap-10 md:grid-cols-2 lg:grid-cols-3">
          <FeatureCard
            icon="M4 5a1 1 0 011-1h14a1 1 0 011 1v2a1 1 0 01-1 1H5a1 1 0 01-1-1V5zM4 13a1 1 0 011-1h6a1 1 0 011 1v6a1 1 0 01-1 1H5a1 1 0 01-1-1v-6zM16 13a1 1 0 011-1h2a1 1 0 011 1v6a1 1 0 01-1 1h-2a1 1 0 01-1-1v-6z"
            title="Stunning Storefronts"
            text="Design a captivating online presence with fully customizable themes that truly represent your brand's unique identity and drive conversions."
            color="green"
            :delay="0"
          />
          <FeatureCard
            icon="M3 10h18M7 15h1m4 0h1m-7 4h12a3 3 0 003-3V8a3 3 0 00-3-3H6a3 3 0 00-3 3v8a3 3 0 003 3z"
            title="Seamless Payments"
            text="Effortlessly accept secure payments worldwide, supporting all major credit cards, digital wallets, and popular payment methods."
            color="amber"
            :delay="200"
          />
          <FeatureCard
            icon="M9 19v-6a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2a2 2 0 002-2zm0 0V9a2 2 0 012-2h2a2 2 0 012 2v10m-6 0a2 2 0 002 2h2a2 2 0 002-2m0 0V5a2 2 0 012-2h2a2 2 0 012 2v14a2 2 0 01-2 2h-2a2 2 0 01-2-2z"
            title="Actionable Analytics"
            text="Unlock powerful insights into your sales, customer behavior, and traffic patterns to make data-driven decisions that fuel growth."
            color="green"
            :delay="400"
          />
          <FeatureCard
            icon="M12 8c-1.657 0-3 .895-3 2s1.343 2 3 2 3 .895 3 2-1.343 2-3 2m0-8c1.11 0 2.08.402 2.599 1M12 8V7m0 1v8m0 0v1m0-1c-1.11 0-2.08-.402-2.599-1"
            title="Smart Inventory"
            text="Keep track of your products with intelligent inventory management, low stock alerts, and automated reorder suggestions."
            color="blue"
            :delay="600"
          />
          <FeatureCard
            icon="M17 20h5v-2a3 3 0 00-5.356-1.857M17 20H7m10 0v-2c0-.656-.126-1.283-.356-1.857M7 20H2v-2a3 3 0 015.356-1.857M7 20v-2c0-.656.126-1.283.356-1.857m0 0a5.002 5.002 0 019.288 0M15 7a3 3 0 11-6 0 3 3 0 016 0zm6 3a2 2 0 11-4 0 2 2 0 014 0zM7 10a2 2 0 11-4 0 2 2 0 014 0z"
            title="Customer Management"
            text="Build lasting relationships with your customers through comprehensive profiles, purchase history, and personalized marketing tools."
            color="purple"
            :delay="800"
          />
          <FeatureCard
            icon="M13 10V3L4 14h7v7l9-11h-7z"
            title="Lightning Fast"
            text="Experience blazing-fast performance with our optimized platform that ensures your store loads quickly and keeps customers engaged."
            color="red"
            :delay="1000"
          />
        </div>

        <!-- CTA Section -->
        <div class="mt-20 text-center">
          <div class="bg-gradient-to-r from-green-50 to-amber-50 rounded-3xl p-12 border border-green-200 shadow-lg">
            <h3 class="text-3xl font-bold text-gray-800 mb-4">
              Ready to Transform Your Business?
            </h3>
            <p class="text-lg text-gray-600 mb-8 max-w-2xl mx-auto">
              Join thousands of successful sellers who've already taken their business to the next level with Seller's Dash.
            </p>
            <router-link
              to="/register"
              class="inline-flex items-center px-10 py-4 bg-gradient-to-r from-green-600 to-green-700 text-white text-xl font-bold rounded-full shadow-lg hover:shadow-xl transition-all duration-300 transform hover:-translate-y-1 focus:outline-none focus:ring-2 focus:ring-green-400 focus:ring-offset-2"
            >
              <svg class="w-6 h-6 mr-3" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 10V3L4 14h7v7l9-11h-7z" />
              </svg>
              Get Started Today
            </router-link>
          </div>
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
/* Added 'Inter' font for better typography */
body {
  font-family: 'Inter', sans-serif;
}

/* Keyframe animations for subtle visual effects */
@keyframes fadeIn {
  from { 
    opacity: 0; 
    transform: translateY(30px); 
  }
  to { 
    opacity: 1; 
    transform: translateY(0); 
  }
}

@keyframes fadeInUp {
  from { 
    opacity: 0; 
    transform: translateY(50px); 
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

.animation-delay-200 { animation-delay: 0.2s; }
.animation-delay-400 { animation-delay: 0.4s; }
.animation-delay-600 { animation-delay: 0.6s; }
.animation-delay-800 { animation-delay: 0.8s; }
.animation-delay-1000 { animation-delay: 1s; }

/* Blob animation for dynamic background elements */
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
  animation: blob 8s infinite ease-in-out;
}

/* Custom utility for inner shadow effect */
.shadow-inner-top {
  box-shadow: inset 0 8px 10px -6px rgba(0, 0, 0, 0.1);
}
</style>
