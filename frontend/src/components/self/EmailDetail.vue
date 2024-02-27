<script setup lang="ts">
import { type Source } from '@/types/api'
import { Badge } from '@/components/ui/badge'

const props = defineProps({
    email: { type: Object as () => Source }
})

</script>

<template>
    <div
        id="email-detail"
        class="w-1/2 py-3 border-l dark:border-gray-700 sticky top-0 right-0 h-screen overflow-y-auto"
        v-if="email !== null"
    >
        <div
            class="px-6 py-4 border rounded-md dark:border-gray-700 dark:bg-gray-800 overflow-visible"
        >
          <span>
            Folders// <Badge v-for="folder in email?.['X-Folder'].split('\\').slice(1,9)" class="m-1 p-0.5" :key="folder" variant="outline">
            {{folder}}
                </Badge>
          </span>

            <div class="flex justify-between items-center">
                <h3 class="text-lg font-semibold text-gray-800 dark:text-gray-200">
                    {{ email?.Subject }}
                </h3>
                <span class="text-sm text-gray-600 dark:text-gray-400">{{ email?.Date }}</span>
            </div>
            <h2 class="mt-4 text-sm text-gray-800 dark:text-gray-200">
              From  {{ email?.['From'] }}
            </h2>
          <h2 class="mt-4 text-sm text-gray-800 dark:text-gray-200">
            To  {{ email?.['To'] }}
          </h2>
          <pre class="mt-2 text-sm text-gray-600 dark:text-gray-400">{{ email?.['CC'] }}</pre>
          <pre class="mt-2 text-sm text-gray-600 dark:text-gray-400">{{ email?.['X-bcc'] }}</pre>
          <pre class="mt-2 text-sm text-gray-600 dark:text-gray-400">{{ email?.Body }}</pre>

          Message-Id
          <Badge>
            {{email?.['Message-Id']}}
          </Badge>
        </div>
    </div>
</template>

<style scoped></style>
