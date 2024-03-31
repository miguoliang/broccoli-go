import {defineConfig} from 'vite'
import react from '@vitejs/plugin-react'

// https://vitejs.dev/config/
export default defineConfig({
    plugins: [react()],
    server: {
        proxy: {
            '/api': 'https://2e5x25uor2.execute-api.us-east-1.amazonaws.com/prod'
        }
    }
})
