# DRPS Project Technical Overview

## Table of Contents
- [Project Structure](#project-structure)
- [Main Technologies](#main-technologies)
- [Backend Architecture](#backend-architecture)
- [Frontend Architecture](#frontend-architecture)
- [Authentication & OAuth](#authentication--oauth)
- [Discount, Cart, and Order Logic](#discount-cart-and-order-logic)
- [Security Principles](#security-principles)
- [Testing & Debugging](#testing--debugging)
- [Key Architectural Decisions](#key-architectural-decisions)

---

## Project Structure

- **backend/**: Go (Golang) REST API server, business logic, MongoDB integration, authentication, discount/order/cart logic.
- **storefront/**: Vue 3 + TypeScript + Vite app for customer-facing shopfront.
- **sellersDashboard/**: Vue 3 + Vite app for sellers to manage shops, products, discounts, and orders.
- **adminPanel/**: Vue 3 + Vite app for platform administrators to manage products, shops, sellers, customers, and analytics.
- **Docs**: Several markdown files detail security, discount, and order system fixes and refactors.

---

## Main Technologies

- **Backend**: Go 1.23+, Gin web framework, MongoDB, JWT, Google OAuth2, CORS, godotenv.
- **Frontend**: Vue 3, Vite, Pinia (state), Vue Router, TailwindCSS, TypeScript (storefront), Axios, Chart.js (sellersDashboard).
- **Authentication**: Google OAuth2 (for sellers and customers), JWT for session management.

---

## Backend Architecture

- **Entry Point**: `main.go` initializes environment, connects to MongoDB, sets up CORS, and registers all API routes.
- **API Structure**:
  - `/auth`: Authentication endpoints for admin, seller, and customer (Google OAuth, JWT refresh/logout, profile).
  - `/admin`: Admin endpoints for managing products, customers, shops, sellers, orders, analytics, and staff.
  - `/seller`: Seller endpoints for shop, product, customer, segment, collection, discount, and order management.
  - `/shops/:shopSlug`: Storefront endpoints for customers (view shop, collections, products, cart, wishlist, orders).
- **Modular Design**: Each domain (admin, sellers, customers, auth) has its own controllers, routes, and services.
- **Database**: MongoDB, with models for Product, Shop, Customer, Seller, Discount, Order, Cart, etc.
- **Discount System**: Only product-level discounts (fixed/percentage), with eligibility (all, specific, segment), usage tracking, and atomic operations for race condition prevention.
- **Order System**: All pricing and discount logic is server-side; minimal data accepted from frontend; stock is reduced after order placement.
- **Cart System**: Cart calculations and discount applications are server-side; frontend only displays server-calculated data.

---

## Frontend Architecture

### Storefront
- **Tech**: Vue 3, TypeScript, Vite, Pinia, TailwindCSS.
- **Purpose**: Customer-facing shopfront for browsing, cart, checkout, and order history.
- **Auth**: Google OAuth for customer login.
- **API**: Communicates with backend `/shops/:shopSlug` and `/auth/customer` endpoints.

### SellersDashboard
- **Tech**: Vue 3, Vite, Pinia, Chart.js, TailwindCSS.
- **Purpose**: Seller dashboard for managing shops, products, discounts, and orders.
- **Auth**: Google OAuth for seller login.
- **API**: Communicates with backend `/seller` and `/auth/seller` endpoints.

### AdminPanel
- **Tech**: Vue 3, Vite, TailwindCSS.
- **Purpose**: Admin dashboard for platform management (products, shops, sellers, customers, analytics).
- **Auth**: Admin login via `/auth/admin` endpoints.
- **API**: Communicates with backend `/admin` endpoints.

---

## Authentication & OAuth

- **Google OAuth2**: Used for both sellers and customers, with separate endpoints and redirect URIs to prevent cross-user-type authentication.
- **JWT**: Used for session management after OAuth login.
- **Environment Variables**: `.env` file in backend for secrets, DB URI, OAuth credentials, and frontend URLs.
- **Security**: Backend validates user type during OAuth callback; separate user models for Seller and Customer.

---

## Discount, Cart, and Order Logic

### Discount System
- **Product-level only**: No order-level, shipping, or buy-x-get-y discounts (future-ready for reintroduction).
- **Eligibility**: All, specific customers, or customer segments.
- **Usage Tracking**: Overall and per-customer usage limits, atomic updates to prevent race conditions.
- **Validation**: All discount logic and validation is server-side; frontend cannot manipulate discount values.

### Cart System
- **Server-side calculations**: All pricing, discount, and total calculations are done on the backend.
- **Frontend**: Only displays server-calculated data; no client-side price calculations.
- **Security**: Minimal data sent from frontend (product_id, variant_id, quantity).

### Order System
- **Order Placement**: Backend accepts only minimal data, validates all input, calculates all prices and discounts, and reduces stock after order creation.
- **Discount Application**: Best eligible discount is selected and applied server-side per product/variant.
- **Security Flags**: API responses include flags indicating server-calculated totals and security notes.

---

## Security Principles

- **Never trust frontend data**: All critical calculations (pricing, discounts, totals) are server-side.
- **Minimal data acceptance**: Backend only accepts essential data from frontend.
- **Comprehensive validation**: All input is validated, including product/variant IDs, quantities, and eligibility.
- **Atomic operations**: Prevent race conditions in discount usage and stock reduction.
- **Clear security messaging**: Frontend displays security notes to users.
- **Robust error handling**: Graceful handling of all edge cases and errors.

---

## Testing & Debugging

- **Comprehensive test suites**: For discount logic, order placement, and security (see test files in backend).
- **Debug endpoints**: For product/variant inspection and price updates.
- **Verification checklists**: Provided in docs for validating fixes and security.

---

## Key Architectural Decisions

- **Separation of concerns**: Distinct backend and frontend apps for admin, sellers, and customers.
- **Strict security**: All sensitive logic is server-side; frontend is a thin client.
- **Extensible discount system**: Product-level only for now, but code is ready for future reintroduction of order/shipping/collection discounts.
- **Modular Go codebase**: Each domain (admin, sellers, customers, auth) is isolated for maintainability.
- **Modern frontend stack**: Vue 3 + Vite + TailwindCSS for all UIs, with TypeScript in the storefront for type safety.

---

# End of Technical Overview 