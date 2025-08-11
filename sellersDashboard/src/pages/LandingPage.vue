<template>
  <div class="min-h-screen bg-gradient-to-br from-gray-900 via-emerald-900 to-gray-900 text-white antialiased font-inter overflow-hidden">
    <!-- Optimized Animated Background -->
    <div class="absolute inset-0 overflow-hidden will-change-transform">
      <!-- Vue Green & Amber gradient mesh -->
      <div class="absolute inset-0 bg-gradient-to-br from-emerald-900/30 via-green-800/20 to-amber-900/25"></div>

      <!-- Optimized animated orbs -->
      <div class="absolute -top-32 -right-32 w-64 h-64 md:w-80 md:h-80 bg-gradient-to-br from-emerald-400/20 to-green-400/15 rounded-full blur-2xl animate-pulse-slow will-change-transform"></div>
      <div class="absolute -bottom-32 -left-32 w-72 h-72 md:w-96 md:h-96 bg-gradient-to-tr from-amber-400/15 to-yellow-400/10 rounded-full blur-2xl animate-pulse-slow animation-delay-1000 will-change-transform"></div>
      <div class="absolute top-1/4 left-1/2 w-48 h-48 md:w-64 md:h-64 bg-gradient-to-br from-green-400/20 to-emerald-400/15 rounded-full blur-2xl animate-pulse-slow animation-delay-2000 will-change-transform"></div>

      <!-- Optimized floating particles -->
      <div class="absolute top-20 left-20 w-1.5 h-1.5 md:w-2 md:h-2 bg-emerald-400 rounded-full animate-float-particle will-change-transform"></div>
      <div class="absolute top-40 right-32 w-1 h-1 bg-amber-400 rounded-full animate-float-particle animation-delay-500 will-change-transform"></div>
      <div class="absolute bottom-32 left-1/4 w-1 h-1 md:w-1.5 md:h-1.5 bg-green-400 rounded-full animate-float-particle animation-delay-1000 will-change-transform"></div>
      <div class="absolute top-1/3 right-1/4 w-1 h-1 bg-yellow-400 rounded-full animate-float-particle animation-delay-1500 will-change-transform"></div>
    </div>

    <!-- Responsive Navigation -->
    <nav class="relative z-50 px-4 sm:px-6 py-4 sm:py-6">
      <div class="max-w-7xl mx-auto flex items-center justify-between">
        <div class="flex items-center space-x-2 sm:space-x-3">
          <div class="w-8 h-8 sm:w-10 sm:h-10 bg-gradient-to-br from-emerald-500 to-green-600 rounded-xl flex items-center justify-center shadow-lg transform hover:scale-105 transition-transform duration-200">
            <svg class="w-4 h-4 sm:w-6 sm:h-6 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 7h8m0 0v8m0-8l-8 8-4-4-6 6" />
            </svg>
          </div>
          <span class="text-xl sm:text-2xl font-black bg-gradient-to-r from-emerald-400 via-green-400 to-amber-400 bg-clip-text text-transparent">Sho24</span>
        </div>

        <div class="hidden md:flex items-center space-x-6 lg:space-x-8">
          <a href="#features" class="text-gray-300 hover:text-emerald-400 transition-colors duration-200 font-medium text-sm lg:text-base">Features</a>
          <a href="#testimonials" class="text-gray-300 hover:text-emerald-400 transition-colors duration-200 font-medium text-sm lg:text-base">Reviews</a>
          <a href="#pricing" class="text-gray-300 hover:text-emerald-400 transition-colors duration-200 font-medium text-sm lg:text-base">Pricing</a>
        </div>

        <!-- Mobile menu button -->
        <button class="md:hidden p-2 text-gray-300 hover:text-emerald-400 transition-colors duration-200">
          <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 6h16M4 12h16M4 18h16" />
          </svg>
        </button>
      </div>
    </nav>

    <!-- Responsive Hero Section -->
    <section class="relative z-10 flex-grow flex items-center justify-center px-4 sm:px-6 py-12 sm:py-20">
      <div class="max-w-6xl mx-auto text-center space-y-8 sm:space-y-12">
        <!-- Hero Title -->
        <div class="animate-fade-in-up">
          <div class="inline-flex items-center px-3 sm:px-4 py-2 bg-emerald-500/20 backdrop-blur-sm border border-emerald-400/30 rounded-full text-emerald-300 text-xs sm:text-sm font-semibold mb-6 sm:mb-8 shadow-lg">
            <svg class="w-3 h-3 sm:w-4 sm:h-4 mr-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 7h8m0 0v8m0-8l-8 8-4-4-6 6" />
            </svg>
            Trusted by 10,000+ sellers worldwide
          </div>

          <h1 class="text-4xl sm:text-6xl md:text-7xl lg:text-8xl font-black leading-tight mb-6 sm:mb-8">
            <span class="block text-white drop-shadow-2xl">Launch Your</span>
            <span class="block text-transparent bg-clip-text bg-gradient-to-r from-emerald-400 via-green-400 to-amber-400 animate-gradient-x">
              Dream Shop
            </span>
          </h1>

          <p class="text-lg sm:text-xl lg:text-2xl text-gray-300 max-w-4xl mx-auto leading-relaxed font-light px-4">
            The fastest e-commerce platform built for modern entrepreneurs.
            <span class="text-transparent bg-clip-text bg-gradient-to-r from-emerald-400 to-amber-400 font-semibold">Start selling in seconds, scale to success.</span>
          </p>
        </div>

        <!-- Responsive Google Login CTA -->
        <div class="animate-fade-in-up animation-delay-400">
          <div class="max-w-sm sm:max-w-lg mx-auto bg-white/10 backdrop-blur-lg rounded-2xl sm:rounded-3xl p-6 sm:p-8 border border-emerald-400/20 shadow-2xl">
            <h3 class="text-2xl sm:text-3xl font-bold mb-3 sm:mb-4 text-white">Start Your Journey</h3>
            <p class="text-gray-300 mb-6 sm:mb-8 text-base sm:text-lg">Join thousands of successful entrepreneurs building their empire</p>

            <button
              @click="loginWithGoogle"
              :disabled="loading"
              class="w-full flex justify-center items-center space-x-3 sm:space-x-4 bg-gradient-to-r from-emerald-500 to-green-600 hover:from-emerald-600 hover:to-green-700 text-white rounded-xl sm:rounded-2xl py-4 sm:py-5 px-6 sm:px-8 transition-all duration-300 shadow-xl hover:shadow-2xl disabled:opacity-60 disabled:cursor-not-allowed group transform hover:scale-105 font-semibold text-base sm:text-xl"
            >
              <svg v-if="loading" class="animate-spin h-6 w-6 sm:h-7 sm:w-7" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24">
                <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"/>
                <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"/>
              </svg>
              <img v-else src="https://www.svgrepo.com/show/355037/google.svg" alt="Google icon" class="h-6 w-6 sm:h-7 sm:w-7 group-hover:scale-110 transition-transform duration-200" />
              <span>
                {{ loading ? 'Setting up your shop...' : 'Start with Google' }}
              </span>
            </button>

            <!-- Responsive Trust badges -->
            <div class="flex flex-wrap items-center justify-center gap-4 sm:gap-8 mt-4 sm:mt-6 text-xs sm:text-sm text-gray-400">
              <div class="flex items-center">
                <svg class="w-3 h-3 sm:w-4 sm:h-4 text-emerald-400 mr-1 sm:mr-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m5.618-4.016A11.955 11.955 0 0112 2.944a11.955 11.955 0 01-8.618 3.04A12.02 12.02 0 003 9c0 5.591 3.824 10.29 9 11.622 5.176-1.332 9-6.03 9-11.622 0-1.042-.133-2.052-.382-3.016z" />
                </svg>
                <span>Secure</span>
              </div>
              <div class="flex items-center">
                <svg class="w-3 h-3 sm:w-4 sm:h-4 text-amber-400 mr-1 sm:mr-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 10V3L4 14h7v7l9-11h-7z" />
                </svg>
                <span>Lightning Fast</span>
              </div>
              <div class="flex items-center">
                <svg class="w-3 h-3 sm:w-4 sm:h-4 text-green-400 mr-1 sm:mr-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7" />
                </svg>
                <span>Free Forever</span>
              </div>
            </div>
          </div>
        </div>

        <!-- Responsive Error/Loading States -->
        <div v-if="error" class="animate-fade-in-up max-w-sm sm:max-w-md mx-auto px-4">
          <div class="bg-red-500/20 backdrop-blur-sm border border-red-400/30 rounded-xl sm:rounded-2xl p-3 sm:p-4 shadow-lg">
            <div class="flex items-center justify-center">
              <svg class="w-4 h-4 sm:w-5 sm:h-5 text-red-400 mr-2 sm:mr-3 flex-shrink-0" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-2.5L13.732 4c-.77-.833-1.964-.833-2.732 0L3.732 16.5c-.77.833.192 2.5 1.732 2.5z" />
              </svg>
              <p class="text-red-300 font-medium text-sm sm:text-base">{{ error }}</p>
            </div>
          </div>
        </div>
      </div>
    </section>

    <!-- Responsive Features Section -->
    <section id="features" class="relative py-16 sm:py-24 px-4 sm:px-6 bg-gradient-to-b from-gray-900 to-gray-800">
      <div class="max-w-7xl mx-auto">
        <div class="text-center mb-12 sm:mb-20 animate-fade-in">
          <h2 class="text-3xl sm:text-5xl lg:text-6xl font-black text-white mb-6 sm:mb-8">
            Everything You Need to <span class="text-transparent bg-clip-text bg-gradient-to-r from-emerald-400 to-amber-400">Succeed</span>
          </h2>
          <p class="text-lg sm:text-xl text-gray-300 max-w-4xl mx-auto leading-relaxed px-4">
            From lightning-fast storefronts to powerful analytics, we've built the ultimate toolkit for modern entrepreneurs ready to scale their business.
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

      </div>
    </section>

    <!-- Responsive Testimonials Section -->
    <section id="testimonials" class="py-16 sm:py-24 bg-gradient-to-b from-gray-800 to-gray-900">
      <div class="max-w-7xl mx-auto px-4 sm:px-6">
        <div class="text-center mb-12 sm:mb-20">
          <h3 class="text-3xl sm:text-5xl font-black text-white mb-4 sm:mb-6">
            Loved by <span class="text-transparent bg-clip-text bg-gradient-to-r from-emerald-400 to-amber-400">Entrepreneurs</span>
          </h3>
          <p class="text-lg sm:text-xl text-gray-300 max-w-2xl mx-auto px-4">Real stories from real sellers who transformed their business with Sho24</p>
        </div>

        <div class="grid sm:grid-cols-2 lg:grid-cols-3 gap-6 sm:gap-8 mb-12 sm:mb-20">
          <!-- Testimonial 1 -->
          <div class="bg-gray-700/50 backdrop-blur-sm rounded-xl sm:rounded-2xl p-6 sm:p-8 border border-emerald-600/20 shadow-2xl transform hover:scale-105 transition-all duration-300 group">
            <div class="flex items-center mb-4 sm:mb-6">
              <div class="w-12 h-12 sm:w-14 sm:h-14 bg-gradient-to-br from-emerald-500 to-green-500 rounded-full flex items-center justify-center text-white font-bold text-lg sm:text-xl shadow-lg">
                S
              </div>
              <div class="ml-3 sm:ml-4">
                <h4 class="font-bold text-white text-base sm:text-lg">Sarah Chen</h4>
                <p class="text-emerald-300 text-xs sm:text-sm">Fashion Boutique Owner</p>
              </div>
            </div>
            <p class="text-gray-300 text-base sm:text-lg leading-relaxed mb-4 sm:mb-6">"Sho24 transformed my business. Sales increased 300% in just 3 months!"</p>
            <div class="flex space-x-1">
              <svg class="w-5 h-5 text-yellow-400" fill="currentColor" viewBox="0 0 20 20">
                <path d="M9.049 2.927c.3-.921 1.603-.921 1.902 0l1.07 3.292a1 1 0 00.95.69h3.462c.969 0 1.371 1.24.588 1.81l-2.8 2.034a1 1 0 00-.364 1.118l1.07 3.292c.3.921-.755 1.688-1.54 1.118l-2.8-2.034a1 1 0 00-1.175 0l-2.8 2.034c-.784.57-1.838-.197-1.539-1.118l1.07-3.292a1 1 0 00-.364-1.118L2.98 8.72c-.783-.57-.38-1.81.588-1.81h3.461a1 1 0 00.951-.69l1.07-3.292z"/>
              </svg>
              <svg class="w-5 h-5 text-yellow-400" fill="currentColor" viewBox="0 0 20 20">
                <path d="M9.049 2.927c.3-.921 1.603-.921 1.902 0l1.07 3.292a1 1 0 00.95.69h3.462c.969 0 1.371 1.24.588 1.81l-2.8 2.034a1 1 0 00-.364 1.118l1.07 3.292c.3.921-.755 1.688-1.54 1.118l-2.8-2.034a1 1 0 00-1.175 0l-2.8 2.034c-.784.57-1.838-.197-1.539-1.118l1.07-3.292a1 1 0 00-.364-1.118L2.98 8.72c-.783-.57-.38-1.81.588-1.81h3.461a1 1 0 00.951-.69l1.07-3.292z"/>
              </svg>
              <svg class="w-5 h-5 text-yellow-400" fill="currentColor" viewBox="0 0 20 20">
                <path d="M9.049 2.927c.3-.921 1.603-.921 1.902 0l1.07 3.292a1 1 0 00.95.69h3.462c.969 0 1.371 1.24.588 1.81l-2.8 2.034a1 1 0 00-.364 1.118l1.07 3.292c.3.921-.755 1.688-1.54 1.118l-2.8-2.034a1 1 0 00-1.175 0l-2.8 2.034c-.784.57-1.838-.197-1.539-1.118l1.07-3.292a1 1 0 00-.364-1.118L2.98 8.72c-.783-.57-.38-1.81.588-1.81h3.461a1 1 0 00.951-.69l1.07-3.292z"/>
              </svg>
              <svg class="w-5 h-5 text-yellow-400" fill="currentColor" viewBox="0 0 20 20">
                <path d="M9.049 2.927c.3-.921 1.603-.921 1.902 0l1.07 3.292a1 1 0 00.95.69h3.462c.969 0 1.371 1.24.588 1.81l-2.8 2.034a1 1 0 00-.364 1.118l1.07 3.292c.3.921-.755 1.688-1.54 1.118l-2.8-2.034a1 1 0 00-1.175 0l-2.8 2.034c-.784.57-1.838-.197-1.539-1.118l1.07-3.292a1 1 0 00-.364-1.118L2.98 8.72c-.783-.57-.38-1.81.588-1.81h3.461a1 1 0 00.951-.69l1.07-3.292z"/>
              </svg>
              <svg class="w-5 h-5 text-yellow-400" fill="currentColor" viewBox="0 0 20 20">
                <path d="M9.049 2.927c.3-.921 1.603-.921 1.902 0l1.07 3.292a1 1 0 00.95.69h3.462c.969 0 1.371 1.24.588 1.81l-2.8 2.034a1 1 0 00-.364 1.118l1.07 3.292c.3.921-.755 1.688-1.54 1.118l-2.8-2.034a1 1 0 00-1.175 0l-2.8 2.034c-.784.57-1.838-.197-1.539-1.118l1.07-3.292a1 1 0 00-.364-1.118L2.98 8.72c-.783-.57-.38-1.81.588-1.81h3.461a1 1 0 00.951-.69l1.07-3.292z"/>
              </svg>
            </div>
          </div>

          <!-- Testimonial 2 -->
          <div class="bg-gray-700/50 backdrop-blur-sm rounded-xl sm:rounded-2xl p-6 sm:p-8 border border-amber-600/20 shadow-2xl transform hover:scale-105 transition-all duration-300 group">
            <div class="flex items-center mb-4 sm:mb-6">
              <div class="w-12 h-12 sm:w-14 sm:h-14 bg-gradient-to-br from-amber-500 to-yellow-500 rounded-full flex items-center justify-center text-white font-bold text-lg sm:text-xl shadow-lg">
                M
              </div>
              <div class="ml-3 sm:ml-4">
                <h4 class="font-bold text-white text-base sm:text-lg">Marcus Johnson</h4>
                <p class="text-amber-300 text-xs sm:text-sm">Electronics Store</p>
              </div>
            </div>
            <p class="text-gray-300 text-base sm:text-lg leading-relaxed mb-4 sm:mb-6">"The speed is incredible. My store loads instantly and customers love the experience."</p>
            <div class="flex space-x-1">
              <svg class="w-5 h-5 text-yellow-400" fill="currentColor" viewBox="0 0 20 20">
                <path d="M9.049 2.927c.3-.921 1.603-.921 1.902 0l1.07 3.292a1 1 0 00.95.69h3.462c.969 0 1.371 1.24.588 1.81l-2.8 2.034a1 1 0 00-.364 1.118l1.07 3.292c.3.921-.755 1.688-1.54 1.118l-2.8-2.034a1 1 0 00-1.175 0l-2.8 2.034c-.784.57-1.838-.197-1.539-1.118l1.07-3.292a1 1 0 00-.364-1.118L2.98 8.72c-.783-.57-.38-1.81.588-1.81h3.461a1 1 0 00.951-.69l1.07-3.292z"/>
              </svg>
              <svg class="w-5 h-5 text-yellow-400" fill="currentColor" viewBox="0 0 20 20">
                <path d="M9.049 2.927c.3-.921 1.603-.921 1.902 0l1.07 3.292a1 1 0 00.95.69h3.462c.969 0 1.371 1.24.588 1.81l-2.8 2.034a1 1 0 00-.364 1.118l1.07 3.292c.3.921-.755 1.688-1.54 1.118l-2.8-2.034a1 1 0 00-1.175 0l-2.8 2.034c-.784.57-1.838-.197-1.539-1.118l1.07-3.292a1 1 0 00-.364-1.118L2.98 8.72c-.783-.57-.38-1.81.588-1.81h3.461a1 1 0 00.951-.69l1.07-3.292z"/>
              </svg>
              <svg class="w-5 h-5 text-yellow-400" fill="currentColor" viewBox="0 0 20 20">
                <path d="M9.049 2.927c.3-.921 1.603-.921 1.902 0l1.07 3.292a1 1 0 00.95.69h3.462c.969 0 1.371 1.24.588 1.81l-2.8 2.034a1 1 0 00-.364 1.118l1.07 3.292c.3.921-.755 1.688-1.54 1.118l-2.8-2.034a1 1 0 00-1.175 0l-2.8 2.034c-.784.57-1.838-.197-1.539-1.118l1.07-3.292a1 1 0 00-.364-1.118L2.98 8.72c-.783-.57-.38-1.81.588-1.81h3.461a1 1 0 00.951-.69l1.07-3.292z"/>
              </svg>
              <svg class="w-5 h-5 text-yellow-400" fill="currentColor" viewBox="0 0 20 20">
                <path d="M9.049 2.927c.3-.921 1.603-.921 1.902 0l1.07 3.292a1 1 0 00.95.69h3.462c.969 0 1.371 1.24.588 1.81l-2.8 2.034a1 1 0 00-.364 1.118l1.07 3.292c.3.921-.755 1.688-1.54 1.118l-2.8-2.034a1 1 0 00-1.175 0l-2.8 2.034c-.784.57-1.838-.197-1.539-1.118l1.07-3.292a1 1 0 00-.364-1.118L2.98 8.72c-.783-.57-.38-1.81.588-1.81h3.461a1 1 0 00.951-.69l1.07-3.292z"/>
              </svg>
              <svg class="w-5 h-5 text-yellow-400" fill="currentColor" viewBox="0 0 20 20">
                <path d="M9.049 2.927c.3-.921 1.603-.921 1.902 0l1.07 3.292a1 1 0 00.95.69h3.462c.969 0 1.371 1.24.588 1.81l-2.8 2.034a1 1 0 00-.364 1.118l1.07 3.292c.3.921-.755 1.688-1.54 1.118l-2.8-2.034a1 1 0 00-1.175 0l-2.8 2.034c-.784.57-1.838-.197-1.539-1.118l1.07-3.292a1 1 0 00-.364-1.118L2.98 8.72c-.783-.57-.38-1.81.588-1.81h3.461a1 1 0 00.951-.69l1.07-3.292z"/>
              </svg>
            </div>
          </div>

          <!-- Testimonial 3 -->
          <div class="bg-gray-700/50 backdrop-blur-sm rounded-xl sm:rounded-2xl p-6 sm:p-8 border border-green-600/20 shadow-2xl transform hover:scale-105 transition-all duration-300 group sm:col-span-2 lg:col-span-1">
            <div class="flex items-center mb-4 sm:mb-6">
              <div class="w-12 h-12 sm:w-14 sm:h-14 bg-gradient-to-br from-green-500 to-emerald-600 rounded-full flex items-center justify-center text-white font-bold text-lg sm:text-xl shadow-lg">
                A
              </div>
              <div class="ml-3 sm:ml-4">
                <h4 class="font-bold text-white text-base sm:text-lg">Alex Rivera</h4>
                <p class="text-green-300 text-xs sm:text-sm">Handmade Crafts</p>
              </div>
            </div>
            <p class="text-gray-300 text-base sm:text-lg leading-relaxed mb-4 sm:mb-6">"The analytics helped me understand my customers better. Revenue doubled with Sho24!"</p>
            <div class="flex space-x-1">
              <svg class="w-5 h-5 text-yellow-400" fill="currentColor" viewBox="0 0 20 20">
                <path d="M9.049 2.927c.3-.921 1.603-.921 1.902 0l1.07 3.292a1 1 0 00.95.69h3.462c.969 0 1.371 1.24.588 1.81l-2.8 2.034a1 1 0 00-.364 1.118l1.07 3.292c.3.921-.755 1.688-1.54 1.118l-2.8-2.034a1 1 0 00-1.175 0l-2.8 2.034c-.784.57-1.838-.197-1.539-1.118l1.07-3.292a1 1 0 00-.364-1.118L2.98 8.72c-.783-.57-.38-1.81.588-1.81h3.461a1 1 0 00.951-.69l1.07-3.292z"/>
              </svg>
              <svg class="w-5 h-5 text-yellow-400" fill="currentColor" viewBox="0 0 20 20">
                <path d="M9.049 2.927c.3-.921 1.603-.921 1.902 0l1.07 3.292a1 1 0 00.95.69h3.462c.969 0 1.371 1.24.588 1.81l-2.8 2.034a1 1 0 00-.364 1.118l1.07 3.292c.3.921-.755 1.688-1.54 1.118l-2.8-2.034a1 1 0 00-1.175 0l-2.8 2.034c-.784.57-1.838-.197-1.539-1.118l1.07-3.292a1 1 0 00-.364-1.118L2.98 8.72c-.783-.57-.38-1.81.588-1.81h3.461a1 1 0 00.951-.69l1.07-3.292z"/>
              </svg>
              <svg class="w-5 h-5 text-yellow-400" fill="currentColor" viewBox="0 0 20 20">
                <path d="M9.049 2.927c.3-.921 1.603-.921 1.902 0l1.07 3.292a1 1 0 00.95.69h3.462c.969 0 1.371 1.24.588 1.81l-2.8 2.034a1 1 0 00-.364 1.118l1.07 3.292c.3.921-.755 1.688-1.54 1.118l-2.8-2.034a1 1 0 00-1.175 0l-2.8 2.034c-.784.57-1.838-.197-1.539-1.118l1.07-3.292a1 1 0 00-.364-1.118L2.98 8.72c-.783-.57-.38-1.81.588-1.81h3.461a1 1 0 00.951-.69l1.07-3.292z"/>
              </svg>
              <svg class="w-5 h-5 text-yellow-400" fill="currentColor" viewBox="0 0 20 20">
                <path d="M9.049 2.927c.3-.921 1.603-.921 1.902 0l1.07 3.292a1 1 0 00.95.69h3.462c.969 0 1.371 1.24.588 1.81l-2.8 2.034a1 1 0 00-.364 1.118l1.07 3.292c.3.921-.755 1.688-1.54 1.118l-2.8-2.034a1 1 0 00-1.175 0l-2.8 2.034c-.784.57-1.838-.197-1.539-1.118l1.07-3.292a1 1 0 00-.364-1.118L2.98 8.72c-.783-.57-.38-1.81.588-1.81h3.461a1 1 0 00.951-.69l1.07-3.292z"/>
              </svg>
            </div>
          </div>
        </div>

        <!-- Responsive Stats Section -->
        <div class="grid grid-cols-2 lg:grid-cols-4 gap-6 sm:gap-8 text-center">
          <div class="animate-fade-in-up">
            <div class="text-3xl sm:text-4xl font-black text-transparent bg-clip-text bg-gradient-to-r from-emerald-400 to-green-400 mb-2">10,000+</div>
            <div class="text-gray-300 font-medium text-sm sm:text-base">Active Sellers</div>
          </div>
          <div class="animate-fade-in-up animation-delay-200">
            <div class="text-3xl sm:text-4xl font-black text-transparent bg-clip-text bg-gradient-to-r from-amber-400 to-yellow-400 mb-2">$2.5M+</div>
            <div class="text-gray-300 font-medium text-sm sm:text-base">Revenue Generated</div>
          </div>
          <div class="animate-fade-in-up animation-delay-400">
            <div class="text-3xl sm:text-4xl font-black text-transparent bg-clip-text bg-gradient-to-r from-green-400 to-emerald-500 mb-2">99.9%</div>
            <div class="text-gray-300 font-medium text-sm sm:text-base">Uptime</div>
          </div>
          <div class="animate-fade-in-up animation-delay-600">
            <div class="text-3xl sm:text-4xl font-black text-transparent bg-clip-text bg-gradient-to-r from-yellow-400 to-amber-500 mb-2">24/7</div>
            <div class="text-gray-300 font-medium text-sm sm:text-base">Support</div>
          </div>
        </div>
      </div>
    </section>

    <!-- Responsive Final CTA Section -->
    <section class="py-16 sm:py-24 bg-gradient-to-t from-gray-900 to-gray-800 border-t border-emerald-700/30">
      <div class="max-w-4xl mx-auto text-center px-4 sm:px-6">
        <h3 class="text-3xl sm:text-5xl font-black text-white mb-6 sm:mb-8">
          Ready to <span class="text-transparent bg-clip-text bg-gradient-to-r from-emerald-400 to-amber-400">Launch</span> Your Shop?
        </h3>
        <p class="text-lg sm:text-xl text-gray-300 mb-8 sm:mb-12 leading-relaxed">
          Join the revolution of successful entrepreneurs who chose to build their empire with Shop24.
        </p>

        <button
          @click="loginWithGoogle"
          :disabled="loading"
          class="inline-flex items-center space-x-3 sm:space-x-4 bg-gradient-to-r from-emerald-600 to-green-600 hover:from-emerald-700 hover:to-green-700 text-white rounded-xl sm:rounded-2xl py-4 sm:py-6 px-8 sm:px-12 transition-all duration-300 shadow-2xl hover:shadow-emerald-500/25 disabled:opacity-60 disabled:cursor-not-allowed group transform hover:scale-105 text-lg sm:text-xl font-bold"
        >
          <svg v-if="loading" class="animate-spin h-6 w-6 sm:h-8 sm:w-8" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24">
            <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"/>
            <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"/>
          </svg>
          <img v-else src="https://www.svgrepo.com/show/355037/google.svg" alt="Google icon" class="h-6 w-6 sm:h-8 sm:w-8 group-hover:scale-110 transition-transform duration-200" />
          <span>{{ loading ? 'Building Your Shop...' : 'Launch Your Shop Now' }}</span>
        </button>
      </div>
    </section>

    <!-- Responsive Modern Footer -->
    <footer class="bg-gray-900 border-t border-emerald-800/30 py-8 sm:py-12">
      <div class="max-w-7xl mx-auto px-4 sm:px-6">
        <div class="flex flex-col sm:flex-row items-center justify-between space-y-4 sm:space-y-0">
          <div class="flex items-center space-x-2 sm:space-x-3">
            <div class="w-8 h-8 sm:w-10 sm:h-10 bg-gradient-to-br from-emerald-500 to-green-600 rounded-xl flex items-center justify-center shadow-lg">
              <svg class="w-4 h-4 sm:w-6 sm:h-6 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 7h8m0 0v8m0-8l-8 8-4-4-6 6" />
              </svg>
            </div>
            <span class="text-xl sm:text-2xl font-black bg-gradient-to-r from-emerald-400 via-green-400 to-amber-400 bg-clip-text text-transparent">Sho24</span>
          </div>

          <div class="flex flex-col sm:flex-row items-center space-y-2 sm:space-y-0 sm:space-x-6 lg:space-x-8 text-gray-400">
            <span class="text-xs sm:text-sm text-center">© 2024 Shop24. Built for ambitious entrepreneurs.</span>
            <div class="flex space-x-4 sm:space-x-6">
              <a href="#" class="hover:text-emerald-400 transition-colors duration-200 text-xs sm:text-sm">Privacy</a>
              <a href="#" class="hover:text-emerald-400 transition-colors duration-200 text-xs sm:text-sm">Terms</a>
              <a href="#" class="hover:text-emerald-400 transition-colors duration-200 text-xs sm:text-sm">Support</a>
            </div>
          </div>
        </div>
      </div>
    </footer>
  </div>
