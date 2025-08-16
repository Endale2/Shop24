<template>
  <div class="min-h-screen bg-gradient-to-br from-white via-emerald-50/30 to-green-50/20 text-gray-800 font-sans antialiased overflow-x-hidden">
    <!-- Cool animated background -->
    <div class="absolute inset-0 z-0">
      <!-- Animated gradient mesh -->
      <div class="absolute inset-0 bg-[radial-gradient(circle_at_25%_25%,rgba(16,185,129,0.15)_0%,transparent_50%),radial-gradient(circle_at_75%_75%,rgba(34,197,94,0.1)_0%,transparent_50%)]"></div>
      <!-- Moving dot pattern -->
      <div class="absolute inset-0 bg-[radial-gradient(#10b981_1px,transparent_1px)] [background-size:32px_32px] opacity-20 animate-pulse"></div>
      <!-- Floating gradient orbs -->
      <div class="absolute top-20 left-10 w-64 h-64 bg-gradient-to-br from-emerald-200/60 to-green-300/40 rounded-full filter blur-3xl opacity-70 animate-float"></div>
      <div class="absolute top-40 right-20 w-48 h-48 bg-gradient-to-bl from-green-200/50 to-emerald-400/30 rounded-full filter blur-2xl opacity-60 animate-float-delayed"></div>
      <div class="absolute bottom-32 left-1/3 w-56 h-56 bg-gradient-to-tr from-emerald-300/40 to-green-200/50 rounded-full filter blur-3xl opacity-50 animate-float-slow"></div>
    </div>

    <!-- Main Content -->
    <div class="relative z-10 flex flex-col min-h-screen">
      <!-- Seamless Navigation integrated with Hero -->
      <nav class="px-4 sm:px-6 py-8 relative z-50">
        <div class="max-w-6xl mx-auto flex items-center justify-between">
          <div class="flex items-center space-x-3 group">
            <div class="w-12 h-12 bg-gradient-to-br from-emerald-500 via-green-500 to-emerald-600 rounded-xl flex items-center justify-center shadow-lg shadow-emerald-500/30 group-hover:shadow-emerald-500/50 transition-all duration-300 group-hover:scale-110 group-hover:rotate-3">
              <svg class="w-7 h-7 text-white group-hover:scale-110 transition-transform duration-300" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2.5" d="M13 7h8m0 0v8m0-8l-8 8-4-4-6 6" />
              </svg>
            </div>
            <span class="text-2xl font-bold bg-gradient-to-r from-emerald-500 via-green-500 to-emerald-600 bg-clip-text text-transparent hover:from-emerald-600 hover:via-green-600 hover:to-emerald-700 transition-all duration-300 animate-gradient-x">Shop24</span>
          </div>

          <div class="hidden md:flex items-center space-x-8">
            <a href="#features" class="relative text-gray-700 hover:text-emerald-600 transition-all duration-300 font-medium group">
              Features
              <span class="absolute -bottom-1 left-0 w-0 h-0.5 bg-gradient-to-r from-emerald-500 to-green-500 transition-all duration-300 group-hover:w-full rounded-full"></span>
            </a>
            <button
              @click="loginWithGoogle"
              class="relative px-6 py-3 bg-gradient-to-r from-emerald-500 to-green-600 hover:from-emerald-600 hover:to-green-700 text-white rounded-xl font-medium transition-all duration-300 shadow-lg shadow-emerald-500/30 hover:shadow-emerald-500/50 hover:-translate-y-1 group overflow-hidden"
            >
              <!-- Shimmer effect -->
              <div class="absolute inset-0 -top-1 -left-1 bg-gradient-to-r from-transparent via-white/20 to-transparent skew-x-12 transform -translate-x-full group-hover:translate-x-full transition-transform duration-700"></div>
              <span class="relative">Get Started</span>
            </button>
          </div>

          <!-- Mobile menu button -->
          <button
            @click="toggleMobileMenu"
            class="md:hidden p-3 text-gray-700 hover:text-emerald-600 hover:bg-white/20 rounded-xl transition-all duration-200 backdrop-blur-sm"
          >
            <svg class="w-6 h-6 transition-transform duration-300" :class="{ 'rotate-90': mobileMenuOpen }" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" :d="mobileMenuOpen ? 'M6 18L18 6M6 6l12 12' : 'M4 6h16M4 12h16M4 18h16'" />
            </svg>
          </button>
        </div>

        <!-- Mobile Menu with slide animation -->
        <div v-if="mobileMenuOpen" class="md:hidden mt-6 pb-4 animate-slide-down">
          <div class="flex flex-col space-y-4 bg-white/80 backdrop-blur-md rounded-2xl p-6 border border-white/30 shadow-xl">
            <a href="#features" @click="closeMobileMenu" class="text-gray-700 hover:text-emerald-600 py-3 px-4 rounded-lg hover:bg-emerald-50 transition-all duration-200">Features</a>
            <button
              @click="loginWithGoogle"
              class="w-full px-6 py-3 bg-gradient-to-r from-emerald-500 to-green-600 hover:from-emerald-600 hover:to-green-700 text-white rounded-xl font-medium transition-all duration-300 shadow-lg"
            >
              Get Started
            </button>
          </div>
        </div>
      </nav>

      <!-- Enhanced Hero Section with Seamless Flow -->
      <section class="flex-grow flex items-center px-4 sm:px-6 pt-8 pb-20 sm:pb-32">
        <div class="max-w-6xl mx-auto">
          <div class="grid lg:grid-cols-2 gap-16 items-center">
            <!-- Hero Content -->
            <div class="text-center lg:text-left space-y-8">
              <!-- Animated badge -->
              <div class="inline-flex items-center space-x-2 bg-gradient-to-r from-emerald-50 to-green-50 border border-emerald-200/50 text-emerald-700 px-4 py-2 rounded-full text-sm font-medium animate-bounce-gentle">
                <div class="w-2 h-2 bg-emerald-500 rounded-full animate-pulse"></div>
                <span>✨ Now in Beta - Free Forever!</span>
              </div>

              <div class="animate-fade-in-up">
                <h1 class="text-4xl sm:text-5xl lg:text-6xl xl:text-7xl font-extrabold text-gray-900 leading-tight tracking-tight">
                  Launch Your
                  <span class="block bg-gradient-to-r from-emerald-500 via-green-500 to-emerald-600 bg-clip-text text-transparent animate-gradient-x">
                    Dream Shop
                  </span>
                  <span class="block text-3xl sm:text-4xl lg:text-5xl xl:text-6xl text-gray-600 font-normal mt-2">
                    in just minutes ⚡
                  </span>
                </h1>

                <p class="text-xl text-gray-600 max-w-xl mx-auto lg:mx-0 leading-relaxed mt-6">
                  The simplest way to start selling online. Beautiful designs, powerful features, and
                  <span class="font-semibold text-emerald-600">zero coding required</span>.
                </p>
              </div>

              <!-- Enhanced CTA Section -->
              <div class="animate-fade-in-up animation-delay-300 space-y-6">
                <button
                  @click="loginWithGoogle"
                  :disabled="loading"
                  class="group relative w-full sm:w-auto bg-gradient-to-r from-emerald-500 via-green-500 to-emerald-600 hover:from-emerald-600 hover:via-green-600 hover:to-emerald-700 text-white rounded-2xl py-5 px-10 font-bold text-lg transition-all duration-300 shadow-xl shadow-emerald-500/30 hover:shadow-2xl hover:shadow-emerald-500/40 hover:-translate-y-2 disabled:opacity-60 flex items-center justify-center space-x-3 overflow-hidden"
                >
                  <!-- Animated background -->
                  <div class="absolute inset-0 bg-gradient-to-r from-emerald-400 to-green-400 opacity-0 group-hover:opacity-100 transition-opacity duration-300"></div>
                  <!-- Shimmer effect -->
                  <div class="absolute inset-0 -top-1 -left-1 bg-gradient-to-r from-transparent via-white/30 to-transparent skew-x-12 transform -translate-x-full group-hover:translate-x-full transition-transform duration-1000"></div>

                  <svg v-if="loading" class="animate-spin h-7 w-7 relative z-10" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24">
                    <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"/>
                    <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"/>
                  </svg>
                  <img v-else src="https://www.svgrepo.com/show/355037/google.svg" alt="Google icon" class="h-7 w-7 group-hover:scale-110 transition-transform duration-200 relative z-10" />
                  <span class="relative z-10">{{ loading ? 'Creating your shop...' : 'Start Building for Free' }}</span>
                </button>

                <!-- Trust indicators -->
                <div class="flex flex-col sm:flex-row items-center justify-center lg:justify-start space-y-2 sm:space-y-0 sm:space-x-6 text-sm text-gray-500">
                  <div class="flex items-center space-x-2">
                    <svg class="w-4 h-4 text-emerald-500" fill="currentColor" viewBox="0 0 20 20">
                      <path fill-rule="evenodd" d="M10 18a8 8 0 100-16 8 8 0 000 16zm3.707-9.293a1 1 0 00-1.414-1.414L9 10.586 7.707 9.293a1 1 0 00-1.414 1.414l2 2a1 1 0 001.414 0l4-4z" clip-rule="evenodd" />
                    </svg>
                    <span>No credit card required</span>
                  </div>
                  <div class="flex items-center space-x-2">
                    <svg class="w-4 h-4 text-emerald-500" fill="currentColor" viewBox="0 0 20 20">
                      <path fill-rule="evenodd" d="M10 18a8 8 0 100-16 8 8 0 000 16zm3.707-9.293a1 1 0 00-1.414-1.414L9 10.586 7.707 9.293a1 1 0 00-1.414 1.414l2 2a1 1 0 001.414 0l4-4z" clip-rule="evenodd" />
                    </svg>
                    <span>5-minute setup</span>
                  </div>
                  <div class="flex items-center space-x-2">
                    <svg class="w-4 h-4 text-emerald-500" fill="currentColor" viewBox="0 0 20 20">
                      <path fill-rule="evenodd" d="M10 18a8 8 0 100-16 8 8 0 000 16zm3.707-9.293a1 1 0 00-1.414-1.414L9 10.586 7.707 9.293a1 1 0 00-1.414 1.414l2 2a1 1 0 001.414 0l4-4z" clip-rule="evenodd" />
                    </svg>
                    <span>Free forever</span>
                  </div>
                </div>
              </div>

              <!-- Error State -->
              <div v-if="error" class="animate-fade-in-up">
                <div class="bg-red-50 border border-red-200 rounded-xl p-4 shadow-sm">
                  <div class="flex items-center">
                    <svg class="w-5 h-5 text-red-500 mr-3" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-2.5L13.732 4c-.77-.833-1.964-.833-2.732 0L3.732 16.5c-.77.833.192 2.5 1.732 2.5z" />
                    </svg>
                    <p class="text-red-700 font-medium">{{ error }}</p>
                  </div>
                </div>
              </div>
            </div>

            <!-- Enhanced Live Store Mockup - Desktop Only -->
            <div class="hidden lg:block lg:order-last animate-fade-in-up animation-delay-500">
              <div class="relative group">
                <!-- Glow effect -->
                <div class="absolute -inset-1 bg-gradient-to-r from-emerald-500 to-green-500 rounded-3xl blur opacity-25 group-hover:opacity-40 transition duration-1000 group-hover:duration-200 animate-pulse"></div>

                <!-- Main mockup container -->
                <div class="relative bg-white/90 backdrop-blur-sm rounded-2xl shadow-2xl shadow-emerald-500/10 p-8 border border-white/50 hover:shadow-3xl hover:shadow-emerald-500/20 transition-all duration-500 group-hover:scale-105">
                  <!-- Browser header -->
                  <div class="flex items-center space-x-3 mb-6 pb-4 border-b border-gray-200/50">
                    <div class="flex space-x-2">
                      <div class="w-3 h-3 bg-red-400 rounded-full animate-pulse"></div>
                      <div class="w-3 h-3 bg-yellow-400 rounded-full animate-pulse" style="animation-delay: 0.2s"></div>
                      <div class="w-3 h-3 bg-green-400 rounded-full animate-pulse" style="animation-delay: 0.4s"></div>
                    </div>
                    <div class="flex-1 bg-gradient-to-r from-gray-100 to-gray-50 rounded-lg px-4 py-2 text-xs text-gray-500 text-center font-mono border">
                      <span class="text-emerald-600 font-semibold">yourstore</span>.shop24.sbs
                    </div>
                  </div>

                  <!-- Store preview -->
                  <div class="space-y-4">
                    <!-- Dynamic Store Header -->
                    <div :class="currentTheme.headerBg" class="h-20 rounded-xl flex items-center justify-center relative overflow-hidden transition-all duration-500">
                      <!-- Animated background pattern -->
                      <div class="absolute inset-0 bg-[radial-gradient(circle_at_50%_50%,rgba(16,185,129,0.1)_0%,transparent_50%)] animate-pulse"></div>
                      <div class="text-center relative z-10">
                        <div class="w-10 h-10 bg-gradient-to-br from-emerald-500 to-green-600 rounded-lg mx-auto mb-1 flex items-center justify-center shadow-lg shadow-emerald-500/30 group-hover:scale-110 transition-transform duration-300">
                          <svg class="w-5 h-5 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 11V7a4 4 0 00-8 0v4M5 9h14l1 12H4L5 9z" />
                          </svg>
                        </div>
                        <div class="text-xs font-bold text-emerald-700">{{ currentTheme.name }} Store</div>
                      </div>
                    </div>

                    <!-- Dynamic Product Grid -->
                    <div class="grid grid-cols-2 gap-2">
                      <!-- Product 1 -->
                      <div class="bg-white rounded-lg border border-gray-200/50 p-2 hover:shadow-md transition-all duration-500 group cursor-pointer">
                        <div :class="currentTheme.productBg1" class="h-10 rounded mb-1 flex items-center justify-center transition-all duration-500">
                          <span :class="currentTheme.productIcon1" class="text-xs transition-all duration-500">{{ currentTheme.icon1 }}</span>
                        </div>
                        <div class="text-xs font-medium text-gray-700 truncate mb-1">{{ currentTheme.product1 }}</div>
                        <div class="flex items-center justify-between">
                          <div class="text-xs text-emerald-600 font-semibold">${{ currentTheme.price1 }}</div>
                          <button class="bg-emerald-500 hover:bg-emerald-600 text-white text-xs px-2 py-0.5 rounded transition-colors duration-200">
                            +
                          </button>
                        </div>
                      </div>

                      <!-- Product 2 -->
                      <div class="bg-white rounded-lg border border-gray-200/50 p-2 hover:shadow-md transition-all duration-500 group cursor-pointer">
                        <div :class="currentTheme.productBg2" class="h-10 rounded mb-1 flex items-center justify-center transition-all duration-500">
                          <span :class="currentTheme.productIcon2" class="text-xs transition-all duration-500">{{ currentTheme.icon2 }}</span>
                        </div>
                        <div class="text-xs font-medium text-gray-700 truncate mb-1">{{ currentTheme.product2 }}</div>
                        <div class="flex items-center justify-between">
                          <div class="text-xs text-emerald-600 font-semibold">${{ currentTheme.price2 }}</div>
                          <button class="bg-emerald-500 hover:bg-emerald-600 text-white text-xs px-2 py-0.5 rounded transition-colors duration-200">
                            +
                          </button>
                        </div>
                      </div>

                      <!-- Product 3 -->
                      <div class="bg-white rounded-lg border border-gray-200/50 p-2 hover:shadow-md transition-all duration-500 group cursor-pointer">
                        <div :class="currentTheme.productBg3" class="h-10 rounded mb-1 flex items-center justify-center transition-all duration-500">
                          <span :class="currentTheme.productIcon3" class="text-xs transition-all duration-500">{{ currentTheme.icon3 }}</span>
                        </div>
                        <div class="text-xs font-medium text-gray-700 truncate mb-1">{{ currentTheme.product3 }}</div>
                        <div class="flex items-center justify-between">
                          <div class="text-xs text-emerald-600 font-semibold">${{ currentTheme.price3 }}</div>
                          <button class="bg-emerald-500 hover:bg-emerald-600 text-white text-xs px-2 py-0.5 rounded transition-colors duration-200">
                            +
                          </button>
                        </div>
                      </div>

                      <!-- Product 4 -->
                      <div class="bg-white rounded-lg border border-gray-200/50 p-2 hover:shadow-md transition-all duration-500 group cursor-pointer">
                        <div :class="currentTheme.productBg4" class="h-10 rounded mb-1 flex items-center justify-center transition-all duration-500">
                          <span :class="currentTheme.productIcon4" class="text-xs transition-all duration-500">{{ currentTheme.icon4 }}</span>
                        </div>
                        <div class="text-xs font-medium text-gray-700 truncate mb-1">{{ currentTheme.product4 }}</div>
                        <div class="flex items-center justify-between">
                          <div class="text-xs text-emerald-600 font-semibold">${{ currentTheme.price4 }}</div>
                          <button class="bg-emerald-500 hover:bg-emerald-600 text-white text-xs px-2 py-0.5 rounded transition-colors duration-200">
                            +
                          </button>
                        </div>
                      </div>
                    </div>

                    <!-- Theme indicator -->
                    <div class="text-center">
                      <div class="inline-flex items-center space-x-2 bg-white/80 backdrop-blur-sm px-3 py-1 rounded-full border border-gray-200/50">
                        <div class="w-2 h-2 bg-emerald-500 rounded-full animate-pulse"></div>
                        <span class="text-xs text-gray-600 font-medium">{{ currentTheme.name }} Store</span>
                      </div>
                    </div>
                  </div>
                </div>

                <!-- Enhanced floating badges -->
                <div class="absolute -top-4 -right-4 bg-gradient-to-r from-green-500 to-emerald-500 text-white px-3 py-2 rounded-full text-xs font-bold animate-bounce shadow-lg">
                  ✨ Live!
                </div>
                <div class="absolute -bottom-4 -left-4 bg-gradient-to-r from-amber-500 to-orange-500 text-white px-3 py-2 rounded-full text-xs font-bold animate-pulse shadow-lg">
                  💰 Free
                </div>
                <div class="absolute top-1/2 -left-6 bg-gradient-to-r from-blue-500 to-purple-500 text-white px-2 py-1 rounded-full text-xs font-bold animate-bounce-slow shadow-lg">
                  🚀 Fast
                </div>
              </div>
            </div>

            <!-- Mobile Simple Visual -->
            <div class="lg:hidden flex justify-center mt-8">
              <div class="relative">
                <div class="w-64 h-40 bg-gradient-to-br from-emerald-100 to-green-100 rounded-2xl flex items-center justify-center shadow-xl border border-emerald-200/50">
                  <div class="text-center">
                    <div class="w-16 h-16 bg-gradient-to-br from-emerald-500 to-green-600 rounded-2xl mx-auto mb-3 flex items-center justify-center shadow-lg">
                      <svg class="w-8 h-8 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 11V7a4 4 0 00-8 0v4M5 9h14l1 12H4L5 9z" />
                      </svg>
                    </div>
                    <div class="text-lg font-bold text-emerald-700">Your Store</div>
                    <div class="text-sm text-emerald-600">Ready in minutes!</div>
                  </div>
                </div>

                <!-- Simple floating badges for mobile -->
                <div class="absolute -top-2 -right-2 bg-green-500 text-white px-2 py-1 rounded-full text-xs font-bold">
                  ✨ Live
                </div>
                <div class="absolute -bottom-2 -left-2 bg-amber-500 text-white px-2 py-1 rounded-full text-xs font-bold">
                  Free
                </div>
              </div>
            </div>
          </div>
        </div>
      </section>
    </div>

    <!-- Enhanced Features Section -->
    <section id="features" class="py-20 sm:py-28 px-4 sm:px-6 bg-gradient-to-br from-gray-50 via-white to-emerald-50/30 relative overflow-hidden">
      <!-- Background decoration -->
      <div class="absolute inset-0 bg-[radial-gradient(circle_at_70%_30%,rgba(16,185,129,0.1)_0%,transparent_50%)]"></div>

      <div class="max-w-6xl mx-auto relative z-10">
        <div class="text-center mb-16 animate-fade-in-up">
          <h2 class="text-4xl sm:text-5xl font-extrabold text-gray-900 mb-6 tracking-tight">
            Everything you need to
            <span class="bg-gradient-to-r from-emerald-500 to-green-600 bg-clip-text text-transparent">succeed</span>
          </h2>
          <p class="text-xl text-gray-600 max-w-3xl mx-auto leading-relaxed">
            Powerful tools designed to help you build, launch, and grow your online business with ease.
          </p>
        </div>

        <div class="grid gap-8 md:grid-cols-3">
          <FeatureCard
            icon="M4 5a1 1 0 011-1h14a1 1 0 011 1v2a1 1 0 01-1 1H5a1 1 0 01-1-1V5zM4 13a1 1 0 011-1h6a1 1 0 011 1v6a1 1 0 01-1 1H5a1 1 0 01-1-1v-6zM16 13a1 1 0 011-1h2a1 1 0 011 1v6a1 1 0 01-1 1h-2a1 1 0 01-1-1v-6z"
            title="🎨 Beautiful Stores"
            text="Create stunning, mobile-responsive storefronts with our intuitive drag-and-drop designer. No coding skills required!"
            color="green"
            :delay="0"
          />
          <FeatureCard
            icon="M17 9V7a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2m2 4h10a2 2 0 002-2v-6a2 2 0 00-2-2H9a2 2 0 00-2 2v6a2 2 0 002 2zm7-5a2 2 0 11-4 0 2 2 0 014 0z"
            title="� Cash on Delivery"
            text="Simple and trusted payment method. Your customers pay when they receive their orders. No online payment hassles!"
            color="amber"
            :delay="200"
          />
          <FeatureCard
            icon="M9 19v-6a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2a2 2 0 002-2zm0 0V9a2 2 0 012-2h2a2 2 0 012 2v10m-6 0a2 2 0 002 2h2a2 2 0 002-2m0 0V5a2 2 0 012-2h2a2 2 0 012 2v14a2 2 0 01-2 2h-2a2 2 0 01-2-2z"
            title="📊 Smart Analytics"
            text="Track sales, monitor customer behavior, and get actionable insights to grow your business with real-time dashboards."
            color="green"
            :delay="400"
          />
        </div>

        <!-- Additional features grid -->
        <div class="grid gap-6 md:grid-cols-2 lg:grid-cols-4 mt-16">
          <div class="text-center p-6 rounded-xl bg-white/60 backdrop-blur-sm border border-white/50 hover:shadow-lg transition-all duration-300 group">
            <div class="w-12 h-12 bg-gradient-to-br from-blue-500 to-blue-600 rounded-lg mx-auto mb-4 flex items-center justify-center group-hover:scale-110 transition-transform duration-300">
              <span class="text-white text-xl">⚡</span>
            </div>
            <h4 class="font-bold text-gray-900 mb-2">Lightning Fast</h4>
            <p class="text-sm text-gray-600">Sub-second loading times</p>
          </div>

          <div class="text-center p-6 rounded-xl bg-white/60 backdrop-blur-sm border border-white/50 hover:shadow-lg transition-all duration-300 group">
            <div class="w-12 h-12 bg-gradient-to-br from-purple-500 to-purple-600 rounded-lg mx-auto mb-4 flex items-center justify-center group-hover:scale-110 transition-transform duration-300">
              <span class="text-white text-xl">📱</span>
            </div>
            <h4 class="font-bold text-gray-900 mb-2">Mobile First</h4>
            <p class="text-sm text-gray-600">Perfect on every device</p>
          </div>

          <div class="text-center p-6 rounded-xl bg-white/60 backdrop-blur-sm border border-white/50 hover:shadow-lg transition-all duration-300 group">
            <div class="w-12 h-12 bg-gradient-to-br from-pink-500 to-pink-600 rounded-lg mx-auto mb-4 flex items-center justify-center group-hover:scale-110 transition-transform duration-300">
              <span class="text-white text-xl">🛡️</span>
            </div>
            <h4 class="font-bold text-gray-900 mb-2">Secure & Reliable</h4>
            <p class="text-sm text-gray-600">Safe data & order management</p>
          </div>

          <div class="text-center p-6 rounded-xl bg-white/60 backdrop-blur-sm border border-white/50 hover:shadow-lg transition-all duration-300 group">
            <div class="w-12 h-12 bg-gradient-to-br from-orange-500 to-orange-600 rounded-lg mx-auto mb-4 flex items-center justify-center group-hover:scale-110 transition-transform duration-300">
              <span class="text-white text-xl">🎯</span>
            </div>
            <h4 class="font-bold text-gray-900 mb-2">SEO Ready</h4>
            <p class="text-sm text-gray-600">Built for search engines</p>
          </div>
        </div>
      </div>
    </section>


    <!-- Enhanced Final CTA Section -->
    <section class="py-20 sm:py-28 bg-gradient-to-br from-emerald-600 via-green-600 to-emerald-700 relative overflow-hidden">
      <!-- Animated background -->
      <div class="absolute inset-0 bg-[radial-gradient(circle_at_30%_20%,rgba(255,255,255,0.1)_0%,transparent_50%),radial-gradient(circle_at_70%_80%,rgba(255,255,255,0.1)_0%,transparent_50%)]"></div>
      <div class="absolute inset-0 bg-[radial-gradient(#ffffff_1px,transparent_1px)] [background-size:32px_32px] opacity-10 animate-pulse"></div>

      <div class="max-w-4xl mx-auto text-center px-4 sm:px-6 relative z-10">
        <div class="animate-fade-in-up">
          <h3 class="text-4xl sm:text-5xl lg:text-6xl font-extrabold text-white mb-6 tracking-tight">
            Ready to launch your
            <span class="block bg-gradient-to-r from-yellow-300 to-orange-300 bg-clip-text text-transparent animate-gradient-x">
              success story?
            </span>
          </h3>
          <p class="text-xl text-emerald-100 mb-10 max-w-2xl mx-auto leading-relaxed">
            Join the next generation of entrepreneurs building amazing online stores.
            Your journey starts with a single click! 🚀
          </p>

          <button
            @click="loginWithGoogle"
            :disabled="loading"
            class="group relative bg-white text-emerald-600 hover:text-emerald-700 rounded-2xl py-5 px-12 font-bold text-xl transition-all duration-300 shadow-2xl hover:shadow-3xl hover:-translate-y-2 disabled:opacity-60 flex items-center justify-center space-x-4 mx-auto overflow-hidden"
          >
            <!-- Animated background -->
            <div class="absolute inset-0 bg-gradient-to-r from-emerald-50 to-green-50 opacity-0 group-hover:opacity-100 transition-opacity duration-300"></div>
            <!-- Shimmer effect -->
            <div class="absolute inset-0 -top-1 -left-1 bg-gradient-to-r from-transparent via-emerald-200/30 to-transparent skew-x-12 transform -translate-x-full group-hover:translate-x-full transition-transform duration-1000"></div>

            <svg v-if="loading" class="animate-spin h-7 w-7 relative z-10" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24">
              <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"/>
              <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"/>
            </svg>
            <img v-else src="https://www.svgrepo.com/show/355037/google.svg" alt="Google icon" class="h-7 w-7 group-hover:scale-110 transition-transform duration-200 relative z-10" />
            <span class="relative z-10">{{ loading ? 'Creating magic...' : 'Start Your Journey - Free!' }}</span>
          </button>

          <!-- Trust indicators -->
          <div class="flex flex-wrap items-center justify-center gap-6 mt-8 text-emerald-100 text-sm">
            <div class="flex items-center space-x-2">
              <span class="text-yellow-300">⭐</span>
              <span>No credit card required</span>
            </div>
            <div class="flex items-center space-x-2">
              <span class="text-yellow-300">⚡</span>
              <span>Setup in minutes</span>
            </div>
            <div class="flex items-center space-x-2">
              <span class="text-yellow-300">💎</span>
              <span>Free forever plan</span>
            </div>
          </div>
        </div>
      </div>
    </section>

    <!-- Enhanced Footer -->
    <footer class="bg-gradient-to-br from-gray-900 via-gray-800 to-gray-900 text-gray-300 py-12 relative overflow-hidden">
      <!-- Background decoration -->
      <div class="absolute inset-0 bg-[radial-gradient(circle_at_50%_50%,rgba(16,185,129,0.1)_0%,transparent_50%)]"></div>

      <div class="max-w-6xl mx-auto px-4 sm:px-6 relative z-10">
        <div class="flex flex-col sm:flex-row items-center justify-between space-y-6 sm:space-y-0">
          <div class="flex items-center space-x-4 group">
            <div class="w-10 h-10 bg-gradient-to-br from-emerald-500 to-green-600 rounded-xl flex items-center justify-center shadow-lg shadow-emerald-500/30 group-hover:scale-110 group-hover:rotate-3 transition-all duration-300">
               <svg class="w-6 h-6 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 7h8m0 0v8m0-8l-8 8-4-4-6 6" />
              </svg>
            </div>
            <div>
              <span class="text-2xl font-bold bg-gradient-to-r from-emerald-400 to-green-400 bg-clip-text text-transparent">Shop24</span>
              <p class="text-xs text-gray-400">Building the future of e-commerce</p>
            </div>
          </div>

          <div class="text-center sm:text-right">
            <div class="text-sm text-gray-400 mb-2">
              © 2024 Shop24. Made with <span class="text-red-400 animate-pulse">♡</span> for creators worldwide.
            </div>
            <div class="flex space-x-4 justify-center sm:justify-end text-xs text-gray-500">
              <a href="#" class="hover:text-emerald-400 transition-colors duration-200">Privacy</a>
              <a href="#" class="hover:text-emerald-400 transition-colors duration-200">Terms</a>
              <a href="#" class="hover:text-emerald-400 transition-colors duration-200">Support</a>
            </div>
          </div>
        </div>
      </div>
    </footer>
  </div>
