<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { instance } from '@/lib/axios'
import { Input } from '@/components/ui/input'
import type { RootObject, Source } from '@/types/api'
import EmailDetail from '@/components/self/EmailDetail.vue'
import EmailPreview from '@/components/self/EmailPreview.vue'

const data = ref<RootObject | null>(null)
const loading = ref(true)
const error = ref('')
let mail = ref<Source | null>(null)

async function fetchData() {
    loading.value = true
    error.value = ''
    try {
        const response = await instance.get('/')
        if (response.status !== 200) {
            throw new Error(`HTTP error! status: ${response.status}`)
        }
        data.value = response.data
    } catch (e) {
        error.value = e.toString()
    } finally {
        loading.value = false
    }
}

onMounted(fetchData)

</script>

<template>
    <div class="flex h-screen bg-gray-100 dark:bg-gray-900">
        <aside class="w-64 bg-white dark:bg-gray-800 border-r dark:border-gray-700">
            <div class="p-4">
                <h2 class="text-lg font-semibold text-gray-800 dark:text-gray-200">Mailbox</h2>
                <nav class="mt-8">
                    <div class="mt-2 -mx-3">
                        <Link
                            class="flex justify-between items-center px-3 py-2 rounded-lg font-medium text-gray-700 dark:text-gray-200 bg-gray-200 dark:bg-gray-700"
                            href="#"
                        >
                            <span>{{$route.params.mail}}</span>
                            <span class="text-sm font-semibold text-gray-700 dark:text-gray-200"
                                >120</span
                            >
                        </Link>
                        <a
                            class="mt-1 flex justify-between items-center px-3 py-2 rounded-lg font-medium text-gray-700 dark:text-gray-200 hover:bg-gray-200 dark:hover:bg-gray-700"
                            href="#"
                        >
                            <span>Sent</span>
                        </a>
                        <a
                            class="mt-1 flex justify-between items-center px-3 py-2 rounded-lg font-medium text-gray-700 dark:text-gray-200 hover:bg-gray-200 dark:hover:bg-gray-700"
                            href="#"
                        >
                            <span>Drafts</span>
                        </a>
                        <a
                            class="mt-1 flex justify-between items-center px-3 py-2 rounded-lg font-medium text-gray-700 dark:text-gray-200 hover:bg-gray-200 dark:hover:bg-gray-700"
                            href="#"
                        >
                            <span>Starred</span>
                        </a>
                        <a
                            class="mt-1 flex justify-between items-center px-3 py-2 rounded-lg font-medium text-gray-700 dark:text-gray-200 hover:bg-gray-200 dark:hover:bg-gray-700"
                            href="#"
                        >
                            <span>Trash</span>
                        </a>
                    </div>
                </nav>
            </div>
        </aside>
        <main class="flex-1 overflow-y-auto">
            <div class="px-4 py-3 border-b dark:border-gray-700 bg-white dark:bg-gray-800">
                <div class="flex items-center">
                    <Input
                        class="w-full px-3 py-2 border rounded-md dark:border-gray-700 dark:bg-gray-700 dark:text-gray-200"
                        placeholder="Search emails"
                        type="search"
                    />
                </div>
            </div>
            <div class="flex relative">
                <ul v-if="data !== null" class="w-1/2 px-4 py-3 space-y-2">
                    <EmailPreview v-bind:emails="data?.hits.hits" />
                </ul>
                <EmailDetail :email="mail" />
            </div>
        </main>
    </div>
</template>
<!--                        v-on:click="selectEmail(email._id)"-->
