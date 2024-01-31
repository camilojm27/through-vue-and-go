<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { instance } from '@/lib/axios'
import { Input } from '@/components/ui/input'
import type { RootObject, Source } from '@/types/api'
import EmailDetail from '@/components/self/EmailDetail.vue'
import EmailPreview from '@/components/self/EmailPreview.vue'

const data = ref<RootObject | null>(null)
const loading = ref(true)
const error = ref()
let mail = ref(<Source | null>null)

async function fetchData() {
    loading.value = true
    error.value = null
    try {
        const response = await instance.get('/')
        if (response.status !== 200) {
            throw new Error(`HTTP error! status: ${response.status}`)
        }
        data.value = response.data
    } catch (e) {
        error.value = e
    } finally {
        loading.value = false
    }
}

function selectEmail(id: string) {
    mail.value = data.value?.hits.hits.find((email) => email._id === id)?._source
    console.log(mail.value)
}

onMounted(fetchData)
</script>

<template>
    <div className="flex h-screen bg-gray-100 dark:bg-gray-900">
        <aside className="w-64 bg-white dark:bg-gray-800 border-r dark:border-gray-700">
            <div className="p-4">
                <h2 className="text-lg font-semibold text-gray-800 dark:text-gray-200">Mailbox</h2>
                <nav className="mt-8">
                    <div className="mt-2 -mx-3">
                        <Link
                            className="flex justify-between items-center px-3 py-2 rounded-lg font-medium text-gray-700 dark:text-gray-200 bg-gray-200 dark:bg-gray-700"
                            href="#"
                        >
                            <span>Inbox</span>
                            <span className="text-sm font-semibold text-gray-700 dark:text-gray-200"
                                >120</span
                            >
                        </Link>
                        <Link
                            className="mt-1 flex justify-between items-center px-3 py-2 rounded-lg font-medium text-gray-700 dark:text-gray-200 hover:bg-gray-200 dark:hover:bg-gray-700"
                            href="#"
                        >
                            <span>Sent</span>
                        </Link>
                        <Link
                            className="mt-1 flex justify-between items-center px-3 py-2 rounded-lg font-medium text-gray-700 dark:text-gray-200 hover:bg-gray-200 dark:hover:bg-gray-700"
                            href="#"
                        >
                            <span>Drafts</span>
                        </Link>
                        <Link
                            className="mt-1 flex justify-between items-center px-3 py-2 rounded-lg font-medium text-gray-700 dark:text-gray-200 hover:bg-gray-200 dark:hover:bg-gray-700"
                            href="#"
                        >
                            <span>Starred</span>
                        </Link>
                        <Link
                            className="mt-1 flex justify-between items-center px-3 py-2 rounded-lg font-medium text-gray-700 dark:text-gray-200 hover:bg-gray-200 dark:hover:bg-gray-700"
                            href="#"
                        >
                            <span>Trash</span>
                        </Link>
                    </div>
                </nav>
            </div>
        </aside>
        <main className="flex-1 overflow-y-auto">
            <div className="px-4 py-3 border-b dark:border-gray-700 bg-white dark:bg-gray-800">
                <div className="flex items-center">
                    <Input
                        className="w-full px-3 py-2 border rounded-md dark:border-gray-700 dark:bg-gray-700 dark:text-gray-200"
                        placeholder="Search emails"
                        type="search"
                    />
                </div>
            </div>
            <div className="flex relative">
                <ul className="w-1/2 px-4 py-3 space-y-2">
                    <EmailPreview
                        v-for="email in data?.hits.hits"
                        v-bind:key="email._id"
                        v-on:click="selectEmail(email._id)"
                        v-bind:email="email._source"
                    />
                </ul>
                <EmailDetail v-bind:email="mail?.value" />
            </div>
        </main>
    </div>
</template>
