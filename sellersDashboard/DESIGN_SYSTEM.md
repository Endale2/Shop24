# Sellers Dashboard Design System

## Overview
This document outlines the standardized design system implemented to ensure consistency across all pages and components in the sellers dashboard.

## Design Inconsistencies Fixed

### Before (Issues Found)
1. **Inconsistent Color Schemes**: Mixed green/amber vs blue themes
2. **Varied Button Styles**: Different padding, colors, and hover effects
3. **Inconsistent Typography**: Mixed font sizes, weights, and line heights
4. **Card Design Variations**: Different shadows, borders, and spacing
5. **Form Element Inconsistencies**: Various input styles and validation states
6. **Status Badge Variations**: Different colors and styling approaches
7. **Spacing Inconsistencies**: Mixed padding and margin values
8. **Component Styling Conflicts**: Some pages used custom CSS, others pure Tailwind

### After (Standardized)
1. **Unified Color Palette**: Consistent green primary, amber secondary theme
2. **Standardized Components**: Reusable button, card, and form components
3. **Typography Scale**: Consistent font sizes, weights, and line heights
4. **Design Tokens**: CSS variables for colors, spacing, and typography
5. **Component Library**: Reusable UI components with consistent APIs

## Design System Structure

### 1. CSS Variables (`src/styles/design-system.css`)
```css
:root {
  /* Primary Brand Colors - Green Theme */
  --color-primary: #10B981;
  --color-primary-50: #ECFDF5;
  /* ... full color scale */
  
  /* Typography */
  --font-family-primary: 'Inter', sans-serif;
  --font-size-xs: 0.75rem;
  /* ... full type scale */
  
  /* Spacing Scale */
  --spacing-1: 0.25rem;
  /* ... full spacing scale */
}
```

### 2. Component Classes
- `.btn` - Base button styles with variants
- `.card` - Consistent card component
- `.form-input` - Standardized form inputs
- `.badge` - Status badges with semantic colors
- `.stat-card` - Dashboard statistics cards

### 3. Layout Components
- `.page-container` - Consistent page layout
- `.stats-grid` - Responsive statistics grid
- `.grid-responsive` - Responsive grid system

## Component Library

### BaseButton (`src/components/ui/BaseButton.vue`)
```vue
<BaseButton variant="primary" size="lg" @click="handleClick">
  Primary Button
</BaseButton>

<BaseButton variant="secondary" :loading="isLoading">
  Loading Button
</BaseButton>

<BaseButton variant="ghost" icon="PlusIcon" icon-only />
```

**Props:**
- `variant`: 'primary' | 'secondary' | 'ghost' | 'danger'
- `size`: 'sm' | 'base' | 'lg'
- `loading`: boolean
- `disabled`: boolean
- `icon`: Vue component
- `fullWidth`: boolean

### BaseCard (`src/components/ui/BaseCard.vue`)
```vue
<BaseCard title="Card Title" subtitle="Subtitle" hoverable>
  Card content
</BaseCard>

<BaseCard clickable @click="handleClick">
  <template #header>Custom Header</template>
  Card body
  <template #footer>Footer content</template>
</BaseCard>
```

**Props:**
- `title`: string
- `subtitle`: string
- `hoverable`: boolean
- `clickable`: boolean
- `padding`: 'none' | 'sm' | 'normal' | 'lg'

### StatusBadge (`src/components/ui/StatusBadge.vue`)
```vue
<StatusBadge variant="success" label="Active" />
<StatusBadge variant="warning" label="Pending" dot />
<StatusBadge variant="error" icon="ExclamationIcon">Error</StatusBadge>
```

**Props:**
- `variant`: 'success' | 'warning' | 'error' | 'info' | 'neutral'
- `size`: 'sm' | 'base' | 'lg'
- `label`: string
- `icon`: Vue component
- `dot`: boolean

## Color Palette

### Primary Colors (Green Theme)
- `--color-primary-50`: #ECFDF5 (lightest)
- `--color-primary-500`: #10B981 (primary)
- `--color-primary-900`: #064E3B (darkest)

### Secondary Colors (Amber Accent)
- `--color-secondary-50`: #FFFBEB
- `--color-secondary-500`: #F59E0B
- `--color-secondary-900`: #78350F

### Semantic Colors
- Success: Primary green
- Warning: Secondary amber
- Error: #EF4444
- Info: #3B82F6

## Typography Scale

### Font Sizes
- `--font-size-xs`: 0.75rem (12px)
- `--font-size-sm`: 0.875rem (14px)
- `--font-size-base`: 1rem (16px)
- `--font-size-lg`: 1.125rem (18px)
- `--font-size-xl`: 1.25rem (20px)
- `--font-size-2xl`: 1.5rem (24px)
- `--font-size-3xl`: 1.875rem (30px)
- `--font-size-4xl`: 2.25rem (36px)

### Font Weights
- Normal: 400
- Medium: 500
- Semibold: 600
- Bold: 700

## Spacing Scale

Based on 4px base unit:
- `--spacing-1`: 0.25rem (4px)
- `--spacing-2`: 0.5rem (8px)
- `--spacing-3`: 0.75rem (12px)
- `--spacing-4`: 1rem (16px)
- `--spacing-5`: 1.25rem (20px)
- `--spacing-6`: 1.5rem (24px)
- `--spacing-8`: 2rem (32px)
- `--spacing-10`: 2.5rem (40px)
- `--spacing-12`: 3rem (48px)
- `--spacing-16`: 4rem (64px)

## Usage Guidelines

### 1. Import Design System
```css
/* In your component or main CSS */
@import "./styles/design-system.css";
```

### 2. Use CSS Variables
```css
.custom-component {
  color: var(--color-text-primary);
  padding: var(--spacing-4);
  border-radius: var(--radius-base);
}
```

### 3. Use Component Classes
```html
<button class="btn btn-primary">Primary Button</button>
<div class="card">Card content</div>
<input class="form-input" type="text" />
<span class="badge badge-success">Success</span>
```

### 4. Import UI Components
```vue
<script setup>
import { BaseButton, BaseCard, StatusBadge } from '@/components/ui'
</script>
```

## Migration Guide

### Updating Existing Components

1. **Replace custom button styles:**
```vue
<!-- Before -->
<button class="px-4 py-2 bg-blue-500 text-white rounded hover:bg-blue-600">
  Button
</button>

<!-- After -->
<BaseButton variant="primary">Button</BaseButton>
```

2. **Replace custom cards:**
```vue
<!-- Before -->
<div class="bg-white rounded-lg shadow-sm border p-6">
  Content
</div>

<!-- After -->
<BaseCard>Content</BaseCard>
```

3. **Replace custom form inputs:**
```vue
<!-- Before -->
<input class="w-full px-3 py-2 border rounded focus:ring-2 focus:ring-blue-500" />

<!-- After -->
<input class="form-input" />
```

## Best Practices

1. **Always use design tokens** instead of hardcoded values
2. **Prefer component classes** over inline styles
3. **Use semantic color names** (primary, success, error) over specific colors
4. **Maintain consistent spacing** using the spacing scale
5. **Test components** in different states (hover, focus, disabled)
6. **Follow accessibility guidelines** for color contrast and focus states

## Browser Support

The design system supports:
- Chrome 90+
- Firefox 88+
- Safari 14+
- Edge 90+

CSS custom properties are used extensively, ensuring modern browser compatibility.
