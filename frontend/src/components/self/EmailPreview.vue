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
        class="px-3 py-2 border rounded-md dark:border-gray-700 dark:bg-gray-800 cursor-pointer hover:bg-black hover:bg-opacity-20"
        @click="redirectToEmail(email._id)"
    >
        <div class="flex justify-between items-center">
            <h3 class="text-sm text-gray-800 dark:text-gray-200">
               From: <span class="font-semibold"> {{ email._source['From'] }} </span>
            </h3>
            <span class="text-sm text-gray-600 dark:text-gray-400">{{
                email._source.Date
            }}</span>
        </div>
      <h3 class="text-sm text-gray-800 dark:text-gray-200">
        To: <span class="font-semibold"> {{ email._source['To'] }} </span>
      </h3>
        <h2 class="mt-1 text-sm text-gray-800 dark:text-gray-200">
            {{ email._source.Subject }}
        </h2>
        <p class="mt-1 text-sm text-gray-600 dark:text-gray-400">
            {{ email._source.Body.substring(0, 150) }}...
        </p>
    </li>
</template>