</template>

<script setup>
import { ref, watch, onMounted } from 'vue'
import { useRouter, useRoute } from 'vue-router'
import { useAuthStore } from '@/store/auth'
import FeatureCard from '@/components/FeatureCard.vue'

const router = useRouter()
const route = useRoute()
const auth = useAuthStore()
const loading = ref(false)
const error = ref(null)

// Helper function to determine where to redirect after authentication
function determineRedirect() {
  const user = auth.user

  // If user doesn't have first_name or last_name, they need to complete profile
  if (!user || !user.first_name || !user.last_name) {
    return { name: 'CompleteProfile' }
  }

  // If user has complete profile, redirect to shop selection
  return { name: 'ShopSelection' }
}

// Google OAuth login
function loginWithGoogle() {
  try {
    loading.value = true
    error.value = null
    auth.loginWithGoogle()
  } catch (err) {
    error.value = 'Failed to initiate Google login. Please try again.'
    loading.value = false
  }
}

// Handle OAuth callback from Google
async function handleOAuthCallback() {
  loading.value = true
  error.value = null

  try {
    await auth.fetchMe()
    const redirect = determineRedirect()
    router.replace(redirect)
  } catch (err) {
    error.value = 'Authentication failed. Please try again.'
    console.error('OAuth callback error:', err)
  } finally {
    loading.value = false
  }
}

