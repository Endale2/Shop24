# 🛒 Shop24

**Shop24** is a multi-tenant e-commerce platform (similar to Shopify) that allows sellers to create and manage their own online stores.  
It’s designed for scalability and includes advanced features like product variants, collections, order and discount management, analytics, and theme customization.  

🌐 Live Demo: [http://shop24.sbs](http://shop24.sbs)  
Stores are accessible via `yourshopname.shop24.sbs`.

---

## 🔹 Key Features
- Multi-tenant architecture with subdomain-based stores
- Product listings with variants and collections
- Order processing and management
- Analytics dashboard with sales metrics
- Discount and coupon management
- Storefront theme and layout customization
- Hosted on AWS with Nginx and Docker for scalability

---

## 🔹 Tech Stack
| Layer        | Technologies |
|-------------|--------------|
| Backend     | Go (Gin Framework) |
| Frontend    | Vue.js, Nuxt.js |
| Database    | MongoDB |
| Deployment  | AWS (EC2, S3, Route53), Docker, Nginx |
| Auth        | JWT Authentication |

---

## 🔹 Architecture Notes
- **Backend:** Modular design with separate packages for `auth`, `admin`, `seller`, `customers`, and `shared` modules.  
- **Frontend:** Vue/Nuxt frontend for storefront and admin dashboards.  
- **Domain Structure:**  
  - Main site: `shop24.sbs`  
  - Individual stores: `storename.shop24.sbs`  

---

## 🔹 Current Status
- Core backend features implemented  
- Seller dashboards and storefront functional  
- Actively adding mobile support with Flutter  
- Ongoing improvements to analytics and theme customization

---

## 🔹 Contact
- **Developer:** Endale Shimelis  
- 📧 Email: endale406@gmail.com    
- 💬 Telegram: [@codejkr](https://t.me/codejkr)
