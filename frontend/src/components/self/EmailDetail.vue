<script setup lang="ts">
import { type Source } from '@/types/api'

const props = defineProps({
    email: { type: Object as () => Source }
})
console.log(props.email)
</script>

<template>
    <div
        id="email-detail"
        className="w-1/2 py-3 border-l dark:border-gray-700 sticky top-0 right-0 h-screen overflow-y-auto"
    >
        <div
            className="px-6 py-4 border rounded-md dark:border-gray-700 dark:bg-gray-800 overflow-visible"
        >
            <div className="flex justify-between items-center">
                <h3 className="text-lg font-semibold text-gray-800 dark:text-gray-200">
                    {{ email?.Subject }}
                </h3>
                <span className="text-sm text-gray-600 dark:text-gray-400">{{ email?.Date }}</span>
            </div>
            <h2 className="mt-2 text-sm text-gray-800 dark:text-gray-200">
                {{ email?.['X-From'] }}
            </h2>
            <p className="mt-2 text-sm text-gray-600 dark:text-gray-400">Dear Team,</p>
            <p className="mt-2 text-sm text-gray-600 dark:text-gray-400">
                {{ email?.Body }}
            </p>
            <p className="mt-2 text-sm text-gray-600 dark:text-gray-400">John Doe</p>
        </div>
    </div>
</template>

<style scoped></style>