onMounted(() => {
  // If user is already authenticated, redirect appropriately
  if (auth.isAuthenticated) {
    const redirect = determineRedirect()
    router.replace(redirect)
  }
})

// Watch for OAuth callback parameters
watch(() => route.query, async (query) => {
  if (query.oauth === 'google' || query.code) {
    await handleOAuthCallback()
  }

  if (query.error) {
    error.value = `Authentication failed: ${query.error}`
    loading.value = false
  }
}, { immediate: true, deep: true })
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

/* Enhanced glassmorphism effect */
.backdrop-blur-sm {
  backdrop-filter: blur(8px);
}

/* Pulse animation for decorative elements */
@keyframes pulse {
  0%, 100% {
    opacity: 0.15;
    transform: scale(1);
  }
  50% {
    opacity: 0.25;
    transform: scale(1.05);
  }
}

.animate-pulse {
  animation: pulse 3s ease-in-out infinite;
}

/* Enhanced hover effects for buttons */
.group:hover .group-hover\:scale-110 {
  transform: scale(1.1);
}

/* Smooth gradient transitions */
.bg-gradient-to-r {
  background-size: 200% 200%;
  animation: gradient-shift 6s ease infinite;
}

@keyframes gradient-shift {
  0% { background-position: 0% 50%; }
  50% { background-position: 100% 50%; }
  100% { background-position: 0% 50%; }
}

