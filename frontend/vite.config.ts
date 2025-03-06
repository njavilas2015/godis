import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react-swc'
import path from 'path';

export default defineConfig({
  plugins: [react()],
  resolve: {
    alias: {
      '@': path.resolve(__dirname, 'src'),
      '@desktop': path.resolve(__dirname, 'src', 'desktop'),
      '@mobile': path.resolve(__dirname, 'src', 'mobile'),
      '@store': path.resolve(__dirname, 'src', 'store'),
      '@components': path.resolve(__dirname, 'src', 'components'),
      '@hooks': path.resolve(__dirname, 'src', 'hooks'),
    },
  },
})
