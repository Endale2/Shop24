<template>
  <div class="min-h-full bg-gray-50">
    <div class="p-4 sm:p-6 lg:p-8">
      
      <!-- Header Section -->
      <div class="mb-6 sm:mb-8">
        <div class="flex items-center mb-3">
          <div class="w-10 h-10 bg-gradient-to-br from-green-500 to-green-600 rounded-lg flex items-center justify-center mr-3 shadow-sm">
            <svg class="w-5 h-5 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10.325 4.317c.426-1.756 2.924-1.756 3.35 0a1.724 1.724 0 002.573 1.066c1.543-.94 3.31.826 2.37 2.37a1.724 1.724 0 001.065 2.572c1.756.426 1.756 2.924 0 3.35a1.724 1.724 0 00-1.066 2.573c.94 1.543-.826 3.31-2.37 2.37a1.724 1.724 0 00-2.572 1.065c-.426 1.756-2.924 1.756-3.35 0a1.724 1.724 0 00-2.573-1.066c-1.543.94-3.31-.826-2.37-2.37a1.724 1.724 0 00-1.065-2.572c-1.756-.426-1.756-2.924 0-3.35a1.724 1.724 0 001.066-2.573c-.94-1.543.826-3.31 2.37-2.37.996.608 2.296.07 2.572-1.065z" />
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z" />
            </svg>
          </div>
          <div>
            <h1 class="text-2xl sm:text-3xl font-bold text-gray-900 tracking-tight">
              Shop Settings
            </h1>
            <p class="text-sm text-gray-600 mt-1">
              Configure your shop appearance and preferences
            </p>
          </div>
        </div>
      </div>

      <!-- Loading State -->
      <div v-if="saving" class="text-center py-16 text-gray-500">
        <div class="flex flex-col items-center justify-center">
          <div class="animate-spin h-8 w-8 text-green-500 mb-3 border-3 border-green-200 border-t-green-500 rounded-full"></div>
          <p class="text-sm font-medium">Saving changes...</p>
        </div>
      </div>

      <div v-else class="space-y-6">
        <!-- Basic Information Card -->
        <div class="bg-white rounded-lg shadow-sm border border-gray-200 overflow-hidden">
          <div class="px-4 py-3 border-b border-gray-100 bg-gray-50">
            <div class="flex items-center">
              <div class="w-8 h-8 bg-gradient-to-br from-blue-500 to-blue-600 rounded-lg flex items-center justify-center mr-3">
                <svg class="w-4 h-4 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 16h-1v-4h-1m1-4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z" />
                </svg>
              </div>
              <div>
                <h2 class="text-lg font-semibold text-gray-900">Basic Information</h2>
                <p class="text-sm text-gray-600">Configure your shop's essential details</p>
              </div>
            </div>
          </div>
          <div class="p-6">
            <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
              <!-- Shop Name -->
              <div>
                <label for="name" class="block text-sm font-medium text-gray-700 mb-2">Shop Name *</label>
                <input
                  id="name"
                  v-model="form.name"
                  type="text"
                  required
                  maxlength="50"
                  class="block w-full border-gray-300 rounded-lg shadow-sm focus:ring-green-500 focus:border-green-500 transition duration-150 ease-in-out"
                  :class="{ 'border-red-500 focus:border-red-500 focus:ring-red-500': errors.name }"
                />
                <p v-if="errors.name" class="mt-1 text-sm text-red-600 flex items-center">
                  <svg class="w-4 h-4 mr-1" fill="currentColor" viewBox="0 0 20 20">
                    <path fill-rule="evenodd" d="M18 10a8 8 0 11-16 0 8 8 0 0116 0zm-7 4a1 1 0 11-2 0 1 1 0 012 0zm-1-9a1 1 0 00-1 1v4a1 1 0 102 0V6a1 1 0 00-1-1z" clip-rule="evenodd" />
                  </svg>
                  {{ errors.name }}
                </p>
                <p class="mt-1 text-sm text-gray-500">{{ form.name.length }}/50 characters</p>
              </div>

              <!-- Contact Email -->
              <div>
                <label for="email" class="block text-sm font-medium text-gray-700 mb-2">Contact Email *</label>
                <input
                  id="email"
                  v-model="form.email"
                  type="email"
                  required
                  class="block w-full border-gray-300 rounded-lg shadow-sm focus:ring-green-500 focus:border-green-500 transition duration-150 ease-in-out"
                  :class="{ 'border-red-500 focus:border-red-500 focus:ring-red-500': errors.email }"
                />
                <p v-if="errors.email" class="mt-1 text-sm text-red-600 flex items-center">
                  <svg class="w-4 h-4 mr-1" fill="currentColor" viewBox="0 0 20 20">
                    <path fill-rule="evenodd" d="M18 10a8 8 0 11-16 0 8 8 0 0116 0zm-7 4a1 1 0 11-2 0 1 1 0 012 0zm-1-9a1 1 0 00-1 1v4a1 1 0 102 0V6a1 1 0 00-1-1z" clip-rule="evenodd" />
                  </svg>
                  {{ errors.email }}
                </p>
              </div>

              <!-- Description -->
              <div class="md:col-span-2">
                <label for="description" class="block text-sm font-medium text-gray-700 mb-2">Description</label>
                <textarea
                  id="description"
                  v-model="form.description"
                  rows="3"
                  maxlength="200"
                  placeholder="Tell customers about your shop..."
                  class="block w-full border-gray-300 rounded-lg shadow-sm focus:ring-green-500 focus:border-green-500 transition duration-150 ease-in-out"
                  :class="{ 'border-red-500 focus:border-red-500 focus:ring-red-500': errors.description }"
                />
                <p v-if="errors.description" class="mt-1 text-sm text-red-600 flex items-center">
                  <svg class="w-4 h-4 mr-1" fill="currentColor" viewBox="0 0 20 20">
                    <path fill-rule="evenodd" d="M18 10a8 8 0 11-16 0 8 8 0 0116 0zm-7 4a1 1 0 11-2 0 1 1 0 012 0zm-1-9a1 1 0 00-1 1v4a1 1 0 102 0V6a1 1 0 00-1-1z" clip-rule="evenodd" />
                  </svg>
                  {{ errors.description }}
                </p>
                <p class="mt-1 text-sm text-gray-500">{{ form.description.length }}/200 characters</p>
              </div>
            </div>
          </div>
        </div>

        <!-- Visual Assets Card -->
        <div class="bg-white rounded-lg shadow-sm border border-gray-200 overflow-hidden">
          <div class="px-4 py-3 border-b border-gray-100 bg-gray-50">
            <div class="flex items-center">
              <div class="w-8 h-8 bg-gradient-to-br from-purple-500 to-purple-600 rounded-lg flex items-center justify-center mr-3">
                <svg class="w-4 h-4 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 16l4.586-4.586a2 2 0 012.828 0L16 16m-2-2l1.586-1.586a2 2 0 012.828 0L20 14m-6-6h.01M6 20h12a2 2 0 002-2V6a2 2 0 00-2-2H6a2 2 0 00-2 2v12a2 2 0 002 2z" />
                </svg>
              </div>
              <div>
                <h2 class="text-lg font-semibold text-gray-900">Visual Assets</h2>
                <p class="text-sm text-gray-600">Upload your shop's logo and banner images</p>
              </div>
            </div>
          </div>
          <div class="p-6">
            <div class="grid grid-cols-1 lg:grid-cols-2 gap-8">
              
              <!-- Logo Section -->
              <div class="space-y-4">
                <div class="flex items-center space-x-2 mb-4">
                  <div class="w-6 h-6 bg-gradient-to-br from-green-500 to-green-600 rounded-md flex items-center justify-center">
                    <svg class="w-3 h-3 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 3v4M3 5h4M6 17v4m-2-2h4m5-16l2.286 6.857L21 12l-5.714 2.143L13 21l-2.286-6.857L5 12l5.714-2.143L13 3z" />
                    </svg>
                  </div>
                  <h3 class="text-md font-semibold text-gray-800">Shop Logo</h3>
                </div>
                
                <!-- Logo URL Input -->
                <div>
                  <label for="logoUrl" class="block text-sm font-medium text-gray-700 mb-2">Logo URL</label>
                  <input
                    id="logoUrl"
                    v-model="form.logoUrl"
                    type="url"
                    placeholder="https://example.com/logo.png"
                    class="block w-full border-gray-300 rounded-lg shadow-sm focus:ring-green-500 focus:border-green-500 transition duration-150 ease-in-out"
                    :class="{ 'border-red-500 focus:border-red-500 focus:ring-red-500': errors.logoUrl }"
                  />
                  <p v-if="errors.logoUrl" class="mt-1 text-sm text-red-600 flex items-center">
                    <svg class="w-4 h-4 mr-1" fill="currentColor" viewBox="0 0 20 20">
                      <path fill-rule="evenodd" d="M18 10a8 8 0 11-16 0 8 8 0 0116 0zm-7 4a1 1 0 11-2 0 1 1 0 012 0zm-1-9a1 1 0 00-1 1v4a1 1 0 102 0V6a1 1 0 00-1-1z" clip-rule="evenodd" />
                    </svg>
                    {{ errors.logoUrl }}
                  </p>
                  <p class="mt-1 text-sm text-gray-500">Enter a direct URL to your logo image</p>
                </div>

                <!-- Logo Preview -->
                <div class="bg-gray-50 rounded-lg border-2 border-dashed border-gray-300 p-4 hover:border-gray-400 transition-colors duration-200">
                  <div class="flex flex-col items-center space-y-3">
                    <div class="w-24 h-24 rounded-lg overflow-hidden bg-white shadow-sm border border-gray-200 flex items-center justify-center">
                      <img
                        v-if="form.logoUrl"
                        :src="form.logoUrl"
                        alt="Logo preview"
                        class="max-h-full max-w-full object-contain"
                        @error="errors.logoUrl = 'Invalid image URL'"
                        @load="errors.logoUrl = ''"
                      />
                      <div v-else class="text-gray-400 text-center">
                        <svg class="h-8 w-8 mx-auto mb-1" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                            d="M4 16l4.586-4.586a2 2 0 012.828 0L16 16m-2-2l1.586-1.586a2 2 0 012.828 0L20 14m-6-6h.01M6 20h12a2 2 0 002-2V6a2 2 0 00-2-2H6a2 2 0 00-2 2v12a2 2 0 002 2z" />
                        </svg>
                      </div>
                    </div>
                    <p class="text-xs text-gray-600 text-center">Logo Preview<br/>Recommended: 200x200px</p>
                  </div>
                </div>
              </div>

              <!-- Banner Section -->
              <div class="space-y-4">
                <div class="flex items-center space-x-2 mb-4">
                  <div class="w-6 h-6 bg-gradient-to-br from-indigo-500 to-indigo-600 rounded-md flex items-center justify-center">
                    <svg class="w-3 h-3 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 6a2 2 0 012-2h2a2 2 0 012 2v2a2 2 0 01-2 2H6a2 2 0 01-2-2V6zM14 6a2 2 0 012-2h2a2 2 0 012 2v2a2 2 0 01-2 2h-2a2 2 0 01-2-2V6zM4 16a2 2 0 012-2h2a2 2 0 012 2v2a2 2 0 01-2 2H6a2 2 0 01-2-2v-2zM14 16a2 2 0 012-2h2a2 2 0 012 2v2a2 2 0 01-2 2h-2a2 2 0 01-2-2v-2z" />
                    </svg>
                  </div>
                  <h3 class="text-md font-semibold text-gray-800">Shop Banner</h3>
                </div>
                
                <!-- Banner URL Input -->
                <div>
                  <label for="bannerUrl" class="block text-sm font-medium text-gray-700 mb-2">Banner URL</label>
                  <input
                    id="bannerUrl"
                    v-model="form.bannerUrl"
                    type="url"
                    placeholder="https://example.com/banner.jpg"
                    class="block w-full border-gray-300 rounded-lg shadow-sm focus:ring-green-500 focus:border-green-500 transition duration-150 ease-in-out"
                    :class="{ 'border-red-500 focus:border-red-500 focus:ring-red-500': errors.bannerUrl }"
                  />
                  <p v-if="errors.bannerUrl" class="mt-1 text-sm text-red-600 flex items-center">
                    <svg class="w-4 h-4 mr-1" fill="currentColor" viewBox="0 0 20 20">
                      <path fill-rule="evenodd" d="M18 10a8 8 0 11-16 0 8 8 0 0116 0zm-7 4a1 1 0 11-2 0 1 1 0 012 0zm-1-9a1 1 0 00-1 1v4a1 1 0 102 0V6a1 1 0 00-1-1z" clip-rule="evenodd" />
                    </svg>
                    {{ errors.bannerUrl }}
                  </p>
                  <p class="mt-1 text-sm text-gray-500">Enter a direct URL to your banner image</p>
                </div>

                <!-- Banner Preview -->
                <div class="bg-gray-50 rounded-lg border-2 border-dashed border-gray-300 p-4 hover:border-gray-400 transition-colors duration-200">
                  <div class="flex flex-col items-center space-y-3">
                    <div class="w-full h-20 rounded-lg overflow-hidden bg-white shadow-sm border border-gray-200 flex items-center justify-center">
                      <img
                        v-if="form.bannerUrl"
                        :src="form.bannerUrl"
                        alt="Banner preview"
                        class="max-h-full max-w-full object-cover rounded-md"
                        @error="errors.bannerUrl = 'Invalid image URL'"
                        @load="errors.bannerUrl = ''"
                      />
                      <div v-else class="text-gray-400 text-center">
                        <svg class="h-6 w-6 mx-auto mb-1" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                            d="M4 16l4.586-4.586a2 2 0 012.828 0L16 16m-2-2l1.586-1.586a2 2 0 012.828 0L20 14m-6-6h.01M6 20h12a2 2 0 002-2V6a2 2 0 00-2-2H6a2 2 0 00-2 2v12a2 2 0 002 2z" />
                        </svg>
                      </div>
                    </div>
                    <p class="text-xs text-gray-600 text-center">Banner Preview<br/>Recommended: 1200x300px</p>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>

        <!-- Theme Customization Card -->
        <div class="bg-white rounded-lg shadow-sm border border-gray-200 overflow-hidden">
          <div class="px-4 py-3 border-b border-gray-100 bg-gray-50">
            <div class="flex items-center">
              <div class="w-8 h-8 bg-gradient-to-br from-amber-500 to-amber-600 rounded-lg flex items-center justify-center mr-3">
                <svg class="w-4 h-4 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 21a4 4 0 01-4-4V5a2 2 0 012-2h4a2 2 0 012 2v12a4 4 0 01-4 4zM21 5a2 2 0 00-2-2h-4a2 2 0 00-2 2v12a4 4 0 004 4h4a2 2 0 002-2V5z" />
                </svg>
              </div>
              <div>
                <h2 class="text-lg font-semibold text-gray-900">Theme Customization</h2>
                <p class="text-sm text-gray-600">Customize your shop's appearance and settings</p>
              </div>
            </div>
          </div>
          <div class="p-6">
            <div class="grid grid-cols-1 md:grid-cols-2 gap-8">
              <!-- Theme Color Picker -->
              <div class="space-y-4">
                <div class="flex items-center space-x-2 mb-4">
                  <div class="w-6 h-6 bg-gradient-to-br from-pink-500 to-pink-600 rounded-md flex items-center justify-center">
                    <svg class="w-3 h-3 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 21a4 4 0 01-4-4V5a2 2 0 012-2h4a2 2 0 012 2v12a4 4 0 01-4 4zM21 5a2 2 0 00-2-2h-4a2 2 0 00-2 2v12a4 4 0 004 4h4a2 2 0 002-2V5z" />
                    </svg>
                  </div>
                  <h3 class="text-md font-semibold text-gray-800">Primary Color</h3>
                </div>
                
                <div class="bg-gray-50 rounded-lg p-4 border border-gray-200">
                  <div class="flex items-center space-x-4">
                    <div class="relative">
                      <input
                        id="themeColor"
                        v-model="form.themeColor"
                        type="color"
                        class="h-16 w-16 border-2 border-white rounded-xl cursor-pointer shadow-lg hover:scale-105 transition-transform duration-200"
                      />
                      <div class="absolute -bottom-1 -right-1 w-5 h-5 bg-white rounded-full border border-gray-200 flex items-center justify-center">
                        <svg class="w-3 h-3 text-gray-500" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15.232 5.232l3.536 3.536m-2.036-5.036a2.5 2.5 0 113.536 3.536L6.5 21.036H3v-3.572L16.732 3.732z" />
                        </svg>
                      </div>
                    </div>
                    <div class="flex-1 space-y-2">
                      <p class="text-sm font-medium text-gray-700">Choose your shop's primary theme color</p>
                      <div class="flex items-center space-x-2">
                        <div 
                          class="w-4 h-4 rounded-full border border-gray-300"
                          :style="{ backgroundColor: form.themeColor }"
                        ></div>
                        <p class="text-xs font-mono text-gray-500 bg-gray-100 px-2 py-1 rounded">{{ form.themeColor.toUpperCase() }}</p>
                      </div>
                      <p class="text-xs text-gray-500">This color will be used for buttons, links, and highlights</p>
                    </div>
                  </div>
                </div>
              </div>

              <!-- Currency -->
              <div class="space-y-4">
                <div class="flex items-center space-x-2 mb-4">
                  <div class="w-6 h-6 bg-gradient-to-br from-emerald-500 to-emerald-600 rounded-md flex items-center justify-center">
                    <svg class="w-3 h-3 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8c-1.657 0-3 .895-3 2s1.343 2 3 2 3 .895 3 2-1.343 2-3 2m0-8c1.11 0 2.08.402 2.599 1M12 8V7m0 1v8m0 0v1m0-1c-1.11 0-2.08-.402-2.599-1" />
                    </svg>
                  </div>
                  <h3 class="text-md font-semibold text-gray-800">Currency</h3>
                </div>
                
                <div class="space-y-3">
                  <label for="currency" class="block text-sm font-medium text-gray-700">Shop Currency</label>
                  <div class="relative">
                    <select
                      id="currency"
                      v-model="form.currency"
                      class="block w-full border-gray-300 rounded-lg shadow-sm focus:ring-green-500 focus:border-green-500 transition duration-150 ease-in-out appearance-none pl-3 pr-10 py-2.5"
                    >
                      <option value="USD">🇺🇸 USD - US Dollar</option>
                      <option value="EUR">🇪🇺 EUR - Euro</option>
                      <option value="GBP">🇬🇧 GBP - British Pound</option>
                      <option value="CAD">🇨🇦 CAD - Canadian Dollar</option>
                      <option value="AUD">🇦🇺 AUD - Australian Dollar</option>
                    </select>
                    <div class="absolute inset-y-0 right-0 flex items-center px-2 pointer-events-none">
                      <svg class="w-4 h-4 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 9l-7 7-7-7" />
                      </svg>
                    </div>
                  </div>
                  <p class="text-sm text-gray-500">Choose the currency for your shop's pricing</p>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- Action Buttons -->
      <div class="mt-8 bg-white rounded-lg shadow-sm border border-gray-200 overflow-hidden">
        <div class="px-6 py-4 bg-gray-50 flex flex-col sm:flex-row justify-between items-center gap-3">
          <div class="flex items-center space-x-2">
            <svg v-if="lastSaved" class="w-4 h-4 text-green-500" fill="currentColor" viewBox="0 0 20 20">
              <path fill-rule="evenodd" d="M10 18a8 8 0 100-16 8 8 0 000 16zm3.707-9.293a1 1 0 00-1.414-1.414L9 10.586 7.707 9.293a1 1 0 00-1.414 1.414l2 2a1 1 0 001.414 0l4-4z" clip-rule="evenodd" />
            </svg>
            <p v-if="lastSaved" class="text-sm text-gray-600">
              Last saved: {{ lastSaved }}
            </p>
            <p v-else class="text-sm text-gray-500">
              No recent saves
            </p>
          </div>
          <div class="flex gap-3">
            <button
              @click="resetForm"
              :disabled="saving"
              class="px-4 py-2.5 bg-white border border-gray-300 rounded-lg text-gray-700 text-sm font-medium hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-gray-500 focus:ring-offset-2 transition duration-150 ease-in-out disabled:opacity-50 disabled:cursor-not-allowed"
            >
              <svg class="w-4 h-4 mr-2 inline" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 4v5h.582m15.356 2A8.001 8.001 0 004.582 9m0 0H9m11 11v-5h-.581m0 0a8.003 8.003 0 01-15.357-2m15.357 2H15" />
              </svg>
              Reset
            </button>
            <button
              @click="save"
              :disabled="saving || !isFormValid"
              class="inline-flex items-center px-6 py-2.5 bg-green-600 text-white text-sm font-medium rounded-lg hover:bg-green-700 focus:outline-none focus:ring-2 focus:ring-green-500 focus:ring-offset-2 transition duration-150 ease-in-out disabled:opacity-50 disabled:cursor-not-allowed shadow-sm"
            >
              <svg v-if="saving" class="animate-spin -ml-1 mr-2 h-4 w-4 text-white" fill="none" viewBox="0 0 24 24">
                <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
                <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
              </svg>
              <svg v-else class="w-4 h-4 mr-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7" />
              </svg>
              {{ saving ? 'Saving...' : 'Save Changes' }}
            </button>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, reactive, onMounted, computed } from 'vue'