</template>

<script setup>
import { ref, watch, onMounted, onUnmounted, computed } from 'vue'
import { useRouter, useRoute } from 'vue-router'
import { useAuthStore } from '@/store/auth'
import FeatureCard from '@/components/FeatureCard.vue'

const router = useRouter()
const route = useRoute()
const auth = useAuthStore()
const loading = ref(false)
const error = ref(null)
const mobileMenuOpen = ref(false)
const currentThemeIndex = ref(0)

// Different store themes that cycle automatically
const storeThemes = [
  {
    name: "Fashion",
    headerBg: "bg-gradient-to-br from-pink-100 via-purple-50 to-pink-100",
    productBg1: "bg-gradient-to-br from-pink-100 to-pink-200",
    productBg2: "bg-gradient-to-br from-purple-100 to-purple-200",
    productBg3: "bg-gradient-to-br from-rose-100 to-rose-200",
    productBg4: "bg-gradient-to-br from-indigo-100 to-indigo-200",
    productIcon1: "text-pink-600",
    productIcon2: "text-purple-600",
    productIcon3: "text-rose-600",
    productIcon4: "text-indigo-600",
    icon1: "👗", icon2: "👠", icon3: "👜", icon4: "💄",
    product1: "Dress", product2: "Heels", product3: "Handbag", product4: "Makeup",
    price1: "45", price2: "65", price3: "35", price4: "25"
  },
  {
    name: "Electronics",
    headerBg: "bg-gradient-to-br from-blue-100 via-cyan-50 to-blue-100",
    productBg1: "bg-gradient-to-br from-blue-100 to-blue-200",
    productBg2: "bg-gradient-to-br from-cyan-100 to-cyan-200",
    productBg3: "bg-gradient-to-br from-slate-100 to-slate-200",
    productBg4: "bg-gradient-to-br from-gray-100 to-gray-200",
    productIcon1: "text-blue-600",
    productIcon2: "text-cyan-600",
    productIcon3: "text-slate-600",
    productIcon4: "text-gray-600",
    icon1: "📱", icon2: "💻", icon3: "🎧", icon4: "⌚",
    product1: "Phone", product2: "Laptop", product3: "Headphones", product4: "Watch",
    price1: "299", price2: "899", price3: "89", price4: "199"
  },
  {
    name: "Food & Drinks",
    headerBg: "bg-gradient-to-br from-orange-100 via-yellow-50 to-orange-100",
    productBg1: "bg-gradient-to-br from-orange-100 to-orange-200",
    productBg2: "bg-gradient-to-br from-yellow-100 to-yellow-200",
    productBg3: "bg-gradient-to-br from-red-100 to-red-200",
    productBg4: "bg-gradient-to-br from-green-100 to-green-200",
    productIcon1: "text-orange-600",
    productIcon2: "text-yellow-600",
    productIcon3: "text-red-600",
    productIcon4: "text-green-600",
    icon1: "🍕", icon2: "☕", icon3: "🍰", icon4: "🥗",
    product1: "Pizza", product2: "Coffee", product3: "Cake", product4: "Salad",
    price1: "12", price2: "5", price3: "8", price4: "10"
  },
  {
    name: "Home & Garden",
    headerBg: "bg-gradient-to-br from-green-100 via-emerald-50 to-green-100",
    productBg1: "bg-gradient-to-br from-green-100 to-green-200",
    productBg2: "bg-gradient-to-br from-emerald-100 to-emerald-200",
    productBg3: "bg-gradient-to-br from-teal-100 to-teal-200",
    productBg4: "bg-gradient-to-br from-lime-100 to-lime-200",
    productIcon1: "text-green-600",
    productIcon2: "text-emerald-600",
    productIcon3: "text-teal-600",
    productIcon4: "text-lime-600",
    icon1: "🪴", icon2: "🛋️", icon3: "💡", icon4: "🌸",
    product1: "Plant", product2: "Sofa", product3: "Lamp", product4: "Flowers",
    price1: "25", price2: "450", price3: "35", price4: "15"
  }
]

