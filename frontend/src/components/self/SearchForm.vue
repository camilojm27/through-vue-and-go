<script setup lang="ts">

import { Input } from '@/components/ui/input'
const emit = defineEmits(['search'])

const search = (e: Event) => {
  emit('search', (e.target as HTMLInputElement).value)
}
</script>

<template>
  <div class="px-4 py-3 border-b dark:border-gray-700 bg-white dark:bg-gray-800">
    <div class="flex items-center">
      <Input
        class="w-full px-3 py-2 border rounded-md dark:border-gray-700 dark:bg-gray-700 dark:text-gray-200"
        placeholder="Search emails"
        type="search"
        @input="search"
      />
    </div>
  </div>
</template>

<style scoped>

</style>