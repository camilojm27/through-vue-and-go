<script setup lang="ts">
import type { Source } from '@/types/api'

const props = defineProps({
    email: { type: Object as () => Source, required: true }
})
</script>

<template>
    <li className="px-3 py-2 border rounded-md dark:border-gray-700 dark:bg-gray-800">
        <div className="flex justify-between items-center">
            <h3 className="text-sm font-semibold text-gray-800 dark:text-gray-200">
                {{ email['X-From'] }}
            </h3>
            <span className="text-sm text-gray-600 dark:text-gray-400">{{ email.Date }}</span>
        </div>
        <h2 className="mt-1 text-sm text-gray-800 dark:text-gray-200">
            {{ email.Subject }}
        </h2>
        <p className="mt-1 text-sm text-gray-600 dark:text-gray-400">
            {{ email.Body.substring(0, 100) }}...
        </p>
    </li>
</template>

<style scoped></style>
