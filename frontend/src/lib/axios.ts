import axios from 'axios'

export const instance = axios.create({
// get the base url from the .env file
  baseURL: import.meta.env.VITE_BASE_URL,
  timeout: 5000,
  headers: { 'Content-Type': 'application/json', Accept: 'application/json' }
})