import { useShopStore } from '@/store/shops'
import { shopService } from '@/services/shop'

const shopStore = useShopStore()
const current = ref(null)
const saving = ref(false)
const lastSaved = ref('')

// reactive form
const form = reactive({
  name: '',
  description: '',
  email: '',
  logoUrl: '',
  bannerUrl: '',
  themeColor: '#10B981',
  currency: 'USD',
})

// Form validation errors
const errors = reactive({
  name: '',
  email: '',
  description: '',
  logoUrl: '',
  bannerUrl: '',
})

// Form validation computed
const isFormValid = computed(() => {
  return form.name.trim().length >= 3 && 
         form.name.trim().length <= 50 &&
         form.email.trim() &&
         isValidEmail(form.email) &&
         form.description.length <= 200 &&
         (form.logoUrl === '' || isValidUrl(form.logoUrl)) &&
         (form.bannerUrl === '' || isValidUrl(form.bannerUrl)) &&
         !Object.values(errors).some(error => error)
})

// Validation helper functions
function isValidEmail(email) {
  const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/
  return emailRegex.test(email)
}

function isValidUrl(url) {
  if (!url) return true // empty is valid
  try {
    new URL(url)
    return true
  } catch {
    return false
  }
}

// Load current shop data into form
function loadCurrent() {
  current.value = shopStore.active
  if (!current.value) return
  
  form.name = current.value.name || ''
  form.description = current.value.description || ''
  form.email = current.value.email || ''
  form.logoUrl = current.value.image || ''
  form.bannerUrl = current.value.banner || ''
  form.themeColor = current.value.themeColor || '#10B981'
  form.currency = current.value.currency || 'USD'
  
  // Clear any previous errors
  Object.keys(errors).forEach(key => errors[key] = '')
}

