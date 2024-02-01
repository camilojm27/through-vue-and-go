<script setup lang="ts">
import type { Hit } from '@/types/api'
import { useRouter } from 'vue-router'
const router = useRouter()

// eslint-disable-next-line @typescript-eslint/no-unused-vars
const props = defineProps({
    emails: { type: Object as () => Hit, required: true }
})

const redirectToEmail = (emailId: string) => {
    // Perform any necessary logic before redirecting
    router.push(`/${emailId}`)
}
</script>

<template>
    <li
        v-for="email in emails"
        v-bind:key="email._id"
        className="px-3 py-2 border rounded-md dark:border-gray-700 dark:bg-gray-800"
        @click="redirectToEmail(email._id)"
    >
        <div className="flex justify-between items-center">
            <h3 className="text-sm font-semibold text-gray-800 dark:text-gray-200">
                {{ email._source['X-From'] }}
            </h3>
            <span className="text-sm text-gray-600 dark:text-gray-400">{{
                email._source.Date
            }}</span>
        </div>
        <h2 className="mt-1 text-sm text-gray-800 dark:text-gray-200">
            {{ email._source.Subject }}
        </h2>
        <p className="mt-1 text-sm text-gray-600 dark:text-gray-400">
            {{ email._source.Body.substring(0, 100) }}...
        </p>
    </li>
</template>