/* Floating animation for decorative elements */
@keyframes float {
  0%, 100% {
    transform: translateY(0px) rotate(0deg);
  }
  33% {
    transform: translateY(-10px) rotate(120deg);
  }
  66% {
    transform: translateY(5px) rotate(240deg);
  }
}

.animate-float {
  animation: float 6s ease-in-out infinite;
}

.animation-delay-800 { animation-delay: 0.8s; }
.animation-delay-1200 { animation-delay: 1.2s; }

/* Additional responsive enhancements */
@media (max-width: 640px) {
  .text-5xl { font-size: 2.5rem; }
  .text-6xl { font-size: 3rem; }
  .text-7xl { font-size: 3.5rem; }
}

/* Enhanced button hover effects */
button:hover {
  transform: translateY(-2px);
}

/* Smooth scroll behavior */
html {
  scroll-behavior: smooth;
}

/* Additional responsive enhancements */
@media (max-width: 640px) {
  .text-5xl { font-size: 2.5rem; }
  .text-6xl { font-size: 3rem; }
  .text-7xl { font-size: 3.5rem; }
}

/* Enhanced button hover effects */
button:hover {
  transform: translateY(-2px);
}

/* Smooth scroll behavior */
html {
  scroll-behavior: smooth;
}

/* Optimized animations for Vue green/amber theme with performance improvements */
@keyframes pulse-slow {
  0%, 100% {
    opacity: 0.15;
    transform: scale(1) translateZ(0);
  }
  50% {
    opacity: 0.25;
    transform: scale(1.05) translateZ(0);
  }
}