// Validate form fields
function validateForm() {
  // Reset errors
  Object.keys(errors).forEach(key => errors[key] = '')
  
  // Name validation
  if (!form.name.trim()) {
    errors.name = 'Shop name is required'
  } else if (form.name.trim().length < 3) {
    errors.name = 'Shop name must be at least 3 characters'
  } else if (form.name.trim().length > 50) {
    errors.name = 'Shop name must be less than 50 characters'
  }
  
  // Email validation
  if (!form.email.trim()) {
    errors.email = 'Contact email is required'
  } else if (!isValidEmail(form.email)) {
    errors.email = 'Please enter a valid email address'
  }
  
  // Description validation
  if (form.description.length > 200) {
    errors.description = 'Description must be less than 200 characters'
  }
  
  // URL validations
  if (form.logoUrl && !isValidUrl(form.logoUrl)) {
    errors.logoUrl = 'Please enter a valid URL'
  }
  
  if (form.bannerUrl && !isValidUrl(form.bannerUrl)) {
    errors.bannerUrl = 'Please enter a valid URL'
  }
  
  return Object.values(errors).every(error => !error)
}

function resetForm() {
  loadCurrent()
}

async function save() {
  if (!current.value) return
  
  // Validate form before saving
  if (!validateForm()) {
    return
  }
  
  saving.value = true
  try {
    const updateData = {
      name: form.name.trim(),
      description: form.description.trim(),
      email: form.email.trim(),
      image: form.logoUrl.trim() || null,
      banner: form.bannerUrl.trim() || null,
      themeColor: form.themeColor,
      currency: form.currency,
    }
    
    const updated = await shopService.updateShop(current.value.id, updateData)
    
    // Update the store with the new data
    shopStore.setActiveShop({ ...current.value, ...updateData })
    
    // Update last saved timestamp
    lastSaved.value = new Date().toLocaleString()
    
    console.log('Shop updated successfully:', updated)
  } catch (err) {
    console.error('Error updating shop:', err)
    // Handle specific error cases
    if (err.response?.data?.error) {
      // Show backend validation error
      if (err.response.data.error.includes('name')) {
        errors.name = err.response.data.error
      } else if (err.response.data.error.includes('email')) {
        errors.email = err.response.data.error
      } else if (err.response.data.error.includes('description')) {
        errors.description = err.response.data.error
      }
    }
  } finally {
    saving.value = false
  }
}

