<script lang="ts" setup>
import { onMounted, ref } from 'vue'
import { instance } from '@/lib/axios'
import type { RootObject, Source } from '@/types/api'
import EmailDetail from '@/components/self/EmailDetail.vue'
import EmailPreview from '@/components/self/EmailPreview.vue'
import SearchForm from '@/components/self/SearchForm.vue'
import LateralPanel from '@/components/self/LateralPanel.vue'
import 'v3-infinite-loading/lib/style.css'
import InfiniteLoading from 'v3-infinite-loading'
import { useRoute } from 'vue-router'

const data = ref<RootObject | null>(null)
const loading = ref(true)
const error = ref('')
let mail = ref<Source | null>(null)
const searchFilter = ref('')
let page = 1
const route = useRoute()

async function fetchData() {
  loading.value = true
  error.value = ''
  try {
    const response = await instance.get('/all/' + page)
    if (response.status !== 200) {
      throw new Error(`HTTP error! status: ${response.status}`)
    }
    if (!data.value) {
      data.value = response.data
    } else {
      data.value.hits.hits.push(...response.data.hits.hits)
    }
    page++
  } catch (e) {
    error.value = e.toString()
  } finally {
    loading.value = false
  }
}

async function searchData() {
  loading.value = true
  error.value = ''
  try {
    const response = await instance.post('/' + searchFilter.value)
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


function handleSearch(query: string) {
  searchFilter.value = query
  searchData()
}

function handleClic() {
  mail.value = data.value?.hits.hits.find((el) => {
    return el._id === route.params.mailID[0]
  })?._source
  console.log(mail.value)
}

onMounted(() => {
    fetchData()
  }
)

</script>

<template>
  <div class="flex h-screen bg-gray-100 dark:bg-gray-900">
    <LateralPanel />
    <main class="flex-1 overflow-y-auto">
      <SearchForm @search="handleSearch" />
      <div class="flex relative">
        <ul v-if="data !== null" class="w-1/2 px-4 py-3 space-y-2" @click="handleClic">
          <EmailPreview v-bind:emails="data?.hits.hits" />
        </ul>
        <EmailDetail :email="mail" />
      </div>
      <InfiniteLoading @infinite="fetchData" />

    </main>

  </div>

</template>
