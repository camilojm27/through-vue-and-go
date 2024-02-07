<script setup lang="ts">

</script>

<template>
  <div class="bg-white relative border rounded-lg">
    <table class="w-full text-sm"></table>
    <thead>
      <tr class="">
        <th class="px-4 py-3">ID</th>
        <th class="px-4 py-3">Email</th>
        <th class="px-4 py-3">Role</th>
        <th class="px-4 py-3">Status</th>
        <th class="px-4 py-3">
          <span class="sr-only">Actions</span>
        </th>
      </tr>
    </thead>
    <tbody>
    <tr v-for="item in items" :key="item.id" class="border-b">

    </tr>
    </tbody>
  </div>
</template>

<style scoped>

</style>
<!--text-left border-b-->