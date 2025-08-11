# Theme Management System for Sellers

This document outlines the comprehensive theme management system built for the DRPS platform, allowing sellers to customize their shop appearance and manage themes.

## 🏗️ Architecture Overview

The theme system follows a clean architecture pattern with:
- **Models**: Theme data structures based on `backend/shared/models/theme.go`
- **Repositories**: Data access layer for theme operations
- **Services**: Business logic layer with validation and authorization
- **Controllers**: HTTP API endpoints for frontend integration
- **Routes**: RESTful API routing for theme management

## 📁 File Structure

```
backend/sellers/
├── controllers/
│   ├── theme_controller.go          # Main theme API endpoints
│   └── theme_seed_controller.go     # Default theme seeding
├── services/
│   └── theme_service.go             # Business logic & validation
├── repositories/
│   └── theme_repository.go          # Database operations
└── routes/
    └── seller_routes.go             # Updated with theme routes

backend/shared/models/
├── theme.go                         # Theme data models
└── shop.go                          # Extended with theme fields

sellersDashboard/src/
├── services/
│   └── theme.js                     # Frontend API service
└── pages/customization/
    ├── CustomizationPage.vue        # Main customization hub
    ├── ThemePage.vue                # Color & theme settings
    ├── LayoutPage.vue               # Layout configuration
    ├── TypographyPage.vue           # Font & typography
    ├── CustomCSSPage.vue            # Custom CSS editor
    ├── MobilePage.vue               # Mobile optimization
    └── SEOPage.vue                  # SEO & meta settings
```

## 🎨 Theme System Components

### 1. Core Models

#### `Theme`
- Base theme definitions (marketplace & custom themes)
- Configuration for colors, fonts, layouts, sections
- Versioning and metadata support

#### `ShopTheme`
- Links shops to themes with customizations
- Override configurations for shop-specific settings
- Publication status management

#### `ThemePreset`
- Pre-configured theme variations
- Quick-apply theme settings
- Default and custom presets

### 2. API Endpoints

#### Theme Management
```
POST   /seller/themes                    # Create custom theme
GET    /seller/themes/marketplace        # List public themes
GET    /seller/themes/my-themes          # List seller's themes
GET    /seller/themes/search             # Search themes
GET    /seller/themes/:id                # Get specific theme
PATCH  /seller/themes/:id                # Update theme
DELETE /seller/themes/:id                # Delete theme
POST   /seller/themes/seed               # Seed default themes
```

#### Shop Theme Configuration
```
POST   /seller/shops/:id/themes                               # Create shop theme config
GET    /seller/shops/:id/themes                               # List shop themes
GET    /seller/shops/:id/themes/active                        # Get active theme
PATCH  /seller/shops/:id/themes/:themeId                      # Update shop theme
POST   /seller/shops/:id/themes/:themeId/publish              # Publish theme
DELETE /seller/shops/:id/themes/:themeId                      # Delete shop theme
POST   /seller/shops/:id/themes/:themeId/presets/:id/apply    # Apply preset
```

#### Customization API
```
GET    /seller/shops/:id/customization    # Get current customization
PATCH  /seller/shops/:id/customization    # Update customization
```

### 3. Frontend Integration

#### Theme Service (`theme.js`)
- Complete API wrapper for all theme operations
- Validation helpers for theme configuration
- Default configuration management
- Error handling and loading states

#### Customization Pages
- **CustomizationPage**: Hub with overview cards for all customization options
- **ThemePage**: Color picker interface with live preview and presets
- **LayoutPage**: Drag & drop layout builder with responsive preview
- **TypographyPage**: Font selection and typography settings
- **CustomCSSPage**: Advanced CSS editor with validation
- **MobilePage**: Mobile-specific optimization settings
- **SEOPage**: SEO and social media meta configuration

## 🔧 Key Features

### 1. Theme Customization
- **Color Management**: Primary, secondary, background, text colors
- **Typography**: Font families, sizes, weights, spacing
- **Layout**: Container width, header styles, grid configurations
- **Responsive Design**: Mobile, tablet, desktop previews

### 2. Theme Marketplace
- **Public Themes**: Curated marketplace themes
- **Custom Themes**: Seller-created themes
- **Theme Search**: Search by name, tags, description
- **Theme Presets**: Quick-apply color schemes

### 3. Advanced Features
- **Live Preview**: Real-time customization preview
- **Theme Export**: Download theme configurations
- **Custom CSS**: Advanced styling capabilities
- **Mobile Optimization**: Mobile-specific settings
- **SEO Integration**: Meta tags and social media optimization