onMounted(loadCurrent)
</script>

<style scoped>
/* Custom styles for enhanced design */
@media (max-width: 640px) {
  .min-w-0 { min-width: 0 !important; }
}

/* Enhanced hover effects */
.group:hover .group-hover\:scale-105 {
  transform: scale(1.05);
}

/* Smooth transitions */
* {
  transition-property: color, background-color, border-color, text-decoration-color, fill, stroke, opacity, box-shadow, transform, filter, backdrop-filter;
  transition-timing-function: cubic-bezier(0.4, 0, 0.2, 1);
  transition-duration: 150ms;
}

/* Custom scrollbar for better UX */
::-webkit-scrollbar {
  width: 4px;
}

::-webkit-scrollbar-track {
  background: transparent;
}

::-webkit-scrollbar-thumb {
  background: #d1d5db;
  border-radius: 2px;
}

::-webkit-scrollbar-thumb:hover {
  background: #9ca3af;
}

/* Enhanced form focus styles */
input:focus, textarea:focus, select:focus {
  box-shadow: 0 0 0 3px rgba(16, 185, 129, 0.1);
}

/* Loading animation */
@keyframes spin {
  to {
    transform: rotate(360deg);
  }
}

.animate-spin {
  animation: spin 1s linear infinite;
}

/* Card hover effects */
.bg-white:hover {
  box-shadow: 0 10px 15px -3px rgba(0, 0, 0, 0.1), 0 4px 6px -2px rgba(0, 0, 0, 0.05);
}

/* Color picker enhancement */
input[type="color"] {
  appearance: none;
  -webkit-appearance: none;
  border: none;
  cursor: pointer;
}

input[type="color"]::-webkit-color-swatch-wrapper {
  padding: 0;
}

input[type="color"]::-webkit-color-swatch {
  border: none;
  border-radius: 0.75rem;
}
</style>