.animate-pulse-slow {
  animation: pulse-slow 6s ease-in-out infinite;
  will-change: transform, opacity;
}

@keyframes float-particle {
  0%, 100% {
    transform: translate3d(0, 0, 0) rotate(0deg);
    opacity: 0.4;
  }
  25% {
    transform: translate3d(10px, -15px, 0) rotate(90deg);
    opacity: 0.8;
  }
  50% {
    transform: translate3d(-8px, -8px, 0) rotate(180deg);
    opacity: 0.6;
  }
  75% {
    transform: translate3d(5px, -20px, 0) rotate(270deg);
    opacity: 0.3;
  }
}

.animate-float-particle {
  animation: float-particle 12s ease-in-out infinite;
  will-change: transform, opacity;
}

@keyframes gradient-x {
  0%, 100% {
    background-size: 200% 200%;
    background-position: left center;
  }
  50% {
    background-size: 200% 200%;
    background-position: right center;
  }
}

.animate-gradient-x {
  animation: gradient-x 4s ease infinite;
  will-change: background-position;
}

/* Performance optimizations */
.will-change-transform {
  will-change: transform;
}

/* Reduced motion for accessibility */
@media (prefers-reduced-motion: reduce) {
  .animate-pulse-slow,
  .animate-float-particle,
  .animate-gradient-x {
    animation: none;
  }
}

/* Additional animation delays */
.animation-delay-500 { animation-delay: 0.5s; }
.animation-delay-1500 { animation-delay: 1.5s; }
.animation-delay-2000 { animation-delay: 2s; }
</style>