const currentTheme = computed(() => storeThemes[currentThemeIndex.value])

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

// Mobile menu functions
function toggleMobileMenu() {
  mobileMenuOpen.value = !mobileMenuOpen.value
}

function closeMobileMenu() {
  mobileMenuOpen.value = false
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

// Theme cycling function
function cycleTheme() {
  currentThemeIndex.value = (currentThemeIndex.value + 1) % storeThemes.length
}

onMounted(() => {
  // If user is already authenticated, redirect appropriately
  if (auth.isAuthenticated) {
    const redirect = determineRedirect()
    router.replace(redirect)
  }

  // Start theme cycling animation only on desktop (lg and above)
  const isDesktop = window.innerWidth >= 1024 // lg breakpoint
  let themeInterval = null

  if (isDesktop) {
    themeInterval = setInterval(cycleTheme, 4000) // Change theme every 4 seconds
  }

  // Cleanup interval on unmount
  onUnmounted(() => {
    if (themeInterval) {
      clearInterval(themeInterval)
    }
  })
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
/* Modern fonts with better weights */
@import url('https://fonts.googleapis.com/css2?family=Inter:wght@300;400;500;600;700;800;900&display=swap');

body {
  font-family: 'Inter', sans-serif;
}

/* Enhanced animations */
@keyframes fadeInUp {
  from {
    opacity: 0;
    transform: translateY(30px);
  }
  to {
    opacity: 1;
    transform: translateY(0);
  }
}

@keyframes fadeInDown {
  from {
    opacity: 0;
    transform: translateY(-20px);
  }
  to {
    opacity: 1;
    transform: translateY(0);
  }
}

@keyframes slideDown {
  from {
    opacity: 0;
    transform: translateY(-10px);
  }
  to {
    opacity: 1;
    transform: translateY(0);
  }
}

@keyframes float {
  0%, 100% {
    transform: translateY(0px);
  }
  50% {
    transform: translateY(-20px);
  }
}

@keyframes float-delayed {
  0%, 100% {
    transform: translateY(0px);
  }
  50% {
    transform: translateY(-15px);
  }
}

@keyframes float-slow {
  0%, 100% {
    transform: translateY(0px);
  }
  50% {
    transform: translateY(-10px);
  }
}

@keyframes bounce-gentle {
  0%, 100% {
    transform: translateY(0);
  }
  50% {
    transform: translateY(-5px);
  }
}

@keyframes bounce-slow {
  0%, 100% {
    transform: translateY(0);
  }
  50% {
    transform: translateY(-8px);
  }
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

@keyframes pulse-glow {
  0%, 100% {
    box-shadow: 0 0 20px rgba(16, 185, 129, 0.3);
  }
  50% {
    box-shadow: 0 0 40px rgba(16, 185, 129, 0.6);
  }
}

/* Animation classes */
.animate-fade-in-up {
  animation: fadeInUp 0.8s ease-out forwards;
  opacity: 0;
}

.animate-fade-in-down {
  animation: fadeInDown 0.5s ease-out forwards;
  opacity: 0;
}

.animate-slide-down {
  animation: slideDown 0.3s ease-out forwards;
}

.animate-float {
  animation: float 6s ease-in-out infinite;
}

.animate-float-delayed {
  animation: float-delayed 6s ease-in-out infinite;
  animation-delay: 2s;
}

.animate-float-slow {
  animation: float-slow 8s ease-in-out infinite;
  animation-delay: 4s;
}

.animate-bounce-gentle {
  animation: bounce-gentle 2s ease-in-out infinite;
}

.animate-bounce-slow {
  animation: bounce-slow 3s ease-in-out infinite;
}

.animate-gradient-x {
  animation: gradient-x 3s ease infinite;
}

.animate-pulse-glow {
  animation: pulse-glow 2s ease-in-out infinite;
}

/* Staggered animation delays */
.animation-delay-100 { animation-delay: 0.1s; }
.animation-delay-200 { animation-delay: 0.2s; }
.animation-delay-300 { animation-delay: 0.3s; }
.animation-delay-400 { animation-delay: 0.4s; }
.animation-delay-500 { animation-delay: 0.5s; }

/* Custom shadows */
.shadow-3xl {
  box-shadow: 0 35px 60px -12px rgba(0, 0, 0, 0.25);
}

/* Smooth scrolling */
html {
  scroll-behavior: smooth;
}

/* Custom scrollbar */
::-webkit-scrollbar {
  width: 8px;
}

::-webkit-scrollbar-track {
  background: #f1f5f9;
}

::-webkit-scrollbar-thumb {
  background: linear-gradient(to bottom, #10b981, #059669);
  border-radius: 4px;
}

::-webkit-scrollbar-thumb:hover {
  background: linear-gradient(to bottom, #059669, #047857);
}

/* Focus styles for accessibility */
button:focus-visible,
a:focus-visible {
  outline: 2px solid #10b981;
  outline-offset: 2px;
}

/* Accessibility improvements */
@media (prefers-reduced-motion: reduce) {
  * {
    animation-duration: 0.01ms !important;
    animation-iteration-count: 1 !important;
    transition-duration: 0.01ms !important;
  }
}

/* Mobile optimizations */
@media (max-width: 768px) {
  .animate-fade-in-up {
    animation-duration: 0.6s;
  }

  .animate-slide-down {
    animation-duration: 0.2s;
  }
}
</style>
