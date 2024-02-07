<script setup lang="ts">

</script>

<template>
  <aside class="w-64 bg-white dark:bg-gray-800 border-r dark:border-gray-700">
    <div class="p-4">
      <h2 class="text-lg font-semibold text-gray-800 dark:text-gray-200">Mailbox</h2>
      <nav class="mt-8">
        <div class="mt-2 -mx-3">
          <Link
            class="flex justify-between items-center px-3 py-2 rounded-lg font-medium text-gray-700 dark:text-gray-200 bg-gray-200 dark:bg-gray-700"
            href="#"
          >
            <span>Inbox</span>
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
</template>

<style scoped>

</style>