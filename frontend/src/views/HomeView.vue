<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { instance } from '@/lib/axios'
import { Input } from '@/components/ui/input'
import type { RootObject } from '@/types/api'

const data = ref<RootObject | null>(null)
const loading = ref(true)
const error = ref()

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
            <div className="flex">
                <ul className="w-1/2 px-4 py-3 space-y-2">
                    <li
                        className="px-3 py-2 border rounded-md dark:border-gray-700 dark:bg-gray-800"
                        v-for="email in data?.hits.hits"
                        v-bind:key="email._id"
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
                </ul>
                <div className="w-1/2 px-4 py-3 border-l dark:border-gray-700 ">
                    <div
                        className="px-6 py-4 border rounded-md dark:border-gray-700 dark:bg-gray-800"
                    >
                        <div className="flex justify-between items-center">
                            <h3 className="text-lg font-semibold text-gray-800 dark:text-gray-200">
                                Meeting Reminder
                            </h3>
                            <span className="text-sm text-gray-600 dark:text-gray-400"
                                >Jan 3, 2024</span
                            >
                        </div>
                        <h2 className="mt-2 text-sm text-gray-800 dark:text-gray-200">
                            From: John Doe
                        </h2>
                        <p className="mt-2 text-sm text-gray-600 dark:text-gray-400">Dear Team,</p>
                        <p className="mt-2 text-sm text-gray-600 dark:text-gray-400">
                            Don't forget about the meeting tomorrow at 10am. We will be discussing
                            the progress of our current project and the plans for the next quarter.
                        </p>
                        <p className="mt-2 text-sm text-gray-600 dark:text-gray-400">Best,</p>
                        <p className="mt-2 text-sm text-gray-600 dark:text-gray-400">John Doe</p>
                    </div>
                </div>
            </div>
        </main>
    </div>
</template>
