// Discount utility functions for the storefront

export interface DiscountInfo {
  type: 'fixed' | 'percentage';
  value: number;
  name?: string;
}

/**
 * Format discount value for display
 */
export function formatDiscountValue(discount: DiscountInfo): string {
  if (discount.type === 'percentage') {
    return `${discount.value}% off`;
  } else {
    return `$${discount.value.toFixed(2)} off`;
  }
}

/**
 * Get discount description for display
 */
export function getDiscountDescription(discount: DiscountInfo): string {
  if (discount.type === 'percentage') {
    return `${discount.value}% discount applied`;
  } else {
    return `$${discount.value.toFixed(2)} discount applied`;
  }
}

/**
 * Get discount type description
 */
export function getDiscountTypeDescription(discount: DiscountInfo): string {
  if (discount.type === 'percentage') {
    return `${discount.value}% off`;
  } else {
    return `$${discount.value.toFixed(2)} off`;
  }
}

/**
 * Calculate discount amount for a given price and quantity
 */
export function calculateDiscountAmount(discount: DiscountInfo, unitPrice: number, quantity: number): number {
  if (discount.type === 'percentage') {
    const totalPrice = unitPrice * quantity;
    return totalPrice * (discount.value / 100);
  } else {
    // For fixed amount discounts, apply per unit
    const discountPerUnit = Math.min(discount.value, unitPrice);
    return discountPerUnit * quantity;
  }
}

/**
 * Calculate final price after discount
 */
export function calculateFinalPrice(discount: DiscountInfo, unitPrice: number, quantity: number): number {
  const discountAmount = calculateDiscountAmount(discount, unitPrice, quantity);
  const totalPrice = unitPrice * quantity;
  return Math.max(0, totalPrice - discountAmount);
}

/**
 * Check if a discount is valid
 */
export function isValidDiscount(discount: DiscountInfo): boolean {
  if (discount.type === 'percentage') {
    return discount.value > 0 && discount.value <= 100;
  } else {
    return discount.value > 0;
  }
} 