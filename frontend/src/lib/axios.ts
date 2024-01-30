import axios from 'axios'

export const instance = axios.create({
  baseURL: 'http://localhost:8080/mail/',
  timeout: 1000,
  headers: { 'Content-Type': 'application/json', Accept: 'application/json' }
})
