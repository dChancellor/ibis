import { defineConfig } from "vite";
import { svelte } from "@sveltejs/vite-plugin-svelte";
import Icons from "unplugin-icons/vite";
import * as path from 'path';

export default defineConfig({
  plugins: [
    svelte(),
    Icons({
      compiler: "svelte",
    }),
  ],
  resolve: {
    alias: {
      '@components': path.resolve(__dirname, 'src/components'),
      '@sections': path.resolve(__dirname, 'src/sections'),
      '@stores': path.resolve(__dirname, 'src/stores'),
      '@helpers': path.resolve(__dirname, 'src/helpers'),
      '@assets': path.resolve(__dirname, 'src/assets'),
      '@typings': path.resolve(__dirname, 'src/typings'),
      '@wails': path.resolve(__dirname, 'wailsjs/go')
    }
  }
});
