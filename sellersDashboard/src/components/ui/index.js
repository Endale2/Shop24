// UI Component Library Exports
export { default as BaseButton } from './BaseButton.vue'
export { default as BaseCard } from './BaseCard.vue'
export { default as StatusBadge } from './StatusBadge.vue'
export { default as LoadingSpinner } from './LoadingSpinner.vue'

// Usage Examples:
/*
// BaseButton Examples:
<BaseButton variant="primary" size="lg" @click="handleClick">
  Primary Button
</BaseButton>

<BaseButton variant="secondary" :loading="isLoading">
  Loading Button
</BaseButton>

<BaseButton variant="ghost" icon="PlusIcon" icon-only />

// BaseCard Examples:
<BaseCard title="Card Title" subtitle="Card subtitle" hoverable>
  Card content goes here
</BaseCard>

<BaseCard clickable @click="handleCardClick">
  <template #header>
    <h3>Custom Header</h3>
  </template>
  Card body content
  <template #footer>
    <BaseButton variant="primary">Action</BaseButton>
  </template>
</BaseCard>

// StatusBadge Examples:
<StatusBadge variant="success" label="Active" />
<StatusBadge variant="warning" label="Pending" dot />
<StatusBadge variant="error" icon="ExclamationIcon">Error</StatusBadge>

// LoadingSpinner Examples:
<LoadingSpinner size="lg" centered />
<LoadingSpinner size="sm" color="primary" />
*/