### 4. Developer Features
- **Theme Seeding**: Default marketplace themes
- **Validation**: Comprehensive input validation
- **Error Handling**: Graceful error management
- **Authorization**: Shop ownership validation
- **Transactions**: Atomic theme publishing

## 🚀 Getting Started

### 1. Backend Setup

1. **Database Collections**: The system uses MongoDB collections:
   - `themes`: Base theme definitions
   - `shop_themes`: Shop-specific theme configurations
   - `theme_presets`: Theme preset configurations

2. **Seed Default Themes**:
   ```bash
   curl -X POST http://localhost:8080/seller/themes/seed \
     -H "Authorization: Bearer YOUR_TOKEN"
   ```

3. **Environment Variables**: Ensure MongoDB connection is configured

### 2. Frontend Setup

1. **Install Dependencies**: All Vue.js dependencies should be installed

2. **API Configuration**: Update `src/services/api.js` with correct backend URL

3. **Navigation**: The customization section is added to the seller dashboard sidebar

### 3. Usage Flow

1. **Access Customization**: Navigate to `/dashboard/customization`
2. **Choose Category**: Select theme, layout, typography, etc.
3. **Customize Settings**: Use the intuitive interface to modify settings
4. **Live Preview**: See changes in real-time
5. **Save Changes**: Persist customizations to the backend
6. **Publish**: Activate the theme for your shop

## 🎯 Default Theme Configurations

The system includes 5 default marketplace themes:

1. **Classic Commerce**: Professional blue theme
2. **Modern Minimalist**: Clean green theme
3. **Bold & Vibrant**: Orange accent theme
4. **Dark Professional**: Dark mode theme
5. **Elegant Fashion**: Pink fashion theme

Each theme includes:
- Color palette (5 colors)
- Font configuration (heading & body)
- Layout settings (header, container, grid)
- Responsive breakpoints

## 🔒 Security & Validation

### Authorization
- Shop ownership validation for all operations
- Theme creator permissions for custom themes
- JWT-based authentication

### Validation
- Color format validation (hex codes)
- Font name validation
- Theme configuration structure validation
- Input sanitization and length limits

### Error Handling
- Graceful API error responses
- Frontend error state management
- Validation error display
- Loading state indicators

## 🎨 Customization Examples

### Color Configuration
```javascript
const colors = {
  primary: '#10B981',      // Main brand color
  secondary: '#F59E0B',    // Accent color
  background: '#FFFFFF',   // Page background
  heading: '#1F2937',      // Heading text
  bodyText: '#6B7280'      // Body text
}
```

### Typography Configuration
```javascript
const fonts = {
  heading: 'Inter',        // Heading font family
  body: 'Inter'            // Body font family
}
```

### Layout Configuration
```javascript
const layout = {
  headerStyle: 'classic',      // Header style
  containerWidth: 'boxed',     // Container width
  sidebarPosition: 'none',     // Sidebar position
  gridColumns: '3'             // Product grid columns
}
```

## 🔄 Future Enhancements

### Planned Features
1. **Theme Marketplace**: Public theme marketplace with ratings
2. **Advanced Sections**: More customizable page sections
3. **Animation Settings**: Transition and animation controls
4. **A/B Testing**: Theme performance testing
5. **Theme Analytics**: Usage and performance metrics
6. **Import/Export**: Theme file import/export functionality
7. **Collaboration**: Multi-user theme editing
8. **Version History**: Theme change tracking

### Technical Improvements
1. **Caching**: Redis caching for theme data
2. **CDN Integration**: Asset optimization and delivery
3. **Real-time Updates**: WebSocket-based live updates
4. **Advanced Validation**: More comprehensive validation rules
5. **Performance Monitoring**: Theme performance tracking
6. **Backup/Restore**: Theme configuration backups

## 📝 API Documentation

For detailed API documentation, refer to the individual controller files:
- `theme_controller.go`: Main theme management endpoints
- `theme_seed_controller.go`: Development and seeding endpoints

## 🤝 Contributing

When extending the theme system:

1. **Follow Patterns**: Use established repository → service → controller pattern
2. **Add Validation**: Include proper input validation
3. **Error Handling**: Implement comprehensive error handling
4. **Documentation**: Update this README with new features
5. **Testing**: Add unit tests for new functionality

## 📞 Support

For issues or questions regarding the theme system:
- Check the error logs in both frontend and backend
- Verify shop ownership and authentication
- Ensure all required fields are provided
- Review validation error messages

The theme management system provides a solid foundation for shop customization while maintaining clean architecture and extensibility for future enhancements.
